    if err != nil {
		input := &svcsdk.DescribeCertificateAuthorityInput{}
		input.CertificateAuthorityArn = desired.ko.Spec.CertificateAuthorityARN

		var describeResp *svcsdk.DescribeCertificateAuthorityOutput
		describeResp, describeErr := rm.sdkapi.DescribeCertificateAuthorityWithContext(ctx, input)
		rm.metrics.RecordAPICall("READ_ONE", "DescribeCertificateAuthority", err)
		if describeErr != nil {
			return nil, describeErr
		}

		if *describeResp.CertificateAuthority.Status == svcsdk.CertificateAuthorityStatusCreating || *describeResp.CertificateAuthority.Status == svcsdk.CertificateAuthorityStatusPendingCertificate {
			return desired, ackrequeue.NeededAfter(err, ackrequeue.DefaultRequeueAfterDuration)
		}

		return nil, err
	}