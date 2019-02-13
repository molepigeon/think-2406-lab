# Summary

## Getting started

* [Getting started](1_getting_started/1_cluster.md)

## Workshop

* [Contents](parts.md)
  * [Securing your IBM Cloud Kubernetes Service workloads by using IBM Cloud Container Registry](2_container_registry/1_registry.md)
    * [Creating a namespace](2_container_registry/2_onboarding.md)
    * [Configuring Identity and Access Management](2_container_registry/3_iam.md)
    * [Using IBM Cloud Container Registry with Kubernetes clusters](2_container_registry/4_kubernetes.md)
    * [Viewing image vulnerabilities](2_container_registry/5_vulnerability_advisor.md)
    * [Protecting your cluster from vulnerable images](2_container_registry/6_cise.md)
    * [Using content trust](2_container_registry/7_content_trust.md)
    * [Stretch goal: Trust pinning](2_container_registry/8_trust_pinning.md)
  * [Cluster-level logging by using IBM Log Analysis with LogDNA](3_logging/1_logging.md)
    * [Provision an instance of IBM Log Analysis with LogDNA](3_logging/provision_an_instance_of_ibm_log_analysis_with_log/README.md)
      * [Step 1: Grant permissions to your user ID](3_logging/provision_an_instance_of_ibm_log_analysis_with_log/step_1_grant_permissions_in_the_ibm_cloud.md)
      * [Step 2: Provisioning an instance from the Observability dashboard](3_logging/provision_an_instance_of_ibm_log_analysis_with_log/step_2_provisioning_an_instance_from_the_observabi.md)
    * [Configure a Kubernetes cluster to forward logs to your logging instance](3_logging/configure_a_kubernetes_cluster_to_forward_logs_to_.md)
    * [Navigate to the LogDNA web UI](3_logging/navigate_to_the_logdna_web_ui.md)
    * [Monitor and manage logs through the LogDNA web UI](3_logging/monitor_and_manage_logs_through_the_logdna_web_ui/README.md)
      * [Create a view by filtering data through tags](3_logging/monitor_and_manage_logs_through_the_logdna_web_ui/create_a_view_by_filtering_data_through_tags.md)
      * [Optional - Export logs to a local file](3_logging/monitor_and_manage_logs_through_the_logdna_web_ui/export_logs_to_a_local_file.md)
      * [Optional - Configure an email alert](3_logging/monitor_and_manage_logs_through_the_logdna_web_ui/configure_an_email_alert.md)
    * [Customize the LogDNA agent](3_logging/customize_the_logdna_agent/README.md)
      * [Add custom tags to your Kubernetes cluster](3_logging/customize_the_logdna_agent/add_custom_tags_to_your_kubernetes_cluster.md)
      * [Exclude log files](3_logging/customize_the_logdna_agent/exclude_log_files.md)
    * [Control what you log by using rules](3_logging/control_what_you_log_by_using_rules/README.md)
      * [Filter out log data from storage for a worker except error entries while keeping visibility of all lines](3_logging/control_what_you_log_by_using_rules/filter_out_log_data_from_storage_for_a_worker_exce.md)
      * [Filter out kube-system logs  from storage in a cluster except error entries](3_logging/control_what_you_log_by_using_rules/filter_out_kube-system_logs_from_storage_in_a_clus.md)
      * [Filter out logs from the UI from a namespace except error entries](3_logging/control_what_you_log_by_using_rules/filter_out_logs_from_the_ui_from_a_namespace_excep.md)
  * [Monitoring a Kubernetes cluster in IBM Cloud by using IBM Cloud Monitoring with Sysdig](4_monitoring/1_monitoring.md)
    * [Provision an instance of IBM Cloud Monitoring with Sysdig](4_monitoring/provision_an_instance_of_ibm_cloud_monitoring_with/README.md)
      * [Step 1: Grant your user permissions to provision an instance in the IBM Cloud](4_monitoring/provision_an_instance_of_ibm_cloud_monitoring_with/step_1_grant_your_user_permissions_to_provision_an.md)
      * [Step 2: Provisioning an instance from the Observability dashboard](4_monitoring/provision_an_instance_of_ibm_cloud_monitoring_with/step_2_provisioning_an_instance_from_the_observabi.md)
    * [Configure a Sysdig agent to monitor a Kubernetes cluster](4_monitoring/configure_a_sysdig_agent_to_monitor_a_kubernetes_c.md)
    * [Navigate to the Sysdig web UI](4_monitoring/navigate_to_the_sysdig_web_ui.md)
    * [Troubleshoot a containerized application](4_monitoring/troubleshoot_a_containerized_application/README.md)
      * [Scenario](4_monitoring/troubleshoot_a_containerized_application/scenario.md)
      * [Download the sample application](4_monitoring/troubleshoot_a_containerized_application/download_the_sample_application.md)
      * [Install the sample ‘ticket-generator’ application](4_monitoring/troubleshoot_a_containerized_application/install_the_sample_ticket-generator_application.md)
      * [Explore the application through the Explore tab of the Sysdig web UI](4_monitoring/troubleshoot_a_containerized_application/explore_the_application_through_the_explore_tab_of.md)
      * [Explore the normal traffic flow of the application](4_monitoring/troubleshoot_a_containerized_application/explore_the_normal_traffic_flow_of_the_application.md)
      * [Explore which processes are running in each container](4_monitoring/troubleshoot_a_containerized_application/explore_which_processes_are_running_in_each_contai.md)
      * [Identify the communication patterns between the different pods](4_monitoring/troubleshoot_a_containerized_application/identify_the_communication_patterns_between_the_di.md)
      * [[Optional] Explore other pre-defined dashboards](4_monitoring/troubleshoot_a_containerized_application/[optional]_explore_other_pre-defined_dashboards.md)
    * [Monitor the application](4_monitoring/monitor_the_application/README.md)
      * [Check for HTTP response codes](4_monitoring/monitor_the_application/check_for_http_response_codes.md)
      * [Monitor the throughput of the application](4_monitoring/monitor_the_application/monitor_the_throughput_of_the_application.md)
      * [Detect error conditions](4_monitoring/monitor_the_application/detect_error_conditions.md)
      * [Monitor the latency of the application](4_monitoring/monitor_the_application/monitor_the_latency_of_the_application.md)
      * [Utilization, Saturations and Errors - USE-style overview of your namespace](4_monitoring/monitor_the_application/utilization,_saturations_and_errors_-_use-style_ov.md)
      * [Monitor data flow through your application components](4_monitoring/monitor_the_application/monitor_data_flow_through_your_application_compone.md)
      * [Monitor the Kubernetes namespace where the application is running](4_monitoring/monitor_the_application/monitor_the_kubernetes_namespace_where_the_applica.md)
    * [Alert to notify on error conditions](4_monitoring/alert_to_notify_on_error_conditions/README.md)
      * [Step 1: Define the alert](4_monitoring/alert_to_notify_on_error_conditions/step_1_define_the_alert.md)
      * [Step 2: Define the notification channel](4_monitoring/alert_to_notify_on_error_conditions/step_2_define_the_notification_channel.md)
      * [Step 3: Add the notification channel to the alert](4_monitoring/alert_to_notify_on_error_conditions/step_3_add_the_notification_channel_to_the_alert.md)

## Feedback

* [Feedback](feedback.md)

## Information

* [Further reading](information.md)

## Legal

* [Disclaimer](disclaimer.md)
