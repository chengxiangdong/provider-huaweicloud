// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

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

type BackupConfigInitParameters struct {

	// It has been deprecated from version 1.58.2. It makes no difference to online config at all Specifys which day the backup action should take place. Valid values: Monday, Tuesday, Wednesday, Thursday, Friday, Saturday and Sunday.
	// Specifys which day the backup action should take place. Valid values: `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday` and `Sunday`.
	BackupPeriod []*string `json:"backupPeriod,omitempty" tf:"backup_period,omitempty"`

	// Specifys what time the backup action should take place. And the time interval should be one hour.
	// Specifys what time the backup action should take place. And the time interval should be one hour.
	BackupTime *string `json:"backupTime,omitempty" tf:"backup_time,omitempty"`
}

type BackupConfigObservation struct {

	// It has been deprecated from version 1.58.2. It makes no difference to online config at all Specifys which day the backup action should take place. Valid values: Monday, Tuesday, Wednesday, Thursday, Friday, Saturday and Sunday.
	// Specifys which day the backup action should take place. Valid values: `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday` and `Sunday`.
	BackupPeriod []*string `json:"backupPeriod,omitempty" tf:"backup_period,omitempty"`

	// Specifys what time the backup action should take place. And the time interval should be one hour.
	// Specifys what time the backup action should take place. And the time interval should be one hour.
	BackupTime *string `json:"backupTime,omitempty" tf:"backup_time,omitempty"`

	// ID of the resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// ID of a redis instance to which the policy will be applied.
	// ID of a redis instance to which the policy will be applied.
	RedisID *string `json:"redisId,omitempty" tf:"redis_id,omitempty"`
}

type BackupConfigParameters struct {

	// It has been deprecated from version 1.58.2. It makes no difference to online config at all Specifys which day the backup action should take place. Valid values: Monday, Tuesday, Wednesday, Thursday, Friday, Saturday and Sunday.
	// Specifys which day the backup action should take place. Valid values: `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday` and `Sunday`.
	// +kubebuilder:validation:Optional
	BackupPeriod []*string `json:"backupPeriod,omitempty" tf:"backup_period,omitempty"`

	// Specifys what time the backup action should take place. And the time interval should be one hour.
	// Specifys what time the backup action should take place. And the time interval should be one hour.
	// +kubebuilder:validation:Optional
	BackupTime *string `json:"backupTime,omitempty" tf:"backup_time,omitempty"`

	// ID of a redis instance to which the policy will be applied.
	// ID of a redis instance to which the policy will be applied.
	// +crossplane:generate:reference:type=Instance
	// +kubebuilder:validation:Optional
	RedisID *string `json:"redisId,omitempty" tf:"redis_id,omitempty"`

	// Reference to a Instance to populate redisId.
	// +kubebuilder:validation:Optional
	RedisIDRef *v1.Reference `json:"redisIdRef,omitempty" tf:"-"`

	// Selector for a Instance to populate redisId.
	// +kubebuilder:validation:Optional
	RedisIDSelector *v1.Selector `json:"redisIdSelector,omitempty" tf:"-"`
}

// BackupConfigSpec defines the desired state of BackupConfig
type BackupConfigSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BackupConfigParameters `json:"forProvider"`
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
	InitProvider BackupConfigInitParameters `json:"initProvider,omitempty"`
}

// BackupConfigStatus defines the observed state of BackupConfig.
type BackupConfigStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BackupConfigObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BackupConfig is the Schema for the BackupConfigs API. Use this resource to create a backup config of redis.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloud}
type BackupConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.backupTime) || (has(self.initProvider) && has(self.initProvider.backupTime))",message="spec.forProvider.backupTime is a required parameter"
	Spec   BackupConfigSpec   `json:"spec"`
	Status BackupConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BackupConfigList contains a list of BackupConfigs
type BackupConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BackupConfig `json:"items"`
}

// Repository type metadata.
var (
	BackupConfig_Kind             = "BackupConfig"
	BackupConfig_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BackupConfig_Kind}.String()
	BackupConfig_KindAPIVersion   = BackupConfig_Kind + "." + CRDGroupVersion.String()
	BackupConfig_GroupVersionKind = CRDGroupVersion.WithKind(BackupConfig_Kind)
)

func init() {
	SchemeBuilder.Register(&BackupConfig{}, &BackupConfigList{})
}