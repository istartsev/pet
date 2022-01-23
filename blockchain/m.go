package main

import (
	"fmt"
	"log"
	"runtime"
	"time"

	"github.com/istartsev/pet/blockchain/block"
)

// from https://habr.com/ru/post/348672/
func getGOMAXPROCS() int {
	return runtime.GOMAXPROCS(0)
}

func main() {
	fmt.Printf("CPUs used: %d\n", getGOMAXPROCS())
	start := time.Now()
	bc := block.NewBlockchain()

	bc.AddBlock("Send 1 BTC to Ivan")
	bc.AddBlock("Send 2 more BTC to Ivan")

	elapsed := time.Since(start)

	bc.Print()
	log.Printf("Took %s", elapsed)
}
