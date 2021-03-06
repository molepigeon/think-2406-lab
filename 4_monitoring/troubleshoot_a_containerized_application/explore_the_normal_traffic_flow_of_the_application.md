# Explore the normal traffic flow of the application {#explore-the-normal-traffic-flow-of-the-application}

The normal traffic flow of the application should be:

```text
Ticket client -> Load Balancer service -> Load Balancer Pod -> Backend service -> Backend pods
```

You should be able to see this flow by completing the following steps:

In the Explore tab, select the **ticket-generator** namespace.

Click ![ ](../images/sysdig_img33a.png).

Select **Topology**.

![ ](../images/sysdig_img33b.png)

Expand Topology, and select **Network Traffic**.

![ ](../images/sysdig_img34.png)

The Network Traffic default dashboard opens.

![ ](../images/sysdig_img35.png)

Click the `+` symbol. Then, double click the image to enlarge it, until you see the following:

![ ](../images/sysdig_img36.png)

**What does the diagram tell you?**
 A ‘ticket’ client requests a new ticket to the load balancer frontend, and this balancer forwards the requests to the two backend server pods. All these pods communicate using HTTP REST-style requests.
