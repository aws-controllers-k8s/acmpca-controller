    // If the certificate is marked for deletion, we will skip
    // calling GetCertificate and Writing the secret as these
    // can lead to terminal errors and unsuccessful deletion
    // in the case where the dependent secret or CA are deleted
    // first
    if r.ko.ObjectMeta.DeletionTimestamp != nil {
      return r, nil
    }
