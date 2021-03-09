package internal

import (
	"fmt"
	"os"
	"strconv"
)

//EArgNExist expected argument wasn't passed
type EArgNExist struct{}

//EArgNNat passed arg is not a natural number
type EArgNNat struct{}

func (e EArgNExist) Error() string {
	return "expecting one argument, pid specifically"
}

func (e EArgNNat) Error() string {
	return "pid should be a natural number"
}

var _ error = (*EArgNExist)(nil)
var _ error = (*EArgNNat)(nil)

//Usage print usage and exit
func Usage() {
	fmt.Println("fddf - minimalistic tool for tracking number of file descriptors\n\n" +
		"Usage: fddf [pid]\n" +
		"\t--help   show this message and exit")
}

//ParseArgs process passed arguments
func ParseArgs(args []string) (pid, error) {
	if len(os.Args) != 2 {
		return 0, EArgNExist{}
	}
	if os.Args[1] == "--help" {
		Usage()
		os.Exit(0)
	}

	pidd, err := strconv.Atoi(os.Args[1])
	if err != nil {
		return 0, EArgNNat{}
	} else if pidd < 0 {
		return 0, EArgNNat{}
	}

	return pidd, nil
}
