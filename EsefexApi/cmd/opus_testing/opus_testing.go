package main

import (
	"fmt"
	"os"
	"os/exec"
	"sync"
	"time"
)

func main() {
	inputFileName := "testsounds/goofy_ahh.s16le"
	// outputFileName := "testsounds/goofy_ahh.opus"

	fi, err := os.Open(inputFileName)
	fist, err := fi.Stat()
	fib := make([]byte, fist.Size())
	fi.Read(fib)
	fi.Close()

	// fo, err := os.Create(outputFileName)
	// defer fo.Close()
	// fob := make([]byte, fist.Size())

	// Create a pipe to connect stdin and stdout
	cmd := exec.Command("opusenc", "--raw", "--raw-bits=16", "--raw-rate=48000", "--raw-chan=2", "-", "-")
	stdin, err := cmd.StdinPipe()
	stdout, err := cmd.StdoutPipe()

	// Start the opusenc command
	err = cmd.Start()

	// Use a WaitGroup to wait for goroutines to finish
	var wg sync.WaitGroup

	// starting time
	start := time.Now()

	// Goroutine to copy from input file to opusenc's stdin
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Copying data to stdin...")
		stdin.Write(fib)
		if err != nil {
			fmt.Println("Error copying data to stdin:", err)
			return
		}
		stdin.Close() // Close stdin to signal the end of input
	}()

	// relay stdout to this process's stdout
	// wg.Add(1)
	go func() {
		// defer wg.Done()

		buf := make([]byte, 1024)

		for {
			n, err := stdout.Read(buf)
			if err != nil {
				fmt.Println("Error reading stdout:", err)
				return
			}
			if n > 0 {
				os.Stdout.Write(buf[:n])
			}
		}

		// stdin.Close() // Close stdin to signal the end of input
	}()

	// Wait for both goroutines to finish
	wg.Wait()

	// Wait for the opusenc command to finish
	err = cmd.Wait()

	// ending time
	end := time.Now()

	// total time taken
	fmt.Printf("Time taken: %s\n", end.Sub(start))

	fmt.Println("Encoding completed.")
}
