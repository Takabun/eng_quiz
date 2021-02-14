FROM node:10.23.2-alpine3.10

WORKDIR /usr/src
COPY ./package.json .
RUN yarn install
COPY ./ .

RUN yarn build

ENV HOST 0.0.0.0

CMD ["yarn", "start"]

EXPOSE 3000
