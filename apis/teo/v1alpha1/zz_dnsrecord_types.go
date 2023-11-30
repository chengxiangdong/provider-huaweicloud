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

type DnsRecordObservation struct {
	Cname *string `json:"cname,omitempty" tf:"cname,omitempty"`

	CreatedOn *string `json:"createdOn,omitempty" tf:"created_on,omitempty"`

	DNSRecordID *string `json:"dnsRecordId,omitempty" tf:"dns_record_id,omitempty"`

	DomainStatus []*string `json:"domainStatus,omitempty" tf:"domain_status,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Locked *bool `json:"locked,omitempty" tf:"locked,omitempty"`

	ModifiedOn *string `json:"modifiedOn,omitempty" tf:"modified_on,omitempty"`
}

type DnsRecordParameters struct {

	// DNS record Content.
	// +kubebuilder:validation:Required
	Content *string `json:"content" tf:"content,omitempty"`

	// Proxy mode. Valid values:- `dns_only`: only DNS resolution of the subdomain is enabled.- `proxied`: subdomain is proxied and accelerated.
	// +kubebuilder:validation:Required
	Mode *string `json:"mode" tf:"mode,omitempty"`

	// DNS record Name.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Priority of the record. Valid value range: 1-50, the smaller value, the higher priority.
	// +kubebuilder:validation:Optional
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// Resolution status. Valid values: `active`, `pending`.
	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Time to live of the DNS record cache in seconds.
	// +kubebuilder:validation:Optional
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`

	// DNS record Type. Valid values: `A`, `AAAA`, `CNAME`, `MX`, `TXT`, `NS`, `CAA`, `SRV`.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`

	// Site ID.
	// +crossplane:generate:reference:type=Zone
	// +kubebuilder:validation:Optional
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`

	// +kubebuilder:validation:Optional
	ZoneIDRef *v1.Reference `json:"zoneIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ZoneIDSelector *v1.Selector `json:"zoneIdSelector,omitempty" tf:"-"`
}

// DnsRecordSpec defines the desired state of DnsRecord
type DnsRecordSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DnsRecordParameters `json:"forProvider"`
}

// DnsRecordStatus defines the observed state of DnsRecord.
type DnsRecordStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DnsRecordObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DnsRecord is the Schema for the DnsRecords API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloudjet}
type DnsRecord struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DnsRecordSpec   `json:"spec"`
	Status            DnsRecordStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DnsRecordList contains a list of DnsRecords
type DnsRecordList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DnsRecord `json:"items"`
}

// Repository type metadata.
var (
	DnsRecord_Kind             = "DnsRecord"
	DnsRecord_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DnsRecord_Kind}.String()
	DnsRecord_KindAPIVersion   = DnsRecord_Kind + "." + CRDGroupVersion.String()
	DnsRecord_GroupVersionKind = CRDGroupVersion.WithKind(DnsRecord_Kind)
)

func init() {
	SchemeBuilder.Register(&DnsRecord{}, &DnsRecordList{})
}