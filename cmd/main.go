package main

import (
	"fmt"

	"github.com/zLukas/go-ssh-tools/pkg/aws"
	"github.com/zLukas/go-ssh-tools/pkg/input"
	"github.com/zLukas/go-ssh-tools/pkg/ssh"
)


func main(){
	var args = input.Input{}
	var keys = ssh.KeyPair{}
	var aws = aws.Aws{}
	args.ParseArgs()

	keys.Name = *args.KeyName
	fmt.Printf("generating keypair: %s, %s.pub \n", keys.Name, keys.Name)
	// err := keys.GenerateKeys()
	// if err != nil{
	// 	fmt.Printf("Error: %s\n", err)
	// }
	// keys.SaveKeys()

	aws.Region = *args.AwsRegion
	aws.ProfileName = *args.AwsSharedProfile
	aws.Ec2login()
}