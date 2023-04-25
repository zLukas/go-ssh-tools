package main

import (
	"fmt"

	"github.com/zLukas/go-ssh-tools/pkg/input"
)


func main(){
	var i = input.Input{}
	i.ParseArgs()
	fmt.Printf("key name:  %s\n", i.KeyName)
	fmt.Printf("key type:  %s\n", i.KeyType)
}