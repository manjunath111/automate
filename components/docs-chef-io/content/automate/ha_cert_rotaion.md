+++
title = "Certificate Rotation"

draft = false

gh_repo = "automate"

[menu]
  [menu.automate]
    title = "Certificate Rotation"
    parent = "automate/deploy_high_availability/certificates"
    identifier = "automate/deploy_high_availability/certificates/ha_cert_rotaion.md Certificate Rotation"
    weight = 220
+++

## What are Certificates?

A security certificate is a small data file used as an Internet security technique through which the identity, authenticity, and reliability of a website or Web application is established.

Certificates should be rotated periodically to ensure optimal security.

## How is Certificate Rotation Helpful?

Certificate rotation means the replacement of existing certificates with new ones when any certificate expires or is based on your organization's policy. A new CA authority is substituted for the old requiring a replacement of root certificate for the cluster.

The certificate rotation is also required when the key for a node, client, or CA is compromised. If compromised, you need to change the contents of a certificate, for example, to add another DNS name or the IP address of a load balancer through which a node can be reached.  In this case, you would need to rotate only the node certificates.

## How to Rotate the Certificates

You can generate the required certificates on your own or use your organization's existing certificates. Ensure you execute all the below commands from the `cd /hab/a2_deploy_workspace` path.

Follow these steps to rotate your certificates that are to be used in Chef Automate High Availability (HA):

1. Navigate to your workspace folder. For example, `cd /hab/a2_deploy_workspace`.
1. Execute the command, `./scripts/credentials set ssl --rotate all`. This command rotates all the certificates of your organization.

  {{< note >}}

  When you run this command first time, a series of certificates are created and saved in `/hab/a2_deploy_workspace/certs` location. You need to identify the appropriate certificates. For example, to rotate certificates for PostgreSQL, use certificate values into *pg_ssl_private.key*,  *pg_ssl_public.pem*, and *ca_root.pem*. Likewise, to rotate certificates for ElasticSearch, use certificate values into *ca_root.pem*, *es_admin_ssl_private.key*, *es_admin_ssl_public.pem*, *es_ssl_private.key*, *es_ssl_public.pem*, *kibana_ssl_private.key*, *kibana_ssl_public.pem*.

  {{< /note >}}

1. For rotating the PostgreSQL certificates, execute the command `./scripts/credentials set ssl --pg-ssl`.

1. For rotating the elasticsearch certificates, execute the command, `./scripts/credentials set ssl --es-ssl` .

<!-- 4. Copy your *x.509 SSL certs* into the appropriate files in `certs/` folder. -->

<!-- - Place your root certificate into `ca_root.pem file`. -->

<!-- - Place your intermediate CA into the `pem` file. -->

1. If your organization issues a certificate from an intermediate CA, place the respective certificate after the server certificate as per order listed. For example, in `certs/pg_ssl_public.pem`, paste it as them as listed:

   - Server Certificate
   - Intermediate CA Certificate 1
   - Intermediate CA Certificate n

1. Execute the command, `./scripts/credentials set ssl` (with the appropriate options). This command deploys the nodes.

1. Type the command, `./scripts/credentials set ssl  --help`. This command provides you with information and a list of commands related to certificate rotation.

1. For rotating the PostgreSQL credentials, execute the command `./scripts/credentials set postgresql --auto`.

1. For rotating the elasticsearch credentials, execute the command, `./scripts/credentials set elasticsearch --auto`.

{{< figure src="/images/automate/ha_cert_rotation.png" alt="Certificate Rotation">}}