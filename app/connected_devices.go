package app

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/grandcat/zeroconf"
)

func ConnectedDevices() {
	service := "_airplay._tcp"
	domain := "local"

	await := 3

	resolver, err := zeroconf.NewResolver(nil)

	if err != nil {
		log.Fatalln("Error while trying to initialize resolver: ", err.Error())
	}

	entries := make(chan *zeroconf.ServiceEntry)

	go func(results <-chan *zeroconf.ServiceEntry) {
		for entry := range results {
			for _, value := range entry.Text {
				fmt.Println(value)
			}
		}
	}(entries)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(await))

	defer cancel()

	err = resolver.Browse(ctx, service, domain, entries)

	if err != nil {
		log.Fatalln("Browsing failed: ", err.Error())
	}

	<-ctx.Done()
}
