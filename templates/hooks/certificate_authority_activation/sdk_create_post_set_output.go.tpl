    if certificateSecret != "" {
        err = rm.writeCertificateChainToSecret(ctx, certificateSecret, certificateChainSecret, desired.ko.ObjectMeta)
        if err != nil {
            return nil, err
        }
    }
    
    if desired.ko.Spec.Status != nil && *desired.ko.Spec.Status == svcsdk.CertificateAuthorityStatusDisabled {
        updateInput := &svcsdk.UpdateCertificateAuthorityInput{}

        updateInput.SetStatus(*desired.ko.Spec.Status)

        if desired.ko.Spec.CertificateAuthorityARN != nil {
            updateInput.SetCertificateAuthorityArn(*desired.ko.Spec.CertificateAuthorityARN)
        }

        _, err = rm.sdkapi.UpdateCertificateAuthorityWithContext(ctx, updateInput)
        rm.metrics.RecordAPICall("UPDATE", "UpdateCertificateAuthority", err)
        if err != nil {
            return nil, err
        }
    }