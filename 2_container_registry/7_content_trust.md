# Deploying trusted content

Signing your images allows you to cryptographically verify the content of that image. Image signatures are stored in the cloud alongside your images. You can use tools such as Container Image Security Enforcement to check image signatures at deploy time, to ensure that you trust the code that's running in your production environments. IBM Cloud Container Registry refers to this feature as Content Trust.

When Content Trust is enabled, Docker pulls the most recent signed image that matches the image name that you specified, which might not be the most recent image.

## Signing your image

When you enable Content Trust, Docker pushes trust information into IBM Cloud Container Registry when you push your image.

1. Set the `DOCKER_CONTENT_TRUST` variable to enable Content Trust.

    `export DOCKER_CONTENT_TRUST=1`

    If you close your terminal window, you must repeat this command to keep using Content Trust.

2. Log in to IBM Cloud Container Registry.

    `ibmcloud cr login`

    The command returns an export line to set the `DOCKER_CONTENT_TRUST_SERVER` variable.

    `export DOCKER_CONTENT_TRUST_SERVER=https://registry.ng.bluemix.net:4443`

3. Run the export line that was returned from the `ibmcloud cr login` command.

4. Give your image a new name. Normally you can sign your image in place, but you will be using the unsigned image later.

    `docker tag registry.ng.bluemix.net/my_namespace/hello-world:latest registry.ng.bluemix.net/my_namespace/signed:latest`

    **Hint**: Don't forget to replace `my_namespace` with your Registry namespace in both image names.

5. Push your image. Once the push is complete, Docker will prompt you to create a number of key passphrases. Make sure to remember your passphrases, because you will need them to add more trust data later.

    `docker push registry.ng.bluemix.net/my_namespace/signed:latest`

6. You can check your image signature by using `registry.ng.bluemix.net/my_namespace/signed:latest`.

## Enforcing content trust by using Container Image Security Enforcement

**TODO**

CISE does trust by mutating the pod spec yadda yadda yadda.

1. Modify your ImagePolicy to enable Content Trust enforcement for images from our namespace.

    `open ~/imagepolicy.yaml`

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

2. Delete your `mypod` pod.

    `kubectl delete --ignore-not-found pod mypod`

3. Try to create the `mypod` pod.

    `kubectl apply -f mypod.yaml`

    The deployment is not allowed because the image is not signed.

    `admission webhook "trust.hooks.securityenforcement.admission.cloud.ibm.com" denied the request: Deny, failed to get content trust information: No valid trust data for latest`

4. Edit the pod definition to use the signed image. Change the image repository from `my_namespace/hello-world` to `my_namespace/signed`.

    `open ~/mypod.yaml`

    Change the `image:` line to use the `signed` image:

    `registry.ng.bluemix.net/my_namespace/signed:latest`

    Save and close the file.

5. Try to create the `mypod` pod.

    The pod creation is allowed because the image is signed.

6. Look at the specification for your pod.

    `kubectl get pod mypod -o yaml`

    Note that the image has changed to a canonical format image name:

    `registry.ng.bluemix.net/my_namespace/signed@sha256:xxxxx`

    Container Image Security Enforcement has modified your pod description to set the image to the digest of the latest signed version. A digest is an immutable reference to the image content. By mutating the pod spec in this way, the cluster will always pull the version of the image that was verified by Container Image Security Enforcement, even when you delete pods created by another workload, like a deployment, later.

7. Delete your `mypod` pod.

    `kubectl delete --ignore-not-found pod mypod`

8. Turn off Content Trust.

    `unset DOCKER_CONTENT_TRUST`

9. Push the vulnerable image over the top of the signed image. Note that because you turned off Content Trust, no trust data is pushed this time.

    `docker tag registry.ng.bluemix.net/my_namespace/hello-world:3.6 registry.ng.bluemix.net/my_namespace/signed:latest`

    `docker push registry.ng.bluemix.net/my_namespace/signed:latest`

10. Try to create the `mypod` pod.

    The pod creation is still allowed because a signed version of the image exists, even though it isn't the latest according to the Container Registry.

11. Look at the specification for your pod again.

    `kubectl get pod mypod -o yaml`

    Note that the image was changed to the same digest as before, because it is still the most recently signed version of the image.

## Recap

You have signed your image and configured Container Image Security Enforcement to require image signatures. You have seen that when it enforces content trust, Container Image Security Enforcement modifies the image name to the digest of the signed image.

## Further reading

**TODO**

make sure to include content trust docs here.

## Next

In the next section of this part of the tutorial, you will configure Container Image Security Enforcement to verify the signature on an image.
