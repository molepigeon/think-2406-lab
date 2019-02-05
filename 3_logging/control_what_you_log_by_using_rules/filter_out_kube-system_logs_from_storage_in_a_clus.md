## Filter out kube-system logs from storage in a cluster except error entries {#filter-out-kube-system-logs-from-storage-in-a-cluster-except-error-entries}

In this step, you will learn how to exclude kube-system data from the cluster while keeping entries that report errors only. You will configure the rule so that you are able to see all log data through views and be able to define alerts on all the data.

Complete the following steps:

1. Select the settings icon ![image52](../images/logdna_img52.png).

2. Select **USAGE**.

    ![image53](../images/logdna_img53.png)

    The _Manage Usage_ page opens.

    ![image54](../images/logdna_img54.png)

3. In the **Exclusion Rules** section, select **Add rule**.

    ![image55](../images/logdna_img55.png)

4. Enter the following query in the **Query** field:

    ```
    Namespace:kube-system -level:error
    ```

    ![image56](../images/logdna_img56.png)