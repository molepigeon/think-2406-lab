## Monitor the Kubernetes namespace where the application is running {#monitor-the-kubernetes-namespace-where-the-application-is-running}

Sysdig collects the kube-state-metrics out of the box. These metrics provide full insight about the state of your Kubernetes cluster. For example, these metrics can tell you:

*   The number of instances that are running.
*   The status of the pods. Are the pods running, available and ready?
*   Which entities compose the application namespace. Are they up and running?

You can use the **Kubernetes State Overview** dashboard to monitor your Kubernetes cluster.

1.  From the _Explore_ tab, select **Deployments and Pods.**
2.  Select **ticket-generator.** This sets the scope to the namespace of the sample application.
3.  Click . Select **Default** **Dashboards**. Select **Kubernetes**. Then, select **Kubernetes State**.
4.  Select **Kubernetes State Overview.**

The page opens:

From these dashboard, you can see the number of containers, pods, deployments, services, etc. You can also detect if any of your pods is restarting frequently or currently unavailable.