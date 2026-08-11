#!/usr/bin/env python3

import json
import subprocess
import sys
from pathlib import Path
from typing import Any


MAX_OUTPUT_CHARS = 12_000


def emit(payload: dict[str, Any]) -> None:
    json.dump(payload, sys.stdout)
    sys.stdout.write("\n")


def shorten(output: str) -> str:
    output = output.strip()
    if len(output) <= MAX_OUTPUT_CHARS:
        return output

    half = MAX_OUTPUT_CHARS // 2
    return f"{output[:half]}\n\n... output truncated ...\n\n{output[-half:]}"


def failure_reason(output: str) -> str:
    details = shorten(output)
    message = (
        "Repository Go verification failed. Fix the failures and run "
        "./scripts/verify-go again before finishing."
    )
    if details:
        return f"{message}\n\n{details}"
    return message


def main() -> int:
    try:
        hook_input = json.load(sys.stdin)
    except (json.JSONDecodeError, TypeError) as error:
        print(f"verify-go hook: invalid input: {error}", file=sys.stderr)
        return 1

    cwd = hook_input.get("cwd")
    try:
        repo_root = Path(
            subprocess.check_output(
                ["git", "rev-parse", "--show-toplevel"],
                cwd=cwd,
                text=True,
                stderr=subprocess.STDOUT,
            ).strip()
        )
    except (OSError, subprocess.CalledProcessError) as error:
        print(f"verify-go hook: cannot locate repository root: {error}", file=sys.stderr)
        return 1

    checker = repo_root / "scripts" / "verify-go"
    try:
        result = subprocess.run(
            [str(checker)],
            cwd=repo_root,
            text=True,
            stdout=subprocess.PIPE,
            stderr=subprocess.STDOUT,
            check=False,
        )
    except OSError as error:
        result_output = f"Could not run {checker}: {error}"
        result_code = 1
    else:
        result_output = result.stdout
        result_code = result.returncode

    if result_code == 0:
        if result_output:
            print(result_output, file=sys.stderr, end="")
        emit({})
        return 0

    reason = failure_reason(result_output)
    if not hook_input.get("stop_hook_active", False):
        emit({"decision": "block", "reason": reason})
        return 0

    emit(
        {
            "continue": False,
            "stopReason": "Go verification still fails after automatic continuation.",
            "systemMessage": reason,
        }
    )
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
