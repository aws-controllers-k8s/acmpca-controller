    resourceARN := (*string)(ko.Status.ACKResourceMetadata.ARN)
    caARN := (*string)(ko.Spec.CertificateAuthorityARN)
    err = rm.writeCertificateToSecret(ctx, *resourceARN, *caARN, desired.ko.ObjectMeta)
    if err != nil {
        return nil, err
    }