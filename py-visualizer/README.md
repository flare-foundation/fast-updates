# FTSO Fast Updates Visualizer

## Install

Using Python 3.12 and [poetry](https://python-poetry.org)

```bash
poetry install --no-root
```

## Usage

This app enables to visualize the fast updates that an instance of the Fast Updates client is submitting and logging in the logs folder. This way the data provider can visualize the effect of its fast updates and monitor the updates of others.

It __needs a running Go client__ that is logging its actions in a file.

### Setting and running the go-client

The go client first needs a feed provider from which the visualizer will collect data. An example feed provider can be started with the following command (from `../go-client/tests`):

```bash
docker compose up value-provider
```

The go client can be started with the following command (from `../go-client`) after setting up the configuration file `config.toml`:

```bash
go run main.go --config config.toml
```

By default the go-client will log current onchain and feed provider values
every time it is eligible to submit a fast update. Add `feed_values_log` value
in the configuration file to set how often should the client log off-chain and
onchain values. For example, set

```toml
[logger]
level = "INFO"
file = "./logger/logs/fast_updates_client.log"
console = true
feed_values_log = 5
```
to log values every 5 blocks.

Once the client is running and a log is being created, we can visualize it with the python app.

## Run

```bash
poetry run python -m app -l ../go-client/logger/logs
```

Internally the app looks for the logs folder and starts a server on a port (default: 8051).

To modify either of these use the `-l` and `-p` flags.

```bash
poetry run python -m app -l <path/to/logs-folder> -p 8055
```