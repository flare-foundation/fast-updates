# Fast Updates Visualizer

## Install

Using [poetry](https://python-poetry.org)

```bash
poetry install --no-root
```

## Run

```bash
poetry run python -m app -l ../logs
```

Internally the app looks for the logs folder and starts a server on a port (default: 8051).

To modify either of these use the `-l` and `-p` flags.

```bash
poetry run python -m app -l <path/to/logs-folder> -p 8055
```
