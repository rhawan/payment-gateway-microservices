FROM node:14.17-alpine

WORKDIR /home/node/app

COPY package*.json ./

RUN npm install

COPY . .

RUN npm run build

#non-root === seu usuario
# USER node 

EXPOSE 3001

CMD [ "npm", "start" ]