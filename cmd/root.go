package cmd

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var RootCmd = &cobra.Command{
	PersistentPreRunE: initializeConfig,
	Use:               "cronitor-kubernetes",
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	RootCmd.PersistentFlags().String("kubeconfig", "", "path to a kubeconfig to use")
	RootCmd.PersistentFlags().String("apikey", "", "Cronitor.io API key")
	RootCmd.PersistentFlags().String("hostname-override", "", "App hostname to use (mainly for testing)")
	RootCmd.PersistentFlags().String("log-level", "", "Minimum log level to print for the agent (TRACE, DEBUG, INFO, WARN, ERROR)")
	_ = RootCmd.PersistentFlags().MarkHidden("hostname-override")
	RootCmd.PersistentFlags().Bool("dev", false, "Set the CLI to dev mode (for things like logs, etc.)")
	_ = RootCmd.PersistentFlags().MarkHidden("dev")
}

func initializeConfig(cmd *cobra.Command, args []string) error {
	_ = viper.BindEnv("apikey", "CRONITOR_API_KEY")
	_ = viper.BindPFlag("apikey", cmd.Flags().Lookup("apikey"))
	_ = viper.BindPFlag("kubeconfig", cmd.Flags().Lookup("kubeconfig"))
	_ = viper.BindPFlag("hostname-override", cmd.Flags().Lookup("hostname-override"))
	_ = viper.BindPFlag("dev", cmd.Flags().Lookup("dev"))
	_ = viper.BindPFlag("log-level", cmd.Flags().Lookup("log-level"))
	_ = viper.BindEnv("version", "APP_VERSION")

	logLevel := viper.GetString("log-level")
	if logLevel != "" {
		level, err := log.ParseLevel(logLevel)
		if err != nil {
			return err
		}
		log.SetLevel(level)
	}

	return nil
}
