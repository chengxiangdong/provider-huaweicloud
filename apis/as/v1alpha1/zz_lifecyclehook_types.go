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

type LifecycleHookObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type LifecycleHookParameters struct {

	// Defines the action the AS group should take when the lifecycle hook timeout elapses or if an unexpected failure occurs. Valid values: `CONTINUE` and `ABANDON`. The default value is `CONTINUE`.
	// +kubebuilder:validation:Optional
	DefaultResult *string `json:"defaultResult,omitempty" tf:"default_result,omitempty"`

	// Defines the amount of time, in seconds, that can elapse before the lifecycle hook times out. Valid value ranges: (30~7200). and default value is `300`.
	// +kubebuilder:validation:Optional
	HeartbeatTimeout *float64 `json:"heartbeatTimeout,omitempty" tf:"heartbeat_timeout,omitempty"`

	// The name of the lifecycle hook.
	// +kubebuilder:validation:Required
	LifecycleHookName *string `json:"lifecycleHookName" tf:"lifecycle_hook_name,omitempty"`

	// The instance state to which you want to attach the lifecycle hook. Valid values: `INSTANCE_LAUNCHING` and `INSTANCE_TERMINATING`.
	// +kubebuilder:validation:Required
	LifecycleTransition *string `json:"lifecycleTransition" tf:"lifecycle_transition,omitempty"`

	// Contains additional information that you want to include any time AS sends a message to the notification target.
	// +kubebuilder:validation:Optional
	NotificationMetadata *string `json:"notificationMetadata,omitempty" tf:"notification_metadata,omitempty"`

	// For CMQ_QUEUE type, a name of queue must be set.
	// +kubebuilder:validation:Optional
	NotificationQueueName *string `json:"notificationQueueName,omitempty" tf:"notification_queue_name,omitempty"`

	// Target type. Valid values: `CMQ_QUEUE`, `CMQ_TOPIC`.
	// +kubebuilder:validation:Optional
	NotificationTargetType *string `json:"notificationTargetType,omitempty" tf:"notification_target_type,omitempty"`

	// For CMQ_TOPIC type, a name of topic must be set.
	// +kubebuilder:validation:Optional
	NotificationTopicName *string `json:"notificationTopicName,omitempty" tf:"notification_topic_name,omitempty"`

	// ID of a scaling group.
	// +crossplane:generate:reference:type=ScalingGroup
	// +kubebuilder:validation:Optional
	ScalingGroupID *string `json:"scalingGroupId,omitempty" tf:"scaling_group_id,omitempty"`

	// +kubebuilder:validation:Optional
	ScalingGroupIDRef *v1.Reference `json:"scalingGroupIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ScalingGroupIDSelector *v1.Selector `json:"scalingGroupIdSelector,omitempty" tf:"-"`
}

// LifecycleHookSpec defines the desired state of LifecycleHook
type LifecycleHookSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LifecycleHookParameters `json:"forProvider"`
}

// LifecycleHookStatus defines the observed state of LifecycleHook.
type LifecycleHookStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LifecycleHookObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LifecycleHook is the Schema for the LifecycleHooks API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloudjet}
type LifecycleHook struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LifecycleHookSpec   `json:"spec"`
	Status            LifecycleHookStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LifecycleHookList contains a list of LifecycleHooks
type LifecycleHookList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LifecycleHook `json:"items"`
}

// Repository type metadata.
var (
	LifecycleHook_Kind             = "LifecycleHook"
	LifecycleHook_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LifecycleHook_Kind}.String()
	LifecycleHook_KindAPIVersion   = LifecycleHook_Kind + "." + CRDGroupVersion.String()
	LifecycleHook_GroupVersionKind = CRDGroupVersion.WithKind(LifecycleHook_Kind)
)

func init() {
	SchemeBuilder.Register(&LifecycleHook{}, &LifecycleHookList{})
}