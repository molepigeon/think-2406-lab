# IBM Cloud Container Registry

In this exercise, you will use IBM Cloud Container Registry to securely deliver container images to your Kubernetes cluster. You will configure IAM policies to control access to images, and configure your cluster to prevent vulnerable or untrusted images from being used in pods.

You can use IBM Cloud Container Registry to safely store and access private Docker images in a highly available and scalable architecture.

IBM Cloud Container Registry provides a multi-tenant, highly available, and scalable private image registry that is hosted and managed by IBM. You can use the private registry by setting up your own image namespace and pushing Docker images to your namespace.

![Image showing how you can interact with IBM Cloud Container Registry.](registry-architecture.png)

## Features

### Highly available and scalable private registry

Set up your own image namespace in a multi-tenant, highly available, scalable private registry that is hosted and managed by IBM.

Securely store your private Docker images and share them with users in your IBM Cloud account.

### Security

Integrated with Identity and Access Management to provide fine-grained access controls to users within your IBM Cloud account.

Vulnerability scanning, deployment policy enforcement, and comprehensive risk assessment and prioritization provide security compliance insight and controls over static images and live containers.

Uses Notary technology to store trusted content; allows you to control and mandate signatures for your images.

Images are encrypted in transit and at rest in IBM Cloud Object Storage.

### Integration

IBM Cloud preferred image registry; pre-integrated with our Kubernetes Service for your DevOps workflow using IBM Open Toolchain or your existing CI/CD toolset.
