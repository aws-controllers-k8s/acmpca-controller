    if certificateSecret != "" {
        err = rm.writeCertificateChainToSecret(ctx, certificateSecret, certificateChainSecret, desired.ko.ObjectMeta)
        if err != nil {
            return nil, err
        }
    }