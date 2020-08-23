# FROM node:12.5.0-alpine

# ENV HOME="/app" \
#     LANG=C.UTF-8 \
#     TZ=Asia/Tokyo 

# WORKDIR ${HOME}

# RUN apk update && \
#     apk upgrade && \
#     yarn add @vue/cli

# ENV HOST 0.0.0.0
# EXPOSE 3000



FROM node:12.5.0-alpine
EXPOSE 3000
ENV HOST 0.0.0.0

WORKDIR /usr/src
COPY ./package.json .
RUN yarn install && yarn cache clean