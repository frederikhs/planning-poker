FROM node:alpine as development
USER 1000
WORKDIR /web

CMD npm install && npm run start

FROM development as intermediate
USER root
COPY --chown=node:node package*.json .
RUN npm ci

COPY --chown=node:node . .
RUN npm run build

FROM nginx:alpine as production
COPY server.conf /etc/nginx/conf.d/nginx.conf
COPY --from=intermediate --chown=nginx:nginx /web/build /usr/share/nginx/html
