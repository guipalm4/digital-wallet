FROM node:18

WORKDIR /usr/src/app
COPY package*.json ./

RUN npm install

COPY . .
COPY ./wait-for-it.sh /usr/bin/wait-for-it

EXPOSE 3003

CMD ["wait-for-it", "kafka:29092", "--", "npm", "run", "start"]
