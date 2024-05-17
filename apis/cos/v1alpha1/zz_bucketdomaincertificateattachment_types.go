// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type BucketDomainCertificateAttachmentInitParameters struct {

	// The certificate of specified doamin.
	// The certificate of specified doamin.
	DomainCertificate []DomainCertificateInitParameters `json:"domainCertificate,omitempty" tf:"domain_certificate,omitempty"`
}

type BucketDomainCertificateAttachmentObservation struct {

	// Bucket name.
	// Bucket name.
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// The certificate of specified doamin.
	// The certificate of specified doamin.
	DomainCertificate []DomainCertificateObservation `json:"domainCertificate,omitempty" tf:"domain_certificate,omitempty"`

	// ID of the resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type BucketDomainCertificateAttachmentParameters struct {

	// Bucket name.
	// Bucket name.
	// +crossplane:generate:reference:type=Bucket
	// +kubebuilder:validation:Optional
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Reference to a Bucket to populate bucket.
	// +kubebuilder:validation:Optional
	BucketRef *v1.Reference `json:"bucketRef,omitempty" tf:"-"`

	// Selector for a Bucket to populate bucket.
	// +kubebuilder:validation:Optional
	BucketSelector *v1.Selector `json:"bucketSelector,omitempty" tf:"-"`

	// The certificate of specified doamin.
	// The certificate of specified doamin.
	// +kubebuilder:validation:Optional
	DomainCertificate []DomainCertificateParameters `json:"domainCertificate,omitempty" tf:"domain_certificate,omitempty"`
}

type CertificateInitParameters struct {

	// Certificate type.
	// Certificate type.
	CertType *string `json:"certType,omitempty" tf:"cert_type,omitempty"`

	// Custom certificate.
	// Custom certificate.
	CustomCert []CustomCertInitParameters `json:"customCert,omitempty" tf:"custom_cert,omitempty"`
}

type CertificateObservation struct {

	// Certificate type.
	// Certificate type.
	CertType *string `json:"certType,omitempty" tf:"cert_type,omitempty"`

	// Custom certificate.
	// Custom certificate.
	CustomCert []CustomCertObservation `json:"customCert,omitempty" tf:"custom_cert,omitempty"`
}

type CertificateParameters struct {

	// Certificate type.
	// Certificate type.
	// +kubebuilder:validation:Optional
	CertType *string `json:"certType" tf:"cert_type,omitempty"`

	// Custom certificate.
	// Custom certificate.
	// +kubebuilder:validation:Optional
	CustomCert []CustomCertParameters `json:"customCert" tf:"custom_cert,omitempty"`
}

type CustomCertInitParameters struct {

	// Public key of certificate.
	// Public key of certificate.
	Cert *string `json:"cert,omitempty" tf:"cert,omitempty"`

	// Private key of certificate.
	// Private key of certificate.
	PrivateKey *string `json:"privateKey,omitempty" tf:"private_key,omitempty"`
}

type CustomCertObservation struct {

	// Public key of certificate.
	// Public key of certificate.
	Cert *string `json:"cert,omitempty" tf:"cert,omitempty"`

	// Private key of certificate.
	// Private key of certificate.
	PrivateKey *string `json:"privateKey,omitempty" tf:"private_key,omitempty"`
}

type CustomCertParameters struct {

	// Public key of certificate.
	// Public key of certificate.
	// +kubebuilder:validation:Optional
	Cert *string `json:"cert" tf:"cert,omitempty"`

	// Private key of certificate.
	// Private key of certificate.
	// +kubebuilder:validation:Optional
	PrivateKey *string `json:"privateKey" tf:"private_key,omitempty"`
}

type DomainCertificateInitParameters struct {

	// Certificate info.
	// Certificate info.
	Certificate []CertificateInitParameters `json:"certificate,omitempty" tf:"certificate,omitempty"`

	// The name of domain.
	// The name of domain.
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`
}

type DomainCertificateObservation struct {

	// Certificate info.
	// Certificate info.
	Certificate []CertificateObservation `json:"certificate,omitempty" tf:"certificate,omitempty"`

	// The name of domain.
	// The name of domain.
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`
}

type DomainCertificateParameters struct {

	// Certificate info.
	// Certificate info.
	// +kubebuilder:validation:Optional
	Certificate []CertificateParameters `json:"certificate" tf:"certificate,omitempty"`

	// The name of domain.
	// The name of domain.
	// +kubebuilder:validation:Optional
	Domain *string `json:"domain" tf:"domain,omitempty"`
}

// BucketDomainCertificateAttachmentSpec defines the desired state of BucketDomainCertificateAttachment
type BucketDomainCertificateAttachmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BucketDomainCertificateAttachmentParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider BucketDomainCertificateAttachmentInitParameters `json:"initProvider,omitempty"`
}

// BucketDomainCertificateAttachmentStatus defines the observed state of BucketDomainCertificateAttachment.
type BucketDomainCertificateAttachmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BucketDomainCertificateAttachmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BucketDomainCertificateAttachment is the Schema for the BucketDomainCertificateAttachments API. Provides a resource to attach/detach the corresponding certificate for the domain name in specified cos bucket.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloud}
type BucketDomainCertificateAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.domainCertificate) || (has(self.initProvider) && has(self.initProvider.domainCertificate))",message="spec.forProvider.domainCertificate is a required parameter"
	Spec   BucketDomainCertificateAttachmentSpec   `json:"spec"`
	Status BucketDomainCertificateAttachmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BucketDomainCertificateAttachmentList contains a list of BucketDomainCertificateAttachments
type BucketDomainCertificateAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BucketDomainCertificateAttachment `json:"items"`
}

// Repository type metadata.
var (
	BucketDomainCertificateAttachment_Kind             = "BucketDomainCertificateAttachment"
	BucketDomainCertificateAttachment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BucketDomainCertificateAttachment_Kind}.String()
	BucketDomainCertificateAttachment_KindAPIVersion   = BucketDomainCertificateAttachment_Kind + "." + CRDGroupVersion.String()
	BucketDomainCertificateAttachment_GroupVersionKind = CRDGroupVersion.WithKind(BucketDomainCertificateAttachment_Kind)
)

func init() {
	SchemeBuilder.Register(&BucketDomainCertificateAttachment{}, &BucketDomainCertificateAttachmentList{})
}