package main

import (
	"github.com/Wuchieh/go-server/cmd"
	_ "github.com/Wuchieh/go-server/internal/utils/validator/autosetup"
)

func main() {
	cmd.Execute()
}
