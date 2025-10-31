package main

import (
    "github.com/pocketbase/pocketbase"
    "github.com/pocketbase/pocketbase/core"
)

// InitializeCollections crea las colecciones necesarias si no existen
// Usando la API de PocketBase v0.31.0
func InitializeCollections(app *pocketbase.PocketBase) error {
    // ===== Colección: c2_clients =====
    _, err := app.FindCollectionByNameOrId("c2_clients")
    if err != nil {
        // Crear colección base
        collection := core.NewBaseCollection("c2_clients")
        
        // Configurar reglas de acceso (solo admins)
        collection.ListRule = nil
        collection.ViewRule = nil
        collection.CreateRule = nil
        collection.UpdateRule = nil
        collection.DeleteRule = nil
        
        // Definir campos usando la nueva API
        collection.Fields.Add(&core.TextField{
            Name:     "client_id",
            Required: true,
            Min:      1,
            Max:      100,
            Pattern:  `^[a-zA-Z0-9_-]+$`,
        })
        
        collection.Fields.Add(&core.TextField{
            Name:     "hostname",
            Required: false,
            Max:      255,
        })
        
        collection.Fields.Add(&core.TextField{
            Name:     "ips",
            Required: false,
            Max:      500,
        })
        
        collection.Fields.Add(&core.TextField{
            Name:     "os",
            Required: false,
            Max:      50,
        })
        
        collection.Fields.Add(&core.TextField{
            Name:     "user",
            Required: false,
            Max:      100,
        })
        
        collection.Fields.Add(&core.NumberField{
            Name:     "pid",
            Required: false,
        })
        
        collection.Fields.Add(&core.DateField{
            Name:     "last_seen",
            Required: true,
        })
        
        collection.Fields.Add(&core.SelectField{
            Name:     "status",
            Required: true,
            Values:   []string{"active", "inactive", "compromised"},
        })
        
        // Crear índice único para client_id
        collection.AddIndex("idx_unique_client_id", true, "client_id", "")
        
        // Guardar colección
        if err := app.Save(collection); err != nil {
            return err
        }
    }
    
    // ===== Colección: c2_commands =====
    _, err = app.FindCollectionByNameOrId("c2_commands")
    if err != nil {
        collection := core.NewBaseCollection("c2_commands")
        collection.ListRule = nil
        collection.ViewRule = nil
        collection.CreateRule = nil
        collection.UpdateRule = nil
        collection.DeleteRule = nil
        
        collection.Fields.Add(&core.TextField{
            Name:     "client_id",
            Required: true,
            Pattern:  `^[a-zA-Z0-9_-]+$`,
        })
        
        collection.Fields.Add(&core.TextField{
            Name:     "command",
            Required: true,
            Min:      1,
            Max:      5000,
        })
        
        collection.Fields.Add(&core.DateField{
            Name:     "issued_at",
            Required: true,
        })
        
        collection.Fields.Add(&core.SelectField{
            Name:     "status",
            Required: true,
            Values:   []string{"pending", "delivered", "executed", "failed"},
        })
        
        // Índices para consultas rápidas
        collection.AddIndex("idx_commands_client_status", false, "client_id", "status")
        
        if err := app.Save(collection); err != nil {
            return err
        }
    }
    
    // ===== Colección: c2_results =====
    _, err = app.FindCollectionByNameOrId("c2_results")
    if err != nil {
        collection := core.NewBaseCollection("c2_results")
        collection.ListRule = nil
        collection.ViewRule = nil
        collection.CreateRule = nil
        collection.UpdateRule = nil
        collection.DeleteRule = nil
        
        collection.Fields.Add(&core.TextField{
            Name:     "client_id",
            Required: true,
            Pattern:  `^[a-zA-Z0-9_-]+$`,
        })
        
        collection.Fields.Add(&core.TextField{
            Name:     "command",
            Required: true,
            Max:      5000,
        })
        
        collection.Fields.Add(&core.TextField{
            Name:     "output",
            Required: false,
            Max:      100000,
        })
        
        collection.Fields.Add(&core.TextField{
            Name:     "hostname",
            Required: false,
            Max:      255,
        })
        
        collection.Fields.Add(&core.TextField{
            Name:     "pwd",
            Required: false,
            Max:      1000,
        })
        
        collection.Fields.Add(&core.DateField{
            Name:     "received_at",
            Required: true,
        })
        
        collection.AddIndex("idx_results_client", false, "client_id", "")
        
        if err := app.Save(collection); err != nil {
            return err
        }
    }
    
    // ===== Colección: c2_files =====
    _, err = app.FindCollectionByNameOrId("c2_files")
    if err != nil {
        collection := core.NewBaseCollection("c2_files")
        collection.ListRule = nil
        collection.ViewRule = nil
        collection.CreateRule = nil
        collection.UpdateRule = nil
        collection.DeleteRule = nil
        
        collection.Fields.Add(&core.TextField{
            Name:     "client_id",
            Required: false,
            Pattern:  `^[a-zA-Z0-9_-]+$`,
        })
        
        collection.Fields.Add(&core.TextField{
            Name:     "filename",
            Required: true,
            Min:      1,
            Max:      255,
        })
        
        collection.Fields.Add(&core.FileField{
            Name:      "file",
            Required:  true,
            MaxSelect: 1,
            MaxSize:   104857600, // 100MB
        })
        
        collection.Fields.Add(&core.SelectField{
            Name:     "file_type",
            Required: true,
            Values:   []string{"upload", "download", "bof"},
        })
        
        collection.Fields.Add(&core.DateField{
            Name:     "uploaded_at",
            Required: true,
        })
        
        if err := app.Save(collection); err != nil {
            return err
        }
    }
    
    return nil
}
