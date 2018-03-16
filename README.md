# Echo server example

Example for listen to event from outside of Geth

## Run

`go run cmd/pongWatcher/main.go --raw-url=wss://ropsten.eth.6120.eu/ws --address=0x85EAb3977d0C1a6F22Fd0a3e37090a234551D2E0`

![go run](https://raw.githubusercontent.com/Magicking/example-event-solidity2go/master/docs/run.gif)

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
