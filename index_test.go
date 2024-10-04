package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestSuite(t *testing.T){
	mockData := "tcp 0 0 192.168.1.2:80 192.168.1.1:12345 ESTABLISHED\n" +
		"udp 0 0 192.168.1.2:53 0.0.0.0:0\n"
	cmd := exec.Command("echo", mockData)
	standOut, _ := cmd.StdoutPipe()
	cmd.Start()

	buf := bytes.Buffer{}
	buf.WriteString("Protocol\tLocal Address\t\tForeign Address\t\tState\n")
	buf.WriteString("--------\t-------------\t\t----------------\t\t-----\n")
	

	scanner := bufio.NewScanner(standOut)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)// split the line into fields
		if len(fields) >= 5 && (fields[0] == "tcp" || fields[0] == "udp"){
			protocol := fields[0]
			localAddr := fields[3]
			foreignAddr := fields[4]
			state := ""
			if len(fields) >= 6 {
				state = fields[5]
			}
			buf.WriteString(fmt.Sprintf("%-8s\t%-22s\t%-22s\t%s\n", protocol, localAddr, foreignAddr, state))
		}
		
	}
	cmd.Wait()
	expectedOutput := strings.TrimSpace("Protocol\tLocal Address\t\tForeign Address\t\tState\n" +
		"--------\t-------------\t\t----------------\t\t-----\n" +
		"tcp     \t192.168.1.2:80         \t192.168.1.1:12345     \tESTABLISHED\n" +
		"udp     \t192.168.1.2:53         \t0.0.0.0:0            \n")
	actualOutput := strings.TrimSpace(buf.String())
	if actualOutput != expectedOutput {
		t.Errorf("Expected output:\n%s\nGot:\n%s", expectedOutput, actualOutput)
	}
}

func TestMain(m *testing.M){
	os.Exit(m.Run())
}