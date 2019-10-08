package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func WaitForServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	fmt.Println(deadline)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil
		}
		log.Printf("Server not responding (%s); retrying...", err)
		time.Sleep(time.Second << uint(tries))
	}
	return fmt.Errorf("Server %s failed to respond after %s", url, timeout)
}

func main() {
	if err := WaitForServer("172.25.153.150:3000"); err != nil {
		fmt.Fprintf(os.Stderr, "Site is down: %v\n", err)
		os.Exit(1)
	}
}
