FROM alpine:node as BUILDER

WORKDIR /app
COPY ./ /app/ 
RUN npm install && npm run build


FROM nginxinc/nginx-unprivileged:1.23-alpine

WORKDIR /usr/share/nginx/html

COPY --101:101 --from BUILDER  