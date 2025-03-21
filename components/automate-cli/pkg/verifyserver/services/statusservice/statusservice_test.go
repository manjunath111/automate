package statusservice_test

import (
	"errors"
	"testing"

	"github.com/chef/automate/components/automate-cli/pkg/verifyserver/constants"
	"github.com/chef/automate/components/automate-cli/pkg/verifyserver/models"
	"github.com/chef/automate/components/automate-cli/pkg/verifyserver/services/statusservice"
	"github.com/chef/automate/lib/logger"
	"github.com/stretchr/testify/assert"
)

const (
	automateStatusOutputOnA2 = `Status from deployment with channel [current] and type [local]

	Service Name            Process State  Health Check  Uptime (s) PID 
	deployment-service      running        ok            1287       4951
	backup-gateway          running        ok            1181       5984
	automate-pg-gateway     running        ok            1180       6060
	automate-es-gateway     running        ok            1181       6019
	automate-ui             running        ok            1180       6079
	pg-sidecar-service      running        ok            1177       6339
	cereal-service          running        ok            1174       6550
	event-service           running        ok            1180       6102
	authz-service           running        ok            1174       6677
	es-sidecar-service      running        ok            1179       6171
	event-feed-service      running        ok            1177       6274
	automate-dex            running        ok            1173       6807
	teams-service           running        ok            1176       6362
	session-service         running        ok            1165       7503
	authn-service           running        ok            1176       6429
	secrets-service         running        ok            1175       6463
	applications-service    running        ok            1174       6573
	notifications-service   running        ok            1175       6520
	nodemanager-service     running        ok            1173       6765
	compliance-service      running        ok            1174       6715
	license-control-service running        ok            1173       6869
	local-user-service      running        ok            1173       6836
	config-mgmt-service     running        ok            1172       6938
	ingest-service          running        ok            1170       7025
	infra-proxy-service     running        ok            1171       7003
	data-feed-service       running        ok            1169       7115
	event-gateway           running        ok            1170       7082
	report-manager-service  running        ok            1169       7151
	user-settings-service   running        ok            1169       7171
	automate-gateway        running        ok            1168       7243
	automate-cs-bookshelf   running        ok            1167       7295
	automate-cs-oc-bifrost  running        ok            1166       7322
	automate-cs-oc-erchef   running        ok            1164       7619
	automate-cs-nginx       running        ok            1163       7671
	automate-load-balancer  running        ok            1164       7640`

	automateStatusOutputOnA2Unhealthy = `Status from deployment with channel [current] and type [local]

	Service Name            Process State  Health Check  Uptime (s) PID  
	deployment-service      running        ok            34         12910
	backup-gateway          running        ok            29         13359
	automate-pg-gateway     running        ok            30         13328
	automate-es-gateway     running        CRITICAL      30         13297
	automate-ui             running        ok            30         13316
	pg-sidecar-service      running        ok            29         13397
	cereal-service          running        ok            29         13425
	event-service           running        ok            29         13454
	authz-service           running        ok            28         13525
	es-sidecar-service      running        ok            28         13502
	event-feed-service      running        ok            27         13608
	automate-dex            running        CRITICAL      27         13583
	teams-service           running        ok            26         13659
	session-service         running        CRITICAL      26         13702
	authn-service           running        ok            25         13763
	secrets-service         running        ok            25         13801
	applications-service    running        CRITICAL      24         13849
	notifications-service   running        ok            24         13873
	nodemanager-service     running        ok            23         13985
	compliance-service      running        CRITICAL      23         14015
	license-control-service running        ok            22         14086
	local-user-service      running        CRITICAL      21         14135
	config-mgmt-service     running        unknown       21         14154
	ingest-service          running        unknown       20         14212
	infra-proxy-service     running        ok            20         14250
	data-feed-service       running        ok            19         14295
	event-gateway           running        ok            19         14325
	report-manager-service  running        ok            18         14382
	user-settings-service   running        ok            16         14471
	automate-gateway        running        unknown       17         14449
	automate-cs-bookshelf   running        ok            16         14529
	automate-cs-oc-bifrost  running        CRITICAL      15         14565
	automate-cs-oc-erchef   running        CRITICAL      14         14606
	automate-cs-nginx       running        CRITICAL      14         14646
	automate-load-balancer  running        ok            13         14662
	
	UnhealthyStatusError: System status is unhealthy: One or more services are unhealthy`

	habSvcStatusWithLicenseOutputOnA2 = `+---------------------------------------------+

	Chef License Acceptance


Before you can continue, 1 product license must be accepted.
View the license at https://www.chef.io/end-user-license-agreement

License that needs accepting:

* Habitat

Do you accept the 1 product license? [yes/No/quit] yes

Accepting 1 product license...
✓  1 product license accepted.

+---------------------------------------------+

package                                            type        desired  state  elapsed (s)  pid   group
chef/backup-gateway/0.1.0/20230223070223           standalone  up       up     1217         5984  backup-gateway.default
chef/report-manager-service/0.1.0/20230309152830   standalone  up       up     1205         7151  report-manager-service.default
chef/user-settings-service/0.1.0/20230417072436    standalone  up       up     1205         7171  user-settings-service.default
chef/compliance-service/1.11.1/20230417072436      standalone  up       up     1210         6715  compliance-service.default
chef/license-control-service/1.0.0/20230223070129  standalone  up       up     1209         6869  license-control-service.default
chef/automate-ui/2.0.0/20230426102739              standalone  up       up     1216         6079  automate-ui.default
chef/config-mgmt-service/0.1.0/20230224142110      standalone  up       up     1208         6938  config-mgmt-service.default
chef/event-gateway/0.1.0/20230223070110            standalone  up       up     1206         7082  event-gateway.default
chef/teams-service/0.1.0/20230223070128            standalone  up       up     1212         6362  teams-service.default
chef/deployment-service/0.1.0/20230502070345       standalone  up       up     1323         4951  deployment-service.default
chef/event-service/0.1.0/20230309152624            standalone  up       up     1216         6102  event-service.default
chef/applications-service/1.0.0/20230223070130     standalone  up       up     1210         6573  applications-service.default
chef/automate-cs-bookshelf/15.4.0/20230410161619   standalone  up       up     1203         7295  automate-cs-bookshelf.default
chef/event-feed-service/1.0.0/20230223070128       standalone  up       up     1213         6274  event-feed-service.default
chef/automate-dex/0.1.0/20230223070225             standalone  up       up     1209         6807  automate-dex.default
chef/session-service/0.1.0/20230223070128          standalone  up       up     1201         7503  session-service.default
chef/automate-es-gateway/0.1.0/20230223070033      standalone  up       up     1217         6019  automate-es-gateway.default
chef/authz-service/0.1.0/20230223070223            standalone  up       up     1210         6677  authz-service.default
chef/local-user-service/0.1.0/20230223065533       standalone  up       up     1209         6836  local-user-service.default
chef/pg-sidecar-service/0.0.1/20230223070131       standalone  up       up     1213         6339  pg-sidecar-service.default
chef/notifications-service/1.0.0/20230323141551    standalone  up       up     1211         6520  notifications-service.default
chef/data-feed-service/1.0.0/20230309152831        standalone  up       up     1205         7115  data-feed-service.default
chef/infra-proxy-service/0.1.0/20230427125632      standalone  up       up     1207         7003  infra-proxy-service.default
chef/automate-pg-gateway/0.0.1/20230130151627      standalone  up       up     1216         6060  automate-pg-gateway.default
chef/automate-load-balancer/0.1.0/20230427090837   standalone  up       up     1200         7640  automate-load-balancer.default
chef/authn-service/0.1.0/20230223070225            standalone  up       up     1212         6429  authn-service.default
chef/nodemanager-service/1.0.0/20230417072436      standalone  up       up     1209         6765  nodemanager-service.default
chef/ingest-service/0.1.0/20230309152831           standalone  up       up     1206         7025  ingest-service.default
chef/automate-cs-oc-bifrost/15.4.0/20230223070128  standalone  up       up     1202         7322  automate-cs-oc-bifrost.default
chef/es-sidecar-service/1.0.0/20230130152441       standalone  up       up     1215         6171  es-sidecar-service.default
chef/automate-gateway/0.1.0/20230417072436         standalone  up       up     1204         7243  automate-gateway.default
chef/automate-cs-nginx/15.4.0/20230223065651       standalone  up       up     1199         7671  automate-cs-nginx.default
chef/cereal-service/0.1.0/20230223070128           standalone  up       up     1210         6550  cereal-service.default
chef/secrets-service/1.0.0/20230223070129          standalone  up       up     1211         6463  secrets-service.default
chef/automate-cs-oc-erchef/15.4.0/20230410161619   standalone  up       up     1200         7619  automate-cs-oc-erchef.default`

	automateStatusOutputOnCS = `Status from deployment with channel [current] and type [local]

	Service Name            Process State  Health Check  Uptime (s) PID 
	deployment-service      running        ok            954        7631
	backup-gateway          running        ok            898        8171
	automate-pg-gateway     running        warn          898        8239
	automate-es-gateway     running        ok            897        8292
	pg-sidecar-service      running        critical      895        8547
	es-sidecar-service      running        ok            897        8311
	license-control-service running        ok            893        8578
	automate-cs-bookshelf   running        ok            891        8633
	automate-cs-oc-bifrost  running        ok            892        8605
	automate-cs-oc-erchef   running        unknown       891        8678
	automate-cs-nginx       running        ok            895        8502
	automate-load-balancer  running        ok            862        9278
	`

	habSvcStatusOutputOnCS = `package                                            type        desired  state  elapsed (s)  pid   group
	chef/deployment-service/0.1.0/20230502070345       standalone  up       up     974          7631  deployment-service.default
	chef/license-control-service/1.0.0/20230223070129  standalone  up       up     913          8578  license-control-service.default
	chef/automate-load-balancer/0.1.0/20230427090837   standalone  up       up     882          9278  automate-load-balancer.default
	chef/backup-gateway/0.1.0/20230223070223           standalone  up       up     918          8171  backup-gateway.default
	chef/pg-sidecar-service/0.0.1/20230223070131       standalone  up       up     915          8547  pg-sidecar-service.default
	chef/automate-cs-oc-bifrost/15.4.0/20230223070128  standalone  up       up     912          8605  automate-cs-oc-bifrost.default
	chef/automate-cs-bookshelf/15.4.0/20230410161619   standalone  up       up     911          8633  automate-cs-bookshelf.default
	chef/automate-cs-nginx/15.4.0/20230223065651       standalone  up       up     915          8502  automate-cs-nginx.default
	chef/automate-es-gateway/0.1.0/20230223070033      standalone  up       up     917          8292  automate-es-gateway.default
	chef/es-sidecar-service/1.0.0/20230130152441       standalone  up       up     917          8311  es-sidecar-service.default
	chef/automate-cs-oc-erchef/15.4.0/20230410161619   standalone  up       up     911          8678  automate-cs-oc-erchef.default
	chef/automate-pg-gateway/0.0.1/20230130151627      standalone  up       up     918          8239  automate-pg-gateway.default
	`

	habSvcStatusOutputOnPG = `package                                            type        desired  state  elapsed (s)  pid   group
	chef/automate-ha-pgleaderchk/0.1.0/20230130152444  standalone  up       up     1507         6325  automate-ha-pgleaderchk.default
	chef/automate-ha-haproxy/2.2.14/20230130151541     standalone  up       up     1507         6343  automate-ha-haproxy.default
	chef/automate-ha-postgresql/13.5.0/20230130151541  standalone  up       up     1501         6462  automate-ha-postgresql.default`

	habSvcStatusOutputOnOS = `
	package                                               type        desired  state  elapsed (s)  pid   group
	chef/automate-ha-opensearch/1.3.7/20230223065900      standalone  up       up     1538         5653  automate-ha-opensearch.default
	chef/automate-ha-elasticsidecar/0.1.0/20230223070538  standalone  up       down     1537         5739  automate-ha-elasticsidecar.default
	`
)

