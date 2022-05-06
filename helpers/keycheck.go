package helpers

import (
	"fmt"
	"os"
)

func check_key() (vtkey string) {
	vtkey = os.Getenv("vtkey")
	if vtkey == "" {
		fmt.Println("No key, set a VT Key for variable \"vtkey\"")
		os.Exit(1)
	}
	return
}