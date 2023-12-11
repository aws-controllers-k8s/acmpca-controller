    if ko.Status.ACKResourceMetadata != nil && ko.Status.ACKResourceMetadata.ARN != nil {
        resourceARN := (*string)(ko.Status.ACKResourceMetadata.ARN)
        csr, err := rm.getCertificateAuthorityCsr(ctx, *resourceARN)
        if err != nil {
            return nil, err
        }
        ko.Status.CSR = []byte(*csr)
    }