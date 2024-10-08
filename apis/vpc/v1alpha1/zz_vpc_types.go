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

type VPCInitParameters struct {

	// List of Assistant CIDR, NOTE: Only `NORMAL` typed CIDRs included, check the Docker CIDR by readonly `assistant_docker_cidrs`.
	// +listType=set
	AssistantCidrs []*string `json:"assistantCidrs,omitempty" tf:"assistant_cidrs,omitempty"`

	// A network address block which should be a subnet of the three internal network segments (10.0.0.0/16, 172.16.0.0/12 and 192.168.0.0/16).
	CidrBlock *string `json:"cidrBlock,omitempty" tf:"cidr_block,omitempty"`

	// The DNS server list of the VPC. And you can specify 0 to 5 servers to this list.
	// +listType=set
	DNSServers []*string `json:"dnsServers,omitempty" tf:"dns_servers,omitempty"`

	// Indicates whether VPC multicast is enabled. The default value is 'true'.
	IsMulticast *bool `json:"isMulticast,omitempty" tf:"is_multicast,omitempty"`

	// The name of the VPC.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Tags of the VPC.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type VPCObservation struct {

	// List of Assistant CIDR, NOTE: Only `NORMAL` typed CIDRs included, check the Docker CIDR by readonly `assistant_docker_cidrs`.
	// +listType=set
	AssistantCidrs []*string `json:"assistantCidrs,omitempty" tf:"assistant_cidrs,omitempty"`

	// A network address block which should be a subnet of the three internal network segments (10.0.0.0/16, 172.16.0.0/12 and 192.168.0.0/16).
	CidrBlock *string `json:"cidrBlock,omitempty" tf:"cidr_block,omitempty"`

	// Creation time of VPC.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// The DNS server list of the VPC. And you can specify 0 to 5 servers to this list.
	// +listType=set
	DNSServers []*string `json:"dnsServers,omitempty" tf:"dns_servers,omitempty"`

	// Default route table id, which created automatically after VPC create.
	DefaultRouteTableID *string `json:"defaultRouteTableId,omitempty" tf:"default_route_table_id,omitempty"`

	// List of Docker Assistant CIDR.
	DockerAssistantCidrs []*string `json:"dockerAssistantCidrs,omitempty" tf:"docker_assistant_cidrs,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Indicates whether it is the default VPC for this region.
	IsDefault *bool `json:"isDefault,omitempty" tf:"is_default,omitempty"`

	// Indicates whether VPC multicast is enabled. The default value is 'true'.
	IsMulticast *bool `json:"isMulticast,omitempty" tf:"is_multicast,omitempty"`

	// The name of the VPC.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Tags of the VPC.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type VPCParameters struct {

	// List of Assistant CIDR, NOTE: Only `NORMAL` typed CIDRs included, check the Docker CIDR by readonly `assistant_docker_cidrs`.
	// +kubebuilder:validation:Optional
	// +listType=set
	AssistantCidrs []*string `json:"assistantCidrs,omitempty" tf:"assistant_cidrs,omitempty"`

	// A network address block which should be a subnet of the three internal network segments (10.0.0.0/16, 172.16.0.0/12 and 192.168.0.0/16).
	// +kubebuilder:validation:Optional
	CidrBlock *string `json:"cidrBlock,omitempty" tf:"cidr_block,omitempty"`

	// The DNS server list of the VPC. And you can specify 0 to 5 servers to this list.
	// +kubebuilder:validation:Optional
	// +listType=set
	DNSServers []*string `json:"dnsServers,omitempty" tf:"dns_servers,omitempty"`

	// Indicates whether VPC multicast is enabled. The default value is 'true'.
	// +kubebuilder:validation:Optional
	IsMulticast *bool `json:"isMulticast,omitempty" tf:"is_multicast,omitempty"`

	// The name of the VPC.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Tags of the VPC.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// VPCSpec defines the desired state of VPC
type VPCSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VPCParameters `json:"forProvider"`
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
	InitProvider VPCInitParameters `json:"initProvider,omitempty"`
}

// VPCStatus defines the observed state of VPC.
type VPCStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VPCObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// VPC is the Schema for the VPCs API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloud}
type VPC struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.cidrBlock) || (has(self.initProvider) && has(self.initProvider.cidrBlock))",message="spec.forProvider.cidrBlock is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   VPCSpec   `json:"spec"`
	Status VPCStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VPCList contains a list of VPCs
type VPCList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VPC `json:"items"`
}

// Repository type metadata.
var (
	VPC_Kind             = "VPC"
	VPC_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VPC_Kind}.String()
	VPC_KindAPIVersion   = VPC_Kind + "." + CRDGroupVersion.String()
	VPC_GroupVersionKind = CRDGroupVersion.WithKind(VPC_Kind)
)

func init() {
	SchemeBuilder.Register(&VPC{}, &VPCList{})
}
