# accelerando

Starter template that gets AI agents up to speed the moment you clone.

## Why the symlinks?

Claude Code, Codex, and Gemini CLI each look for their own instruction filename
(`CLAUDE.md`, `AGENTS.md`, `GEMINI.md`). This repo keeps one real file — `AGENTS.md` —
and symlinks the other two to it, so all three tools read identical instructions with no
duplication to keep in sync.

## Using this template

Clone it and work directly in it. Once it becomes a real project, detach it from this
template and push it as its own repo — see the "Detach / relaunch workflow" section in
[`AGENTS.md`](AGENTS.md) for the exact commands.
