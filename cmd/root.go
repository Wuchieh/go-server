package cmd

import (
	"log"
	"os"
	"path/filepath"

	"github.com/Wuchieh/go-server/internal/bootstrap"
	"github.com/Wuchieh/go-server/internal/flags"
	"github.com/spf13/cobra"
)

var (
	baseName = filepath.Base(os.Args[0])
	rootCmd  = &cobra.Command{
		Use: baseName,
		Run: func(cmd *cobra.Command, args []string) {
			bootstrap.Run()
		},
	}
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&flags.Env, "env", "e", ".env", "env")

	rootCmd.AddCommand(configCmd)
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	os.Exit(0)
}
