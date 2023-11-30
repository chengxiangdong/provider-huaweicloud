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

type RolePolicyAttachmentObservation struct {
	CreateMode *float64 `json:"createMode,omitempty" tf:"create_mode,omitempty"`

	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	PolicyType *string `json:"policyType,omitempty" tf:"policy_type,omitempty"`
}

type RolePolicyAttachmentParameters struct {

	// Name of the policy.
	// +crossplane:generate:reference:type=Policy
	// +kubebuilder:validation:Optional
	PolicyName *string `json:"policyName,omitempty" tf:"policy_name,omitempty"`

	// +kubebuilder:validation:Optional
	PolicyNameRef *v1.Reference `json:"policyNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	PolicyNameSelector *v1.Selector `json:"policyNameSelector,omitempty" tf:"-"`

	// Name of the attached CAM role.
	// +crossplane:generate:reference:type=Role
	// +kubebuilder:validation:Optional
	RoleName *string `json:"roleName,omitempty" tf:"role_name,omitempty"`

	// +kubebuilder:validation:Optional
	RoleNameRef *v1.Reference `json:"roleNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	RoleNameSelector *v1.Selector `json:"roleNameSelector,omitempty" tf:"-"`
}

// RolePolicyAttachmentSpec defines the desired state of RolePolicyAttachment
type RolePolicyAttachmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RolePolicyAttachmentParameters `json:"forProvider"`
}

// RolePolicyAttachmentStatus defines the observed state of RolePolicyAttachment.
type RolePolicyAttachmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RolePolicyAttachmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RolePolicyAttachment is the Schema for the RolePolicyAttachments API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloudjet}
type RolePolicyAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RolePolicyAttachmentSpec   `json:"spec"`
	Status            RolePolicyAttachmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RolePolicyAttachmentList contains a list of RolePolicyAttachments
type RolePolicyAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RolePolicyAttachment `json:"items"`
}

// Repository type metadata.
var (
	RolePolicyAttachment_Kind             = "RolePolicyAttachment"
	RolePolicyAttachment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RolePolicyAttachment_Kind}.String()
	RolePolicyAttachment_KindAPIVersion   = RolePolicyAttachment_Kind + "." + CRDGroupVersion.String()
	RolePolicyAttachment_GroupVersionKind = CRDGroupVersion.WithKind(RolePolicyAttachment_Kind)
)

func init() {
	SchemeBuilder.Register(&RolePolicyAttachment{}, &RolePolicyAttachmentList{})
}