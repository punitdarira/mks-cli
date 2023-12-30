package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Deployment struct {
	App     string `yaml:"app"`
	Image   string `yaml:"image"`
	Port    int32  `yaml:"port"`
	Replica int32  `yaml:"replica"`
}

func main() {
	// Read YAML file
	yamlFile, err := os.ReadFile("mks.yaml")
	if err != nil {
		fmt.Printf("Error reading YAML file: %v", err)
		return
	}

	// Unmarshal YAML content into a struct
	var deployment Deployment
	err = yaml.Unmarshal(yamlFile, &deployment)
	if err != nil {
		fmt.Printf("Error unmarshalling YAML: %v", err)
		return
	}

	// Print the values to check if they are read correctly
	fmt.Printf("Image: %s\n", deployment.Image)
	fmt.Printf("Replica: %d\n", deployment.Replica)

	// Now you can proceed to interact with the Kubernetes API
	// ...
	connectCluster(deployment)
}
