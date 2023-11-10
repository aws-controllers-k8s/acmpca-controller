    if ko.Status.ACKResourceMetadata != nil && ko.Status.ACKResourceMetadata.ARN != nil {
        resourceARN := (*string)(ko.Status.ACKResourceMetadata.ARN)
        csr, err := rm.getCertificateSigningRequest(ctx, *resourceARN)
        if err != nil {
            return nil, err
        }
        ko.Status.CertificateSigningRequest = csr
    }
