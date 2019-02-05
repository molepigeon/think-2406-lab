## Scenario {#scenario}

**Sample application**: A ticket token generator microservices app

The ticket generator application emulates a simple consumer/producer (or pub/sub) workflow where N consumers need to communicate with N producers. This communication is mediated and balanced by a proxy.

*   The application runs in a Kubernetes cluster.
*   The application has 3 services and 4 containers: client, balancer and 2 servers.
*   The client periodically requests a new ticket token.
*   The load balancer distributes HTTP requests through servers using HAproxy.
*   The servers run a simple sample Python application that generates a unique ticket token per request.

The server application is a simple python script that produces a unique Ticket ID per requests. The application will start by default with 1 consumer, 1 balancer and 2 servers, but you can easily scale it using the supporting [Kubernetes deployments](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/).

**The** **problem**: The application works well on the developer&#039;s laptop using Docker, but the same image triggers some 502 HTTP errors when it is deployed on a cluster that runs in the IBM Cloud Kubernetes service.