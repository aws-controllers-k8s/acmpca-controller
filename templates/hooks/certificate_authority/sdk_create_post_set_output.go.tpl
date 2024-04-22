    if ko.Status.ACKResourceMetadata != nil && ko.Status.ACKResourceMetadata.ARN != nil {
        resourceARN := (*string)(ko.Status.ACKResourceMetadata.ARN)
        ko.Status.CertificateSigningRequest, err = rm.getCertificateAuthorityCsr(ctx, *resourceARN)
        for err != nil && strings.HasPrefix(err.Error(), "RequestInProgressException") {
            time.Sleep(1 * time.Second)
            ko.Status.CertificateSigningRequest, err = rm.getCertificateAuthorityCsr(ctx, *resourceARN)
        }
        if err != nil {
            return nil, err
        }
    }