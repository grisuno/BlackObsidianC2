# API

## app.py

### setup_modern_theme `def setup_modern_theme()`
- Defined: `app.py:73`

### login `def login()`
- Defined: `app.py:662`

### show_notification `def show_notification(message, type)`
- Defined: `app.py:683`
- Doc: Mostrar notificación en el sistema

### refresh_clients `def refresh_clients()`
- Defined: `app.py:699`

### select_client `def select_client(client_id)`
- Defined: `app.py:718`

### create_beacon_tab `def create_beacon_tab(client_id)`
- Defined: `app.py:729`
- Doc: Crea una nueva pestaña para un beacon, con subpestañas de Consola e Intel.

### load_implant_config `def load_implant_config(client_id)`
- Defined: `app.py:756`
- Doc: Carga la configuración del implant desde el archivo JSON.

### create_intel_tab `def create_intel_tab(parent, client_id)`
- Defined: `app.py:769`
- Doc: Crea la pestaña de Inteligencia/Recon para un beacon específico.

### create_section_header `def create_section_header(parent, title)`
- Defined: `app.py:850`
- Doc: Crea un encabezado de sección para la pestaña de Intel.

### load_intel_data `def load_intel_data(client_id)`
- Defined: `app.py:857`
- Doc: Extrae y parsea datos de inteligencia del archivo .log del cliente.

### is_client_active `def is_client_active(client_id)`
- Defined: `app.py:922`

### load_latest_client_info_dict `def load_latest_client_info_dict(client_id)`
- Defined: `app.py:928`

### refresh_table_view `def refresh_table_view(tree)`
- Defined: `app.py:963`

### create_table_view_frame `def create_table_view_frame(parent)`
- Defined: `app.py:987`

### toggle_global_view `def toggle_global_view()`
- Defined: `app.py:1021`

### create_section_header `def create_section_header(parent, title)`
- Defined: `app.py:1042`
- Doc: Crea un encabezado de sección para la pestaña de Intel.

### load_intel_data `def load_intel_data(client_id)`
- Defined: `app.py:1049`
- Doc: Extrae y parsea datos de inteligencia del archivo .log del cliente.

### process_queue `def process_queue()`
- Defined: `app.py:1163`

### start_polling `def start_polling()`
- Defined: `app.py:1181`

### stop_polling `def stop_polling()`
- Defined: `app.py:1190`

### auto_refresh_clients `def auto_refresh_clients()`
- Defined: `app.py:1197`

### load_implants_data `def load_implants_data()`
- Defined: `app.py:1206`

### load_banners_data `def load_banners_data()`
- Defined: `app.py:1218`

### load_access_log `def load_access_log()`
- Defined: `app.py:1226`

### upload_file `def upload_file()`
- Defined: `app.py:1244`

### create_modern_ui `def create_modern_ui()`
- Defined: `app.py:1265`

### create_tools_grid `def create_tools_grid(parent)`
- Defined: `app.py:1427`
- Doc: Crear grid de herramientas

### create_tool_card `def create_tool_card(parent, tool)`
- Defined: `app.py:1479`
- Doc: Crear card de herramienta

### create_modern_menu `def create_modern_menu()`
- Defined: `app.py:1503`
- Doc: Crear menú moderno

### setup_keyboard_shortcuts `def setup_keyboard_shortcuts()`
- Defined: `app.py:1541`
- Doc: Configurar atajos de teclado

### export_logs `def export_logs()`
- Defined: `app.py:1551`
- Doc: Exportar logs a archivo

### manage_payloads `def manage_payloads()`
- Defined: `app.py:1586`
- Doc: Ventana de gestión de payloads

### show_statistics `def show_statistics()`
- Defined: `app.py:1626`
- Doc: Mostrar ventana de estadísticas

### show_settings `def show_settings()`
- Defined: `app.py:1668`
- Doc: Mostrar ventana de configuración

### show_help `def show_help()`
- Defined: `app.py:1721`
- Doc: Mostrar ayuda

### show_about `def show_about()`
- Defined: `app.py:1792`
- Doc: Mostrar información del programa

### on_closing `def on_closing()`
- Defined: `app.py:1804`
- Doc: Manejar cierre de aplicación

### main `def main()`
- Defined: `app.py:1811`
- Doc: Función principal

### __init__ `def __init__(self, parent)`
- Defined: `app.py:173`

### update_status `def update_status(self, text, connected)`
- Defined: `app.py:188`

### update_time `def update_time(self)`
- Defined: `app.py:194`

### __init__ `def __init__(self, parent, columns, data_loader)`
- Defined: `app.py:200`

### refresh_data `def refresh_data(self)`
- Defined: `app.py:241`

### set_title `def set_title(self, title)`
- Defined: `app.py:253`

### __init__ `def __init__(self, parent, client_id)`
- Defined: `app.py:257`

