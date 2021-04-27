# build environment
FROM node:alpine as build
WORKDIR /app
COPY . .
RUN yarn
RUN yarn build
# production environment
FROM nginx:alpine
WORKDIR /usr/share/nginx/html
RUN rm -rf ./*
COPY --from=build /app/build/ ./
COPY --from=build /app/nginx/nginx.conf /etc/nginx/nginx.conf
# https://www.freecodecamp.org/news/how-to-implement-runtime-environment-variables-with-create-react-app-docker-and-nginx-7f9d42a91d70/
COPY ./env.sh .
COPY .env .
RUN apk add --no-cache bash
RUN chmod +x env.sh
ENTRYPOINT ["/bin/bash", "-c", "/usr/share/nginx/html/env.sh && nginx -g \"daemon off;\""]