# Concepts

Nouns map atomically to file sets (EXTRACTED); verbs aggregate structural edges (INFERRED).

- `con` | files=3 | mentions=5 | `app.py`, `crypto.go`, `handlers.go`
- `load` | files=2 | mentions=12 | `app.py`, `web/js/dashboard.js`
- `log` | files=2 | mentions=8 | `app.py`, `handlers.go`
- `command` | files=2 | mentions=6 | `app.py`, `handlers.go`
- `get` | files=2 | mentions=6 | `crypto.go`, `handlers.go`
- `clients` | files=2 | mentions=5 | `app.py`, `handlers.go`
- `crea` | files=2 | mentions=5 | `app.py`, `schemas.go`
- `datos` | files=2 | mentions=4 | `app.py`, `crypto.go`
- `desde` | files=2 | mentions=4 | `app.py`, `crypto.go`
- `handler` | files=2 | mentions=4 | `app.py`, `handlers.go`
- `open` | files=2 | mentions=4 | `app.py`, `web/js/dashboard.js`
- `compatible` | files=2 | mentions=3 | `crypto.go`, `handlers.go`
- `login` | files=2 | mentions=3 | `app.py`, `handlers.go`
- `logs` | files=2 | mentions=3 | `app.py`, `handlers.go`
- `usando` | files=2 | mentions=3 | `crypto.go`, `schemas.go`
- `funci` | files=2 | mentions=2 | `app.py`, `handlers.go`
- `implants` | files=2 | mentions=2 | `app.py`, `web/js/dashboard.js`
- `list` | files=2 | mentions=2 | `handlers.go`, `web/js/dashboard.js`
- `nueva` | files=2 | mentions=2 | `app.py`, `handlers.go`
- `python` | files=2 | mentions=2 | `crypto.go`, `handlers.go`
- `upload` | files=2 | mentions=2 | `app.py`, `handlers.go`

## Dialectic

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
