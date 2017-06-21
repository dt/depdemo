package main

import (
	"fmt"
	"os"

	"github.com/cockroachdb/cockroach/pkg/util/uuid"
)

func main() {
	fmt.Println(uuid.FromString(os.Args[1]))
}
