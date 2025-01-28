FROM node:23

WORKDIR /frontend

COPY package*.json ./

RUN npm ci

COPY . .
