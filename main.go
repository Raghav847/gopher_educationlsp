package main

import (
	"bufio"
	"fmt"
	"gopher_educationlsp/rpc"
	"os"
)

func main() {
	fmt.Println("hi")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)
}

func handleMessage(_ any) {

}
