    resourceARN := (*string)(ko.Status.ACKResourceMetadata.ARN)
    tags, err := rm.getTags(ctx, *resourceARN)
    if err != nil {
        return nil, err
    }
    ko.Spec.Tags = tags

    ko.Status.CertificateSigningRequest, err = rm.getCertificateAuthorityCsr(ctx, *resourceARN)
    if err != nil && strings.HasPrefix(err.Error(), "RequestInProgressException") {
        return nil, ackrequeue.NeededAfter(err, ackrequeue.DefaultRequeueAfterDuration)
    }
    if err != nil {
        return nil, err
    }