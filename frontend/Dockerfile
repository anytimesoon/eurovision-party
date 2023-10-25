FROM node:19 as frontend
ENV NODE_ENV=production


WORKDIR /app

COPY package.json ./
COPY package-lock.json ./
RUN npm install
COPY . ./
RUN npm run build


FROM node:19-alpine3.16

WORKDIR /app
COPY --from=frontend /app .


ENV HOST=0.0.0.0
EXPOSE 3000
CMD ["node", "build"]