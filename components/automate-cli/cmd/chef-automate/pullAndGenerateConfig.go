package main

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/url"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	dc "github.com/chef/automate/api/config/deployment"
	"github.com/chef/automate/components/automate-cli/pkg/status"
	mtoml "github.com/chef/automate/components/automate-deployment/pkg/toml"
	"github.com/chef/automate/lib/io/fileutils"
	"github.com/chef/automate/lib/stringutils"
	"github.com/chef/toml"
	"github.com/sirupsen/logrus"
)

const GET_OS_PASSWORD = "sudo HAB_LICENSE=accept-no-persist hab pkg exec chef/automate-platform-tools secrets-helper show userconfig.os_password"
const GET_AWS_OS_PASSWORD = "sudo HAB_LICENSE=accept-no-persist hab pkg exec chef/automate-platform-tools secrets-helper show userconfig.aws_os_password"
const GET_PG_SUPERUSER_PASSWORD = "sudo HAB_LICENSE=accept-no-persist hab pkg exec chef/automate-platform-tools secrets-helper show userconfig.pg_superuser_password"
const GET_PG_DBUSER_PASSWORD = "sudo HAB_LICENSE=accept-no-persist hab pkg exec chef/automate-platform-tools secrets-helper show userconfig.pg_dbuser_password"
const AUTOMATE_HA_WORKSPACE_GOOGLE_SERVICE_FILE = "/hab/a2_deploy_workspace/googleServiceAccount.json"

type ConfigKeys struct {
	rootCA     string
	privateKey string
	publicKey  string
	adminCert  string
	adminKey   string
}

type ObjectStorageConfig struct {
	accessKey                    string
	secrectKey                   string
	endpoint                     string
	bucketName                   string
	RoleArn                      string
	location                     string
	automateBasePath             string
	opensearchBasePath           string
	gcsServiceAccountCredentials string
}

type HAAwsAutoTfvars struct {
	AwsProfile                         string   `json:"aws_profile"`
	AwsRegion                          string   `json:"aws_region"`
	Endpoint                           string   `json:"endpoint"`
	SecretKey                          string   `json:"secret_key"`
	AccessKey                          string   `json:"access_key"`
	BucketName                         string   `json:"bucket_name"`
	Architecture                       string   `json:"architecture"`
	SshKeyFileName                     string   `json:"aws_ssh_key_pair_name"`
	AutomateDcToken                    string   `json:"automate_dc_token"`
	AwsVpcId                           string   `json:"aws_vpc_id"`
	AmiID                              string   `json:"aws_ami_id"`
	AwsCidrBlockAddr                   string   `json:"aws_cidr_block_addr"`
	PrivateCustomSubnets               []string `json:"private_custom_subnets"`
	PublicCustomSubnets                []string `json:"public_custom_subnets"`
	AutomateLbCertificateArn           string   `json:"automate_lb_certificate_arn"`
	ChefServerLbCertificateArn         string   `json:"chef_server_lb_certificate_arn"`
	AwsManagedRdsPostgresqlCertificate string   `json:"managed_rds_certificate"`
	ManagedRdsDbuserPassword           string   `json:"managed_rds_dbuser_password"`
	ManagedRdsDbuserUsername           string   `json:"managed_rds_dbuser_username"`
	ManagedRdsSuperuserPassword        string   `json:"managed_rds_superuser_password"`
	ManagedRdsSuperuserUsername        string   `json:"managed_rds_superuser_username"`
	ManagedRdsInstanceUrl              string   `json:"managed_rds_instance_url"`
	OsSnapshotUserAccessKeySecret      string   `json:"os_snapshot_user_access_key_secret"`
	OsSnapshotUserAccessKeyId          string   `json:"os_snapshot_user_access_key_id"`
	AwsOsSnapshotRoleArn               string   `json:"aws_os_snapshot_role_arn"`
	ManagedOpensearchCertificate       string   `json:"managed_opensearch_certificate"`
	ManagedOpensearchUserPassword      string   `json:"managed_opensearch_user_password"`
	ManagedOpensearchUsername          string   `json:"managed_opensearch_username"`
	ManagedOpensearchDomainUrl         string   `json:"managed_opensearch_domain_url"`
	ManagedOpensearchDomainName        string   `json:"managed_opensearch_domain_name"`
	AwsSetupManagedServices            bool     `json:"setup_managed_services"`
	ExistingPostgresqlPrivateIps       []string `json:"existing_postgresql_private_ips"`
	ExistingOpensearchPrivateIps       []string `json:"existing_opensearch_private_ips"`
	ExistingChefServerPrivateIps       []string `json:"existing_chef_server_private_ips"`
	ExistingAutomatePrivateIps         []string `json:"existing_automate_private_ips"`
	BackupConfigS3                     string   `json:"backup_config_s3"`
	BackupConfigEFS                    string   `json:"backup_config_efs"`
	LBAccessLogs                       string   `json:"lb_access_logs"`
	DeleteOnTermination                string   `json:"delete_on_termination"`
	AutomateServerInstanceType         string   `json:"automate_server_instance_type"`
	AutomateEbsVolumeIops              string   `json:"automate_ebs_volume_iops"`
	AutomateEbsVolumeSize              string   `json:"automate_ebs_volume_size"`
	AutomateEbsVolumeType              string   `json:"automate_ebs_volume_type"`
	ChefServerInstanceType             string   `json:"chef_server_instance_type"`
	ChefEbsVolumeIops                  string   `json:"chef_ebs_volume_iops"`
	ChefEbsVolumeSize                  string   `json:"chef_ebs_volume_size"`
	ChefEbsVolumeType                  string   `json:"chef_ebs_volume_type"`
	OpensearchServerInstanceType       string   `json:"opensearch_server_instance_type"`
	OpensearchEbsVolumeIops            string   `json:"opensearch_ebs_volume_iops"`
	OpensearchEbsVolumeSize            string   `json:"opensearch_ebs_volume_size"`
	OpensearchEbsVolumeType            string   `json:"opensearch_ebs_volume_type"`
	PostgresqlServerInstanceType       string   `json:"postgresql_server_instance_type"`
	PostgresqlEbsVolumeIops            string   `json:"postgresql_ebs_volume_iops"`
	PostgresqlEbsVolumeSize            string   `json:"postgresql_ebs_volume_size"`
	PostgresqlEbsVolumeType            string   `json:"postgresql_ebs_volume_type"`
}

type HATfvars struct {
	Region                       string      `json:"region"`
	Endpoint                     string      `json:"endpoint"`
	SecretKey                    string      `json:"secret_key"`
	AccessKey                    string      `json:"access_key"`
	BucketName                   string      `json:"bucket_name"`
	SshKeyFile                   string      `json:"ssh_key_file"`
	SshPort                      string      `json:"ssh_port"`
	SshUser                      string      `json:"ssh_user"`
	AutomateDcToken              string      `json:"automate_dc_token"`
	SSHGroupName                 string      `json:"ssh_group_name"`
	HabitatUidGid                string      `json:"habitat_uid_gid"`
	PostgresqlArchiveDiskFsPath  string      `json:"postgresql_archive_disk_fs_path"`
	PostgresqlInstanceCount      int         `json:"postgresql_instance_count"`
	NfsMountPath                 string      `json:"nfs_mount_path"`
	OpensearchCertsByIp          interface{} `json:"opensearch_certs_by_ip"`
	PostgresqlCertsByIp          interface{} `json:"postgresql_certs_by_ip"`
	ChefServerCertsByIp          interface{} `json:"chef_server_certs_by_ip"`
	AutomateCertsByIp            interface{} `json:"automate_certs_by_ip"`
	OpensearchNodesDn            string      `json:"opensearch_nodes_dn"`
	OpensearchAdminDn            string      `json:"opensearch_admin_dn"`
	OpensearchCustomCertsEnabled bool        `json:"opensearch_custom_certs_enabled"`
	PostgresqlCustomCertsEnabled bool        `json:"postgresql_custom_certs_enabled"`
	ChefServerCustomCertsEnabled bool        `json:"chef_server_custom_certs_enabled"`
	AutomateCustomCertsEnabled   bool        `json:"automate_custom_certs_enabled"`
	PostgresqlPublicKey          string      `json:"postgresql_public_key"`
	OpensearchAdminCert          string      `json:"opensearch_admin_cert"`
	OpensearchPublicKey          string      `json:"opensearch_public_key"`
	ChefServerPublicKey          string      `json:"chef_server_public_key"`
	AutomatePublicKey            string      `json:"automate_public_key"`
	PostgresqlPrivateKey         string      `json:"postgresql_private_key"`
	OpensearchPrivateKey         string      `json:"opensearch_private_key"`
	OpensearchAdminKey           string      `json:"opensearch_admin_key"`
	ChefServerPrivateKey         string      `json:"chef_server_private_key"`
	AutomatePrivateKey           string      `json:"automate_private_key"`
	PostgresqlRootCa             string      `json:"postgresql_root_ca"`
	OpensearchRootCa             string      `json:"opensearch_root_ca"`
	AutomateRootCa               string      `json:"automate_root_ca"`
	OpensearchInstanceCount      int         `json:"opensearch_instance_count"`
	ChefServerInstanceCount      int         `json:"chef_server_instance_count"`
	AutomateInstanceCount        int         `json:"automate_instance_count"`
	AutomateFqdn                 string      `json:"automate_fqdn"`
	AutomateConfigFile           string      `json:"automate_config_file"`
	OpensearchRootCert           string      `json:"opensearch_root_cert"`
	PostgresqlRootCert           string      `json:"postgresql_root_cert"`
	AwsVpcId                     string      `json:"aws_vpc_id"`
	AmiID                        string      `json:"ami_id"`
	AwsCidrBlockAddr             string      `json:"aws_cidr_block_addr"`
	PrivateCustomSubnets         []string    `json:"private_custom_subnets"`
	PublicCustomSubnets          []string    `json:"public_custom_subnets"`
	SSHKeyPairName               string      `json:"ssh_key_pair_name"`

	BucketNameDep                         string `json:"bucket_name_deployment"`
	ManagedOpensearchDomainUrlDep         string `json:"managed_opensearch_domain_url_deployment"`
	ManagedOpensearchUsernameDep          string `json:"managed_opensearch_username_deployment"`
	ManagedOpensearchUserPasswordDep      string `json:"managed_opensearch_user_password_deployment"`
	ManagedOpensearchCertificateDep       string `json:"managed_opensearch_certificate_deployment"`
	AwsOsSnapshotRoleArnDep               string `json:"aws_os_snapshot_role_arn_deployment"`
	OsSnapshotUserAccessKeyIdDep          string `json:"os_snapshot_user_access_key_id_deployment"`
	OsSnapshotUserAccessKeySecretDep      string `json:"os_snapshot_user_access_key_secret_deployment"`
	ManagedRdsInstanceUrlDep              string `json:"managed_rds_instance_url_deployment"`
	ManagedRdsSuperuserUsernameDep        string `json:"managed_rds_superuser_username_deployment"`
	ManagedRdsSuperuserPasswordDep        string `json:"managed_rds_superuser_password_deployment"`
	ManagedRdsDbuserUsernameDep           string `json:"managed_rds_dbuser_username_deployment"`
	ManagedRdsDbuserPasswordDep           string `json:"managed_rds_dbuser_password_deployment"`
	AwsManagedRdsPostgresqlCertificateDep string `json:"managed_rds_certificate_deployment"`

	ManagedRdsDbuserPassword      string   `json:"managed_rds_dbuser_password"`
	ManagedRdsDbuserUsername      string   `json:"managed_rds_dbuser_username"`
	ManagedRdsSuperuserPassword   string   `json:"managed_rds_superuser_password"`
	ManagedRdsSuperuserUsername   string   `json:"managed_rds_superuser_username"`
	ManagedRdsInstanceUrl         string   `json:"managed_rds_instance_url"`
	OsSnapshotUserAccessKeySecret string   `json:"os_snapshot_user_access_key_secret"`
	OsSnapshotUserAccessKeyId     string   `json:"os_snapshot_user_access_key_id"`
	AwsOsSnapshotRoleArn          string   `json:"aws_os_snapshot_role_arn"`
	ManagedOpensearchUserPassword string   `json:"managed_opensearch_user_password"`
	ManagedOpensearchUsername     string   `json:"managed_opensearch_username"`
	ManagedOpensearchDomainUrl    string   `json:"managed_opensearch_domain_url"`
	ManagedOpensearchDomainName   string   `json:"managed_opensearch_domain_name"`
	SetupSelfManagedServices      string   `json:"setup_self_managed_services"`
	SetupManagedServices          string   `json:"setup_managed_services"`
	ExistingPostgresqlPrivateIps  []string `json:"existing_postgresql_private_ips"`
	ExistingOpensearchPrivateIps  []string `json:"existing_opensearch_private_ips"`
	ExistingChefServerPrivateIps  []string `json:"existing_chef_server_private_ips"`
	ExistingAutomatePrivateIps    []string `json:"existing_automate_private_ips"`
	BackupConfigS3                string   `json:"backup_config_s3"`
	BackupConfigEFS               string   `json:"backup_config_efs"`
	AutomateAdminPassword         string   `json:"automate_admin_password"`
	TeamsPort                     int      `json:"teams_port"`
	SecretsKeyFile                string   `json:"secrets_key_file"`
	SecretsStoreFile              string   `json:"secrets_store_file"`
}
type PullConfigs interface {
	pullOpensearchConfigs(removeUnreachableNodes bool) (map[string]*ConfigKeys, []string, error)
	pullPGConfigs(removeUnreachableNodes bool) (map[string]*ConfigKeys, []string, error)
	pullAutomateConfigs(removeUnreachableNodes bool) (map[string]*dc.AutomateConfig, []string, error)
	pullChefServerConfigs(removeUnreachableNodes bool) (map[string]*dc.AutomateConfig, []string, error)
	fetchInfraConfig(removeUnreachableNodes bool) (*ExistingInfraConfigToml, map[string][]string, error)
	generateInfraConfig(removeUnreachableNodes bool) (*ExistingInfraConfigToml, map[string][]string, error)
	fetchAwsConfig(removeUnreachableNodes bool) (*AwsConfigToml, map[string][]string, error)
	generateAwsConfig(removeUnreachableNodes bool) (*AwsConfigToml, map[string][]string, error)
	getExceptionIps() []string
	setExceptionIps(ips []string)
	getOsCertsByIp(map[string]*ConfigKeys) []CertByIP
	setInfraAndSSHUtil(*AutomateHAInfraDetails, SSHUtil)
	getBackupPathFromAutomateConfig(a2ConfigMap map[string]*dc.AutomateConfig, backupLocation string) (string, error)
	getBackupPathFromOpensearchConfig() (string, error)
}

