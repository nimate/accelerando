# accelerando

Generic, language-agnostic starter template for AI-assisted projects. Clone it, then
either work directly in it or detach it into its own project (see below). It carries two
things by default: an agent-instruction-file convention, and a detach/relaunch workflow.

## Agent-instruction-file convention

`CLAUDE.md` and `GEMINI.md` in this repo are symlinks to this file (`AGENTS.md`). Edit
`AGENTS.md`; never edit the symlinks directly — they exist purely so Claude Code, Codex,
and Gemini CLI each find their expected filename with identical content, no duplication.

When adding this convention to a new subdirectory (e.g. a subproject with its own
context), repeat the same pattern there: pick one real file, symlink the other two agent
filenames to it. Don't overwrite an existing real instruction file with a symlink — if a
directory already has a specific one, preserve it and point the missing filenames at it.

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
