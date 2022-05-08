package main

import (
  "bytes"
  "fmt"
  "log"
  "os/exec"
)

func Shell(command string) (string, string, error) {
  var stdout bytes.Buffer
  var stderr bytes.Buffer
	
  cmd := exec.Command("bash", "-c", command)
  cmd.Stdout = &stdout
  cmd.Stderr = &stderr
  err := cmd.Run()
	
  return stderr.String(), stdout.String(), err
}

func main() {
	stderr, stdout, err := Shell("cd image-builder/images/capi && PACKER_FLAGS=\"-var=ami_regions=us-east-1 -var=kubernetes_series=v1.24 -var=kubernetes_semver=v1.24.0 -var=kubernetes_rpm_version=1.24.0-0 -var=kubernetes_deb_version=1.24.0-00 \" make build-ami-amazon-2")
	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}
	if stderr != "" {
		log.Fatalf("STDERR: %v\n", stderr)
	}

	fmt.Println("--- stdout ---")
	fmt.Println(stdout)
}
