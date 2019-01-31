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

When you create a pod from an image in the same IBM Cloud account as your Kubernetes cluster, you don't need to specify these default secrets. When you start a pod, Kubernetes automatically adds these Image Pull Secrets to your pod definition, because the secrets are listed in the default Service Account.

1. List Service Accounts in the default Kubernetes namespace.

    `kubectl get serviceaccounts`

2. Get the default Service Account as YAML.

    `kubectl get sa default -o yaml`

    The default Image Pull Secrets are listed in the `imagePullSecrets` section.

## Creating an Image Pull Secret

The token in the default secret has access to pull images that are owned by the cluster's account. Since your images are stored in your account, you must create your own Image Pull Secret to access your images.

1. Create an Image Pull Secret called `think-registry-demo` for your Service ID. Replace `<apikey>` with your Service ID API key.

    Note that the `docker-server` value includes your Registry namespace. Because the secret is configured in this way, Kubernetes will only use this Image Pull Secret for images in your namespace, and it will use the default secret for other images.

    `kubectl create secret docker-registry think-registry-demo --docker-server registry.ng.bluemix.net/my_namespace --docker-username iamapikey --docker-password <apikey> --docker-email a@b.com`

2. As with the default Image Pull Secrets, you can add your new Image Pull Secret to the default Service Account so that you don't have to manually add the secret to each pod that you run. You can do this interactively or by using `kubectl patch`.

    1. Edit the default Service Account interactively, and add your new secret to the list:

        1. Open the default Service Account for editing. The default editor for `kubectl edit` is `vim`.

            `kubectl edit sa default`

        2. Add the following YAML to the bottom of the existing `imagePullSecrets` list. In vim, press `I` to enter insert mode.

            ```yaml
            - name: think-registry-demo
            ```

        3. Save and close the file.

            `Esc, :wq, Enter`

    2. Patch the default Service Account. `kubectl patch` allows you to specify a patch, in this case a JSON patch, to apply to the resource. You can use this to automate rollouts.

        `kubectl patch sa default --type="json" --patch '[{"op":"add", "path":"/imagePullSecrets/-", "value": {"name":"think-registry-demo"}}]'`

## Starting a pod

1. Create a YAML file. Call it `mypod.yaml`. Add the following YAML.

    ```yaml
    apiVersion: v1
    kind: Pod
    metadata:
    name: nginx
    labels:
        name: nginx
    spec:
        containers:
        - name: nginx
            image: registry.ng.bluemix.net/my_namespace/
            ports:
            - containerPort: 80
    ```

    **Hint**: Don't forget to replace `my_namespace` with your Registry namespace.

2. Apply the yaml.

    `kubectl apply -f mypod.yaml`

3. List pods in your cluster.

    `kubectl get pods`

    A pod called `nginx` appears in the list.

4. Describe the pod.

    `kubectl describe pod nginx`

    **TODO** There's got to be some information that's worth looking at in here!

5. Get the pod specification.

    `kubectl get pod nginx -o yaml`

    Note that your Image Pull Secret has been injected from the default Service Account.

## Recap

You have created an Image Pull Secret for your IAM Service ID and used it to create a pod from an image in your IBM Cloud Container Registry namespace.

You can have multiple Image Pull Secrets in your cluster for the same registry. This allows you to finely control what images a given Kubernetes resource can access.

## Further Reading

**TODO**

## Next

In the next section, you will use Vulnerability Advisor to detect vulnerabilities in your private images.
