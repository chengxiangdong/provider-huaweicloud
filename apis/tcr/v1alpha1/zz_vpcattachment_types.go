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

type VpcAttachmentInitParameters struct {

	// Whether to enable public domain dns. Default value is false.
	// Whether to enable public domain dns. Default value is `false`.
	EnablePublicDomainDNS *bool `json:"enablePublicDomainDns,omitempty" tf:"enable_public_domain_dns,omitempty"`

	// Whether to enable vpc domain dns. Default value is false.
	// Whether to enable vpc domain dns. Default value is `false`.
	EnableVPCDomainDNS *bool `json:"enableVpcDomainDns,omitempty" tf:"enable_vpc_domain_dns,omitempty"`

	// this argument was deprecated, use region_name instead. ID of region. Conflict with region_name, can not be set at the same time.
	// ID of region. Conflict with region_name, can not be set at the same time.
	RegionID *float64 `json:"regionId,omitempty" tf:"region_id,omitempty"`

	// Name of region. Conflict with region_id, can not be set at the same time.
	// Name of region. Conflict with region_id, can not be set at the same time.
	RegionName *string `json:"regionName,omitempty" tf:"region_name,omitempty"`
}

type VpcAttachmentObservation struct {

	// IP address of the internal access.
	// IP address of the internal access.
	AccessIP *string `json:"accessIp,omitempty" tf:"access_ip,omitempty"`

	// Whether to enable public domain dns. Default value is false.
	// Whether to enable public domain dns. Default value is `false`.
	EnablePublicDomainDNS *bool `json:"enablePublicDomainDns,omitempty" tf:"enable_public_domain_dns,omitempty"`

	// Whether to enable vpc domain dns. Default value is false.
	// Whether to enable vpc domain dns. Default value is `false`.
	EnableVPCDomainDNS *bool `json:"enableVpcDomainDns,omitempty" tf:"enable_vpc_domain_dns,omitempty"`

	// ID of the resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// ID of the TCR instance.
	// ID of the TCR instance.
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// this argument was deprecated, use region_name instead. ID of region. Conflict with region_name, can not be set at the same time.
	// ID of region. Conflict with region_name, can not be set at the same time.
	RegionID *float64 `json:"regionId,omitempty" tf:"region_id,omitempty"`

	// Name of region. Conflict with region_id, can not be set at the same time.
	// Name of region. Conflict with region_id, can not be set at the same time.
	RegionName *string `json:"regionName,omitempty" tf:"region_name,omitempty"`

	// Status of the internal access.
	// Status of the internal access.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// ID of subnet.
	// ID of subnet.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// ID of VPC.
	// ID of VPC.
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`
}

type VpcAttachmentParameters struct {

	// Whether to enable public domain dns. Default value is false.
	// Whether to enable public domain dns. Default value is `false`.
	// +kubebuilder:validation:Optional
	EnablePublicDomainDNS *bool `json:"enablePublicDomainDns,omitempty" tf:"enable_public_domain_dns,omitempty"`

	// Whether to enable vpc domain dns. Default value is false.
	// Whether to enable vpc domain dns. Default value is `false`.
	// +kubebuilder:validation:Optional
	EnableVPCDomainDNS *bool `json:"enableVpcDomainDns,omitempty" tf:"enable_vpc_domain_dns,omitempty"`

	// ID of the TCR instance.
	// ID of the TCR instance.
	// +crossplane:generate:reference:type=Instance
	// +kubebuilder:validation:Optional
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Reference to a Instance to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDRef *v1.Reference `json:"instanceIdRef,omitempty" tf:"-"`

	// Selector for a Instance to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDSelector *v1.Selector `json:"instanceIdSelector,omitempty" tf:"-"`

	// this argument was deprecated, use region_name instead. ID of region. Conflict with region_name, can not be set at the same time.
	// ID of region. Conflict with region_name, can not be set at the same time.
	// +kubebuilder:validation:Optional
	RegionID *float64 `json:"regionId,omitempty" tf:"region_id,omitempty"`

	// Name of region. Conflict with region_id, can not be set at the same time.
	// Name of region. Conflict with region_id, can not be set at the same time.
	// +kubebuilder:validation:Optional
	RegionName *string `json:"regionName,omitempty" tf:"region_name,omitempty"`

	// ID of subnet.
	// ID of subnet.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1.Subnet
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// ID of VPC.
	// ID of VPC.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1.VPC
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Reference to a VPC in vpc to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// Selector for a VPC in vpc to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`
}

// VpcAttachmentSpec defines the desired state of VpcAttachment
type VpcAttachmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VpcAttachmentParameters `json:"forProvider"`
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
	InitProvider VpcAttachmentInitParameters `json:"initProvider,omitempty"`
}

// VpcAttachmentStatus defines the observed state of VpcAttachment.
type VpcAttachmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VpcAttachmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VpcAttachment is the Schema for the VpcAttachments API. Use this resource to attach tcr instance with the vpc and subnet network.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloud}
type VpcAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VpcAttachmentSpec   `json:"spec"`
	Status            VpcAttachmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VpcAttachmentList contains a list of VpcAttachments
type VpcAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VpcAttachment `json:"items"`
}

// Repository type metadata.
var (
	VpcAttachment_Kind             = "VpcAttachment"
	VpcAttachment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VpcAttachment_Kind}.String()
	VpcAttachment_KindAPIVersion   = VpcAttachment_Kind + "." + CRDGroupVersion.String()
	VpcAttachment_GroupVersionKind = CRDGroupVersion.WithKind(VpcAttachment_Kind)
)

func init() {
	SchemeBuilder.Register(&VpcAttachment{}, &VpcAttachmentList{})
}