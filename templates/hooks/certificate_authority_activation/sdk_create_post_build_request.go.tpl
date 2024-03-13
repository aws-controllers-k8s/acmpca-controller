    certificateSecret := ""
	certificateChainSecret := ""
	if desired.ko.Spec.Certificate != nil {
		certificateSecret, err = rm.rr.SecretValueFromReference(ctx, desired.ko.Spec.Certificate)
		if err != nil {
			return nil, ackrequeue.Needed(err)
		}
		if certificateSecret != "" {
			input.SetCertificate([]byte(certificateSecret))
		}
	}
	if desired.ko.Spec.CertificateChain != nil {
		certificateChainSecret, err = rm.rr.SecretValueFromReference(ctx, desired.ko.Spec.CertificateChain)
		if err != nil {
			return nil, ackrequeue.Needed(err)
		}
		if certificateChainSecret != "" {
			input.SetCertificateChain([]byte(certificateChainSecret))
		}
	}