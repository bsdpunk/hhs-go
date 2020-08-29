package general

import "os"
import "fmt"

func Rhymes() {
	os.Exit(0)
}

func Clear() {
	fmt.Println("\033[2J")
}
