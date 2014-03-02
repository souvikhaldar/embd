// +build ignore

package main

import (
	"log"
	"time"

	"github.com/kidoman/embd"
	"github.com/kidoman/embd/sensor/lsm303"
)

func main() {
	if err := embd.InitI2C(); err != nil {
		panic(err)
	}
	defer embd.CloseI2C()

	bus := embd.NewI2CBus(1)

	mems := lsm303.New(bus)
	defer mems.Close()

	for {
		heading, err := mems.Heading()
		if err != nil {
			log.Panic(err)
		}
		log.Printf("Heading is %v", heading)

		time.Sleep(500 * time.Millisecond)
	}
}
