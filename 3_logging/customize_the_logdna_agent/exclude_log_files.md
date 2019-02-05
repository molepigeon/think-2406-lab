## Exclude log files {#exclude-log-files}

Complete the following steps to exclude log files:

1. From the terminal where you set the cluster context in a previous step, generate the configuration file of the agent by running the following command:

    ```
    kubectl get daemonset logdna-agent -o=yaml &gt; prod-logdna-agent-ds.yaml
    ```

2. Make changes. Add the section **LOGDNA_EXCLUDE**, and exclude *calico-node* logs. Add to the yaml file:

    ```
    - name: LOGDNA_EXCLUDE
      value: /var/log/containers/calico-node*
    ```

    Other examples of exclusions:

    To exclude all of the _kube-system_ logs, enter:

    ```
    - name: LOGDNA_EXCLUDE
      value: /var/log/containers/*_kube-system_*
    ```

    To exclude all non- container logs, that is to exclude files as shown in the All Apps filter view, enter:

    ![image42](images/logdna_img42.png)

    ```
    - name: LOGDNA_EXCLUDE
      value: /var/log/!(containers)*
    ```

    To exclude calico logs, enter:

    ```
    - name: LOGDNA_EXCLUDE
      value: /var/log/containers/calico*
    ```

    To exclude all of the _kube-system_ logs and all non-container logs, enter:

    ```
    - name: LOGDNA_EXCLUDE
      value: /var/log/!(containers)*,/var/log/containers/*_kube-system_*
    ```

    For example, the following section shows where to add tags in the configuration file:

    ```
    apiVersion: extensions/v1beta1
    kind: DaemonSet
    metadata:
      name: logdna-agent
    spec:
      template:
        metadata:
          labels:
            app: logdna-agent
        spec:
          containers:
          - name: logdna-agent
            image: logdna/logdna-agent:latest
            imagePullPolicy: Always
            env:
            - name: LOGDNA_AGENT_KEY
              valueFrom:
                secretKeyRef:
                name: logdna-agent-key
                key: logdna-agent-key
            - name: LDAPIHOST
              value: api.us-south.logging.cloud.ibm.com
            - name: LDLOGHOST
              value: logs.us-south.logging.cloud.ibm.com
            - name: LOGDNA_PLATFORM
              value: k8s
            - name: LOGDNA_TAGS
              value: tag1,tag2,tag3
            - name: LOGDNA_EXCLUDE
              value: /var/log/containers/calico-node*
    ```
    
    ![image43](images/logdna_img43.png)


3. Apply the configuration changes. Run the following command:

    ```
    kubectl apply -f prod-logdna-agent-ds.yaml
    ```

4. Get the logdna-agent pods. Run the following command:

    ```
    kubectl get pods
    ```

    ![image44](images/logdna_img44.png)

5. Delete all the logdna pods that are listed in the previous step.

    ```
    kubectl delete pod PodName
    ```

    For example:

    ```
    $ kubectl delete pod logdna-agent-g6qv6
    pod "logdna-agent-g6qv6" deleted
    ```

6. Verify that calico-node entries are not showing in the LogDNA web UI.