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

type EipAssociationInitParameters struct {

	// The ID of EIP.
	// +crossplane:generate:reference:type=Eip
	EIPID *string `json:"eipId,omitempty" tf:"eip_id,omitempty"`

	// Reference to a Eip to populate eipId.
	// +kubebuilder:validation:Optional
	EIPIDRef *v1.Reference `json:"eipIdRef,omitempty" tf:"-"`

	// Selector for a Eip to populate eipId.
	// +kubebuilder:validation:Optional
	EIPIDSelector *v1.Selector `json:"eipIdSelector,omitempty" tf:"-"`

	// The CVM or CLB instance id going to bind with the EIP. This field is conflict with `network_interface_id` and `private_ip fields`.
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Indicates the network interface id like `eni-xxxxxx`. This field is conflict with `instance_id`.
	NetworkInterfaceID *string `json:"networkInterfaceId,omitempty" tf:"network_interface_id,omitempty"`

	// Indicates an IP belongs to the `network_interface_id`. This field is conflict with `instance_id`.
	PrivateIP *string `json:"privateIp,omitempty" tf:"private_ip,omitempty"`
}

type EipAssociationObservation struct {

	// The ID of EIP.
	EIPID *string `json:"eipId,omitempty" tf:"eip_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The CVM or CLB instance id going to bind with the EIP. This field is conflict with `network_interface_id` and `private_ip fields`.
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Indicates the network interface id like `eni-xxxxxx`. This field is conflict with `instance_id`.
	NetworkInterfaceID *string `json:"networkInterfaceId,omitempty" tf:"network_interface_id,omitempty"`

	// Indicates an IP belongs to the `network_interface_id`. This field is conflict with `instance_id`.
	PrivateIP *string `json:"privateIp,omitempty" tf:"private_ip,omitempty"`
}

type EipAssociationParameters struct {

	// The ID of EIP.
	// +crossplane:generate:reference:type=Eip
	// +kubebuilder:validation:Optional
	EIPID *string `json:"eipId,omitempty" tf:"eip_id,omitempty"`

	// Reference to a Eip to populate eipId.
	// +kubebuilder:validation:Optional
	EIPIDRef *v1.Reference `json:"eipIdRef,omitempty" tf:"-"`

	// Selector for a Eip to populate eipId.
	// +kubebuilder:validation:Optional
	EIPIDSelector *v1.Selector `json:"eipIdSelector,omitempty" tf:"-"`

	// The CVM or CLB instance id going to bind with the EIP. This field is conflict with `network_interface_id` and `private_ip fields`.
	// +kubebuilder:validation:Optional
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Indicates the network interface id like `eni-xxxxxx`. This field is conflict with `instance_id`.
	// +kubebuilder:validation:Optional
	NetworkInterfaceID *string `json:"networkInterfaceId,omitempty" tf:"network_interface_id,omitempty"`

	// Indicates an IP belongs to the `network_interface_id`. This field is conflict with `instance_id`.
	// +kubebuilder:validation:Optional
	PrivateIP *string `json:"privateIp,omitempty" tf:"private_ip,omitempty"`
}

// EipAssociationSpec defines the desired state of EipAssociation
type EipAssociationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EipAssociationParameters `json:"forProvider"`
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
	InitProvider EipAssociationInitParameters `json:"initProvider,omitempty"`
}

// EipAssociationStatus defines the observed state of EipAssociation.
type EipAssociationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EipAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// EipAssociation is the Schema for the EipAssociations API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloud}
type EipAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EipAssociationSpec   `json:"spec"`
	Status            EipAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EipAssociationList contains a list of EipAssociations
type EipAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EipAssociation `json:"items"`
}

// Repository type metadata.
var (
	EipAssociation_Kind             = "EipAssociation"
	EipAssociation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EipAssociation_Kind}.String()
	EipAssociation_KindAPIVersion   = EipAssociation_Kind + "." + CRDGroupVersion.String()
	EipAssociation_GroupVersionKind = CRDGroupVersion.WithKind(EipAssociation_Kind)
)

func init() {
	SchemeBuilder.Register(&EipAssociation{}, &EipAssociationList{})
}
