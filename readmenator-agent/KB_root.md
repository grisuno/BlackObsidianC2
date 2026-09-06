# Subsystem: root

## app.py
- Layer: utility
- Doc: _*_ coding: utf8 _*_
- Language: py
- Symbols:
  - `setup_modern_theme` (function, line 73) `def setup_modern_theme()`
  - `StatusBar` (class, line 172) `class StatusBar(Frame)`
  - `ModernTreeview` (class, line 199) `class ModernTreeview(Frame)`
  - `ModernConsole` (class, line 256) `class ModernConsole(Frame)`
  - `ImplantCard` (class, line 391) `class ImplantCard(Frame)`
  - `login` (method, line 662) `def login()`
  - `show_notification` (method, line 683) `def show_notification(message, type)`
  - `refresh_clients` (method, line 699) `def refresh_clients()`
  - `select_client` (method, line 718) `def select_client(client_id)`
  - `create_beacon_tab` (method, line 729) `def create_beacon_tab(client_id)`
  - `load_implant_config` (method, line 756) `def load_implant_config(client_id)`
  - `create_intel_tab` (method, line 769) `def create_intel_tab(parent, client_id)`
  - `create_section_header` (method, line 850) `def create_section_header(parent, title)`
  - `load_intel_data` (method, line 857) `def load_intel_data(client_id)`
  - `is_client_active` (method, line 922) `def is_client_active(client_id)`
  - `load_latest_client_info_dict` (method, line 928) `def load_latest_client_info_dict(client_id)`
  - `refresh_table_view` (method, line 963) `def refresh_table_view(tree)`
  - `create_table_view_frame` (method, line 987) `def create_table_view_frame(parent)`
  - `toggle_global_view` (method, line 1021) `def toggle_global_view()`
  - `create_section_header` (method, line 1042) `def create_section_header(parent, title)`
  - `load_intel_data` (method, line 1049) `def load_intel_data(client_id)`
  - `LogHandler` (class, line 1115) `class LogHandler(FileSystemEventHandler)`
  - `process_queue` (method, line 1163) `def process_queue()`
  - `start_polling` (method, line 1181) `def start_polling()`
  - `stop_polling` (method, line 1190) `def stop_polling()`
  - `auto_refresh_clients` (method, line 1197) `def auto_refresh_clients()`
  - `load_implants_data` (method, line 1206) `def load_implants_data()`
  - `load_banners_data` (method, line 1218) `def load_banners_data()`
  - `load_access_log` (method, line 1226) `def load_access_log()`
  - `upload_file` (method, line 1244) `def upload_file()`
  - `create_modern_ui` (method, line 1265) `def create_modern_ui()`
  - `create_tools_grid` (method, line 1427) `def create_tools_grid(parent)`
  - `create_tool_card` (method, line 1479) `def create_tool_card(parent, tool)`
  - `create_modern_menu` (method, line 1503) `def create_modern_menu()`
  - `setup_keyboard_shortcuts` (method, line 1541) `def setup_keyboard_shortcuts()`
  - `export_logs` (method, line 1551) `def export_logs()`
  - `manage_payloads` (method, line 1586) `def manage_payloads()`
  - `show_statistics` (method, line 1626) `def show_statistics()`
  - `show_settings` (method, line 1668) `def show_settings()`
  - `show_help` (method, line 1721) `def show_help()`
  - `show_about` (method, line 1792) `def show_about()`
  - `on_closing` (method, line 1804) `def on_closing()`
  - `main` (method, line 1811) `def main()`
  - `__init__` (method, line 173) `def __init__(self, parent)`
  - `update_status` (method, line 188) `def update_status(self, text, connected)`
  - `update_time` (method, line 194) `def update_time(self)`
  - `__init__` (method, line 200) `def __init__(self, parent, columns, data_loader)`
  - `refresh_data` (method, line 241) `def refresh_data(self)`
  - `set_title` (method, line 253) `def set_title(self, title)`
  - `__init__` (method, line 257) `def __init__(self, parent, client_id)`
  - `load_command_history` (method, line 310) `def load_command_history(self)`
  - `send_command` (method, line 333) `def send_command(self, client_id)`
  - `navigate_history_up` (method, line 358) `def navigate_history_up(self, event)`
  - `navigate_history_down` (method, line 372) `def navigate_history_down(self, event)`
  - `add_text` (method, line 387) `def add_text(self, text, tag)`
  - `__init__` (method, line 392) `def __init__(self, parent, client_id, on_select)`
  - `load_latest_client_info` (method, line 489) `def load_latest_client_info(self)`
  - `update_status_indicator` (method, line 539) `def update_status_indicator(self)`
  - `on_card_click` (method, line 554) `def on_card_click(self, event)`
  - `open_console` (method, line 558) `def open_console(self)`
  - `open_files` (method, line 562) `def open_files(self)`
  - `open_processes` (method, line 566) `def open_processes(self)`
  - `load_os_image` (method, line 570) `def load_os_image(self, client_id)`
  - `toggle_view` (method, line 601) `def toggle_view(self)`
  - `show_table_view` (method, line 611) `def show_table_view(self)`
  - `show_card_view` (method, line 656) `def show_card_view(self)`
  - `on_table_select` (method, line 1011) `def on_table_select(event)`
  - `__init__` (method, line 1116) `def __init__(self, log_dir)`
  - `on_modified` (method, line 1120) `def on_modified(self, event)`
  - `process_log_file` (method, line 1131) `def process_log_file(self, client_id)`
  - `configure_canvas_width` (method, line 1292) `def configure_canvas_width(event)`
  - `configure_scroll_region` (method, line 1342) `def configure_scroll_region(event)`

## crypto.go
- Layer: utility
- Language: go
- Symbols:
  - `GetAESKey` (function, line 21) `func GetAESKey(`
  - `AESEncrypt` (function, line 40) `func AESEncrypt(`
  - `AESDecrypt` (function, line 71) `func AESDecrypt(`

## handlers.go
- Layer: presentation
- Language: go
- Symbols:
  - `RegisterC2Routes` (function, line 30) `func RegisterC2Routes(`
  - `HandleLogin` (function, line 83) `func HandleLogin(`
  - `HandleGetConnectedClients` (function, line 125) `func HandleGetConnectedClients(`
  - `HandleIssueCommandLegacy` (function, line 154) `func HandleIssueCommandLegacy(`
  - `HandleGetCommand` (function, line 195) `func HandleGetCommand(`
  - `HandleUpload` (function, line 267) `func HandleUpload(`
  - `HandleDownload` (function, line 308) `func HandleDownload(`
  - `HandleIssueCommand` (function, line 343) `func HandleIssueCommand(`
  - `HandlePostResult` (function, line 383) `func HandlePostResult(`
  - `writeLogCSV` (function, line 526) `func writeLogCSV(`
  - `HandleListClients` (function, line 595) `func HandleListClients(`
  - `HandleGetResults` (function, line 626) `func HandleGetResults(`
  - `truncate` (function, line 661) `func truncate(`

## install.sh
- Layer: utility
- Language: sh

## main.go
- Layer: utility
- Language: go
- Symbols:
  - `main` (function, line 13) `func main(`

## schemas.go
- Layer: utility
- Language: go
- Symbols:
  - `InitializeCollections` (function, line 10) `func InitializeCollections(`
