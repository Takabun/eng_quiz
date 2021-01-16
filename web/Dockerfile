FROM node:12.16.1-alpine
# EXPOSE 3000  #なくても問題ない
ENV HOST 0.0.0.0

# WORKDIR /app
# RUN apk update

WORKDIR /usr/src
# ↑こうしたら"ts-loader"の件が解決

COPY ./package.json .
RUN yarn install
COPY ./ .

ENV HOST 0.0.0.0

CMD [ "yarn", "dev"]