func TestStatusService(t *testing.T) {
	log, err := logger.NewLogger("text", "debug")
	assert.NoError(t, err)

	type testCaseInfo struct {
		testCaseDescription string
		input               func(cmd string) ([]byte, error)
		expectedOutput      *[]models.ServiceDetails
		isError             bool
		errorMsg            string
	}

	testCases := []testCaseInfo{
		{
			testCaseDescription: "FE Node",
			input: func(cmd string) ([]byte, error) {
				if cmd == constants.AUTOMATESTATUSCMD {
					return []byte(automateStatusOutputOnCS), nil
				}
				return []byte(habSvcStatusOutputOnCS), nil
			},
			expectedOutput: &[]models.ServiceDetails{{ServiceName: "deployment-service", Status: "OK", Version: "chef/deployment-service/0.1.0/20230502070345"}, {ServiceName: "license-control-service", Status: "OK", Version: "chef/license-control-service/1.0.0/20230223070129"}, {ServiceName: "automate-load-balancer", Status: "OK", Version: "chef/automate-load-balancer/0.1.0/20230427090837"}, {ServiceName: "backup-gateway", Status: "OK", Version: "chef/backup-gateway/0.1.0/20230223070223"}, {ServiceName: "pg-sidecar-service", Status: "CRITICAL", Version: "chef/pg-sidecar-service/0.0.1/20230223070131"}, {ServiceName: "automate-cs-oc-bifrost", Status: "OK", Version: "chef/automate-cs-oc-bifrost/15.4.0/20230223070128"}, {ServiceName: "automate-cs-bookshelf", Status: "OK", Version: "chef/automate-cs-bookshelf/15.4.0/20230410161619"}, {ServiceName: "automate-cs-nginx", Status: "OK", Version: "chef/automate-cs-nginx/15.4.0/20230223065651"}, {ServiceName: "automate-es-gateway", Status: "OK", Version: "chef/automate-es-gateway/0.1.0/20230223070033"}, {ServiceName: "es-sidecar-service", Status: "OK", Version: "chef/es-sidecar-service/1.0.0/20230130152441"}, {ServiceName: "automate-cs-oc-erchef", Status: "UNKNOWN", Version: "chef/automate-cs-oc-erchef/15.4.0/20230410161619"}, {ServiceName: "automate-pg-gateway", Status: "WARN", Version: "chef/automate-pg-gateway/0.0.1/20230130151627"}},
			isError:        false,
			errorMsg:       "",
		},
		{
			testCaseDescription: "BE Node",
			input: func(cmd string) ([]byte, error) {
				if cmd == constants.AUTOMATESTATUSCMD {
					return []byte(constants.AUTOMATESTATUSONBEERROR), errors.New("exit status 90")
				}
				return []byte(habSvcStatusOutputOnOS), nil
			},
			expectedOutput: &[]models.ServiceDetails{{ServiceName: "automate-ha-opensearch", Status: "OK", Version: "chef/automate-ha-opensearch/1.3.7/20230223065900"}, {ServiceName: "automate-ha-elasticsidecar", Status: "CRITICAL", Version: "chef/automate-ha-elasticsidecar/0.1.0/20230223070538"}},
			isError:        false,
			errorMsg:       "",
		},
		{
			testCaseDescription: "Hab Error",
			input: func(cmd string) ([]byte, error) {
				if cmd == constants.AUTOMATESTATUSCMD {
					return []byte(automateStatusOutputOnCS), nil
				}
				return []byte("Failed to execute hab command"), errors.New("exit status 2")
			},
			expectedOutput: nil,
			isError:        true,
			errorMsg:       "error getting services from hab svc status",
		},
		{
			testCaseDescription: "Automate Error",
			input: func(cmd string) ([]byte, error) {
				if cmd == constants.AUTOMATESTATUSCMD {
					return []byte("Failed to execute automate command"), errors.New("exit status 2")
				}
				return []byte(habSvcStatusWithLicenseOutputOnA2), nil
			},
			expectedOutput: nil,
			isError:        true,
			errorMsg:       "error getting services from chef-automate status",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testCaseDescription, func(t *testing.T) {

			ss := statusservice.NewStatusService(tc.input, log)
			actualOutput, err := ss.GetServices()
			if tc.isError {
				assert.Error(t, err)
				assert.Equal(t, tc.errorMsg, err.Error())
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tc.expectedOutput, actualOutput)
		})
	}
}

