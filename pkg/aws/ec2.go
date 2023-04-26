package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

type Aws struct {
	cfg         aws.Config
	ProfileName string
	Ec2         *ec2.Client
	Region      string
}

func (a *Aws) Ec2login() error {
	var err error
	if a.ProfileName != "" {
		a.cfg, err = config.LoadDefaultConfig(context.TODO(),
			config.WithRegion(a.Region))
	} else {
		a.cfg, err = config.LoadDefaultConfig(context.TODO(),
			config.WithSharedConfigProfile(a.ProfileName))
	}

	if err == nil {
		a.Ec2 = ec2.NewFromConfig(a.cfg)
		fmt.Printf("succesfully logged to AWS\n")
	} else {
		fmt.Printf("Error while connecting: %v", err)
	}

	return err
}

func (a *Aws) ImportKeyPairToAws(KeyName string, pubKey []byte, region string) (*string, error) {

	out, err := a.Ec2.ImportKeyPair(context.TODO(), &ec2.ImportKeyPairInput{
		KeyName: &KeyName, PublicKeyMaterial: pubKey},
		func(opt *ec2.Options) {
			opt.Region = region
		})
	if err != nil {
		fmt.Printf("Error while importing the key: %v", err)
		return nil, err
	}
	return out.KeyPairId, nil
}
