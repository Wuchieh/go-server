package cmd

import (
	"log"
	"os"

	"github.com/Wuchieh/go-server/internal/bootstrap"
	"github.com/spf13/cobra"
)

var (
	configCmd = &cobra.Command{
		Use:   "config",
		Short: "設定檔管理指令",
	}

	createCmd = &cobra.Command{
		Use:   "create [type]",
		Short: "創建設定檔",
		Long:  "創建設定檔，支援的格式：env, json, yaml, toml",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			configType := args[0]
			if err := bootstrap.CreateConfigFile(configType); err != nil {
				log.Printf("創建設定檔失敗: %v\n", err)
				os.Exit(1)
			}
			log.Printf("設定檔創建成功\n")
		},
	}
)

func init() {
	configCmd.AddCommand(createCmd)
}
