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

type CustomDomainObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Status *float64 `json:"status,omitempty" tf:"status,omitempty"`
}

type CustomDomainParameters struct {

	// Unique certificate ID of the custom domain name to be bound. You can choose to upload for the `protocol` attribute value `https` or `http&https`.
	// +kubebuilder:validation:Optional
	CertificateID *string `json:"certificateId,omitempty" tf:"certificate_id,omitempty"`

	// Default domain name.
	// +kubebuilder:validation:Required
	DefaultDomain *string `json:"defaultDomain" tf:"default_domain,omitempty"`

	// Whether the default path mapping is used. The default value is `true`. When it is `false`, it means custom path mapping. In this case, the `path_mappings` attribute is required.
	// +kubebuilder:validation:Optional
	IsDefaultMapping *bool `json:"isDefaultMapping,omitempty" tf:"is_default_mapping,omitempty"`

	// Network type. Valid values: `OUTER`, `INNER`.
	// +kubebuilder:validation:Required
	NetType *string `json:"netType" tf:"net_type,omitempty"`

	// Custom domain name path mapping. The data format is: `path#environment`. Optional values for the environment are `test`, `prepub`, and `release`.
	// +kubebuilder:validation:Optional
	PathMappings []*string `json:"pathMappings,omitempty" tf:"path_mappings,omitempty"`

	// Protocol supported by service. Valid values: `http`, `https`, `http&https`.
	// +kubebuilder:validation:Required
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`

	// Unique service ID.
	// +crossplane:generate:reference:type=Service
	// +kubebuilder:validation:Optional
	ServiceID *string `json:"serviceId,omitempty" tf:"service_id,omitempty"`

	// +kubebuilder:validation:Optional
	ServiceIDRef *v1.Reference `json:"serviceIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ServiceIDSelector *v1.Selector `json:"serviceIdSelector,omitempty" tf:"-"`

	// Custom domain name to be bound.
	// +kubebuilder:validation:Required
	SubDomain *string `json:"subDomain" tf:"sub_domain,omitempty"`
}

// CustomDomainSpec defines the desired state of CustomDomain
type CustomDomainSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CustomDomainParameters `json:"forProvider"`
}

// CustomDomainStatus defines the observed state of CustomDomain.
type CustomDomainStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CustomDomainObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CustomDomain is the Schema for the CustomDomains API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloudjet}
type CustomDomain struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CustomDomainSpec   `json:"spec"`
	Status            CustomDomainStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CustomDomainList contains a list of CustomDomains
type CustomDomainList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CustomDomain `json:"items"`
}

// Repository type metadata.
var (
	CustomDomain_Kind             = "CustomDomain"
	CustomDomain_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CustomDomain_Kind}.String()
	CustomDomain_KindAPIVersion   = CustomDomain_Kind + "." + CRDGroupVersion.String()
	CustomDomain_GroupVersionKind = CRDGroupVersion.WithKind(CustomDomain_Kind)
)

func init() {
	SchemeBuilder.Register(&CustomDomain{}, &CustomDomainList{})
}