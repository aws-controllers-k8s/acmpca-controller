    if err != nil {
		input := &svcsdk.DescribeCertificateAuthorityInput{}
		input.CertificateAuthorityArn = desired.ko.Spec.CertificateAuthorityARN

		var describeResp *svcsdk.DescribeCertificateAuthorityOutput
		describeResp, describeErr := rm.sdkapi.DescribeCertificateAuthorityWithContext(ctx, input)
		rm.metrics.RecordAPICall("READ_ONE", "DescribeCertificateAuthority", err)
		if describeErr != nil {
			return desired, ackrequeue.NeededAfter(describeErr, ackrequeue.DefaultRequeueAfterDuration)
		}

		if *describeResp.CertificateAuthority.Status != svcsdk.CertificateAuthorityStatusFailed && *describeResp.CertificateAuthority.Status != svcsdk.CertificateAuthorityStatusDeleted && *describeResp.CertificateAuthority.Status != svcsdk.CertificateAuthorityStatusDisabled {
			return desired, ackrequeue.NeededAfter(err, ackrequeue.DefaultRequeueAfterDuration)
		}

		return nil, err
	}