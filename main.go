package main

import (
	"fmt"
	"strings"
	"time"
)

// - Insert into the "output" var all Nike + Adidas + Puma product IDs that start with the digit "1"
// - Do not change the the 3 Get functions in any way
// - Utilize go routines (i.e. you will want to be calling the Get functions and inserting into "output" concurrently)

func main() {
	start := time.Now()

	var output []string
	chan1 := make(chan []string)
	chan2 := make(chan []string)
	chan3 := make(chan []string)

	go func() {
		chan1 <- GetNikeProductIDs()
	}()
	go func() {
		chan2 <- GetAdidasProductIDs()
	}()
	go func() {
		chan3 <- GetPumaProductIDs()
	}()
	// recieve data from goroutines
	nikeIDS := <-chan1
	adidasIDS := <-chan2
	pumaIDS := <-chan3
	// filter ids which start with 1
	for _, id := range append(append(nikeIDS, adidasIDS...), pumaIDS...) {
		if strings.HasPrefix(id, "1") {
			output = append(output, id)
		}
	}
	fmt.Println(output)
	fmt.Printf("time taken: %s\n", time.Since(start))
}

func GetNikeProductIDs() (out []string) {
	for i := 0; i < 100; i += 10 {
		out = append(out, fmt.Sprintf("%d_%s", i, "nike"))
	}
	return out
}

func GetAdidasProductIDs() (out []string) {
	for i := 0; i < 100; i += 5 {
		out = append(out, fmt.Sprintf("%d_%s", i, "adidas"))
	}
	return out
}

func GetPumaProductIDs() (out []string) {
	for i := 0; i < 100; i += 2 {
		out = append(out, fmt.Sprintf("%d_%s", i, "puma"))
	}
	return out
}
