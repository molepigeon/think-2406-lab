# Identify the communication patterns between the different pods {#identify-the-communication-patterns-between-the-different-pods}

To identify the communication patterns between the different pods., you can segment a single metric like **net.http.request.count** by **container.name**.

The HTTP request count metric counts the number of HTTP requests.

![ ](../images/sysdig_img41.png)

## 1. View the HTTP request count metric{#1}

Complete the following steps in the *Explore* tab:

Select **Deployment and Pods**.

Select the namespace **ticket-generator**.

Click ![ ](../images/sysdig_img33a.png).

Select **Metrics** > **Network** > **net.http.request.count**

![ ](../images/sysdig_img42.png)

The metric panel shows.

![ ](../images/sysdig_img43.png)

### 2. Segment the data by **container.id**{#2}

Look for the field **Segmented by**.

![ ](../images/sysdig_img44.png)

Search and select the **container.id** label.

![ ](../images/sysdig_img45.png)

The metric displays the average value for each container ID:

![ ](../images/sysdig_img46.png)

### Get a topology perspective of the data to see each pod and how it interacts{#3}

Select **More options**.

![ ](../images/sysdig_img47.png)

Choose **Topology.**

![ ](../images/sysdig_img48a.png)

![ ](../images/sysdig_img48.png)

The panel changes and displays the following view:

![ ](../images/sysdig_img49.png)

Expand **ticket-generator** and double click it to enlarge it:

![ ](../images/sysdig_img50.png)

You can see the value of HTTP requests per container ID and their communication pattern.
