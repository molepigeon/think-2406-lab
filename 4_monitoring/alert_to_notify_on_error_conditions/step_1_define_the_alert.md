## Step 1: Define the alert {#step-1-define-the-alert}

### 1. Add an alert{#1}

In the *Alerts* tab, select **Add Alert**.

![](../images/sysdig_img76.png)

Select the Alert Type **Metric**.

![](../images/sysdig_img77.png)

The _New Metric Alert_ page opens.

![](../images/sysdig_img78.png)

Enter the following name for the metric: **[APM] Ticket generator HTTP errors**

1. Click **New Metric Alert**

2. Enter the name of the alert

Choose the severity of the alert:

1. Click **warning**.

2. Select **error**.

![](../images/sysdig_img79.png)

### 2. Define metric details{#2}

Select how to group the data for the metric on which you want to define the alert. Choose **Average**.

![](../images/sysdig_img80.png)

Choose the metric **net.http.request.count** for which you want to define the alert.

![](../images/sysdig_img81.png)

### 3. Set the scope{#3}

**Filter the scope of the data to be monitored to the namespace where the application is running.**

Select **kubernetes.namespace.name**.

![](../images/sysdig_img82.png)

Select the operand **is**.

![](../images/sysdig_img83.png)

Select the value **ticket-generator**.

![](../images/sysdig_img84.png)

**Add a second rule to limit the scope so the alert is triggered when specific status codes are identified.**

Select the label **net.http.statusCode**.

![](../images/sysdig_img85.png)

And the values that you want to alert on: **500, 400, 404**

![](../images/sysdig_img86.png)

**NOTE**: When you look for the first time for code 500, you cannot see it. However, type it in the search bar and select it. It will show in the list.

![](../images/sysdig_img87.png)

### 4. Set the trigger condition{#4}

Set the trigger condition as shown in the image:

![](../images/sysdig_img88.png)

The definition section should look like:

![](../images/sysdig_img89.png)

Click **Save**.

You can see the alert defined in the _Alerts_ dashboard.

![](../images/sysdig_img90.png)