type PullConfigsImpl struct {
	infra        *AutomateHAInfraDetails
	sshUtil      SSHUtil
	exceptionIps []string
}

func NewPullConfigs(infra *AutomateHAInfraDetails, sshUtil SSHUtil) PullConfigs {
	return &PullConfigsImpl{
		infra:   infra,
		sshUtil: sshUtil,
	}
}

func (p *PullConfigsImpl) getExceptionIps() []string {
	return p.exceptionIps
}

func (p *PullConfigsImpl) setInfraAndSSHUtil(infra *AutomateHAInfraDetails, sshUtil SSHUtil) {
	p.infra = infra
	p.sshUtil = sshUtil
}

func (p *PullConfigsImpl) setExceptionIps(ips []string) {
	p.exceptionIps = ips
}

func (p *PullConfigsImpl) pullOpensearchConfigs(removeUnreachableNodes bool) (map[string]*ConfigKeys, []string, error) {
	var unreachableNodes []string
	ipConfigMap := make(map[string]*ConfigKeys)
	for _, ip := range p.infra.Outputs.OpensearchPrivateIps.Value {
		if stringutils.SliceContains(p.exceptionIps, ip) {
			continue
		}
		p.sshUtil.getSSHConfig().hostIP = ip
		scriptCommands := fmt.Sprintf(GET_CONFIG, opensearch_const)
		rawOutput, err := p.sshUtil.connectAndExecuteCommandOnRemote(scriptCommands, true)
		if err != nil {
			if removeUnreachableNodes {
				unreachableNodes = append(unreachableNodes, ip)
				continue
			}
			return nil, unreachableNodes, err
		}
		var src OpensearchConfig
		if _, err := toml.Decode(cleanToml(rawOutput), &src); err != nil {
			if removeUnreachableNodes {
				unreachableNodes = append(unreachableNodes, ip)
				continue
			}
			return nil, unreachableNodes, err
		}
		ipConfigMap[ip] = &ConfigKeys{
			rootCA:     src.TLS.RootCA,
			privateKey: src.TLS.SslKey,
			publicKey:  src.TLS.SslCert,
			adminCert:  src.TLS.AdminCert,
			adminKey:   src.TLS.AdminKey,
		}
	}
	return ipConfigMap, unreachableNodes, nil
}

func (p *PullConfigsImpl) pullPGConfigs(removeUnreachableNodes bool) (map[string]*ConfigKeys, []string, error) {
	var unreachableNodes []string
	ipConfigMap := make(map[string]*ConfigKeys)
	for _, ip := range p.infra.Outputs.PostgresqlPrivateIps.Value {
		if stringutils.SliceContains(p.exceptionIps, ip) {
			continue
		}
		p.sshUtil.getSSHConfig().hostIP = ip
		scriptCommands := fmt.Sprintf(GET_CONFIG, postgresql)
		rawOutput, err := p.sshUtil.connectAndExecuteCommandOnRemote(scriptCommands, true)
		if err != nil {
			if removeUnreachableNodes {
				unreachableNodes = append(unreachableNodes, ip)
				continue
			}
			return nil, unreachableNodes, err
		}
		var src PostgresqlConfig
		if _, err := toml.Decode(cleanToml(rawOutput), &src); err != nil {
			if removeUnreachableNodes {
				unreachableNodes = append(unreachableNodes, ip)
				continue
			}
			return nil, unreachableNodes, err
		}
		ipConfigMap[ip] = &ConfigKeys{
			rootCA:     src.Ssl.IssuerCert,
			privateKey: src.Ssl.SslKey,
			publicKey:  src.Ssl.SslCert,
		}
	}
	return ipConfigMap, unreachableNodes, nil
}

func (p *PullConfigsImpl) pullAutomateConfigs(removeUnreachableNodes bool) (map[string]*dc.AutomateConfig, []string, error) {
	var unreachableNodes []string
	ipConfigMap := make(map[string]*dc.AutomateConfig)
	for _, ip := range p.infra.Outputs.AutomatePrivateIps.Value {
		if stringutils.SliceContains(p.exceptionIps, ip) {
			continue
		}
		p.sshUtil.getSSHConfig().hostIP = ip
		rawOutput, err := p.sshUtil.connectAndExecuteCommandOnRemote(fmt.Sprintf(GET_FRONTEND_CONFIG, ""), true)
		if err != nil {
			if removeUnreachableNodes {
				unreachableNodes = append(unreachableNodes, ip)
				continue
			}
			return nil, unreachableNodes, err
		}
		var src dc.AutomateConfig
		if _, err := toml.Decode(cleanToml(rawOutput), &src); err != nil {
			if removeUnreachableNodes {
				unreachableNodes = append(unreachableNodes, ip)
				continue
			}
			return nil, unreachableNodes, err
		}
		ipConfigMap[ip] = &src
	}

	return ipConfigMap, unreachableNodes, nil

}

func (p *PullConfigsImpl) pullChefServerConfigs(removeUnreachableNodes bool) (map[string]*dc.AutomateConfig, []string, error) {
	var unreachableNodes []string
	ipConfigMap := make(map[string]*dc.AutomateConfig)
	for _, ip := range p.infra.Outputs.ChefServerPrivateIps.Value {
		if stringutils.SliceContains(p.exceptionIps, ip) {
			continue
		}
		p.sshUtil.getSSHConfig().hostIP = ip
		rawOutput, err := p.sshUtil.connectAndExecuteCommandOnRemote(fmt.Sprintf(GET_FRONTEND_CONFIG, ""), true)
		if err != nil {
			if removeUnreachableNodes {
				unreachableNodes = append(unreachableNodes, ip)
				continue
			}
			return nil, unreachableNodes, err
		}
		var src dc.AutomateConfig
		if _, err := toml.Decode(cleanToml(rawOutput), &src); err != nil {
			if removeUnreachableNodes {
				unreachableNodes = append(unreachableNodes, ip)
				continue
			}
			return nil, unreachableNodes, err
		}
		ipConfigMap[ip] = &src
	}
	return ipConfigMap, unreachableNodes, nil
}

type BkpLocation string

const BKP_LOCATION_S3 BkpLocation = "s3"
const BKP_LOCATION_FS BkpLocation = "fs"
const BKP_LOCATION_GCS BkpLocation = "gcs"

func determineBkpConfig(a2ConfigMap map[string]*dc.AutomateConfig, currConfig, objectStorage, fileSystem string) (string, string, error) {
	for _, ele := range a2ConfigMap {
		if ele.Global.V1.External.Opensearch != nil {
			osBkpLocation := ""
			if ele.Global.V1.External.Opensearch != nil &&
				ele.Global.V1.External.Opensearch.Backup != nil &&
				ele.Global.V1.External.Opensearch.Backup.Location != nil {
				osBkpLocation = ele.Global.V1.External.Opensearch.Backup.Location.Value
			}
			if ele.Global.V1.Backups == nil && ele.Global.V1.External.Opensearch.Backup == nil {
				return "", "", nil
			} else if ele.Global.V1.Backups != nil &&
				ele.Global.V1.Backups.Location != nil &&
				ele.Global.V1.Backups.Location.Value == string(BKP_LOCATION_S3) &&
				osBkpLocation == string(BKP_LOCATION_S3) {
				return objectStorage, osBkpLocation, nil
			} else if ele.Global.V1.Backups != nil &&
				ele.Global.V1.Backups.Filesystem != nil &&
				ele.Global.V1.Backups.Filesystem.Path != nil &&
				len(ele.Global.V1.Backups.Filesystem.Path.Value) > 0 &&
				osBkpLocation == string(BKP_LOCATION_FS) {
				return fileSystem, osBkpLocation, nil
			} else if ele.Global.V1.Backups != nil &&
				ele.Global.V1.Backups.Location != nil &&
				ele.Global.V1.Backups.Location.Value == string(BKP_LOCATION_GCS) &&
				osBkpLocation == string(BKP_LOCATION_GCS) {
				return objectStorage, osBkpLocation, nil
			} else {
				return "", "", errors.New("automate backup config mismatch in Global.V1.Backups and Global.V1.External.Opensearch.Backup")
			}
		} else {
			return "", "", errors.New("automate config Global.V1.External.Opensearch missing")
		}
	}
	return currConfig, "", nil
}

func determineDBType(a2ConfigMap map[string]*dc.AutomateConfig, dbtype string) (string, error) {
	if dbtype == TYPE_AWS || dbtype == TYPE_SELF_MANAGED {
		for _, ele := range a2ConfigMap {
			if ele.Global.V1.External.Opensearch != nil &&
				ele.Global.V1.External.Opensearch.Auth != nil &&
				ele.Global.V1.External.Opensearch.Auth.Scheme != nil {
				if ele.Global.V1.External.Opensearch.Auth.Scheme.Value == "basic_auth" {
					return TYPE_SELF_MANAGED, nil
				} else if ele.Global.V1.External.Opensearch.Auth.Scheme.Value == "aws_os" {
					return TYPE_AWS, nil
				} else {
					return "", errors.New("automate config Value in Global.V1.External.Opensearch.Auth.Scheme can be either basic_auth or aws_os")
				}
			} else {
				return "", errors.New("automate config error found")
			}
		}
	} else if dbtype == "" {
		return dbtype, nil
	}
	return "", errors.New(`unsupported db type. It should be either "aws" or "self-managed" or ""`)
}

