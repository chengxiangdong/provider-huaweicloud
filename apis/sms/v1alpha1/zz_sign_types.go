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

type SignObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type SignParameters struct {

	// Power of attorney, which should be submitted if SignPurpose is for use by others. You should Base64-encode the image first, remove the prefix data:image/jpeg;base64, from the resulted string, and then use it as the value of this parameter. Note: this field will take effect only when SignPurpose is 1 (for user by others).
	// +kubebuilder:validation:Optional
	CommissionImage *string `json:"commissionImage,omitempty" tf:"commission_image,omitempty"`

	// DocumentType is used for enterprise authentication, or website, app authentication, etc. DocumentType: 0, 1, 2, 3, 4, 5, 6, 7, 8.
	// +kubebuilder:validation:Required
	DocumentType *float64 `json:"documentType" tf:"document_type,omitempty"`

	// Whether it is Global SMS: 0: Mainland China SMS; 1: Global SMS.
	// +kubebuilder:validation:Required
	International *float64 `json:"international" tf:"international,omitempty"`

	// You should Base64-encode the image of the identity certificate corresponding to the signature first, remove the prefix data:image/jpeg;base64, from the resulted string, and then use it as the value of this parameter.
	// +kubebuilder:validation:Required
	ProofImage *string `json:"proofImage" tf:"proof_image,omitempty"`

	// Signature application remarks.
	// +kubebuilder:validation:Optional
	Remark *string `json:"remark,omitempty" tf:"remark,omitempty"`

	// Sms sign name, unique.
	// +kubebuilder:validation:Required
	SignName *string `json:"signName" tf:"sign_name,omitempty"`

	// Signature purpose: 0: for personal use; 1: for others.
	// +kubebuilder:validation:Required
	SignPurpose *float64 `json:"signPurpose" tf:"sign_purpose,omitempty"`

	// Sms sign type: 0, 1, 2, 3, 4, 5, 6.
	// +kubebuilder:validation:Required
	SignType *float64 `json:"signType" tf:"sign_type,omitempty"`
}

// SignSpec defines the desired state of Sign
type SignSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SignParameters `json:"forProvider"`
}

// SignStatus defines the observed state of Sign.
type SignStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SignObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Sign is the Schema for the Signs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloudjet}
type Sign struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SignSpec   `json:"spec"`
	Status            SignStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SignList contains a list of Signs
type SignList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Sign `json:"items"`
}

// Repository type metadata.
var (
	Sign_Kind             = "Sign"
	Sign_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Sign_Kind}.String()
	Sign_KindAPIVersion   = Sign_Kind + "." + CRDGroupVersion.String()
	Sign_GroupVersionKind = CRDGroupVersion.WithKind(Sign_Kind)
)

func init() {
	SchemeBuilder.Register(&Sign{}, &SignList{})
}