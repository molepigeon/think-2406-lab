## Identify the communication patterns between the different pods {#identify-the-communication-patterns-between-the-different-pods}

To identify the communication patterns between the different pods., you can segment a single metric like **net.http.request.count** by **container.name**.

The HTTP request count metric counts the number of HTTP requests.

1.  View this metric. Complete the following steps in the _Explore_ tab:
    1.  Select **Deployment and Pods**.
    2.  Select the namespace **ticket-generator**.
    3.  Select the dashboard or metric that displays in your UI. In this figure, it is HTTP.
    4.  Select **Metrics.**
    5.  Select **Network**.
    6.  Select **net.http.request.count**.

The metric panel shows.

1.  Segment the data by **container.id**.

Select the **container.id** label:

The metric displays the average value for each container ID:

1.  Get a topology perspective of the data to see each pod and how it interacts.
    1.  Select **More options**.
    2.  Choose **Topology.**

The panel changes and displays the following view:

*   1.  Expand **ticket-generator** and double click it to enlarge it:

You can see the value of HTTP requests per container ID and how their communication pattern.