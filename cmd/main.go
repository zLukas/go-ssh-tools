package main

import (
	"fmt"

	"github.com/zLukas/go-ssh-tools/pkg/aws"
	"github.com/zLukas/go-ssh-tools/pkg/input"
	"github.com/zLukas/go-ssh-tools/pkg/ssh"
)

func main() {
	var args = input.Input{}
	var keys = ssh.KeyPair{}
	var awsClient = aws.Client{}
	args.ParseArgs()

	awsClient.ProfileName = *args.AwsSharedProfile
	awsClient.Ec2login()
	keys.Name = *args.KeyName

	if *args.RemoveKeys {
		fmt.Print("Cleaning old keys... \n")
		for _, reg := range args.AwsRegions {
			err := awsClient.RemoveKeyPairFromAws(keys.Name, reg)
			if err != nil {
				if awsError, ok := err.(aws.AwsError); ok {
					fmt.Printf("Api Error(%s): %s, Operation: %s\n", reg, awsError.ApiError, awsError.Operation)
				}
			} else {
				fmt.Printf("succesfully removed %s key from %s region \n", keys.Name, reg)
			}

		}
	}


	fmt.Printf("generating keypair: %s, %s.pub \n", keys.Name, keys.Name)
	err := keys.GenerateKeys()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	keys.SaveKeys()

	for _, reg := range args.AwsRegions {
		id, err := awsClient.ImportKeyPairToAws(keys.Name, keys.PublicKey, reg)
		if err != nil {
			if awsError, ok := err.(aws.AwsError); ok {
				fmt.Printf("Api Error(%s): %s\n", reg, awsError.ApiError)
			}
		} else {
			fmt.Printf("succesfully imported key with id %s to %s region \n", *id, reg)
		}

	}

}
