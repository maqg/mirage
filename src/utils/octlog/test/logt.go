package main

import (
	"fmt"

	"octlink/mirage/src/utils/octlog"
)

func main() {
	fmt.Printf("running in logger\n")

	octlog.SetLevel(octlog.FatalLevel)

	octlog.Errorf("jjjjjj")
}
