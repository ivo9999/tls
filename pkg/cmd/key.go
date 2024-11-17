package cmd

import (
	"fmt"
	"tls-server/pkg/key"

	"github.com/spf13/cobra"
)

var (
	keyOut    string
	keyLength int
)

func init() {
	createCmd.AddCommand(keyCreateCmd)
	keyCreateCmd.Flags().StringVarP(&keyOut, "key-out", "k", "key.pem", "destination for generated key")
	keyCreateCmd.Flags().IntVarP(&keyLength, "key-length", "l", 4096, "key length in bits")
}

var keyCreateCmd = &cobra.Command{
	Use:   "key",
	Short: "key commands",
	Long:  `commands to create keys`,
	Run: func(cmd *cobra.Command, args []string) {
		err := key.CreateRSAPrivateKeyAndSave(keyOut, 4096)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("key created with length ", keyLength)
	},
}
