// Copyright © 2017 Chef Software

package main

import (
	"bytes"
	"errors"
	"fmt"
	"text/template"

	"github.com/chef/automate/components/automate-cli/pkg/docs"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	infoCommandTemp = `AUTOMATE DETAILS:{{- "\n"}}
	{{- printf "%50s" "automate_admin_user: "}}{{.Outputs.AutomateAdminUser.Value}}{{- "\n"}}
	{{- printf "%50s" "automate_data_collector_token: "}}{{.Outputs.AutomateDataCollectorToken.Value }}{{- "\n"}}
	{{- printf "%50s" "automate_frontend_urls: "}}{{.Outputs.AutomateURL.Value }}{{- "\n"}}
	
	{{- range $index, $el := .Outputs.AutomatePrivateIps.Value}}{{if eq $index  0}}{{- printf "%50s" "automate_private_ips: "}}{{ $el }}{{"\n"}}{{else}}{{- printf "%50s" ""}}{{ $el }}{{"\n"}}{{end}}{{end}}
	{{- range $index, $el := .Outputs.AutomateSSH.Value}}{{if eq $index  0}}{{- printf "%50s" "automate_ssh: "}}{{ $el }}{{"\n"}}{{else}}{{- printf "%50s" ""}}{{ $el }}{{"\n"}}{{end}}{{end}}

	{{- printf "%50s" "automate_url: "}}{{.Outputs.AutomateURL.Value }}{{- "\n"}}
	
	{{- range $index, $el := .Outputs.ChefServerPrivateIps.Value}}{{if eq $index  0}}{{- printf "%50s" "chef_server_private_ips: "}}{{ $el }}{{"\n"}}{{else}}{{- printf "%50s" ""}}{{ $el }}{{"\n"}}{{end}}{{end}}
	{{- range $index, $el := .Outputs.ChefServerSSH.Value}}{{if eq $index  0}}{{- printf "%50s" "chef_server_ssh: "}}{{ $el }}{{"\n"}}{{else}}{{- printf "%50s" ""}}{{ $el }}{{"\n"}}{{end}}{{end}}

	{{- range $index, $el := .Outputs.OpensearchPrivateIps.Value}}{{if eq $index  0}}{{- printf "%50s" "opensearch_private_ips: "}}{{ $el }}{{"\n"}}{{else}}{{- printf "%50s" ""}}{{ $el }}{{"\n"}}{{end}}{{end}}
	{{- range $index, $el := .Outputs.OpensearchSSH.Value}}{{if eq $index  0}}{{- printf "%50s" "opensearch_ssh: "}}{{ $el }}{{"\n"}}{{else}}{{- printf "%50s" ""}}{{ $el }}{{"\n"}}{{end}}{{end}}
	
	{{- range $index, $el := .Outputs.PostgresqlPrivateIps.Value}}{{if eq $index  0}}{{- printf "%50s" "postgresql_private_ips: "}}{{ $el }}{{"\n"}}{{else}}{{- printf "%50s" ""}}{{ $el }}{{"\n"}}{{end}}{{end}}
	{{- range $index, $el := .Outputs.PostgresqlSSH.Value}}{{if eq $index  0}}{{- printf "%50s" "postgresql_ssh: "}}{{ $el }}{{"\n"}}{{else}}{{- printf "%50s" ""}}{{ $el }}{{"\n"}}{{end}}{{end}}
	
	{{- printf "%50s" "ssh_key_file: "}}{{.Outputs.SSHKeyFile.Value}}{{- "\n"}}
	{{- printf "%50s" "ssh_port: "}}{{.Outputs.SSHPort.Value}}{{- "\n"}}
	{{- printf "%50s" "ssh_user: "}}{{.Outputs.SSHUser.Value}}{{- "\n"}}
`
)

func init() {
	infoCmd.SetUsageTemplate(infoHelpDocs)
	RootCmd.AddCommand(infoCmd)
}

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Info about Automate HA",
	Long:  "Info for Automate HA cluster",
	Annotations: map[string]string{
		NoCheckVersionAnnotation: NoCheckVersionAnnotation,
		docs.Compatibility:       docs.CompatiblewithHA,
	},
	PersistentPreRunE: checkLicenseStatusForExpiry,
	RunE:              runInfoConfigCmd,
}

func runInfoConfigCmd(cmd *cobra.Command, args []string) error {
	if isA2HARBFileExist() {
		return execInfo()
	}
	return errors.New(AUTOMATE_HA_INVALID_BASTION)
}

func execInfo() error {

	automate, err := getAutomateHAInfraDetails()
	if err != nil {
		return err

	}
	var b bytes.Buffer
	err = printInfo(infoCommandTemp, automate, &b)
	if err != nil {
		return err
	}
	fmt.Print(b.String())
	return nil
}

func printInfo(infoCommandTemplate string, automate *AutomateHAInfraDetails, writer *bytes.Buffer) error {
	tmpl, err := template.New("output").Parse(infoCommandTemplate)
	if err != nil {
		logrus.Errorf("Error: %v", err)
		return err
	}
	if err := tmpl.Execute(writer, automate); err != nil {
		logrus.Errorf("Error: %v", err)
		return err
	}
	return nil
}
