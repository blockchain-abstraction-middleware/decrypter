package main

import (
	"fmt"

	"github.com/blockchain-abstraction-middleware/encrypt/decrypt"
	"gopkg.in/yaml.v2"
)

// Config is defined to take the decrypted byte array from the decrypter
type Config struct {
	Environment string `yaml:"environment"`
	Name        string `yaml:"name"`
	Port        int    `yaml:"port"`
	Metrics     bool   `ymal:"metrics"`
}

// Sample main to decrypt the review.yml config file into the sample Config struct
func main() {
	var config Config

	fl := "./config/"
	f := "review.yml"

	decryptedFile, err := decrypt.Decrypt(fl, f)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = yaml.Unmarshal(decryptedFile, &config)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Environmnent:", config.Environment)
	fmt.Println("Name:", config.Name)
	fmt.Println("Port:", config.Port)
	fmt.Println("Metrics:", config.Metrics)
}
