    err = rm.writeCertificateToSecret(ctx, resp.Certificate, r.ko.ObjectMeta)
    if err != nil {
        return nil, err
    }
    