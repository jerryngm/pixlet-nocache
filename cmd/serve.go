package cmd

import (
	"github.com/spf13/cobra"

	"tidbyt.dev/pixlet/server"
)

var (
	host  string
	port  int
	watch bool
	configOutFile string
)

func init() {
	ServeCmd.Flags().StringVarP(&configOutFile,"saveconfig","o","", "Output file for config changes")
	ServeCmd.Flags().StringVarP(&host, "host", "i", "127.0.0.1", "Host interface for serving rendered images")
	ServeCmd.Flags().IntVarP(&port, "port", "p", 8080, "Port for serving rendered images")
	ServeCmd.Flags().BoolVarP(&watch, "watch", "w", true, "Reload scripts on change")
	ServeCmd.Flags().IntVarP(&maxDuration, "max_duration", "d", 15000, "Maximum allowed animation duration (ms)")
}

var ServeCmd = &cobra.Command{
	Use:   "serve [script]",
	Short: "Serve a Pixlet app in a web server",
	Args:  cobra.ExactArgs(1),
	RunE:  serve,
}

func serve(cmd *cobra.Command, args []string) error {
	s, err := server.NewServer(host, port, watch, args[0], maxDuration,configOutFile)
	if err != nil {
		return err
	}
	return s.Run()
}
