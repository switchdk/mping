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

func extractData(resultSplit []string) map[string]string {

	outputData := make(map[string]string)

	for i, rs := range resultSplit {
		switch rs {
		case "loss,":
			outputData["loss"] = resultSplit[i-2]
		case "min/avg/max/mdev":
			outputData["avg"] = resultSplit[i+2]
		}
	}
	return outputData
}

func multiPing(sleep *int, ping *string, target string) {

	var sumAverage float64
	var sumAverageLoss float64
	var counter int
	var out bytes.Buffer

	for true {
		cmd := exec.Command(*ping, "-c 1", target)
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}

		outputData := extractData(strings.Fields(out.String()))
		packetLoss, _ := strconv.ParseFloat(strings.Replace(outputData["loss"], "%", "", -1), 64)
		averagePing, _ := strconv.ParseFloat(strings.Split(outputData["avg"], "/")[2], 64)
		sumAverage += averagePing
		sumAverageLoss += packetLoss
		counter++

		fmt.Printf("Target: %s\t\tCount: %d\tPackets Lost: %.2f%%\t Average: %.3f\t Total Average: %.3f\t\tTotal Loss: %.2f%%\n",
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