### load_command_history `def load_command_history(self)`
- Defined: `app.py:310`
- Doc: Carga el historial de comandos desde el archivo .log del cliente

### send_command `def send_command(self, client_id)`
- Defined: `app.py:333`

### navigate_history_up `def navigate_history_up(self, event)`
- Defined: `app.py:358`
- Doc: Navegar hacia arriba en el historial de comandos

### navigate_history_down `def navigate_history_down(self, event)`
- Defined: `app.py:372`
- Doc: Navegar hacia abajo en el historial de comandos

### add_text `def add_text(self, text, tag)`
- Defined: `app.py:387`

### __init__ `def __init__(self, parent, client_id, on_select)`
- Defined: `app.py:392`

### load_latest_client_info `def load_latest_client_info(self)`
- Defined: `app.py:489`
- Doc: Carga la información más reciente del cliente desde su archivo .log Y su archivo .json de configuración.

### update_status_indicator `def update_status_indicator(self)`
- Defined: `app.py:539`
- Doc: Actualiza el indicador de estado basado en la última actividad

### on_card_click `def on_card_click(self, event)`
- Defined: `app.py:554`

### open_console `def open_console(self)`
- Defined: `app.py:558`

### open_files `def open_files(self)`
- Defined: `app.py:562`

### open_processes `def open_processes(self)`
- Defined: `app.py:566`

### load_os_image `def load_os_image(self, client_id)`
- Defined: `app.py:570`
- Doc: Cargar imagen del sistema operativo basado en el client_id

### toggle_view `def toggle_view(self)`
- Defined: `app.py:601`

### show_table_view `def show_table_view(self)`
- Defined: `app.py:611`

### show_card_view `def show_card_view(self)`
- Defined: `app.py:656`

### on_table_select `def on_table_select(event)`
- Defined: `app.py:1011`

### __init__ `def __init__(self, log_dir)`
- Defined: `app.py:1116`

### on_modified `def on_modified(self, event)`
- Defined: `app.py:1120`

### process_log_file `def process_log_file(self, client_id)`
- Defined: `app.py:1131`

### configure_canvas_width `def configure_canvas_width(event)`
- Defined: `app.py:1292`

### configure_scroll_region `def configure_scroll_region(event)`
- Defined: `app.py:1342`

## crypto.go

### GetAESKey `func GetAESKey(`
- Defined: `crypto.go:21`
- Doc: GetAESKey obtiene la clave AES desde variables de entorno con validación

### AESEncrypt `func AESEncrypt(`
- Defined: `crypto.go:40`
- Doc: AESEncrypt cifra datos usando AES-256-CFB (compatible con Python)

### AESDecrypt `func AESDecrypt(`
- Defined: `crypto.go:71`
- Doc: AESDecrypt descifra datos usando AES-256-CFB

## handlers.go

### RegisterC2Routes `func RegisterC2Routes(`
- Defined: `handlers.go:30`

### HandleLogin `func HandleLogin(`
- Defined: `handlers.go:83`
- Doc: ===== HANDLER: Login Compatible con GUI Python =====

### HandleGetConnectedClients `func HandleGetConnectedClients(`
- Defined: `handlers.go:125`
- Doc: ===== HANDLER: Get Connected Clients (Compatible con GUI) =====

### HandleIssueCommandLegacy `func HandleIssueCommandLegacy(`
- Defined: `handlers.go:154`
- Doc: ===== HANDLER: Issue Command Legacy (sin autenticación) =====

### HandleGetCommand `func HandleGetCommand(`
- Defined: `handlers.go:195`

### HandleUpload `func HandleUpload(`
- Defined: `handlers.go:267`

### HandleDownload `func HandleDownload(`
- Defined: `handlers.go:308`

### HandleIssueCommand `func HandleIssueCommand(`
- Defined: `handlers.go:343`

### HandlePostResult `func HandlePostResult(`
- Defined: `handlers.go:383`

### writeLogCSV `func writeLogCSV(`
- Defined: `handlers.go:526`
- Doc: ✅ FUNCIÓN NUEVA: Escribir logs en CSV

### HandleListClients `func HandleListClients(`
- Defined: `handlers.go:595`

### HandleGetResults `func HandleGetResults(`
- Defined: `handlers.go:626`

### truncate `func truncate(`
- Defined: `handlers.go:661`

## main.go

### main `func main(`
- Defined: `main.go:13`

## schemas.go

### InitializeCollections `func InitializeCollections(`
- Defined: `schemas.go:10`
- Doc: InitializeCollections crea las colecciones necesarias si no existen Usando la API de PocketBase v0.31.0

## web/js/dashboard.js

### loadDashboard
- Defined: `web/js/dashboard.js:10`

### loadImplantsList
- Defined: `web/js/dashboard.js:36`

### openTerminal
- Defined: `web/js/dashboard.js:61`

### logout
- Defined: `web/js/dashboard.js:65`
