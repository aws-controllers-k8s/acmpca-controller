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

package v1alpha1

import (
	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CertificateAuthoritySpec defines the desired state of CertificateAuthority.
//
// Contains information about your private certificate authority (CA). Your
// private CA can issue and revoke X.509 digital certificates. Digital certificates
// verify that the entity named in the certificate Subject field owns or controls
// the public key contained in the Subject Public Key Info field. Call the CreateCertificateAuthority
// (https://docs.aws.amazon.com/privateca/latest/APIReference/API_CreateCertificateAuthority.html)
// action to create your private CA. You must then call the GetCertificateAuthorityCertificate
// (https://docs.aws.amazon.com/privateca/latest/APIReference/API_GetCertificateAuthorityCertificate.html)
// action to retrieve a private CA certificate signing request (CSR). Sign the
// CSR with your Amazon Web Services Private CA-hosted or on-premises root or
// subordinate CA certificate. Call the ImportCertificateAuthorityCertificate
// (https://docs.aws.amazon.com/privateca/latest/APIReference/API_ImportCertificateAuthorityCertificate.html)
// action to import the signed certificate into Certificate Manager (ACM).
type CertificateAuthoritySpec struct {

	// Name and bit size of the private key algorithm, the name of the signing algorithm,
	// and X.500 certificate subject information.
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Value is immutable once set"
	// +kubebuilder:validation:Required
	CertificateAuthorityConfiguration *CertificateAuthorityConfiguration `json:"certificateAuthorityConfiguration"`
	// Specifies a cryptographic key management compliance standard used for handling
	// CA keys.
	//
	// Default: FIPS_140_2_LEVEL_3_OR_HIGHER
	//
	// Some Amazon Web Services Regions do not support the default. When creating
	// a CA in these Regions, you must provide FIPS_140_2_LEVEL_2_OR_HIGHER as the
	// argument for KeyStorageSecurityStandard. Failure to do this results in an
	// InvalidArgsException with the message, "A certificate authority cannot be
	// created in this region with the specified security standard."
	//
	// For information about security standard support in various Regions, see Storage
	// and security compliance of Amazon Web Services Private CA private keys (https://docs.aws.amazon.com/privateca/latest/userguide/data-protection.html#private-keys).
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Value is immutable once set"
	KeyStorageSecurityStandard *string `json:"keyStorageSecurityStandard,omitempty"`
	// Contains information to enable support for Online Certificate Status Protocol
	// (OCSP), certificate revocation list (CRL), both protocols, or neither. By
	// default, both certificate validation mechanisms are disabled.
	//
	// The following requirements apply to revocation configurations.
	//
	//   - A configuration disabling CRLs or OCSP must contain only the Enabled=False
	//     parameter, and will fail if other parameters such as CustomCname or ExpirationInDays
	//     are included.
	//
	//   - In a CRL configuration, the S3BucketName parameter must conform to Amazon
	//     S3 bucket naming rules (https://docs.aws.amazon.com/AmazonS3/latest/userguide/bucketnamingrules.html).
	//
	//   - A configuration containing a custom Canonical Name (CNAME) parameter
	//     for CRLs or OCSP must conform to RFC2396 (https://www.ietf.org/rfc/rfc2396.txt)
	//     restrictions on the use of special characters in a CNAME.
	//
	//   - In a CRL or OCSP configuration, the value of a CNAME parameter must
	//     not include a protocol prefix such as "http://" or "https://".
	//
	// For more information, see the OcspConfiguration (https://docs.aws.amazon.com/privateca/latest/APIReference/API_OcspConfiguration.html)
	// and CrlConfiguration (https://docs.aws.amazon.com/privateca/latest/APIReference/API_CrlConfiguration.html)
	// types.
	RevocationConfiguration *RevocationConfiguration `json:"revocationConfiguration,omitempty"`
	// Key-value pairs that will be attached to the new private CA. You can associate
	// up to 50 tags with a private CA. For information using tags with IAM to manage
	// permissions, see Controlling Access Using IAM Tags (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_iam-tags.html).
	Tags []*Tag `json:"tags,omitempty"`
	// The type of the certificate authority.
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Value is immutable once set"
	// +kubebuilder:validation:Required
	Type *string `json:"type,omitempty"`
	// Specifies whether the CA issues general-purpose certificates that typically
	// require a revocation mechanism, or short-lived certificates that may optionally
	// omit revocation because they expire quickly. Short-lived certificate validity
	// is limited to seven days.
	//
	// The default value is GENERAL_PURPOSE.
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Value is immutable once set"
	UsageMode *string `json:"usageMode,omitempty"`
}

// CertificateAuthorityStatus defines the observed state of CertificateAuthority
type CertificateAuthorityStatus struct {
	// All CRs managed by ACK have a common `Status.ACKResourceMetadata` member
	// that is used to contain resource sync state, account ownership,
	// constructed ARN for the resource
	// +kubebuilder:validation:Optional
	ACKResourceMetadata *ackv1alpha1.ResourceMetadata `json:"ackResourceMetadata"`
	// All CRs managed by ACK have a common `Status.Conditions` member that
	// contains a collection of `ackv1alpha1.Condition` objects that describe
	// the various terminal states of the CR and its backend AWS service API
	// resource
	// +kubebuilder:validation:Optional
	Conditions []*ackv1alpha1.Condition `json:"conditions"`
	// The base64 PEM-encoded certificate signing request (CSR) for your private
	// CA certificate.
	// +kubebuilder:validation:Optional
	CertificateSigningRequest *string `json:"certificateSigningRequest,omitempty"`
	// Date and time at which your private CA was created.
	// +kubebuilder:validation:Optional
	CreatedAt *metav1.Time `json:"createdAt,omitempty"`
	// Reason the request to create your private CA failed.
	// +kubebuilder:validation:Optional
	FailureReason *string `json:"failureReason,omitempty"`
	// Date and time at which your private CA was last updated.
	// +kubebuilder:validation:Optional
	LastStateChangeAt *metav1.Time `json:"lastStateChangeAt,omitempty"`
	// Date and time after which your private CA certificate is not valid.
	// +kubebuilder:validation:Optional
	NotAfter *metav1.Time `json:"notAfter,omitempty"`
	// Date and time before which your private CA certificate is not valid.
	// +kubebuilder:validation:Optional
	NotBefore *metav1.Time `json:"notBefore,omitempty"`
	// The Amazon Web Services account ID that owns the certificate authority.
	//
	// Regex Pattern: `^[0-9]+$`
	// +kubebuilder:validation:Optional
	OwnerAccount *string `json:"ownerAccount,omitempty"`
	// The period during which a deleted CA can be restored. For more information,
	// see the PermanentDeletionTimeInDays parameter of the DeleteCertificateAuthorityRequest
	// (https://docs.aws.amazon.com/privateca/latest/APIReference/API_DeleteCertificateAuthorityRequest.html)
	// action.
	// +kubebuilder:validation:Optional
	RestorableUntil *metav1.Time `json:"restorableUntil,omitempty"`
	// Serial number of your private CA.
	// +kubebuilder:validation:Optional
	Serial *string `json:"serial,omitempty"`
	// Status of your private CA.
	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty"`
}

// CertificateAuthority is the Schema for the CertificateAuthorities API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type CertificateAuthority struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CertificateAuthoritySpec   `json:"spec,omitempty"`
	Status            CertificateAuthorityStatus `json:"status,omitempty"`
}

// CertificateAuthorityList contains a list of CertificateAuthority
// +kubebuilder:object:root=true
type CertificateAuthorityList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CertificateAuthority `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CertificateAuthority{}, &CertificateAuthorityList{})
}
