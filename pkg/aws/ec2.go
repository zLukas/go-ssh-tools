package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

type Client struct {
	cfg         aws.Config
	ProfileName string
	ec2Client   *ec2.Client
	Region      string
}

func (c *Client) Ec2login() error {
	var err error
	if c.ProfileName != "" {
		c.cfg, err = config.LoadDefaultConfig(context.TODO(),
			config.WithRegion(c.Region))
	} else {
		c.cfg, err = config.LoadDefaultConfig(context.TODO(),
			config.WithSharedConfigProfile(c.ProfileName))
	}

	if err == nil {
		c.ec2Client = ec2.NewFromConfig(c.cfg)
		fmt.Printf("succesfully logged to AWS\n")
	} else {
		fmt.Printf("Error while connecting: %v \n", err)
	}

	return err
}

func (c *Client) ImportKeyPairToAws(KeyName string, pubKey []byte, region string) (*string, error) {

	out, err := c.ec2Client.ImportKeyPair(context.TODO(), &ec2.ImportKeyPairInput{
		KeyName: &KeyName, PublicKeyMaterial: pubKey},
		func(opt *ec2.Options) {
			opt.Region = region
		})

	if err != nil {
		return nil, formatAwsError(err)
	}

	return out.KeyPairId, nil
}

func (c * Client) RemoveKeyPairFromAws( keyName string, region string) error {
	_, err := c.ec2Client.DeleteKeyPair(context.TODO(), &ec2.DeleteKeyPairInput{
		KeyName: &keyName},
		func(opt *ec2.Options) {
			opt.Region = region
		})
	if err != nil {
		 return formatAwsError(err)
	}
	return nil
}