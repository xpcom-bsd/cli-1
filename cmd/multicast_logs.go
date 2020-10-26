package cmd

import (
	"fmt"

	helper "github.com/home-assistant/cli/client"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var multicastLogsCmd = &cobra.Command{
	Use:     "logs",
	Aliases: []string{"log", "lg"},
	Short:   "View the log output of the Home Assistant Multicast server",
	Long: `
Allowing you to look at the log output generated by the Home Assistant Multicast server.
`,
	Example: `
  ha multicast logs
`,
	Run: func(cmd *cobra.Command, args []string) {
		log.WithField("args", args).Debug("multicast logs")

		section := "multicast"
		command := "logs"
		base := viper.GetString("endpoint")

		url, err := helper.URLHelper(base, section, command)
		if err != nil {
			fmt.Printf("Error: %v", err)
			return
		}

		request := helper.GetRequest()
		resp, err := request.SetHeader("Accept", "text/plain").Get(url)

		if err != nil {
			fmt.Println(err)
			ExitWithError = true
		} else {
			fmt.Println(resp.String())
		}
	},
}

func init() {
	multicastCmd.AddCommand(multicastLogsCmd)
}
