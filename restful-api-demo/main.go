package main

import (
	"fmt"
	"github.com/go-practice/restful-api-demo/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
	}

}
