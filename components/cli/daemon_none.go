// +build !daemon

package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"runtime"
	"strings"
)

func newDaemonCommand() *cobra.Command {
	return &cobra.Command{
		Use:    "daemon",
		Hidden: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runDaemon()
		},
	}
}

func runDaemon() error {
	return fmt.Errorf(
		"`docker daemon` is not supported on %s. Please run `dockerd` directly",
		strings.Title(runtime.GOOS))
}
