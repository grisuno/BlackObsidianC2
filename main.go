package main

import (
    "crypto/tls"
    "log"
    "os"
    "path/filepath"
    "net/http"
    "github.com/pocketbase/pocketbase"
    "github.com/pocketbase/pocketbase/core"
)

func main() {
    app := pocketbase.New()
    
    if os.Getenv("C2_AES_KEY") == "" {
        log.Fatal("[CRÍTICO] Variable C2_AES_KEY no configurada")
    }
    
    certFile := os.Getenv("TLS_CERT")
    keyFile := os.Getenv("TLS_KEY")
    
    if certFile == "" || keyFile == "" {
        log.Fatal("[CRÍTICO] Variables TLS_CERT y TLS_KEY requeridas")
    }
    
    cert, err := tls.LoadX509KeyPair(certFile, keyFile)
    if err != nil {
        log.Fatalf("[CRÍTICO] Error cargando certificados: %v", err)
    }
    log.Println("[✓] Certificados cargados correctamente")
    
    app.OnServe().BindFunc(func(se *core.ServeEvent) error {
        se.Server.TLSConfig = &tls.Config{
            Certificates: []tls.Certificate{cert},
            MinVersion:   tls.VersionTLS12,
        }
        
        if err := InitializeCollections(app); err != nil {
            return err
        }
        
        // ✅ REGISTRAR RUTAS DE API
        RegisterC2Routes(se, app)
        
        // ✅ SERVIR ARCHIVOS HTML ESPECÍFICOS
        webDir, _ := filepath.Abs("./web")
        
        // Rutas específicas
        se.Router.GET("/", func(e *core.RequestEvent) error {
            http.ServeFile(e.Response, e.Request, filepath.Join(webDir, "dashboard.html"))
            return nil
        })
        
        se.Router.GET("/dashboard.html", func(e *core.RequestEvent) error {
            http.ServeFile(e.Response, e.Request, filepath.Join(webDir, "dashboard.html"))
            return nil
        })
        
        se.Router.GET("/index.html", func(e *core.RequestEvent) error {
            http.ServeFile(e.Response, e.Request, filepath.Join(webDir, "index.html"))
            return nil
        })
        
        se.Router.GET("/css/{file}", func(e *core.RequestEvent) error {
            file := e.Request.PathValue("file")
            http.ServeFile(e.Response, e.Request, filepath.Join(webDir, "css", file))
            return nil
        })
        
        se.Router.GET("/js/{file}", func(e *core.RequestEvent) error {
            file := e.Request.PathValue("file")
            http.ServeFile(e.Response, e.Request, filepath.Join(webDir, "js", file))
            return nil
        })
        
        se.Router.GET("/img/{file}", func(e *core.RequestEvent) error {
            file := e.Request.PathValue("file")
            http.ServeFile(e.Response, e.Request, filepath.Join(webDir, "img", file))
            return nil
        })
        
        log.Println("[✓] Black Obsidian C2 iniciado")
        log.Println("[✓] Dashboard: https://127.0.0.1:4444/")
        return se.Next()
    })
    
    if err := app.Start(); err != nil {
        log.Fatal(err)
    }
}
