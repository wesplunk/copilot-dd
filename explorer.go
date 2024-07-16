package main

import (
	"fmt"
	"math/rand"
	"time"
)

func seedRandomNumberGenerator() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	seedRandomNumberGenerator()
	fmt.Println("Random number generator seeded!")
}