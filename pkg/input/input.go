package input

import "flag"

type Input struct {
	KeyName string
	KeyType string
}

func (i *Input) ParseArgs() {
	i.KeyName = *flag.String("n", "my-key", "keyname, default my-key")
	i.KeyType = *flag.String("t", "rsa", "type of encryption, default rsa")
	flag.Parse()
}
