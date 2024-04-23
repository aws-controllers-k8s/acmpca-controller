    resourceARN := (*string)(ko.Status.ACKResourceMetadata.ARN)
    tags, err := rm.getTags(ctx, *resourceARN)
    if err != nil {
        return nil, err
    }
    ko.Spec.Tags = tags

    if ko.Spec.KeyStorageSecurityStandard == nil {
        ko.Spec.KeyStorageSecurityStandard = aws.String("FIPS_140_2_LEVEL_3_OR_HIGHER")
    }

    if ko.Spec.UsageMode == nil {
        ko.Spec.UsageMode = aws.String("GENERAL_PURPOSE")
    }

    if ko.Spec.RevocationConfiguration == nil {
        revocationConfiguration := &svcapitypes.RevocationConfiguration{}

		revocationConfigurationCRLConfiguration := &svcapitypes.CRLConfiguration{}
		revocationConfigurationCRLConfiguration.Enabled = aws.Bool(false)
        revocationConfiguration.CRLConfiguration = revocationConfigurationCRLConfiguration

        revocationConfigurationOCSPConfiguration := &svcapitypes.OCSPConfiguration{}
		revocationConfigurationOCSPConfiguration.Enabled = aws.Bool(false)
        revocationConfiguration.OCSPConfiguration = revocationConfigurationOCSPConfiguration

		
		ko.Spec.RevocationConfiguration = revocationConfiguration
	}
    
    ko.Status.CertificateSigningRequest, err = rm.getCertificateAuthorityCsr(ctx, *resourceARN)
    if err != nil && strings.HasPrefix(err.Error(), "RequestInProgressException") {
        return nil, ackrequeue.NeededAfter(err, ackrequeue.DefaultRequeueAfterDuration)
    }
    if err != nil {
        return nil, err
    }