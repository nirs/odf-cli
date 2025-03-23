package dr

import (
	"github.com/ramendr/ramenctl/cmd/commands"
	"github.com/spf13/cobra"
)

// DrCmd is the dr sub command
var DrCmd = commands.RootCmd

func init() {
	// Modify ramenctl RootCmd for odf-cli
	DrCmd.Use = "dr"
	DrCmd.Short = "Troubleshoot ODF DR"
	DrCmd.Annotations = map[string]string{
		cobra.CommandDisplayNameAnnotation: "odf dr",
	}

	// Add a subset of ramenctl commands suitable for "odf dr"
	DrCmd.AddCommand(commands.InitCmd)
	DrCmd.AddCommand(commands.TestCmd)

	/*
		TODO: set build info for ramenctl report.

		build.Version = "v1.2.3"
		build.Commit = "eb92ed81e2715d286bfd8ce173c76d4ecda9e2b4"
	*/
}
