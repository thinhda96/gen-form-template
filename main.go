package main

import (
	"github.com/thinhda96/gen-form-template/cmd"
	"github.com/thinhda96/gen-form-template/utils"
)

func main() {
	hd, err := utils.HomeDir()
	utils.NoError(err)

	cmd.HomeDir = hd
	cmd.Execute()
}
