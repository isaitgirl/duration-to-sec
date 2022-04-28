package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {

	precisionPtr := flag.Int("precision", 1, "an int to multiply seconds by")
	flag.Parse()
	precision := *precisionPtr

	reader := bufio.NewReader(os.Stdin)
	input, _, err := reader.ReadLine()
	if err != nil {
		fmt.Printf("Could not read input duration: %s", err)
		os.Exit(1)
	}
	duration := string(input)

	duration_sec, _ := time.ParseDuration(duration)
	duration_sec_final := duration_sec.Seconds() * float64(precision)
	fmt.Println(duration_sec_final)

	os.Exit(0)
}
