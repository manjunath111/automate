package main

import (
	"fmt"
	"os"
	"testing"

	dc "github.com/chef/automate/api/config/deployment"
	"github.com/chef/automate/api/config/shared"
	w "github.com/chef/automate/api/config/shared/wrappers"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestGetScriptCommandsForConfigChangedLogging(t *testing.T) {
	type args struct {
		reqConfig   *dc.AutomateConfig
		existConfig *dc.AutomateConfig
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Get Script Commands if both the config is requested for disabling",
			args: args{
				reqConfig: &dc.AutomateConfig{
					Global: &shared.GlobalConfig{
						V1: &shared.V1{
							Log: &shared.Log{
								RedirectSysLog:      w.Bool(false),
								RedirectLogFilePath: w.String("Testing"),
							},
						},
					},
				},
				existConfig: &dc.AutomateConfig{
					Global: &shared.GlobalConfig{
						V1: &shared.V1{
							Log: &shared.Log{
								RedirectSysLog:      w.Bool(true),
								RedirectLogFilePath: w.String("Testing"),
							},
						},
					},
				},
			},
			want: " sudo rm /etc/rsyslog.d/automate.conf; sudo rm /etc/logrotate.d/automate; sudo systemctl restart rsyslog.service;",
		}, {
			name: "Get Script Commands if both the config is requested for logroate only and not file path",
			args: args{
				reqConfig: &dc.AutomateConfig{
					Global: &shared.GlobalConfig{
						V1: &shared.V1{
							Log: &shared.Log{
								RedirectSysLog:       w.Bool(true),
								RedirectLogFilePath:  w.String("Testing"),
								MaxSizeRotateLogs:    w.String("100M"),
								MaxNumberRotatedLogs: w.Int32(10),
								CompressRotatedLogs:  w.Bool(true),
							},
						},
					},
				},
				existConfig: &dc.AutomateConfig{
					Global: &shared.GlobalConfig{
						V1: &shared.V1{
							Log: &shared.Log{
								RedirectSysLog:       w.Bool(true),
								RedirectLogFilePath:  w.String("Testing"),
								MaxSizeRotateLogs:    w.String("10M"),
								MaxNumberRotatedLogs: w.Int32(5),
							},
						},
					},
				},
			},
			want: `sudo sh -c 'echo "Testing/automate.log {
	size 100M
	rotate 10
	missingok
	copytruncate
	compress
	dateext
	dateformat -%Y%m%d%s
}
" > /etc/logrotate.d/automate'`,
		}, {
			name: "Get Script Commands if both the config is requested for file path changed and logrotate values",
			args: args{
				reqConfig: &dc.AutomateConfig{
					Global: &shared.GlobalConfig{
						V1: &shared.V1{
							Log: &shared.Log{
								RedirectSysLog:       w.Bool(true),
								RedirectLogFilePath:  w.String("Testing1"),
								MaxSizeRotateLogs:    w.String("100M"),
								MaxNumberRotatedLogs: w.Int32(10),
								CompressRotatedLogs:  w.Bool(true),
							},
						},
					},
				},
				existConfig: &dc.AutomateConfig{
					Global: &shared.GlobalConfig{
						V1: &shared.V1{
							Log: &shared.Log{
								RedirectSysLog:       w.Bool(true),
								RedirectLogFilePath:  w.String("Testing"),
								MaxSizeRotateLogs:    w.String("10M"),
								MaxNumberRotatedLogs: w.Int32(5),
							},
						},
					},
				},
			},
			want: `sudo sh -c 'echo "\$imjournalRatelimitBurst 20000
\$imjournalRatelimitInterval 600
if \$programname == \"bash\" then Testing1/automate.log
& stop
" > /etc/rsyslog.d/automate.conf'; 
 sudo sh -c 'echo "Testing1/automate.log {
	size 100M
	rotate 10
	missingok
	copytruncate
	compress
	dateext
	dateformat -%Y%m%d%s
}
" > /etc/logrotate.d/automate'; 
 sudo systemctl restart rsyslog.service;`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println(getScriptCommandsForConfigChangedLogging(tt.args.reqConfig, tt.args.existConfig))
			assert.Equalf(t, tt.want, getScriptCommandsForConfigChangedLogging(tt.args.reqConfig, tt.args.existConfig), "getScriptCommandsForConfigChangedLogging(%v, %v)", tt.args.reqConfig, tt.args.existConfig)
		})
	}
}

