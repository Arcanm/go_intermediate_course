// COMMANDS TO ADD MODULES //

// go mod init [NAME]
// go mod init github.com/Arcanm/test

// GO GET [REPOSITORY OF MODULE]
// go get github.com/donvito/hellomod

// INSTALL OTHER VERSION OF THE PACKAGE
// go get github.com/donvito/hellomod/v2

// REMOVE ALL UNUSED DEPENDENCIES
// go mod tidy

// VIEW ALL MODULES DOWNLOADED
// go mod download -json

package main

import (
	"github.com/donvito/hellomod"
	hellomodV2 "github.com/donvito/hellomod/v2"
)

func main() {
	hellomod.SayHello()
	hellomodV2.SayHello("Andres")
}
