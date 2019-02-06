## Install the sample ‘ticket-generator’ application {#install-the-sample-ticket-generator-application}

The ticket-generator directory contains simple scripts to create and destroy the application. The application is auto-contained in its own Kubernetes namespace.

Complete the following steps to deploy the sample application in your cluster:

### 1. Log in to IBM Cloud{#1}

Log in to IBM Cloud. Run the following command:

```
ibmcloud login -a api.ng.bluemix.net
```

**Note**: If you have a federated ID, use `ibmcloud login --sso` to log in to the IBM Cloud.

After you log in with your user ID and password, the IBM Cloud Dashboard opens.

### 2. Set up the cluster environment{#2}

1. List the clusters. Run the following command:

    ```
    ibmcloud ks clusters
    ```

2. Set the resource group where the cluster is available. Run the following command:

    ```
    ibmcloud target -g ResourceGroupName
    ```

    You can run the command `ibmcloud ks cluster-get ClusterName` to find out the resource group.

3. Set the region where the cluster is available. Although you have logged in to the us-south region, you must set the region as part of setting the command line to work with the IBM Cloud Kubernetes service.

    ```
    ibmcloud ks region-set us-south
    ```

4. Get the command to set the environment variable and download the Kubernetes configuration files. 

    ```
    ibmcloud ks cluster-config CLUSTERNAME
    ```
    When the download of the configuration file is finished, a command is displayed.  

5. Copy and paste that command in your terminal to set the KUBECONFIG environment variable.


**Note:** Every time you log in to the IBM Cloud Kubernetes Service CLI to work with clusters, you must run these commands to set the path to the cluster&#039;s configuration file as a session variable. The Kubernetes CLI uses this variable to find a local configuration file and certificates that are necessary to connect with the cluster in IBM Cloud. 


### Deploy the application{#3}

From the command line, run the following command to deploy the application in your cluster:

```
./create.sh
```

The outcome of running this command is the following:

```
$ ./create.sh
namespace "ticket-generator" created
deployment "ticket-server" created
service "ticket-service" created
deployment "ticket-balancer" created
service "tickets" created
deployment "ticket-client" created
```

When the deployment is completed, the following resources are created:

*  The ticket-generator namespace
*  A ticket-server backend deployment composed of 2 ticket-server pods
*  A Kubernetes service to communicate with the backend service abstraction
*  A ticket-balancer deployment composed of 1 HAproxy pod
*  A Kubernetes service to communicate with the load balancer service abstraction
*  A ticket-client deployment composed of 1 ticket-consumer pod

### Verify the application is deployed and running in your cluster{#4}

Run the following command to verify that the namespace **ticket-generator** has been created:

```
kubectl get namespaces
```

Check the status of the pods. Run the following command:

```
kubectl get pods -n ticket-generator
```

Wait until all the pods reach **Running** state.