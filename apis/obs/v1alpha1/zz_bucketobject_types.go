// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type BucketObjectInitParameters struct {

	// The ACL policy to apply. Defaults to private.
	ACL *string `json:"acl,omitempty" tf:"acl,omitempty"`

	// The name of the bucket to put the file in.
	// +crossplane:generate:reference:type=github.com/huaweicloud/provider-huaweicloud/apis/obs/v1alpha1.Bucket
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("bucket",true)
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Reference to a Bucket in obs to populate bucket.
	// +kubebuilder:validation:Optional
	BucketRef *v1.Reference `json:"bucketRef,omitempty" tf:"-"`

	// Selector for a Bucket in obs to populate bucket.
	// +kubebuilder:validation:Optional
	BucketSelector *v1.Selector `json:"bucketSelector,omitempty" tf:"-"`

	// The literal content being uploaded to the bucket.
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// A standard MIME type describing the format of the object data, e.g.
	// application/octet-stream. All Valid MIME Types are valid for this input.
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`

	// Whether enable server-side encryption of the object in SSE-KMS mode.
	Encryption *bool `json:"encryption,omitempty" tf:"encryption,omitempty"`

	// Specifies the unique identifier of the object content. It can be used to trigger updates.
	// The only meaningful value is md5(file("path_to_file")).
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	// The ID of the kms key. If omitted, the default master key will be used.
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// The name of the object once it is in the bucket.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// The region in which to create the OBS bucket object resource. If omitted, the
	// provider-level region will be used. Changing this creates a new OBS bucket object resource.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The path to the source file being uploaded to the bucket.
	Source *string `json:"source,omitempty" tf:"source,omitempty"`

	// Specifies the storage class of the object. Defaults to STANDARD.
	StorageClass *string `json:"storageClass,omitempty" tf:"storage_class,omitempty"`
}

type BucketObjectObservation struct {

	// The ACL policy to apply. Defaults to private.
	ACL *string `json:"acl,omitempty" tf:"acl,omitempty"`

	// The name of the bucket to put the file in.
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// The literal content being uploaded to the bucket.
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// A standard MIME type describing the format of the object data, e.g.
	// application/octet-stream. All Valid MIME Types are valid for this input.
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`

	// Whether enable server-side encryption of the object in SSE-KMS mode.
	Encryption *bool `json:"encryption,omitempty" tf:"encryption,omitempty"`

	// Specifies the unique identifier of the object content. It can be used to trigger updates.
	// The only meaningful value is md5(file("path_to_file")).
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	// the key of the resource supplied above.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the kms key. If omitted, the default master key will be used.
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// The name of the object once it is in the bucket.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// The region in which to create the OBS bucket object resource. If omitted, the
	// provider-level region will be used. Changing this creates a new OBS bucket object resource.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// the size of the object in bytes.
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`

	// The path to the source file being uploaded to the bucket.
	Source *string `json:"source,omitempty" tf:"source,omitempty"`

	// Specifies the storage class of the object. Defaults to STANDARD.
	StorageClass *string `json:"storageClass,omitempty" tf:"storage_class,omitempty"`

	// A unique version ID value for the object, if bucket versioning is enabled.
	VersionID *string `json:"versionId,omitempty" tf:"version_id,omitempty"`
}

type BucketObjectParameters struct {

	// The ACL policy to apply. Defaults to private.
	// +kubebuilder:validation:Optional
	ACL *string `json:"acl,omitempty" tf:"acl,omitempty"`

	// The name of the bucket to put the file in.
	// +crossplane:generate:reference:type=github.com/huaweicloud/provider-huaweicloud/apis/obs/v1alpha1.Bucket
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("bucket",true)
	// +kubebuilder:validation:Optional
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Reference to a Bucket in obs to populate bucket.
	// +kubebuilder:validation:Optional
	BucketRef *v1.Reference `json:"bucketRef,omitempty" tf:"-"`

	// Selector for a Bucket in obs to populate bucket.
	// +kubebuilder:validation:Optional
	BucketSelector *v1.Selector `json:"bucketSelector,omitempty" tf:"-"`

	// The literal content being uploaded to the bucket.
	// +kubebuilder:validation:Optional
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// A standard MIME type describing the format of the object data, e.g.
	// application/octet-stream. All Valid MIME Types are valid for this input.
	// +kubebuilder:validation:Optional
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`

	// Whether enable server-side encryption of the object in SSE-KMS mode.
	// +kubebuilder:validation:Optional
	Encryption *bool `json:"encryption,omitempty" tf:"encryption,omitempty"`

	// Specifies the unique identifier of the object content. It can be used to trigger updates.
	// The only meaningful value is md5(file("path_to_file")).
	// +kubebuilder:validation:Optional
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	// The ID of the kms key. If omitted, the default master key will be used.
	// +kubebuilder:validation:Optional
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// The name of the object once it is in the bucket.
	// +kubebuilder:validation:Optional
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// The region in which to create the OBS bucket object resource. If omitted, the
	// provider-level region will be used. Changing this creates a new OBS bucket object resource.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The path to the source file being uploaded to the bucket.
	// +kubebuilder:validation:Optional
	Source *string `json:"source,omitempty" tf:"source,omitempty"`

	// Specifies the storage class of the object. Defaults to STANDARD.
	// +kubebuilder:validation:Optional
	StorageClass *string `json:"storageClass,omitempty" tf:"storage_class,omitempty"`
}

// BucketObjectSpec defines the desired state of BucketObject
type BucketObjectSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BucketObjectParameters `json:"forProvider"`
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
	InitProvider BucketObjectInitParameters `json:"initProvider,omitempty"`
}

// BucketObjectStatus defines the observed state of BucketObject.
type BucketObjectStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BucketObjectObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// BucketObject is the Schema for the BucketObjects API. ""
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,huaweicloud}
type BucketObject struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.key) || (has(self.initProvider) && has(self.initProvider.key))",message="spec.forProvider.key is a required parameter"
	Spec   BucketObjectSpec   `json:"spec"`
	Status BucketObjectStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BucketObjectList contains a list of BucketObjects
type BucketObjectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BucketObject `json:"items"`
}

// Repository type metadata.
var (
	BucketObject_Kind             = "BucketObject"
	BucketObject_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BucketObject_Kind}.String()
	BucketObject_KindAPIVersion   = BucketObject_Kind + "." + CRDGroupVersion.String()
	BucketObject_GroupVersionKind = CRDGroupVersion.WithKind(BucketObject_Kind)
)

func init() {
	SchemeBuilder.Register(&BucketObject{}, &BucketObjectList{})
}
