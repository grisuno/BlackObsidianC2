#!/bin/bash
go build -o c2-server

# Ejecutar
export TLS_CERT=./cert.pem
export TLS_KEY=./key_go.pem
export C2_AES_KEY=18547a9428b62fdf2ba11cebc786bccbca8a941748d3acf4aad100ac65d0477f
./c2-server serve --https=0.0.0.0:4444
