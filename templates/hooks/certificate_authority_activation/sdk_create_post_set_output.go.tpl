    readOneInput := &svcsdk.DescribeCertificateAuthorityInput{}
	readOneInput.CertificateAuthorityArn = ko.Spec.CertificateAuthorityARN

	var readOneResp *svcsdk.DescribeCertificateAuthorityOutput
	readOneResp, err = rm.sdkapi.DescribeCertificateAuthorityWithContext(ctx, readOneInput)
	rm.metrics.RecordAPICall("READ_ONE", "DescribeCertificateAuthority", err)
	if err != nil {
		return nil, err
	}
    ko.Status.Status = readOneResp.CertificateAuthority.Status