# Second Brain

*Last synthesized: 2026-10-07 | 7 files | 1 concept pages | offline, zero tokens*

> Raw sources -> readmenator wiki -> links (Karpathy LLM Wiki Pattern, deterministic).
> Start here, then open one community page. Prefer grep over full reads.

## Vault Overview

The codebase centres on `app.py`, `handlers.go`, `dashboard.js`. Architecturally it is 2 layers, dominant utility (6 files) across 1 import-based communities. Recorded risk surface: 0 security findings and 0 dependency cycles.

Communities are self-contained in the resolved import graph; no cross-boundary bridges were recorded.

Open work clusters around documentation (14% file coverage), 0 security findings, 1 taint paths, and 5 suggested exploration questions in `queries.md`.

## Stats

| Metric | Value |
|--------|-------|
| Files | 7 |
| Symbols | 94 |
| Resolved imports | 0 |
| Languages | go, js, py, sh |
| Communities | 1 |
| Doc coverage | 14% (1/7 files) |
| Security findings | 0 |
| Estimated read cost | ~1306 tokens (chars/4, offline so $0) |

## Reading Order

1. Skim Stats and God Nodes below for blast radius.
2. Open the largest community page first, then follow Connections.
3. Use `queries.md` for the next question; log the answer there.

```
grep -rn '<keyword>' index.md community_*.md
readmenator query "<question>" --target readmenator_BlackObsidianC2_syhiohg_
```

## Concept Wiki

- [root (7 files, cohesion 1.00)](./community_0_root.md)

## God Nodes

| File | Score |
|------|-------|
| `app.py` | 7.2 |
| `handlers.go` | 1.3 |
| `web/js/dashboard.js` | 0.4 |
| `crypto.go` | 0.3 |
| `main.go` | 0.1 |

## Strongest Connections

- No cross-community connections recorded.

## Navigation Tips

- Obsidian Graph View works: every community page links back here.
- `connections.json` is machine-readable for GraphRAG pipelines.
- `REPORT.md` states what was extracted vs inferred and current limits.
- Regenerate offline: `readmenator . --rebuild` (no network, no tokens).
