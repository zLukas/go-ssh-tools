package input

import "flag"

type Input struct {
	KeyName          *string
	KeyType          *string
	AwsSharedProfile *string
	AwsRegion        *string
}

func (i *Input) ParseArgs() {
	i.KeyName = flag.String("n", "my-key", "keyname, default my-key")
	i.KeyType = flag.String("t", "rsa", "type of encryption, default rsa")
	i.AwsSharedProfile = flag.String("p", "", "(optional)aws shared profile name")
	i.AwsRegion = flag.String("rg", "None", "region to log in, optional if specified in shared profile")
	flag.Parse()
}
