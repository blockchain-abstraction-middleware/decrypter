# gpg-secrets-decryption

## Overview
GPG Secrets Decryption is used for decrypting sops encrypted configs at run time.

Decrypt a config file using private key defined in `$(HOME)/.gnupg`

### Prerequisites

Ensure you have `gpg` installed, this will allow you to generate keys to be used for encryption/decryption
```shell
gpg --version
```

Ensure you have `sops` installed, this will allow you to encrypt the `.yaml` files
1. [Install from source](https://github.com/mozilla/sops)
2. [Download binary](https://github.com/mozilla/sops/releases)
   
```
sops --version
```

Generate a key to use to encrypt files
```shell
gpg --generate-key
```

Get the fingerprint of the key
```shell
gpg --list-keys
```

Export secret key to an application specific location
```shell
gpg --export-secret-keys <fingerprint> > ~/.gnupg/secring.gpg
gpg --export <fingerprint> > ~/.gnupg/pubring.gpg 
```

---

### Sharing keys

To export keys
```shell
gpg --output pubring.gpg --armor --export <fingerprint>
gpg --output secring.gpg --armor --export-secret-key <fingerprint>
```

To import keys
```shell
gpg --import ./pubring.gpg
gpg --allow-secret-key-import --import ./secring.gpg
```

---

### Runtime

#### Encryption

Encrypt a file with sops
```shell
sops -pgp <fingerprint> -e config/review.yml
```

#### To Run

To Run:
```shell
cd cmd/decrypt && go run main.go
```

#### Decryption

Encrypt a file with sops
```shell
sops -d config/review.yml
```
