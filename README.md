<img width="1024" height="1024" alt="image" src="https://github.com/user-attachments/assets/1ee5607c-dcbe-4a58-8f30-e7c720e0062e" />


# 🖤 BLACK OBSIDIAN C2
Modern, realtime, minimalist Command & Control server built in Go with PocketBase

```text
⬛ BLACK OBSIDIAN - Control Panel for Red Team Operations
   ├── Realtime Implant Management
   ├── AES-256 Encryption
   ├── Multi-Platform Support
   └── Modern Web Dashboard
```
## 📋 Description
Black Obsidian is a next-generation C2 (Command & Control) server designed for professional Red Team operations. Built with Go and PocketBase, it offers a modern and specialized alternative to complex C2 servers while maintaining agility, security, and customization ease.

Part of the [LazyOwn](https://github.com/grisuno/LazyOwn) ecosystem, [Black Obsidian](https://github.com/grisuno/BlackObsidianC2) integrates seamlessly with implants such as:

[Black Basalt Beacon (Windows)](https://github.com/grisuno/beacon)

[Black Sand Beacon (Linux ELF)](https://github.com/grisuno/blacksandbeacon)

[Black Serpentine (Python)](https://github.com/grisuno/BlackSerpentine)

[Maleable Implant (Go Multi-Platform)](https://github.com/grisuno/LazyOwn)

## ✨ Key Features
- ✅ HTTPS/TLS Server with self-signed certificates

- ✅ AES-256 Encryption for all communications

- ✅ Realtime Management of connected implants

- ✅ Modern Web Dashboard with dark hacker theme

- ✅ Integrated Database (PocketBase SQLite)

- ✅ Automatic CSV Logs per client

- ✅ Pending Commands with execution status

- ✅ Multi-Implant parallel support

- ✅ Fully Functional REST API

- ✅ Scalable and customizable architecture

## 🛠️ Requirements
Go 1.20+

- OpenSSL (for certificate generation)

- Linux/macOS/Windows with Go support

- Port 4444 available (or customizable)

## Go Dependencies
```bash
go get github.com/pocketbase/pocketbase
```
## 📦 Installation

1. Clone Repository

```bash
git clone https://github.com/grisuno/BlackObsidianC2.git
cd BlackObsidianC2
```
2. Generate SSL Certificates
```bash
# Generate private key
openssl genpkey -algorithm RSA -out key.pem -pkeyopt rsa_keygen_bits:2048

# Generate self-signed certificate
openssl req -new -x509 -key key.pem -out cert.pem -days 365 \
  -subj "/C=CL/ST=Hack/L=World/O=C2/CN=127.0.0.1"

# Convert to PKCS#8 (required by Go)
openssl pkcs8 -topk8 -nocrypt -in key.pem -out key_go.pem
```
3. Compile
```bash
go build -o c2-server
```
4. Run
```bash
export TLS_CERT=./cert.pem
export TLS_KEY=./key_go.pem
export C2_AES_KEY=18547a9428b62fdf2ba11cebc786bccbca8a941748d3acf4aad100ac65d0477f

./c2-server serve --https=0.0.0.0:4444
```
## 🚀 Quick Start
Access Dashboard
```text
https://127.0.0.1:4444/dashboard.html
```
Create Admin User
```bash
./c2-server superuser create LazyOwn@local.local LazyOwn
```
Login Credentials
Username: LazyOwn@loacl.local

Password: LazyOwn

## 📡 API Endpoints
Endpoint	Method	Description
- /login	POST	Authenticate user
- /get_connected_clients	GET	List active implants
- /pleasesubscribe/v1/users/{client_id}	GET	Fetch pending commands
- /pleasesubscribe/v1/users/{client_id}	POST	Submit command results
- /issue_command	POST	Issue command to implant
## 🔐 Security Features
Encryption
All communications use AES-256-GCM encryption:

```go
// Example encryption in handlers
encryptedData, err := AESEncrypt([]byte(sensitiveData))
```
## TLS Configuration
TLS 1.2 - 1.3

Strong cipher suites only

Certificate pinning support

## Authentication
Basic auth for initial login

Session tokens with expiration

## Superuser role management

## 📊 Project Structure
```text
black-obsidian/
├── main.go              # Entry point
├── handlers.go          # API handlers
├── crypto.go            # AES encryption functions
├── schemas.go           # Database schemas
├── web/                 # Web dashboard
│   ├── dashboard.html
│   ├── css/
│   ├── js/
│   └── img/
├── pb_data/             # PocketBase database
├── sessions/            # Client logs (CSV)
├── go.mod
├── go.sum
└── README.md
```
## 🔧 Configuration
Environment Variables
```bash
TLS_CERT           # Path to SSL certificate
TLS_KEY            # Path to SSL private key
C2_AES_KEY         # AES encryption key (64 hex chars)
POCKETBASE_DIR     # Database location (default: ./pb_data)
```
Encryption Key Generation
```bash
# Generate random 256-bit AES key
openssl rand -hex 32
# Output: 18547a9428b62fdf2ba11cebc786bccbca8a941748d3acf4aad100ac65d0477f
```
## 📝 Usage Examples

### Issue Command

```bash
python3  app.py 
```
or manual

```bash
curl -k -X POST https://127.0.0.1:4444/issue_command \
  -d "client_id=linux_go&command=whoami" \
  -H "Authorization: Bearer YOUR_TOKEN"
```
Get Connected Clients
```bash
curl -k -X GET https://127.0.0.1:4444/get_connected_clients \
  -H "Authorization: Bearer YOUR_TOKEN"
```
## 📚 Integration with LazyOwn Ecosystem
Black Obsidian is designed to work with LazyOwn implants:

```bash
# Example: Black Sand Beacon (Linux)
./black-sand-beacon \
  --c2=https://127.0.0.1:4444 \
  --client-id=linux_go \
  --aes-key=18547a9428b62fdf2ba11cebc786bccbca8a941748d3acf4aad100ac65d0477f
```
## 🛡️ OPSEC Recommendations
Use VPN/Proxy for C2 infrastructure

Rotate AES keys periodically

Enable firewall rules to restrict access

Use domain fronting for HTTPS traffic

Implement jitter in beacon callbacks

Monitor logs for anomalies

## 🤝 Contributing
Contributions welcome! Please:

Fork the repository

Create feature branch (git checkout -b feature/AmazingFeature)

Commit changes (git commit -m 'Add AmazingFeature')

Push branch (git push origin feature/AmazingFeature)

Open Pull Request

## ⚖️ Legal Disclaimer
Black Obsidian is designed for authorized penetration testing and red team operations only. Unauthorized access to computer systems is illegal. Ensure you have explicit written authorization before conducting any offensive security activities.

## 📜 License
This project is licensed under the GPLv3 - see LICENSE file for details.

## 👨‍💻 Author
grisun0 - LazyOwn Red Team Operator & Security Researcher

GitHub: @grisuno

Twitter: @lazyown.redteam

Medium: @lazyown.redteam

## 🔗 Related Projects

- [LazyOwn - Full RedTeam Framework](https://github.com/grisuno/LazyOwn)

- [Black Basalt Beacon - Windows Implant](https://github.com/grisuno/beacon)

- [Black Sand Beacon - Linux Implant](https://github.com/grisuno/blacksandbeacon)

- [Black Serpentine - Python Implant](https://github.com/grisuno/BlackSerpentine)

- [Maleable Implant](https://github.com/grisuno/LazyOwn)

- [LazyLoader](https://github.com/grisuno/LazyLoader/)

- [CompressLoader Only for Patreons](https://github.com/grisuno/CompressLoader)

- [Pro Bof's for Black Basalt Beacon Only for patreons](https://github.com/grisuno/probofs) loader.x64.o loader_fluc.x64.o shellcodeloader.x64.o sudo.x64.o getsystem.x64.o (and their soources)

## 📞 Support
For issues, questions, or suggestions:

Open an Issue

Start a Discussion

Built with ⬛ obsidian by LazyOwn Red Team operators, for the world Red Team operators.

## ☕  Ko-fi: [ko-fi.com/grisuno](https://ko-fi.com/grisuno) (Buy me coffee. I’ll use it to compile more BOFs that vanish mid-execution.)
## 💛 Patreon: [https://www.patreon.com/c/LazyOwn](https://www.patreon.com/c/LazyOwn)
## 🎙️ Podcast: [https://www.podbean.com/ew/pb-gyy75-199a35d](https://www.podbean.com/ew/pb-gyy75-199a35d)
## 🗄️ Hashnode: [https://lazyown.hashnode.dev/](https://lazyown.hashnode.dev/)
## ✒️ Medium: [https://medium.com/@lazyown.redteam](https://medium.com/@lazyown.redteam)
## 📡 Reddit: [https://www.reddit.com/r/LazyOwn/](https://www.reddit.com/r/LazyOwn/)
## 🎞️ Youtube: [https://www.youtube.com/@KillerMonkyRecordz](https://www.youtube.com/@KillerMonkyRecordz)

![Python](https://img.shields.io/badge/python-3670A0?style=for-the-badge&logo=python&logoColor=ffdd54) ![Shell Script](https://img.shields.io/badge/shell_script-%23121011.svg?style=for-the-badge&logo=gnu-bash&logoColor=white) ![Flask](https://img.shields.io/badge/flask-%23000.svg?style=for-the-badge&logo=flask&logoColor=white) [![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)

[![ko-fi](https://ko-fi.com/img/githubbutton_sm.svg)](https://ko-fi.com/Y8Y2Z73AV)


---
### Grisuno Offensive Security Ecosystem
This tool is part of a broader, synergistic RedTeam workflow:
- [LazyOwn](https://github.com/grisuno/LazyOwn): RedTeam/APT framework with AI-powered C&C, rootkits and malleable implants (Windows/Linux/Mac).
- [LazyOwnBT](https://github.com/grisuno/LazyOwnBT): Advanced complementary toolkit for BlueTeam professionals.
- [Lazymapd](https://github.com/grisuno/Lazymapd): Fast, customizable port scanner for firewall evasion.
