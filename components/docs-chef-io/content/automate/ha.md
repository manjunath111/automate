+++
title = "High Availability Overview"

draft = false

gh_repo = "automate"
[menu]
  [menu.automate]
    title = "Overview"
    parent = "automate/deploy_high_availability"
    identifier = "automate/deploy_high_availability/ha.md High Availability Overview"
    weight = 10
+++

{{< note >}}
{{% automate/ha-warn %}}
{{< /note >}}

**High availability (HA)** refers to a system or application that offers high operational availability. This means the entire site or application will not be down if one server goes down due to traffic overload or other issues. HA represents the application remains available with no interruption. We achieve high availability when an application continues to operate even when one or more underlying components fail.

Thus, HA is designed to avoid loss of service by reducing or managing failures and minimizing unscheduled downtime (when your system or network is not available for use or is unresponsive) that happens due to power outages or failure of a component.

## Chef Automate High Availability (HA)

The Chef Automate HA equates to reliability, efficiency, and productivity, built on **Redundancy** and **Fail-over**. It aids in addressing significant issues like service failure and zone failure.

## Chef Automate HA Architecture

HA architecture includes the cluster of the *Chef Automate*, *Chef Server*, *PostgreSQL*, and *OpenSearch*.

{{< note >}}
Port **7799** must be accessible from the bastion host to all nodes within the Chef Automate cluster.
Although this requirement is not explicitly illustrated in the network architecture diagram for the sake of visual clarity, it is essential for proper cluster operation. The `chef-automate verify` command depends on successful connectivity to port **7799** on each node to perform its validations correctly.
{{< /note >}}

### Chef Automate HA Architecture for OnPremise / Cloud Non-Managed

![High Availability Architecture](/images/automate/ha_arch_onprem.png)

{{< note >}}
In Chef Automate HA architecture for On-Premise or non-managed Cloud deployments, frontend nodes connect to PostgreSQL over port **5432** and use port **6432** to perform leader checks.

Chef has deprecated the earlier configuration that required frontend nodes to use port **7432** for PostgreSQL connectivity.
{{< /note >}}

### Chef Automate HA Architecture for AWS Managed

![High Availability Architecture](/images/automate/ha_arch_aws_managedservices.png)

{{< note >}}
Chef Automate HA for Managed Services has default port 5432 for Managed PostgreSQL and 9200 for Managed OpenSearch. You can also change to your custom port.
{{< /note >}}

### Chef Automate HA Architecture for On-Premises Non-Managed Minimum Node Cluster

The following shows a five-node cluster, which is a supported deployment pattern. Work with your Progress technical teams to determine the appropriate cluster configuration for optimal performance based on parameters such as node count and data size.

![High Availability Architecture](/images/automate/ha_arch_minnode_cluster.png)

{{< note >}}
In Chef Automate HA architecture for On-Premise or non-managed Cloud deployments, frontend nodes connect to PostgreSQL over port **5432** and use port **6432** to perform leader checks.

Chef has deprecated the earlier configuration that required frontend nodes to use port **7432** for PostgreSQL connectivity.
{{< /note >}}

{{< warning >}}

- Choose Minimum node deployment type when you have VM constraints.
- Minimum node deployment is only for on-premises deployments
- Minimum node deployment is not supported for AWS deployments

{{< /warning >}}

## Chef Automate HA Topology

The Chef Automate HA Architecture involves the following clusters as part of the main cluster:

- **Backend Cluster** (Persistent Services)
  - **PostgreSQL:** Database requires a minimum of three nodes. PostgreSQL database uses the *Leader-Follower* strategy, where one becomes a leader, and the other two are the followers.

  - **OpenSearch:** Database requires a minimum of three nodes. OpenSearch database manages the [cluster internally](https://opensearch.org/docs/latest/opensearch/cluster/).

- **Frontend Cluster** (Application Services)
  - [Chef Automate](https://docs.chef.io/automate/)
  - [Chef Server](https://docs.chef.io/server/)

## Provisioning

Chef Automate's high availability solution can run on cloud providers and on-premises infrastructure systems. Appropriately provisioned backend, frontend, and bastion systems ensure a smooth deployment and installation experience.

- On-premises provisioning
- Cloud provisioning

### On-premises provisioning

  The customer can provision virtual machines or bare metal machines on a supported operating system with the required system settings to deploy the Automate HA solution.

### Cloud provisioning

Systems and services from the following cloud providers are supported:

- [AWS](https://docs.chef.io/automate/ha_aws_deploy_steps/#steps-to-provision)
- Azure
- Google

Deploy the Automate HA on the cloud infrastructure after provisioning the cloud systems. We have a simplified provisioning utility for AWS, Azure, and Google, and we expect to provision the systems manually.

## Deployment Methods

Chef Automate High Availability (HA) supports two types of deployment:

- [on-premises Deployment (Existing Node) Deployment](/automate/ha_onprim_deployment_procedure/)
- [Amazon Web Services (AWS) Deployment](/automate/ha_aws_deploy_steps/)

### On-premises Deployment (Existing Node/Bare Infrastructure)

In this, we expect VM (Virtual machine) or Bare Metal machines (Physical machine) that are already created and have initial Operating System (OS) setup done. Including Ports and Security policies changed according to requirements.

After this, installation steps will Deploy Chef Automate, Chef Infra Server, PostgreSQL DB, and OpenSearch DB to the relevant VMs or Physical Machines as provided in Config.

See the [performance benchmarking documentation](https://docs.chef.io/automate/ha_performance_benchmarks/#performance-benchmarks) for more information.

### Cloud Deployment using Amazon Web Services (AWS)

The two-step deployment process is as shown below:

- Provisioning Infrastructure. (Optional, if already manually done)
- Deployment of services on the provisioned infrastructure.
  - Installation of *PostgreSQL*, *OpenSearch*, *Chef Automate*, and *Chef Infra Server* will be done in this step.

### Cloud Deployment using Azure

The two-step deployment process is as shown below:

- Provisioning Infrastructure: Manually provision the infrastructure
- Deployment of services on the provisioned infrastructure (follow the [on-premises deployment procedure](/automate/ha_onprim_deployment_procedure/)).
  - Installation of *PostgreSQL*, *OpenSearch*, *Chef Automate*, and *Chef Infra Server* will be done in this step.
- Only File System Backup and Restore are supported.

### Cloud Deployment using Google Cloud Platform (GCP)

The two-step deployment process is as shown below:

- Provisioning Infrastructure: Manually provision the infrastructure
- Deployment of services on the provisioned infrastructure (follow the [on-premises deployment procedure](/automate/ha_onprim_deployment_procedure/)).
  - Installation of *PostgreSQL*, *OpenSearch*, *Chef Automate*, and *Chef Infra Server* will be done in this step.

## Performance (Benchmarking)

Please refer to the [Performance Benchmarking document](/automate/ha_performance_benchmarks/) for the detailed performance benchmark numbers
