// Declare the package,t his should match the folder name
package helpers

import (
	"fmt"
	"os"
)
// Function name has to be capitalized

func Check_key() (vtkey string) {
	vtkey = os.Getenv("vtkey")
	if vtkey == "" {
		fmt.Println("No key, set a VT Key for variable \"vtkey\"")
		os.Exit(1)
	}
	return
}