func TestCheckIfBENode(t *testing.T) {
	log, err := logger.NewLogger("text", "debug")
	assert.NoError(t, err)
	ss := statusservice.NewStatusService(func(cmd string) ([]byte, error) {
		return nil, nil
	}, log)

	type testCaseInfo struct {
		testCaseDescription string
		input               string
		expected            bool
	}

	testCases := []testCaseInfo{
		{
			testCaseDescription: "Positive",
			input:               constants.AUTOMATESTATUSONBEERROR,
			expected:            true,
		},
		{
			testCaseDescription: "Negative",
			input:               "some error output",
			expected:            false,
		},
		{
			testCaseDescription: "Empty",
			input:               "",
			expected:            false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testCaseDescription, func(t *testing.T) {
			assert.Equal(t, tc.expected, ss.CheckIfBENode(tc.input))
		})
	}
}

func TestParseChefAutomateStatus(t *testing.T) {
	log, err := logger.NewLogger("text", "debug")
	assert.NoError(t, err)
	ss := statusservice.NewStatusService(func(cmd string) ([]byte, error) {
		return nil, nil
	}, log)

	type testCaseInfo struct {
		testCaseDescription string
		input               string
		expected            int
		isError             bool
		errorMsg            string
	}

	testCases := []testCaseInfo{
		{
			testCaseDescription: "A2 Node",
			input:               automateStatusOutputOnA2,
			expected:            35,
			isError:             false,
			errorMsg:            "",
		},
		{
			testCaseDescription: "CS Node",
			input:               automateStatusOutputOnCS,
			expected:            12,
			isError:             false,
			errorMsg:            "",
		},
		{
			testCaseDescription: "BE Node",
			input:               constants.AUTOMATESTATUSONBEERROR,
			expected:            0,
			isError:             true,
			errorMsg:            "no table found in output",
		},
		{
			testCaseDescription: "Empty Response",
			input:               "",
			expected:            0,
			isError:             true,
			errorMsg:            "no table found in output",
		},
		{
			testCaseDescription: "A2 Node Unhealthy",
			input:               automateStatusOutputOnA2Unhealthy,
			expected:            35,
			isError:             false,
			errorMsg:            "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testCaseDescription, func(t *testing.T) {
			out, err := ss.ParseChefAutomateStatus(tc.input)
			if tc.isError {
				assert.Error(t, err)
				assert.Equal(t, tc.errorMsg, err.Error())
			} else {
				assert.NoError(t, err)
			}
			if out != nil {
				assert.Equal(t, tc.expected, len(*out))
			}
		})
	}
}

