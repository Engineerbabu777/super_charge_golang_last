package main

import (
	"fmt"
	"log"
	"time"
)

func main() {

	start := time.Now()
	result, err := TPF()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The response took %v, %+v ", time.Since(start), result)

	
}

func TPF() (string, error) {

	time.Sleep(time.Millisecond * 100)

	return "some response!", nil
}
