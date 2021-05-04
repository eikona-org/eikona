<h1 align="center">Welcome to εικόνα 👋</h1>
<p>
  <img alt="Version" src="https://img.shields.io/badge/version-1.0-blue.svg?cacheSeconds=2592000" />
  <a href="https://github.com/eikona-org/eikona/blob/main/LICENSE" target="_blank">
    <img alt="License: Apache License 2.0" src="https://img.shields.io/badge/License-Apache License 2.0-yellow.svg" />
  </a>
</p>

> εικόνα resizes, adjust quality, convert images on-the-fly and provides a simple REST API. This project was built for the OST DS1 Challenge Task.


⚠️ This is a school project and not production-ready! Run at your own risk.

### 🏠 [Homepage](https://github.com/eikona-org/eikona)

### ✨ [Demo](https://eikona.pesc.xyz/)

## Infrastructure

![Infrastructure overview](https://github.com/eikona-org/eikona/blob/main/doc/images/Infrastruktur.png)
## Running in Production

```sh
git clone https://github.com/eikona-org/eikona
```
```sh
# Copy your certs - TODO: Adding Letsencrypt to Traefik:
cp cert.pem key.pem infrastructure/ssl/
```
```sh
#Rename and adjust ALL env files in the root
cp backend.env.sample backend.env
vim backend.env
```
```sh
# Running
docker-compose -f docker-compose.yml -f docker-compose.prod.yml up -d
```

Access: `https://[domain]/`

## Running in Development

```sh
git clone https://github.com/eikona-org/eikona
```
Follow the instructions to get mkcert: https://github.com/FiloSottile/mkcert
```sh
# If it's the firt install of mkcert, run
mkcert -install

# Generate certificate for domain "docker.localhost", "domain.local" and their sub-domains
mkcert -cert-file infrastructure/ssl/cert.pem -key-file infrastructure/ssl/key.pem \
"docker.localhost" "*.docker.localhost" "domain.local" "*.domain.local"
```
```sh
#Rename and adjust ALL env files in the root
cp backend.env.sample backend.env
vim backend.env
```
```sh
# Running
docker-compose up
```
Generating Swagger Docs
Run `go get -u github.com/swaggo/swag/cmd/swag` in the `backend` folder and then `swag init`

Access:
```sh
LB: https://docker.localhost/

Traefik Dashboard: https://docker.localhost:8082/dashboard/#/

Backend: http://docker.localhost:8081/

Frontend: http://docker.localhost:8080/

MinIO: http://docker.localhost:9000/

Database: docker.localhost:5432
```


## Author

👤 **Lukas Ribi, Dominik Castelberg, Pascal Christen**

* Github: [@lribi](https://github.com/lribi) [@dcastelberg](https://github.com/dcastelberg) [@pesc](https://github.com/pesc))

## 🤝 Contributing

Contributions, issues and feature requests are welcome!<br />Feel free to check [issues page](https://github.com/eikona-org/eikona/issues). 

## Show your support

Give a ⭐️ if this project helped you!

## 📝 License

Copyright © 2021 [Lukas Ribi](https://github.com/lribi), [Dominik Castelberg](https://github.com/dcastelberg), [Pascal Christen](https://github.com/pesc).<br />
This project is [Apache License 2.0](https://github.com/eikona-org/eikona/blob/main/LICENSE) licensed.

***
_This README was generated with ❤️ by [readme-md-generator](https://github.com/kefranabg/readme-md-generator)_
