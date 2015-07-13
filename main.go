package main

import (
	"fmt"
	"github.com/docopt/docopt-go"
	"os"
	"strconv"
)

func main() {
	functions := []func(){
		First,
		Second,
		Third,
		Fourth,
		Fifth,
		Sixth,
	}
	usage := `Go euler

Usage:
  go-euler
  go-euler <nr>

Options:
  -h --help     Show this screen`
	arguments, _ := docopt.Parse(usage, nil, true, "0.1", false)

	if arguments["<nr>"] != nil {
		num, err := strconv.Atoi(arguments["<nr>"].(string))
		if err != nil {
			fmt.Println(err)
		} else if num >= len(functions)+1 {
			fmt.Printf("Solution for problem '%v' does not exist\n", num)
			os.Exit(1)
		} else {
			function := functions[num-1]
			fmt.Printf("#### %v ####\n", num)
			function()
		}
	} else {
		for num, function := range functions {
			fmt.Printf("#### %v ####\n", num+1)
			function()
		}
	}

}
