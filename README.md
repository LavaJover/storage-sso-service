# SSO Authentication Service
This Single Sign-On service was developed for safe and rapid authentication and authorization

## Capability
- Users registration and login
- Access-JWT generation and validation
- Refresh-JWT generation and validation
- Updating clients access-JWT with refresh-JWT
- Safe password storing (encrypted)
- Using Postgres

## Technologies
- **Golang**
- **PostgreSQL**
- **JWT**
- **Docker**

---

## Installation
### Clone git repo
```sh
git clone https://github.com/LavaJover/storage-sso-service.git
cd storage-sso-service
```
### Install dependencies
```sh
go mod tidy
```
### Setup environment and run
```sh
export SSO_CONFIG_PATH="path/to/config/file"
cd storage-sso-service/sso-service
go run cmd/sso-service/main.go
```
