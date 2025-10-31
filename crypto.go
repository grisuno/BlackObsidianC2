package main

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "encoding/base64"
    "encoding/hex"
    "errors"
    "io"
    "os"
)

var (
    ErrInvalidKeyLength = errors.New("clave AES inválida: debe ser exactamente 32 bytes (256 bits)")
    ErrInvalidCiphertext = errors.New("ciphertext inválido: menor a 16 bytes")
    ErrDecryptionFailed = errors.New("fallo la desencriptación")
)

// GetAESKey obtiene la clave AES desde variables de entorno con validación
func GetAESKey() ([]byte, error) {
    keyHex := os.Getenv("C2_AES_KEY")
    if keyHex == "" {
        return nil, errors.New("variable C2_AES_KEY no configurada")
    }
    
    key, err := hex.DecodeString(keyHex)
    if err != nil {
        return nil, errors.New("C2_AES_KEY debe ser una cadena hexadecimal válida")
    }
    
    if len(key) != 32 {
        return nil, ErrInvalidKeyLength
    }
    
    return key, nil
}

// AESEncrypt cifra datos usando AES-256-CFB (compatible con Python)
func AESEncrypt(plaintext []byte) (string, error) {
    key, err := GetAESKey()
    if err != nil {
        return "", err
    }
    
    // Crear cipher block
    block, err := aes.NewCipher(key)
    if err != nil {
        return "", err
    }
    
    // Generar IV aleatorio de 16 bytes
    iv := make([]byte, aes.BlockSize)
    if _, err := io.ReadFull(rand.Reader, iv); err != nil {
        return "", err
    }
    
    // Cifrar usando CFB mode
    ciphertext := make([]byte, len(plaintext))
    stream := cipher.NewCFBEncrypter(block, iv)
    stream.XORKeyStream(ciphertext, plaintext)
    
    // Concatenar: IV + ciphertext
    combined := append(iv, ciphertext...)
    
    // Encodear en base64
    return base64.StdEncoding.EncodeToString(combined), nil
}

// AESDecrypt descifra datos usando AES-256-CFB
func AESDecrypt(encodedData string) ([]byte, error) {
    key, err := GetAESKey()
    if err != nil {
        return nil, err
    }
    
    // Decodear base64
    combined, err := base64.StdEncoding.DecodeString(encodedData)
    if err != nil {
        return nil, err
    }
    
    // Validar longitud mínima
    if len(combined) < aes.BlockSize {
        return nil, ErrInvalidCiphertext
    }
    
    // Extraer IV y ciphertext
    iv := combined[:aes.BlockSize]
    ciphertext := combined[aes.BlockSize:]
    
    // Crear cipher block
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }
    
    // Descifrar usando CFB mode
    plaintext := make([]byte, len(ciphertext))
    stream := cipher.NewCFBDecrypter(block, iv)
    stream.XORKeyStream(plaintext, ciphertext)
    
    return plaintext, nil
}
