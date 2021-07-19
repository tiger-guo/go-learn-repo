package common

import (
	"fmt"
	"os"
)

func CheckErr(err error) {
	fmt.Fprintf(os.Stderr, "%+v", err)
}