func TestGetScriptCommandsForLogging(t *testing.T) {
	type args struct {
		reqConfig   *dc.AutomateConfig
		existConfig *dc.AutomateConfig
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Get Script Commands if there is no change in the requested and existed config",
			args: args{
				reqConfig: &dc.AutomateConfig{
					Global: &shared.GlobalConfig{
						V1: &shared.V1{
							Log: &shared.Log{
								RedirectSysLog:      w.Bool(true),
								RedirectLogFilePath: w.String("Testing"),
							},
						},
					},
				},
				existConfig: &dc.AutomateConfig{
					Global: &shared.GlobalConfig{
						V1: &shared.V1{
							Log: &shared.Log{
								RedirectSysLog:      w.Bool(true),
								RedirectLogFilePath: w.String("Testing"),
							},
						},
					},
				},
			},
			want: "",
		},
		{
			name: "Get Script Commands if there is config changed in the requested and existed config with false value",
			args: args{
				reqConfig: &dc.AutomateConfig{
					Global: &shared.GlobalConfig{
						V1: &shared.V1{
							Log: &shared.Log{
								RedirectSysLog:      w.Bool(false),
								RedirectLogFilePath: w.String("Testing"),
							},
						},
					},
				},
				existConfig: &dc.AutomateConfig{
					Global: &shared.GlobalConfig{
						V1: &shared.V1{
							Log: &shared.Log{
								RedirectSysLog:      w.Bool(true),
								RedirectLogFilePath: w.String("Testing"),
							},
						},
					},
				},
			},
			want: ` sudo rm /etc/rsyslog.d/automate.conf; sudo rm /etc/logrotate.d/automate; sudo systemctl restart rsyslog.service;`,
		},
		{
			name: "Get Script Commands if there is config changed in the requested and existed config has more values",
			args: args{
				reqConfig: &dc.AutomateConfig{
					Global: &shared.GlobalConfig{
						V1: &shared.V1{
							Log: &shared.Log{
								RedirectSysLog:      w.Bool(true),
								RedirectLogFilePath: w.String("Testing1"),
							},
						},
					},
				},
				existConfig: &dc.AutomateConfig{
					Global: &shared.GlobalConfig{
						V1: &shared.V1{
							Log: &shared.Log{
								RedirectSysLog:       w.Bool(true),
								RedirectLogFilePath:  w.String("Testing"),
								MaxSizeRotateLogs:    w.String("10k"),
								MaxNumberRotatedLogs: w.Int32(50),
							},
						},
					},
				},
			},
			want: `sudo sh -c 'echo "\$imjournalRatelimitBurst 20000
\$imjournalRatelimitInterval 600
if \$programname == \"bash\" then Testing1/automate.log
& stop
" > /etc/rsyslog.d/automate.conf'; 
 sudo sh -c 'echo "Testing1/automate.log {
	size 10k
	rotate 50
	missingok
	copytruncate
	dateext
	dateformat -%Y%m%d%s
}
" > /etc/logrotate.d/automate'; 
 sudo systemctl restart rsyslog.service;`,
		},
		{
			name: "When there is no exist config and take default values",
			args: args{
				reqConfig: &dc.AutomateConfig{
					Global: &shared.GlobalConfig{
						V1: &shared.V1{
							Log: &shared.Log{
								RedirectSysLog:      w.Bool(true),
								RedirectLogFilePath: w.String("Testing"),
							},
						},
					},
				},
			},
			want: `sudo sh -c 'echo "\$imjournalRatelimitBurst 20000
\$imjournalRatelimitInterval 600
if \$programname == \"bash\" then Testing/automate.log
& stop
" > /etc/rsyslog.d/automate.conf'; 
 sudo sh -c 'echo "Testing/automate.log {
	size 100M
	rotate 10
	missingok
	copytruncate
	dateext
	dateformat -%Y%m%d%s
}
" > /etc/logrotate.d/automate'; 
 sudo systemctl restart rsyslog.service;`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, getScriptCommandsForLogging(tt.args.reqConfig, tt.args.existConfig), "getScriptCommandsForLogging(%v, %v)", tt.args.reqConfig, tt.args.existConfig)
		})
	}
}

func TestLogDecodeFromInput(t *testing.T) {
	t.Run("If there are some parameters", func(t *testing.T) {
		req := `[global.v1.log]
redirect_sys_log = true
redirect_log_file_path = "/var/tmp/"
compress_rotated_logs = true
max_size_rotate_logs = "10M"
max_number_rotated_logs = 5
`

		config, _ := decodeLogConfig(req)

		assert.Equal(t, config.GetGlobal().GetV1().GetLog().GetRedirectLogFilePath().GetValue(), "/var/tmp/")
	})
	t.Run("If there are no parameters", func(t *testing.T) {
		req := ""
		config, _ := decodeLogConfig(req)
		assert.Nil(t, config.GetGlobal().GetV1().GetLog())
	})
	t.Run("If there are some parameters", func(t *testing.T) {
		req := `[global.v1.log]
redirect_sys_log = true
redirect_log_file_path = "/var/tmp/"

`

		config, _ := decodeLogConfig(req)
		fmt.Println(config)

		assert.Equal(t, config.GetGlobal().GetV1().GetLog().GetRedirectLogFilePath().GetValue(), "/var/tmp/")
	})
}

