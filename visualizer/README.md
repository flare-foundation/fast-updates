# Fast Updates Visualizer

## Install

Using [poetry](https://python-poetry.org)

```bash
poetry install --no-root
```

## Run

```bash
poetry run python -m app
```

Internally the app looks for the logs folder (default `../logs`) and starts a server on port 8051.

To modify either of these use the `-l` and `-p` flags.

```bash
poetry run python -m app -l <path/to/logs-folder> -p 8055
```
