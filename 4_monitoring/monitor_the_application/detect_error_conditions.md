## Detect error conditions {#detect-error-conditions}

You can monitor the following **metric net.http.error.count** to detect error conditions in your application. With this metric, you monitor the number of failed HTTP requests.

**NOTE**: This metric is a good candidate to set up an alert condition. In a follow up step, you will learn how to set an alert.

1.  From the _Explore_ tab, select **Deployments and Pods.**
2.  Select **ticket-generator.**
3.  Click , and select **Metrics**.
4.  Select **net.http.error.count.**
5.  Segment the data by hostname. Choose **host.hostname** in the Segment by field.

The metric panel opens and displays data segmented by hostname. If you hover over the colour line in the graph, you get the legend with the hostnames whose data is being displayed.

If the metric reports a number greater than 0 for any hostname,

1.  You can go back to the **HTTP dashboard** and check the **panel Status Codes Over Time** that represents the metric **net.http.statusCode.** This metric reports the response status codes from the application HTTP requests.
2.  You can select **More Options**.
3.  Select Topology.
4.  Expand the ticket-generator box to see in which component of the application you are receiving responses with error codes: