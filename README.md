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
mkcert -cert-file infrastructure/ssl/local-cert.pem -key-file infrastructure/ssl/local-key.pem \
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

Keep in mind that live reload with air is currently not working on Windows due to filesystem limitations.

Runs the app with DB and MinIO\
Open [http://localhost:8080](http://localhost:8080) to view the Backend in the browser.\
Open [http://localhost:9000](http://localhost:9000) to view the MinIO in the browser.\
Connect to [http://localhost:5432](http://localhost:5432) to connect to the DB
