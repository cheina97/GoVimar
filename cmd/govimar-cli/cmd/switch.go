package cmd

import (
	"fmt"

	"github.com/cheina97/govimar/pkg/config"
	"github.com/cheina97/govimar/pkg/lightswitch"
	"github.com/spf13/cobra"
)

type SwitchAction string

const (
	On  SwitchAction = "on"
	Off SwitchAction = "off"
)

func NewSwitchCmd(s *config.Switch, key string) *cobra.Command {
	ls := lightswitch.NewSwitch(s, key)
	return &cobra.Command{
		Use:       s.Name + " [on|off]",
		Short:     fmt.Sprintf("Control %s switch", s.Name),
		Args:      checkArgsValidity,
		ValidArgs: []string{string(On), string(Off)},
		RunE: func(cmd *cobra.Command, args []string) error {
			var err error
			switch SwitchAction(args[0]) {
			case On:
				_, _, err = ls.On()
			case Off:
				_, _, err = ls.Off()
			}
			if err != nil {
				return err
			}
			return nil
		},
	}
}

func checkArgsValidity(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("accepts 1 arg, received %d", len(args))
	}
	if args[0] != "on" && args[0] != "off" {
		return fmt.Errorf("invalid argument %q, use \"on\" or \"off\"", args[0])
	}
	return nil
}