func (p *PullConfigsImpl) fetchInfraConfig(removeUnreachableNodes bool) (*ExistingInfraConfigToml, map[string][]string, error) {
	unreachableNodes := make(map[string][]string)
	sharedConfigToml, err := getExistingHAConfig()
	if err != nil {
		return nil, nil, status.Wrap(err, status.ConfigError, "unable to fetch HA config")
	}
	a2ConfigMap, automateUnreachableNodes, err := p.pullAutomateConfigs(removeUnreachableNodes)
	if err != nil {
		return nil, nil, status.Wrap(err, status.ConfigError, "unable to fetch Automate config")
	}
	unreachableNodes[AUTOMATE] = automateUnreachableNodes

	csConfigMap, chefServerUnreachableNodes, err := p.pullChefServerConfigs(removeUnreachableNodes)
	if err != nil {
		return nil, unreachableNodes, status.Wrap(err, status.ConfigError, "unable to fetch Chef Server config")
	}
	unreachableNodes[CHEF_SERVER] = chefServerUnreachableNodes

	gcsObjStorageConfig, err := getGcsBackupConfig(a2ConfigMap, fileutils.NewFileSystemUtils())
	if err != nil {
		return nil, unreachableNodes, status.New(status.ConfigError, err.Error())
	}
	if len(gcsObjStorageConfig.location) > 0 {
		sharedConfigToml.ObjectStorage.Config.Location = gcsObjStorageConfig.location
	}
	if len(gcsObjStorageConfig.bucketName) > 0 {
		sharedConfigToml.ObjectStorage.Config.BucketName = gcsObjStorageConfig.bucketName
	}
	if len(gcsObjStorageConfig.gcsServiceAccountCredentials) > 0 {
		sharedConfigToml.ObjectStorage.Config.GoogleServiceAccountFile = AUTOMATE_HA_WORKSPACE_GOOGLE_SERVICE_FILE
	}

	bktype, bkpLocation, err := determineBkpConfig(a2ConfigMap, sharedConfigToml.Architecture.ConfigInitials.BackupConfig, "object_storage", "file_system")
	if err != nil {
		return nil, unreachableNodes, status.New(status.ConfigError, err.Error())
	}
	sharedConfigToml.Architecture.ConfigInitials.BackupConfig = bktype

	dbtype, err := determineDBType(a2ConfigMap, sharedConfigToml.ExternalDB.Database.Type)
	if err != nil {
		return nil, unreachableNodes, status.New(status.ConfigError, err.Error())
	}
	sharedConfigToml.ExternalDB.Database.Type = dbtype

	// checking onprem with managed or self managed services
	logrus.Debug(sharedConfigToml.ExternalDB.Database.Type)
	if len(strings.TrimSpace(sharedConfigToml.ExternalDB.Database.Type)) < 1 {
		osConfigMap, osUnreachableNodes, err := p.pullOpensearchConfigs(removeUnreachableNodes)
		if err != nil {
			return nil, unreachableNodes, status.Wrap(err, status.ConfigError, "unable to fetch Opensearch config")
		}
		unreachableNodes[OPENSEARCH] = osUnreachableNodes
		pgConfigMap, pgUnreachableNode, err := p.pullPGConfigs(removeUnreachableNodes)
		if err != nil {
			return nil, unreachableNodes, status.Wrap(err, status.ConfigError, "unable to fetch Postgresql config")
		}
		unreachableNodes[POSTGRESQL] = pgUnreachableNode
		// Build CertsByIP for Opensearch
		osCerts := p.getOsCertsByIp(osConfigMap)

		sharedConfigToml.Opensearch.Config.CertsByIP = osCerts

		if osRootCA := getOSORPGRootCA(osConfigMap); len(osRootCA) > 0 {
			sharedConfigToml.Opensearch.Config.RootCA = osRootCA
		}
		osAdminCert, osAdminKey := getOSAdminCertAndAdminKey(osConfigMap)
		if len(osAdminCert) > 0 {
			sharedConfigToml.Opensearch.Config.AdminCert = osAdminCert
		}
		if len(osAdminKey) > 0 {
			sharedConfigToml.Opensearch.Config.AdminKey = osAdminKey
		}
		adminDn, err := getDistinguishedNameFromKey(sharedConfigToml.Opensearch.Config.AdminCert)
		if err != nil {
			writer.Fail(err.Error())
		}
		sharedConfigToml.Opensearch.Config.AdminDn = strings.ReplaceAll(fmt.Sprintf("%v", adminDn), `\`, `\\\\\\\\`)
		sharedConfigToml.Opensearch.Config.EnableCustomCerts = true

		// Build CertsByIP for Postgresql
		var pgCerts []CertByIP
		for key, ele := range pgConfigMap {
			certByIP := CertByIP{
				IP:         key,
				PrivateKey: ele.privateKey,
				PublicKey:  ele.publicKey,
			}
			pgCerts = append(pgCerts, certByIP)
		}
		sharedConfigToml.Postgresql.Config.CertsByIP = pgCerts

		if pgRootCA := getOSORPGRootCA(pgConfigMap); len(pgRootCA) > 0 {
			sharedConfigToml.Postgresql.Config.RootCA = pgRootCA
		}
		sharedConfigToml.Postgresql.Config.EnableCustomCerts = true
	} else {
		externalOsDetails, err := p.getExternalOpensearchDetails(a2ConfigMap, sharedConfigToml.ExternalDB.Database.Type)
		if err != nil {
			return nil, unreachableNodes, err
		}
		if externalOsDetails != nil {
			sharedConfigToml.ExternalDB.Database.Opensearch.OpensearchDomainName = externalOsDetails.OpensearchDomainName
			sharedConfigToml.ExternalDB.Database.Opensearch.OpensearchInstanceURL = externalOsDetails.OpensearchInstanceURL
			sharedConfigToml.ExternalDB.Database.Opensearch.OpensearchRootCert = externalOsDetails.OpensearchRootCert
			sharedConfigToml.ExternalDB.Database.Opensearch.OpensearchSuperUserName = externalOsDetails.OpensearchSuperUserName
			sharedConfigToml.ExternalDB.Database.Opensearch.OpensearchSuperUserPassword = externalOsDetails.OpensearchSuperUserPassword
			sharedConfigToml.ExternalDB.Database.Opensearch.AWS.AwsOsSnapshotRoleArn = externalOsDetails.AWS.AwsOsSnapshotRoleArn
			sharedConfigToml.ExternalDB.Database.Opensearch.AWS.OsUserAccessKeyId = externalOsDetails.AWS.OsUserAccessKeyId
			sharedConfigToml.ExternalDB.Database.Opensearch.AWS.OsUserAccessKeySecret = externalOsDetails.AWS.OsUserAccessKeySecret
		}
		externalPgDetails, err := p.getExternalPGDetails(a2ConfigMap)
		if externalPgDetails != nil {
			sharedConfigToml.ExternalDB.Database.PostgreSQL.PostgreSQLDBUserName = externalPgDetails.PostgreSQLDBUserName
			sharedConfigToml.ExternalDB.Database.PostgreSQL.PostgreSQLDBUserPassword = externalPgDetails.PostgreSQLDBUserPassword
			sharedConfigToml.ExternalDB.Database.PostgreSQL.PostgreSQLInstanceURL = externalPgDetails.PostgreSQLInstanceURL
			sharedConfigToml.ExternalDB.Database.PostgreSQL.PostgreSQLRootCert = externalPgDetails.PostgreSQLRootCert
			sharedConfigToml.ExternalDB.Database.PostgreSQL.PostgreSQLSuperUserName = externalPgDetails.PostgreSQLSuperUserName
			sharedConfigToml.ExternalDB.Database.PostgreSQL.PostgreSQLSuperUserPassword = externalPgDetails.PostgreSQLSuperUserPassword
		}
	}

	// Build CertsByIP for Automate
	var a2Certs []CertByIP
	for key, ele := range a2ConfigMap {
		certByIP := CertByIP{
			IP:         key,
			PrivateKey: ele.Global.V1.FrontendTls[0].Key,
			PublicKey:  ele.Global.V1.FrontendTls[0].Cert,
		}
		a2Certs = append(a2Certs, certByIP)
	}
	sharedConfigToml.Automate.Config.CertsByIP = a2Certs

	sharedConfigToml.Automate.Config.EnableCustomCerts = true

	// Build CertsByIP for ChefServer
	var csCerts []CertByIP
	for key, ele := range csConfigMap {
		if len(ele.Global.V1.FrontendTls) > 0 {
			certByIP := CertByIP{
				IP:         key,
				PrivateKey: ele.Global.V1.FrontendTls[0].Key,
				PublicKey:  ele.Global.V1.FrontendTls[0].Cert,
			}
			csCerts = append(csCerts, certByIP)
		}
	}
	sharedConfigToml.ChefServer.Config.CertsByIP = csCerts

	if csRootCA := getRootCAFromCS(csConfigMap); len(csRootCA) > 0 {
		sharedConfigToml.Automate.Config.RootCA = csRootCA
	}

	if csToken := getTokenFromCS(csConfigMap); len(csToken) > 0 {
		sharedConfigToml.Architecture.ConfigInitials.AutomateDcToken = csToken
	}
	sharedConfigToml.ChefServer.Config.EnableCustomCerts = true

	objectStorageConfig := getS3BackConfig(a2ConfigMap)
	if len(objectStorageConfig.location) > 0 {
		sharedConfigToml.ObjectStorage.Config.Location = objectStorageConfig.location
	}
	if len(objectStorageConfig.accessKey) > 0 {
		sharedConfigToml.ObjectStorage.Config.AccessKey = objectStorageConfig.accessKey
	}
	if len(objectStorageConfig.secrectKey) > 0 {
		sharedConfigToml.ObjectStorage.Config.SecretKey = objectStorageConfig.secrectKey

	}
	if len(objectStorageConfig.bucketName) > 0 {
		sharedConfigToml.ObjectStorage.Config.BucketName = objectStorageConfig.bucketName
	}

	if len(objectStorageConfig.endpoint) > 0 {
		sharedConfigToml.ObjectStorage.Config.Endpoint = objectStorageConfig.endpoint
	}

	if bkpLocation == string(BKP_LOCATION_S3) {
		sharedConfigToml.Architecture.ConfigInitials.AutomateBasePath = objectStorageConfig.automateBasePath
		sharedConfigToml.Architecture.ConfigInitials.OpensearchBasePath = objectStorageConfig.opensearchBasePath
	} else if bkpLocation == string(BKP_LOCATION_FS) {
		var fullPathA2, fullPathOS string
		for _, ele := range a2ConfigMap {
			if ele.Global.V1.Backups != nil &&
				ele.Global.V1.Backups.Filesystem != nil &&
				ele.Global.V1.Backups.Filesystem.Path != nil &&
				len(ele.Global.V1.Backups.Filesystem.Path.Value) > 0 {
				fullPathA2 = ele.Global.V1.Backups.Filesystem.Path.Value
			}
			if ele.Global.V1.External.Opensearch != nil &&
				ele.Global.V1.External.Opensearch.Backup != nil &&
				ele.Global.V1.External.Opensearch.Backup.Fs != nil &&
				ele.Global.V1.External.Opensearch.Backup.Fs.Path != nil {
				fullPathOS = ele.Global.V1.External.Opensearch.Backup.Fs.Path.Value
			}
			common, a2Base, osBase := findCommonPath(fullPathA2, fullPathOS)
			sharedConfigToml.Architecture.ConfigInitials.AutomateBasePath = a2Base
			sharedConfigToml.Architecture.ConfigInitials.OpensearchBasePath = osBase
			sharedConfigToml.Architecture.ConfigInitials.BackupMount = common
		}
	} else if bkpLocation == string(BKP_LOCATION_GCS) {
		sharedConfigToml.Architecture.ConfigInitials.AutomateBasePath = gcsObjStorageConfig.automateBasePath
		sharedConfigToml.Architecture.ConfigInitials.OpensearchBasePath = gcsObjStorageConfig.opensearchBasePath
	}

	if len(sharedConfigToml.Architecture.ConfigInitials.AutomateBasePath) == 0 {
		sharedConfigToml.Architecture.ConfigInitials.AutomateBasePath = "automate"
	}
	if len(sharedConfigToml.Architecture.ConfigInitials.OpensearchBasePath) == 0 {
		sharedConfigToml.Architecture.ConfigInitials.OpensearchBasePath = "elasticsearch"
	}

	a2Fqdn := getA2fqdn(a2ConfigMap)
	if a2Fqdn != "" {
		sharedConfigToml.Automate.Config.Fqdn = a2Fqdn
	}

	fqdn, root_ca := getChefServerFqdnAndLBRootCA(a2ConfigMap)
	sharedConfigToml.ChefServer.Config.Fqdn, sharedConfigToml.ChefServer.Config.RootCA = fqdn, root_ca
	return sharedConfigToml, unreachableNodes, nil
}

// Function to find the common path and unique parts
func findCommonPath(path1, path2 string) (common, unique1, unique2 string) {
	// Check if paths start with "/"
	hasLeadingSlash1 := strings.HasPrefix(path1, "/")
	hasLeadingSlash2 := strings.HasPrefix(path2, "/")

	// Split paths into components
	components1 := strings.Split(filepath.Clean(path1), string(filepath.Separator))
	components2 := strings.Split(filepath.Clean(path2), string(filepath.Separator))

	// Handle leading slash explicitly
	if hasLeadingSlash1 {
		components1 = append([]string{""}, components1...)
	}
	if hasLeadingSlash2 {
		components2 = append([]string{""}, components2...)
	}

	// Find the common prefix
	var commonComponents []string
	i := 0
	for i < len(components1) && i < len(components2) && components1[i] == components2[i] {
		commonComponents = append(commonComponents, components1[i])
		i++
	}

	// Build the common path
	common = filepath.Join(commonComponents...)
	if len(commonComponents) > 0 && commonComponents[0] == "" {
		common = "/" + common // Ensure leading slash if present in input
	}

	// Build the unique parts
	unique1 = filepath.Join(components1[i:]...)
	unique2 = filepath.Join(components2[i:]...)

	// Adjust unique paths if no common prefix
	if common == "" {
		if hasLeadingSlash1 {
			unique1 = "/" + unique1
		}
		if hasLeadingSlash2 {
			unique2 = "/" + unique2
		}
	}

	return
}

func (p *PullConfigsImpl) getPasswordFromSecretHelper(pwdConfigValue string) (string, error) {
	for _, ip := range p.infra.Outputs.AutomatePrivateIps.Value {
		if stringutils.SliceContains(p.exceptionIps, ip) {
			continue
		}
		p.sshUtil.getSSHConfig().hostIP = ip
		rawOutput, err := p.sshUtil.connectAndExecuteCommandOnRemote(pwdConfigValue, true)
		if err != nil {
			return "", err
		}
		return strings.TrimSpace(rawOutput), nil
	}
	return "", nil
}

func (p *PullConfigsImpl) getExternalOpensearchDetails(a2ConfigMap map[string]*dc.AutomateConfig, dbType string) (*ExternalOpensearchToml, error) {
	for _, ele := range a2ConfigMap {
		roleArn := ""
		if ele.Global.V1.External.Opensearch.Backup != nil &&
			ele.Global.V1.External.Opensearch.Backup.S3 != nil &&
			ele.Global.V1.External.Opensearch.Backup.S3.Settings != nil &&
			ele.Global.V1.External.Opensearch.Backup.S3.Settings.RoleArn != nil {
			roleArn = ele.Global.V1.External.Opensearch.Backup.S3.Settings.RoleArn.Value
		}

		if dbType == TYPE_AWS {
			if ele.Global.V1.External.Opensearch != nil &&
				ele.Global.V1.External.Opensearch.Auth != nil &&
				ele.Global.V1.External.Opensearch.Auth.AwsOs != nil {
				var osPwd string
				if ele.Global.V1.External.Opensearch.Auth.AwsOs.Password != nil && ele.Global.V1.External.Opensearch.Auth.AwsOs.Password.Value != "" {
					osPwd = ele.Global.V1.External.Opensearch.Auth.AwsOs.Password.Value
				} else {
					osPass, err := p.getPasswordFromSecretHelper(GET_AWS_OS_PASSWORD)
					if err != nil {
						return nil, status.Wrap(err, status.ConfigError, "unable to fetch Opensearch password")
					}
					osPwd = osPass
				}
				return setExternalOpensearchDetails(ele.Global.V1.External.Opensearch.Nodes[0].Value,
					ele.Global.V1.External.Opensearch.Auth.AwsOs.Username.Value,
					base64.StdEncoding.EncodeToString([]byte(osPwd)),
					ele.Global.V1.External.Opensearch.Ssl.RootCert.Value,
					ele.Global.V1.External.Opensearch.Ssl.ServerName.Value,
					ele.Global.V1.External.Opensearch.Auth.AwsOs.AccessKey.Value,
					ele.Global.V1.External.Opensearch.Auth.AwsOs.SecretKey.Value,
					roleArn,
				), nil
			}
		} else if dbType == TYPE_SELF_MANAGED {
			osPass, err := p.getPasswordFromSecretHelper(GET_OS_PASSWORD)
			if err != nil {
				return nil, status.Wrap(err, status.ConfigError, "unable to fetch Opensearch password")
			}
			if ele.Global.V1.External.Opensearch != nil &&
				ele.Global.V1.External.Opensearch.Auth != nil &&
				ele.Global.V1.External.Opensearch.Auth.BasicAuth != nil {
				return setExternalOpensearchDetails(ele.Global.V1.External.Opensearch.Nodes[0].Value,
					ele.Global.V1.External.Opensearch.Auth.BasicAuth.Username.Value,
					base64.StdEncoding.EncodeToString([]byte(osPass)),
					ele.Global.V1.External.Opensearch.Ssl.RootCert.Value,
					ele.Global.V1.External.Opensearch.Ssl.ServerName.Value,
					"",
					"",
					roleArn,
				), nil
			}
		} else {
			return nil, status.New(status.ConfigError, fmt.Sprintf("Database type %s is invalid", dbType))
		}
	}

	return nil, nil
}

func setExternalOpensearchDetails(instanceUrl, superUserName, superPassword, rootCert, domainName, accessKey, secretKey, roleArn string) *ExternalOpensearchToml {
	nodeUrl, _ := url.Parse(instanceUrl)
	return &ExternalOpensearchToml{
		OpensearchInstanceURL:       nodeUrl.Host,
		OpensearchSuperUserName:     superUserName,
		OpensearchSuperUserPassword: superPassword,
		OpensearchRootCert:          rootCert,
		OpensearchDomainName:        domainName,
		AWS: ExternalAwsToml{
			OsUserAccessKeyId:     accessKey,
			OsUserAccessKeySecret: secretKey,
			AwsOsSnapshotRoleArn:  roleArn,
		},
	}
}

func (p *PullConfigsImpl) getExternalPGDetails(a2ConfigMap map[string]*dc.AutomateConfig) (*ExternalPostgreSQLToml, error) {
	for _, ele := range a2ConfigMap {
		if ele.Global.V1.External.Postgresql.Nodes != nil &&
			ele.Global.V1.External.Postgresql.Auth.Password.Superuser != nil &&
			ele.Global.V1.External.Postgresql.Auth.Password.Dbuser != nil {
			var spwd, dpwd string
			if ele.Global.V1.External.Postgresql.Auth.Password.Superuser.Password != nil && ele.Global.V1.External.Postgresql.Auth.Password.Superuser.Password.Value != "" {
				spwd = ele.Global.V1.External.Postgresql.Auth.Password.Superuser.Password.Value
			} else {
				supwd, err := p.getPasswordFromSecretHelper(GET_PG_SUPERUSER_PASSWORD)
				if err != nil {
					return nil, status.Wrap(err, status.ConfigError, "unable to fetch Postgres superuser password")
				}
				spwd = supwd
			}
			if ele.Global.V1.External.Postgresql.Auth.Password.Dbuser.Password != nil && ele.Global.V1.External.Postgresql.Auth.Password.Dbuser.Password.Value != "" {
				dpwd = ele.Global.V1.External.Postgresql.Auth.Password.Dbuser.Password.Value
			} else {
				dbpwd, err := p.getPasswordFromSecretHelper(GET_PG_DBUSER_PASSWORD)
				if err != nil {
					return nil, status.Wrap(err, status.ConfigError, "unable to fetch Postgres Dbuser password")
				}
				dpwd = dbpwd
			}
			return setExternalPGDetails(
				ele.Global.V1.External.Postgresql.Nodes[0].Value,
				ele.Global.V1.External.Postgresql.Auth.Password.Superuser.Username.Value,
				base64.StdEncoding.EncodeToString([]byte(spwd)),
				ele.Global.V1.External.Postgresql.Auth.Password.Dbuser.Username.Value,
				base64.StdEncoding.EncodeToString([]byte(dpwd)),
				ele.Global.V1.External.Postgresql.Ssl.RootCert.Value,
			), nil
		}
	}
	return nil, nil
}

func setExternalPGDetails(instanceUrl, superUserName, superUserPassword, dBUserName, dBUserPassword, rootCerts string) *ExternalPostgreSQLToml {
	return &ExternalPostgreSQLToml{
		PostgreSQLInstanceURL:       instanceUrl,
		PostgreSQLSuperUserName:     superUserName,
		PostgreSQLSuperUserPassword: superUserPassword,
		PostgreSQLDBUserName:        dBUserName,
		PostgreSQLDBUserPassword:    dBUserPassword,
		PostgreSQLRootCert:          rootCerts,
	}
}

func getChefServerFqdnAndLBRootCA(config map[string]*dc.AutomateConfig) (string, string) {
	for _, ele := range config {
		if ele.Global != nil && ele.Global.V1 != nil && ele.Global.V1.ChefServer != nil && len(strings.TrimSpace(ele.Global.V1.ChefServer.Fqdn.Value)) > 0 && len(strings.TrimSpace(ele.Global.V1.ChefServer.RootCa.Value)) > 0 {
			return ele.Global.V1.ChefServer.Fqdn.Value, ele.Global.V1.ChefServer.RootCa.Value
		}
	}
	return "", ""
}

func (p *PullConfigsImpl) getOsCertsByIp(osConfigMap map[string]*ConfigKeys) []CertByIP {
	var osCerts []CertByIP
	nodesDnMap := make(map[string]bool)
	allNodesDn := ""

	for i := 0; i < len(p.infra.Outputs.OpensearchPrivateIps.Value); i++ {
		currentIp := p.infra.Outputs.OpensearchPrivateIps.Value[i]

		if stringutils.SliceContains(p.exceptionIps, currentIp) {
			continue
		}

		nodeConfig, _ := osConfigMap[currentIp]
		if nodeConfig == nil {
			continue
		}
		nodeDn, err := getDistinguishedNameFromKey(nodeConfig.publicKey)
		if err != nil {
			writer.Fail(err.Error())
		}
		certByIP := CertByIP{
			IP:         currentIp,
			PrivateKey: strings.TrimSpace(nodeConfig.privateKey),
			PublicKey:  strings.TrimSpace(nodeConfig.publicKey),
		}

		nodeDnStr := fmt.Sprintf("%v", nodeDn)
		_, isPresent := nodesDnMap[nodeDnStr]

		nodeDnStrReplaced := strings.ReplaceAll(nodeDnStr, `\`, `\\\\`)

		if !isPresent {
			if len(allNodesDn) == 0 {
				allNodesDn = allNodesDn + fmt.Sprintf("%v", nodeDnStrReplaced) + "\\n  "
			} else {
				allNodesDn = allNodesDn + fmt.Sprintf("- %v", nodeDnStrReplaced) + "\\n  "
			}
		}

		nodesDnMap[nodeDnStr] = true
		osCerts = append(osCerts, certByIP)
	}

	for i := 0; i < len(osCerts); i++ {
		osCerts[i].NodesDn = strings.TrimSpace(allNodesDn)
	}

	return osCerts
}

func (p *PullConfigsImpl) generateInfraConfig(removeUnreachableNodes bool) (*ExistingInfraConfigToml, map[string][]string, error) {
	sharedConfigToml, unreachableNodes, err := p.fetchInfraConfig(removeUnreachableNodes)
	if err != nil {
		return nil, unreachableNodes, err
	}

	shardConfig, err := mtoml.Marshal(sharedConfigToml)
	if err != nil {
		return nil, unreachableNodes, status.Wrap(err, status.ConfigError, "unable to marshal config to file")
	}
	err = ioutil.WriteFile(filepath.Join(initConfigHabA2HAPathFlag.a2haDirPath, "config.toml"), shardConfig, 0644) // nosemgrep
	if err != nil {
		return nil, unreachableNodes, status.Wrap(err, status.ConfigError, "unable to write config toml to file")
	}

	return sharedConfigToml, unreachableNodes, nil
}

func (p *PullConfigsImpl) fetchAwsConfig(removeUnreachableNodes bool) (*AwsConfigToml, map[string][]string, error) {
	unreachableNodes := make(map[string][]string)
	sharedConfigToml, err := getAwsHAConfig()
	if err != nil {
		return nil, nil, status.Wrap(err, status.ConfigError, "unable to fetch HA config")
	}
	archBytes, err := ioutil.ReadFile(filepath.Join(initConfigHabA2HAPathFlag.a2haDirPath, "terraform", ".tf_arch")) // nosemgrep
	if err != nil {
		writer.Errorf("%s", err.Error())
		return nil, nil, err
	}
	var arch = strings.Trim(string(archBytes), "\n")
	sharedConfigToml.Architecture.ConfigInitials.Architecture = arch

	a2ConfigMap, automateUnreachableNodes, err := p.pullAutomateConfigs(removeUnreachableNodes)
	if err != nil {
		return nil, unreachableNodes, status.Wrap(err, status.ConfigError, "unable to fetch Automate config")
	}
	unreachableNodes[AUTOMATE] = automateUnreachableNodes
	csConfigMap, chefServerUnreachableNodes, err := p.pullChefServerConfigs(removeUnreachableNodes)
	if err != nil {
		return nil, unreachableNodes, status.Wrap(err, status.ConfigError, "unable to fetch Chef Server config")
	}
	unreachableNodes[CHEF_SERVER] = chefServerUnreachableNodes
	bktype, bkpLocation, err := determineBkpConfig(a2ConfigMap, sharedConfigToml.Architecture.ConfigInitials.BackupConfig, "s3", "efs")
	if err != nil {
		return nil, unreachableNodes, err
	}
	sharedConfigToml.Architecture.ConfigInitials.BackupConfig = bktype

	// checking AWS with managed services or Non managed services
	logrus.Debug(sharedConfigToml.Aws.Config.SetupManagedServices)
	if !isManagedServicesOn() {
		osConfigMap, opensearchUnreachableNodes, err := p.pullOpensearchConfigs(removeUnreachableNodes)
		if err != nil {
			return nil, unreachableNodes, status.Wrap(err, status.ConfigError, "unable to fetch Opensearch config")
		}
		unreachableNodes[OPENSEARCH] = opensearchUnreachableNodes
		pgConfigMap, pgUnreachableNodes, err := p.pullPGConfigs(removeUnreachableNodes)
		if err != nil {
			return nil, unreachableNodes, status.Wrap(err, status.ConfigError, "unable to fetch Postgresql config")
		}
		unreachableNodes[POSTGRESQL] = pgUnreachableNodes
		// Build CertsByIP for Opensearch
		osCerts := p.getOsCertsByIp(osConfigMap)

		sharedConfigToml.Opensearch.Config.CertsByIP = osCerts

		if osRootCA := getOSORPGRootCA(osConfigMap); len(osRootCA) > 0 {
			sharedConfigToml.Opensearch.Config.RootCA = osRootCA
		}
		osAdminCert, osAdminKey := getOSAdminCertAndAdminKey(osConfigMap)
		if len(osAdminCert) > 0 {
			sharedConfigToml.Opensearch.Config.AdminCert = osAdminCert
		}
		if len(osAdminKey) > 0 {
			sharedConfigToml.Opensearch.Config.AdminKey = osAdminKey
		}

		osPrivKey, osPubKey := getPrivateKeyAndPublicKeyFromBE(osConfigMap)
		if len(osPrivKey) > 0 {
			sharedConfigToml.Opensearch.Config.PrivateKey = osPrivKey
		}
		if len(osPubKey) > 0 {
			sharedConfigToml.Opensearch.Config.PublicKey = osPubKey
		}
		nodeDn, err := getDistinguishedNameFromKey(sharedConfigToml.Opensearch.Config.PublicKey)
		if err != nil {
			writer.Fail(err.Error())
		}
		adminDn, err := getDistinguishedNameFromKey(sharedConfigToml.Opensearch.Config.AdminCert)
		if err != nil {
			writer.Fail(err.Error())
		}
		sharedConfigToml.Opensearch.Config.NodesDn = strings.ReplaceAll(fmt.Sprintf("%v", nodeDn), `\`, `\\\\`)
		sharedConfigToml.Opensearch.Config.AdminDn = strings.ReplaceAll(fmt.Sprintf("%v", adminDn), `\`, `\\\\\\\\`)
		sharedConfigToml.Opensearch.Config.EnableCustomCerts = true

		// Build CertsByIP for Postgresql
		var pgCerts []CertByIP
		for key, ele := range pgConfigMap {
			certByIP := CertByIP{
				IP:         key,
				PrivateKey: ele.privateKey,
				PublicKey:  ele.publicKey,
			}
			pgCerts = append(pgCerts, certByIP)
		}
		sharedConfigToml.Postgresql.Config.CertsByIP = pgCerts

		if pgRootCA := getOSORPGRootCA(pgConfigMap); len(pgRootCA) > 0 {
			sharedConfigToml.Postgresql.Config.RootCA = pgRootCA
		}
		pgPrivKey, pgPubKey := getPrivateKeyAndPublicKeyFromBE(pgConfigMap)
		if len(pgPrivKey) > 0 {
			sharedConfigToml.Postgresql.Config.PrivateKey = pgPrivKey
		}
		if len(pgPubKey) > 0 {
			sharedConfigToml.Postgresql.Config.PublicKey = pgPubKey
		}
		sharedConfigToml.Postgresql.Config.EnableCustomCerts = true
	} else {
		externalOsDetails, err := p.getExternalOpensearchDetails(a2ConfigMap, TYPE_AWS)
		if err != nil {
			return nil, unreachableNodes, err
		}
		if externalOsDetails != nil {
			sharedConfigToml.Aws.Config.OpensearchDomainName = externalOsDetails.OpensearchDomainName
			sharedConfigToml.Aws.Config.OpensearchDomainUrl = externalOsDetails.OpensearchInstanceURL
			sharedConfigToml.Aws.Config.OpensearchCertificate = externalOsDetails.OpensearchRootCert
			sharedConfigToml.Aws.Config.OpensearchUsername = externalOsDetails.OpensearchSuperUserName
			sharedConfigToml.Aws.Config.OpensearchUserPassword = externalOsDetails.OpensearchSuperUserPassword
			sharedConfigToml.Aws.Config.AwsOsSnapshotRoleArn = externalOsDetails.AWS.AwsOsSnapshotRoleArn
			sharedConfigToml.Aws.Config.OsUserAccessKeyId = externalOsDetails.AWS.OsUserAccessKeyId
			sharedConfigToml.Aws.Config.OsUserAccessKeySecret = externalOsDetails.AWS.OsUserAccessKeySecret
		}
		externalPgDetails, err := p.getExternalPGDetails(a2ConfigMap)
		if externalPgDetails != nil {
			sharedConfigToml.Aws.Config.RDSDBUserName = externalPgDetails.PostgreSQLDBUserName
			sharedConfigToml.Aws.Config.RDSDBUserPassword = externalPgDetails.PostgreSQLDBUserPassword
			sharedConfigToml.Aws.Config.RDSInstanceUrl = externalPgDetails.PostgreSQLInstanceURL
			sharedConfigToml.Aws.Config.RDSCertificate = externalPgDetails.PostgreSQLRootCert
			sharedConfigToml.Aws.Config.RDSSuperUserName = externalPgDetails.PostgreSQLSuperUserName
			sharedConfigToml.Aws.Config.RDSSuperUserPassword = externalPgDetails.PostgreSQLSuperUserPassword
		}
	}

	// Build CertsByIP for Automate
	var a2Certs []CertByIP
	for key, ele := range a2ConfigMap {
		certByIP := CertByIP{
			IP:         key,
			PrivateKey: ele.Global.V1.FrontendTls[0].Key,
			PublicKey:  ele.Global.V1.FrontendTls[0].Cert,
		}
		a2Certs = append(a2Certs, certByIP)
	}
	sharedConfigToml.Automate.Config.CertsByIP = a2Certs

	a2Fqdn := getA2fqdn(a2ConfigMap)
	if len(a2Fqdn) > 0 {
		sharedConfigToml.Automate.Config.Fqdn = a2Fqdn
	}
	if a2PrivKey, a2PubKey := getPrivateAndPublicKeyFromFE(a2ConfigMap); len(a2PrivKey) > 0 && len(a2PubKey) > 0 {
		sharedConfigToml.Automate.Config.PrivateKey = a2PrivKey
		sharedConfigToml.Automate.Config.PublicKey = a2PubKey
	}
	sharedConfigToml.Automate.Config.EnableCustomCerts = true

	// Build CertsByIP for ChefServer
	var csCerts []CertByIP
	for key, ele := range csConfigMap {
		if len(ele.Global.V1.FrontendTls) > 0 {
			certByIP := CertByIP{
				IP:         key,
				PrivateKey: ele.Global.V1.FrontendTls[0].Key,
				PublicKey:  ele.Global.V1.FrontendTls[0].Cert,
			}
			csCerts = append(csCerts, certByIP)
		}
	}
	sharedConfigToml.ChefServer.Config.CertsByIP = csCerts

	if csRootCA := getRootCAFromCS(csConfigMap); len(csRootCA) > 0 {
		sharedConfigToml.Automate.Config.RootCA = csRootCA
	}
	if csToken := getTokenFromCS(csConfigMap); len(csToken) > 0 {
		sharedConfigToml.Architecture.ConfigInitials.AutomateDcToken = csToken
	}
	if csPrivKey, csPubKey := getPrivateAndPublicKeyFromFE(csConfigMap); len(csPrivKey) > 0 && len(csPubKey) > 0 {
		sharedConfigToml.ChefServer.Config.PrivateKey = csPrivKey
		sharedConfigToml.ChefServer.Config.PublicKey = csPubKey
	}
	sharedConfigToml.ChefServer.Config.EnableCustomCerts = true

	objStorageConfig := getOpenSearchObjectStorageConfig(a2ConfigMap)
	if len(objStorageConfig.bucketName) > 0 {
		sharedConfigToml.Architecture.ConfigInitials.S3BucketName = objStorageConfig.bucketName
	}
	if len(objStorageConfig.RoleArn) > 0 {
		sharedConfigToml.Aws.Config.AwsOsSnapshotRoleArn = objStorageConfig.RoleArn
	}
	if len(objStorageConfig.accessKey) > 0 {
		sharedConfigToml.Aws.Config.OsUserAccessKeyId = objStorageConfig.accessKey
	}
	if len(objStorageConfig.secrectKey) > 0 {
		sharedConfigToml.Aws.Config.OsUserAccessKeySecret = objStorageConfig.secrectKey
	}

	if bkpLocation == string(BKP_LOCATION_S3) {
		sharedConfigToml.Architecture.ConfigInitials.AutomateBasePath = objStorageConfig.automateBasePath
		sharedConfigToml.Architecture.ConfigInitials.OpensearchBasePath = objStorageConfig.opensearchBasePath
	} else if bkpLocation == string(BKP_LOCATION_FS) {
		var fullPathA2, fullPathOS string
		for _, ele := range a2ConfigMap {
			if ele.Global.V1.Backups != nil &&
				ele.Global.V1.Backups.Filesystem != nil &&
				ele.Global.V1.Backups.Filesystem.Path != nil &&
				len(ele.Global.V1.Backups.Filesystem.Path.Value) > 0 {
				fullPathA2 = ele.Global.V1.Backups.Filesystem.Path.Value
			}
			if ele.Global.V1.External.Opensearch != nil &&
				ele.Global.V1.External.Opensearch.Backup != nil &&
				ele.Global.V1.External.Opensearch.Backup.Fs != nil &&
				ele.Global.V1.External.Opensearch.Backup.Fs.Path != nil {
				fullPathOS = ele.Global.V1.External.Opensearch.Backup.Fs.Path.Value
			}
			common, a2Base, osBase := findCommonPath(fullPathA2, fullPathOS)
			sharedConfigToml.Architecture.ConfigInitials.AutomateBasePath = a2Base
			sharedConfigToml.Architecture.ConfigInitials.OpensearchBasePath = osBase
			sharedConfigToml.Architecture.ConfigInitials.BackupMount = common
		}
	}

	if len(sharedConfigToml.Architecture.ConfigInitials.AutomateBasePath) == 0 {
		sharedConfigToml.Architecture.ConfigInitials.AutomateBasePath = "automate"
	}
	if len(sharedConfigToml.Architecture.ConfigInitials.OpensearchBasePath) == 0 {
		sharedConfigToml.Architecture.ConfigInitials.OpensearchBasePath = "elasticsearch"
	}

	fqdn, root_ca := getChefServerFqdnAndLBRootCA(a2ConfigMap)
	sharedConfigToml.ChefServer.Config.Fqdn, sharedConfigToml.ChefServer.Config.RootCA = fqdn, root_ca

	return sharedConfigToml, unreachableNodes, nil
}

func (p *PullConfigsImpl) generateAwsConfig(removeUnreachableNodes bool) (*AwsConfigToml, map[string][]string, error) {
	sharedConfigToml, unreachableNodes, err := p.fetchAwsConfig(removeUnreachableNodes)
	if err != nil {
		return nil, unreachableNodes, err
	}

	shardConfig, err := mtoml.Marshal(sharedConfigToml)
	if err != nil {
		return nil, unreachableNodes, status.Wrap(err, status.ConfigError, "unable to marshal config to file")
	}
	err = ioutil.WriteFile(filepath.Join(initConfigHabA2HAPathFlag.a2haDirPath, "config.toml"), shardConfig, 0644) // nosemgrep
	if err != nil {
		return nil, unreachableNodes, status.Wrap(err, status.ConfigError, "unable to write config toml to file")
	}

	return sharedConfigToml, unreachableNodes, nil
}

func getExistingHAConfig() (*ExistingInfraConfigToml, error) {
	if checkSharedConfigFile() {
		sharedConfigToml, err := getExistingInfraConfig(filepath.Join(initConfigHabA2HAPathFlag.a2haDirPath, "config.toml"))
		if err != nil {
			return nil, err
		}
		return sharedConfigToml, nil
	} else {
		jsonString := convTfvarToJson(filepath.Join(initConfigHabA2HAPathFlag.a2haDirPath, "terraform", "terraform.tfvars"))
		tfvarConfig, err := getJsonFromTerraformTfVarsFile(jsonString)
		if err != nil {
			return nil, err
		}
		return getExistingHAConfigFromTFVars(tfvarConfig)
	}
}

func getAwsHAConfig() (*AwsConfigToml, error) {
	if checkSharedConfigFile() {
		sharedConfigToml, err := getAwsConfig(filepath.Join(initConfigHabA2HAPathFlag.a2haDirPath, "config.toml"))
		if err != nil {
			return nil, err
		}
		return sharedConfigToml, nil
	} else {
		jsonString := convTfvarToJson(filepath.Join(initConfigHabA2HAPathFlag.a2haDirPath, "terraform", "terraform.tfvars"))
		tfvarConfig, err := getJsonFromTerraformTfVarsFile(jsonString)
		if err != nil {
			return nil, err
		}
		AwsConfigJsonString := convTfvarToJson(filepath.Join(initConfigHabA2HAPathFlag.a2haDirPath, "terraform", "aws.auto.tfvars"))
		awsAutoTfvarConfig, err := getJsonFromTerraformAwsAutoTfVarsFile(AwsConfigJsonString)
		if err != nil {
			return nil, err
		}
		return getAwsHAConfigFromTFVars(tfvarConfig, awsAutoTfvarConfig)
	}
}

func getExistingHAConfigFromTFVars(tfvarConfig *HATfvars) (*ExistingInfraConfigToml, error) {
	sharedConfigToml := &ExistingInfraConfigToml{}
	sharedConfigToml.Architecture.ConfigInitials.Architecture = "existing_nodes"
	if len(strings.TrimSpace(tfvarConfig.SecretsKeyFile)) < 1 {
		sharedConfigToml.Architecture.ConfigInitials.SecretsKeyFile = filepath.Join(initConfigHabA2HAPathFlag.a2haDirPath, "secrets.key")
	} else {
		sharedConfigToml.Architecture.ConfigInitials.SecretsKeyFile = tfvarConfig.SecretsKeyFile
	}

	if len(strings.TrimSpace(tfvarConfig.SecretsStoreFile)) < 1 {
		sharedConfigToml.Architecture.ConfigInitials.SecretsStoreFile = filepath.Join(initConfigHabA2HAPathFlag.a2haDirPath, "secrets.json")
	} else {
		sharedConfigToml.Architecture.ConfigInitials.SecretsStoreFile = tfvarConfig.SecretsStoreFile
	}

	if strings.EqualFold(strings.TrimSpace(tfvarConfig.BackupConfigS3), "true") {
		sharedConfigToml.Architecture.ConfigInitials.BackupConfig = "object_storage"
	}
	if strings.EqualFold(strings.TrimSpace(tfvarConfig.BackupConfigEFS), "true") {
		sharedConfigToml.Architecture.ConfigInitials.BackupConfig = "file_system"
	}
	if tfvarConfig.TeamsPort > 0 {
		sharedConfigToml.Automate.Config.TeamsPort = strconv.Itoa(tfvarConfig.TeamsPort)
	}
	sharedConfigToml.Architecture.ConfigInitials.BackupMount = strings.TrimSpace(tfvarConfig.NfsMountPath)
	sharedConfigToml.Architecture.ConfigInitials.HabitatUIDGid = strings.TrimSpace(tfvarConfig.HabitatUidGid)
	sharedConfigToml.Architecture.ConfigInitials.SSHKeyFile = strings.TrimSpace(tfvarConfig.SshKeyFile)
	sharedConfigToml.Architecture.ConfigInitials.SSHPort = strings.TrimSpace(tfvarConfig.SshPort)
	if len(strings.TrimSpace(tfvarConfig.SshPort)) == 0 {
		sharedConfigToml.Architecture.ConfigInitials.SSHPort = "22"
	}

	sharedConfigToml.Architecture.ConfigInitials.SSHUser = strings.TrimSpace(tfvarConfig.SshUser)
	sharedConfigToml.Architecture.ConfigInitials.SSHGroupName = strings.TrimSpace(tfvarConfig.SSHGroupName)
	sharedConfigToml.Architecture.ConfigInitials.WorkspacePath = AUTOMATE_HA_WORKSPACE_DIR
	sharedConfigToml.Architecture.ConfigInitials.AutomateDcToken = strings.TrimSpace(tfvarConfig.AutomateDcToken)
	sharedConfigToml.Automate.Config.Fqdn = strings.TrimSpace(tfvarConfig.AutomateFqdn)
	sharedConfigToml.Automate.Config.InstanceCount = strconv.Itoa(tfvarConfig.AutomateInstanceCount)
	sharedConfigToml.Automate.Config.ConfigFile = strings.TrimSpace(tfvarConfig.AutomateConfigFile)
	sharedConfigToml.Automate.Config.EnableCustomCerts = tfvarConfig.AutomateCustomCertsEnabled
	sharedConfigToml.Automate.Config.AdminPassword = strings.TrimSpace(tfvarConfig.AutomateAdminPassword)
	sharedConfigToml.ChefServer.Config.EnableCustomCerts = tfvarConfig.ChefServerCustomCertsEnabled
	sharedConfigToml.ChefServer.Config.InstanceCount = strconv.Itoa(tfvarConfig.ChefServerInstanceCount)
	sharedConfigToml.Postgresql.Config.EnableCustomCerts = tfvarConfig.PostgresqlCustomCertsEnabled
	sharedConfigToml.Postgresql.Config.InstanceCount = strconv.Itoa(tfvarConfig.PostgresqlInstanceCount)
	sharedConfigToml.Opensearch.Config.EnableCustomCerts = tfvarConfig.OpensearchCustomCertsEnabled
	sharedConfigToml.Opensearch.Config.InstanceCount = strconv.Itoa(tfvarConfig.OpensearchInstanceCount)
	sharedConfigToml.Opensearch.Config.AdminDn = strings.TrimSpace(tfvarConfig.OpensearchAdminDn)
	sharedConfigToml.Opensearch.Config.NodesDn = strings.TrimSpace(tfvarConfig.OpensearchNodesDn)
	sharedConfigToml.ExistingInfra.Config.AutomatePrivateIps = tfvarConfig.ExistingAutomatePrivateIps
	sharedConfigToml.ExistingInfra.Config.ChefServerPrivateIps = tfvarConfig.ExistingChefServerPrivateIps
	sharedConfigToml.ExistingInfra.Config.PostgresqlPrivateIps = tfvarConfig.ExistingPostgresqlPrivateIps
	sharedConfigToml.ExistingInfra.Config.OpensearchPrivateIps = tfvarConfig.ExistingOpensearchPrivateIps
	sharedConfigToml.ObjectStorage.Config.AccessKey = strings.TrimSpace(tfvarConfig.AccessKey)
	sharedConfigToml.ObjectStorage.Config.SecretKey = strings.TrimSpace(tfvarConfig.SecretKey)
	sharedConfigToml.ObjectStorage.Config.BucketName = strings.TrimSpace(tfvarConfig.BucketName)
	sharedConfigToml.ObjectStorage.Config.Endpoint = strings.TrimSpace(tfvarConfig.Endpoint)
	sharedConfigToml.ObjectStorage.Config.Region = strings.TrimSpace(tfvarConfig.Region)
	sharedConfigToml.ExternalDB.Database.Opensearch.OpensearchInstanceURL = strings.TrimSpace(tfvarConfig.ManagedOpensearchDomainUrl)
	sharedConfigToml.ExternalDB.Database.Opensearch.OpensearchDomainName = strings.TrimSpace(tfvarConfig.ManagedOpensearchDomainName)
	sharedConfigToml.ExternalDB.Database.Opensearch.OpensearchRootCert = strings.TrimSpace(tfvarConfig.OpensearchRootCert)
	sharedConfigToml.ExternalDB.Database.Opensearch.OpensearchSuperUserName = strings.TrimSpace(tfvarConfig.ManagedOpensearchUsername)
	sharedConfigToml.ExternalDB.Database.Opensearch.OpensearchSuperUserPassword = strings.TrimSpace(tfvarConfig.ManagedOpensearchUserPassword)
	sharedConfigToml.ExternalDB.Database.Opensearch.AWS.AwsOsSnapshotRoleArn = strings.TrimSpace(tfvarConfig.AwsOsSnapshotRoleArn)
	sharedConfigToml.ExternalDB.Database.Opensearch.AWS.OsUserAccessKeyId = strings.TrimSpace(tfvarConfig.OsSnapshotUserAccessKeyId)
	sharedConfigToml.ExternalDB.Database.Opensearch.AWS.OsUserAccessKeySecret = strings.TrimSpace(tfvarConfig.OsSnapshotUserAccessKeySecret)
	sharedConfigToml.ExternalDB.Database.PostgreSQL.PostgreSQLDBUserName = strings.TrimSpace(tfvarConfig.ManagedRdsDbuserUsername)
	sharedConfigToml.ExternalDB.Database.PostgreSQL.PostgreSQLDBUserPassword = strings.TrimSpace(tfvarConfig.ManagedRdsDbuserPassword)
	sharedConfigToml.ExternalDB.Database.PostgreSQL.PostgreSQLInstanceURL = strings.TrimSpace(tfvarConfig.ManagedRdsInstanceUrl)
	sharedConfigToml.ExternalDB.Database.PostgreSQL.PostgreSQLRootCert = strings.TrimSpace(tfvarConfig.PostgresqlRootCert)
	sharedConfigToml.ExternalDB.Database.PostgreSQL.PostgreSQLSuperUserName = strings.TrimSpace(tfvarConfig.ManagedRdsSuperuserUsername)
	sharedConfigToml.ExternalDB.Database.PostgreSQL.PostgreSQLSuperUserPassword = strings.TrimSpace(tfvarConfig.ManagedRdsSuperuserPassword)
	setupManagedServices := strings.TrimSpace(tfvarConfig.SetupManagedServices)
	setupSelfManagedServices := strings.TrimSpace(tfvarConfig.SetupSelfManagedServices)
	if strings.EqualFold(setupManagedServices, "true") && strings.EqualFold(setupSelfManagedServices, "true") {
		sharedConfigToml.ExternalDB.Database.Type = "self-managed"
	}
	if strings.EqualFold(setupManagedServices, "true") && strings.EqualFold(setupSelfManagedServices, "false") {
		sharedConfigToml.ExternalDB.Database.Type = "aws"
	}
	return sharedConfigToml, nil
}

func getAwsHAConfigFromTFVars(tfvarConfig *HATfvars, awsAutoTfvarConfig *HAAwsAutoTfvars) (*AwsConfigToml, error) {
	sharedConfigToml := &AwsConfigToml{}
	sharedConfigToml.Architecture.ConfigInitials.Architecture = "deployment"

	if len(strings.TrimSpace(tfvarConfig.SecretsKeyFile)) < 1 {
		sharedConfigToml.Architecture.ConfigInitials.SecretsKeyFile = filepath.Join(initConfigHabA2HAPathFlag.a2haDirPath, "secrets.key")
	} else {
		sharedConfigToml.Architecture.ConfigInitials.SecretsKeyFile = tfvarConfig.SecretsKeyFile
	}

	if len(strings.TrimSpace(tfvarConfig.SecretsStoreFile)) < 1 {
		sharedConfigToml.Architecture.ConfigInitials.SecretsStoreFile = filepath.Join(initConfigHabA2HAPathFlag.a2haDirPath, "secrets.json")
	} else {
		sharedConfigToml.Architecture.ConfigInitials.SecretsStoreFile = tfvarConfig.SecretsStoreFile
	}

	if strings.EqualFold(strings.TrimSpace(awsAutoTfvarConfig.BackupConfigS3), "true") {
		sharedConfigToml.Architecture.ConfigInitials.BackupConfig = "s3"
	}

	if strings.EqualFold(strings.TrimSpace(awsAutoTfvarConfig.BackupConfigEFS), "true") {
		sharedConfigToml.Architecture.ConfigInitials.BackupConfig = "efs"
	}

	if tfvarConfig.TeamsPort > 0 {
		sharedConfigToml.Automate.Config.TeamsPort = strconv.Itoa(tfvarConfig.TeamsPort)
	}

	sharedConfigToml.Architecture.ConfigInitials.BackupMount = strings.TrimSpace(tfvarConfig.NfsMountPath)
	sharedConfigToml.Architecture.ConfigInitials.HabitatUIDGid = strings.TrimSpace(tfvarConfig.HabitatUidGid)
	sharedConfigToml.Architecture.ConfigInitials.Architecture = strings.TrimSpace(awsAutoTfvarConfig.Architecture)
	sharedConfigToml.Architecture.ConfigInitials.SSHKeyFile = strings.TrimSpace(tfvarConfig.SshKeyFile)
	sharedConfigToml.Architecture.ConfigInitials.SSHPort = strings.TrimSpace(tfvarConfig.SshPort)
	if len(strings.TrimSpace(tfvarConfig.SshPort)) == 0 {
		sharedConfigToml.Architecture.ConfigInitials.SSHPort = "22"
	}
	sharedConfigToml.Architecture.ConfigInitials.SSHUser = strings.TrimSpace(tfvarConfig.SshUser)
	sharedConfigToml.Architecture.ConfigInitials.SSHGroupName = strings.TrimSpace(tfvarConfig.SSHGroupName)
	sharedConfigToml.Architecture.ConfigInitials.AutomateDcToken = strings.TrimSpace(tfvarConfig.AutomateDcToken)
	sharedConfigToml.Architecture.ConfigInitials.WorkspacePath = AUTOMATE_HA_WORKSPACE_DIR
	sharedConfigToml.Automate.Config.InstanceCount = strconv.Itoa(tfvarConfig.AutomateInstanceCount)
	sharedConfigToml.Automate.Config.ConfigFile = strings.TrimSpace(tfvarConfig.AutomateConfigFile)
	sharedConfigToml.Automate.Config.EnableCustomCerts = tfvarConfig.AutomateCustomCertsEnabled
	sharedConfigToml.Automate.Config.AdminPassword = strings.TrimSpace(tfvarConfig.AutomateAdminPassword)
	sharedConfigToml.ChefServer.Config.EnableCustomCerts = tfvarConfig.ChefServerCustomCertsEnabled
	sharedConfigToml.ChefServer.Config.InstanceCount = strconv.Itoa(tfvarConfig.ChefServerInstanceCount)
	sharedConfigToml.Postgresql.Config.EnableCustomCerts = tfvarConfig.PostgresqlCustomCertsEnabled
	sharedConfigToml.Postgresql.Config.InstanceCount = strconv.Itoa(tfvarConfig.PostgresqlInstanceCount)
	sharedConfigToml.Opensearch.Config.EnableCustomCerts = tfvarConfig.OpensearchCustomCertsEnabled
	sharedConfigToml.Opensearch.Config.InstanceCount = strconv.Itoa(tfvarConfig.OpensearchInstanceCount)
	sharedConfigToml.Opensearch.Config.AdminDn = strings.TrimSpace(tfvarConfig.OpensearchAdminDn)
	sharedConfigToml.Opensearch.Config.NodesDn = strings.TrimSpace(tfvarConfig.OpensearchNodesDn)
	sharedConfigToml.Aws.Config.Profile = strings.TrimSpace(awsAutoTfvarConfig.AwsProfile)
	sharedConfigToml.Aws.Config.Region = strings.TrimSpace(awsAutoTfvarConfig.AwsRegion)
	sharedConfigToml.Aws.Config.AwsVpcId = strings.TrimSpace(awsAutoTfvarConfig.AwsVpcId)
	sharedConfigToml.Aws.Config.AmiID = strings.TrimSpace(awsAutoTfvarConfig.AmiID)
	sharedConfigToml.Aws.Config.SSHKeyPairName = strings.TrimSpace(awsAutoTfvarConfig.SshKeyFileName)
	sharedConfigToml.Aws.Config.AwsCidrBlockAddr = strings.TrimSpace(awsAutoTfvarConfig.AwsCidrBlockAddr)
	sharedConfigToml.Aws.Config.PrivateCustomSubnets = awsAutoTfvarConfig.PrivateCustomSubnets
	sharedConfigToml.Aws.Config.PublicCustomSubnets = awsAutoTfvarConfig.PublicCustomSubnets
	sharedConfigToml.Aws.Config.SSHKeyPairName = strings.TrimSpace(awsAutoTfvarConfig.SshKeyFileName)
	sharedConfigToml.Aws.Config.AutomateLbCertificateArn = strings.TrimSpace(awsAutoTfvarConfig.AutomateLbCertificateArn)
	sharedConfigToml.Aws.Config.ChefServerLbCertificateArn = strings.TrimSpace(awsAutoTfvarConfig.ChefServerLbCertificateArn)
	sharedConfigToml.Aws.Config.OpensearchDomainUrl = strings.TrimSpace(tfvarConfig.ManagedOpensearchDomainUrlDep)
	sharedConfigToml.Aws.Config.OpensearchDomainName = strings.TrimSpace(awsAutoTfvarConfig.ManagedOpensearchDomainName)
	sharedConfigToml.Aws.Config.OpensearchCertificate = strings.TrimSpace(tfvarConfig.ManagedOpensearchCertificateDep)
	sharedConfigToml.Aws.Config.OpensearchUsername = strings.TrimSpace(tfvarConfig.ManagedOpensearchUsernameDep)
	sharedConfigToml.Aws.Config.OpensearchUserPassword = strings.TrimSpace(tfvarConfig.ManagedOpensearchUserPasswordDep)
	sharedConfigToml.Aws.Config.AwsOsSnapshotRoleArn = strings.TrimSpace(tfvarConfig.AwsOsSnapshotRoleArnDep)
	sharedConfigToml.Aws.Config.OsUserAccessKeyId = strings.TrimSpace(tfvarConfig.OsSnapshotUserAccessKeyIdDep)
	sharedConfigToml.Aws.Config.OsUserAccessKeySecret = strings.TrimSpace(tfvarConfig.OsSnapshotUserAccessKeySecretDep)
	sharedConfigToml.Aws.Config.RDSCertificate = strings.TrimSpace(tfvarConfig.AwsManagedRdsPostgresqlCertificateDep)
	sharedConfigToml.Aws.Config.RDSDBUserName = strings.TrimSpace(tfvarConfig.ManagedRdsDbuserUsernameDep)
	sharedConfigToml.Aws.Config.RDSDBUserPassword = strings.TrimSpace(tfvarConfig.ManagedRdsDbuserPasswordDep)
	sharedConfigToml.Aws.Config.RDSInstanceUrl = strings.TrimSpace(tfvarConfig.ManagedRdsInstanceUrlDep)
	sharedConfigToml.Aws.Config.RDSSuperUserName = strings.TrimSpace(tfvarConfig.ManagedRdsSuperuserUsernameDep)
	sharedConfigToml.Aws.Config.RDSSuperUserPassword = strings.TrimSpace(tfvarConfig.ManagedRdsSuperuserPasswordDep)
	sharedConfigToml.Aws.Config.LBAccessLogs = strings.TrimSpace(awsAutoTfvarConfig.LBAccessLogs)
	sharedConfigToml.Aws.Config.DeleteOnTermination, _ = strconv.ParseBool(awsAutoTfvarConfig.DeleteOnTermination)
	sharedConfigToml.Aws.Config.AutomateServerInstanceType = strings.TrimSpace(awsAutoTfvarConfig.AutomateServerInstanceType)
	sharedConfigToml.Aws.Config.AutomateEbsVolumeIops = strings.TrimSpace(awsAutoTfvarConfig.AutomateEbsVolumeIops)
	sharedConfigToml.Aws.Config.AutomateEbsVolumeSize = strings.TrimSpace(awsAutoTfvarConfig.AutomateEbsVolumeSize)
	sharedConfigToml.Aws.Config.AutomateEbsVolumeType = strings.TrimSpace(awsAutoTfvarConfig.AutomateEbsVolumeType)
	sharedConfigToml.Aws.Config.ChefServerInstanceType = strings.TrimSpace(awsAutoTfvarConfig.ChefServerInstanceType)
	sharedConfigToml.Aws.Config.ChefEbsVolumeSize = strings.TrimSpace(awsAutoTfvarConfig.ChefEbsVolumeSize)
	sharedConfigToml.Aws.Config.ChefEbsVolumeType = strings.TrimSpace(awsAutoTfvarConfig.ChefEbsVolumeType)
	sharedConfigToml.Aws.Config.OpensearchServerInstanceType = strings.TrimSpace(awsAutoTfvarConfig.OpensearchServerInstanceType)
	sharedConfigToml.Aws.Config.OpensearchEbsVolumeIops = strings.TrimSpace(awsAutoTfvarConfig.OpensearchEbsVolumeIops)
	sharedConfigToml.Aws.Config.OpensearchEbsVolumeSize = strings.TrimSpace(awsAutoTfvarConfig.OpensearchEbsVolumeSize)
	sharedConfigToml.Aws.Config.OpensearchEbsVolumeType = strings.TrimSpace(awsAutoTfvarConfig.OpensearchEbsVolumeType)
	sharedConfigToml.Aws.Config.PostgresqlServerInstanceType = strings.TrimSpace(awsAutoTfvarConfig.PostgresqlServerInstanceType)
	sharedConfigToml.Aws.Config.PostgresqlEbsVolumeIops = strings.TrimSpace(awsAutoTfvarConfig.PostgresqlEbsVolumeIops)
	sharedConfigToml.Aws.Config.PostgresqlEbsVolumeSize = strings.TrimSpace(awsAutoTfvarConfig.PostgresqlEbsVolumeSize)
	sharedConfigToml.Aws.Config.PostgresqlEbsVolumeType = strings.TrimSpace(awsAutoTfvarConfig.PostgresqlEbsVolumeType)
	sharedConfigToml.Aws.Config.SetupManagedServices = awsAutoTfvarConfig.AwsSetupManagedServices
	sharedConfigToml.Aws.Config.ChefEbsVolumeIops, _ = getTheValueFromA2HARB("chef_ebs_volume_iops")
	return sharedConfigToml, nil
}

func getTheValueFromA2HARB(key string) (string, error) {
	wordCountCmd := `HAB_LICENSE=accept-no-persist hab pkg exec core/grep grep %s %s | wc -l`
	WordCountF := fmt.Sprintf(wordCountCmd, key, AUTOMATE_HA_WORKSPACE_A2HARB_FILE)
	output, err := exec.Command("/bin/sh", "-c", WordCountF).Output()
	if err != nil {
		return "", err
	}
	value := string(output)
	count, _ := strconv.Atoi(value)
	if count > 1 {
		logrus.Debug("The Key has more than one value it cannot be executed")
		return "", nil
	} else {
		GrepCmd := `HAB_LICENSE=accept-no-persist hab pkg exec core/grep grep %s %s | hab pkg exec core/gawk gawk '{print $2}'`
		GrepCmdF := fmt.Sprintf(GrepCmd, key, AUTOMATE_HA_WORKSPACE_A2HARB_FILE)
		output, err := exec.Command("/bin/sh", "-c", GrepCmdF).Output()
		if err != nil {
			return "", err
		}
		value := string(output)
		read_line := strings.TrimSuffix(value, "\n")
		KeyV, err := strconv.Unquote(read_line)
		if err != nil {
			return "", err
		}
		return strings.TrimSpace(KeyV), nil
	}
}

func getOSORPGRootCA(config map[string]*ConfigKeys) string {
	for _, ele := range config {
		if len(strings.TrimSpace(ele.rootCA)) > 0 {
			return strings.TrimSpace(ele.rootCA)
		}
	}
	return ""
}

func getRootCAFromCS(config map[string]*dc.AutomateConfig) string {
	if config == nil {
		return ""
	}
	for _, ele := range config {
		if ele.Global.V1.External != nil && ele.Global.V1.External.Automate != nil && ele.Global.V1.External.Automate.Ssl != nil && ele.Global.V1.External.Automate.Ssl.RootCert != nil {
			return ele.GetGlobal().V1.External.Automate.Ssl.RootCert.Value
		}
	}
	return ""
}

func getTokenFromCS(config map[string]*dc.AutomateConfig) string {
	if config == nil {
		return ""
	}
	for _, ele := range config {
		if ele.Global.V1.External != nil && ele.Global.V1.External.Automate != nil && ele.Global.V1.External.Automate.Auth != nil && ele.Global.V1.External.Automate.Auth.Token != nil {
			return ele.GetGlobal().V1.External.Automate.Auth.Token.Value
		}
	}
	return ""
}

func getPrivateAndPublicKeyFromFE(config map[string]*dc.AutomateConfig) (string, string) {
	if config == nil {
		return "", ""
	}
	for _, ele := range config {
		if ele.Global.V1.FrontendTls[0] != nil && ele.Global.V1.FrontendTls[0].Key != "" && ele.Global.V1.FrontendTls[0].Cert != "" {
			return ele.GetGlobal().V1.FrontendTls[0].Key, ele.GetGlobal().V1.FrontendTls[0].Cert
		}
	}
	return "", ""
}

func getOSAdminCertAndAdminKey(config map[string]*ConfigKeys) (string, string) {
	for _, ele := range config {
		if len(strings.TrimSpace(ele.adminCert)) > 0 && len(strings.TrimSpace(ele.adminKey)) > 0 {
			return strings.TrimSpace(ele.adminCert), strings.TrimSpace(ele.adminKey)
		} else if len(strings.TrimSpace(ele.adminCert)) > 0 {
			return strings.TrimSpace(ele.adminCert), ""
		} else if len(strings.TrimSpace(ele.adminKey)) > 0 {
			return "", strings.TrimSpace(ele.adminKey)
		}
	}
	return "", ""
}

func getPrivateKeyAndPublicKeyFromBE(config map[string]*ConfigKeys) (string, string) {
	for _, ele := range config {
		if len(strings.TrimSpace(ele.privateKey)) > 0 && len(strings.TrimSpace(ele.publicKey)) > 0 {
			return strings.TrimSpace(ele.privateKey), strings.TrimSpace(ele.publicKey)
		} else if len(strings.TrimSpace(ele.privateKey)) > 0 {
			return strings.TrimSpace(ele.privateKey), ""
		} else if len(strings.TrimSpace(ele.publicKey)) > 0 {
			return "", strings.TrimSpace(ele.publicKey)
		}
	}
	return "", ""
}

func getA2ORCSRootCA(config map[string]*dc.AutomateConfig) string {
	for _, ele := range config {
		if ele.Global.V1.Sys != nil && ele.Global.V1.Sys.Tls != nil && len(strings.TrimSpace(ele.Global.V1.Sys.Tls.RootCertContents)) > 0 {
			return ele.Global.V1.Sys.Tls.RootCertContents
		}
	}
	return ""
}

func getA2fqdn(config map[string]*dc.AutomateConfig) string {
	for _, ele := range config {
		if ele.Global != nil && ele.Global.V1 != nil && len(strings.TrimSpace(ele.Global.V1.Fqdn.Value)) > 0 {
			return ele.Global.V1.Fqdn.Value
		}
	}
	return ""
}

func getS3Bucket(config map[string]*dc.AutomateConfig) string {
	for _, ele := range config {
		if ele.Global.V1.External.Opensearch != nil && ele.Global.V1.External.Opensearch.Backup != nil && ele.Global.V1.External.Opensearch.Backup.S3 != nil && len(strings.TrimSpace(ele.Global.V1.External.Opensearch.Backup.S3.Bucket.Value)) > 0 {
			return ele.Global.V1.External.Opensearch.Backup.S3.Bucket.Value
		}
	}
	return ""
}

func getS3BucketOsBasePath(config map[string]*dc.AutomateConfig) string {
	for _, ele := range config {
		if ele.Global.V1.External.Opensearch != nil && ele.Global.V1.External.Opensearch.Backup != nil && ele.Global.V1.External.Opensearch.Backup.S3 != nil && ele.Global.V1.External.Opensearch.Backup.S3.BasePath != nil && len(strings.TrimSpace(ele.Global.V1.External.Opensearch.Backup.S3.BasePath.Value)) > 0 {
			return ele.Global.V1.External.Opensearch.Backup.S3.BasePath.Value
		}
	}
	return ""
}

func getS3BucketA2BasePath(config map[string]*dc.AutomateConfig) string {
	for _, ele := range config {
		if ele.Global.V1.Backups != nil &&
			ele.Global.V1.Backups.S3 != nil &&
			ele.Global.V1.Backups.S3.Bucket != nil &&
			ele.Global.V1.Backups.S3.Bucket.BasePath != nil {
			return ele.Global.V1.Backups.S3.Bucket.BasePath.Value
		}
	}
	return ""
}

func getOsRoleArn(config map[string]*dc.AutomateConfig) string {
	for _, ele := range config {
		if ele.Global.V1.External.Opensearch != nil && ele.Global.V1.External.Opensearch.Backup != nil && ele.Global.V1.External.Opensearch.Backup.S3 != nil && ele.Global.V1.External.Opensearch.Backup.S3.Settings != nil && len(strings.TrimSpace(ele.Global.V1.External.Opensearch.Backup.S3.Settings.RoleArn.Value)) > 0 {
			return ele.Global.V1.External.Opensearch.Backup.S3.Settings.RoleArn.Value
		}
	}
	return ""
}

func getOsAccessKey(config map[string]*dc.AutomateConfig) string {
	for _, ele := range config {
		if ele.Global.V1.External.Opensearch != nil && ele.Global.V1.External.Opensearch.Auth != nil && ele.Global.V1.External.Opensearch.Auth.AwsOs != nil && len(strings.TrimSpace(ele.Global.V1.External.Opensearch.Auth.AwsOs.AccessKey.Value)) > 0 {
			return ele.Global.V1.External.Opensearch.Auth.AwsOs.AccessKey.Value
		}
	}
	return ""
}

func getOsSecretKey(config map[string]*dc.AutomateConfig) string {
	for _, ele := range config {
		if ele.Global.V1.External.Opensearch != nil && ele.Global.V1.External.Opensearch.Auth != nil && ele.Global.V1.External.Opensearch.Auth.AwsOs != nil && len(strings.TrimSpace(ele.Global.V1.External.Opensearch.Auth.AwsOs.SecretKey.Value)) > 0 {
			return ele.Global.V1.External.Opensearch.Auth.AwsOs.SecretKey.Value
		}
	}
	return ""
}

func getOpenSearchObjectStorageConfig(config map[string]*dc.AutomateConfig) *ObjectStorageConfig {
	objStoage := &ObjectStorageConfig{}
	objStoage.accessKey = getOsAccessKey(config)
	objStoage.secrectKey = getOsSecretKey(config)
	objStoage.RoleArn = getOsRoleArn(config)
	objStoage.bucketName = getS3Bucket(config)
	objStoage.opensearchBasePath = getS3BucketOsBasePath(config)
	objStoage.automateBasePath = getS3BucketA2BasePath(config)
	return objStoage
}

func getJsonFromTerraformTfVarsFile(jsonString string) (*HATfvars, error) {
	params := HATfvars{}
	err := json.Unmarshal([]byte(jsonString), &params)
	if err != nil {
		writer.Fail(err.Error())
		return nil, err
	}
	return &params, nil

}
func getJsonFromTerraformAwsAutoTfVarsFile(jsonString string) (*HAAwsAutoTfvars, error) {
	params := HAAwsAutoTfvars{}
	err := json.Unmarshal([]byte(jsonString), &params)
	if err != nil {
		writer.Fail(err.Error())
		return nil, err
	}
	return &params, nil

}

func getModeOfDeployment() string {
	contentByte, err := ioutil.ReadFile(filepath.Join(initConfigHabA2HAPathFlag.a2haDirPath, "terraform", ".tf_arch")) // nosemgrep
	if err != nil {
		writer.Println(err.Error())
		return ""
	}
	deploymentMode := strings.TrimSpace(string(contentByte))
	if strings.EqualFold(deploymentMode, "existing_nodes") {
		return EXISTING_INFRA_MODE
	} else if strings.EqualFold(deploymentMode, "aws") {
		return AWS_MODE
	}
	return AWS_MODE
}

func getS3BackConfig(config map[string]*dc.AutomateConfig) *ObjectStorageConfig {
	objStoage := &ObjectStorageConfig{}
	for _, ele := range config {
		if ele.Global.V1.External != nil &&
			ele.Global.V1.External.Opensearch != nil &&
			ele.Global.V1.External.Opensearch.Backup != nil &&
			ele.Global.V1.External.Opensearch.Backup.S3 != nil &&
			ele.Global.V1.External.Opensearch.Backup.S3.BasePath != nil &&
			len(strings.TrimSpace(ele.Global.V1.External.Opensearch.Backup.S3.BasePath.Value)) > 0 {
			objStoage.opensearchBasePath = ele.Global.V1.External.Opensearch.Backup.S3.BasePath.Value
		}
		if ele.Global.V1.Backups != nil && ele.Global.V1.Backups.S3 != nil {
			if len(strings.TrimSpace(ele.Global.V1.Backups.Location.Value)) > 0 {
				objStoage.location = ele.Global.V1.Backups.Location.Value
			}
			if ele.Global.V1.Backups.S3.Credentials != nil {
				objStoage.accessKey = ele.Global.V1.Backups.S3.Credentials.AccessKey.Value
				objStoage.secrectKey = ele.Global.V1.Backups.S3.Credentials.SecretKey.Value
			}
			if ele.Global.V1.Backups.S3.Bucket != nil {
				objStoage.bucketName = ele.Global.V1.Backups.S3.Bucket.Name.Value
				objStoage.endpoint = ele.Global.V1.Backups.S3.Bucket.Endpoint.Value
			}
			if ele.Global.V1.Backups.S3.Bucket.BasePath != nil {
				objStoage.automateBasePath = ele.Global.V1.Backups.S3.Bucket.BasePath.Value
			}
			break
		}
	}
	return objStoage
}

func getGcsBackupConfig(a2ConfigMap map[string]*dc.AutomateConfig, fileUtils fileutils.FileUtils) (*ObjectStorageConfig, error) {
	objStoage := new(ObjectStorageConfig)
	for _, ele := range a2ConfigMap {
		if ele.Global.V1 != nil &&
			ele.Global.V1.Backups != nil &&
			ele.Global.V1.Backups.Gcs != nil &&
			ele.Global.V1.Backups.Location != nil {
			if len(strings.TrimSpace(ele.Global.V1.Backups.Location.Value)) > 0 {
				objStoage.location = ele.Global.V1.Backups.Location.Value
			}
			if ele.Global.V1.Backups.Gcs.Bucket != nil && len(strings.TrimSpace(ele.Global.V1.Backups.Gcs.Bucket.Name.Value)) > 0 {
				objStoage.bucketName = ele.Global.V1.Backups.Gcs.Bucket.Name.Value
			}
			if ele.Global.V1.Backups.Gcs.Bucket != nil &&
				ele.Global.V1.Backups.Gcs.Bucket.BasePath != nil &&
				len(strings.TrimSpace(ele.Global.V1.Backups.Gcs.Bucket.BasePath.Value)) > 0 {
				objStoage.automateBasePath = ele.Global.V1.Backups.Gcs.Bucket.BasePath.Value
			}

			if ele.Global.V1.External != nil &&
				ele.Global.V1.External.Opensearch != nil &&
				ele.Global.V1.External.Opensearch.Backup != nil &&
				ele.Global.V1.External.Opensearch.Backup.Gcs != nil &&
				ele.Global.V1.External.Opensearch.Backup.Gcs.BasePath != nil &&
				len(strings.TrimSpace(ele.Global.V1.External.Opensearch.Backup.Gcs.BasePath.Value)) > 0 {
				objStoage.opensearchBasePath = ele.Global.V1.External.Opensearch.Backup.Gcs.BasePath.Value
			}

			if ele.Global.V1.Backups.Gcs.Credentials != nil && len(strings.TrimSpace(ele.Global.V1.Backups.Gcs.Credentials.Json.Value)) > 0 {
				objStoage.gcsServiceAccountCredentials = ele.Global.V1.Backups.Gcs.Credentials.Json.Value
			}
			break
		}
	}
	// Update the /hab/a2_deploy_workspace/googleServiceAccount.json file (if required)
	if len(objStoage.gcsServiceAccountCredentials) > 0 {
		err := fileUtils.WriteFile(AUTOMATE_HA_WORKSPACE_GOOGLE_SERVICE_FILE, []byte(objStoage.gcsServiceAccountCredentials), 0644)
		if err != nil {
			return nil, err
		}
	}
	return objStoage, nil
}

func (p *PullConfigsImpl) getBackupPathFromAutomateConfig(a2ConfigMap map[string]*dc.AutomateConfig, backupLocation string) (string, error) {
	for _, ele := range a2ConfigMap {
		path := ""
		if ele.Global.V1.External.Opensearch != nil && ele.Global.V1.External.Opensearch.Backup != nil {
			switch backupLocation {
			case "fs":
				if ele.Global.V1.External.Opensearch.Backup.Fs != nil &&
					ele.Global.V1.External.Opensearch.Backup.Fs.Path != nil {
					path = ele.Global.V1.External.Opensearch.Backup.Fs.Path.GetValue()
					logrus.Debugf("Backup path set for opensearch settings in automate config: %s and backup location: %s", path, backupLocation)
					return path, nil
				}
			case "s3":
				if ele.Global.V1.External.Opensearch.Backup.S3 != nil &&
					ele.Global.V1.External.Opensearch.Backup.S3.Bucket != nil &&
					ele.Global.V1.External.Opensearch.Backup.S3.BasePath != nil {
					path = ele.Global.V1.External.Opensearch.Backup.S3.BasePath.GetValue()
					logrus.Debugf("Backup path set for opensearch settings in automate config: %s and backup location: %s", path, backupLocation)
					return path, nil
				}
			case "gcs":
				if ele.Global.V1.External.Opensearch.Backup.Gcs != nil &&
					ele.Global.V1.External.Opensearch.Backup.Gcs.Bucket != nil &&
					ele.Global.V1.External.Opensearch.Backup.Gcs.BasePath != nil {
					path = ele.Global.V1.External.Opensearch.Backup.Gcs.BasePath.GetValue()
					logrus.Debugf("Backup path set for opensearch settings in automate config: %s and backup location: %s", path, backupLocation)
					return path, nil
				}
			}
		}
		return "", errors.New("automate config Global.V1.External.Opensearch configurations are missing")
	}
	return "", errors.New("backup path from automate node could not be determined")
}

func (p *PullConfigsImpl) getBackupPathFromOpensearchConfig() (string, error) {
	if len(p.infra.Outputs.OpensearchPrivateIps.Value) == 0 {
		return "", errors.New("the cluster has no opensearch nodes")
	}
	ip := p.infra.Outputs.OpensearchPrivateIps.Value[0]
	p.sshUtil.getSSHConfig().hostIP = ip
	scriptCommands := fmt.Sprintf(GET_CONFIG, opensearch_const)
	rawOutput, err := p.sshUtil.connectAndExecuteCommandOnRemote(scriptCommands, true)
	if err != nil {
		return "", err
	}
	var src OpensearchConfig
	if _, err := toml.Decode(cleanToml(rawOutput), &src); err != nil {
		return "", err
	}
	logrus.Debugf("Backup path from opensearch config: %s", src.Path.Repo)
	return src.Path.Repo, nil
}
