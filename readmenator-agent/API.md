# API

## app.py
- `setup_modern_theme` (function) `app.py:73` `def setup_modern_theme()`
- `StatusBar.__init__` (method) `app.py:173` `def __init__(self, parent)`
- `StatusBar.update_status` (method) `app.py:188` `def update_status(self, text, connected)`
- `StatusBar.update_time` (method) `app.py:194` `def update_time(self)`
- `ModernTreeview.__init__` (method) `app.py:200` `def __init__(self, parent, columns, data_loader)`
- `ModernTreeview.refresh_data` (method) `app.py:241` `def refresh_data(self)`
- `ModernTreeview.set_title` (method) `app.py:253` `def set_title(self, title)`
- `ModernConsole.__init__` (method) `app.py:257` `def __init__(self, parent, client_id)`
- `ModernConsole.load_command_history` (method) `app.py:310` `def load_command_history(self)` -- Carga el historial de comandos desde el archivo .log del cliente
- `ModernConsole.send_command` (method) `app.py:333` `def send_command(self, client_id)`
- `ModernConsole.navigate_history_up` (method) `app.py:358` `def navigate_history_up(self, event)` -- Navegar hacia arriba en el historial de comandos
- `ModernConsole.navigate_history_down` (method) `app.py:372` `def navigate_history_down(self, event)` -- Navegar hacia abajo en el historial de comandos
- `ModernConsole.add_text` (method) `app.py:387` `def add_text(self, text, tag)`
- `ImplantCard.__init__` (method) `app.py:392` `def __init__(self, parent, client_id, on_select)`
- `ImplantCard.load_latest_client_info` (method) `app.py:489` `def load_latest_client_info(self)` -- Carga la información más reciente del cliente desde su archivo .log Y su archivo .json de configuración.
- `ImplantCard.update_status_indicator` (method) `app.py:539` `def update_status_indicator(self)` -- Actualiza el indicador de estado basado en la última actividad
- `ImplantCard.on_card_click` (method) `app.py:554` `def on_card_click(self, event)`
- `ImplantCard.open_console` (method) `app.py:558` `def open_console(self)`
- `ImplantCard.open_files` (method) `app.py:562` `def open_files(self)`
- `ImplantCard.open_processes` (method) `app.py:566` `def open_processes(self)`
- `ImplantCard.load_os_image` (method) `app.py:570` `def load_os_image(self, client_id)` -- Cargar imagen del sistema operativo basado en el client_id
- `ImplantCard.toggle_view` (method) `app.py:601` `def toggle_view(self)`
- `ImplantCard.show_table_view` (method) `app.py:611` `def show_table_view(self)`
- `ImplantCard.show_card_view` (method) `app.py:656` `def show_card_view(self)`
- `ImplantCard.login` (method) `app.py:662` `def login()`
- `ImplantCard.show_notification` (method) `app.py:683` `def show_notification(message, type)` -- Mostrar notificación en el sistema
- `ImplantCard.refresh_clients` (method) `app.py:699` `def refresh_clients()`
- `ImplantCard.select_client` (method) `app.py:718` `def select_client(client_id)`
- `ImplantCard.create_beacon_tab` (method) `app.py:729` `def create_beacon_tab(client_id)` -- Crea una nueva pestaña para un beacon, con subpestañas de Consola e Intel.
- `ImplantCard.load_implant_config` (method) `app.py:756` `def load_implant_config(client_id)` -- Carga la configuración del implant desde el archivo JSON.
- `ImplantCard.create_intel_tab` (method) `app.py:769` `def create_intel_tab(parent, client_id)` -- Crea la pestaña de Inteligencia/Recon para un beacon específico.
- `ImplantCard.create_section_header` (method) `app.py:850` `def create_section_header(parent, title)` -- Crea un encabezado de sección para la pestaña de Intel.
- `ImplantCard.load_intel_data` (method) `app.py:857` `def load_intel_data(client_id)` -- Extrae y parsea datos de inteligencia del archivo .log del cliente.
- `ImplantCard.is_client_active` (method) `app.py:922` `def is_client_active(client_id)`
- `ImplantCard.load_latest_client_info_dict` (method) `app.py:928` `def load_latest_client_info_dict(client_id)`
- `ImplantCard.refresh_table_view` (method) `app.py:963` `def refresh_table_view(tree)`
- `ImplantCard.create_table_view_frame` (method) `app.py:987` `def create_table_view_frame(parent)`
- `ImplantCard.on_table_select` (method) `app.py:1011` `def on_table_select(event)`
- `ImplantCard.toggle_global_view` (method) `app.py:1021` `def toggle_global_view()`
- `ImplantCard.create_section_header` (method) `app.py:1042` `def create_section_header(parent, title)` -- Crea un encabezado de sección para la pestaña de Intel.
- `ImplantCard.load_intel_data` (method) `app.py:1049` `def load_intel_data(client_id)` -- Extrae y parsea datos de inteligencia del archivo .log del cliente.
- `LogHandler.__init__` (method) `app.py:1116` `def __init__(self, log_dir)`
- `LogHandler.on_modified` (method) `app.py:1120` `def on_modified(self, event)`
- `LogHandler.process_log_file` (method) `app.py:1131` `def process_log_file(self, client_id)`
- `LogHandler.process_queue` (method) `app.py:1163` `def process_queue()`
- `LogHandler.start_polling` (method) `app.py:1181` `def start_polling()`
- `LogHandler.stop_polling` (method) `app.py:1190` `def stop_polling()`
- `LogHandler.auto_refresh_clients` (method) `app.py:1197` `def auto_refresh_clients()`
- `LogHandler.load_implants_data` (method) `app.py:1206` `def load_implants_data()`
- `LogHandler.load_banners_data` (method) `app.py:1218` `def load_banners_data()`
- `LogHandler.load_access_log` (method) `app.py:1226` `def load_access_log()`
- `LogHandler.upload_file` (method) `app.py:1244` `def upload_file()`
- `LogHandler.create_modern_ui` (method) `app.py:1265` `def create_modern_ui()`
- `LogHandler.configure_canvas_width` (method) `app.py:1292` `def configure_canvas_width(event)`
- `LogHandler.configure_scroll_region` (method) `app.py:1342` `def configure_scroll_region(event)`
- `LogHandler.create_tools_grid` (method) `app.py:1427` `def create_tools_grid(parent)` -- Crear grid de herramientas
- `LogHandler.create_tool_card` (method) `app.py:1479` `def create_tool_card(parent, tool)` -- Crear card de herramienta
- `LogHandler.create_modern_menu` (method) `app.py:1503` `def create_modern_menu()` -- Crear menú moderno
- `LogHandler.setup_keyboard_shortcuts` (method) `app.py:1541` `def setup_keyboard_shortcuts()` -- Configurar atajos de teclado
- `LogHandler.export_logs` (method) `app.py:1551` `def export_logs()` -- Exportar logs a archivo
- `LogHandler.manage_payloads` (method) `app.py:1586` `def manage_payloads()` -- Ventana de gestión de payloads
- `LogHandler.show_statistics` (method) `app.py:1626` `def show_statistics()` -- Mostrar ventana de estadísticas
- `LogHandler.show_settings` (method) `app.py:1668` `def show_settings()` -- Mostrar ventana de configuración
- `LogHandler.show_help` (method) `app.py:1721` `def show_help()` -- Mostrar ayuda
- `LogHandler.show_about` (method) `app.py:1792` `def show_about()` -- Mostrar información del programa
- `LogHandler.on_closing` (method) `app.py:1804` `def on_closing()` -- Manejar cierre de aplicación
- `LogHandler.main` (method) `app.py:1811` `def main()` -- Función principal

