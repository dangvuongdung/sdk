# SDK

SDK is a Golang library for Decrypt and Encrypt a String.

## Installation

Use the package manager to install SDK.

```bash
go get https://github.com/inc-backend/sdk
```

## Usage

```go
package main

import (
	sdk "github.com/inc-backend/sdk/encryption"
	"fmt"
)

func main () {
      privateKey, err := sdk.DecryptToString("private_key_encrypt", "key_name") # returns a private key
      fmt.Println(privateKey)

      encryptPrivateKey, err := sdk.EncryptToString("private_key", "key_name") # returns a encrypt private key
      fmt.Println(encryptPrivateKey)
}
```
