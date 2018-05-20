# Echo server example

Example of listening events Ethereum events in Golang

## Run

`go run cmd/pongWatcher/main.go --raw-url=wss://ropsten.eth.6120.eu/ws --address=0x85EAb3977d0C1a6F22Fd0a3e37090a234551D2E0`

![go run](https://raw.githubusercontent.com/Magicking/example-event-solidity2go/master/docs/run.gif)

## Generate wrapper

From base directory: `go generate`

## Dependencies

For wrapper generation: [abigen](https://github.com/ethereum/go-ethereum#executables) `go get github.com/ethereum/go-ethereum/cmd/abigen`

## Example code

```golang
	c, err := ethclient.Dial(opts.RawURL)
[...]
	ppf, err := examples.NewPingPongFilterer(addr, c)
[...]
	sub, err := ppf.WatchPong(&bind.WatchOpts{nil, nil}, sink)
[...]
	for e := range sink {
		log.Printf("PingPongPong: %+v", e)
	}
```

## Donation

BTC: 1MYiMU3GfsgEP4EYHHonG9Fy6DkA1JC3B5
ETH: 0xc8f8371BDd6FB64388F0D65F43A0040926Ee38be
BCC: bitcoincash:qphs9n0apf553jgwawgrx24wfrnutjse45av6k5h8x
