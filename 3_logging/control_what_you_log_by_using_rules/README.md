# Control what you log by using rules {#control-what-you-log-by-using-rules}

**NOTE: To complete this step you need a paid plan. Upgrade the service plan.**

So far in this lab, you have learned how to exclude log files by configuring the LogDNA agent through the LOGDNA_EXCLUDE parameter. This configuration allows you to stop logs from being forwarded to your logging instance. But what can you do if you want to receive some of those log lines, for example, what can you do to only receive log lines that report errors? Or if you want to see the log lines but not store them so they do not add to your cost? How can you stop logs from counting towards your quota?

In this part of the lab, you will learn how to control log data and how to apply rules to control log data. In the LogDNA web UI, the section **USAGE** is where you can apply controls to the log data that you receive.