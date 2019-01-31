# Part 3: Monitoring a Kubernetes cluster in IBM Cloud by using IBM Cloud Monitoring with Sysdig

## IBM Cloud Monitoring with Sysdig

IBM Cloud Monitoring with Sysdig is a co-branded cloud-native, and container- intelligence management system that you can include as part of your IBM Cloud architecture. Use it to gain operational visibility into the performance and health of your applications, services, and platforms. It offers administrators, DevOps teams, and developers full stack telemetry with advanced features to monitor and troubleshoot performance issues, define alerts, and design custom dashboards. IBM Cloud Monitoring with Sysdig is operated by Sysdig in partnership with IBM.

The following diagram shows the components overview for the IBM Cloud Monitoring with Sysdig service that is running on IBM Cloud:

[monitoring_ov.png](Components overview of IBM Cloud Monitoring with Sysdig)

## Features

**Accelerate the diagnosis and resolution of performance incidents.**

IBM Cloud Monitoring with Sysdig offers deep visibility into your infrastructure and applications with the ability to troubleshoot from service level all the way down to the system level. Pre-defined dashboards and alerts simplify identification of potential threats or problems. By using IBM Cloud Monitoring with Sysdig, developers and DevOps teams monitor and troubleshoot performance issues in real-time, identify the source of errors, and eliminate problems.

**Control the cost of your monitoring infrastructure.**

IBM Cloud Monitoring with Sysdig includes functionality that helps you to control the cost of your monitoring infrastructure in IBM Cloud. You can configure the metric sources for which you want to monitor performance. You can enable a pre-defined alert to warn you of usage changes that will impact your billing.

**Explore and visualize easily your entire environment.**

IBM Cloud Monitoring with Sysdig makes it easier to visually explore your environment. Dynamic topology maps provide a view of dependencies between services. Multi- dimensional queries across high churn, high cardinality, high frequency metrics accelerate troubleshooting. Customizable dashboards allow you to visualize what matters most.

**Get critical Kubernetes and container insights for dynamic microservice monitoring.**

IBM Cloud Monitoring with Sysdig auto-discovers Kubernetes environments and provides out-of-the-box dashboards and alerts for clusters, nodes, namespaces, services, deployments, pods, and more. A single agent per node dynamically discovers all microservices and auto-collects metrics and events from various sources including Kubernetes, hosts, networks, containers, processes, applications, and custom metrics like Prometheus, JMX, and StatsD.

**Mitigate the impact of abnormal situations with proactive notifications.**

IBM Cloud Monitoring with Sysdig includes alerts and multi-channel notifications that you can use to reduce the impact on your day-to-day operations and accelerate your reaction and response time to anomalies, downtime, and performance degradation. Notification channels that you can easily configure
include email, Slack, PagerDuty, webhooks, Opsgenie, and VictorOps.

## Audience

This section of the lab is suitable for any users that want to deep dive into how to use the IBM Cloud Monitoring with Sysdig service and learn how to use this service to monitor the health of deployed resources in a Kubernetes cluster.
Knowledge

You are expected to understand the basic concepts of microservice applications, containers, and Kubernetes.

## Objectives

The different sections in this part of the lab will show you how to:

* Provision the IBM Cloud Monitoring with Sysdig service
* Configure a Kubernetes cluster so that you can monitor its health and performance by using the IBM Cloud Monitoring with Sysdig service
* Launch the Sysdig web UI from where you can monitor and manage your environments
* Explore a microservices application by using IBM Cloud Monitoring with Sysdig
* Monitor an application that is running in a cluster
* Define an alert to notify you about errors

## Duration

The estimated time to complete this part of the lab is 45 minutes.