## crypto.go
- `GetAESKey` (function) `crypto.go:21` `func GetAESKey(` -- GetAESKey obtiene la clave AES desde variables de entorno con validación
- `AESEncrypt` (function) `crypto.go:40` `func AESEncrypt(` -- AESEncrypt cifra datos usando AES-256-CFB (compatible con Python)
- `AESDecrypt` (function) `crypto.go:71` `func AESDecrypt(` -- AESDecrypt descifra datos usando AES-256-CFB

## handlers.go
- `RegisterC2Routes` (function) `handlers.go:30` `func RegisterC2Routes(`
- `HandleLogin` (function) `handlers.go:83` `func HandleLogin(`
- `HandleGetConnectedClients` (function) `handlers.go:125` `func HandleGetConnectedClients(`
- `HandleIssueCommandLegacy` (function) `handlers.go:154` `func HandleIssueCommandLegacy(`
- `HandleGetCommand` (function) `handlers.go:195` `func HandleGetCommand(`
- `HandleUpload` (function) `handlers.go:267` `func HandleUpload(`
- `HandleDownload` (function) `handlers.go:308` `func HandleDownload(`
- `HandleIssueCommand` (function) `handlers.go:343` `func HandleIssueCommand(`
- `HandlePostResult` (function) `handlers.go:383` `func HandlePostResult(`
- `writeLogCSV` (function) `handlers.go:526` `func writeLogCSV(` -- ✅ FUNCIÓN NUEVA: Escribir logs en CSV
- `HandleListClients` (function) `handlers.go:595` `func HandleListClients(`
- `HandleGetResults` (function) `handlers.go:626` `func HandleGetResults(`
- `truncate` (function) `handlers.go:661` `func truncate(`

## main.go
- `main` (function) `main.go:13` `func main(`

## schemas.go
- `InitializeCollections` (function) `schemas.go:10` `func InitializeCollections(` -- InitializeCollections crea las colecciones necesarias si no existen Usando la API de PocketBase v0.31.0

## web/js/dashboard.js
- `loadDashboard` (function) `web/js/dashboard.js:10`
- `loadImplantsList` (function) `web/js/dashboard.js:36`
- `openTerminal` (function) `web/js/dashboard.js:61`
- `logout` (function) `web/js/dashboard.js:65`
