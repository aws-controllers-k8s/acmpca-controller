    if err != nil {
		input := &svcsdk.DescribeCertificateAuthorityInput{}
		input.CertificateAuthorityArn = desired.ko.Spec.CertificateAuthorityARN

		var describeResp *svcsdk.DescribeCertificateAuthorityOutput
		describeResp, describeErr := rm.sdkapi.DescribeCertificateAuthority(ctx, input)
		rm.metrics.RecordAPICall("READ_ONE", "DescribeCertificateAuthority", err)
		if describeErr != nil {
			return desired, ackrequeue.NeededAfter(describeErr, ackrequeue.DefaultRequeueAfterDuration)
		}

		if describeResp.CertificateAuthority.Status != svcsdktypes.CertificateAuthorityStatusFailed && describeResp.CertificateAuthority.Status != svcsdktypes.CertificateAuthorityStatusDeleted && describeResp.CertificateAuthority.Status != svcsdktypes.CertificateAuthorityStatusDisabled {
			return desired, ackrequeue.NeededAfter(err, ackrequeue.DefaultRequeueAfterDuration)
		}

		return nil, err
	}