func TestRemoveOrUpdateCentralisedLog(t *testing.T) {
	configDoesnotContainAnythingRelatedToCentralisedLogging := "notCentralisedLoggingconfig.toml"
	CentralisedLoggingEnabledWithPathConfigfileName := "validCentralisedLogging.toml"
	pathNotGivenForCentralisedLogging := "pathNotGivenForCentralisedLogging.toml"
	err := createFileWithContent(configDoesnotContainAnythingRelatedToCentralisedLogging, "[global.v1.log]\nrate_limit_interval = 20\nrate_limit_burst = 20000")
	assert.NoError(t, err)
	err = createFileWithContent(CentralisedLoggingEnabledWithPathConfigfileName, "[global.v1.log]\nredirect_sys_log = true\nredirect_log_file_path = \"test\"")
	assert.NoError(t, err)
	err = createFileWithContent(pathNotGivenForCentralisedLogging, "[global.v1.log]\nredirect_sys_log = true\n")
	assert.NoError(t, err)
	tests := []struct {
		name        string
		sshUtil     SSHUtil
		remoteIp    []string
		remoteType  string
		args        []string
		expectedErr error
	}{
		{
			name:        "Config passed by user doesn't exist",
			sshUtil:     &SSHUtilImpl{},
			remoteIp:    []string{"1.1.1.1"},
			remoteType:  "postgresql",
			args:        []string{"file_doesn't_exist.toml"},
			expectedErr: errors.New(""),
		},
		{
			name: "Config passed by user doesn't contain anything related to Ratelimit",
			sshUtil: &MockSSHUtilsImpl{
				getSSHConfigFunc: func() *SSHConfig {
					return &SSHConfig{}
				},
				connectAndExecuteCommandOnRemoteFunc: func(remoteCommands string, spinner bool) (string, error) {
					return "Executed Successfully", nil
				},
			},
			remoteIp:    []string{"1.1.1.1"},
			remoteType:  "opensearch",
			args:        []string{configDoesnotContainAnythingRelatedToCentralisedLogging},
			expectedErr: nil,
		},
		{
			name: "User wants to enable the rateLimit",
			sshUtil: &MockSSHUtilsImpl{
				getSSHConfigFunc: func() *SSHConfig {
					return &SSHConfig{}
				},
				connectAndExecuteCommandOnRemoteFunc: func(remoteCommands string, spinner bool) (string, error) {
					return "Executed Successfully", nil
				},
			},
			remoteIp:    []string{"1.1.1.1"},
			remoteType:  "postgresql",
			args:        []string{CentralisedLoggingEnabledWithPathConfigfileName},
			expectedErr: nil,
		},
		{
			name: "User wants to enable the rateLimit - But forgot to give path",
			sshUtil: &MockSSHUtilsImpl{
				getSSHConfigFunc: func() *SSHConfig {
					return &SSHConfig{}
				},
				connectAndExecuteCommandOnRemoteFunc: func(remoteCommands string, spinner bool) (string, error) {
					return "Executed Successfully", nil
				},
			},
			remoteIp:    []string{"1.1.1.1"},
			remoteType:  "postgresql",
			args:        []string{pathNotGivenForCentralisedLogging},
			expectedErr: errors.New(""),
		},
		{
			name: "User wants to enable the rateLimit - Bot got some error while enabling it",
			sshUtil: &MockSSHUtilsImpl{
				getSSHConfigFunc: func() *SSHConfig {
					return &SSHConfig{}
				},
				connectAndExecuteCommandOnRemoteFunc: func(remoteCommands string, spinner bool) (string, error) {
					return "", errors.New("Got some error")
				},
			},
			remoteIp:    []string{"1.1.1.1"},
			remoteType:  "postgresql",
			args:        []string{CentralisedLoggingEnabledWithPathConfigfileName},
			expectedErr: errors.New("Got some error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := removeOrUpdateCentralisedLog(tt.args, tt.remoteType, tt.sshUtil, tt.remoteIp)
			if tt.expectedErr != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
	err = os.Remove(CentralisedLoggingEnabledWithPathConfigfileName)
	assert.NoError(t, err)
	err = os.Remove(configDoesnotContainAnythingRelatedToCentralisedLogging)
	assert.NoError(t, err)
	err = os.Remove(pathNotGivenForCentralisedLogging)
	assert.NoError(t, err)
}
