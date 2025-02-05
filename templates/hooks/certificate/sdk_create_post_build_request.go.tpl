    if desired.ko.Spec.CertificateSigningRequest != nil {
		input.Csr = []byte(*desired.ko.Spec.CertificateSigningRequest)
	}
	input.IdempotencyToken = aws.String(string(desired.ko.ObjectMeta.UID))