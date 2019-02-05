# Provision an instance of IBM Cloud Monitoring with Sysdig {#provision-an-instance-of-ibm-cloud-monitoring-with-sysdig}

Before you can monitor and manage metrics with Sysdig, you must provision an instance of the IBM Cloud Monitoring with Sysdig service in IBM Cloud.

You can provision an instance of IBM Cloud Monitoring with Sysdig in any of the following ways:

*   From the Observability section in the IBM Cloud UI
*   [From the Catalog](https://console.bluemix.net/docs/services/Log-Analysis-with-LogDNA/provision.html)
*   [From the command line](https://console.bluemix.net/docs/services/Log-Analysis-with-LogDNA/provision.html)

You provision an instance within a resource group.

**Note**: After you provision an instance in a resource group, you cannot change it.

To provision an instance, your IBM ID needs 2 different permissions:

1.  Permission within a resource group. This is the resource group where you are going to provision the monitoring service.
2.  Permission to work with the IBM Cloud Monitoring with Sysdig service. Your user ID must have a policy for the service with platform permissions in the IBM Cloud that allow you to provision it in a resource group and manage it. For example, you must have a policy for the service IBM Cloud Monitoring with Sysdig with one of the following platform roles: administrator, editor