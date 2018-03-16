# Echo server example

Example for listen to event from outside of Geth

## Run

`docker-compose up --build`

![docker-compose up --build](https://raw.githubusercontent.com/Magicking/example-event-solidity2go/master/docs/run.gif)

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
