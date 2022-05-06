package main

import (
	"fmt"
	H "github.com/remotephone/vt-go/helpers"
	// "os"
	// "strings"
)


func main() {
	vtkey := H.check_key()
	fmt.Println(vtkey)
}