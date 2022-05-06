package main

import (
	"fmt"
	"flag"
	// Importing from local folder, we're aliasing our package to H
	H "github.com/remotephone/vt-go/helpers"
	// "os"
	// "strings"
)


func main() {
	// AliasName.FunctionName()
	vtkey := H.Check_key()

	wordPtr := flag.String("hash")
	flag.Parse()

	fmt.Println(H.Check_hash(hash, vtkey))
}