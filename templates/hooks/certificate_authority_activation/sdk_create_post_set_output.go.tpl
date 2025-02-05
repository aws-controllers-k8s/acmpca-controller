    if certificateSecret != "" {
        err = rm.writeCertificateChainToSecret(ctx, certificateSecret, certificateChainSecret, desired.ko.GetNamespace(), desired.ko.Spec.CompleteCertificateChainOutput)
        if err != nil {
            return nil, err
        }
    }
    
    if desired.ko.Spec.Status != nil && *desired.ko.Spec.Status == string(svcsdktypes.CertificateAuthorityStatusDisabled) {
        updateInput := &svcsdk.UpdateCertificateAuthorityInput{}

        updateInput.Status = svcsdktypes.CertificateAuthorityStatus(*desired.ko.Spec.Status)

        if desired.ko.Spec.CertificateAuthorityARN != nil {
            updateInput.CertificateAuthorityArn = desired.ko.Spec.CertificateAuthorityARN
        }

        _, err = rm.sdkapi.UpdateCertificateAuthority(ctx, updateInput)
        rm.metrics.RecordAPICall("UPDATE", "UpdateCertificateAuthority", err)
        if err != nil {
            return nil, err
        }
    }