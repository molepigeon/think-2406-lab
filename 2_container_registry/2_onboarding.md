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

## Creating an image

Containers are created from images, which contain the initial state of a container's filesystem as well as important metadata about how to start the image, such as which binary should be executed. Docker container images are defined in a Dockerfile, which defines a series of steps that should be executed in order that produce the layers of the image. Each step results in one layer in the image.

A container image name is made up of two parts, a repository and a tag, in the format `repository:tag`. The repository is a URI of where on the Internet the image can be found, and the tag is a version for the image. For example, in `icr.io/ibm/datashield-server:0.2.38`, the repository is `icr.io/ibm/datashield-server`, and the tag is `0.2.38`.

Because the domain name of the registry is in the image name, to upload an image we name it with the correct address and then issue `docker push` to upload it.

1. Create a folder for our image and navigate into it.

    `mkdir ~/myimage; cd ~/myimage`

    **Best practice tip**: Keep the folder that contains your Dockerfile as minimal as possible, and only include objects that are required for your image build. When you run `docker build` later, the contents of the folder that contains the Dockerfile and all sub-directories are added to the build context. When a lot of files are in the build context, it can take a while for the build to start while Docker transfers the build context to its daemon or the remote build server.

2. Create a Dockerfile.

    `touch Dockerfile`

3. Open the Dockerfile.

    `open Dockerfile`

4. Copy the following instructions into the Dockerfile, then save the file and exit the editor.

    ```Dockerfile
    FROM molepigeon/alpine:3.6
    ADD https://raw.githubusercontent.com/molepigeon/think-2406-lab/master/hello-world-app/hello-world /hello-world
    RUN chmod +x /hello-world
    EXPOSE 8080
    CMD ["/hello-world"]
    ```

    Let's explore this Dockerfile a bit.

    * The `FROM` instruction is an image that exists in a registry already that is used as a starting point for our image. This particular image is a copy of Alpine Linux.
    * The `ADD` instruction adds files to the image from either the filesystem or, as in this case, a web address for a Hello World web application.
    * The `EXPOSE` instruction tells the container runtime to expect that this container will be running a server on the given port, in this case 8080.
    * The `CMD` instruction defines the command to be executed when a container is started from this image.

5. Build the image. Name the image ready for it to be sent to IBM Cloud Container Registry

    `docker build -t registry.ng.bluemix.net/my_namespace/hello-world:3.6 .`

    **Hint**: `registry.ng.bluemix.net` is the address of the US South instance of IBM Container Registry. Don't forget to replace `my_namespace` with your chosen namespace name.

## Pushing your image

1. Log in to IBM Cloud Container Registry for pushing images.

    `ibmcloud cr login`

    This command wraps `docker login` to authenticate your Docker command line with IBM Cloud Container Registry.

2. Push the image to IBM Cloud Container Registry.

    `docker push registry.ng.bluemix.net/my_namespace/hello-world:3.6`

3. Confirm that the image has been uploaded successfully.

    `ibmcloud cr images`

    Note that the image list shows vulnerabilities in the status column. You will explore Vulnerability Advisor and resolve these vulnerabilities later in this tutorial.

## Recap

You have set up an IBM Cloud Container Registry namespace in your account, built an image, and pushed your image into your namespace.

## Further reading

[IBM Cloud Container Registry product page](https://icr.io)

[Planning Namespaces (IBM Cloud Docs)](https://console.bluemix.net/docs/services/Registry/registry_overview.html#registry_namespaces)

[Adding images to your namespace (IBM Cloud Docs)](https://console.bluemix.net/docs/services/Registry/registry_images_.html#registry_images_)

## Next

In the next section, you will use IBM Cloud Identity and Access Management (IAM) to control access to your images.
