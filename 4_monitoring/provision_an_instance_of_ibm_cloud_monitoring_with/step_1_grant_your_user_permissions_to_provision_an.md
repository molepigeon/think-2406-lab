## Step 1: Grant your user permissions to provision an instance in the IBM Cloud {#step-1-grant-your-user-permissions-to-provision-an-instance-in-the-ibm-cloud}

**Note:** This step must be completed by the account owner or an administrator of the IBM Cloud Monitoring with Sysdig service on the IBM Cloud.

Complete the following steps to assign a user administrator role to the IBM Cloud Monitoring with Sysdig service:

1.  From the menu bar, click **Manage** &gt; **Access (IAM)**.
2.  Create an access group: sysdig-admins
    1.  Select Access Groups.
    2.  Click **Create**.
    3.  Enter the name of the access group: **sysdig-admins**
    4.  [Optional] Enter a description

1.  Add an access policy for theresource group where you are going to provision the IBM Cloud Monitoring with Sysdig service. In the lab, use Default.
    1.  Click Access Policies.
    2.  Click Assign Access.
    3.  Select Assign Access within a Resource Group.
    4.  Select the name of the resource group. Choose Default.
    5.  Select the role. Choose Administrator.
    6.  Click **Assign**.
2.  Add an access policy to work with the IBM Cloud Monitoring with Sysdig service.
    1.  Click Access Policies.
    2.  Click Assign Access.
    3.  Select Assign Access to resources.
    4.  Select IBM Cloud Monitoring with Sysdig.
    5.  Select All instances.
    6.  Select administrator role for the platform role.
    7.  Select manager role for the service role.
    8.  Click **Assign**.
3.  Select **Users**.
    1.  Click Add users.
    2.  Choose your IBM ID.
    3.  Click Add to group.