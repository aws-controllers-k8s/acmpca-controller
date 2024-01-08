    if desired.ko.Spec.Status != nil && (*desired.ko.Spec.Status == svcsdk.CertificateAuthorityStatusActive || *desired.ko.Spec.Status == svcsdk.CertificateAuthorityStatusDisabled) {
		input.SetStatus(*desired.ko.Spec.Status)
	}