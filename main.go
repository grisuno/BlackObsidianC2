package main

import (
    "crypto/tls"
    "log"
    "os"
    
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
    
    // ✅ Cargar certificados ANTES
    cert, err := tls.LoadX509KeyPair(certFile, keyFile)
    if err != nil {
        log.Fatalf("[CRÍTICO] Error cargando certificados: %v\nUSA: openssl pkcs8 -topk8 -nocrypt -in key.pem -out key_go.pem", err)
    }
    log.Println("[✓] Certificados cargados correctamente")
    
    app.OnServe().BindFunc(func(se *core.ServeEvent) error {
        // ✅ Configurar TLS en OnServe (hook correcto)
        se.Server.TLSConfig = &tls.Config{
            Certificates: []tls.Certificate{cert},
            MinVersion:   tls.VersionTLS12,
            MaxVersion:   tls.VersionTLS13,
            ClientAuth:   tls.NoClientCert,
            CipherSuites: []uint16{
                tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
                tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
                tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
                tls.TLS_RSA_WITH_AES_128_GCM_SHA256,
            },
        }
        log.Println("[✓] TLS configurado en el servidor")
        
        if err := InitializeCollections(app); err != nil {
            log.Printf("[ERROR] Fallo al inicializar colecciones: %v", err)
            return err
        }
        
        RegisterC2Routes(se, app)
        
        log.Println("[✓] Servidor C2 iniciado con HTTPS")
        return se.Next()
    })
    
    if err := app.Start(); err != nil {
        log.Fatal(err)
    }
}
