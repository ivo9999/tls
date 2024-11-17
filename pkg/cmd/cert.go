package cmd

import (
	"fmt"
	"os"
	"tls-server/pkg/cert"

	"github.com/spf13/cobra"
)

var (
	certKeyPath string
	certPath    string
	certName    string
)

func init() {
	createCmd.AddCommand(certCreateCmd)
	certCreateCmd.Flags().StringVarP(&certKeyPath, "key-out", "k", "server.key", "destination path for cert key")
	certCreateCmd.Flags().StringVarP(&certPath, "cert-out", "o", "server.crt", "destination path for cert cert")
	certCreateCmd.Flags().StringVarP(&certName, "name", "n", "", "name of the certificate in the config file")
	certCreateCmd.Flags().StringVar(&caKey, "ca-key", "ca.key", "ca key to sign the certificate")
	certCreateCmd.Flags().StringVar(&caCert, "ca-cert", "ca.crt", "ca cert for the certificate")
	certCreateCmd.MarkFlagRequired("ca-key")
	certCreateCmd.MarkFlagRequired("ca-cert")
	certCreateCmd.MarkFlagRequired("name")
}

var certCreateCmd = &cobra.Command{
	Use:   "cert",
	Short: "cert commands",
	Long:  `commands to create the certificates`,
	Run: func(cmd *cobra.Command, args []string) {
		caKeyBytes, err := os.ReadFile(caKey)
		if err != nil {
			fmt.Println("Ca key file not found")
			return
		}
		caCertBytes, err := os.ReadFile(caCert)
		if err != nil {
			fmt.Println("Ca cert file not found")
			return
		}

		certConfig, ok := config.Cert[certName]
		if !ok {
			fmt.Println("Certificate not found in config")
			return
		}

		err = cert.CreateCert(certConfig, caKeyBytes, caCertBytes, certKeyPath, certPath)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Certificate created. Key: %s, Cert: %s \n", certKeyPath, certPath)
	},
}
