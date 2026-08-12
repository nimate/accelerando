# accelerando

Generic, language-agnostic starter template for AI-assisted projects. Clone it, then
either work directly in it or detach it into its own project (see below). It carries two
things by default: an agent-instruction-file convention, and a detach/relaunch workflow.

## Agent-instruction-file convention

`AGENTS.md` holds the shared instructions and is what Codex reads directly. `CLAUDE.md`
and `GEMINI.md` are real files, not symlinks: each starts with an `@AGENTS.md` import
line, followed by any tool-specific instructions. Claude Code and Gemini CLI both
resolve `@path/to/file` imports, inlining the target's content at load time — so the
import line pulls in everything in `AGENTS.md` plus whatever follows it locally.

Edit shared content in `AGENTS.md`. Edit Claude-only or Gemini-only instructions
directly in `CLAUDE.md` / `GEMINI.md`, below the import line.

When adding this convention to a new subdirectory (e.g. a subproject with its own
context), repeat the same pattern: an `AGENTS.md` with the shared content, and
`CLAUDE.md`/`GEMINI.md` each starting with `@AGENTS.md` plus their own additions. Don't
overwrite an existing real instruction file — if a directory already has tool-specific
content, preserve it and just add the `@AGENTS.md` import line at the top if it's
missing.

## Detach / relaunch workflow

`nimate/accelerando` is registered as a GitHub template repository. Starting a new
project from it never touches this clone — it always creates a fresh repo + fresh local
folder, seeded with this repo's files but with its own commit history (no template
history baggage):

```sh
gh repo create <new-project-name> --private --template nimate/accelerando --clone
cd <new-project-name>
```

This `accelerando` folder itself stays put as the reusable template — its `origin` never
changes, nothing gets renamed or removed. Don't run the old remove-origin/mv dance here;
that workflow is retired in favor of `--template`.

If you don't use `gh`: click "Use this template" on the repo's GitHub page instead, which
does the same thing through the web UI.

In the new project's folder, update `AGENTS.md`'s title/purpose section to describe the
actual project instead of the template.

## Skills

If a skill like `mac-plan` (planner + critic subagents for design docs) ends up copied
into `.claude/skills/` here, it stays inert unless explicitly invoked with `/mac-plan` or
its description clearly matches the task at hand — its mere presence in the repo doesn't
trigger anything automatically.

See `RECAP.md` for the fuller context behind why this repo exists.

## Local agent safety policy

Agents may create files and directories, edit files, and run local tests, `make`,
and other project tooling without asking for routine confirmation.

Before deleting an existing file or directory, agents must ask the user for
confirmation. This includes `rm`, `rmdir`, `git clean`, `git reset`, `git
checkout`, and `git restore` when they would remove or overwrite existing work.
