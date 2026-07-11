# Polyglot Codebase Knowledge Graph

> Generated offline by **readmenator**. Supports C, C++, Python, Go, Rust, JS/TS, Java, C#, Shell, PHP, Dart, GDScript, Nim, ASM.
> No LLMs. No tokens. Pure static analysis.

**Total Files Parsed:** 7 | **Total Symbols Extracted:** 94 | **Total Imports:** 48

## Structural Knowledge Map
```mermaid
graph TD
    classDef mod fill:#1e1e1e,stroke:#ff6666,stroke-width:2px,color:#fff;
    classDef cls fill:#2d2d2d,stroke:#4ec9b0,stroke-width:2px,color:#fff;
    classDef fn fill:#333,stroke:#dcdcaa,stroke-width:1px,color:#dcdcaa;
    classDef ext fill:#111,stroke:#666,stroke-dasharray: 5 5,color:#aaa;
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
- `GetAESKey` (line 21) - *GetAESKey obtiene la clave AES desde variables de entorno con validación*
- `AESEncrypt` (line 40) - *AESEncrypt cifra datos usando AES-256-CFB (compatible con Python)*
- `AESDecrypt` (line 71) - *AESDecrypt descifra datos usando AES-256-CFB*

#### `handlers.go`
**Path:** `handlers.go`

**Functions:**
- `RegisterC2Routes` (line 30)
- `HandleLogin` (line 83) - *===== HANDLER: Login Compatible con GUI Python =====*
- `HandleGetConnectedClients` (line 125) - *===== HANDLER: Get Connected Clients (Compatible con GUI) =====*
- `HandleIssueCommandLegacy` (line 154) - *===== HANDLER: Issue Command Legacy (sin autenticación) =====*
- `HandleGetCommand` (line 195) - *===== RESTO DE HANDLERS (igual que antes) =====*
- `HandleUpload` (line 267)
- `HandleDownload` (line 308)
- `HandleIssueCommand` (line 343)
- `HandlePostResult` (line 383)
- `writeLogCSV` (line 526) - *✅ FUNCIÓN NUEVA: Escribir logs en CSV*
- `HandleListClients` (line 595)
- `HandleGetResults` (line 626)
- `truncate` (line 661)

#### `main.go`
**Path:** `main.go`

**Functions:**
- `main` (line 13)

#### `schemas.go`
**Path:** `schemas.go`

**Functions:**
- `InitializeCollections` (line 10) - *InitializeCollections crea las colecciones necesarias si no existen Usando la API de PocketBase v0.31.0*

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

**Classs:**
- `StatusBar` (line 172)
- `ModernTreeview` (line 199)
- `ModernConsole` (line 256)
- `ImplantCard` (line 391)
- `LogHandler` (line 1115)

**Functions:**
- `setup_modern_theme` (line 73)
- `login` (line 662)
- `show_notification` (line 683) - *Mostrar notificación en el sistema*
- `refresh_clients` (line 699)
- `select_client` (line 718)
- `create_beacon_tab` (line 729) - *Crea una nueva pestaña para un beacon, con subpestañas de Consola e Intel.*
- `load_implant_config` (line 756) - *Carga la configuración del implant desde el archivo JSON.*
- `create_intel_tab` (line 769) - *Crea la pestaña de Inteligencia/Recon para un beacon específico.*
- `create_section_header` (line 850) - *Crea un encabezado de sección para la pestaña de Intel.*
- `load_intel_data` (line 857) - *Extrae y parsea datos de inteligencia del archivo .log del cliente.*
- `is_client_active` (line 922)
- `load_latest_client_info_dict` (line 928)
- `refresh_table_view` (line 963)
- `create_table_view_frame` (line 987)
- `toggle_global_view` (line 1021)
- `create_section_header` (line 1042) - *Crea un encabezado de sección para la pestaña de Intel.*
- `load_intel_data` (line 1049) - *Extrae y parsea datos de inteligencia del archivo .log del cliente.*
- `process_queue` (line 1163)
- `start_polling` (line 1181)
- `stop_polling` (line 1190)
- `auto_refresh_clients` (line 1197)
- `load_implants_data` (line 1206)
- `load_banners_data` (line 1218)
- `load_access_log` (line 1226)
- `upload_file` (line 1244)
- `create_modern_ui` (line 1265)
- `create_tools_grid` (line 1427) - *Crear grid de herramientas*
- `create_tool_card` (line 1479) - *Crear card de herramienta*
- `create_modern_menu` (line 1503) - *Crear menú moderno*
- `setup_keyboard_shortcuts` (line 1541) - *Configurar atajos de teclado*
- `export_logs` (line 1551) - *Exportar logs a archivo*
- `manage_payloads` (line 1586) - *Ventana de gestión de payloads*
- `show_statistics` (line 1626) - *Mostrar ventana de estadísticas*
- `show_settings` (line 1668) - *Mostrar ventana de configuración*
- `show_help` (line 1721) - *Mostrar ayuda*
- `show_about` (line 1792) - *Mostrar información del programa*
- `on_closing` (line 1804) - *Manejar cierre de aplicación*
- `main` (line 1811) - *Función principal*
- `__init__` (line 173)
- `update_status` (line 188)
- `update_time` (line 194)
- `__init__` (line 200)
- `refresh_data` (line 241)
- `set_title` (line 253)
- `__init__` (line 257)
- `load_command_history` (line 310) - *Carga el historial de comandos desde el archivo .log del cliente*
- `send_command` (line 333)
- `navigate_history_up` (line 358) - *Navegar hacia arriba en el historial de comandos*
- `navigate_history_down` (line 372) - *Navegar hacia abajo en el historial de comandos*
- `add_text` (line 387)
- `__init__` (line 392)
- `load_latest_client_info` (line 489) - *Carga la información más reciente del cliente desde su archivo .log Y su archivo .json de configuración.*
- `update_status_indicator` (line 539) - *Actualiza el indicador de estado basado en la última actividad*
- `on_card_click` (line 554)
- `open_console` (line 558)
- `open_files` (line 562)
- `open_processes` (line 566)
- `load_os_image` (line 570) - *Cargar imagen del sistema operativo basado en el client_id*
- `toggle_view` (line 601)
- `show_table_view` (line 611)
- `show_card_view` (line 656)
- `on_table_select` (line 1011)
- `__init__` (line 1116)
- `on_modified` (line 1120)
- `process_log_file` (line 1131)
- `configure_canvas_width` (line 1292)
- `configure_scroll_region` (line 1342)

### SH (1 files)

#### `install.sh`
**Path:** `install.sh`

*No symbols extracted*
