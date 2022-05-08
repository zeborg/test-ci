package main

import (
  "bytes"
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
	stderr, stdout, err := Shell("./clusterawsadm ami list --kubernetes-version 1.24.0 --owner-id 570412231501")
	if err != nil {
		log.Fatalf("ERROR: %v", err)
	}
	if stderr != "" {
		log.Fatalf("STDERR: %v", stderr)
	} else if stdout == "" {
		stderr, stdout, err := Shell("cd image-builder/images/capi && PACKER_FLAGS=\"-var=ami_regions=us-east-1 -var=kubernetes_series=v1.24 -var=kubernetes_semver=v1.24.0 -var=kubernetes_rpm_version=1.24.0-0 -var=kubernetes_deb_version=1.24.0-00 \" make build-ami-amazon-2")
		if err != nil {
			log.Fatalf("ERROR: %v", err)
		}
		if stderr != "" {
			log.Fatalf("STDERR: %v", stderr)
		}

		log.Println("--- stdout ---")
		log.Println(stdout)	
	} else {
		log.Println("Info: AMI for Kubernetes v1.24.0 already exists. Skipping image building.")
	}
}
