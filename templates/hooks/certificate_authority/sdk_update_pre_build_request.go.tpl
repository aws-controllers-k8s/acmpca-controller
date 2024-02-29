	if delta.DifferentAt("Spec.Tags") {
		err := rm.syncTags(ctx, desired, latest)
		if err != nil {
			return nil, err
		}
	}