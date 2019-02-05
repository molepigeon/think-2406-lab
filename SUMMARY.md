# Summary

## Getting started

* [Getting started](1_getting_started/1_cluster.md)

## Workshop

* [Part 1: Securing your IBM Cloud Kubernetes Service workloads by using IBM Cloud Container Registry](2_container_registry/1_registry.md)
  * [Creating a namespace](2_container_registry/2_onboarding.md)
  * [Configuring Identity and Access Management](2_container_registry/3_iam.md)
  * [Using IBM Cloud Container Registry with Kubernetes clusters](2_container_registry/4_kubernetes.md)
  * [Viewing image vulnerabilities](2_container_registry/5_vulnerability_advisor.md)
  * [Protecting your cluster from vulnerable images](2_container_registry/6_cise.md)
  * [Using content trust](2_container_registry/7_content_trust.md)
  * [Stretch goal: Trust pinning](2_container_registry/8_trust_pinning.md)
* [Part 2: Cluster-level logging by using IBM Log Analysis with LogDNA](3_logging/1_logging.md)
  * [Provision an instance of IBM Log Analysis with LogDNA](3_logging/provision_an_instance_of_ibm_log_analysis_with_log/README.md)
    * [Step 1: Grant permissions in the IBM Cloud](3_logging/provision_an_instance_of_ibm_log_analysis_with_log/step_1_grant_permissions_in_the_ibm_cloud.md)
    * [Step 2: Provisioning an instance from the Observability dashboard](3_logging/provision_an_instance_of_ibm_log_analysis_with_log/step_2_provisioning_an_instance_from_the_observabi.md)
  * [Configure a Kubernetes cluster to forward logs to your logging instance](3_logging/configure_a_kubernetes_cluster_to_forward_logs_to_.md)
  * [Navigate to the LogDNA web UI ](3_logging/navigate_to_the_logdna_web_ui.md)
  * [Monitor and manage logs through the LogDNA web UI](3_logging/monitor_and_manage_logs_through_the_logdna_web_ui/README.md)
    * [Create a view by filtering data through tags](3_logging/monitor_and_manage_logs_through_the_logdna_web_ui/create_a_view_by_filtering_data_through_tags.md)
    * [Optional - Export logs to a local file](3_logging/monitor_and_manage_logs_through_the_logdna_web_ui/export_logs_to_a_local_file.md)
    * [Optional - Configure an email alert](3_logging/monitor_and_manage_logs_through_the_logdna_web_ui/configure_an_email_alert.md)
  * [Customize the LogDNA agent ](3_logging/customize_the_logdna_agent/README.md)
    * [Add custom tags to your Kubernetes cluster](3_logging/customize_the_logdna_agent/add_custom_tags_to_your_kubernetes_cluster.md)
    * [Exclude log files](3_logging/customize_the_logdna_agent/exclude_log_files.md)
  * [Control what you log by using rules](3_logging/control_what_you_log_by_using_rules/README.md)
    * [Filter out log data from storage for a worker except error entries while keeping visibility of all lines](3_logging/control_what_you_log_by_using_rules/filter_out_log_data_from_storage_for_a_worker_exce.md)
    * [Filter out kube-system logs  from storage in a cluster except error entries](3_logging/control_what_you_log_by_using_rules/filter_out_kube-system_logs_from_storage_in_a_clus.md)
    * [Filter out logs from the UI from a namespace except error entries ](3_logging/control_what_you_log_by_using_rules/filter_out_logs_from_the_ui_from_a_namespace_excep.md)
* [Part 3: Monitoring a Kubernetes cluster in IBM Cloud by using IBM Cloud Monitoring with Sysdig](4_monitoring/1_monitoring.md)

## Feedback

* [Feedback](feedback.md)

## Legal

* [Disclaimer](disclaimer.md)
