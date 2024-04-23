# FTSO Fast Updates Simulation

This repo contains MVP of an implementation of the new Fast Updates proposal.

## Setup

### Node.JS

Install [NVM](https://github.com/nvm-sh/nvm), and Node v18 (LTS):

```bash
nvm install 18
```

### Yarn

Install dependencies with `yarn`:

```bash
corepack enable
yarn install
```

### Project

To compile smart contracts:

```bash
yarn compile
```

To format with prettier:

```bash
yarn format
```

To lint with eslint:

```bash
yarn lint
```

## Docker

To build a docker image, run:

```bash
docker build -t fast-updates .
```

To test the protocol using docker, navigate to the `test` directory and run docker compose:

```bash
cd test
docker compose up
```

## Simulating locally

Start the chain:

```bash
yarn hardhat node
```

Deploy the contracts and run the admin daemon:

```bash
yarn hardhat deploy-contracts --network localhost && yarn hardhat run-admin-daemon --network localhost
```

Run the fast updates providers:

```bash
yarn ts-node client/run-fast-updates-provider.ts $ID --network localhost
```

where `$ID` is the ID of the provider (1, 2, ...)

## Simulating with Docker

A simulation of the protocol using docker is available. It includes deploying and setting a chain node,
deploying contracts, running a daemon, and running multiple fast updates providers. Navigate to
`deployment/simulation` (assuming that the docker image `fast-updates` was build, as explained above)
and run docker compose:

```bash
cd deployment
docker compose up
```

## Notes

Strict type checked (in addition to strict mode) is enforced throughout the project (see `tsconfig.json` and `.eslintrc.js`).

For testing purposes (under `test/`), `typechain-truffle` is used to generate typechain artifacts from the contracts. The contracts are locally instantiated and used in the tests.

For deployment purposes (under `deployment/`), `typechain-web3-v1` is used to interact with the contracts.
