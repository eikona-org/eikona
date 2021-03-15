# ImgProcessing
Welcome to the DS1 Challenge Project: ImageProcessing

## Infrastructure

//TODO Image

## Usage
## 1. Clone
```bash
git clone https://github.com/imgProcessing/???
```
## 2.1 mkcert (developing only)
Follow the instructions to get mkcert: https://github.com/FiloSottile/mkcert
```bash
# If it's the firt install of mkcert, run
mkcert -install

# Generate certificate for domain "docker.localhost", "domain.local" and their sub-domains
mkcert -cert-file infrastructure/ssl/cert.pem -key-file infrastructure/ssl/key.pem \
"docker.localhost" "*.docker.localhost" "domain.local" "*.domain.local"
```

## 2.2 copy cert (prod only)
```bash
# Copy your certs:
cp cert.pem key.pem infrastructure/ssl/
```

## 3. Environment
Rename the *.env.sample files to *.env and edit them.
```bash

```
## 4. Run
Scale 2 instances of frontend and 2 instances of backend
```bash
docker-compose up --scale frontend=2 --scale backend=2
```

### 4.1 Run on dev system
Scale 2 instances of frontend and 2 instances of backend
```bash
docker-compose -f development.yml up
```
## 5. Access
Traefik Dashboard: http://localhost:8080/dashboard/#/

Backend: https://backend.docker.localhost/

Frontend: https://frontend.docker.localhost/

Keep in mind that live reload with air is currently not working on Windows due to filesystem limitations.

Runs the app with DB and MinIO\
Open [http://localhost:8080](http://localhost:8080) to view the Backend in the browser.\
Open [http://localhost:9000](http://localhost:9000) to view the MinIO in the browser.\
Connect to [http://localhost:5432](http://localhost:5432) to connect to the DB
