package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/keyvault/azsecrets"
)

func fetchAzureSecret() (string, error) {
	// Azure Key Vault configuration
	vaultName := os.Getenv("AZURE_VAULT_NAME")
	secretName := os.Getenv("AZURE_SECRET_NAME")

	if vaultName == "" || secretName == "" {
		return "", fmt.Errorf("AZURE_VAULT_NAME and AZURE_SECRET_NAME must be set")
	}

	vaultURI := fmt.Sprintf("https://%s.vault.azure.net", vaultName)

	// Create a credential using the NewDefaultAzureCredential type.
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}

	// Establish a connection to the Key Vault client
	client, err := azsecrets.NewClient(vaultURI, cred, nil)

	if err != nil {
		log.Fatalf("failed to get the Client: %v", err)
		return "", err
	}

	version := ""
	resp, err := client.GetSecret(context.TODO(), secretName, version, nil)
	if err != nil {
		log.Fatalf("failed to get the secret: %v", err)
		return "", err
	}

	// Print the secret value
	fmt.Printf("The secret value is: %s\n", *resp.Value)

	return *resp.Value, nil

}

// func fetchAWSSecret() {
// 	//os.Setenv("SECRET_NAME", "catsndogs")
// 	//os.Setenv("AWS_REGION", "us-east-1")
// 	SecretName := os.Getenv("SECRET_NAME")
// 	AWSRegion := os.Getenv("AWS_REGION")
// 	sess, err := session.NewSession()
// 	svc := secretsmanager.New(sess, &aws.Config{
// 		Region: aws.String(AWSRegion),
// 	})
// 	input := &secretsmanager.GetSecretValueInput{
// 		SecretId:     aws.String(SecretName),
// 		VersionStage: aws.String("AWSCURRENT"),
// 	}

// 	result, err := svc.GetSecretValue(input)
// 	if err != nil {
// 		if aerr, ok := err.(awserr.Error); ok {
// 			switch aerr.Code() {
// 			case secretsmanager.ErrCodeResourceNotFoundException:
// 				fmt.Println(secretsmanager.ErrCodeResourceNotFoundException, aerr.Error())
// 			case secretsmanager.ErrCodeInvalidParameterException:
// 				fmt.Println(secretsmanager.ErrCodeInvalidParameterException, aerr.Error())
// 			case secretsmanager.ErrCodeInvalidRequestException:
// 				fmt.Println(secretsmanager.ErrCodeInvalidRequestException, aerr.Error())
// 			case secretsmanager.ErrCodeDecryptionFailure:
// 				fmt.Println(secretsmanager.ErrCodeDecryptionFailure, aerr.Error())
// 			case secretsmanager.ErrCodeInternalServiceError:
// 				fmt.Println(secretsmanager.ErrCodeInternalServiceError, aerr.Error())
// 			default:
// 				fmt.Println(aerr.Error())
// 			}
// 		} else {
// 			// Print the error, cast err to awserr.Error to get the Code and
// 			// Message from an error.
// 			fmt.Println(err.Error())
// 		}
// 		return
// 	}
// 	// Decrypts secret using the associated KMS CMK.
// 	// Depending on whether the secret is a string or binary, one of these fields will be populated.
// 	var secretString, decodedBinarySecret string
// 	if result.SecretString != nil {
// 		secretString = *result.SecretString
// 		writeOutput(secretString)
// 	} else {
// 		decodedBinarySecretBytes := make([]byte, base64.StdEncoding.DecodedLen(len(result.SecretBinary)))
// 		len, err := base64.StdEncoding.Decode(decodedBinarySecretBytes, result.SecretBinary)
// 		if err != nil {
// 			fmt.Println("Base64 Decode Error:", err)
// 			return
// 		}
// 		decodedBinarySecret = string(decodedBinarySecretBytes[:len])
// 		writeOutput(decodedBinarySecret)
// 	}

// 	return secretString, nil
// }
