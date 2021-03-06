# Using IBM Cloud Container Registry with your Kubernetes cluster

In this section, you will configure your Kubernetes cluster to access your images in IBM Cloud Container Registry.

When you create an IBM Cloud Kubernetes Service cluster in your IBM Cloud account, the cluster is configured automatically to access all images in your account. If you're using your own cluster, or would like to control which images your cluster can access, you can configure the cluster yourself. It's useful to know what the automation is doing for you, so we'll work through configuring your cluster to use the Service ID that you created in the previous section to access your images.

## [optional] Exploring default Image Pull Secrets

Kubernetes stores authentication information for container registries in resources called Image Pull Secrets. Each Image Pull Secret contains a single username and password for a single registry.

Let's look at the Image Pull Secrets that have been added to your cluster by default.

1. List the secrets in your cluster.

    ```bash
    kubectl get secrets
    ```

2. Get the secret called `bluemix-default-secret` as YAML.

    ```bash
    kubectl get secret bluemix-default-secret -o yaml
    ```

    The credentials are stored in a section called `data`, encoded as base64. Look for a long string of letters and numbers that starts with `eyJ`.

3. Inspect the Image Pull Secret contents. Copy the secret data and paste it into the base64 decode command.

    ```bash
    echo -n <paste here> | base64 --decode
    ```

    The data is a JSON object containing a username and password for `registry.ng.bluemix.net`. The username is `token`.

When you create a pod from an image in the same IBM Cloud account as your Kubernetes cluster, you don't need to specify these default secrets. When you start a pod, Kubernetes automatically adds these Image Pull Secrets to your pod definition, because the secrets are listed in the default Service Account.

1. List Service Accounts in the default Kubernetes namespace.

    ```bash
    kubectl get serviceaccounts
    ```

2. Get the default Service Account as YAML.

    ```bash
    kubectl get sa default -o yaml
    ```

    The default Image Pull Secrets are listed in the `imagePullSecrets` section.

## Creating an Image Pull Secret

The token in the default secret has access to pull images that are owned by the cluster's account. Since your images are stored in your account, you must create your own Image Pull Secret to access your images.

1. Create an Image Pull Secret called `think-registry-demo` for your Service ID. Replace `<apikey>` with your Service ID API key.

    ```bash
    kubectl create secret docker-registry think-registry-demo --docker-server registry.ng.bluemix.net --docker-username iamapikey --docker-password <apikey> --docker-email a@b.com
    ```

2. Add your new Image Pull Secret to the default Service Account so that you don't have to manually add the secret to each pod that you run.

    ```bash
    kubectl patch sa default --type="json" --patch '[{"op":"add", "path":"/imagePullSecrets/-", "value": {"name":"think-registry-demo"}}]'
    ```

## Starting a pod

1. Create a YAML file. Call it `~/mypod.yaml`. Add the following YAML.

    ```yaml
    apiVersion: v1
    kind: Pod
    metadata:
        name: mypod
        labels:
            app: mypod
    spec:
        containers:
          - name: mypod
            image: registry.ng.bluemix.net/my_namespace/hello-world:3.6
            ports:
            - containerPort: 8080
    ```

    **Hint**: Don't forget to replace `my_namespace` with your Registry namespace.

2. Apply the YAML file.

    ```bash
    kubectl apply -f ~/mypod.yaml
    ```

3. List pods in your cluster.

    ```bash
    kubectl get pods
    ```

    A pod called `mypod` appears in the list.

4. Describe the pod.

    ```bash
    kubectl describe pod mypod
    ```

    If your Image Pull Secret has worked correctly, `Successfully pulled image` and `Started container` are shown in the events section.

5. Get the pod specification.

    ```bash
    kubectl get pod mypod -o yaml
    ```

    Note that your Image Pull Secret has been injected from the default Service Account.

6. Delete the pod.

    ```bash
    kubectl delete pod mypod
    ```

## Recap

You have created an Image Pull Secret for your IAM Service ID and used it to create a pod from an image in your IBM Cloud Container Registry namespace.

You can have multiple Image Pull Secrets in your cluster for the same registry. This allows you to finely control what images a given Kubernetes resource can access.

## Further reading

[Automating access to IBM Cloud Container Registry (IBM Cloud Docs)](https://console.bluemix.net/docs/services/Registry/registry_tokens.html#registry_access)

[Building containers from images (IBM Cloud Docs)](https://console.bluemix.net/docs/containers/cs_images.html#images)

[Specifying ImagePullSecrets on a Pod (Kubernetes Docs)](https://kubernetes.io/docs/concepts/containers/images/#specifying-imagepullsecrets-on-a-pod)

## Next

In the next section, you will use Vulnerability Advisor to detect vulnerabilities in your private images.
