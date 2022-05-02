package main

import (
  "bytes"
  "fmt"
  "os/exec"
)

func Shell(command string) (error, string, string) {
  var stdout bytes.Buffer
  var stderr bytes.Buffer
	
  cmd := exec.Command("bash", "-c", command)
  cmd.Stdout = &stdout
  cmd.Stderr = &stderr
  err := cmd.Run()
	
  return err, stdout.String(), stderr.String()
}

func main() {
	err, out, errout := Shell("./clusterawsadm ami list --kubernetes-version v1.24.0")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	fmt.Println("--- stdout ---")
	if out != "" {
		fmt.Println(out)
	} else {
		fmt.Println("info: v1.24.0 release not found")
	}
	fmt.Println("--- stderr ---")
	fmt.Println(errout)
}
