FROM node:18-alpine

WORKDIR /usr/src/Frontend

COPY package.json package-lock.json ./

RUN npm install

COPY . .

EXPOSE 5173

CMD [ "npm", "run", "dev" ]