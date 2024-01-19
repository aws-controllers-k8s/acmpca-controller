    if certificateSecret != "" {
        err = rm.writeCertificateChainToSecret(ctx, certificateSecret, certificateChainSecret, desired)
        if err != nil {
            return nil, err
        }
    }