# Contributing

Thank you for considering improving our source code.
All contributions are welcome.

## Issues

_Sensitive security-related issues should not be reported publicly.
See [Security](SECURITY.md)._

To share ideas, considerations, or concerns open an issue.
Before filing an issue make sure the issue has not been already raised.
In the issue, answer the following questions:

- What is the issue?
- Why it is an issue?
- How do you propose to change it?

### Pull request

Before opening a pull request open an issue on why the request is needed.

To contribute: fork the repo, make your improvements, commit, and open a pull request.
The maintainers will review the request.

The request must:

- Reference the relevant issue,
- Follow standard golang guidelines,
- Be well documented,
- Be well tested,
- Compile,
- Pass all the tests,
- Pass all the linters,
- Be opened against the `main` branch.

## Setting the environment

The Go client lives in [go-client](go-client/). Make sure you are using go with higher or equal version to the one specified in [go-client/go.mod](go-client/go.mod).

Get all the dependencies

```bash
cd go-client
go mod tidy
```

### Tests

Run unit tests with

```bash
cd go-client
make test-unit
```

Run the end-to-end scenario suite (requires Docker) with

```bash
cd go-client
make test-e2e
```

See the [Tests section of the go-client README](go-client/README.md#tests) for details on what the suites cover, runtime tuning, and Docker platform requirements.

### Linting

Run linters (make sure you have [golangci-lint](https://golangci-lint.run/) installed) with

```bash
cd go-client
golangci-lint run
```

## AI Assistance

Any significant use of the AI assistance in the contribution MUST be disclosed in the pull request along with the extent of the use.

An example disclosure:

> This PR was written primarily by Claude Code.

Or a more detailed disclosure:

> I consulted ChatGPT for the following code snippets: ...
