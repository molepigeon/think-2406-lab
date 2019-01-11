# Container Image Security Enforcement

Intro to CISE, what it is and what it does

Install chart (Probably from yaml because let's just not deal with Helm)

Add an ImagePolicy for default kube namespace

Set va to false for images from our namespace

Change image to vulnerable image, show that it allows the rollout because va is not enforced

Set va to true

Add an annotation to our pod, and try to apply, show that it now blocks the rollout because the image is vulnerable

Change image in pod to clean image, show that it allows the rollout
