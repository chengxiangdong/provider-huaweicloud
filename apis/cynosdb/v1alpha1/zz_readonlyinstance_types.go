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

type ReadonlyInstanceObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	InstanceStatus *string `json:"instanceStatus,omitempty" tf:"instance_status,omitempty"`

	InstanceStorageSize *float64 `json:"instanceStorageSize,omitempty" tf:"instance_storage_size,omitempty"`
}

type ReadonlyInstanceParameters struct {

	// Cluster ID which the readonly instance belongs to.
	// +crossplane:generate:reference:type=Cluster
	// +kubebuilder:validation:Optional
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// +kubebuilder:validation:Optional
	ClusterIDRef *v1.Reference `json:"clusterIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ClusterIDSelector *v1.Selector `json:"clusterIdSelector,omitempty" tf:"-"`

	// Indicate whether to delete readonly instance directly or not. Default is false. If set true, instance will be deleted instead of staying recycle bin. Note: works for both `PREPAID` and `POSTPAID_BY_HOUR` cluster.
	// +kubebuilder:validation:Optional
	ForceDelete *bool `json:"forceDelete,omitempty" tf:"force_delete,omitempty"`

	// The number of CPU cores of read-write type instance in the CynosDB cluster. Required while creating normal cluster. Note: modification of this field will take effect immediately, if want to upgrade on maintenance window, please upgrade from console.
	// +kubebuilder:validation:Optional
	InstanceCPUCore *float64 `json:"instanceCpuCore,omitempty" tf:"instance_cpu_core,omitempty"`

	// Duration time for maintenance, unit in second. `3600` by default.
	// +kubebuilder:validation:Optional
	InstanceMaintainDuration *float64 `json:"instanceMaintainDuration,omitempty" tf:"instance_maintain_duration,omitempty"`

	// Offset time from 00:00, unit in second. For example, 03:00am should be `10800`. `10800` by default.
	// +kubebuilder:validation:Optional
	InstanceMaintainStartTime *float64 `json:"instanceMaintainStartTime,omitempty" tf:"instance_maintain_start_time,omitempty"`

	// Weekdays for maintenance. `["Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"]` by default.
	// +kubebuilder:validation:Optional
	InstanceMaintainWeekdays []*string `json:"instanceMaintainWeekdays,omitempty" tf:"instance_maintain_weekdays,omitempty"`

	// Memory capacity of read-write type instance, unit in GB. Required while creating normal cluster. Note: modification of this field will take effect immediately, if want to upgrade on maintenance window, please upgrade from console.
	// +kubebuilder:validation:Optional
	InstanceMemorySize *float64 `json:"instanceMemorySize,omitempty" tf:"instance_memory_size,omitempty"`

	// Name of instance.
	// +kubebuilder:validation:Required
	InstanceName *string `json:"instanceName" tf:"instance_name,omitempty"`
}

// ReadonlyInstanceSpec defines the desired state of ReadonlyInstance
type ReadonlyInstanceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ReadonlyInstanceParameters `json:"forProvider"`
}

// ReadonlyInstanceStatus defines the observed state of ReadonlyInstance.
type ReadonlyInstanceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ReadonlyInstanceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ReadonlyInstance is the Schema for the ReadonlyInstances API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloudjet}
type ReadonlyInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ReadonlyInstanceSpec   `json:"spec"`
	Status            ReadonlyInstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ReadonlyInstanceList contains a list of ReadonlyInstances
type ReadonlyInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ReadonlyInstance `json:"items"`
}

// Repository type metadata.
var (
	ReadonlyInstance_Kind             = "ReadonlyInstance"
	ReadonlyInstance_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ReadonlyInstance_Kind}.String()
	ReadonlyInstance_KindAPIVersion   = ReadonlyInstance_Kind + "." + CRDGroupVersion.String()
	ReadonlyInstance_GroupVersionKind = CRDGroupVersion.WithKind(ReadonlyInstance_Kind)
)

func init() {
	SchemeBuilder.Register(&ReadonlyInstance{}, &ReadonlyInstanceList{})
}
