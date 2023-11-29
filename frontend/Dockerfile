FROM node:19 as frontend
ENV NODE_ENV=production

ARG GO_HOST
ARG GO_PORT
ARG CHAT

ENV PUBLIC_GO_HOST=$GO_HOST
ENV PUBLIC_GO_PORT=$GO_PORT
ENV PUBLIC_CHAT=$CHAT

WORKDIR /app

COPY package.json ./
COPY package-lock.json ./
RUN npm install
COPY . ./

RUN npm run build

FROM node:19-alpine3.16
LABEL maintainer="anytimesoon"
WORKDIR /app
COPY --from=frontend /app .

ENV BODY_SIZE_LIMIT=0

CMD ["node", "-r", "dotenv/config", "build"]