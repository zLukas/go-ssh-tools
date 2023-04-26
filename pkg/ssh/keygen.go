package ssh

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"

	"golang.org/x/crypto/ssh"
)

type KeyPair struct{
	PrivateKey []byte
	PublicKey  []byte
	Name string		
}

func (kp *KeyPair)GenerateKeys() (error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return err
	}

	privateKeyPEM := &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privateKey)}

	pubKey, err := ssh.NewPublicKey(&privateKey.PublicKey)
	if err != nil {
		return err
	}
	kp.PrivateKey = pem.EncodeToMemory(privateKeyPEM)
	kp.PublicKey = ssh.MarshalAuthorizedKey(pubKey)
	return  nil
}

func (kp *KeyPair)SaveKeys() {
	if kp.Name == "" {
		kp.Name = "my-key"
	}
	var err error
	if err = os.WriteFile(kp.Name, kp.PrivateKey, 0600); err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	pubkeyName := kp.Name + ".pub"
	if err = os.WriteFile(pubkeyName, kp.PublicKey, 0644); err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

}