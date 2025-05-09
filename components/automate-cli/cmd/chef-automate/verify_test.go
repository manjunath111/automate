package main

import (
	"errors"
	"net/http"
	"strings"
	"testing"

	"github.com/chef/automate/components/automate-cli/pkg/verifyserver/constants"
	"github.com/chef/automate/components/automate-cli/pkg/verifyserver/models"
	"github.com/chef/automate/components/automate-cli/pkg/verifysystemdcreate"
	"github.com/chef/automate/lib/config"
	"github.com/chef/automate/lib/httputils"
	"github.com/chef/automate/lib/majorupgrade_utils"
	"github.com/chef/automate/lib/reporting"
	"github.com/chef/automate/lib/sshutils"
	"github.com/chef/automate/lib/version"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

const (
	CONFIG_FILE                                   = "/config_valid_config_parser.toml"
	STATUS_API_RESPONSE                           = `{"status":"SUCCESS","result":{"status":"OK","services":[],"cli_version":"20230622174936","error":"error getting services from hab svc status"}}`
	BATCH_CHECK_REQUEST                           = `{"status":"SUCCESS","result":{"passed":true,"node_result":[]}}`
	BATCH_CHECK_REQUEST_HAB_ID_WITH_NFS           = `{"status":"SUCCESS","result":{"passed":true,"node_result":[{"node_type":"opensearch","ip":"10.0.164.66","tests":[{"passed":true,"msg":"System User Check","check":"system-user","checks":[{"title":"User creation/validation check","passed":true,"success_msg":"User is created or found successfully","skipped":false}],"skipped":false,"id":{"user_id":"1001","group_id":"1001"}},{"passed": true,"msg": "NFS Backup Config Check","check": "nfs-backup-config"}]},{"node_type":"opensearch","ip":"10.0.133.98","tests":[{"passed":true,"msg":"System User Check","check":"system-user","checks":[{"title":"User creation/validation check","passed":true,"success_msg":"User is created or found successfully","skipped":false}],"skipped":false,"id":{"user_id":"1001","group_id":"1001"}}]},{"node_type":"opensearch","ip":"10.0.146.192","tests":[{"passed":true,"msg":"System User Check","check":"system-user","checks":[{"title":"User creation/validation check","passed":true,"success_msg":"User is created or found successfully","skipped":false}],"skipped":false,"id":{"user_id":"1001","group_id":"1001"}}]},{"node_type":"chef-infra-server","ip":"10.0.128.167","tests":[{"passed":true,"msg":"System User Check","check":"system-user","checks":[{"title":"User creation/validation check","passed":true,"success_msg":"User is created or found successfully","skipped":false}],"skipped":false,"id":{"user_id":"1001","group_id":"1001"}}]},{"node_type":"automate","ip":"10.0.130.38","tests":[{"passed":true,"msg":"System User Check","check":"system-user","checks":[{"title":"User creation/validation check","passed":true,"success_msg":"User is created or found successfully","skipped":false}],"skipped":false,"id":{"user_id":"1001","group_id":"1001"}}]},{"node_type":"postgresql","ip":"10.0.139.234","tests":[{"passed":true,"msg":"System User Check","check":"system-user","checks":[{"title":"User creation/validation check","passed":true,"success_msg":"User is created or found successfully","skipped":false}],"skipped":false,"id":{"user_id":"1001","group_id":"1001"}}]},{"node_type":"postgresql","ip":"10.0.168.59","tests":[{"passed":true,"msg":"System User Check","check":"system-user","checks":[{"title":"User creation/validation check","passed":true,"success_msg":"User is created or found successfully","skipped":false}],"skipped":false,"id":{"user_id":"1001","group_id":"1001"}}]},{"node_type":"postgresql","ip":"10.0.159.200","tests":[{"passed":true,"msg":"System User Check","check":"system-user","checks":[{"title":"User creation/validation check","passed":true,"success_msg":"User is created or found successfully","skipped":false}],"skipped":false,"id":{"user_id":"1001","group_id":"1001"}}]}]}}`
	BATCH_CHECK_REQUEST_HAB_ID_WITH_NFS_FAILURE   = `{"status":"SUCCESS","result":{"passed":true,"node_result":[{"node_type":"opensearch","ip":"10.0.164.66","tests":[{"passed":true,"msg":"System User Check","check":"system-user","checks":[{"title":"User creation/validation check","passed":true,"success_msg":"User is created or found successfully","skipped":false}],"skipped":false,"id":{"user_id":"1001","group_id":"1001"}},{"passed": true,"msg": "NFS Backup Config Check","check": "nfs-backup-config"}]},{"node_type":"opensearch","ip":"10.0.133.98","tests":[{"passed":true,"msg":"System User Check","check":"system-user","checks":[{"title":"User creation/validation check","passed":true,"success_msg":"User is created or found successfully","skipped":false}],"skipped":false,"id":{"user_id":"1001","group_id":"1001"}}]},{"node_type":"opensearch","ip":"10.0.146.192","tests":[{"passed":true,"msg":"System User Check","check":"system-user","checks":[{"title":"User creation/validation check","passed":true,"success_msg":"User is created or found successfully","skipped":false}],"skipped":false,"id":{"user_id":"1002","group_id":"1002"}}]},{"node_type":"chef-infra-server","ip":"10.0.128.167","tests":[{"passed":true,"msg":"System User Check","check":"system-user","checks":[{"title":"User creation/validation check","passed":true,"success_msg":"User is created or found successfully","skipped":false}],"skipped":false,"id":{"user_id":"1001","group_id":"1001"}}]},{"node_type":"automate","ip":"10.0.130.38","tests":[{"passed":true,"msg":"System User Check","check":"system-user","checks":[{"title":"User creation/validation check","passed":true,"success_msg":"User is created or found successfully","skipped":false}],"skipped":false,"id":{"user_id":"1001","group_id":"1001"}}]},{"node_type":"postgresql","ip":"10.0.139.234","tests":[{"passed":true,"msg":"System User Check","check":"system-user","checks":[{"title":"User creation/validation check","passed":true,"success_msg":"User is created or found successfully","skipped":false}],"skipped":false,"id":{"user_id":"1001","group_id":"1001"}}]},{"node_type":"postgresql","ip":"10.0.168.59","tests":[{"passed":true,"msg":"System User Check","check":"system-user","checks":[{"title":"User creation/validation check","passed":true,"success_msg":"User is created or found successfully","skipped":false}],"skipped":false,"id":{"user_id":"1001","group_id":"1001"}}]},{"node_type":"postgresql","ip":"10.0.159.200","tests":[{"passed":true,"msg":"System User Check","check":"system-user","checks":[{"title":"User creation/validation check","passed":true,"success_msg":"User is created or found successfully","skipped":false}],"skipped":false,"id":{"user_id":"1001","group_id":"1001"}}]}]}}`
	BATCH_CHECK_REQUEST_HAB_ID_WITH_NFS_FAILURE_2 = `{"status":"SUCCESS","result":{"passed":true,"node_result":[{"node_type":"opensearch","ip":"10.0.164.66","tests":[{"passed":true,"msg":"System User Check","check":"system-user","checks":[],"skipped":false,"id":{"user_id":"1001","group_id":"1001"}},{"passed": true,"msg": "NFS Backup Config Check","check": "nfs-backup-config"}]},{"node_type":"opensearch","ip":"10.0.133.98","tests":[{"passed":true,"msg":"System User Check","check":"system-user","checks":[],"skipped":false,"id":{"user_id":"1001","group_id":"1001"}}]}]}}`
	AWS_CONFIG_FILE                               = "/valid_config.toml"
	DARWIN                                        = "darwin"
)

var AwsAutoTfvarsJsonStringEmpty = `
{"automate_private_ips":["10.0.130.162", "10.0.153.152"], "chef_server_private_ips":["10.0.136.84", "10.0.149.79"], "postgresql_private_ips":["10.0.135.82", "10.0.159.64", "10.0.161.255"], "opensearch_private_ips":["10.0.133.254", "10.0.144.64", "10.0.163.250"], "automate_lb_fqdn":"A2-df60949d-automate-lb-1317318119.eu-west-2.elb.amazonaws.com", "automate_frontend_url":"https://A2-df60949d-automate-lb-1317318119.eu-west-2.elb.amazonaws.com", "bucket_name":"test", "aws_os_snapshot_role_arn":" ", "os_snapshot_user_access_key_id":" ", "os_snapshot_user_access_key_secret":" "}
`

func TestRunVerifyCmd(t *testing.T) {
	tests := []struct {
		description              string
		mockHttputils            *httputils.MockHTTPClient
		mockCreateSystemdService *verifysystemdcreate.MockCreateSystemdService
		mockSystemdCreateUtils   *verifysystemdcreate.MockSystemdCreateUtils
		mockSSHUtil              *sshutils.MockSSHUtilsImpl
		mockVerifyCmdDeps        *verifyCmdDeps
		configFile               string
		wantErr                  error
		IsAws                    bool
		ConvTfvarToJsonFunc      func(string) string
	}{
		{
			description: "bastion with aws automate-verify - success",
			IsAws:       true,
			mockHttputils: &httputils.MockHTTPClient{
				MakeRequestFunc: func(requestMethod, url string, body interface{}) (*http.Response, []byte, error) {
					if strings.Contains(url, "batch-check") {
						return &http.Response{
							StatusCode: http.StatusOK,
							Body:       nil,
						}, []byte(BATCH_CHECK_REQUEST), nil
					}
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       nil,
					}, []byte(STATUS_API_RESPONSE), nil
				},
			},
			mockCreateSystemdService: &verifysystemdcreate.MockCreateSystemdService{
				CreateFun: func() error {
					return nil
				},
			},
			mockSystemdCreateUtils: &verifysystemdcreate.MockSystemdCreateUtils{
				GetBinaryPathFunc: func() (string, error) {
					return "", nil
				},
			},
			mockSSHUtil: &sshutils.MockSSHUtilsImpl{
				ExecuteConcurrentlyFunc: func(sshConfig sshutils.SSHConfig, cmd string, hostIPs []string) []sshutils.Result {
					return []sshutils.Result{
						{
							HostIP: "",
							Error:  nil,
							Output: "",
						},
					}
				},
				CopyFileToRemoteConcurrentlyFunc: func(sshConfig sshutils.SSHConfig, srcFilePath string, destFileName string, destDir string, removeFile bool, hostIPs []string) []sshutils.Result {
					return []sshutils.Result{
						{
							HostIP: "",
							Error:  nil,
							Output: "",
						},
					}
				},
				CopyFileToRemoteConcurrentlyInHomeDirFunc: func(sshConfig sshutils.SSHConfig, srcFilePath, destFileName, destDir string, removeFile bool, hostIPs []string) []sshutils.Result {
					return []sshutils.Result{
						{
							HostIP: "",
							Error:  nil,
							Output: "",
						},
					}
				},
			},
			configFile: CONFIG_AWS_TOML_PATH + AWS_CONFIG_FILE,
			wantErr:    nil,
			ConvTfvarToJsonFunc: func(string) string {
				return AwsAutoTfvarsJsonStringEmpty
			},
		},
		{
			description: "bastion aws without automate-verify - success",
			IsAws:       true,
			mockHttputils: &httputils.MockHTTPClient{
				MakeRequestFunc: func(requestMethod, url string, body interface{}) (*http.Response, []byte, error) {
					if strings.Contains(url, "batch-check") {
						return &http.Response{
							StatusCode: http.StatusOK,
							Body:       nil,
						}, []byte(BATCH_CHECK_REQUEST), nil
					}
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       nil,
					}, []byte(STATUS_API_RESPONSE), nil
				},
			},
			mockCreateSystemdService: &verifysystemdcreate.MockCreateSystemdService{
				CreateFun: func() error {
					return nil
				},
			},
			mockSystemdCreateUtils: &verifysystemdcreate.MockSystemdCreateUtils{
				GetBinaryPathFunc: func() (string, error) {
					return "", nil
				},
			},
			mockSSHUtil: &sshutils.MockSSHUtilsImpl{
				ExecuteConcurrentlyFunc: func(sshConfig sshutils.SSHConfig, cmd string, hostIPs []string) []sshutils.Result {
					return []sshutils.Result{
						{
							HostIP: "",
							Error:  nil,
							Output: "",
						},
					}
				},
				CopyFileToRemoteConcurrentlyFunc: func(sshConfig sshutils.SSHConfig, srcFilePath string, destFileName string, destDir string, removeFile bool, hostIPs []string) []sshutils.Result {
					return []sshutils.Result{
						{
							HostIP: "",
							Error:  nil,
							Output: "",
						},
					}
				},
				CopyFileToRemoteConcurrentlyInHomeDirFunc: func(sshConfig sshutils.SSHConfig, srcFilePath, destFileName, destDir string, removeFile bool, hostIPs []string) []sshutils.Result {
					return []sshutils.Result{
						{
							HostIP: "",
							Error:  nil,
							Output: "",
						},
					}
				},
			},
			configFile: CONFIG_AWS_TOML_PATH + AWS_CONFIG_FILE,
			wantErr:    nil,
			ConvTfvarToJsonFunc: func(string) string {
				return AwsAutoTfvarsJsonStringEmpty
			},
		},
		{
			description: "bastion with existing automate-verify - success",
			mockHttputils: &httputils.MockHTTPClient{
				MakeRequestFunc: func(requestMethod, url string, body interface{}) (*http.Response, []byte, error) {
					if strings.Contains(url, "batch-check") {
						return &http.Response{
							StatusCode: http.StatusOK,
							Body:       nil,
						}, []byte(BATCH_CHECK_REQUEST), nil
					}
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       nil,
					}, []byte(STATUS_API_RESPONSE), nil
				},
			},
			mockCreateSystemdService: &verifysystemdcreate.MockCreateSystemdService{
				CreateFun: func() error {
					return nil
				},
			},
			mockSystemdCreateUtils: &verifysystemdcreate.MockSystemdCreateUtils{
				GetBinaryPathFunc: func() (string, error) {
					return "", nil
				},
			},
			mockSSHUtil: &sshutils.MockSSHUtilsImpl{
				ExecuteConcurrentlyFunc: func(sshConfig sshutils.SSHConfig, cmd string, hostIPs []string) []sshutils.Result {
					return []sshutils.Result{
						{
							HostIP: "",
							Error:  nil,
							Output: "",
						},
					}
				},
				CopyFileToRemoteConcurrentlyFunc: func(sshConfig sshutils.SSHConfig, srcFilePath string, destFileName string, destDir string, removeFile bool, hostIPs []string) []sshutils.Result {
					return []sshutils.Result{
						{
							HostIP: "",
							Error:  nil,
							Output: "",
						},
					}
				},
				CopyFileToRemoteConcurrentlyInHomeDirFunc: func(sshConfig sshutils.SSHConfig, srcFilePath, destFileName, destDir string, removeFile bool, hostIPs []string) []sshutils.Result {
					return []sshutils.Result{
						{
							HostIP: "",
							Error:  nil,
							Output: "",
						},
					}
				},
			},
			mockVerifyCmdDeps: &verifyCmdDeps{
				getAutomateHAInfraDetails: func() (*AutomateHAInfraDetails, error) {
					return &AutomateHAInfraDetails{}, nil
				},
				PopulateHaCommonConfig: func(configPuller PullConfigs) (haDeployConfig *config.HaDeployConfig, err error) {
					return &config.HaDeployConfig{}, nil
				},
			},
			configFile: CONFIG_TOML_PATH + CONFIG_FILE,
			wantErr:    nil,
		},
		{
			description: "bastion without automate-verify",
			mockHttputils: &httputils.MockHTTPClient{
				MakeRequestFunc: func(requestMethod, url string, body interface{}) (*http.Response, []byte, error) {
					if strings.Contains(url, "batch-check") {
						return &http.Response{
							StatusCode: http.StatusOK,
							Body:       nil,
						}, []byte(BATCH_CHECK_REQUEST), nil
					}
					return nil, nil, errors.New("some error occurred")
				},
			},
			mockCreateSystemdService: &verifysystemdcreate.MockCreateSystemdService{
				CreateFun: func() error {
					return nil
				},
			},
			mockSystemdCreateUtils: &verifysystemdcreate.MockSystemdCreateUtils{
				GetBinaryPathFunc: func() (string, error) {
					return "", nil
				},
			},
			mockSSHUtil: &sshutils.MockSSHUtilsImpl{
				ExecuteConcurrentlyFunc: func(sshConfig sshutils.SSHConfig, cmd string, hostIPs []string) []sshutils.Result {
					return []sshutils.Result{
						{
							HostIP: "",
							Error:  nil,
							Output: "",
						},
					}
				},
				CopyFileToRemoteConcurrentlyFunc: func(sshConfig sshutils.SSHConfig, srcFilePath string, destFileName string, destDir string, removeFile bool, hostIPs []string) []sshutils.Result {
					return []sshutils.Result{
						{
							HostIP: "",
							Error:  nil,
							Output: "",
						},
					}
				},
				CopyFileToRemoteConcurrentlyInHomeDirFunc: func(sshConfig sshutils.SSHConfig, srcFilePath, destFileName, destDir string, removeFile bool, hostIPs []string) []sshutils.Result {
					return []sshutils.Result{
						{
							HostIP: "",
							Error:  nil,
							Output: "",
						},
					}
				},
			},
			mockVerifyCmdDeps: &verifyCmdDeps{
				getAutomateHAInfraDetails: func() (*AutomateHAInfraDetails, error) {
					return &AutomateHAInfraDetails{}, nil
				},
				PopulateHaCommonConfig: func(configPuller PullConfigs) (haDeployConfig *config.HaDeployConfig, err error) {
					return &config.HaDeployConfig{}, nil
				},
			},
			configFile: CONFIG_TOML_PATH + CONFIG_FILE,
			wantErr:    errors.New("This might be due to not enabling the 7799 port"),
		},
		{
			description: "Failed to get automate HA infra details",
			mockHttputils: &httputils.MockHTTPClient{
				MakeRequestFunc: func(requestMethod, url string, body interface{}) (*http.Response, []byte, error) {
					if strings.Contains(url, "batch-check") {
						return &http.Response{
							StatusCode: http.StatusOK,
							Body:       nil,
						}, []byte(BATCH_CHECK_REQUEST), nil
					}
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       nil,
					}, []byte(STATUS_API_RESPONSE), nil
				},
			},
			mockCreateSystemdService: &verifysystemdcreate.MockCreateSystemdService{
				CreateFun: func() error {
					return nil
				},
			},
			mockSystemdCreateUtils: &verifysystemdcreate.MockSystemdCreateUtils{
				GetBinaryPathFunc: func() (string, error) {
					return "", nil
				},
			},
			mockSSHUtil: &sshutils.MockSSHUtilsImpl{
				ExecuteConcurrentlyFunc: func(sshConfig sshutils.SSHConfig, cmd string, hostIPs []string) []sshutils.Result {
					return []sshutils.Result{
						{
							HostIP: "",
							Error:  nil,
							Output: "",
						},
					}
				},
				CopyFileToRemoteConcurrentlyFunc: func(sshConfig sshutils.SSHConfig, srcFilePath string, destFileName string, destDir string, removeFile bool, hostIPs []string) []sshutils.Result {
					return []sshutils.Result{
						{
							HostIP: "",
							Error:  nil,
							Output: "",
						},
					}
				},
				CopyFileToRemoteConcurrentlyInHomeDirFunc: func(sshConfig sshutils.SSHConfig, srcFilePath, destFileName, destDir string, removeFile bool, hostIPs []string) []sshutils.Result {
					return []sshutils.Result{
						{
							HostIP: "",
							Error:  nil,
							Output: "",
						},
					}
				},
			},
			mockVerifyCmdDeps: &verifyCmdDeps{
				getAutomateHAInfraDetails: func() (*AutomateHAInfraDetails, error) {
					return nil, errors.New("Unable to get automate HA infra details")
				},
				PopulateHaCommonConfig: func(configPuller PullConfigs) (haDeployConfig *config.HaDeployConfig, err error) {
					return &config.HaDeployConfig{}, nil
				},
			},
			configFile: "",
			wantErr:    errors.New("Unable to get automate HA infra details"),
		},
		{
			description: "Failed to populate HA common config",
			mockHttputils: &httputils.MockHTTPClient{
				MakeRequestFunc: func(requestMethod, url string, body interface{}) (*http.Response, []byte, error) {
					if strings.Contains(url, "batch-check") {
						return &http.Response{
							StatusCode: http.StatusOK,
							Body:       nil,
						}, []byte(BATCH_CHECK_REQUEST), nil
					}
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       nil,
					}, []byte(STATUS_API_RESPONSE), nil
				},
			},
			mockCreateSystemdService: &verifysystemdcreate.MockCreateSystemdService{
				CreateFun: func() error {
					return nil
				},
			},
			mockSystemdCreateUtils: &verifysystemdcreate.MockSystemdCreateUtils{
				GetBinaryPathFunc: func() (string, error) {
					return "", nil
				},
			},
			mockSSHUtil: &sshutils.MockSSHUtilsImpl{
				ExecuteConcurrentlyFunc: func(sshConfig sshutils.SSHConfig, cmd string, hostIPs []string) []sshutils.Result {
					return []sshutils.Result{
						{
							HostIP: "",
							Error:  nil,
							Output: "",
						},
					}
				},
				CopyFileToRemoteConcurrentlyFunc: func(sshConfig sshutils.SSHConfig, srcFilePath string, destFileName string, destDir string, removeFile bool, hostIPs []string) []sshutils.Result {
					return []sshutils.Result{
						{
							HostIP: "",
							Error:  nil,
							Output: "",
						},
					}
				},
				CopyFileToRemoteConcurrentlyInHomeDirFunc: func(sshConfig sshutils.SSHConfig, srcFilePath, destFileName, destDir string, removeFile bool, hostIPs []string) []sshutils.Result {
					return []sshutils.Result{
						{
							HostIP: "",
							Error:  nil,
							Output: "",
						},
					}
				},
			},
			mockVerifyCmdDeps: &verifyCmdDeps{
				getAutomateHAInfraDetails: func() (*AutomateHAInfraDetails, error) {
					return &AutomateHAInfraDetails{}, nil
				},
				PopulateHaCommonConfig: func(configPuller PullConfigs) (haDeployConfig *config.HaDeployConfig, err error) {
					return nil, errors.New("Failed to populate HA common config")
				},
			},
			configFile: "",
			wantErr:    errors.New("Failed to populate HA common config"),
		},
		{
			description: "bastion with aws automate-verify to check hab if with NFS - failure",
			IsAws:       true,
			mockHttputils: &httputils.MockHTTPClient{
				MakeRequestFunc: func(requestMethod, url string, body interface{}) (*http.Response, []byte, error) {
					if strings.Contains(url, "batch-check") {
						return &http.Response{
							StatusCode: http.StatusOK,
							Body:       nil,
						}, []byte(BATCH_CHECK_REQUEST_HAB_ID_WITH_NFS_FAILURE), nil
					}
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       nil,
					}, []byte(STATUS_API_RESPONSE), nil
				},
			},
			mockCreateSystemdService: &verifysystemdcreate.MockCreateSystemdService{
				CreateFun: func() error {
					return nil
				},
			},
			mockSystemdCreateUtils: &verifysystemdcreate.MockSystemdCreateUtils{
				GetBinaryPathFunc: func() (string, error) {
					return "", nil
				},
			},
			mockSSHUtil: &sshutils.MockSSHUtilsImpl{
				ExecuteConcurrentlyFunc: func(sshConfig sshutils.SSHConfig, cmd string, hostIPs []string) []sshutils.Result {
					return []sshutils.Result{
						{
							HostIP: "",
							Error:  nil,
							Output: "",
						},
					}
				},
				CopyFileToRemoteConcurrentlyFunc: func(sshConfig sshutils.SSHConfig, srcFilePath string, destFileName string, destDir string, removeFile bool, hostIPs []string) []sshutils.Result {
					return []sshutils.Result{
						{
							HostIP: "",
							Error:  nil,
							Output: "",
						},
					}
				},
				CopyFileToRemoteConcurrentlyInHomeDirFunc: func(sshConfig sshutils.SSHConfig, srcFilePath, destFileName, destDir string, removeFile bool, hostIPs []string) []sshutils.Result {
					return []sshutils.Result{
						{
							HostIP: "",
							Error:  nil,
							Output: "",
						},
					}
				},
			},
			configFile: CONFIG_AWS_TOML_PATH + AWS_CONFIG_FILE,
			wantErr:    nil,
			ConvTfvarToJsonFunc: func(string) string {
				return AwsAutoTfvarsJsonStringEmpty
			},
		},
		{
			description: "bastion with aws automate-verify to check hab if with NFS - failure_2",
			IsAws:       true,
			mockHttputils: &httputils.MockHTTPClient{
				MakeRequestFunc: func(requestMethod, url string, body interface{}) (*http.Response, []byte, error) {
					if strings.Contains(url, "batch-check") {
						return &http.Response{
							StatusCode: http.StatusOK,
							Body:       nil,
						}, []byte(BATCH_CHECK_REQUEST_HAB_ID_WITH_NFS_FAILURE_2), nil
					}
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       nil,
					}, []byte(STATUS_API_RESPONSE), nil
				},
			},
			mockCreateSystemdService: &verifysystemdcreate.MockCreateSystemdService{
				CreateFun: func() error {
					return nil
				},
			},
			mockSystemdCreateUtils: &verifysystemdcreate.MockSystemdCreateUtils{
				GetBinaryPathFunc: func() (string, error) {
					return "", nil
				},
			},
			mockSSHUtil: &sshutils.MockSSHUtilsImpl{
				ExecuteConcurrentlyFunc: func(sshConfig sshutils.SSHConfig, cmd string, hostIPs []string) []sshutils.Result {
					return []sshutils.Result{
						{
							HostIP: "",
							Error:  nil,
							Output: "",
						},
					}
				},
				CopyFileToRemoteConcurrentlyFunc: func(sshConfig sshutils.SSHConfig, srcFilePath string, destFileName string, destDir string, removeFile bool, hostIPs []string) []sshutils.Result {
					return []sshutils.Result{
						{
							HostIP: "",
							Error:  nil,
							Output: "",
						},
					}
				},
				CopyFileToRemoteConcurrentlyInHomeDirFunc: func(sshConfig sshutils.SSHConfig, srcFilePath, destFileName, destDir string, removeFile bool, hostIPs []string) []sshutils.Result {
					return []sshutils.Result{
						{
							HostIP: "",
							Error:  nil,
							Output: "",
						},
					}
				},
			},
			configFile: CONFIG_AWS_TOML_PATH + AWS_CONFIG_FILE,
			wantErr:    nil,
			ConvTfvarToJsonFunc: func(string) string {
				return AwsAutoTfvarsJsonStringEmpty
			},
		},
		{
			description: "bastion with aws automate-verify to check hab if with NFS - success",
			IsAws:       true,
			mockHttputils: &httputils.MockHTTPClient{
				MakeRequestFunc: func(requestMethod, url string, body interface{}) (*http.Response, []byte, error) {
					if strings.Contains(url, "batch-check") {
						return &http.Response{
							StatusCode: http.StatusOK,
							Body:       nil,
						}, []byte(BATCH_CHECK_REQUEST_HAB_ID_WITH_NFS), nil
					}
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       nil,
					}, []byte(STATUS_API_RESPONSE), nil
				},
			},
			mockCreateSystemdService: &verifysystemdcreate.MockCreateSystemdService{
				CreateFun: func() error {
					return nil
				},
			},
			mockSystemdCreateUtils: &verifysystemdcreate.MockSystemdCreateUtils{
				GetBinaryPathFunc: func() (string, error) {
					return "", nil
				},
			},
			mockSSHUtil: &sshutils.MockSSHUtilsImpl{
				ExecuteConcurrentlyFunc: func(sshConfig sshutils.SSHConfig, cmd string, hostIPs []string) []sshutils.Result {
					return []sshutils.Result{
						{
							HostIP: "",
							Error:  nil,
							Output: "",
						},
					}
				},
				CopyFileToRemoteConcurrentlyFunc: func(sshConfig sshutils.SSHConfig, srcFilePath string, destFileName string, destDir string, removeFile bool, hostIPs []string) []sshutils.Result {
					return []sshutils.Result{
						{
							HostIP: "",
							Error:  nil,
							Output: "",
						},
					}
				},
				CopyFileToRemoteConcurrentlyInHomeDirFunc: func(sshConfig sshutils.SSHConfig, srcFilePath, destFileName, destDir string, removeFile bool, hostIPs []string) []sshutils.Result {
					return []sshutils.Result{
						{
							HostIP: "",
							Error:  nil,
							Output: "",
						},
					}
				},
			},
			configFile: CONFIG_AWS_TOML_PATH + AWS_CONFIG_FILE,
			wantErr:    nil,
			ConvTfvarToJsonFunc: func(string) string {
				return AwsAutoTfvarsJsonStringEmpty
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			oldBuildTimeVal := version.BuildTime
			version.BuildTime = "20230622174934"
			ConvTfvarToJsonFunc = tt.ConvTfvarToJsonFunc
			cw := majorupgrade_utils.NewCustomWriter()

			vc := NewVerifyCmdFlow(tt.mockHttputils,
				tt.mockCreateSystemdService,
				tt.mockSystemdCreateUtils,
				config.NewHaDeployConfig(),
				tt.mockSSHUtil,
				cw.CliWriter,
				tt.mockVerifyCmdDeps,
			)

			flagsObj := &verifyCmdFlags{
				config: tt.configFile,
			}

			err := vc.runVerifyCmd(nil, nil, flagsObj)
			if tt.wantErr != nil {
				assert.Error(t, err)
				assert.ErrorContains(t, err, tt.wantErr.Error())
			} else {
				assert.NoError(t, err)
			}
			version.BuildTime = oldBuildTimeVal
		})
	}
}

