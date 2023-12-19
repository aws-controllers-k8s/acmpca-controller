    err = rm.writeCertificateToSecret(ctx, resp.Certificate, ko.ObjectMeta.GetAnnotations())
    if err != nil {
        return nil, err
    }
    