func TestParseHabSvcStatus(t *testing.T) {

	type testCaseInfo struct {
		testCaseDescription string
		input               string
		expected            int
		isError             bool
		errorMsg            string
	}

	testCases := []testCaseInfo{
		{
			testCaseDescription: "A2 Node",
			input:               habSvcStatusWithLicenseOutputOnA2,
			expected:            35,
			isError:             false,
			errorMsg:            "",
		},
		{
			testCaseDescription: "CS Node",
			input:               habSvcStatusOutputOnCS,
			expected:            12,
			isError:             false,
			errorMsg:            "",
		},
		{
			testCaseDescription: "PG Node",
			input:               habSvcStatusOutputOnPG,
			expected:            3,
			isError:             false,
			errorMsg:            "",
		},
		{
			testCaseDescription: "OS Node",
			input:               habSvcStatusOutputOnOS,
			expected:            2,
			isError:             false,
			errorMsg:            "",
		},
		{
			testCaseDescription: "Empty Response",
			input:               "",
			expected:            0,
			isError:             true,
			errorMsg:            "no table found in output",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testCaseDescription, func(t *testing.T) {
			out, err := statusservice.ParseHabSvcStatus(tc.input)
			if tc.isError {
				assert.Error(t, err)
				assert.Equal(t, tc.errorMsg, err.Error())
			} else {
				assert.NoError(t, err)
			}
			if out != nil {
				assert.Equal(t, tc.expected, len(*out))
			}
		})
	}
}

