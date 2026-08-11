import importlib.util
import io
import json
import os
import subprocess
import tempfile
import unittest
from pathlib import Path
from unittest import mock


REPO_ROOT = Path(__file__).resolve().parents[1]
HOOK_PATH = REPO_ROOT / ".codex" / "hooks" / "verify_go_stop.py"
VERIFY_SCRIPT = REPO_ROOT / "scripts" / "verify-go"


spec = importlib.util.spec_from_file_location("verify_go_stop", HOOK_PATH)
verify_go_stop = importlib.util.module_from_spec(spec)
assert spec.loader is not None
spec.loader.exec_module(verify_go_stop)


class VerifyGoStopHookTests(unittest.TestCase):
    def run_hook(self, payload, *, checker_returncode=0, checker_output=""):
        fake_result = subprocess.CompletedProcess(
            args=[str(REPO_ROOT / "scripts" / "verify-go")],
            returncode=checker_returncode,
            stdout=checker_output,
        )
        stdout = io.StringIO()
        stderr = io.StringIO()
        with (
            mock.patch.object(verify_go_stop.subprocess, "check_output", return_value=str(REPO_ROOT)),
            mock.patch.object(verify_go_stop.subprocess, "run", return_value=fake_result),
            mock.patch.object(verify_go_stop.sys, "stdin", io.StringIO(json.dumps(payload))),
            mock.patch.object(verify_go_stop.sys, "stdout", stdout),
            mock.patch.object(verify_go_stop.sys, "stderr", stderr),
        ):
            return_code = verify_go_stop.main()
        return return_code, json.loads(stdout.getvalue()), stderr.getvalue()

    def test_success_allows_stop_and_forwards_checker_output_to_stderr(self):
        return_code, payload, stderr = self.run_hook(
            {"cwd": str(REPO_ROOT), "stop_hook_active": False},
            checker_output="all checks passed\n",
        )

        self.assertEqual(return_code, 0)
        self.assertEqual(payload, {})
        self.assertEqual(stderr, "all checks passed\n")

    def test_first_failure_requests_continuation(self):
        return_code, payload, _ = self.run_hook(
            {"cwd": str(REPO_ROOT), "stop_hook_active": False},
            checker_returncode=1,
            checker_output="go vet failed\n",
        )

        self.assertEqual(return_code, 0)
        self.assertEqual(payload["decision"], "block")
        self.assertIn("go vet failed", payload["reason"])

    def test_second_failure_stops_hook_loop(self):
        return_code, payload, _ = self.run_hook(
            {"cwd": str(REPO_ROOT), "stop_hook_active": True},
            checker_returncode=1,
            checker_output="go test failed\n",
        )

        self.assertEqual(return_code, 0)
        self.assertFalse(payload["continue"])
        self.assertIn("still fails", payload["stopReason"])
        self.assertIn("go test failed", payload["systemMessage"])

    def test_invalid_input_is_reported(self):
        stdout = io.StringIO()
        stderr = io.StringIO()
        with (
            mock.patch.object(verify_go_stop.sys, "stdin", io.StringIO("not json")),
            mock.patch.object(verify_go_stop.sys, "stdout", stdout),
            mock.patch.object(verify_go_stop.sys, "stderr", stderr),
        ):
            return_code = verify_go_stop.main()

        self.assertEqual(return_code, 1)
        self.assertIn("invalid input", stderr.getvalue())


class VerifyGoScriptTests(unittest.TestCase):
    def run_git(self, repo, *args):
        return subprocess.run(
            ["git", *args],
            cwd=repo,
            check=True,
            capture_output=True,
            text=True,
        )

    def test_checks_tracked_staged_unstaged_and_untracked_go_files(self):
        with tempfile.TemporaryDirectory() as temp_dir:
            repo = Path(temp_dir)
            self.run_git(repo, "init", "-q")
            self.run_git(repo, "config", "user.email", "tests@example.com")
            self.run_git(repo, "config", "user.name", "Hook tests")
            (repo / "tracked.go").write_text("package main\n")
            self.run_git(repo, "add", "tracked.go")
            self.run_git(repo, "commit", "-qm", "initial")

            (repo / "tracked.go").write_text("package main\n\n// unstaged\n")
            (repo / "staged.go").write_text("package main\n")
            self.run_git(repo, "add", "staged.go")
            (repo / "untracked.go").write_text("package main\n")
            (repo / ".gitignore").write_text("ignored.go\n")
            (repo / "ignored.go").write_text("package main\n")

            fake_bin = repo / "fake-bin"
            fake_bin.mkdir()
            (fake_bin / "gofmt").write_text(
                "#!/bin/sh\nprintf '%s\\n' \"$@\" > \"$FAKE_LOG/gofmt\"\n"
            )
            (fake_bin / "go").write_text(
                "#!/bin/sh\nprintf '%s\\n' \"$@\" >> \"$FAKE_LOG/go\"\n"
            )
            (fake_bin / "gofmt").chmod(0o755)
            (fake_bin / "go").chmod(0o755)
            log_dir = repo / "logs"
            log_dir.mkdir()

            env = os.environ.copy()
            env["PATH"] = f"{fake_bin}{os.pathsep}{env['PATH']}"
            env["FAKE_LOG"] = str(log_dir)
            result = subprocess.run(
                [str(VERIFY_SCRIPT)],
                cwd=repo,
                env=env,
                check=False,
                capture_output=True,
                text=True,
            )

            self.assertEqual(result.returncode, 0, result.stdout + result.stderr)
            formatted = (log_dir / "gofmt").read_text().splitlines()
            self.assertEqual(formatted[0], "-w")
            self.assertEqual(
                set(formatted[1:]),
                {"tracked.go", "staged.go", "untracked.go"},
            )
            self.assertNotIn("ignored.go", formatted)
            self.assertEqual(
                (log_dir / "go").read_text().splitlines(),
                ["vet", "./...", "test", "./...", "test", "-race", "./..."],
            )


if __name__ == "__main__":
    unittest.main()
