# Recap — how this repo came to be

Read this first if you're a fresh AI session picking this repo up cold.

## Origin story

Maciek has an upcoming job interview with a live-coding segment: building some
integration with Ethereum, most likely in Go using `go-ethereum` (`ethclient`, etc.). AI
use is explicitly allowed during the interview, which is a newer interview format he
wasn't familiar with. Practice materials for that specific interview live in his other
repo, `experience/learning/eth/` — `TASK.md` (a practice task), `ai_rules/checklist.md`
(what AI-allowed interviews evaluate and how to behave in one), `ai_rules/sources.md`
(background reading), plus its own local `AGENTS.md`/symlinks. None of that ETH-specific
content is copied into `accelerando` — this repo is the generic template that grew out of
the process, not the practice project itself.

Research summary (from `ai_rules/checklist.md`, worth restating here since it shaped the
detach-workflow thinking): AI-allowed interviews grade problem decomposition *before*
touching AI, prompt precision, critical reading of AI output, and narration — not
memorized syntax or raw typing speed. First 5 minutes should be spent restating the
problem and asking clarifying questions before opening an editor or an AI tool at all.

## Why this repo exists

While setting up the ETH practice folder, the conversation noticed a recurring personal
convention worth templating: every project should ship with one real agent-instruction
file (`AGENTS.md`) and symlinks for the other filenames Claude Code / Codex / Gemini CLI
each look for (`CLAUDE.md`, `GEMINI.md`), so all three tools read identical instructions
without duplication. This is the same pattern already used in the `experience` repo (root
level, and in its `jobs/` subdirectory).

Rather than re-setting this up by hand on every new project, `accelerando` is meant to be
a clonable starting point: clone it, work in it, and once it becomes a real project, run
the detach workflow (documented in `AGENTS.md`) to rename it and push it to its own new
GitHub repo, disconnected from this template.

## The name

Playful, not arbitrary: *Accelerando* is a Charles Stross novel about AI agents, uploaded
minds, and post-singularity economics — and separately, an actual musical term meaning
"gradually speeding up." Both meanings fit: a starter template that gets a new project up
to speed the moment you clone it. A folder-level nickname `aineko` (the uploaded-cat
character in the novel who becomes an autonomous economic agent) was floated for any
future *non-tool-required* folder holding agent notes/prompts — explicitly **not** for
anything a tool actually depends on. Required paths (`.claude/skills/`, `AGENTS.md`,
`CLAUDE.md`, `GEMINI.md`) must keep their literal standard names or the tools that read
them won't find them.

## The `mac-plan` skill

`~/.claude/skills/mac-plan` (global, not yet copied into this repo) runs two independent
subagents — a Planner that writes a structured improvement proposal, and a Critic that
reviews it cold (no shared context with the Planner) — and saves the result as
`docs/improvements/IMP-NNN-slug.md`. Discussed as a possible future addition to this
template, but:
- It's built for durable, reviewable design docs on non-trivial changes (refactors,
  architecture decisions) — not for fast, narrated iteration on a scoped live-coding task.
  Spinning up two subagents mid-interview would burn time budget for no benefit there.
- It's inert by default: skills only run when invoked explicitly (`/mac-plan`) or when
  their description clearly matches the task, so having it present in `.claude/skills/`
  doesn't risk it firing unintentionally.

Not copied into `accelerando` yet — left as a "maybe later" in case a project started from
this template actually needs that kind of planning workflow.

## What's actually in this repo right now

- `README.md` — default, not yet rewritten
- `AGENTS.md` (real) + `CLAUDE.md`/`GEMINI.md` (symlinks) — the convention + detach
  workflow, written out concretely with shell commands
- This file

## Open threads / possible next steps

- Confirm with the interview's recruiter whether bringing personal tooling/config into
  their live-coding environment is even allowed — some run in a locked-down shared sandbox
  where none of this would be reachable anyway.
- Decide whether to eventually copy `mac-plan` in, and/or add other reusable skills.
- Rewrite `README.md` to actually describe the repo instead of the default GitHub stub.
- Nothing here has been committed yet — `AGENTS.md`, the two symlinks, and this file are
  untracked as of this recap.
