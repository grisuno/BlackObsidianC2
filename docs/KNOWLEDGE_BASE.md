# Polyglot Codebase Knowledge Graph

> Generated offline by **readmenator**. Supports C, C++, Python, Go, Rust, JS/TS, Java, C#, Shell, PHP, Dart, GDScript, Nim, ASM.
> No LLMs. No tokens. Pure static analysis. See more [here](https://github.com/grisuno/ReadMenator)

**Total Files Parsed:** 7 | **Total Symbols Extracted:** 94 | **Total Imports:** 48

## Structural Knowledge Map
```mermaid
graph TD
    classDef mod fill:#1e1e1e,stroke:#ff6666,stroke-width:2px,color:#fff;
    classDef cls fill:#2d2d2d,stroke:#4ec9b0,stroke-width:2px,color:#fff;
    classDef fn fill:#333,stroke:#dcdcaa,stroke-width:1px,color:#dcdcaa;
    classDef ext fill:#111,stroke:#666,stroke-dasharray:5 5,color:#aaa;
    app_py["app.py (py)"]
    class app_py mod;
    app_py_setup_modern_theme["setup_modern_theme"]
    class app_py_setup_modern_theme fn;
    app_py --> app_py_setup_modern_theme
    app_py_StatusBar["StatusBar"]
    class app_py_StatusBar cls;
    app_py --> app_py_StatusBar
    app_py_ModernTreeview["ModernTreeview"]
    class app_py_ModernTreeview cls;
    app_py --> app_py_ModernTreeview
    app_py_ModernConsole["ModernConsole"]
    class app_py_ModernConsole cls;
    app_py --> app_py_ModernConsole
    app_py_ImplantCard["ImplantCard"]
    class app_py_ImplantCard cls;
    app_py --> app_py_ImplantCard
    handlers_go["handlers.go (go)"]
    class handlers_go mod;
    handlers_go_RegisterC2Routes["RegisterC2Routes"]
    class handlers_go_RegisterC2Routes fn;
    handlers_go --> handlers_go_RegisterC2Routes
    handlers_go_HandleLogin["HandleLogin"]
    class handlers_go_HandleLogin fn;
    handlers_go --> handlers_go_HandleLogin
    handlers_go_HandleGetConnectedClients["HandleGetConnectedClients"]
    class handlers_go_HandleGetConnectedClients fn;
    handlers_go --> handlers_go_HandleGetConnectedClients
    handlers_go_HandleIssueCommandLegacy["HandleIssueCommandLegacy"]
    class handlers_go_HandleIssueCommandLegacy fn;
    handlers_go --> handlers_go_HandleIssueCommandLegacy
    handlers_go_HandleGetCommand["HandleGetCommand"]
    class handlers_go_HandleGetCommand fn;
    handlers_go --> handlers_go_HandleGetCommand
    crypto_go["crypto.go (go)"]
    class crypto_go mod;
    crypto_go_GetAESKey["GetAESKey"]
    class crypto_go_GetAESKey fn;
    crypto_go --> crypto_go_GetAESKey
    crypto_go_AESEncrypt["AESEncrypt"]
    class crypto_go_AESEncrypt fn;
    crypto_go --> crypto_go_AESEncrypt
    crypto_go_AESDecrypt["AESDecrypt"]
    class crypto_go_AESDecrypt fn;
    crypto_go --> crypto_go_AESDecrypt
    main_go["main.go (go)"]
    class main_go mod;
    main_go_main["main"]
    class main_go_main fn;
    main_go --> main_go_main
    schemas_go["schemas.go (go)"]
    class schemas_go mod;
    schemas_go_InitializeCollections["InitializeCollections"]
    class schemas_go_InitializeCollections fn;
    schemas_go --> schemas_go_InitializeCollections
    web_js_dashboard_js["dashboard.js (js)"]
    class web_js_dashboard_js mod;
    web_js_dashboard_js_loadDashboard["loadDashboard"]
    class web_js_dashboard_js_loadDashboard fn;
    web_js_dashboard_js --> web_js_dashboard_js_loadDashboard
    web_js_dashboard_js_loadImplantsList["loadImplantsList"]
    class web_js_dashboard_js_loadImplantsList fn;
    web_js_dashboard_js --> web_js_dashboard_js_loadImplantsList
    web_js_dashboard_js_openTerminal["openTerminal"]
    class web_js_dashboard_js_openTerminal fn;
    web_js_dashboard_js --> web_js_dashboard_js_openTerminal
    web_js_dashboard_js_logout["logout"]
    class web_js_dashboard_js_logout fn;
    web_js_dashboard_js --> web_js_dashboard_js_logout
    install_sh["install.sh (sh)"]
    class install_sh mod;
    ext_tkinter["tkinter"]
    class ext_tkinter ext;
    app_py -.->|imports| ext_tkinter
    app_py -.->|imports| ext_tkinter
    ext_requests["requests"]
    class ext_requests ext;
    app_py -.->|imports| ext_requests
    ext_os["os"]
    class ext_os ext;
    app_py -.->|imports| ext_os
    ext_csv["csv"]
    class ext_csv ext;
    app_py -.->|imports| ext_csv
    ext_time["time"]
    class ext_time ext;
    app_py -.->|imports| ext_time
    ext_threading["threading"]
    class ext_threading ext;
    app_py -.->|imports| ext_threading
    ext_json["json"]
    class ext_json ext;
    app_py -.->|imports| ext_json
    ext_queue["queue"]
    class ext_queue ext;
    app_py -.->|imports| ext_queue
    ext_re["re"]
    class ext_re ext;
    app_py -.->|imports| ext_re
    ext_datetime["datetime"]
    class ext_datetime ext;
    app_py -.->|imports| ext_datetime
    ext_PIL["PIL"]
    class ext_PIL ext;
    app_py -.->|imports| ext_PIL
    ext_watchdog_observers["watchdog.observers"]
    class ext_watchdog_observers ext;
    app_py -.->|imports| ext_watchdog_observers
    ext_watchdog_events["watchdog.events"]
    class ext_watchdog_events ext;
    app_py -.->|imports| ext_watchdog_events
    app_py -.->|imports| ext_PIL
    app_py -.->|imports| ext_watchdog_observers
    app_py -.->|imports| ext_watchdog_events
    ext_crypto_aes["aes"]
    class ext_crypto_aes ext;
    crypto_go -.->|imports| ext_crypto_aes
    ext_crypto_cipher["cipher"]
    class ext_crypto_cipher ext;
    crypto_go -.->|imports| ext_crypto_cipher
    ext_crypto_rand["rand"]
    class ext_crypto_rand ext;
    crypto_go -.->|imports| ext_crypto_rand
    ext_encoding_base64["base64"]
    class ext_encoding_base64 ext;
    crypto_go -.->|imports| ext_encoding_base64
    ext_encoding_hex["hex"]
    class ext_encoding_hex ext;
    crypto_go -.->|imports| ext_encoding_hex
    ext_errors["errors"]
    class ext_errors ext;
    crypto_go -.->|imports| ext_errors
    ext_io["io"]
    class ext_io ext;
    crypto_go -.->|imports| ext_io
    crypto_go -.->|imports| ext_os
    ext_encoding_json["json"]
    class ext_encoding_json ext;
    handlers_go -.->|imports| ext_encoding_json
    ext_encoding_csv["csv"]
    class ext_encoding_csv ext;
    handlers_go -.->|imports| ext_encoding_csv
    ext_fmt["fmt"]
    class ext_fmt ext;
    handlers_go -.->|imports| ext_fmt
    handlers_go -.->|imports| ext_io
    handlers_go -.->|imports| ext_os
    ext_net_http["http"]
    class ext_net_http ext;
    handlers_go -.->|imports| ext_net_http
    ext_path_filepath["filepath"]
    class ext_path_filepath ext;
    handlers_go -.->|imports| ext_path_filepath
    ext_regexp["regexp"]
    class ext_regexp ext;
    handlers_go -.->|imports| ext_regexp
    ext_strings["strings"]
    class ext_strings ext;
    handlers_go -.->|imports| ext_strings
    handlers_go -.->|imports| ext_time
    ext_github_com_pocketbase_dbx["dbx"]
    class ext_github_com_pocketbase_dbx ext;
    handlers_go -.->|imports| ext_github_com_pocketbase_dbx
    ext_github_com_pocketbase_pocketbase["pocketbase"]
    class ext_github_com_pocketbase_pocketbase ext;
    handlers_go -.->|imports| ext_github_com_pocketbase_pocketbase
    ext_github_com_pocketbase_pocketbase_apis["apis"]
    class ext_github_com_pocketbase_pocketbase_apis ext;
    handlers_go -.->|imports| ext_github_com_pocketbase_pocketbase_apis
    ext_github_com_pocketbase_pocketbase_core["core"]
    class ext_github_com_pocketbase_pocketbase_core ext;
    handlers_go -.->|imports| ext_github_com_pocketbase_pocketbase_core
    ext_crypto_tls["tls"]
    class ext_crypto_tls ext;
    main_go -.->|imports| ext_crypto_tls
    ext_log["log"]
    class ext_log ext;
    main_go -.->|imports| ext_log
    main_go -.->|imports| ext_os
    main_go -.->|imports| ext_path_filepath
    main_go -.->|imports| ext_net_http
    main_go -.->|imports| ext_github_com_pocketbase_pocketbase
    main_go -.->|imports| ext_github_com_pocketbase_pocketbase_core
    schemas_go -.->|imports| ext_github_com_pocketbase_pocketbase
    schemas_go -.->|imports| ext_github_com_pocketbase_pocketbase_core
```

---

## Architecture Reference

### GO (4 files)

#### `crypto.go`
**Path:** `crypto.go`

**Functions:**
- `GetAESKey` (line 21) `func GetAESKey(` - *GetAESKey obtiene la clave AES desde variables de entorno con validación*
- `AESEncrypt` (line 40) `func AESEncrypt(` - *AESEncrypt cifra datos usando AES-256-CFB (compatible con Python)*
- `AESDecrypt` (line 71) `func AESDecrypt(` - *AESDecrypt descifra datos usando AES-256-CFB*

#### `handlers.go`
**Path:** `handlers.go`

**Functions:**
- `RegisterC2Routes` (line 30) `func RegisterC2Routes(`
- `HandleLogin` (line 83) `func HandleLogin(` - *===== HANDLER: Login Compatible con GUI Python =====*
- `HandleGetConnectedClients` (line 125) `func HandleGetConnectedClients(` - *===== HANDLER: Get Connected Clients (Compatible con GUI) =====*
- `HandleIssueCommandLegacy` (line 154) `func HandleIssueCommandLegacy(` - *===== HANDLER: Issue Command Legacy (sin autenticación) =====*
- `HandleGetCommand` (line 195) `func HandleGetCommand(`
- `HandleUpload` (line 267) `func HandleUpload(`
- `HandleDownload` (line 308) `func HandleDownload(`
- `HandleIssueCommand` (line 343) `func HandleIssueCommand(`
- `HandlePostResult` (line 383) `func HandlePostResult(`
- `writeLogCSV` (line 526) `func writeLogCSV(` - *✅ FUNCIÓN NUEVA: Escribir logs en CSV*
- `HandleListClients` (line 595) `func HandleListClients(`
- `HandleGetResults` (line 626) `func HandleGetResults(`
- `truncate` (line 661) `func truncate(`

#### `main.go`
**Path:** `main.go`

**Functions:**
- `main` (line 13) `func main(`

#### `schemas.go`
**Path:** `schemas.go`

**Functions:**
- `InitializeCollections` (line 10) `func InitializeCollections(` - *InitializeCollections crea las colecciones necesarias si no existen Usando la API de PocketBase v0.31.0*

### JS (1 files)

#### `dashboard.js`
**Path:** `web/js/dashboard.js`

**Functions:**
- `loadDashboard` (line 10)
- `loadImplantsList` (line 36)
- `openTerminal` (line 61)
- `logout` (line 65)

### PY (1 files)

#### `app.py`
**Path:** `app.py`

**Classes:**
- `StatusBar` (line 172) `class StatusBar`
- `ModernTreeview` (line 199) `class ModernTreeview`
- `ModernConsole` (line 256) `class ModernConsole`
- `ImplantCard` (line 391) `class ImplantCard`
- `LogHandler` (line 1115) `class LogHandler(FileSystemEventHandler)`

**Functions:**
- `setup_modern_theme` (line 73) `def setup_modern_theme()`
- `login` (line 662) `def login()`
- `show_notification` (line 683) `def show_notification(message, type)` - *Mostrar notificación en el sistema*
- `refresh_clients` (line 699) `def refresh_clients()`
- `select_client` (line 718) `def select_client(client_id)`
- `create_beacon_tab` (line 729) `def create_beacon_tab(client_id)` - *Crea una nueva pestaña para un beacon, con subpestañas de Consola e Intel.*
- `load_implant_config` (line 756) `def load_implant_config(client_id)` - *Carga la configuración del implant desde el archivo JSON.*
- `create_intel_tab` (line 769) `def create_intel_tab(parent, client_id)` - *Crea la pestaña de Inteligencia/Recon para un beacon específico.*
- `create_section_header` (line 850) `def create_section_header(parent, title)` - *Crea un encabezado de sección para la pestaña de Intel.*
- `load_intel_data` (line 857) `def load_intel_data(client_id)` - *Extrae y parsea datos de inteligencia del archivo .log del cliente.*
- `is_client_active` (line 922) `def is_client_active(client_id)`
- `load_latest_client_info_dict` (line 928) `def load_latest_client_info_dict(client_id)`
- `refresh_table_view` (line 963) `def refresh_table_view(tree)`
- `create_table_view_frame` (line 987) `def create_table_view_frame(parent)`
- `toggle_global_view` (line 1021) `def toggle_global_view()`
- `create_section_header` (line 1042) `def create_section_header(parent, title)` - *Crea un encabezado de sección para la pestaña de Intel.*
- `load_intel_data` (line 1049) `def load_intel_data(client_id)` - *Extrae y parsea datos de inteligencia del archivo .log del cliente.*
- `process_queue` (line 1163) `def process_queue()`
- `start_polling` (line 1181) `def start_polling()`
- `stop_polling` (line 1190) `def stop_polling()`
- `auto_refresh_clients` (line 1197) `def auto_refresh_clients()`
- `load_implants_data` (line 1206) `def load_implants_data()`
- `load_banners_data` (line 1218) `def load_banners_data()`
- `load_access_log` (line 1226) `def load_access_log()`
- `upload_file` (line 1244) `def upload_file()`
- `create_modern_ui` (line 1265) `def create_modern_ui()`
- `create_tools_grid` (line 1427) `def create_tools_grid(parent)` - *Crear grid de herramientas*
- `create_tool_card` (line 1479) `def create_tool_card(parent, tool)` - *Crear card de herramienta*
- `create_modern_menu` (line 1503) `def create_modern_menu()` - *Crear menú moderno*
- `setup_keyboard_shortcuts` (line 1541) `def setup_keyboard_shortcuts()` - *Configurar atajos de teclado*
- `export_logs` (line 1551) `def export_logs()` - *Exportar logs a archivo*
- `manage_payloads` (line 1586) `def manage_payloads()` - *Ventana de gestión de payloads*
- `show_statistics` (line 1626) `def show_statistics()` - *Mostrar ventana de estadísticas*
- `show_settings` (line 1668) `def show_settings()` - *Mostrar ventana de configuración*
- `show_help` (line 1721) `def show_help()` - *Mostrar ayuda*
- `show_about` (line 1792) `def show_about()` - *Mostrar información del programa*
- `on_closing` (line 1804) `def on_closing()` - *Manejar cierre de aplicación*
- `main` (line 1811) `def main()` - *Función principal*
- `__init__` (line 173) `def __init__(self, parent)`
- `update_status` (line 188) `def update_status(self, text, connected)`
- `update_time` (line 194) `def update_time(self)`
- `__init__` (line 200) `def __init__(self, parent, columns, data_loader)`
- `refresh_data` (line 241) `def refresh_data(self)`
- `set_title` (line 253) `def set_title(self, title)`
- `__init__` (line 257) `def __init__(self, parent, client_id)`
- `load_command_history` (line 310) `def load_command_history(self)` - *Carga el historial de comandos desde el archivo .log del cliente*
- `send_command` (line 333) `def send_command(self, client_id)`
- `navigate_history_up` (line 358) `def navigate_history_up(self, event)` - *Navegar hacia arriba en el historial de comandos*
- `navigate_history_down` (line 372) `def navigate_history_down(self, event)` - *Navegar hacia abajo en el historial de comandos*
- `add_text` (line 387) `def add_text(self, text, tag)`
- `__init__` (line 392) `def __init__(self, parent, client_id, on_select)`
- `load_latest_client_info` (line 489) `def load_latest_client_info(self)` - *Carga la información más reciente del cliente desde su archivo .log Y su archivo .json de configuración.*
- `update_status_indicator` (line 539) `def update_status_indicator(self)` - *Actualiza el indicador de estado basado en la última actividad*
- `on_card_click` (line 554) `def on_card_click(self, event)`
- `open_console` (line 558) `def open_console(self)`
- `open_files` (line 562) `def open_files(self)`
- `open_processes` (line 566) `def open_processes(self)`
- `load_os_image` (line 570) `def load_os_image(self, client_id)` - *Cargar imagen del sistema operativo basado en el client_id*
- `toggle_view` (line 601) `def toggle_view(self)`
- `show_table_view` (line 611) `def show_table_view(self)`
- `show_card_view` (line 656) `def show_card_view(self)`
- `on_table_select` (line 1011) `def on_table_select(event)`
- `__init__` (line 1116) `def __init__(self, log_dir)`
- `on_modified` (line 1120) `def on_modified(self, event)`
- `process_log_file` (line 1131) `def process_log_file(self, client_id)`
- `configure_canvas_width` (line 1292) `def configure_canvas_width(event)`
- `configure_scroll_region` (line 1342) `def configure_scroll_region(event)`

### SH (1 files)

#### `install.sh`
**Path:** `install.sh`

*No symbols extracted*
