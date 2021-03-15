# Infrastructure
This is the infrastructure part of the imgProcessing Service

## 1. Clone
```bash
git clone https://github.com/imgProcessing/infrastructure.git
```
## 2. mkcert (developing only)
Follow the instructions to get mkcert: https://github.com/FiloSottile/mkcert
```bash
# If it's the firt install of mkcert, run
mkcert -install

# Generate certificate for domain "docker.localhost", "domain.local" and their sub-domains
mkcert -cert-file ssl/local-cert.pem -key-file ssl/local-key.pem \
"docker.localhost" "*.docker.localhost" "domain.local" "*.domain.local"
```

## 3. Environment
Rename the *.env.sample files to *.env and edit them.

## 4. Run
Scale 2 instances of frontend and 2 instances of backend
```bash
docker-compose up --scale frontend=2 --scale backend=2
```
## 5. Access
http://localhost:8080/dashboard/#/
https://backend.docker.localhost/
https://frontend.docker.localhost/

