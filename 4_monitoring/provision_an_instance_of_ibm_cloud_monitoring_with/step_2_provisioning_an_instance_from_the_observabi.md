## Step 2: Provisioning an instance from the Observability dashboard {#step-2-provisioning-an-instance-from-the-observability-dashboard}

To provision an instance of IBM Cloud Monitoring with Sysdig by using the IBM Cloud UI, complete the following steps:

1.  Log in to your IBM Cloud account.

The IBM Cloud dashboard can be found at: [https://cloud.ibm.com](https://cloud.ibm.com)

After you log in with your user ID and password, the IBM Cloud UI opens.

1.  Click the Navigation Menu.
2.  Select Observability.
3.  Click Create monitoring instance.
4.  Enter a name for the service instance. For example, _sysdig-marisa._
5.  Select the **Default** resource group.

By default, the **Default** resource group is set.

**Note:** If you are not able to select a resource group, check that you have editing permissions on the resource group where you want to provision the instance.

1.  Select the **Trial** service plan.

By default, the **Trial** plan is set.

1.  Click **Create**.

After you provision an instance, the _Observability_ dashboard opens.

Next, configure a Sysdig agent. This agent is responsible for collecting and forwarding metrics to Sysdig.