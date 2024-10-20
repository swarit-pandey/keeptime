package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/swarit-pandey/keeptime/pkg/format"
)

var formatOpts = &format.Options{}

// formatCmd represents the format command
var formatCmd = &cobra.Command{
	Use:   "format",
	Short: "Format time",
	Long: `Format helps you format in a specified way, be default the formatting
is unix timestamp. You can give the timezone with flag system, the currently supported
timezones are IST, EST, PST, GMT, JST, AEDT, UTC, and CET

Examples:
keeptime format --system=UTC --time=now --format=ansic
keeptime format --system=IST --time=1m --format=unix


You can define relative time as well in terms of seconds(s), hours(h), days(d), months(m)
or years(y) in place of time. 
`,
	Run: func(cmd *cobra.Command, args []string) {
		formatted := format.NewFormat(formatOpts)
		err := formatted.Start()
		if err != nil {
			fmt.Printf("Error: %v", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(formatCmd)

	formatCmd.Flags().StringVar(&formatOpts.System, "system", "UTC", "Time zone, UTC is default")
	formatCmd.Flags().StringVar(&formatOpts.Time, "time", "", "Time to format")
	formatCmd.Flags().StringVar(&formatOpts.Format, "format", "unix", "Output format, unix is default")
}
