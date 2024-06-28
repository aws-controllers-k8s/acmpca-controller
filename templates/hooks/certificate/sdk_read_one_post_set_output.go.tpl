    err = rm.writeCertificateToSecret(ctx, *resp.Certificate, r.ko.ObjectMeta)
    // If the Secret cannot be written to requeue and wait for secret to exist
    if err != nil {
        return &resource{ko}, ackrequeue.NeededAfter(err, ackrequeue.DefaultRequeueAfterDuration)
    }
