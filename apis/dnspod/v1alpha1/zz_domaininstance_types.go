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

type DomainInstanceInitParameters struct {

	// The Domain.
	// The Domain.
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// The Group Id of Domain.
	// The Group Id of Domain.
	GroupID *float64 `json:"groupId,omitempty" tf:"group_id,omitempty"`

	// Whether to Mark the Domain.
	// Whether to Mark the Domain.
	IsMark *string `json:"isMark,omitempty" tf:"is_mark,omitempty"`

	// The remark of Domain.
	// The remark of Domain.
	Remark *string `json:"remark,omitempty" tf:"remark,omitempty"`

	// The status of Domain.
	// The status of Domain.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type DomainInstanceObservation struct {

	// Create time of the domain.
	// Create time of the domain.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// The Domain.
	// The Domain.
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// The Group Id of Domain.
	// The Group Id of Domain.
	GroupID *float64 `json:"groupId,omitempty" tf:"group_id,omitempty"`

	// ID of the resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Whether to Mark the Domain.
	// Whether to Mark the Domain.
	IsMark *string `json:"isMark,omitempty" tf:"is_mark,omitempty"`

	// The remark of Domain.
	// The remark of Domain.
	Remark *string `json:"remark,omitempty" tf:"remark,omitempty"`

	// Is secondary DNS enabled.
	// Is secondary DNS enabled.
	SlaveDNS *string `json:"slaveDns,omitempty" tf:"slave_dns,omitempty"`

	// The status of Domain.
	// The status of Domain.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type DomainInstanceParameters struct {

	// The Domain.
	// The Domain.
	// +kubebuilder:validation:Optional
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// The Group Id of Domain.
	// The Group Id of Domain.
	// +kubebuilder:validation:Optional
	GroupID *float64 `json:"groupId,omitempty" tf:"group_id,omitempty"`

	// Whether to Mark the Domain.
	// Whether to Mark the Domain.
	// +kubebuilder:validation:Optional
	IsMark *string `json:"isMark,omitempty" tf:"is_mark,omitempty"`

	// The remark of Domain.
	// The remark of Domain.
	// +kubebuilder:validation:Optional
	Remark *string `json:"remark,omitempty" tf:"remark,omitempty"`

	// The status of Domain.
	// The status of Domain.
	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

// DomainInstanceSpec defines the desired state of DomainInstance
type DomainInstanceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DomainInstanceParameters `json:"forProvider"`
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
	InitProvider DomainInstanceInitParameters `json:"initProvider,omitempty"`
}

// DomainInstanceStatus defines the observed state of DomainInstance.
type DomainInstanceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DomainInstanceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DomainInstance is the Schema for the DomainInstances API. Provide a resource to create a DnsPod Domain instance.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloud}
type DomainInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.domain) || (has(self.initProvider) && has(self.initProvider.domain))",message="spec.forProvider.domain is a required parameter"
	Spec   DomainInstanceSpec   `json:"spec"`
	Status DomainInstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DomainInstanceList contains a list of DomainInstances
type DomainInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DomainInstance `json:"items"`
}

// Repository type metadata.
var (
	DomainInstance_Kind             = "DomainInstance"
	DomainInstance_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DomainInstance_Kind}.String()
	DomainInstance_KindAPIVersion   = DomainInstance_Kind + "." + CRDGroupVersion.String()
	DomainInstance_GroupVersionKind = CRDGroupVersion.WithKind(DomainInstance_Kind)
)

func init() {
	SchemeBuilder.Register(&DomainInstance{}, &DomainInstanceList{})
}