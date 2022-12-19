FROM node:alpine as BUILDER
WORKDIR /app
COPY ./ /app/
RUN npm install && npm run build

FROM nginxinc/nginx-unprivileged:1.23-alpine
#COPY ./nginx/nginx.conf /etc/nginx/conf.d/default.conf
COPY ./nginx/ui.conf /etc/nginx/conf.d/default.conf
COPY --from=BUILDER --chown=101:101  /app/build/ /usr/share/nginx/html