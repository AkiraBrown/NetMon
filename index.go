package main

import (
	"bufio"
	"fmt"
	"os/exec"
	"strings"
)


func streamNetworkActivity(cmdName string, cmdArgs ...string) {
	cmd := exec.Command(cmdName, cmdArgs...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Printf("Error creating StdoutPipe: %v\n", err)
		return
	}

	if err := cmd.Start(); err != nil {
		fmt.Printf("Error starting netstat command: %v\n", err)
		return
	}

	fmt.Println("Protocol\tLocal Address\t\tForeign Address\t\tState")
	fmt.Println("--------\t-------------\t\t----------------\t\t-----")

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if len(fields) >= 5 && (fields[0] == "tcp" || fields[0] == "udp") {
			protocol := fields[0]
			localAddr := fields[3]
			foreignAddr := fields[4]
			state := ""
			if len(fields) >= 6 {
				state = fields[5]
			}
			fmt.Printf("%-8s\t%-22s\t%-22s\t%s\n", protocol, localAddr, foreignAddr, state)
		}
	}

	if err := cmd.Wait(); err != nil {
		fmt.Printf("Error waiting for netstat command: %v\n", err)
	}
}



func main() {
	streamNetworkActivity("netstat", "-antup")
}
