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

type TopicObservation struct {
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	ForwardCosBucket *string `json:"forwardCosBucket,omitempty" tf:"forward_cos_bucket,omitempty"`

	ForwardInterval *float64 `json:"forwardInterval,omitempty" tf:"forward_interval,omitempty"`

	ForwardStatus *float64 `json:"forwardStatus,omitempty" tf:"forward_status,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	MessageStorageLocation *string `json:"messageStorageLocation,omitempty" tf:"message_storage_location,omitempty"`

	SegmentBytes *float64 `json:"segmentBytes,omitempty" tf:"segment_bytes,omitempty"`
}

type TopicParameters struct {

	// Clear log policy, log clear mode, default is `delete`. `delete`: logs are deleted according to the storage time. `compact`: logs are compressed according to the key. `compact, delete`: logs are compressed according to the key and will be deleted according to the storage time.
	// +kubebuilder:validation:Optional
	CleanUpPolicy *string `json:"cleanUpPolicy,omitempty" tf:"clean_up_policy,omitempty"`

	// Whether to open the ip whitelist, `true`: open, `false`: close.
	// +kubebuilder:validation:Optional
	EnableWhiteList *bool `json:"enableWhiteList,omitempty" tf:"enable_white_list,omitempty"`

	// Ip whitelist, quota limit, required when enableWhileList=true.
	// +kubebuilder:validation:Optional
	IPWhiteList []*string `json:"ipWhiteList,omitempty" tf:"ip_white_list,omitempty"`

	// Ckafka instance ID.
	// +crossplane:generate:reference:type=Instance
	// +kubebuilder:validation:Optional
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// +kubebuilder:validation:Optional
	InstanceIDRef *v1.Reference `json:"instanceIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	InstanceIDSelector *v1.Selector `json:"instanceIdSelector,omitempty" tf:"-"`

	// Max message bytes. min: 1024 Byte(1KB), max: 8388608 Byte(8MB).
	// +kubebuilder:validation:Optional
	MaxMessageBytes *float64 `json:"maxMessageBytes,omitempty" tf:"max_message_bytes,omitempty"`

	// The subject note. It must start with a letter, and the remaining part can contain letters, numbers and dashes (-).
	// +kubebuilder:validation:Optional
	Note *string `json:"note,omitempty" tf:"note,omitempty"`

	// The number of partition.
	// +kubebuilder:validation:Required
	PartitionNum *float64 `json:"partitionNum" tf:"partition_num,omitempty"`

	// The number of replica.
	// +kubebuilder:validation:Required
	ReplicaNum *float64 `json:"replicaNum" tf:"replica_num,omitempty"`

	// Message can be selected. Retention time, unit is ms, the current minimum value is 60000ms.
	// +kubebuilder:validation:Optional
	Retention *float64 `json:"retention,omitempty" tf:"retention,omitempty"`

	// Segment scrolling time, in ms, the current minimum is 3600000ms.
	// +kubebuilder:validation:Optional
	Segment *float64 `json:"segment,omitempty" tf:"segment,omitempty"`

	// Min number of sync replicas, Default is `1`.
	// +kubebuilder:validation:Optional
	SyncReplicaMinNum *float64 `json:"syncReplicaMinNum,omitempty" tf:"sync_replica_min_num,omitempty"`

	// Name of the CKafka topic. It must start with a letter, the rest can contain letters, numbers and dashes(-).
	// +kubebuilder:validation:Required
	TopicName *string `json:"topicName" tf:"topic_name,omitempty"`

	// Whether to allow unsynchronized replicas to be selected as leader, default is `false`, `true: `allowed, `false`: not allowed.
	// +kubebuilder:validation:Optional
	UncleanLeaderElectionEnable *bool `json:"uncleanLeaderElectionEnable,omitempty" tf:"unclean_leader_election_enable,omitempty"`
}

// TopicSpec defines the desired state of Topic
type TopicSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TopicParameters `json:"forProvider"`
}

// TopicStatus defines the observed state of Topic.
type TopicStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TopicObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Topic is the Schema for the Topics API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloudjet}
type Topic struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TopicSpec   `json:"spec"`
	Status            TopicStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TopicList contains a list of Topics
type TopicList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Topic `json:"items"`
}

// Repository type metadata.
var (
	Topic_Kind             = "Topic"
	Topic_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Topic_Kind}.String()
	Topic_KindAPIVersion   = Topic_Kind + "." + CRDGroupVersion.String()
	Topic_GroupVersionKind = CRDGroupVersion.WithKind(Topic_Kind)
)

func init() {
	SchemeBuilder.Register(&Topic{}, &TopicList{})
}