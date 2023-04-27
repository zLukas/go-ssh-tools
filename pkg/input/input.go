package input

import (
	"flag"
	"strings"
)

type Input struct {
	KeyName          *string
	KeyType          *string
	RemoveKeys		 *bool
	AwsSharedProfile *string
	AwsRegions       []string
	ConfigFile       *string
}

func (i *Input) ParseArgs() {

	i.KeyName = flag.String("n", "my-key", "keyname, default my-key")
	i.RemoveKeys = flag.Bool("c", true, "clean keys before uploading (default true)")
	i.KeyType = flag.String("t", "rsa", "type of encryption, default rsa")
	i.AwsSharedProfile = flag.String("p", "", "(optional)aws shared profile name")
	awsRegions := flag.String("rg", "None", "region to log in, optional if specified in shared profile")
	flag.Parse()
	i.AwsRegions = strings.Split(*awsRegions, ",")
}
