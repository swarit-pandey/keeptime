package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/swarit-pandey/keeptime/pkg/convert"
)

var convertOpts = &convert.Options{}

// convertCmd represents the convert command
var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert between timezones",
	Long: `Convert lets you convert from a give time zone to another time
zone, and also comes with a formatting option. Currently supported timezones are, IST,
EST, PST, GMT, JST, AEDT, UTC, and CET. 

Examples:
keeptime convert --from=IST --to=UTC --time=now --format=ansic
keeptime convert --from=JST --to=GMT --time=19972607082310
keeptime convert --from=IST --to=EST --time=3d

You can define relative time as well in terms of seconds(s), hours(h), days(d), months(m)
or years(y) in place of time. 
Time layout is YYYYMMDDHHMMSS, and default format is RFC3339
`,
	Run: func(cmd *cobra.Command, args []string) {
		conv := convert.NewConvert(convertOpts)
		err := conv.Start()
		if err != nil {
			fmt.Printf("Error: %v", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(convertCmd)

	convertCmd.Flags().StringVar(&convertOpts.From, "from", "", "Source time zone")
	convertCmd.Flags().StringVar(&convertOpts.To, "to", "", "Target time zone")
	convertCmd.Flags().StringVar(&convertOpts.Time, "time", "", "Time to convert")
	convertCmd.Flags().StringVar(&convertOpts.Format.Format, "format", "", "Output format")
}
