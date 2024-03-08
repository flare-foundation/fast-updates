FROM node:20
WORKDIR /fast-updates
COPY . .
RUN corepack enable
RUN yarn install --immutable
RUN yarn compile
ENV NETWORK="docker"
ENTRYPOINT [] 
