FROM node:alpine

COPY package*.json .

RUN npm i

COPY . .

CMD time node app.js
