# eikona
Welcome to the DS1 Challenge Project: eikona

## Infrastructure

![Infrastructure overview](https://github.com/eikona-org/eikona/blob/main/doc/images/Infrastruktur.png)

## Usage Prod
## 1. Clone
```bash
git clone https://github.com/eikona-org/eikona
```
## 2. copy cert - TODO: Let's encrypt
```bash
# Copy your certs:
cp cert.pem key.pem infrastructure/ssl/
```
## 3. Environment
Rename the *.env.sample files to *.env and edit them.
```bash

```
## 4. Run
```bash
docker-compose -f docker-compose.yml -f docker-compose.prod.yml up -d
```
## 5. Access

LB: https://[domain]/

## Usage Dev

## 1. Clone
```bash
git clone https://github.com/eikona-org/eikona
```
## 2. mkcert
Follow the instructions to get mkcert: https://github.com/FiloSottile/mkcert
```bash
# If it's the firt install of mkcert, run
mkcert -install

# Generate certificate for domain "docker.localhost", "domain.local" and their sub-domains
mkcert -cert-file infrastructure/ssl/cert.pem -key-file infrastructure/ssl/key.pem \
"docker.localhost" "*.docker.localhost" "domain.local" "*.domain.local"
```
## 3. Environment
Rename the *.env.sample files to *.env and edit them.
```bash

```
## 4. Run
```bash
docker-compose up
```
## 5. Generate docs
Run `go get -u github.com/swaggo/swag/cmd/swag` in the `backend` folder and then `swag init`

## 6. Access

LB: https://docker.localhost/

Traefik Dashboard: http://docker.localhost:8082/dashboard/#/

Backend: http://docker.localhost:8081/

Frontend: http://docker.localhost:8080/

MinIO: http://docker.localhost:9000/

Database: docker.localhost:5432
