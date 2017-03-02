package main

import (
	"fmt"
	"octlink/mirage/src/utils/uuid"
)

func main() {
	test := uuid.Generate().Simple()
	fmt.Printf("got new uuid %s\n", test)
}
