# Configure a Sysdig agent to monitor a Kubernetes cluster {#configure-a-sysdig-agent-to-monitor-a-kubernetes-cluster}

Open a terminal in your local environment and complete the following steps to configure a Sysdig agent on a Kubernetes cluster that runs in the IBM Cloud Kubernetes Service:

1.  Log in to IBM Cloud.

ibmcloud login -a api.ng.bluemix.net

**Note**: If you have a federated ID, use ibmcloud login --sso to log in to the IBM Cloud.

1.  Set the region where the cluster is available.

ibmcloud ks region-set us-south

1.  Set up the cluster environment. Run the following commands:
    1.  Get the command to set the environment variable and download the Kubernetes configuration files.

ibmcloud ks cluster-config CLUSTERNAME

When the download of the configuration files is finished, a command is displayed.

*   1.  Copy and paste that command in your terminal to set the KUBECONFIG environment variable.

**Note:** Every time you log in to the IBM Cloud Kubernetes Service CLI to work with clusters, you must run these commands to set the path to the cluster&#039;s configuration file as a session variable. The Kubernetes CLI uses this variable to find a local configuration file and certificates that are necessary to connect with the cluster in IBM Cloud.

1.  Deploy the Sysdig agent by running a script.
    1.  From the _Observability_ dashboard, select **Edit sources**.
    2.  Copy the command from the section **Install Sysdig Agent to your cluster**. Before you run it, add the following tags: workshop:lab,location:madrid The command looks as follows:

curl -sL https://raw.githubusercontent.com/draios/sysdig-cloud-scripts/master/agent_deploy/IBMCloud-Kubernetes-Service/install-agent-k8s.sh | bash -s -- -a 3f049b1f-402d-409b-82fd-1442dd8efe4f -c ingest.us-south.monitoring.cloud.ibm.com -t workshop:lab,location:madrid -ac &#039;sysdig_capture_enabled: false&#039;

Notice that the command already includes the correct values for the access key and the ingestion URL.

Information about the command:

curl -sL https://raw.githubusercontent.com/draios/sysdig-cloud-scripts/master/agent_deploy/IBMCloud-Kubernetes-Service/install-agent-k8s.sh | bash -s -- -a SYSDIG_ACCESS_KEY -c COLLECTOR_ENDPOINT -t TAG_DATA -ac &#039;sysdig_capture_enabled: false&#039;

where:

*   SYSDIG_ACCESS_KEY is the ingestion key for the monitoring instance.
*   COLLECTOR_ENDPOINT is the ingestion URL ( ingest.us-south.monitoring.cloud.ibm.com).
*   TAG_DATA are comma-separated tags that are formatted as _TAG_NAME:TAG_VALUE_. You can associate one or more tags to your Sysdig agent.

To disable the Sysdig capture feature, set **sysdig_capture_enabled** to _false_. By default is set to _true_.

1.  Verify that the Sysdig agent is running. Cluster resources are located in the **ibm-observe** namespace.
    1.  Check that the pods are running. Your cluster has two workers running, therefore, you should see two sysdig-agent pods.
    2.  Run the following command:

kubectl --namespace=ibm-observe get pods

*   1.  To check that the secret containing the access key has been created, you can run the following command:

kubectl --namespace=ibm-observe get secrets