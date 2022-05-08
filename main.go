package main

import (
  "bytes"
  "fmt"
  "log"
  "os/exec"
  "strings"
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
	v := "v1.24.0"
	ami_regions := "us-east-1"
	supportedOS := []string{"amazon-2"}
	
	stderr, stdout, err := Shell("./clusterawsadm ami list --kubernetes-version 1.24.0 --owner-id 570412231501")
	log.Printf("Info: Building AMI for Kubernetes %s.", v)
	
	if err != nil {
		log.Fatalf("ERROR: %v", err)
	}
	if stderr != "" {
		log.Fatalf("STDERR: %v", stderr)
	} else if stdout == "" {
		kubernetes_semver := v
		kubernetes_rpm_version := strings.TrimPrefix(v, "v") + "-0"
		kubernetes_deb_version := strings.TrimPrefix(v, "v") + "-00"
		kubernetes_series := strings.Split(v, ".")[0] + "." + strings.Split(v, ".")[1]

		flags := fmt.Sprintf("-var=ami_regions=%s -var=kubernetes_series=%s -var=kubernetes_semver=%s -var=kubernetes_rpm_version=%s -var=kubernetes_deb_version=%s ", ami_regions, kubernetes_series, kubernetes_semver, kubernetes_rpm_version, kubernetes_deb_version)

		for _, os := range supportedOS {
			switch os {
			case "amazon-2":
				log.Println(fmt.Sprintf("Info: Building AMI for OS %s", os))
				log.Println(fmt.Sprintf("Info: flags:  \"%s\"", flags))
				
				stderr, stdout, err := custom.Shell(fmt.Sprintf("cd image-builder/images/capi && PACKER_FLAGS=\"%s\" make build-ami-%s", flags, os))
				if err != nil {
					log.Fatalf("ERROR: %v", err)
				}
				if stderr != "" {
					log.Fatalf("STDERR: %v", stderr)
				}

				log.Println("--- stdout ---")
				log.Println(stdout)
			}
		}
	} else {
		log.Println("Info: AMI for Kubernetes v1.24.0 already exists. Skipping image building.")
	}
}
