package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"tls-server/pkg/cert"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

type Config struct {
	CACert *cert.CACert          `yaml:"caCert"`
	Cert   map[string]*cert.Cert `yaml:"certs"`
}

var (
	cfgFilePath string
	config      Config
)

var rootCmd = &cobra.Command{
	Use:   "tls",
	Short: "tls is a command line tool for generating self-signed TLS certificates",
	Long:  `tls is a command line tool for generating self-signed TLS certificates, mainly used for generating x509 certificates`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&cfgFilePath, "config", "c", "", "config file (default is tls.yaml)")
}

func initConfig() {
	if cfgFilePath == "" {
		cfgFilePath = "tls.yaml"
	}

	cfgFileBytes, err := ioutil.ReadFile(cfgFilePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = yaml.Unmarshal(cfgFileBytes, &config)
	if err != nil {
		return
	}
}
