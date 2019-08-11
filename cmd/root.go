package cmd

import (
	"fmt"
	"os"

	"github.com/seapy/decolink/internal"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Parser ...
var Parser string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "decolink",
	Short: "URL to link with title for markdown(or another)",
	Long: `URL to link with site title. currently only support markdown. For example:

./decolink https://www.github.com
-> 
`,
	Run: func(cmd *cobra.Command, args []string) {
		internal.Init(args)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&Parser, "parser", "p", "goquery", "choose parser")
	viper.BindPFlag("parser", rootCmd.PersistentFlags().Lookup("parser"))
}
