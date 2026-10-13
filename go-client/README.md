# FTSO Fast Updates Client

This repository contains an implementation of a FTSO Fast Updates client that can be used to submit
fast updates to the [FTSOv2 contracts](https://github.com/flare-foundation/flare-smart-contracts-v2).
The FTSO Fast Updates Client is implemented to connect to a RPC blockchain node
and continuously generate verifiable random numbers, based on a sortition private key and current
block number, to determine if it can submit a price update. In the latter case, it makes a transaction
from one of the provided accounts calling the SubmitUpdates function on the
[FastUpdater contract](https://github.com/flare-foundation/flare-smart-contracts-v2/blob/main/contracts/fastUpdates/implementation/FastUpdater.sol).
Multiple accounts can be provided for the submissions in the case of multiple simultaneous
fast updates, to not miss the submission window.

## Prerequisites

The client is implemented in Go (requires version 1.27 or newer, see `go.mod`).

## Configuration

The configuration is read from a `toml` file and/or the environment variables. Config file can be specified using the
command line parameter `--config`, e.g., `./fast-updates-client --config config.toml`. The default config file name is `config.toml`.
Here is a list of example configuration parameters. Note that to participate as a fast updater, one must be registered at the
FTSOv2 Top Level system (see [link](https://github.com/flare-foundation/ftso-v2-provider-deployment?tab=readme-ov-file#register-accounts)),
where it registers its **sortition private key** (see
below on how to generate the sortition key).

Set private information in environment variables

```bash
# voters private key registered in the VoterRegistry, aka signingPolicy private key
SIGNING_PRIVATE_KEY="0xd49743deccbccc5dc7baa8e69e5be03298da8688a15dd202e20f15d5e0e9a9fb"
# voters sortition key registered in the VoterRegistry contract that enables generating verifiable
# randomness to determine the order of clients submitting the fast updates
SORTITION_PRIVATE_KEY="0xd49743deccbccc5dc7baa8e69e5be03298da8688a15dd202e20f15d5e0e9a9fb"
# private keys of accounts from which the fast updates will be
# submitted - the client needs multiple addresses to not miss the
# submission window in case multiple fast updates can be submitted
# for blocks in a short interval
ACCOUNTS="0xd49743deccbccc5dc7baa8e69e5be03298da8688a15dd202e20f15d5e0e9a9fb,0x23c601ae397441f3ef6f1075dcb0031ff17fb079837beadaf3c84d96c6f3e569,0xee9d129c1997549ee09c0757af5939b2483d80ad649a0eda68e8b0357ad11131"
```

```toml
[client]
# address of the FastUpdater contract
fast_updater_address = "0xbe65A1F9a31D5E81d5e2B863AEf15bF9b3d92891"
# address of the Submission contract to which the updates are sent
submission_address = "0x18b9306737eaf6E8FC8e737F488a1AE077b18053"
# address of the top level FlareSystemManager contract
flare_system_manager = "0x919b4b4B561C72c990DC868F751328eF127c45F4"
# address of the FastUpdatesIncentiveManager contract
incentive_manager_address = "0x919b4b4B561C72c990DC868F751328eF127c45F4"
# parameter defining when a fast update can be submitted
submission_window = 10
# url of the off-chain data provider
value_provider_url = "http://127.0.0.1:3101/feed-values/0"
# (optional) when (off-chain) fetching feeds values, an address with sufficient
# balance should be provided, even though the transactions never happen;
# the below address is the default one, which on flare chains has enough balance
fetch_current_feeds_address = "0x000000000000000000000000000000000000dEaD"
fetch_current_feeds_value = "1000000000000000000000000"

[transactions]
gas_limit = 8000000
value = 0
# Submissions are EIP-1559 (type 2). gas_fee_cap = baseFee * base_fee_multiplier + gas_tip_cap,
# gas_tip_cap = clamp(baseFee * max_priority_fee_multiplier, minimal_max_priority_fee, maximal_max_priority_fee).
# All optional; defaults: base_fee_multiplier=4, max_priority_fee_multiplier=2,
# minimal_max_priority_fee=100 Gwei, maximal_max_priority_fee=5000 Gwei.
base_fee_multiplier = 4
max_priority_fee_multiplier = 2
minimal_max_priority_fee = 100_000_000_000
maximal_max_priority_fee = 5000_000_000_000


[logger]
level = "DEBUG"
# logs go to standard output, format is "console" (default) or "json"
format = "console"
# when the balance (in WEI) of the provided accounts falls bellow this value
# the logger will show warnings
min_balance = 10000000000000000000

[chain]
node_url = "https://coston2-api.flare.network/ext/C/rpc"
# optional rpc api key, can also be set via API_KEY env variable
api_key = ""
chain_id = 114
```

The private key, sortition private key, and accounts private keys can also be set in
the configuration file, but we strongly suggest to set them using environment variables
to avoid accidentally exposing them.

## Price Updates Provider

To provide meaningful updates of prices/feeds on the FastUpdates contract a user needs
to implement its own updates logic. For this it needs to implement a struct with
functionality that fits the following interface.

```go
type ValuesProvider interface {
	GetValues(feeds []FeedId) ([]*float64, error)
}
```

See `provider/feed_provider.go` for the interface. A `nil` entry in the
returned slice marks a feed the provider does not support; the client then
never submits updates for it. We provide an implementation of obtaining
feeds values based on API calls to the [FTSOv2 Example value provider](https://github.com/flare-foundation/ftso-v2-example-value-provider),
see `provider/http_feed_provider.go` (the example value provider can also
serve fixed or random test values via its `VALUE_PROVIDER_IMPL` setting).
When a custom implementation is provided, one can define in `main.go` which
price provider will be used.

## Running the FTSO Fast Updates Client

Assuming that the configuration file and configuration environment variables were set and
the provider is registered, simply run

```bash
go run main.go --config config.toml
```

or build and run the binaries with

```bash
go build .
./fast-updates-client --config config.toml
```

## Handling failed transactions

The FTSO Fast Updates Client is implemented to submit fast updates to the
specified FastUpdater contract through the Submission contract.
In file `client/transaction_queue.go` there is an implementation of a queue
that accepts tasks (transaction requests
for fast updates) and executes them on parallel threads. In the case of
a failed transaction, which can happen for multiple reasons such as failed
connection to the RPC node, missed submission block, etc., an error handler
is implemented, see

```go
func (txQueue *TransactionQueue) ErrorHandler()
```

function in the file `client/transaction_queue.go`. Currently, in the
case of an error, the client only logs the error and dismisses the
transaction. Alternative actions, such as resubmitting the transaction
can be implemented, but this can be risky since the submission window
might already be closed.

## Generating sortition key and signing

To participate in the Fast Updates protocol, a client needs to have a private/public sortition
key, that needs to be registered at the top level FTSO voter registry.

To generate the key one can run

```bash
go run keygen/keygen.go
```

and the key should be printed in the console. Additionally, to register the public key at the
top level FTSO voter registry, one needs to provide a proof of the correctness of the public
key. For this a user needs to sign its _entity_ address. Use

```bash
go run keygen/keygen.go --key 0x1512de600a10a0aac01580dbfc080965b89ed2329a7b2bf538f4c7e09e34aa1 --address 0xd4e934C2749CA8C1618659D02E7B28B074bf4df7
```

where the key value needs to be replaced by the generated private key and the address value needs
to be replaced by the actual _entity_ address.

In case you forgot your public key, but saved the private one, you can run the following for outputting the key:

```bash
go run keygen/keygen.go --key 0x1512de600a10a0aac01580dbfc080965b89ed2329a7b2bf538f4c7e09e34aa1
```

Alternatively, one can save and read the (encrypted) key from a file and save the signature with:

```bash
go run keygen/keygen.go --key_out keys.out --pass secret_password # generate and save a key to file
go run keygen/keygen.go --key_file keys.out --pass secret_password --address 0xd4e934C2749CA8C1618659D02E7B28B074bf4df7 --sig_out sig.out # read and decrypt a key from a file, sign and write the signature to a file
go run keygen/keygen.go --key_file keys.out --pass secret_password # read and decrypt a key from file, just print it out
```

where the address value needs to be replaced by the actual address that will be used to sign
the updates. The signature should be saved in the specified file. Signing of the entity address
is provided also in [this registration script](https://github.com/flare-foundation/flare-smart-contracts-v2/blob/main/deployment/tasks/register-public-keys.ts#L38).

## Tests

### End to end scenario tests

The `tests/e2e` package contains a scenario test suite that exercises the
whole client against the real Fast Updates contracts on a local Ganache chain
(managed automatically via `docker`), with the Flare system mocked by
`FlareSystemMock` and feed values provided by an in-process,
scenario-controlled provider. The chain mines a block every 2 seconds by
default — sortition alone costs ~1s per block at the test weight, so a
faster chain outruns the client on slower machines; on a fast idle machine
`FAST_UPDATES_E2E_BLOCK_TIME=1` roughly halves the runtime.

The suite is built to catch subtle bugs in the submitted price deltas. After
each scenario run, a verifier replays the chain block by block, independently
of the client code, and checks that

- every feed only ever moves toward the value the provider was reporting
  (allowing for the volatility-band jitter near the target),
- feeds without a provider value, delisted feeds, and feeds already at their
  target never move at all,
- no submission transaction reverted, and every update the client reports as
  successful is actually mined on chain,
- all feeds end the run within a few scale steps of the provider values.

Every feed gets a distinct target direction and magnitude, so updates
applied to the wrong feed index fail the run immediately. Scenarios
cover convergence up/down, feeds unsupported by the provider (index
alignment), a delisted (zeroed) feed id, and a feed removed and replaced
by a new feed id at the same index while the client is running (the
client must switch to the new feed id after its periodic feed refresh).

```bash
make test-e2e   # or: go test ./tests/e2e/ -v -count=1 -timeout 30m
```

Each scenario boots a fresh chain and takes a few minutes (~12 minutes for
the whole suite at the default block time). The suite is skipped in `-short`
mode or when `docker` is unavailable.

Note that Ganache has no working arm64 build (its arm64 image variant ships
an amd64 binary and upstream is discontinued), so the suite always runs the
amd64 image and needs a host able to execute it: any amd64 machine, or an
ARM host with emulation (Docker Desktop on Apple Silicon supports this via
Rosetta; on Linux ARM install binfmt, e.g.
`docker run --privileged --rm tonistiigi/binfmt --install amd64`). Without
emulation the suite skips with a message explaining this.

### Unit tests

```bash
make test-unit   # or: go test ./client/ ./provider/ ./sortition/ -short -count=1
```

The delta computation and packing logic (`provider.GetDeltas`,
`provider.StringToDeltas`) and the on-chain value conversion
(`client.RawChainValuesToFloats`) have table-driven tests covering update
directions, the volatility band, missing-feed index mapping, packing
boundaries and negative decimals.

### Simulation of the Fast Updates protocol using multiple clients on a local Hardhat node

Using the repository `flare-smart-contracts-v2` found [here](https://github.com/flare-foundation/flare-smart-contracts-v2/tree/main)
one can deploy all the Fast Updates contracts
together with the whole Flare system and voter repository. Navigate to the
repository and run

```bash
yarn install
yarn compile
yarn sim-node & yarn sim-run
```

This will start a Flare system on a local Hardhat node, register a couple of
data providers and start a simulation of FTSO v2 feed providers.
To additionally run a simulation of Fast Updates providers navigate to
`go-client/tests/` folder and run feed value provider with 3 clients using Docker:

```bash
docker compose up value-provider client1 client2 client3
```

The configuration files `tests/configs/config1.toml`, `tests/configs/config2.toml`,
and `tests/configs/config3.toml` should be set so that the clients can participate
in the protocol.

## Compiled ABI of Flare-Smart-Contracts-v2

The Fast Updates Go client uses an interface to the contracts that was compiled using `solc` compiler and `abigen` tool. In the case
that the contracts are changed, the interface needs to be changed as well. Use the provided
`Makefile` to compile the new interfaces. In the makefile set the destination of the `flare-smart-contracts-v2` repo and run

```bash
make compile
```
