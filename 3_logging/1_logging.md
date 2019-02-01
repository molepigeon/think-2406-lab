# Part 2: Cluster-level logging by using IBM Log Analysis with LogDNA

## Logging in IBM Cloud by using IBM Log Analysis with LogDNA

IBM Log Analysis with LogDNA is a co-branded service that you can include as part of your IBM Cloud architecture to add log management capabilities. IBM Log Analysis with LogDNA is operated by LogDNA in partnership with IBM.

You can use IBM Log Analysis with LogDNA to manage system and application logs in IBM Cloud.

IBM Log Analysis with LogDNA offers administrators, DevOps teams, and developers advanced features to filter, search, and tail log data, define alerts, and design custom views to monitor application and system logs.

The following figure shows the components overview for the IBM Log Analysis with LogDNA service that is running on IBM Cloud:

![Components overview of IBM Log Analysis with LogDNA](logging_ov.png)

## Features

**Troubleshoot logs in real-time to diagnose issues and identify problems.**

By using the live streaming tail feature, developers and DevOps teams can diagnose issues, analyze stack traces and exceptions, identify the source of errors, and monitor different log sources through a single view. This feature is available through the command line and through the web interface.

**Issue alerts so that you are notified about important actions.**

To act promptly for application and services events that you identify as critical, DevOps teams can configure alert notification integrations with the following systems: email, Slack, HipChat, webhook, PagerDuty, and Opsgenie.

**Export logs to a local file for analysis or to an archive service to meet auditing requirements.**

Export specific log lines to a local copy or archive logs from IBM Log Analysis with LogDNA to IBM Cloud Object Storage. Log lines are exported in JSON line format. Logs are archived in JSON format and preserve the metadata that is associated with each line.

**Control logging infrastructure costs by customizing the logs that you want to manage through IBM Log Analysis with LogDNA.**

Control the cost of your logging infrastructure in IBM Cloud by configuring the log sources for which you want to collect and manage logs.

## Audience

This section of the lab is suitable for any users with no previous knowledge of the service that want to learn how to use the IBM Log Analysis with LogDNA service.

## Objectives

The different sections in this part of the lab will show you how to:

* Provision the IBM Log Analysis with LogDNA service
* Configure a Kubernetes cluster so that you can monitor its logs by using the IBM Log Analysis with LogDNA service
* Launch the LogDNA web UI from where you can monitor and manage your logs
* Explore the LogDNA web to learn about the different sections

The following sections are optional and require a paid service plan. This lab include sections to show you how to:

* Create an alert
* Export log data

## Duration

The estimated time to complete this part of the lab is 20 minutes.