func TestGetServicesFromHabSvcStatus(t *testing.T) {
	log, err := logger.NewLogger("text", "debug")
	assert.NoError(t, err)

	type testCaseInfo struct {
		testCaseDescription string
		input               func(cmd string) ([]byte, error)
		expected            int
		isError             bool
		errorMsg            string
	}

	testCases := []testCaseInfo{
		{
			testCaseDescription: "Sucessful Response",
			input: func(cmd string) ([]byte, error) {
				return []byte(habSvcStatusWithLicenseOutputOnA2), nil
			},
			expected: 35,
			isError:  false,
			errorMsg: "",
		},
		{
			testCaseDescription: "Failed Response",
			input: func(cmd string) ([]byte, error) {
				return []byte("Failed to execute command"), errors.New("exit status 1")
			},
			expected: 0,
			isError:  true,
			errorMsg: "error getting services from hab svc status",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testCaseDescription, func(t *testing.T) {

			ss := statusservice.NewStatusService(tc.input, log)
			out, err := ss.GetServicesFromHabSvcStatus()
			if tc.isError {
				assert.Error(t, err)
				assert.Equal(t, tc.errorMsg, err.Error())
			} else {
				assert.NoError(t, err)
			}
			if out != nil {
				assert.Equal(t, tc.expected, len(*out))
			}
		})
	}
}

