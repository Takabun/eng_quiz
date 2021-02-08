# FROM node:12.16.1-alpine
# # EXPOSE 3000  #なくても問題ない
# ENV HOST 0.0.0.0

# # WORKDIR /app
# # RUN apk update

# WORKDIR /usr/src
# # ↑こうしたら"ts-loader"の件が解決

# COPY ./package.json .
# RUN yarn install
# COPY ./ .

# ENV HOST 0.0.0.0

# CMD [ "yarn", "dev"]








# FROM node:12.16.1-alpine
# FROM node:9.2.0
FROM node:10.23.2-alpine3.10

# WORKDIR /todoweb
# COPY . /todoweb
WORKDIR /usr/src
COPY ./package.json .
RUN yarn install
COPY ./ .

# RUN npm install -g vue-cli@2.9.3
# RUN npm install
RUN yarn build

ENV HOST 0.0.0.0

CMD ["yarn", "start"]

EXPOSE 3000
