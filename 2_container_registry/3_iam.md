# Securing access to your private registry

In this section, you will configure access controls for IBM Cloud Container Registry using IBM Cloud Identity and Access Management (IAM). IAM allows you to configure policies that control who can access certain resources in your IBM Cloud account. You can configure policies for users in your account, and you can create and bind policies to Service IDs for use in automation.

## Creating an IAM Service ID and API Key

An IAM Service ID is a special user that can access IBM Cloud APIs, but cannot log in to the IBM Cloud CLI or the Dashboard. You can bind policies to a Service ID to control what resources the ID can access, and you can manage API keys for the Service ID in IAM.

1. Log in to your IBM Cloud account.

    ```bash
    ibmcloud login
    ```

2. Create an IAM Service ID called `think-registry-demo`. The description field can be any text that you want to help you to recognise the Service ID later.

    ```bash
    ibmcloud iam service-id-create think-registry-demo --description "Service ID created for session 2406 at Think 2019"
    ```

3. List your Service IDs to see your new Service ID in the list.

    ```bash
    ibmcloud iam service-ids
    ```

    ```text
    Getting all services IDs bound to current account as michaelh@uk.ibm.com...
    OK
    UUID                                             Name                  Created At              Last Updated            Description                                         Locked
    ServiceId-ba934ca1-36a4-4d6c-a649-af45a2fc184b   think-registry-demo   2019-01-29T10:57+0000   2019-01-29T10:57+0000   Service ID created for session 2406 at Think 2019   false
    ```

4. Create an API key for your new Service ID called `think-registry-demo-key`.

    ```bash
    ibmcloud iam service-api-key-create think-registry-demo-key think-registry-demo
    ```

    Find your API key in the output. You will need the API key in later steps.

## Enabling IAM policy enforcement for IBM Cloud Container Registry

Make sure that IAM policy enforcement is enabled in IBM Cloud Container Registry for your account.

**Note:** If multiple people use IBM Cloud Container Registry in your IBM Cloud account, you must have policies in place for your users before you enable policy enforcement so that your users retain access to images in your account. If you do not have policies configured for your users, use a separate IBM Cloud account for this tutorial.

1. Make sure that you are targeting the correct IBM Cloud account.

    ```bash
    ibmcloud target
    ```

2. Make sure that you are targeting the US South instance of IBM Cloud Container Registry.

    ```bash
    ibmcloud cr region-set us-south
    ```

3. Enable IAM policy enforcement for IBM Cloud Container Registry in your account.

    ```bash
    ibmcloud cr iam-policies-enable
    ```

## Configuring IAM Policies

Now that you have IAM policy enforcement enabled for your account, you must create policies to control the resources that your Service ID can access.

Let's prove that we can't pull images using the Service ID.

1. Log in to Container Registry using the Service ID's API key.

    ```bash
    docker login -u iamapikey --password "<your_api_key>" registry.ng.bluemix.net
    ```

    **Hint**: `registry.ng.bluemix.net` is the address of the US South instance of IBM Container Registry.

2. Try to pull the image that you pushed earlier.

    ```bash
    docker pull registry.ng.bluemix.net/my_namespace/hello-world:3.6
    ```

    **Hint**: Don't forget to replace `my_namespace` with your chosen namespace name.

    The pull fails with an error message:

    `Error response from daemon: pull access denied for ...`

Now let's create an IAM policy to allow your Service ID to pull, but not push, images in your namespace.

1. Create an IAM policy to grant Reader role on your namespace to the service ID.

    ```bash
    ibmcloud iam service-policy-create think-registry-demo --roles Reader --service-name container-registry --region us-south --resource-type namespace --resource my_namespace
    ```

    **Hint**: Don't forget to replace `my_namespace` with your chosen namespace name.

2. Try your image pull again.

    ```bash
    docker pull registry.ng.bluemix.net/my_namespace/hello-world:3.6
    ```

    The pull works because your IAM policy allows the Service ID to access the image.

3. Try to push the image back to the registry.

    ```bash
    docker push registry.ng.bluemix.net/my_namespace/hello-world:3.6
    ```

    The push is not allowed because your Service ID does not have Writer access to IBM Cloud Container Registry.

    ```example
    The push refers to repository [registry.ng.bluemix.net/my_namespace/hello-world]
    aba2478e24ae: Layer already exists
    01b17e92d16a: Layer already exists
    5bef08742407: Layer already exists
    errors:
    denied: requested access to the resource is denied
    unauthorized: authentication required
    ```

## Recap

You have created an IBM Cloud IAM Service ID and granted it reader access to your IBM Cloud Container Registry namespace. The Service ID can pull images in your namespace, but it cannot push images or access resources in any other IBM Cloud Container Registry regions, or other IBM Cloud services.

You can use Service IDs to finely control automation access to the resources in your account. You can add more than one policy to a Service ID to give it the access that it needs.

You can also create policies for other users in your account to control what resources they have access to.

## Further reading

[Automating access to IBM Cloud Container Registry (IBM Cloud Docs)](https://console.bluemix.net/docs/services/Registry/registry_tokens.html#registry_access)

[Defining user access role policies (IBM Cloud Docs)](https://console.bluemix.net/docs/services/Registry/registry_users.html#user)

[Managing user access with Identity and Access Management (IBM Cloud Docs)](https://console.bluemix.net/docs/services/Registry/iam.html#iam)

## Next

In the next section, you will set up your Kubernetes cluster to access your images.