func TestGetHostIPsWithNoLatestCLI(t *testing.T) {
	tests := []struct {
		description           string
		mockHttputils         *httputils.MockHTTPClient
		mockHostIPs           []string
		expectedIpsListLength int
		BuildVersion          string
		wantErr               error
	}{

		{
			description: "remote with existing automate-verify - success",
			mockHttputils: &httputils.MockHTTPClient{
				MakeRequestFunc: func(requestMethod, url string, body interface{}) (*http.Response, []byte, error) {
					return &http.Response{
							StatusCode: http.StatusOK,
							Body:       nil,
						}, []byte(`{
							"status": "SUCCESS",
							"result": {
									"status": "OK",
									"services": [],
									"cli_version": "2023",
									"error": "error getting services from hab svc status"
							}
					}`), nil
				},
			},
			mockHostIPs:           []string{"10.0.0.1", "10.0.0.2", "10.0.0.3", "10.0.0.4", "10.0.0.5"},
			expectedIpsListLength: 0,
			BuildVersion:          "2023",
			wantErr:               nil,
		},
		{
			description: "remote with some unreachable IP",
			mockHttputils: &httputils.MockHTTPClient{
				MakeRequestFunc: func(requestMethod, url string, body interface{}) (*http.Response, []byte, error) {
					if strings.Contains(url, "10.0.0.2") {
						return nil, nil, errors.New("failed to make HTTP request")
					}
					return &http.Response{
							StatusCode: http.StatusOK,
							Body:       nil,
						}, []byte(`{
							"status": "SUCCESS",
							"result": {
									"status": "OK",
									"services": [],
									"cli_version": "2023",
									"error": "error getting services from hab svc status"
							}
					}`), nil
				},
			},
			mockHostIPs:           []string{"10.0.0.1", "10.0.0.2", "10.0.0.3", "10.0.0.4", "10.0.0.5"},
			expectedIpsListLength: 1,
			BuildVersion:          "2023",
			wantErr:               nil,
		},
		{
			description: "remote with some of the machines have old cli version",
			mockHttputils: &httputils.MockHTTPClient{
				MakeRequestFunc: func(requestMethod, url string, body interface{}) (*http.Response, []byte, error) {
					if strings.Contains(url, "10.0.0.2") {
						return &http.Response{
								StatusCode: http.StatusOK,
							}, []byte(`{
							"status": "SUCCESS",
							"result": {
									"status": "OK",
									"services": [],
									"cli_version": "2022",
									"error": "error getting services from hab svc status"
							}
					}`), nil
					}
					return &http.Response{
							StatusCode: http.StatusOK,
						}, []byte(`{
							"status": "SUCCESS",
							"result": {
									"status": "OK",
									"services": [],
									"cli_version": "2023",
									"error": "error getting services from hab svc status"
							}
					}`), nil
				},
			},
			mockHostIPs:           []string{"10.0.0.1", "10.0.0.2", "10.0.0.3", "10.0.0.4", "10.0.0.5"},
			expectedIpsListLength: 1,
			BuildVersion:          "2023",
			wantErr:               nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			version.BuildTime = tt.BuildVersion
			cw := majorupgrade_utils.NewCustomWriter()

			vc := newMockVerifyCmdFlow(tt.mockHttputils, cw)

			hostIpsList, err := vc.getHostIPsWithNoLatestCLI(tt.mockHostIPs)
			if tt.wantErr != nil {
				assert.Error(t, err)
				assert.ErrorContains(t, err, tt.wantErr.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedIpsListLength, len(hostIpsList))
			}
		})
	}
}

