package main

import (
	"fmt"
	"os"
)

func main() {
	// if len(os.Args) != 3 {
	// 	fmt.Println("Usage: generic-secrets-fetcher <cloud_provider> <output_file>")
	// 	os.Exit(1)
	// }

	// cloudProvider := os.Args[1]
	// outputFile := os.Args[2]

	var secretValue string
	var err error

	// switch cloudProvider {
	// case "aws":
	// 	secretValue, err = fetchAWSSecret()
	// case "gcp":
	// 	secretValue, err = fetchGCPSecret()
	// case "azure":
	// 	secretValue, err = fetchAzureSecret()
	// default:
	// 	fmt.Println("Invalid cloud provider. Supported values: aws, gcp, azure")
	// 	os.Exit(1)
	// }

	secretValue, err = fetchAzureSecret()

	if err != nil {
		fmt.Println("Error fetching secret:", err)
		os.Exit(1)
	}

	writeOutput(secretValue)

	if err != nil {
		fmt.Println("Error writing secret to file:", err)
		os.Exit(1)
	}

}

func writeOutput(output string) {
	f, err := os.Create("/tmp/secret")
	if err != nil {
		return
	}
	defer f.Close()

	f.WriteString(output)
	fmt.Println("Secret successfully fetched and saved to /tmp/secret")
}
