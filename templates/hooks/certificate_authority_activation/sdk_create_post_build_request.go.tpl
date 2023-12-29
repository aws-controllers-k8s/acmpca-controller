    if desired.ko.Spec.Certificate != nil {
		tmpSecret, err := rm.rr.SecretValueFromReference(ctx, desired.ko.Spec.Certificate)
		if err != nil {
			return nil, ackrequeue.Needed(err)
		}
		if tmpSecret != "" {
			input.SetCertificate([]byte(tmpSecret))
		}
	}