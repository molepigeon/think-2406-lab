# Trusted delegation roles and trust pinning

Stretch goal, this might not get reached within 1 hour and has potential to take a long time. Users should strongly consider skipping this and coming back to it once they've learned about the other topics!

Delegation roles let you have individuals hold keys that mean something to you when they sign an image.

You might want to require that a specific person or team or automation has signed an image before it can deploy, perhaps to indicate that they approve it going to prod. CISE can do trust pinning and enforce that it is signed by certain keys.

Make a key

Add the key as a signing role

Add the public key as a secret in the cluster

Configure CISE to pin to our new role. Also turn off VA so that we don't have to wait for VA to pass anymore.

Change an annotation in the pod, make sure image is set to the signed one. Show that CISE blocks the deploy.

Sign again, show that it uses the new key because you have it.

Push the newly signed image.

Try to deploy the new image. Show that CISE allows the deploy.
