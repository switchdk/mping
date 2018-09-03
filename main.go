package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"
)

func multiPing(sleep *int, ping *string, target string) {
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
		packetLoss := resultSplit[25]
		averagePing := strings.Split(resultSplit[33], "/")[2]
		fmt.Println("Target:", target, "Packets Lost:", packetLoss, "Average:", averagePing)
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
	var input string
	fmt.Scanln(&input)
}
