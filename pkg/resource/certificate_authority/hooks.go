// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package certificate_authority

import (
	"context"

	svcapitypes "github.com/aws-controllers-k8s/acmpca-controller/apis/v1alpha1"
	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	ackrtlog "github.com/aws-controllers-k8s/runtime/pkg/runtime/log"
	svcsdk "github.com/aws/aws-sdk-go-v2/service/acmpca"
	svcsdktypes "github.com/aws/aws-sdk-go-v2/service/acmpca/types"
)

func (rm *resourceManager) getCertificateAuthorityCsr(
	ctx context.Context,
	resourceARN string,
) (*string, error) {
	input := &svcsdk.GetCertificateAuthorityCsrInput{}
	input.CertificateAuthorityArn = &resourceARN
	resp, err := rm.sdkapi.GetCertificateAuthorityCsr(ctx, input)
	rm.metrics.RecordAPICall("READ_ONE", "GetCertificateAuthorityCsr", err)
	if err != nil {
		return nil, err
	}
	csr := resp.Csr
	return csr, err
}

// getTags gets tags from given CA.
func (rm *resourceManager) getTags(
	ctx context.Context,
	resourceARN string,
) ([]*svcapitypes.Tag, error) {
	resp, err := rm.sdkapi.ListTags(
		ctx,
		&svcsdk.ListTagsInput{
			CertificateAuthorityArn: &resourceARN,
		},
	)
	rm.metrics.RecordAPICall("GET", "ListTags", err)
	if err != nil {
		return nil, err
	}
	tags := resourceTagsFromSDKTags(resp.Tags)
	return tags, nil
}

// syncTags updates tags of given CA to desired tags.
func (rm *resourceManager) syncTags(
	ctx context.Context,
	desired *resource,
	latest *resource,
) (err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.syncTags")
	defer func(err error) { exit(err) }(err)

	arn := (*string)(latest.ko.Status.ACKResourceMetadata.ARN)

	desiredTags, _ := convertToOrderedACKTags(desired.ko.Spec.Tags)
	latestTags, _ := convertToOrderedACKTags(latest.ko.Spec.Tags)

	added, _, removed := ackcompare.GetTagsDifference(latestTags, desiredTags)

	toAdd := fromACKTags(added, nil)
	toRemove := fromACKTags(removed, nil)

	/*var toDelete []*string
	for _, removedElement := range toRemove {
		toDelete = append(toDelete, removedElement.Key)
	}*/

	if len(toRemove) > 0 {
		rlog.Debug("removing tags from CertificateAuthority", "tags", toRemove)
		_, err = rm.sdkapi.UntagCertificateAuthority(
			ctx,
			&svcsdk.UntagCertificateAuthorityInput{
				CertificateAuthorityArn: arn,
				Tags:                    sdkTagsFromResourceTags(toRemove),
			},
		)
		rm.metrics.RecordAPICall("UPDATE", "UntagCertificateAuthority", err)
		if err != nil {
			return err
		}
	}

	if len(toAdd) > 0 {
		rlog.Debug("adding tags to CertificateAuthority", "tags", toAdd)
		_, err = rm.sdkapi.TagCertificateAuthority(
			ctx,
			&svcsdk.TagCertificateAuthorityInput{
				CertificateAuthorityArn: arn,
				Tags:                    sdkTagsFromResourceTags(toAdd),
			},
		)
		rm.metrics.RecordAPICall("UPDATE", "TagCertificateAuthority", err)
		if err != nil {
			return err
		}
	}

	return nil
}

func sdkTagsFromResourceTags(
	rTags []*svcapitypes.Tag,
) []svcsdktypes.Tag {
	tags := make([]svcsdktypes.Tag, len(rTags))
	for i := range rTags {
		tags[i] = svcsdktypes.Tag{
			Key:   rTags[i].Key,
			Value: rTags[i].Value,
		}
	}
	return tags
}

func resourceTagsFromSDKTags(
	sdkTags []svcsdktypes.Tag,
) []*svcapitypes.Tag {
	tags := make([]*svcapitypes.Tag, len(sdkTags))
	for i := range sdkTags {
		tags[i] = &svcapitypes.Tag{
			Key:   sdkTags[i].Key,
			Value: sdkTags[i].Value,
		}
	}
	return tags
}
