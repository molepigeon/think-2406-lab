## Filter out logs from the UI from a namespace except error entries {#filter-out-logs-from-the-ui-from-a-namespace-except-error-entries}

Stop showing logs from a namespace, for example, ibm-observe namespace logs, except the ones that report an error:

In this step, you will learn how to exclude data from the namespace ibm-observe from the cluster while keeping entries that report errors only. You will configure the rule so that you are NOT able to see all log data through views and therefore not be able to define alerts on all the data.

Complete the following steps:

1. Select the settings icon ![image52](../images/logdna_img52.png).

2. Select **USAGE**.

    ![image53](../images/logdna_img53.png)

    The _Manage Usage_ page opens.

    ![image58](../images/logdna_img58.png)

3. In the **Exclusion Rules** section, select **Add rule**.

    ![image5](../images/logdna_img59.png)

4. Enter the following query in the **Query** field:

    ```
    Namespace:ibm-observe -level:error
    ```

    ![image60](../images/logdna_img60.png)
