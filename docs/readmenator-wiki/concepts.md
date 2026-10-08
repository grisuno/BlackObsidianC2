# Concepts

Second-brain semantic layer: nouns map atomically to file sets (EXTRACTED); verbs aggregate structural edges (INFERRED).

| Concept | Files | Mentions | Top Files |
|---------|-------|----------|-----------|
| `con` | 3 | 5 | `app.py`, `crypto.go`, `handlers.go` |
| `load` | 2 | 12 | `app.py`, `web/js/dashboard.js` |
| `log` | 2 | 8 | `app.py`, `handlers.go` |
| `command` | 2 | 6 | `app.py`, `handlers.go` |
| `get` | 2 | 6 | `crypto.go`, `handlers.go` |
| `clients` | 2 | 5 | `app.py`, `handlers.go` |
| `crea` | 2 | 5 | `app.py`, `schemas.go` |
| `datos` | 2 | 4 | `app.py`, `crypto.go` |
| `desde` | 2 | 4 | `app.py`, `crypto.go` |
| `handler` | 2 | 4 | `app.py`, `handlers.go` |
| `open` | 2 | 4 | `app.py`, `web/js/dashboard.js` |
| `compatible` | 2 | 3 | `crypto.go`, `handlers.go` |
| `login` | 2 | 3 | `app.py`, `handlers.go` |
| `logs` | 2 | 3 | `app.py`, `handlers.go` |
| `usando` | 2 | 3 | `crypto.go`, `schemas.go` |
| `funci` | 2 | 2 | `app.py`, `handlers.go` |
| `implants` | 2 | 2 | `app.py`, `web/js/dashboard.js` |
| `list` | 2 | 2 | `handlers.go`, `web/js/dashboard.js` |
| `nueva` | 2 | 2 | `app.py`, `handlers.go` |
| `python` | 2 | 2 | `crypto.go`, `handlers.go` |
| `upload` | 2 | 2 | `app.py`, `handlers.go` |

## Dialectic Prompts

- Thesis: `clients` centralizes 2 files; Antithesis: `command` pulls 2 files with 2 shared (Jaccard 1.00); Synthesis: should they merge, split by layer, or keep `bridges` explicit?
- Thesis: `clients` centralizes 2 files; Antithesis: `con` pulls 3 files with 2 shared (Jaccard 0.67); Synthesis: should they merge, split by layer, or keep `bridges` explicit?
- Thesis: `clients` centralizes 2 files; Antithesis: `funci` pulls 2 files with 2 shared (Jaccard 1.00); Synthesis: should they merge, split by layer, or keep `bridges` explicit?
- Thesis: `clients` centralizes 2 files; Antithesis: `handler` pulls 2 files with 2 shared (Jaccard 1.00); Synthesis: should they merge, split by layer, or keep `bridges` explicit?
- Thesis: `clients` centralizes 2 files; Antithesis: `log` pulls 2 files with 2 shared (Jaccard 1.00); Synthesis: should they merge, split by layer, or keep `bridges` explicit?
- Thesis: `clients` centralizes 2 files; Antithesis: `login` pulls 2 files with 2 shared (Jaccard 1.00); Synthesis: should they merge, split by layer, or keep `bridges` explicit?
- Thesis: `clients` centralizes 2 files; Antithesis: `logs` pulls 2 files with 2 shared (Jaccard 1.00); Synthesis: should they merge, split by layer, or keep `bridges` explicit?
- Thesis: `clients` centralizes 2 files; Antithesis: `nueva` pulls 2 files with 2 shared (Jaccard 1.00); Synthesis: should they merge, split by layer, or keep `bridges` explicit?
- Thesis: `clients` centralizes 2 files; Antithesis: `upload` pulls 2 files with 2 shared (Jaccard 1.00); Synthesis: should they merge, split by layer, or keep `bridges` explicit?
- Thesis: `command` centralizes 2 files; Antithesis: `con` pulls 3 files with 2 shared (Jaccard 0.67); Synthesis: should they merge, split by layer, or keep `bridges` explicit?
