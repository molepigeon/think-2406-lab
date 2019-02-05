## Install the sample ‘ticket-generator’ application {#install-the-sample-ticket-generator-application}

The ticket-generator directory contains simple scripts to create and destroy the application. The application is auto-contained in its own Kubernetes namespace.

Complete the following steps to deploy the sample application in your cluster:

1.  From a terminal, log in to IBM Cloud.

ibmcloud login -a api.ng.bluemix.net

**Note**: If you have a federated ID, use ibmcloud login --sso to log in to the IBM Cloud.

1.  Set the region where the cluster is available.

ibmcloud ks region-set us-south

1.  Set up the cluster environment. Run the following commands:
    1.  Get the command to set the environment variable and download the Kubernetes configuration files.

ibmcloud ks cluster-config CLUSTERNAME

When the download of the configuration files is finished, a command is displayed.

*   1.  Run that command in your terminal to set the path to the local Kubernetes configuration file as an environment variable. Another command is displayed.
    2.  Copy and paste that command in your terminal to set the KUBECONFIG environment variable.

**Note:** Every time you log in to the IBM Cloud Kubernetes Service CLI to work with clusters, you must run these commands to set the path to the cluster&#039;s configuration file as a session variable. The Kubernetes CLI uses this variable to find a local configuration file and certificates that are necessary to connect with the cluster in IBM Cloud.

1.  Run the following command to deploy the application in your cluster:

./create.sh

The outcome of the running this command is the following:

$ ./create.sh

namespace &quot;ticket-generator&quot; created

deployment &quot;ticket-server&quot; created

service &quot;ticket-service&quot; created

deployment &quot;ticket-balancer&quot; created

service &quot;tickets&quot; created

deployment &quot;ticket-client&quot; created

When the deployment is completed, the following resources are created:

*   The ticket-generator namespace
*   A ticket-server backend deployment composed of 2 ticket-server pods
*   A Kubernetes service to communicate with the backend service abstraction
*   A ticket-balancer deployment composed of 1 HAproxy pod
*   A Kubernetes service to communicate with the load balancer service abstraction
*   A ticket-client deployment composed of 1 ticket-consumer pod

1.  Verify the application is deployed and running in your cluster.
    1.  Run the following command to verify that the namespace **ticket-generator** has been created:

kubectl get namespaces

*   1.  Check the status of the pods.

Wait until all the pods reach Running state.

Run the following command:

kubectl get pods -n ticket-generator