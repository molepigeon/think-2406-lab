# Getting Started

Before you start, you need the following resources and information:

* **An IBM Cloud account**. If you do not have one, contact one of the instructors to get a promotional code that you can apply in your account to complete the lab.

* Access to an IBM Cloud Lab account which contains **pre-provisioned clusters**. Each lab attendee will be granted access to one cluster to complete this lab.

    Note: You can also use a cluster that you might have provisioned in your account.

* Login details to the lab **VM image**.

    To complete your lab by using one of the provided environments in the lab room, you must use the lab VM image. The image is preinstalled with all the CLIs that are required to complete the lab.

## [optional] Sign up for an IBM Cloud account

1. Create your own [IBM Cloud account](https://cloud.ibm.com).

2. Wait for the verification email. Follow the instructions in the email.

3. Log in to [IBM Cloud](https://cloud.ibm.com).

## Add your promotional code{#4}

1. Login to [IBM Cloud account](https://cloud.ibm.com) with your IBM ID.

2. Select **Manage > Account > Account settings**.

3. Look for the section **Feature (Promo) Codes**.

4. Enter your promotional code.

## Request a pre-provisioned cluster

Go to [Get Cluster](https://think-iks.mybluemix.net/) web page and enter your IBM ID (the email you used to sign up). Contact the lab instructors to get the lab key.

Unless instructed otherwise, choose the **US South** region.

![cluster assignment tool](https://raw.githubusercontent.com/rvennam/istio101/master/workshop/README_images/get-cluster.png)

You will be added to the Lab account and granted access to a cluster.

Refresh your [IBM Cloud Dashboard](https://cloud.ibm.com)

Switch to the **IBM** account by clicking on the account selection drop down in the top navigation bar.

Click on **View all** in the Resource Summary tile.

Under **Kubernetes Clusters**, one cluster is shown. You will use this cluster for this lab. **Note that this is a *Standard/Paid* cluster (as opposed to FREE cluster.)**

![cluster dashboard](https://raw.githubusercontent.com/rvennam/istio101/master/workshop/README_images/dashboard.png)

## Configure kubectl to control your Kubernetes cluster

`kubectl` is the command-line tool for administering Kubernetes clusters. In this lab, you will be using `kubectl` to control the cluster that you were assigned in the _Getting Started_ section. If you haven't had a cluster assigned to you, go back and complete that section now.

You must download a configuration file to allow `kubectl` to control your cluster. Switch to the `IBM CSS` tab on your lab machine to access your lab VM, and then open _Terminal Emulator_ and enter the following commands:

1. Log in to IBM Cloud. When prompted to select an account, choose the one called **IBM**.

    ```bash
    ibmcloud login
    ```

2. Make sure that you are targeting the IBM Cloud Kubernetes Service region that you selected in the Get Cluster tool.

    ```bash
    ibmcloud ks region-set us-south
    ```

3. List your clusters.

    ```bash
    ibmcloud ks clusters
    ```

4. Download the configuration for your cluster.

    ```bash
    ibmcloud ks cluster-config cluster_name --export
    ```

    Replace `cluster_name` with the name of your assigned cluster.

    The command returns an export line to set the `KUBECONFIG` variable. If the command does not return an export line, make sure that you typed the cluster name correctly.

    ```bash
    export KUBECONFIG=/home/sysadmin/.bluemix/plugins/container-service/clusters/think-iks-100/kube-config-sjc04-think-iks-100.yaml
    ```

5. Run the export line that was returned from the `ibmcloud ks cluster-config` command.

6. List namespaces in your cluster to confirm that kubectl is configured correctly.

    ```bash
    kubectl get namespaces
    ```
