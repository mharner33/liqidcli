package cmd

import (
	"fmt"
	"net"

	"github.com/spf13/cobra"
)

var ipAddress string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "liqidcli",
	Short: "A CLI for interacting with Liqid API",
	Long: `A Command Line Interface for calling Liqid API commands.  Typing just liqidcli and the 
	IP address will return a summary of the current system toplogy, similar to using the "list" flag.

	An example of this would be
		liqidcli --ip 10.204.103.30`,
	Version: "1.0",
	//Args: func(rootCmd *cobra.Command, args []string) error{},
	// Uncomment the following line if your bare application
	// has an action associated with it:
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if net.ParseIP(ipAddress) == nil {
			return fmt.Errorf("invalid IP specified: %s", ipAddress)
		}
		return nil
	},
	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Printf("You typed IP: %s", ipAddress)
	// },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	//cobra.OnInitialize(initConfig)
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	//rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.liqidcli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().StringVarP(&ipAddress, "ip", "i", "", "IP address of the Liqid UI")
	rootCmd.MarkFlagRequired("ip")
}

// func initConfig(){

// }
