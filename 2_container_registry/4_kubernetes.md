# Using IBM Cloud Container Registry with your Kubernetes cluster

In this section, you will configure your Kubernetes cluster to access your images in IBM Cloud Container Registry.

When you create an IBM Cloud Kubernetes Service cluster in your IBM Cloud account, the cluster is configured automatically to access all images in your account. If you're using your own cluster, or would like to control which images your cluster can access, you can configure the cluster yourself. It's useful to know what the automation is doing for you, so we'll work through configuring your cluster to use the Service ID that you created in the previous section to access your images.

## Configuring kubectl to control your Kubernetes cluster

`kubectl` is the command-line tool for administering Kubernetes clusters. In this lab, you will be using `kubectl` to control the cluster that you were assigned in the _Getting Started_ section. If you haven't had a cluster assigned to you, go back and complete that section now.

You must download a configuration file to allow `kubectl` to control your cluster.

1. Log in to IBM Cloud. When prompted to select an account, choose the one called **IBM**.

    `ibmcloud login`

2. Make sure that you are targeting the IBM Cloud Kubernetes Service region that you selected in the Get Cluster tool.

    `ibmcloud ks region-set`

3. List your clusters.

    `ibmcloud ks clusters`

4. Download the configuration for your cluster.

    `$(ibmcloud ks cluster-config <cluster-name> --export)`

5. List pods in your cluster to confirm that kubectl is configured correctly.

    `kubectl get pods`

## Exploring default Image Pull Secrets

Kubernetes stores authentication information for container registries in resources called Image Pull Secrets. Each Image Pull Secret contains a single username and password for a single registry.

Let's look at the Image Pull Secrets that have been added to your cluster by default.

1. List the secrets in your cluster.

    `kubectl get secrets`

2. Get the secret called `bluemix-default-secret` as YAML.

    `kubectl get secret bluemix-default-secret -o yaml`

    The credentials are stored in a section called `data`, encoded as base64. Look for a long string of letters and numbers that starts with `eyJ`.

3. Inspect the Image Pull Secret contents. Copy the secret data and paste it into the base64 decode command.

    `echo -n <paste here> | base64 --decode | jq`

    The data is a JSON object containing a username and password for `registry.ng.bluemix.net`. The username is `token`.

## Creating an Image Pull Secret

The token in the default secret has access to pull images that are owned by the cluster's account. Since your images are stored in your account, you must create your own Image Pull Secret to access your images.

1. Create an Image Pull Secret for your Service ID. Replace `<apikey>` with your Service ID API key.

    `kubectl create secret docker-registry --docker-server registry.ng.bluemix.net --docker-username iamapikey --docker-password <apikey> --docker-email a@b.com`

Add IPS to default SA?

Create pod from image in our namespace

If added to default SA, show the IPS being injected automatically

## Recap

You can have multiple Image Pull Secrets in your cluster for the same registry.
