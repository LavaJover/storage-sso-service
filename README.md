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
```sh
git clone https://github.com/LavaJover/storage-sso-service.git
export SSO_CONFIG_PATH="path/to/config/file"
cd storage-sso-service/sso-service
go run cmd/sso-service/main.go
```
