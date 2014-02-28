package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	wpm := flag.Int("w", 250, "wpm(words per minute), must be greater than 0")
	flag.Usage = usage
	flag.Parse()

	if len(flag.Args()) != 1 || *wpm <= 0 {
		flag.Usage()
		os.Exit(0)
	}

	sleepDuration := time.Duration(float32(60) / float32(*wpm) * float32(time.Second))

	wordC := read(flag.Args()[0])

	prevWord := ""

	fmt.Printf("\n")

	for word := range wordC {
		fmt.Printf("\r%s%s\r", strings.Repeat(" ", 10), strings.Repeat(" ", len(prevWord)))
		fmt.Printf("\r%s%s\r", strings.Repeat(" ", 10), word)
		prevWord = word
		time.Sleep(sleepDuration)
	}
}

func read(path string) <-chan string {
	wordC := make(chan string, 256)
	go func() {
		file, err := os.Open(path)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		defer file.Close()
		defer close(wordC)
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			words := strings.Split(scanner.Text(), " ")
			for _, word := range words {
				wordC <- word
			}
		}
		if err := scanner.Err(); err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}()
	return wordC
}

func usage() {
	fmt.Println("Usage: gospritz [OPTIONS] FILEPATH")
	fmt.Println("Options:")
	flag.PrintDefaults()
}
