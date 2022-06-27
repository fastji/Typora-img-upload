/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"github.com/hu-jinwen/Typora-img-upload/cmd"
	_ "github.com/hu-jinwen/Typora-img-upload/cmd/sub"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	cmd.Execute()
}
