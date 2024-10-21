# `Pod` stuck on `CrashLoopBackoff` and no body cares

**Scenario:** One-Time Database Migration/Backup `Job`

Imagine you have a Kubernetes `Job` that performs a one-time database migration or backup.
The `Job` is supposed to copy data from a source to a destination
(e.g., migrate or backup database files to a persistent storage volume).
On its first run, the `Job` completes successfully,
migrating all the necessary data to the destination directory.

However, the next time the `Job` is triggered (intentionally or accidentally),
it finds that all files already exist in the destination directory (the same files are already backed up).

Now the `Job` cannot copy the files as they already exist,
it crashed.

Kubernetes tries to restart the `Job`, which leads to the `Pod` repeatedly crashing and entering the `CrashLoopBackOff` state.
