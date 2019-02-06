## Step 2: Provisioning an instance from the Observability dashboard {#step-2-provisioning-an-instance-from-the-observability-dashboard}

To provision an instance of IBM Log Analysis with LogDNA by using the IBM Cloud UI, complete the following steps:

### 1.  Log in to your IBM Cloud account{#1}

Log in to your IBM Cloud account

The IBM Cloud dashboard can be found at: [https://cloud.ibm.com](https://cloud.ibm.com)

After you log in with your user ID and password, the IBM Cloud UI opens.

### 2. Launch the Observability page{#2}

1. Click the Navigation Menu.

    ![](../images/logdna_img7.png)

2. Select Observability.

    ![](../images/logdna_img8.png)

### 3. Create a logging instance{#3}

From the Observability page, click **Create logging instance**.

![](../images/logdna_img9.png)

Enter a name for the service instance. For example, _logdna-marisa_

![](../images/logdna_img10.png)

Select a resource group. By default, the **Default** resource group is set.

You can select the **Default** resource group or a different one. In the sample images you can see that the DevOps resource group is selected. 

**Note:** If you cannot see a resource group or if you are not able to select a specific resource group, check that you have editing permissions on the resource group where you want to provision the instance.

Select the **Lite** service plan. By default, the **Lite** plan is set.

![](../images/logdna_img11.png)

Click **Create**.

After you provision an instance, the _Observability_ dashboard opens.

![](../images/logdna_img12.png)

Next, configure a LogDNA agent. This agent is responsible for collecting and forwarding logs to LogDNA.