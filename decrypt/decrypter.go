package decrypt

import (
	"fmt"
	"io/ioutil"

	"go.mozilla.org/sops/aes"
	sopsyaml "go.mozilla.org/sops/stores/yaml"
)

// Decrypt decrypts a configuration file and passes back a unencrypted byte Array
func Decrypt(fileLocation string, file string) ([]byte, error) {
	encryptedData, err := ioutil.ReadFile(fileLocation + file)
	if err != nil {
		return nil, fmt.Errorf("Could not read config file: %s", err)
	}

	store := &sopsyaml.Store{}

	tree, err := store.LoadEncryptedFile(encryptedData)
	if err != nil {
		return nil, fmt.Errorf("Could not load encrypted data into store: %s", err)
	}

	key, err := tree.Metadata.GetDataKey()
	if err != nil {
		return nil, fmt.Errorf("Could not load key: %s", err)
	}

	cipher := aes.NewCipher()
	tree.Decrypt(key, cipher)

	decryptedFile, err := store.EmitPlainFile(tree.Branches)
	if err != nil {
		return nil, fmt.Errorf("Could not decrypt value: %s", err)
	}

	return decryptedFile, nil
}
