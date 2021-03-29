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
ENTRYPOINT ["nginx", "-g", "daemon off;"]
