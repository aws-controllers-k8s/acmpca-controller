    err = rm.writeCertificateToSecret(ctx, *resp.Certificate, r.ko.ObjectMeta)
    if err != nil && strings.HasPrefix(err.Error(), "RequestInProgressException") {
        return &resource{ko}, ackrequeue.NeededAfter(err, ackrequeue.DefaultRequeueAfterDuration)
    }
    if err != nil {
        return nil, err
    }