func TestGetServicesFromAutomateStatus(t *testing.T) {
	log, err := logger.NewLogger("text", "debug")
	assert.NoError(t, err)

	type testCaseInfo struct {
		testCaseDescription string
		input               func(cmd string) ([]byte, error)
		expected            int
		isError             bool
		errorMsg            string
	}

	testCases := []testCaseInfo{
		{
			testCaseDescription: "Sucessful Response",
			input: func(cmd string) ([]byte, error) {
				return []byte(automateStatusOutputOnCS), nil
			},
			expected: 12,
			isError:  false,
			errorMsg: "",
		},
		{
			testCaseDescription: "BE Node",
			input: func(cmd string) ([]byte, error) {
				return []byte(constants.AUTOMATESTATUSONBEERROR), errors.New("exit status 90")
			},
			expected: 0,
			isError:  false,
			errorMsg: "",
		},
		{
			testCaseDescription: "Failed Response",
			input: func(cmd string) ([]byte, error) {
				return []byte("Failed to execute command"), errors.New("exit status 1")
			},
			expected: 0,
			isError:  true,
			errorMsg: "error getting services from chef-automate status",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testCaseDescription, func(t *testing.T) {

			ss := statusservice.NewStatusService(tc.input, log)
			out, err := ss.GetServicesFromAutomateStatus()
			if tc.isError {
				assert.Error(t, err)
				assert.Equal(t, tc.errorMsg, err.Error())
			} else {
				assert.NoError(t, err)
			}
			if out != nil {
				assert.Equal(t, tc.expected, len(*out))
			}
		})
	}
}
