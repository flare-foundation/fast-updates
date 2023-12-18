FROM ubuntu:22.04
RUN apt-get update && apt-get install -y sudo curl 

ENV NVM_DIR /usr/local/nvm
ENV NODE_VERSION v18.0.0
RUN mkdir -p /usr/local/nvm && apt-get update && echo "y" | apt-get install curl
RUN curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.1/install.sh | bash
RUN /bin/bash -c "source $NVM_DIR/nvm.sh && nvm install $NODE_VERSION && nvm use --delete-prefix $NODE_VERSION"
ENV NODE_PATH $NVM_DIR/versions/node/$NODE_VERSION/bin
ENV PATH $NODE_PATH:$PATH

RUN npm install -g yarn
ADD . /fast-updates
WORKDIR /fast-updates
RUN ls
RUN npm install

ENTRYPOINT yarn hardhat node