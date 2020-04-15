package cmd

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd *cobra.Command

func init() {
	rootCmd = &cobra.Command{
		Use:           filepath.Base(os.Args[0]),
		Short:         "artifact archival",
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}
	viper.AutomaticEnv()
	rootCmd.AddCommand(archiveToS3Cmd())
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
