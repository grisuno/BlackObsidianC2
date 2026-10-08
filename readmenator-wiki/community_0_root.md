# root

*Community 0 | 7 files | cohesion 1.00*

## Definition

This community groups 7 file(s) rooted at `root` with dominant language go (cohesion 1.00). Central symbols: `AESDecrypt`, `AESEncrypt`, `GetAESKey`, `HandleDownload`, `HandleGetCommand`, `HandleGetConnectedClients`, `HandleGetResults`, `HandleIssueCommand`. Core file: `app.py` (72 symbols). Documented purpose: Autor: Gris Iscomeback Correo electrónico: grisiscomeback[at]gmail[dot]com Fecha de creación: xx/xx/xxxx Licencia: GPL v3  Descripción:.

## Files

| File | Language | Layer | Symbols | Doc |
|------|----------|-------|---------|-----|
| `app.py` | py | utility | 72 | yes |
| `crypto.go` | go | utility | 3 | no |
| `handlers.go` | go | presentation | 13 | no |
| `install.sh` | sh | utility | 0 | no |
| `main.go` | go | utility | 1 | no |
| `schemas.go` | go | utility | 1 | no |
| `web/js/dashboard.js` | js | utility | 4 | no |

## Key Symbols

- `setup_modern_theme` (function, `app.py:73`) `def setup_modern_theme()`
- `StatusBar` (class, `app.py:172`) `class StatusBar(Frame)`
- `__init__` (method, `app.py:173`) `def __init__(self, parent)`
- `update_status` (method, `app.py:188`) `def update_status(self, text, connected)`
- `update_time` (method, `app.py:194`) `def update_time(self)`
- `ModernTreeview` (class, `app.py:199`) `class ModernTreeview(Frame)`
- `__init__` (method, `app.py:200`) `def __init__(self, parent, columns, data_loader)`
- `refresh_data` (method, `app.py:241`) `def refresh_data(self)`
- `set_title` (method, `app.py:253`) `def set_title(self, title)`
- `ModernConsole` (class, `app.py:256`) `class ModernConsole(Frame)`
- `__init__` (method, `app.py:257`) `def __init__(self, parent, client_id)`
- `load_command_history` (method, `app.py:310`) `def load_command_history(self)` - Carga el historial de comandos desde el archivo .log del cliente
- `send_command` (method, `app.py:333`) `def send_command(self, client_id)`
- `navigate_history_up` (method, `app.py:358`) `def navigate_history_up(self, event)` - Navegar hacia arriba en el historial de comandos
- `navigate_history_down` (method, `app.py:372`) `def navigate_history_down(self, event)` - Navegar hacia abajo en el historial de comandos
- `add_text` (method, `app.py:387`) `def add_text(self, text, tag)`
- `ImplantCard` (class, `app.py:391`) `class ImplantCard(Frame)`
- `__init__` (method, `app.py:392`) `def __init__(self, parent, client_id, on_select)`
- `load_latest_client_info` (method, `app.py:489`) `def load_latest_client_info(self)` - Carga la información más reciente del cliente desde su archivo .log Y su archivo .json de configurac
- `update_status_indicator` (method, `app.py:539`) `def update_status_indicator(self)` - Actualiza el indicador de estado basado en la última actividad
- `on_card_click` (method, `app.py:554`) `def on_card_click(self, event)`
- `open_console` (method, `app.py:558`) `def open_console(self)`
- `open_files` (method, `app.py:562`) `def open_files(self)`
- `open_processes` (method, `app.py:566`) `def open_processes(self)`
- `load_os_image` (method, `app.py:570`) `def load_os_image(self, client_id)` - Cargar imagen del sistema operativo basado en el client_id
- `toggle_view` (method, `app.py:601`) `def toggle_view(self)`
- `show_table_view` (method, `app.py:611`) `def show_table_view(self)`
- `show_card_view` (method, `app.py:656`) `def show_card_view(self)`
- `login` (method, `app.py:662`) `def login()`
- `show_notification` (method, `app.py:683`) `def show_notification(message, type)` - Mostrar notificación en el sistema

## Internal vs External Edges

- Internal resolved imports (EXTRACTED): 0
- Cross-boundary resolved imports (EXTRACTED): 0

## Connections

- No cross-community bridges recorded. This community is self-contained.

## Risks

- [taint medium] `app.py` -> `app.py` via `requests` (0 hops)
- [dataflow UNCHECKED_ALLOC] `app.py:590` `load_os_image` `img`: Result of allocator stored in `img` is never checked against NULL.

## Open Questions

- Why do 6 file(s) lack file-level docs (e.g. `crypto.go`)? What purpose do they serve?
- Is the dangerous import `requests` in `app.py` still required, or can it be isolated?
- What would break if the most connected file in root changed?
- Should root be split, given cohesion 1.00?

## Sources

- `app.py`
- `crypto.go`
- `handlers.go`
- `install.sh`
- `main.go`
- `schemas.go`
- `web/js/dashboard.js`
