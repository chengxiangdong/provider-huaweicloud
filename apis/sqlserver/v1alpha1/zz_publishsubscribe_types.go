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

type DatabaseTuplesInitParameters struct {

	// Publish the database.
	// Publish the database.
	PublishDatabase *string `json:"publishDatabase,omitempty" tf:"publish_database,omitempty"`

	// Subscribe the database.
	// Subscribe the database.
	SubscribeDatabase *string `json:"subscribeDatabase,omitempty" tf:"subscribe_database,omitempty"`
}

type DatabaseTuplesObservation struct {

	// Publish the database.
	// Publish the database.
	PublishDatabase *string `json:"publishDatabase,omitempty" tf:"publish_database,omitempty"`

	// Subscribe the database.
	// Subscribe the database.
	SubscribeDatabase *string `json:"subscribeDatabase,omitempty" tf:"subscribe_database,omitempty"`
}

type DatabaseTuplesParameters struct {

	// Publish the database.
	// Publish the database.
	// +kubebuilder:validation:Optional
	PublishDatabase *string `json:"publishDatabase" tf:"publish_database,omitempty"`

	// Subscribe the database.
	// Subscribe the database.
	// +kubebuilder:validation:Optional
	SubscribeDatabase *string `json:"subscribeDatabase" tf:"subscribe_database,omitempty"`
}

type PublishSubscribeInitParameters struct {

	// Database Publish and Publish relationship list. The elements inside can be deleted and added individually, but modification is not allowed.
	// Database Publish and Publish relationship list. The elements inside can be deleted and added individually, but modification is not allowed.
	DatabaseTuples []DatabaseTuplesInitParameters `json:"databaseTuples,omitempty" tf:"database_tuples,omitempty"`

	// Whether to delete the subscriber database when deleting the Publish and Subscribe. true for deletes the subscribe database, false for does not delete the subscribe database. default is false.
	// Whether to delete the subscriber database when deleting the Publish and Subscribe. `true` for deletes the subscribe database, `false` for does not delete the subscribe database. default is `false`.
	DeleteSubscribeDB *bool `json:"deleteSubscribeDb,omitempty" tf:"delete_subscribe_db,omitempty"`

	// ID of the SQL Server instance which publish.
	// ID of the SQL Server instance which publish.
	PublishInstanceID *string `json:"publishInstanceId,omitempty" tf:"publish_instance_id,omitempty"`

	// The name of the Publish and Subscribe. Default is default_name.
	// The name of the Publish and Subscribe. Default is `default_name`.
	PublishSubscribeName *string `json:"publishSubscribeName,omitempty" tf:"publish_subscribe_name,omitempty"`

	// ID of the SQL Server instance which subscribe.
	// ID of the SQL Server instance which subscribe.
	SubscribeInstanceID *string `json:"subscribeInstanceId,omitempty" tf:"subscribe_instance_id,omitempty"`
}

type PublishSubscribeObservation struct {

	// Database Publish and Publish relationship list. The elements inside can be deleted and added individually, but modification is not allowed.
	// Database Publish and Publish relationship list. The elements inside can be deleted and added individually, but modification is not allowed.
	DatabaseTuples []DatabaseTuplesObservation `json:"databaseTuples,omitempty" tf:"database_tuples,omitempty"`

	// Whether to delete the subscriber database when deleting the Publish and Subscribe. true for deletes the subscribe database, false for does not delete the subscribe database. default is false.
	// Whether to delete the subscriber database when deleting the Publish and Subscribe. `true` for deletes the subscribe database, `false` for does not delete the subscribe database. default is `false`.
	DeleteSubscribeDB *bool `json:"deleteSubscribeDb,omitempty" tf:"delete_subscribe_db,omitempty"`

	// ID of the resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// ID of the SQL Server instance which publish.
	// ID of the SQL Server instance which publish.
	PublishInstanceID *string `json:"publishInstanceId,omitempty" tf:"publish_instance_id,omitempty"`

	// The name of the Publish and Subscribe. Default is default_name.
	// The name of the Publish and Subscribe. Default is `default_name`.
	PublishSubscribeName *string `json:"publishSubscribeName,omitempty" tf:"publish_subscribe_name,omitempty"`

	// ID of the SQL Server instance which subscribe.
	// ID of the SQL Server instance which subscribe.
	SubscribeInstanceID *string `json:"subscribeInstanceId,omitempty" tf:"subscribe_instance_id,omitempty"`
}

