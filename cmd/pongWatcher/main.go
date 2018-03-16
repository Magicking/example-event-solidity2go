package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Magicking/example-event-solidity2go/examples"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	flags "github.com/jessevdk/go-flags"
)

var opts struct {
	RawURL          string `env:"RAW_URL" long:"raw-url" required:"true" description:"Websocket raw url"`
	ContractAddress string `env:"CONTRACT_ADDRESS" long:"address" required:"true" description:"Address of Ethereum contract to watch"`
}

func main() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs
		done <- true
	}()
	_, err := flags.Parse(&opts)
	if err != nil {
		log.Fatalf("%v", err)
	}
	c, err := ethclient.Dial(opts.RawURL)
	if err != nil {
		log.Fatalf("Could not initialize client: %v", err)
	}
	addr := common.HexToAddress(opts.ContractAddress)
	ppf, err := examples.NewPingPongFilterer(addr, c)
	if err != nil {
		log.Fatalf("Could not watch for events: %v", err)
	}
	sink := make(chan *examples.PingPongPong)
	sub, err := ppf.WatchPong(&bind.WatchOpts{nil, nil}, sink)
	if err != nil {
		log.Fatalf("Could not register for event: %v", err)
	}
	go func() {
		log.Println(<-sub.Err())
		done <- true
	}()

	log.Printf("Watching events at %v", addr.Hex())

	for e := range sink {
		log.Printf("PingPongPong: %+v", e)
	}

	<-done
	sub.Unsubscribe()
}
