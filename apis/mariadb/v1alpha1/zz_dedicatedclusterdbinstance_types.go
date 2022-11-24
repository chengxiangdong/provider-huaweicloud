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

type DedicatedclusterDBInstanceObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type DedicatedclusterDBInstanceParameters struct {

	// dedicated cluster id.
	// +kubebuilder:validation:Required
	ClusterID *string `json:"clusterId" tf:"cluster_id,omitempty"`

	// db engine version, default to 0.
	// +kubebuilder:validation:Optional
	DBVersionID *string `json:"dbVersionId,omitempty" tf:"db_version_id,omitempty"`

	// number of instance.
	// +kubebuilder:validation:Required
	GoodsNum *float64 `json:"goodsNum" tf:"goods_num,omitempty"`

	// name of this instance.
	// +kubebuilder:validation:Optional
	InstanceName *string `json:"instanceName,omitempty" tf:"instance_name,omitempty"`

	// instance memory.
	// +kubebuilder:validation:Required
	Memory *float64 `json:"memory" tf:"memory,omitempty"`

	// instance disk storage.
	// +kubebuilder:validation:Required
	Storage *float64 `json:"storage" tf:"storage,omitempty"`

	// subnet id, it&amp;#39;s required when vpcId is set.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1.Subnet
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// Tag description list.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// vpc id.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1.VPC
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcidRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcidSelector,omitempty" tf:"-"`
}

// DedicatedclusterDBInstanceSpec defines the desired state of DedicatedclusterDBInstance
type DedicatedclusterDBInstanceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DedicatedclusterDBInstanceParameters `json:"forProvider"`
}

// DedicatedclusterDBInstanceStatus defines the observed state of DedicatedclusterDBInstance.
type DedicatedclusterDBInstanceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DedicatedclusterDBInstanceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DedicatedclusterDBInstance is the Schema for the DedicatedclusterDBInstances API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloudjet}
type DedicatedclusterDBInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DedicatedclusterDBInstanceSpec   `json:"spec"`
	Status            DedicatedclusterDBInstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DedicatedclusterDBInstanceList contains a list of DedicatedclusterDBInstances
type DedicatedclusterDBInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DedicatedclusterDBInstance `json:"items"`
}

// Repository type metadata.
var (
	DedicatedclusterDBInstance_Kind             = "DedicatedclusterDBInstance"
	DedicatedclusterDBInstance_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DedicatedclusterDBInstance_Kind}.String()
	DedicatedclusterDBInstance_KindAPIVersion   = DedicatedclusterDBInstance_Kind + "." + CRDGroupVersion.String()
	DedicatedclusterDBInstance_GroupVersionKind = CRDGroupVersion.WithKind(DedicatedclusterDBInstance_Kind)
)

func init() {
	SchemeBuilder.Register(&DedicatedclusterDBInstance{}, &DedicatedclusterDBInstanceList{})
}