type PublishSubscribeParameters struct {

	// Database Publish and Publish relationship list. The elements inside can be deleted and added individually, but modification is not allowed.
	// Database Publish and Publish relationship list. The elements inside can be deleted and added individually, but modification is not allowed.
	// +kubebuilder:validation:Optional
	DatabaseTuples []DatabaseTuplesParameters `json:"databaseTuples,omitempty" tf:"database_tuples,omitempty"`

	// Whether to delete the subscriber database when deleting the Publish and Subscribe. true for deletes the subscribe database, false for does not delete the subscribe database. default is false.
	// Whether to delete the subscriber database when deleting the Publish and Subscribe. `true` for deletes the subscribe database, `false` for does not delete the subscribe database. default is `false`.
	// +kubebuilder:validation:Optional
	DeleteSubscribeDB *bool `json:"deleteSubscribeDb,omitempty" tf:"delete_subscribe_db,omitempty"`

	// ID of the SQL Server instance which publish.
	// ID of the SQL Server instance which publish.
	// +kubebuilder:validation:Optional
	PublishInstanceID *string `json:"publishInstanceId,omitempty" tf:"publish_instance_id,omitempty"`

	// The name of the Publish and Subscribe. Default is default_name.
	// The name of the Publish and Subscribe. Default is `default_name`.
	// +kubebuilder:validation:Optional
	PublishSubscribeName *string `json:"publishSubscribeName,omitempty" tf:"publish_subscribe_name,omitempty"`

	// ID of the SQL Server instance which subscribe.
	// ID of the SQL Server instance which subscribe.
	// +kubebuilder:validation:Optional
	SubscribeInstanceID *string `json:"subscribeInstanceId,omitempty" tf:"subscribe_instance_id,omitempty"`
}

// PublishSubscribeSpec defines the desired state of PublishSubscribe
type PublishSubscribeSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PublishSubscribeParameters `json:"forProvider"`
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
	InitProvider PublishSubscribeInitParameters `json:"initProvider,omitempty"`
}

// PublishSubscribeStatus defines the observed state of PublishSubscribe.
type PublishSubscribeStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PublishSubscribeObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PublishSubscribe is the Schema for the PublishSubscribes API. Provides a SQL Server PublishSubscribe resource belongs to SQL Server instance.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloud}
type PublishSubscribe struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.databaseTuples) || (has(self.initProvider) && has(self.initProvider.databaseTuples))",message="spec.forProvider.databaseTuples is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.publishInstanceId) || (has(self.initProvider) && has(self.initProvider.publishInstanceId))",message="spec.forProvider.publishInstanceId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.subscribeInstanceId) || (has(self.initProvider) && has(self.initProvider.subscribeInstanceId))",message="spec.forProvider.subscribeInstanceId is a required parameter"
	Spec   PublishSubscribeSpec   `json:"spec"`
	Status PublishSubscribeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PublishSubscribeList contains a list of PublishSubscribes
type PublishSubscribeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PublishSubscribe `json:"items"`
}

// Repository type metadata.
var (
	PublishSubscribe_Kind             = "PublishSubscribe"
	PublishSubscribe_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PublishSubscribe_Kind}.String()
	PublishSubscribe_KindAPIVersion   = PublishSubscribe_Kind + "." + CRDGroupVersion.String()
	PublishSubscribe_GroupVersionKind = CRDGroupVersion.WithKind(PublishSubscribe_Kind)
)

func init() {
	SchemeBuilder.Register(&PublishSubscribe{}, &PublishSubscribeList{})
}