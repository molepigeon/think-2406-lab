# Trusted delegation roles and trust pinning

Content Trust allows you to sign images, but it gives control over image signatures to the first person who pushes signed data into the trust server. When using trusted content in the way that you did in the previous section, you trusted that the IBM Cloud Container Registry trust server was returning valid signed digests.

A proper security model would hold true even when the trust server were compromised. Container Image Security Enforcement should be able to detect when the signature is invalid, and prevent deployment of the invalid image.

With Content Trust, you can add multiple people's keys as signers in each image repository, and you can configure Container Image Security Enforcement to require a signature from a given signer. When it is configured to require a signature from a particular signer, Container Image Security Enforcement also verifies their signature by using that signer's public key.

You can also use multiple signers to indicate someone's approval of an image in a deployment process. In this case, Container Image Security Enforcement could verify this signature to ensure that they have signed off on an image before it is deployed to production.

**Note:** This section will add another 15 minutes to the duration of this part of the tutorial. If you would like enough time to explore the Logging and Monitoring parts in detail during the session, you should consider returning to this section once you have completed those parts.

## Creating a signing key

1. Create a signing key called `thinkdemo`.

    `docker trust key generate thinkdemo`

    The private key goes in to your Docker Content Trust directory. The public key is saved to `thinkdemo.pub` in your working directory.

## Adding your signing key to your image

1. Set the `DOCKER_CONTENT_TRUST` variable to enable Content Trust.

    `export DOCKER_CONTENT_TRUST=1`

    If you close your terminal window, you must repeat this command to keep using Content Trust.

2. Log in to IBM Cloud Container Registry.

    `ibmcloud cr login`

    The command returns an export line to set the `DOCKER_CONTENT_TRUST_SERVER` variable.

    `export DOCKER_CONTENT_TRUST_SERVER=https://registry.ng.bluemix.net:4443`

3. Run the export line that was returned from the `ibmcloud cr login` command.

4. Add your key as a signer.

    `docker trust signer add --key=thinkdemo.pub thinkdemo registry.ng.bluemix.net/my_namespace/signed`

## Enforcing specific signers in Kubernetes

To configure Container Image Security Enforcement to require a signature from a specific signer, you must add the signer's public key to the cluster as a secret, and then reference the secret in your ImagePolicy.

1. Create a secret for the signer's public key.

    `kubectl create secret generic thinkdemo --from-literal=name=thinkdemo --from-file=publicKey=thinkdemo.pub`

2. Add the signer secret as a required key in your ImagePolicy.

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
                signerSecrets:
                - name: thinkdemo
            va:
                enabled: false
    ```

3. Delete your `mypod` pod.

    `kubectl delete --ignore-not-found pod mypod`

4. Try to create the `mypod` pod.

    `kubectl apply -f mypod.yaml`

    The deployment is not allowed because while the image is signed, it is not signed by the `thinkdemo` key.

5. Sign the image.

    `docker trust sign registry.ng.bluemix.net/my_namespace/signed:latest`

    Because you have the private key for `thinkdemo` on your workstation, Docker adds a signature using that key.

6. Check your image signature.

    `registry.ng.bluemix.net/my_namespace/signed:latest`

    Note that a signature from `thinkdemo` is shown as well as a signature from `Repo Admin`.

7. Try to create the `mypod` pod.

    `kubectl apply -f mypod.yaml`

    The deployment is allowed because the image is signed by the required `thinkdemo` key.

## Cleaning up

1. Uninstall Container Image Security Enforcement.

    `helm delete --purge cise`

2. Delete your `mypod` pod.

    `kubectl delete --ignore-not-found pod mypod`

## Recap

You have created a signing key and configured Container Image Security Enforcement to require images in your namespace to be signed using it. You can use the same signing key in multiple repositories to require a common root of trust.

That concludes the IBM Cloud Container Registry part of this tutorial.

## Further Reading

**TODO**
