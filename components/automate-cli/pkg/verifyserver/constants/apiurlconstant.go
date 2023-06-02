package constants

const (
	LOCAL_HOST_URL                           = "http://localhost"
	LOCALHOST                                = "127.0.0.1"
	CONTENT_TYPE                             = "Content-Type"
	TYPE_JSON                                = "application/json"
	HARDWARE_RESOURCE_CHECK_API_PATH         = "/api/v1/checks/hardware-resource-count"
	NFS_MOUNT_API_PATH                       = "/api/v1/checks/nfs-mount"
	NFS_MOUNT_LOC_API_PATH                   = "/api/v1/fetch/nfs-mount-loc"
	SOFTWARE_VERSION_CHECK_API_PATH          = "/api/v1/checks/software-versions"
	SYSTEM_RESOURCE_CHECK_API_PATH           = "/api/v1/checks/system-resource"
	SYSTEM_USER_CHECK_API_PATH               = "/api/v1/checks/system-user"
	FIREWALL_API_PATH                        = "/api/v1/checks/firewall"
	SSH_USER_CHECK_API_PATH                  = "/api/v1/checks/ssh-users"
	CERTIFICATE_CHECK_API_PATH               = "/api/v1/checks/certificate"
	S3_BACKUP_CHECK_API_PATH                 = "/api/v1/checks/s3-config"
	FQDN_LOAD_BALANCER_CHECK                 = "/api/v1/checks/fqdn"
	AWS_OPENSEARCH_S3_BUCKET_ACCESS_API_PATH = "/api/v1/checks/aws-opensearch-s3-bucket-access"
	PORT_REACHABLE_API_PATH                  = "/api/v1/checks/port-reachable"
)