func TestVerifyCmdFunc(t *testing.T) {
	tests := []struct {
		test                string
		flagsObj            *verifyCmdFlags
		ConvTfvarToJsonFunc func(string) string
	}{
		{
			test: "Existing Infra",
			flagsObj: &verifyCmdFlags{
				config: CONFIG_TOML_PATH + CONFIG_FILE,
				debug:  true,
			},
			ConvTfvarToJsonFunc: func(string) string {
				return AwsAutoTfvarsJsonStringEmpty
			},
		}, {
			test: "Aws Infra",
			flagsObj: &verifyCmdFlags{
				config: CONFIG_AWS_TOML_PATH + AWS_CONFIG_FILE,
				debug:  true,
			},
			ConvTfvarToJsonFunc: func(string) string {
				return AwsAutoTfvarsJsonStringEmpty
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.test, func(t *testing.T) {
			ConvTfvarToJsonFunc = tt.ConvTfvarToJsonFunc
			vf := verifyCmdFunc(tt.flagsObj)
			assert.NotNil(t, vf, "verifyCmdFunc should not be nil")

			err := vf(nil, nil)
			assert.Error(t, err, "expected an error")
			assert.Contains(t, err.Error(), "Cannot create automate-verify service", "unexpected error message")
		})
	}
}

func newMockVerifyCmdFlow(mockHttputils *httputils.MockHTTPClient, cw *majorupgrade_utils.CustomWriter) *verifyCmdFlow {

	mockCreateSystemdService := &verifysystemdcreate.MockCreateSystemdService{
		CreateFun: func() error {
			return nil
		},
	}
	mockSystemdCreateUtils := &verifysystemdcreate.MockSystemdCreateUtils{
		GetBinaryPathFunc: func() (string, error) {
			return "", nil
		},
	}
	mockSSHUtil := &sshutils.MockSSHUtilsImpl{
		ExecuteConcurrentlyFunc: func(sshConfig sshutils.SSHConfig, cmd string, hostIPs []string) []sshutils.Result {
			return []sshutils.Result{
				{
					HostIP: "",
					Error:  nil,
					Output: "",
				},
			}
		},
		CopyFileToRemoteConcurrentlyFunc: func(sshConfig sshutils.SSHConfig, srcFilePath string, destFileName string, destDir string, removeFile bool, hostIPs []string) []sshutils.Result {
			return []sshutils.Result{
				{
					HostIP: "",
					Error:  nil,
					Output: "",
				},
			}
		},
	}
	mockVerifyCmdDeps := &verifyCmdDeps{
		getAutomateHAInfraDetails: func() (*AutomateHAInfraDetails, error) {
			return &AutomateHAInfraDetails{}, nil
		},
		PopulateHaCommonConfig: func(configPuller PullConfigs) (haDeployConfig *config.HaDeployConfig, err error) {
			return &config.HaDeployConfig{}, nil
		},
	}
	return NewVerifyCmdFlow(mockHttputils,
		mockCreateSystemdService,
		mockSystemdCreateUtils,
		config.NewHaDeployConfig(),
		mockSSHUtil,
		cw.CliWriter,
		mockVerifyCmdDeps,
	)
}

func TestBuildReports(t *testing.T) {
	type args struct {
		batchCheckResults []models.BatchCheckResult
	}
	tests := []struct {
		name string
		args args
		want []reporting.VerificationReport
	}{
		{
			name: "If Report is successfully created",
			args: args{
				batchCheckResults: []models.BatchCheckResult{
					{
						NodeType: "automate",
						Ip:       "1.1.1.1",
						Tests: []models.ApiResult{
							{
								Passed:  true,
								Message: "SSH User Access Check",
								Check:   "ssh-user",
								Checks: []models.Checks{
									{
										Title:         "SSH user accessible",
										Passed:        true,
										SuccessMsg:    "SSH user is accessible for the node: 1.1.1.1",
										ErrorMsg:      "",
										ResolutionMsg: "",
									},
									{
										Title:         "Sudo password valid",
										Passed:        true,
										SuccessMsg:    "SSH user sudo password is valid for the node: 1.1.1.1",
										ErrorMsg:      "",
										ResolutionMsg: "",
									},
								},
								Skipped: false,
							},
						},
					},
				},
			},
			want: []reporting.VerificationReport{
				{
					TableKey: "automate",
					Report: reporting.Info{
						Hostip:    "1.1.1.1",
						Parameter: "ssh-user",
						Status:    "Success",
						StatusMessage: &reporting.StatusMessage{
							MainMessage: "SSH User Access Check - Success",
							SubMessage:  nil,
						},
						SummaryInfo: &reporting.SummaryInfo{
							SuccessfulCount: 2,
							FailedCount:     0,
							ToResolve:       nil,
						},
					},
				},
			},
		},
		{
			name: "Error was there form the handler or Trigger response",
			args: args{
				batchCheckResults: []models.BatchCheckResult{
					{
						NodeType: "automate",
						Ip:       "1.1.1.1",
						Tests: []models.ApiResult{
							{
								Passed:  false,
								Message: "SSH User Access Check",
								Check:   "ssh-user",
								Checks:  nil,
								Error: &fiber.Error{
									Code:    400,
									Message: "Permissions on the ssh key file do not satisfy the requirement",
								},
								Skipped: false,
							},
						},
					},
				},
			},
			want: []reporting.VerificationReport{
				{
					TableKey: "automate",
					Report: reporting.Info{
						Hostip:    "1.1.1.1",
						Parameter: "ssh-user",
						Status:    "Failed",
						StatusMessage: &reporting.StatusMessage{
							MainMessage: "SSH User Access Check - Failed",
							SubMessage:  nil,
						},
						SummaryInfo: &reporting.SummaryInfo{
							SuccessfulCount: 0,
							FailedCount:     1,
							ToResolve:       []string{"Permissions on the ssh key file do not satisfy the requirement"},
						},
					},
				},
			},
		},
		{
			name: "If Report is successfully created with warning message",
			args: args{
				batchCheckResults: []models.BatchCheckResult{
					{
						NodeType: "automate",
						Ip:       "1.1.1.1",
						Tests: []models.ApiResult{
							{
								Passed:  true,
								Message: constants.CERTIFICATE_MSG,
								Check:   constants.CERTIFICATE,
								Checks: []models.Checks{
									{
										Title:         "Certificate check",
										Passed:        true,
										SuccessMsg:    "Certificate will expire in 15 days",
										ErrorMsg:      "Certs are expiring soon",
										ResolutionMsg: "Certificates will expire in 15 days. Please update the certificates!",
									},
									{
										Title:         "Certificate check",
										Passed:        true,
										SuccessMsg:    "Certificate will expire in 15 days",
										ErrorMsg:      "Certs are expiring soon",
										ResolutionMsg: "Certificates will expire in 15 days. Please update the certificates!",
									},
								},
								Skipped: false,
							},
						},
					},
				},
			},
			want: []reporting.VerificationReport{
				{
					TableKey: "automate",
					Report: reporting.Info{
						Hostip:    "1.1.1.1",
						Parameter: constants.CERTIFICATE,
						Status:    "Success",
						StatusMessage: &reporting.StatusMessage{
							MainMessage: "Certificate Check - Success",
							SubMessage:  []string{"Certs are expiring soon", "Certs are expiring soon"},
						},
						SummaryInfo: &reporting.SummaryInfo{
							SuccessfulCount: 2,
							FailedCount:     0,
							ToResolve:       []string{"Certificates will expire in 15 days. Please update the certificates!", "Certificates will expire in 15 days. Please update the certificates!"},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := buildReports(tt.args.batchCheckResults)
			assert.Equal(t, got, tt.want)

		})
	}
}
