//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by ack-generate. DO NOT EDIT.

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	corev1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIPassthrough) DeepCopyInto(out *APIPassthrough) {
	*out = *in
	if in.Extensions != nil {
		in, out := &in.Extensions, &out.Extensions
		*out = new(Extensions)
		(*in).DeepCopyInto(*out)
	}
	if in.Subject != nil {
		in, out := &in.Subject, &out.Subject
		*out = new(ASN1Subject)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIPassthrough.
func (in *APIPassthrough) DeepCopy() *APIPassthrough {
	if in == nil {
		return nil
	}
	out := new(APIPassthrough)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ASN1Subject) DeepCopyInto(out *ASN1Subject) {
	*out = *in
	if in.CommonName != nil {
		in, out := &in.CommonName, &out.CommonName
		*out = new(string)
		**out = **in
	}
	if in.Country != nil {
		in, out := &in.Country, &out.Country
		*out = new(string)
		**out = **in
	}
	if in.CustomAttributes != nil {
		in, out := &in.CustomAttributes, &out.CustomAttributes
		*out = make([]*CustomAttribute, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(CustomAttribute)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.DistinguishedNameQualifier != nil {
		in, out := &in.DistinguishedNameQualifier, &out.DistinguishedNameQualifier
		*out = new(string)
		**out = **in
	}
	if in.GenerationQualifier != nil {
		in, out := &in.GenerationQualifier, &out.GenerationQualifier
		*out = new(string)
		**out = **in
	}
	if in.GivenName != nil {
		in, out := &in.GivenName, &out.GivenName
		*out = new(string)
		**out = **in
	}
	if in.Initials != nil {
		in, out := &in.Initials, &out.Initials
		*out = new(string)
		**out = **in
	}
	if in.Locality != nil {
		in, out := &in.Locality, &out.Locality
		*out = new(string)
		**out = **in
	}
	if in.Organization != nil {
		in, out := &in.Organization, &out.Organization
		*out = new(string)
		**out = **in
	}
	if in.OrganizationalUnit != nil {
		in, out := &in.OrganizationalUnit, &out.OrganizationalUnit
		*out = new(string)
		**out = **in
	}
	if in.Pseudonym != nil {
		in, out := &in.Pseudonym, &out.Pseudonym
		*out = new(string)
		**out = **in
	}
	if in.SerialNumber != nil {
		in, out := &in.SerialNumber, &out.SerialNumber
		*out = new(string)
		**out = **in
	}
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(string)
		**out = **in
	}
	if in.Surname != nil {
		in, out := &in.Surname, &out.Surname
		*out = new(string)
		**out = **in
	}
	if in.Title != nil {
		in, out := &in.Title, &out.Title
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ASN1Subject.
func (in *ASN1Subject) DeepCopy() *ASN1Subject {
	if in == nil {
		return nil
	}
	out := new(ASN1Subject)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessDescription) DeepCopyInto(out *AccessDescription) {
	*out = *in
	if in.AccessLocation != nil {
		in, out := &in.AccessLocation, &out.AccessLocation
		*out = new(GeneralName)
		(*in).DeepCopyInto(*out)
	}
	if in.AccessMethod != nil {
		in, out := &in.AccessMethod, &out.AccessMethod
		*out = new(AccessMethod)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessDescription.
func (in *AccessDescription) DeepCopy() *AccessDescription {
	if in == nil {
		return nil
	}
	out := new(AccessDescription)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessMethod) DeepCopyInto(out *AccessMethod) {
	*out = *in
	if in.AccessMethodType != nil {
		in, out := &in.AccessMethodType, &out.AccessMethodType
		*out = new(string)
		**out = **in
	}
	if in.CustomObjectIdentifier != nil {
		in, out := &in.CustomObjectIdentifier, &out.CustomObjectIdentifier
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessMethod.
func (in *AccessMethod) DeepCopy() *AccessMethod {
	if in == nil {
		return nil
	}
	out := new(AccessMethod)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CRLConfiguration) DeepCopyInto(out *CRLConfiguration) {
	*out = *in
	if in.CustomCNAME != nil {
		in, out := &in.CustomCNAME, &out.CustomCNAME
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.ExpirationInDays != nil {
		in, out := &in.ExpirationInDays, &out.ExpirationInDays
		*out = new(int64)
		**out = **in
	}
	if in.S3BucketName != nil {
		in, out := &in.S3BucketName, &out.S3BucketName
		*out = new(string)
		**out = **in
	}
	if in.S3ObjectACL != nil {
		in, out := &in.S3ObjectACL, &out.S3ObjectACL
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CRLConfiguration.
func (in *CRLConfiguration) DeepCopy() *CRLConfiguration {
	if in == nil {
		return nil
	}
	out := new(CRLConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CSRExtensions) DeepCopyInto(out *CSRExtensions) {
	*out = *in
	if in.KeyUsage != nil {
		in, out := &in.KeyUsage, &out.KeyUsage
		*out = new(KeyUsage)
		(*in).DeepCopyInto(*out)
	}
	if in.SubjectInformationAccess != nil {
		in, out := &in.SubjectInformationAccess, &out.SubjectInformationAccess
		*out = make([]*AccessDescription, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(AccessDescription)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CSRExtensions.
func (in *CSRExtensions) DeepCopy() *CSRExtensions {
	if in == nil {
		return nil
	}
	out := new(CSRExtensions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Certificate) DeepCopyInto(out *Certificate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Certificate.
func (in *Certificate) DeepCopy() *Certificate {
	if in == nil {
		return nil
	}
	out := new(Certificate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Certificate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateAuthority) DeepCopyInto(out *CertificateAuthority) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateAuthority.
func (in *CertificateAuthority) DeepCopy() *CertificateAuthority {
	if in == nil {
		return nil
	}
	out := new(CertificateAuthority)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CertificateAuthority) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateAuthorityActivation) DeepCopyInto(out *CertificateAuthorityActivation) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateAuthorityActivation.
func (in *CertificateAuthorityActivation) DeepCopy() *CertificateAuthorityActivation {
	if in == nil {
		return nil
	}
	out := new(CertificateAuthorityActivation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CertificateAuthorityActivation) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateAuthorityActivationList) DeepCopyInto(out *CertificateAuthorityActivationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CertificateAuthorityActivation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateAuthorityActivationList.
func (in *CertificateAuthorityActivationList) DeepCopy() *CertificateAuthorityActivationList {
	if in == nil {
		return nil
	}
	out := new(CertificateAuthorityActivationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CertificateAuthorityActivationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateAuthorityActivationSpec) DeepCopyInto(out *CertificateAuthorityActivationSpec) {
	*out = *in
	if in.Certificate != nil {
		in, out := &in.Certificate, &out.Certificate
		*out = new(corev1alpha1.SecretKeyReference)
		**out = **in
	}
	if in.CertificateAuthorityARN != nil {
		in, out := &in.CertificateAuthorityARN, &out.CertificateAuthorityARN
		*out = new(string)
		**out = **in
	}
	if in.CertificateAuthorityRef != nil {
		in, out := &in.CertificateAuthorityRef, &out.CertificateAuthorityRef
		*out = new(corev1alpha1.AWSResourceReferenceWrapper)
		(*in).DeepCopyInto(*out)
	}
	if in.CertificateChain != nil {
		in, out := &in.CertificateChain, &out.CertificateChain
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.RevocationConfiguration != nil {
		in, out := &in.RevocationConfiguration, &out.RevocationConfiguration
		*out = new(RevocationConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateAuthorityActivationSpec.
func (in *CertificateAuthorityActivationSpec) DeepCopy() *CertificateAuthorityActivationSpec {
	if in == nil {
		return nil
	}
	out := new(CertificateAuthorityActivationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateAuthorityActivationStatus) DeepCopyInto(out *CertificateAuthorityActivationStatus) {
	*out = *in
	if in.ACKResourceMetadata != nil {
		in, out := &in.ACKResourceMetadata, &out.ACKResourceMetadata
		*out = new(corev1alpha1.ResourceMetadata)
		(*in).DeepCopyInto(*out)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]*corev1alpha1.Condition, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(corev1alpha1.Condition)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateAuthorityActivationStatus.
func (in *CertificateAuthorityActivationStatus) DeepCopy() *CertificateAuthorityActivationStatus {
	if in == nil {
		return nil
	}
	out := new(CertificateAuthorityActivationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateAuthorityConfiguration) DeepCopyInto(out *CertificateAuthorityConfiguration) {
	*out = *in
	if in.CSRExtensions != nil {
		in, out := &in.CSRExtensions, &out.CSRExtensions
		*out = new(CSRExtensions)
		(*in).DeepCopyInto(*out)
	}
	if in.KeyAlgorithm != nil {
		in, out := &in.KeyAlgorithm, &out.KeyAlgorithm
		*out = new(string)
		**out = **in
	}
	if in.SigningAlgorithm != nil {
		in, out := &in.SigningAlgorithm, &out.SigningAlgorithm
		*out = new(string)
		**out = **in
	}
	if in.Subject != nil {
		in, out := &in.Subject, &out.Subject
		*out = new(ASN1Subject)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateAuthorityConfiguration.
func (in *CertificateAuthorityConfiguration) DeepCopy() *CertificateAuthorityConfiguration {
	if in == nil {
		return nil
	}
	out := new(CertificateAuthorityConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateAuthorityList) DeepCopyInto(out *CertificateAuthorityList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CertificateAuthority, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateAuthorityList.
func (in *CertificateAuthorityList) DeepCopy() *CertificateAuthorityList {
	if in == nil {
		return nil
	}
	out := new(CertificateAuthorityList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CertificateAuthorityList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateAuthoritySpec) DeepCopyInto(out *CertificateAuthoritySpec) {
	*out = *in
	if in.CertificateAuthorityConfiguration != nil {
		in, out := &in.CertificateAuthorityConfiguration, &out.CertificateAuthorityConfiguration
		*out = new(CertificateAuthorityConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.CertificateAuthorityType != nil {
		in, out := &in.CertificateAuthorityType, &out.CertificateAuthorityType
		*out = new(string)
		**out = **in
	}
	if in.KeyStorageSecurityStandard != nil {
		in, out := &in.KeyStorageSecurityStandard, &out.KeyStorageSecurityStandard
		*out = new(string)
		**out = **in
	}
	if in.RevocationConfiguration != nil {
		in, out := &in.RevocationConfiguration, &out.RevocationConfiguration
		*out = new(RevocationConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]*Tag, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Tag)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.UsageMode != nil {
		in, out := &in.UsageMode, &out.UsageMode
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateAuthoritySpec.
func (in *CertificateAuthoritySpec) DeepCopy() *CertificateAuthoritySpec {
	if in == nil {
		return nil
	}
	out := new(CertificateAuthoritySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateAuthorityStatus) DeepCopyInto(out *CertificateAuthorityStatus) {
	*out = *in
	if in.ACKResourceMetadata != nil {
		in, out := &in.ACKResourceMetadata, &out.ACKResourceMetadata
		*out = new(corev1alpha1.ResourceMetadata)
		(*in).DeepCopyInto(*out)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]*corev1alpha1.Condition, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(corev1alpha1.Condition)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.CSR != nil {
		in, out := &in.CSR, &out.CSR
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateAuthorityStatus.
func (in *CertificateAuthorityStatus) DeepCopy() *CertificateAuthorityStatus {
	if in == nil {
		return nil
	}
	out := new(CertificateAuthorityStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateAuthority_SDK) DeepCopyInto(out *CertificateAuthority_SDK) {
	*out = *in
	if in.ARN != nil {
		in, out := &in.ARN, &out.ARN
		*out = new(string)
		**out = **in
	}
	if in.CertificateAuthorityConfiguration != nil {
		in, out := &in.CertificateAuthorityConfiguration, &out.CertificateAuthorityConfiguration
		*out = new(CertificateAuthorityConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.CreatedAt != nil {
		in, out := &in.CreatedAt, &out.CreatedAt
		*out = (*in).DeepCopy()
	}
	if in.FailureReason != nil {
		in, out := &in.FailureReason, &out.FailureReason
		*out = new(string)
		**out = **in
	}
	if in.KeyStorageSecurityStandard != nil {
		in, out := &in.KeyStorageSecurityStandard, &out.KeyStorageSecurityStandard
		*out = new(string)
		**out = **in
	}
	if in.LastStateChangeAt != nil {
		in, out := &in.LastStateChangeAt, &out.LastStateChangeAt
		*out = (*in).DeepCopy()
	}
	if in.NotAfter != nil {
		in, out := &in.NotAfter, &out.NotAfter
		*out = (*in).DeepCopy()
	}
	if in.NotBefore != nil {
		in, out := &in.NotBefore, &out.NotBefore
		*out = (*in).DeepCopy()
	}
	if in.OwnerAccount != nil {
		in, out := &in.OwnerAccount, &out.OwnerAccount
		*out = new(string)
		**out = **in
	}
	if in.RestorableUntil != nil {
		in, out := &in.RestorableUntil, &out.RestorableUntil
		*out = (*in).DeepCopy()
	}
	if in.RevocationConfiguration != nil {
		in, out := &in.RevocationConfiguration, &out.RevocationConfiguration
		*out = new(RevocationConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.Serial != nil {
		in, out := &in.Serial, &out.Serial
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.UsageMode != nil {
		in, out := &in.UsageMode, &out.UsageMode
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateAuthority_SDK.
func (in *CertificateAuthority_SDK) DeepCopy() *CertificateAuthority_SDK {
	if in == nil {
		return nil
	}
	out := new(CertificateAuthority_SDK)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateList) DeepCopyInto(out *CertificateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Certificate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateList.
func (in *CertificateList) DeepCopy() *CertificateList {
	if in == nil {
		return nil
	}
	out := new(CertificateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CertificateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateSpec) DeepCopyInto(out *CertificateSpec) {
	*out = *in
	if in.APIPassthrough != nil {
		in, out := &in.APIPassthrough, &out.APIPassthrough
		*out = new(APIPassthrough)
		(*in).DeepCopyInto(*out)
	}
	if in.CertificateAuthorityARN != nil {
		in, out := &in.CertificateAuthorityARN, &out.CertificateAuthorityARN
		*out = new(string)
		**out = **in
	}
	if in.CertificateAuthorityRef != nil {
		in, out := &in.CertificateAuthorityRef, &out.CertificateAuthorityRef
		*out = new(corev1alpha1.AWSResourceReferenceWrapper)
		(*in).DeepCopyInto(*out)
	}
	if in.CSR != nil {
		in, out := &in.CSR, &out.CSR
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.CSRRef != nil {
		in, out := &in.CSRRef, &out.CSRRef
		*out = new(corev1alpha1.AWSResourceReferenceWrapper)
		(*in).DeepCopyInto(*out)
	}
	if in.SigningAlgorithm != nil {
		in, out := &in.SigningAlgorithm, &out.SigningAlgorithm
		*out = new(string)
		**out = **in
	}
	if in.TemplateARN != nil {
		in, out := &in.TemplateARN, &out.TemplateARN
		*out = new(string)
		**out = **in
	}
	if in.Validity != nil {
		in, out := &in.Validity, &out.Validity
		*out = new(Validity)
		(*in).DeepCopyInto(*out)
	}
	if in.ValidityNotBefore != nil {
		in, out := &in.ValidityNotBefore, &out.ValidityNotBefore
		*out = new(Validity)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateSpec.
func (in *CertificateSpec) DeepCopy() *CertificateSpec {
	if in == nil {
		return nil
	}
	out := new(CertificateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateStatus) DeepCopyInto(out *CertificateStatus) {
	*out = *in
	if in.ACKResourceMetadata != nil {
		in, out := &in.ACKResourceMetadata, &out.ACKResourceMetadata
		*out = new(corev1alpha1.ResourceMetadata)
		(*in).DeepCopyInto(*out)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]*corev1alpha1.Condition, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(corev1alpha1.Condition)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateStatus.
func (in *CertificateStatus) DeepCopy() *CertificateStatus {
	if in == nil {
		return nil
	}
	out := new(CertificateStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomAttribute) DeepCopyInto(out *CustomAttribute) {
	*out = *in
	if in.ObjectIdentifier != nil {
		in, out := &in.ObjectIdentifier, &out.ObjectIdentifier
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomAttribute.
func (in *CustomAttribute) DeepCopy() *CustomAttribute {
	if in == nil {
		return nil
	}
	out := new(CustomAttribute)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomExtension) DeepCopyInto(out *CustomExtension) {
	*out = *in
	if in.Critical != nil {
		in, out := &in.Critical, &out.Critical
		*out = new(bool)
		**out = **in
	}
	if in.ObjectIdentifier != nil {
		in, out := &in.ObjectIdentifier, &out.ObjectIdentifier
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomExtension.
func (in *CustomExtension) DeepCopy() *CustomExtension {
	if in == nil {
		return nil
	}
	out := new(CustomExtension)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EDIPartyName) DeepCopyInto(out *EDIPartyName) {
	*out = *in
	if in.NameAssigner != nil {
		in, out := &in.NameAssigner, &out.NameAssigner
		*out = new(string)
		**out = **in
	}
	if in.PartyName != nil {
		in, out := &in.PartyName, &out.PartyName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EDIPartyName.
func (in *EDIPartyName) DeepCopy() *EDIPartyName {
	if in == nil {
		return nil
	}
	out := new(EDIPartyName)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExtendedKeyUsage) DeepCopyInto(out *ExtendedKeyUsage) {
	*out = *in
	if in.ExtendedKeyUsageObjectIdentifier != nil {
		in, out := &in.ExtendedKeyUsageObjectIdentifier, &out.ExtendedKeyUsageObjectIdentifier
		*out = new(string)
		**out = **in
	}
	if in.ExtendedKeyUsageType != nil {
		in, out := &in.ExtendedKeyUsageType, &out.ExtendedKeyUsageType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExtendedKeyUsage.
func (in *ExtendedKeyUsage) DeepCopy() *ExtendedKeyUsage {
	if in == nil {
		return nil
	}
	out := new(ExtendedKeyUsage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Extensions) DeepCopyInto(out *Extensions) {
	*out = *in
	if in.CertificatePolicies != nil {
		in, out := &in.CertificatePolicies, &out.CertificatePolicies
		*out = make([]*PolicyInformation, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(PolicyInformation)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.CustomExtensions != nil {
		in, out := &in.CustomExtensions, &out.CustomExtensions
		*out = make([]*CustomExtension, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(CustomExtension)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.ExtendedKeyUsage != nil {
		in, out := &in.ExtendedKeyUsage, &out.ExtendedKeyUsage
		*out = make([]*ExtendedKeyUsage, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ExtendedKeyUsage)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.KeyUsage != nil {
		in, out := &in.KeyUsage, &out.KeyUsage
		*out = new(KeyUsage)
		(*in).DeepCopyInto(*out)
	}
	if in.SubjectAlternativeNames != nil {
		in, out := &in.SubjectAlternativeNames, &out.SubjectAlternativeNames
		*out = make([]*GeneralName, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(GeneralName)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Extensions.
func (in *Extensions) DeepCopy() *Extensions {
	if in == nil {
		return nil
	}
	out := new(Extensions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GeneralName) DeepCopyInto(out *GeneralName) {
	*out = *in
	if in.DirectoryName != nil {
		in, out := &in.DirectoryName, &out.DirectoryName
		*out = new(ASN1Subject)
		(*in).DeepCopyInto(*out)
	}
	if in.DNSName != nil {
		in, out := &in.DNSName, &out.DNSName
		*out = new(string)
		**out = **in
	}
	if in.EDIPartyName != nil {
		in, out := &in.EDIPartyName, &out.EDIPartyName
		*out = new(EDIPartyName)
		(*in).DeepCopyInto(*out)
	}
	if in.IPAddress != nil {
		in, out := &in.IPAddress, &out.IPAddress
		*out = new(string)
		**out = **in
	}
	if in.OtherName != nil {
		in, out := &in.OtherName, &out.OtherName
		*out = new(OtherName)
		(*in).DeepCopyInto(*out)
	}
	if in.RegisteredID != nil {
		in, out := &in.RegisteredID, &out.RegisteredID
		*out = new(string)
		**out = **in
	}
	if in.RFC822Name != nil {
		in, out := &in.RFC822Name, &out.RFC822Name
		*out = new(string)
		**out = **in
	}
	if in.UniformResourceIdentifier != nil {
		in, out := &in.UniformResourceIdentifier, &out.UniformResourceIdentifier
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GeneralName.
func (in *GeneralName) DeepCopy() *GeneralName {
	if in == nil {
		return nil
	}
	out := new(GeneralName)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeyUsage) DeepCopyInto(out *KeyUsage) {
	*out = *in
	if in.CRLSign != nil {
		in, out := &in.CRLSign, &out.CRLSign
		*out = new(bool)
		**out = **in
	}
	if in.DataEncipherment != nil {
		in, out := &in.DataEncipherment, &out.DataEncipherment
		*out = new(bool)
		**out = **in
	}
	if in.DecipherOnly != nil {
		in, out := &in.DecipherOnly, &out.DecipherOnly
		*out = new(bool)
		**out = **in
	}
	if in.DigitalSignature != nil {
		in, out := &in.DigitalSignature, &out.DigitalSignature
		*out = new(bool)
		**out = **in
	}
	if in.EncipherOnly != nil {
		in, out := &in.EncipherOnly, &out.EncipherOnly
		*out = new(bool)
		**out = **in
	}
	if in.KeyAgreement != nil {
		in, out := &in.KeyAgreement, &out.KeyAgreement
		*out = new(bool)
		**out = **in
	}
	if in.KeyCertSign != nil {
		in, out := &in.KeyCertSign, &out.KeyCertSign
		*out = new(bool)
		**out = **in
	}
	if in.KeyEncipherment != nil {
		in, out := &in.KeyEncipherment, &out.KeyEncipherment
		*out = new(bool)
		**out = **in
	}
	if in.NonRepudiation != nil {
		in, out := &in.NonRepudiation, &out.NonRepudiation
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeyUsage.
func (in *KeyUsage) DeepCopy() *KeyUsage {
	if in == nil {
		return nil
	}
	out := new(KeyUsage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OCSPConfiguration) DeepCopyInto(out *OCSPConfiguration) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.OCSPCustomCNAME != nil {
		in, out := &in.OCSPCustomCNAME, &out.OCSPCustomCNAME
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OCSPConfiguration.
func (in *OCSPConfiguration) DeepCopy() *OCSPConfiguration {
	if in == nil {
		return nil
	}
	out := new(OCSPConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OtherName) DeepCopyInto(out *OtherName) {
	*out = *in
	if in.TypeID != nil {
		in, out := &in.TypeID, &out.TypeID
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OtherName.
func (in *OtherName) DeepCopy() *OtherName {
	if in == nil {
		return nil
	}
	out := new(OtherName)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Permission) DeepCopyInto(out *Permission) {
	*out = *in
	if in.CertificateAuthorityARN != nil {
		in, out := &in.CertificateAuthorityARN, &out.CertificateAuthorityARN
		*out = new(string)
		**out = **in
	}
	if in.CreatedAt != nil {
		in, out := &in.CreatedAt, &out.CreatedAt
		*out = (*in).DeepCopy()
	}
	if in.SourceAccount != nil {
		in, out := &in.SourceAccount, &out.SourceAccount
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Permission.
func (in *Permission) DeepCopy() *Permission {
	if in == nil {
		return nil
	}
	out := new(Permission)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyInformation) DeepCopyInto(out *PolicyInformation) {
	*out = *in
	if in.CertPolicyID != nil {
		in, out := &in.CertPolicyID, &out.CertPolicyID
		*out = new(string)
		**out = **in
	}
	if in.PolicyQualifiers != nil {
		in, out := &in.PolicyQualifiers, &out.PolicyQualifiers
		*out = make([]*PolicyQualifierInfo, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(PolicyQualifierInfo)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyInformation.
func (in *PolicyInformation) DeepCopy() *PolicyInformation {
	if in == nil {
		return nil
	}
	out := new(PolicyInformation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyQualifierInfo) DeepCopyInto(out *PolicyQualifierInfo) {
	*out = *in
	if in.PolicyQualifierID != nil {
		in, out := &in.PolicyQualifierID, &out.PolicyQualifierID
		*out = new(string)
		**out = **in
	}
	if in.Qualifier != nil {
		in, out := &in.Qualifier, &out.Qualifier
		*out = new(Qualifier)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyQualifierInfo.
func (in *PolicyQualifierInfo) DeepCopy() *PolicyQualifierInfo {
	if in == nil {
		return nil
	}
	out := new(PolicyQualifierInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Qualifier) DeepCopyInto(out *Qualifier) {
	*out = *in
	if in.CPSURI != nil {
		in, out := &in.CPSURI, &out.CPSURI
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Qualifier.
func (in *Qualifier) DeepCopy() *Qualifier {
	if in == nil {
		return nil
	}
	out := new(Qualifier)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RevocationConfiguration) DeepCopyInto(out *RevocationConfiguration) {
	*out = *in
	if in.CRLConfiguration != nil {
		in, out := &in.CRLConfiguration, &out.CRLConfiguration
		*out = new(CRLConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.OCSPConfiguration != nil {
		in, out := &in.OCSPConfiguration, &out.OCSPConfiguration
		*out = new(OCSPConfiguration)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RevocationConfiguration.
func (in *RevocationConfiguration) DeepCopy() *RevocationConfiguration {
	if in == nil {
		return nil
	}
	out := new(RevocationConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Tag) DeepCopyInto(out *Tag) {
	*out = *in
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Tag.
func (in *Tag) DeepCopy() *Tag {
	if in == nil {
		return nil
	}
	out := new(Tag)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Validity) DeepCopyInto(out *Validity) {
	*out = *in
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Validity.
func (in *Validity) DeepCopy() *Validity {
	if in == nil {
		return nil
	}
	out := new(Validity)
	in.DeepCopyInto(out)
	return out
}
