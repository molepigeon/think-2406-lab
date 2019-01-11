# IBM Cloud Container Registry

Use IBM Cloud Container Registry to safely store and access private Docker images in a highly available and scalable architecture.

IBM Cloud Container Registry provides a multi-tenant, highly available, and scalable private image registry that is hosted and managed by IBM. You can use the private registry by setting up your own image namespace and pushing Docker images to your namespace.

![Image showing how you can interact with IBM Cloud Container Registry. Container Registry contains both private and public repositories, and APIs to interact with the service. Your local Docker client can pull and push images to and from your private repositories in the registry, and can pull public repositories. The IBM Cloud web UI (console) interacts with the Container Registry API to list images. The Container Registry CLI interacts with the API to list, create, inspect, and remove images, as well as other administrative functions. Your local Docker client can also pull and push images from your local image store to other registries.](https://console.bluemix.net/docs/api/content/services/Registry/images/registry_architecture1.svg)

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