/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type EipObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	PublicIP *string `json:"publicIp,omitempty" tf:"public_ip,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type EipParameters struct {

	// The zone of anycast. Valid value: `ANYCAST_ZONE_GLOBAL` and `ANYCAST_ZONE_OVERSEAS`.
	// +kubebuilder:validation:Optional
	AnycastZone *string `json:"anycastZone,omitempty" tf:"anycast_zone,omitempty"`

	// Indicates whether the anycast eip can be associated to a CLB.
	// +kubebuilder:validation:Optional
	ApplicableForClb *bool `json:"applicableForClb,omitempty" tf:"applicable_for_clb,omitempty"`

	// The charge type of eip. Valid value: `BANDWIDTH_PACKAGE`, `BANDWIDTH_POSTPAID_BY_HOUR` and `TRAFFIC_POSTPAID_BY_HOUR`.
	// +kubebuilder:validation:Optional
	InternetChargeType *string `json:"internetChargeType,omitempty" tf:"internet_charge_type,omitempty"`

	// The bandwidth limit of EIP, unit is Mbps.
	// +kubebuilder:validation:Optional
	InternetMaxBandwidthOut *float64 `json:"internetMaxBandwidthOut,omitempty" tf:"internet_max_bandwidth_out,omitempty"`

	// Internet service provider of eip. Valid value: `BGP`, `CMCC`, `CTCC` and `CUCC`.
	// +kubebuilder:validation:Optional
	InternetServiceProvider *string `json:"internetServiceProvider,omitempty" tf:"internet_service_provider,omitempty"`

	// The name of eip.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The tags of eip.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The type of eip. Valid value:  `EIP` and `AnycastEIP` and `HighQualityEIP`. Default is `EIP`.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// EipSpec defines the desired state of Eip
type EipSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EipParameters `json:"forProvider"`
}

// EipStatus defines the observed state of Eip.
type EipStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EipObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Eip is the Schema for the Eips API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloudjet}
type Eip struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EipSpec   `json:"spec"`
	Status            EipStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EipList contains a list of Eips
type EipList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Eip `json:"items"`
}

// Repository type metadata.
var (
	Eip_Kind             = "Eip"
	Eip_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Eip_Kind}.String()
	Eip_KindAPIVersion   = Eip_Kind + "." + CRDGroupVersion.String()
	Eip_GroupVersionKind = CRDGroupVersion.WithKind(Eip_Kind)
)

func init() {
	SchemeBuilder.Register(&Eip{}, &EipList{})
}