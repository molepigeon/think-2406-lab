# Start using IBM Cloud Container Registry

In this section, you will create a namespace in IBM Cloud Container Registry for your account, and push an image to it.

## Using the IBM Cloud Container Registry CLI

You can use the `ibmcloud cr` command line tool to interact with IBM Cloud Container Registry. The Container Registry plugin is installed on your lab machine, but you can install it on your machine once `ibmcloud` is installed by typing `ibmcloud plugin install container-registry`.

1. Run `ibmcloud cr`. A list of valid commands is displayed.
2. Run `ibmcloud cr images --help`. Valid parameters for the image list command is displayed.

You can continue to explore options if you want. Carry on to the next section when you are done.

## Creating a namespace

In IBM Cloud Container Registry, you store your container images in a location known as a namespace. Your namespace is used as part of your image name, and is unique across a region. For this section, you will need to think of a name for your namespace. Namespaces can have lowercase letters, numbers, dashes or full stops, and must be at least 4 and no longer than 30 characters long.

From now on, this guide will use `my_namespace` as the namespace. When you see `my_namespace` in commands, replace it with your chosen namespace name before running the command.

1. Log in to your IBM Cloud account.

    `ibmcloud login`

2. Make sure that you are targeting the US South instance of IBM Cloud Container Registry.

    `ibmcloud cr region-set us-south`

3. Create your namespace.

    `ibmcloud cr namespace-add my_namespace`

    If your chosen namespace name is available, it is assigned to your account. If you see a message saying that your chosen namespace is already in use, try again with a different namespace name.

## Pushing an image

A container image name is made up of two parts, a repository and a tag, in the format `repository:tag`. The repository is a URI of where on the Internet the image can be found, and the tag is a version for the image. For example, in `icr.io/ibm/datashield-server:0.2.38`, the repository is `icr.io/ibm/datashield-server`, and the tag is `0.2.38`.

Because the domain name of the registry is in the image name, to upload an image we name it with the correct address and then issue `docker push` to upload it.

1. To upload, or push, an image into Container Registry, we must first download an image. Let's use `gliderlabs/alpine:3.6`, an image stored in Docker Hub.

    `docker pull gliderlabs/alpine:3.6`

    If the repository does not include a domain name, the `docker` command line assumes that we want to use Docker Hub.

2. Rename the image ready for it to be sent to IBM Cloud Container Registry.

    `docker tag gliderlabs/alpine:3.6 us.icr.io/my_namespace/alpine:3.6`

    **Hint**: `us.icr.io` is the address of the US South instance of IBM Container Registry. Don't forget to replace `my_namespace` with your chosen namespace name.

3. Examine the container images stored your computer.

    `docker images`

    Look for `gliderlabs/alpine:3.6` and note the value in the Image ID field, then compare it to the value for `us.icr.io/my_namespace/alpine:3.6`. They should be the same, because that image is now referred to by both names on your computer.

4. Push the image to IBM Cloud Container Registry.

    `docker push us.icr.io/my_namespace/alpine:3.6`

5. Confirm that the image has been uploaded successfully.

    `ibmcloud cr images`

## Recap

You have set up an IBM Cloud Container Registry namespace in your account and pushed an image to it.

## Further Reading

**TODO**

## Next

In the next section, you will use IBM Cloud Identity and Access Management (IAM) to control access to your images.
