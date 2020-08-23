FROM node:12.5.0-alpine
# EXPOSE 3000  #なくても問題ない
ENV HOST 0.0.0.0

WORKDIR /usr/src

COPY ./package.json .
RUN yarn install && yarn cache clean