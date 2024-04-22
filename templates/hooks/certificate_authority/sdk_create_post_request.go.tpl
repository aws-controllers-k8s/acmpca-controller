    if err != nil && strings.HasPrefix(err.Error(), "RequestInProgressException") {
            return nil, ackrequeue.Needed(err)
    }