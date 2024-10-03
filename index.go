package main

import (
	"fmt"
	"os/exec"
	"time"
)

func monitorNetStat(interval time.Duration){
	for {
		cmd := exec.Command("netstat", "-an")
		output,err := cmd.Output()
		if err != nil {
			fmt.Printf("Error executing netstat: %v\n", err)
			return
		}
		fmt.Println(string(output))
		time.Sleep(interval)
	}
}

func main(){
	interval :=5 * time.Second
	monitorNetStat(interval)
}