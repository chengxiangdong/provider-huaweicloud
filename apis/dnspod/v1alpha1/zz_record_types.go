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

type RecordInitParameters struct {

	// The Domain.
	// The Domain.
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// MX priority, valid when the record type is MX, range 1-20. Note: must set when record type equal MX.
	// MX priority, valid when the record type is MX, range 1-20. Note: must set when record type equal MX.
	Mx *float64 `json:"mx,omitempty" tf:"mx,omitempty"`

	// The record line.
	// The record line.
	RecordLine *string `json:"recordLine,omitempty" tf:"record_line,omitempty"`

	// The record type.
	// The record type.
	RecordType *string `json:"recordType,omitempty" tf:"record_type,omitempty"`

	// The Remark of record.
	// The Remark of record.
	Remark *string `json:"remark,omitempty" tf:"remark,omitempty"`

	// Records the initial state, with values ranging from ENABLE and DISABLE. The default is ENABLE, and if DISABLE is passed in, resolution will not take effect and the limits of load balancing will not be verified.
	// Records the initial state, with values ranging from ENABLE and DISABLE. The default is ENABLE, and if DISABLE is passed in, resolution will not take effect and the limits of load balancing will not be verified.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// The host records, default value is @.
	// The host records, default value is `@`.
	SubDomain *string `json:"subDomain,omitempty" tf:"sub_domain,omitempty"`

	// TTL, the range is 1-604800, and the minimum value of different levels of domain names is different. Default is 600.
	// TTL, the range is 1-604800, and the minimum value of different levels of domain names is different. Default is 600.
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`

	// The record value.
	// The record value.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`

	// Weight information. An integer from 0 to 100. Only enterprise VIP domain names are available, 0 means off, does not pass this parameter, means that the weight information is not set. Default is 0.
	// Weight information. An integer from 0 to 100. Only enterprise VIP domain names are available, 0 means off, does not pass this parameter, means that the weight information is not set. Default is 0.
	Weight *float64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

type RecordObservation struct {

	// The Domain.
	// The Domain.
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// ID of the resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The monitoring status of the record.
	// The monitoring status of the record.
	MonitorStatus *string `json:"monitorStatus,omitempty" tf:"monitor_status,omitempty"`

	// MX priority, valid when the record type is MX, range 1-20. Note: must set when record type equal MX.
	// MX priority, valid when the record type is MX, range 1-20. Note: must set when record type equal MX.
	Mx *float64 `json:"mx,omitempty" tf:"mx,omitempty"`

	// The record line.
	// The record line.
	RecordLine *string `json:"recordLine,omitempty" tf:"record_line,omitempty"`

	// The record type.
	// The record type.
	RecordType *string `json:"recordType,omitempty" tf:"record_type,omitempty"`

	// The Remark of record.
	// The Remark of record.
	Remark *string `json:"remark,omitempty" tf:"remark,omitempty"`

	// Records the initial state, with values ranging from ENABLE and DISABLE. The default is ENABLE, and if DISABLE is passed in, resolution will not take effect and the limits of load balancing will not be verified.
	// Records the initial state, with values ranging from ENABLE and DISABLE. The default is ENABLE, and if DISABLE is passed in, resolution will not take effect and the limits of load balancing will not be verified.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// The host records, default value is @.
	// The host records, default value is `@`.
	SubDomain *string `json:"subDomain,omitempty" tf:"sub_domain,omitempty"`

	// TTL, the range is 1-604800, and the minimum value of different levels of domain names is different. Default is 600.
	// TTL, the range is 1-604800, and the minimum value of different levels of domain names is different. Default is 600.
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`

	// The record value.
	// The record value.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`

	// Weight information. An integer from 0 to 100. Only enterprise VIP domain names are available, 0 means off, does not pass this parameter, means that the weight information is not set. Default is 0.
	// Weight information. An integer from 0 to 100. Only enterprise VIP domain names are available, 0 means off, does not pass this parameter, means that the weight information is not set. Default is 0.
	Weight *float64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

type RecordParameters struct {

	// The Domain.
	// The Domain.
	// +kubebuilder:validation:Optional
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// MX priority, valid when the record type is MX, range 1-20. Note: must set when record type equal MX.
	// MX priority, valid when the record type is MX, range 1-20. Note: must set when record type equal MX.
	// +kubebuilder:validation:Optional
	Mx *float64 `json:"mx,omitempty" tf:"mx,omitempty"`

	// The record line.
	// The record line.
	// +kubebuilder:validation:Optional
	RecordLine *string `json:"recordLine,omitempty" tf:"record_line,omitempty"`

	// The record type.
	// The record type.
	// +kubebuilder:validation:Optional
	RecordType *string `json:"recordType,omitempty" tf:"record_type,omitempty"`

	// The Remark of record.
	// The Remark of record.
	// +kubebuilder:validation:Optional
	Remark *string `json:"remark,omitempty" tf:"remark,omitempty"`

	// Records the initial state, with values ranging from ENABLE and DISABLE. The default is ENABLE, and if DISABLE is passed in, resolution will not take effect and the limits of load balancing will not be verified.
	// Records the initial state, with values ranging from ENABLE and DISABLE. The default is ENABLE, and if DISABLE is passed in, resolution will not take effect and the limits of load balancing will not be verified.
	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// The host records, default value is @.
	// The host records, default value is `@`.
	// +kubebuilder:validation:Optional
	SubDomain *string `json:"subDomain,omitempty" tf:"sub_domain,omitempty"`

	// TTL, the range is 1-604800, and the minimum value of different levels of domain names is different. Default is 600.
	// TTL, the range is 1-604800, and the minimum value of different levels of domain names is different. Default is 600.
	// +kubebuilder:validation:Optional
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`

	// The record value.
	// The record value.
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`

	// Weight information. An integer from 0 to 100. Only enterprise VIP domain names are available, 0 means off, does not pass this parameter, means that the weight information is not set. Default is 0.
	// Weight information. An integer from 0 to 100. Only enterprise VIP domain names are available, 0 means off, does not pass this parameter, means that the weight information is not set. Default is 0.
	// +kubebuilder:validation:Optional
	Weight *float64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

// RecordSpec defines the desired state of Record
type RecordSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RecordParameters `json:"forProvider"`
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
	InitProvider RecordInitParameters `json:"initProvider,omitempty"`
}

// RecordStatus defines the observed state of Record.
type RecordStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RecordObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Record is the Schema for the Records API. Provide a resource to create a DnsPod record.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloud}
type Record struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.domain) || (has(self.initProvider) && has(self.initProvider.domain))",message="spec.forProvider.domain is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.recordLine) || (has(self.initProvider) && has(self.initProvider.recordLine))",message="spec.forProvider.recordLine is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.recordType) || (has(self.initProvider) && has(self.initProvider.recordType))",message="spec.forProvider.recordType is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.value) || (has(self.initProvider) && has(self.initProvider.value))",message="spec.forProvider.value is a required parameter"
	Spec   RecordSpec   `json:"spec"`
	Status RecordStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RecordList contains a list of Records
type RecordList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Record `json:"items"`
}

// Repository type metadata.
var (
	Record_Kind             = "Record"
	Record_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Record_Kind}.String()
	Record_KindAPIVersion   = Record_Kind + "." + CRDGroupVersion.String()
	Record_GroupVersionKind = CRDGroupVersion.WithKind(Record_Kind)
)

func init() {
	SchemeBuilder.Register(&Record{}, &RecordList{})
}