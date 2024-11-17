package cmd

import (
	"fmt"
	"tls-server/pkg/cert"

	"github.com/spf13/cobra"
)

var (
	caKey  string
	caCert string
)

func init() {
	createCmd.AddCommand(caCreateCmd)
	caCreateCmd.Flags().StringVarP(&caKey, "key-out", "k", "ca.key", "destination path for ca key")
	caCreateCmd.Flags().StringVarP(&caCert, "cert-out", "o", "ca.crt", "destination path for ca cert")
}

var caCreateCmd = &cobra.Command{
	Use:   "ca",
	Short: "ca commands",
	Long:  `commands to create the CA`,
	Run: func(cmd *cobra.Command, args []string) {
		err := cert.CreateCACert(config.CACert, caKey, caCert)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("CA created. Key: %s, Cert: %s \n", caKey, caCert)
	},
}
