# FTSO Fast Updates

This repo contains MVP of an implementation of the new Fast Updates proposal.

# Setup

## Local install

### Node.JS

-   Install [NVM](https://github.com/nvm-sh/nvm).
-   Install Node v18 (LTS):
    ```
    nvm install 18
    ```
-   Set version 18 as default:
    ```
    nvm alias default 18
    ```

### Project

-   Install `yarn` and dependencies:
    ```
    npm install -g yarn
    yarn install
    ```
-   To compile smart contracts run:
    ```
    yarn c
    ```

## Docker

Instead of installing a local version, one can use docker. To build a docker image, run

```
docker build . -t fast-updates
```

# Tests

-   To run all tests run:

```
yarn test
```

-   With docker (assuming the image `fast-updates` was build), move to `test` folder and run docker compose:

```
cd test
docker compose up
```

# Simulation

A simulation of the protocol using docker is available. It includes deploying and setting a chain node,
deploying contracts, running a daemon, and running multiple fast updates providers. Navigate to
`deployment/simulation` (assuming that the docker image `fast-updates` was build, as explained above)
and run docker compose:

```
cd deployment/simulation
docker compose up
```

# Development

Recommended editor to use is [VSCode](https://code.visualstudio.com/).

## Code formatting

We use `Prettier` for code formatting, with settings defined under `package.json`.

You can install the VSCode extension and use the shortcut `Alt/Option` + `Shift` + `F` to auto-format the current file.
