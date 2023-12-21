FROM ubuntu:22.04
RUN apt-get update && apt-get install -y sudo curl python3 build-essential

ENV NVM_DIR /usr/local/nvm
ENV NODE_VERSION v18.19.0
RUN mkdir -p /usr/local/nvm && apt-get update && echo "y" | apt-get install curl
RUN curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.1/install.sh | bash
RUN /bin/bash -c "source $NVM_DIR/nvm.sh && nvm install $NODE_VERSION && nvm use --delete-prefix $NODE_VERSION"
ENV NODE_PATH $NVM_DIR/versions/node/$NODE_VERSION/bin
ENV PATH $NODE_PATH:$PATH

RUN npm install -g yarn
ADD . /fast-updates
WORKDIR /fast-updates
RUN npm cache clean --force
RUN rm -r node_modules || true
RUN rm yarn.lock || true
RUN yarn install

RUN yarn c
ENV CHAIN_CONFIG="docker"
ENV DEPLOYER_PRIVATE_KEY="0xc5e8f61d1ab959b397eecc0a37a6517b8e67a0e7cf1f4bce5591f3ed80199122"
ENTRYPOINT sleep 5 && yarn hardhat chain-config --network docker && yarn hardhat deploy-contracts --network docker
