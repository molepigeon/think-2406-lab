# Deploying trusted content

Signing your images allows you to cryptographically verify the content of that image. Image signatures are stored in the cloud alongside your images. You can use tools such as Container Image Security Enforcement to check image signatures at deploy time, to ensure that you trust the code that's running in your production environments. IBM Cloud Container Registry refers to this feature as Content Trust.

IBM Cloud Container Registry runs Notary as its trust server. Notary is a tool that manages signatures. It implements The Update Framework (TUF), a specification for securing application deployment. Both Notary and TUF are "Incubating" CNCF projects.

When Content Trust is enabled, Docker pulls the most recent signed image that matches the image name that you specified, which might not be the most recent image.

## Signing your image

When you enable Content Trust, Docker pushes trust information into IBM Cloud Container Registry when you push your image.

1. Set the `DOCKER_CONTENT_TRUST` variable to enable Content Trust.

    ```bash
    export DOCKER_CONTENT_TRUST=1
    ```

    If you close your terminal window, you must repeat this command to keep using Content Trust.

2. Log in to IBM Cloud Container Registry.

    ```bash
    ibmcloud cr login
    ```

    The command returns an export line to set the `DOCKER_CONTENT_TRUST_SERVER` variable.

    ```bash
    export DOCKER_CONTENT_TRUST_SERVER=https://registry.ng.bluemix.net:4443
    ```

3. Run the export line that was returned from the `ibmcloud cr login` command.

4. Give your image a new name. Normally you can sign your image in place, but you will be using the unsigned image later.

    ```bash
    docker tag registry.ng.bluemix.net/my_namespace/hello-world:latest registry.ng.bluemix.net/my_namespace/signed:latest
    ```

    **Hint**: Don't forget to replace `my_namespace` with your Registry namespace in both image names.

5. Push your image. Once the push is complete, Docker will prompt you to create a number of key passphrases. Make sure to remember your passphrases, because you will need them to add more trust data later.

    ```bash
    docker push registry.ng.bluemix.net/my_namespace/signed:latest
    ```

6. Check your image signature.

    ```bash
    docker trust inspect --pretty registry.ng.bluemix.net/my_namespace/signed:latest
    ```

## Enforcing content trust by using Container Image Security Enforcement

Container Image Security Enforcement can be configured to implement Content Trust in the same way as Docker. When you create a new workload in Kubernetes, Container Image Security Enforcement modifies the image tag in your workload specification to refer to the most recent signed version of the image before the Kubernetes API Server distributes the workload to your cluster workers. Your cluster workers pull the signed version of the image instead of the most recently tagged version, if they are different. When pods are re-scheduled by your cluster, the new workers also pull the same signed version, rather than the tagged version that may have changed in between creating the workload and re-scheduling the pod.

1. Modify your ImagePolicy to enable Content Trust enforcement for images from our namespace.

    ```bash
    vi ~/imagepolicy.yaml
    ```

    ```yaml
    apiVersion: securityenforcement.admission.cloud.ibm.com/v1beta1
    kind: ImagePolicy
    metadata:
        name: myimagepolicy
        namespace: default
    spec:
        repositories:
        - name: registry.ng.bluemix.net/my_namespace/*
          policy:
            trust:
                enabled: true
            va:
                enabled: false
    ```

    Note that the repository name has changed to a `*` wildcard character. All images in your namespace have the policy applied.

2. Apply the new ImagePolicy.

    ```bash
    kubectl apply -f ~/imagepolicy.yaml
    ```

3. Delete your `mypod` pod.

    ```bash
    kubectl delete --ignore-not-found pod mypod
    ```

4. Try to create the `mypod` pod.

    ```bash
    kubectl apply -f ~/mypod.yaml
    ```

    The deployment is not allowed because the image is not signed.

    `admission webhook "trust.hooks.securityenforcement.admission.cloud.ibm.com" denied the request: Deny, failed to get content trust information: registry.ng.bluemix.net:4443 does not have trust data for latest`

5. Edit the pod definition to use the signed image. Change the image repository from `my_namespace/hello-world` to `my_namespace/signed`.

    ```bash
    vi ~/mypod.yaml
    ```

    Change the `image:` line to use the `signed` image:

    ```bash
    registry.ng.bluemix.net/my_namespace/signed:latest
    ```

    Save and close the file.

6. Try to create the `mypod` pod.

    ```bash
    kubectl apply -f ~/mypod.yaml
    ```

    The pod creation is allowed because the image is signed.

7. Look at the specification for your pod.

    ```bash
    kubectl get pod mypod -o yaml
    ```

    Note that the image has changed to a canonical format image name:

    `registry.ng.bluemix.net/my_namespace/signed@sha256:xxxxx`

    Container Image Security Enforcement has modified your pod description to set the image to the digest of the latest signed version. A digest is an immutable reference to the image content. By mutating the pod spec in this way, the cluster will always pull the version of the image that was verified by Container Image Security Enforcement, even when you delete pods created by another workload, like a deployment, later.

8. Delete your `mypod` pod.

    ```bash
    kubectl delete --ignore-not-found pod mypod
    ```

9. Turn off Content Trust.

    ```bash
    unset DOCKER_CONTENT_TRUST
    ```

10. Push the vulnerable image over the top of the signed image. Note that because you turned off Content Trust, no trust data is pushed this time.

    ```bash
    docker tag registry.ng.bluemix.net/my_namespace/hello-world:3.6 registry.ng.bluemix.net/my_namespace/signed:latest
    ```

    ```bash
    docker push registry.ng.bluemix.net/my_namespace/signed:latest
    ```

11. Try to create the `mypod` pod.

    ```bash
    kubectl apply -f ~/mypod.yaml
    ```

    The pod creation is still allowed because a signed version of the image exists, even though it isn't the latest according to the Container Registry.

12. Look at the specification for your pod again.

    ```bash
    kubectl get pod mypod -o yaml
    ```

    Note that the image was changed to the same digest as before, because it is still the most recently signed version of the image.

## Recap

You have signed your image and configured Container Image Security Enforcement to require image signatures. You have seen that when it enforces content trust, Container Image Security Enforcement modifies the image name to the digest of the signed image.

## Further reading

[Signing images for trusted content (IBM Cloud Docs)](https://console.bluemix.net/docs/services/Registry/registry_trusted_content.html#registry_trustedcontent)

[Enforcing container image security (Beta) (IBM Cloud Docs)](https://console.bluemix.net/docs/services/Registry/registry_security_enforce.html#security_enforce)

[The Update Framework](https://theupdateframework.github.io/)

[Notary Project](https://github.com/theupdateframework/notary)

## Next

In the next section of this part of the tutorial, you will configure Container Image Security Enforcement to verify the signature on an image.
