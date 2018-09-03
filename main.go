package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func multiPing(sleep *int, ping *string, target string) {

	// TODO: Use map instead of slice for Stdout
	// TODO: Determine why one goroutine exiting fails other goroutines
	// TODO: Ensure utility does not exit when ping fails

	var sumAverage float64
	var sumAverageLoss float64
	var counter int

	for true {
		cmd := exec.Command(*ping, "-c 1", target)
		var out bytes.Buffer
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
		result := out.String()
		resultSplit := strings.Fields(result)
		packetLoss, _ := strconv.ParseFloat(strings.Replace(resultSplit[25], "%", "", -1), 64)
		averagePing, _ := strconv.ParseFloat(strings.Split(resultSplit[33], "/")[2], 64)
		sumAverage += averagePing
		sumAverageLoss += packetLoss
		counter++
		fmt.Printf("Target: %s\t\tCount: %d\tPackets Lost: %.1f%%\t Average: %.3f\t Total Average: %.3f\t\tTotal Loss: %.2f%%\n",
			target, counter, packetLoss, averagePing, sumAverage/float64(counter), sumAverageLoss/float64(counter))
		time.Sleep(time.Second * time.Duration(*sleep))
	}
}

func main() {

	// Validate the ping command is available
	path, err := exec.LookPath("ping")
	if err != nil {
		log.Fatal("Not able to find ping")
	}

	// TODO: Add more flags, such as the -c, -w parameters from PING
	// Define flags
	maxWaitPtr := flag.Int("sleep", 1, "the maximum time in seconds to sleep")
	flag.Parse()

	// Retrieve list of arguments provided at runtime
	targetList := flag.Args()
	if len(targetList) == 0 {
		targetList = []string{"9.9.9.9"}
	}

	for _, t := range targetList {
		go multiPing(maxWaitPtr, &path, t)
	}

	fmt.Println("Hit Enter to end execution")
	defer fmt.Println("test")
	var input string
	fmt.Scanln(&input)
}
