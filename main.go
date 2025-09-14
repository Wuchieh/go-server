//go:generate swag fmt
//go:generate swag init -d internal/route -g route.go --requiredByDefault --parseInternal --pd
package main

import (
	"github.com/Wuchieh/go-server/cmd"
	_ "github.com/Wuchieh/go-server/docs"
	_ "github.com/Wuchieh/go-server/internal/utils/validator/autosetup"
)

func main() {
	cmd.Execute()
}
