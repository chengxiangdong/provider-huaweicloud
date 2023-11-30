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

type ImageSpriteTemplateObservation struct {
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time,omitempty"`
}

type ImageSpriteTemplateParameters struct {

	// Subimage column count of an image sprite.
	// +kubebuilder:validation:Required
	ColumnCount *float64 `json:"columnCount" tf:"column_count,omitempty"`

	// Template description. Length limit: 256 characters.
	// +kubebuilder:validation:Optional
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// Fill refers to the way of processing a screenshot when its aspect ratio is different from that of the source video. The following fill types are supported: `stretch`: stretch. The screenshot will be stretched frame by frame to match the aspect ratio of the source video, which may make the screenshot shorter or longer; `black`: fill with black. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with black color blocks. Default value: `black`.
	// +kubebuilder:validation:Optional
	FillType *string `json:"fillType,omitempty" tf:"fill_type,omitempty"`

	// Maximum value of the `height` (or short side) of a screenshot in px. Value range: 0 and [128, 4,096]. If both `width` and `height` are `0`, the resolution will be the same as that of the source video; If `width` is `0`, but `height` is not `0`, `width` will be proportionally scaled; If `width` is not `0`, but `height` is `0`, `height` will be proportionally scaled; If both `width` and `height` are not `0`, the custom resolution will be used. Default value: `0`.
	// +kubebuilder:validation:Optional
	Height *float64 `json:"height,omitempty" tf:"height,omitempty"`

	// Name of a time point screen capturing template. Length limit: 64 characters.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Resolution adaption. Valid values: `true`,`false`. `true`: enabled. In this case, `width` represents the long side of a video, while `height` the short side; `false`: disabled. In this case, `width` represents the width of a video, while `height` the height. Default value: `true`.
	// +kubebuilder:validation:Optional
	ResolutionAdaptive *bool `json:"resolutionAdaptive,omitempty" tf:"resolution_adaptive,omitempty"`

	// Subimage row count of an image sprite.
	// +kubebuilder:validation:Required
	RowCount *float64 `json:"rowCount" tf:"row_count,omitempty"`

	// Sampling interval. If `sample_type` is `Percent`, sampling will be performed at an interval of the specified percentage. If `sample_type` is `Time`, sampling will be performed at the specified time interval in seconds.
	// +kubebuilder:validation:Required
	SampleInterval *float64 `json:"sampleInterval" tf:"sample_interval,omitempty"`

	// Sampling type. Valid values: `Percent`, `Time`. `Percent`: by percent. `Time`: by time interval.
	// +kubebuilder:validation:Required
	SampleType *string `json:"sampleType" tf:"sample_type,omitempty"`

	// Subapplication ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
	// +kubebuilder:validation:Optional
	SubAppID *float64 `json:"subAppId,omitempty" tf:"sub_app_id,omitempty"`

	// Maximum value of the `width` (or long side) of a screenshot in px. Value range: 0 and [128, 4,096]. If both `width` and `height` are `0`, the resolution will be the same as that of the source video; If `width` is `0`, but `height` is not `0`, width will be proportionally scaled; If `width` is not `0`, but `height` is `0`, `height` will be proportionally scaled; If both `width` and `height` are not `0`, the custom resolution will be used. Default value: `0`.
	// +kubebuilder:validation:Optional
	Width *float64 `json:"width,omitempty" tf:"width,omitempty"`
}

// ImageSpriteTemplateSpec defines the desired state of ImageSpriteTemplate
type ImageSpriteTemplateSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ImageSpriteTemplateParameters `json:"forProvider"`
}

// ImageSpriteTemplateStatus defines the observed state of ImageSpriteTemplate.
type ImageSpriteTemplateStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ImageSpriteTemplateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ImageSpriteTemplate is the Schema for the ImageSpriteTemplates API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloudjet}
type ImageSpriteTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ImageSpriteTemplateSpec   `json:"spec"`
	Status            ImageSpriteTemplateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ImageSpriteTemplateList contains a list of ImageSpriteTemplates
type ImageSpriteTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ImageSpriteTemplate `json:"items"`
}

// Repository type metadata.
var (
	ImageSpriteTemplate_Kind             = "ImageSpriteTemplate"
	ImageSpriteTemplate_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ImageSpriteTemplate_Kind}.String()
	ImageSpriteTemplate_KindAPIVersion   = ImageSpriteTemplate_Kind + "." + CRDGroupVersion.String()
	ImageSpriteTemplate_GroupVersionKind = CRDGroupVersion.WithKind(ImageSpriteTemplate_Kind)
)

func init() {
	SchemeBuilder.Register(&ImageSpriteTemplate{}, &ImageSpriteTemplateList{})
}