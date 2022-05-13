package main

import (
	"flag"
	//"fmt"
	//"io"
	"strings"
	"os"
	"log"
	lab2 "github.com/roman-mazur/architecture-lab-2"
)


var (
	inputExpression = flag.String("e", "", "Expression to compute")
	inputFile = flag.String("f", "", "File with expression to compute")
	outputFile = flag.String("o", "", "File for the result")
)


func main() {
	flag.Parse()
	var handler lab2.ComputeHandler
	if *inputExpression != "" {
		handler = lab2.ComputeHandler{
			Input : strings.NewReader(*inputExpression),
		}
	} else {
		file, err := os.Open(*inputFile)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		handler = lab2.ComputeHandler{
			Input : file,
		}
	}
	
	// TODO: Change this to accept input from the command line arguments as described in the task and
	//       output the results using the ComputeHandler instance.
	//       handler := &lab2.ComputeHandler{
	//           Input: {construct io.Reader according the command line parameters},
	//           Output: {construct io.Writer according the command line parameters},
	//       }
	//       err := handler.Compute()

	//res, _ := lab2.PrefixToPostfix("+ 2 2")
	//fmt.Println(res)
	handler.Compute()
}
