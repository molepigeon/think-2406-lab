# Container Image Security Enforcement

With Container Image Security Enforcement, you can verify your container images before you deploy them to your cluster in IBM Cloud Kubernetes Service. You can control where images are deployed from, enforce Vulnerability Advisor policies, and ensure that content trust is properly applied to the image. If an image does not meet your policy requirements, the pod is not deployed to your cluster or updated.

Container Image Security Enforcement retrieves information about image content trust and vulnerabilities from IBM Cloud Container Registry. You can choose to block or to allow deployment of images that are stored in other registries, but you cannot use vulnerability or trust enforcement for these images.

In this section, you will install Container Image Security Enforcement into your cluster, and use it to control deployment of your image.

## Installing Container Image Security Enforcement

1. Install Helm into the cluster.

    `helm init`

    This command installs a component called Tiller into your cluster. For more information about Helm, see the Further Reading section below.

2. Wait for Tiller to come up. When you can list charts in your cluster without getting an error, Tiller is ready for use.

    `helm list`

3. Add the IBM Helm repository into your list of repositories.

    `helm repo add ibm https://registry.bluemix.net/helm/ibm`

4. Install the Container Image Security Enforcement chart.

    `helm install --name cise ibm/ibmcloud-image-enforcement`

## Exploring Image Policies

Container Image Security Enforcement uses policies to allow you to define what enforcement rules to apply to your images. You can create a ClusterImagePolicy to apply to the whole cluster, or you can create an ImagePolicy inside a Kubernetes namespace to control workloads inside that namespace. If an ImagePolicy exists in the namespace, it completely overrides the ClusterImagePolicy. If multiple ImagePolicies exist in a namespace, they are merged.

When you installed Container Image Security Enforcement, a default ClusterImagePolicy was created for you.

1. View the default ClusterImagePolicy.

    `kubectl get ClusterImagePolicy -o yaml`

    Note that `trust` and `va` are both set to `true` for `*`. Container Image Security Enforcement will only allow images to be deployed into the cluster if they have no vulnerabilities and are signed using Content Trust. You will explore Content Trust in a later section.

2. Create a YAML file for an ImagePolicy.

    `touch ~/imagepolicy.yaml`

3. Open your ImagePolicy file.

    `open ~/imagepolicy.yaml`

4. Configure your ImagePolicy to require Vulnerability Advisor for your image. Then save and close the editor.

    ```yaml
    apiVersion: securityenforcement.admission.cloud.ibm.com/v1beta1
    kind: ImagePolicy
    metadata:
        name: myimagepolicy
        namespace: default
    spec:
        repositories:
        - name: registry.ng.bluemix.net/my_namespace/hello-world
          policy:
            trust:
                enabled: false
            va:
                enabled: true
    ```
5. Apply your ImagePolicy.

    `kubectl apply -f ~/imagepolicy.yaml`

    Your ImagePolicy is added to the cluster.

## Deploying a vulnerable image

In the previous sections, you deployed a pod into your cluster, and then you identified vulnerabilities in the image. Now that you have configured Container Image Security Enforcement, deploy the pod again to have Container Image Security Enforcement block the deployment.

1. Make sure that the `mypod` is gone from your cluster.

    `kubectl delete --ignore-not-found pod mypod`

2. Try to create the `mypod` pod.

    `kubectl apply -f mypod.yaml`

    The deployment is not allowed because of the vulnerabilities in the image.

    ``admission webhook "va.hooks.securityenforcement.admission.cloud.ibm.com" denied the request: The Vulnerability Advisor image scan assessment found issues with the container image that are not exempted. Refer to your image vulnerability report for more details by using the command `ibmcloud cr va`.``

3. Edit the pod definition to use the image without vulnerabilities. Change the image tag from `3.6` to `latest`.

    `open ~/mypod.yaml`

    Change the `image:` line to use the `latest` image:

    `registry.ng.bluemix.net/my_namespace/hello-world:latest`

    Save and close the file.

4. Try to create the `mypod` pod.

    The pod creation is allowed because the image does not have any reported vulnerabilities in Vulnerability Advisor.

## Recap

You have installed Container Image Security Enforcement and configured it to block deployment of images that have vulnerabilities in Vulnerability Advisor.

## Further reading

[Enforcing container image security (Beta) (IBM Cloud Docs)](https://console.bluemix.net/docs/services/Registry/registry_security_enforce.html#security_enforce)

[Helm Documentation](https://docs.helm.sh/)

## Next

In the next section, you will use Content Trust to sign your image, and configure Container Image Security Enforcement to require that the image is signed.
