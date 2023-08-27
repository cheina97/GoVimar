/*
Copyright © 2023 cheina97
*/
package cmd

import (
	"github.com/cheina97/govimar/pkg/config"
	"github.com/spf13/cobra"
)

func NewRootCmd(cfg *config.Config) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "govimar-cli",
		Args:  cobra.NoArgs,
		Short: `govimar-cli is a CLI tool to control a Vimar light switch`,
		Run:   func(cmd *cobra.Command, args []string) {},
	}
	for _, s := range cfg.Switches {
		switchCmd := NewSwitchCmd(&s, cfg.Key)
		rootCmd.AddCommand(switchCmd)
	}
	return rootCmd
}
