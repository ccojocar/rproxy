package cmd

import (
	"os"

	"github.com/ccojocar/rproxy/pkg/config"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile     string
	proxyConfig *config.ProxyCfg
	verbose     bool
)

var rootCmd = &cobra.Command{
	Short: "Simple Reverse Proxy",
}

// Execute main entrypoint
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.rproxy.yaml)")
	rootCmd.PersistentFlags().BoolVar(&verbose, "verbose", true, "enable verbose output logging")
}

// initConfig parse the configuration file and sets up the global configuration
func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		if err != nil {
			log.Fatalf("Reading home directory: %s", err)
		}
		viper.AddConfigPath(home)
		viper.SetConfigName(".rproxy")
	}

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Loading config: %s", err)
	}

	proxyConfig = &config.ProxyCfg{}
	if err := viper.Unmarshal(proxyConfig); err != nil {
		log.Fatalf("Parsing config: %s", err)
	}
	initLogging(verbose)
	log.Infof("Using config: %+v", proxyConfig)
}

// initLogging sets up the logging configuration
func initLogging(verbose bool) {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
	if verbose {
		log.SetLevel(log.InfoLevel)
	}
	log.SetFormatter(&log.JSONFormatter{})
}
