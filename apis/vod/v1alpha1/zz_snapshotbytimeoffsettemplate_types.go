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

type SnapshotByTimeOffsetTemplateInitParameters struct {

	// Template description. Length limit: 256 characters.
	// Template description. Length limit: 256 characters.
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// Fill refers to the way of processing a screenshot when its aspect ratio is different from that of the source video. The following fill types are supported: stretch: stretch. The screenshot will be stretched frame by frame to match the aspect ratio of the source video, which may make the screenshot shorter or longer; black: fill with black. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with black color blocks. white: fill with white. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with white color blocks. gauss: fill with Gaussian blur. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with Gaussian blur. Default value: black.
	// Fill refers to the way of processing a screenshot when its aspect ratio is different from that of the source video. The following fill types are supported: `stretch`: stretch. The screenshot will be stretched frame by frame to match the aspect ratio of the source video, which may make the screenshot `shorter` or `longer`; `black`: fill with black. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with black color blocks. `white`: fill with white. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with white color blocks. `gauss`: fill with Gaussian blur. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with Gaussian blur. Default value: `black`.
	FillType *string `json:"fillType,omitempty" tf:"fill_type,omitempty"`

	// Image format. Valid values: jpg, png. Default value: jpg.
	// Image format. Valid values: `jpg`, `png`. Default value: `jpg`.
	Format *string `json:"format,omitempty" tf:"format,omitempty"`

	// Maximum value of the height (or short side) of a screenshot in px. Value range: 0 and [128, 4,096]. If both width and height are 0, the resolution will be the same as that of the source video; If width is 0, but height is not 0, width will be proportionally scaled; If width is not 0, but height is 0, height will be proportionally scaled; If both width and height are not 0, the custom resolution will be used. Default value: 0.
	// Maximum value of the `height` (or short side) of a screenshot in px. Value range: 0 and [128, 4,096]. If both `width` and `height` are `0`, the resolution will be the same as that of the source video; If `width` is `0`, but `height` is not `0`, `width` will be proportionally scaled; If `width` is not `0`, but `height` is `0`, `height` will be proportionally scaled; If both `width` and `height` are not `0`, the custom resolution will be used. Default value: `0`.
	Height *float64 `json:"height,omitempty" tf:"height,omitempty"`

	// Name of a time point screen capturing template. Length limit: 64 characters.
	// Name of a time point screen capturing template. Length limit: 64 characters.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Resolution adaption. Valid values: true,false. true: enabled. In this case, width represents the long side of a video, while height the short side; false: disabled. In this case, width represents the width of a video, while height the height. Default value: true.
	// Resolution adaption. Valid values: `true`,`false`. `true`: enabled. In this case, `width` represents the long side of a video, while `height` the short side; `false`: disabled. In this case, `width` represents the width of a video, while `height` the height. Default value: `true`.
	ResolutionAdaptive *bool `json:"resolutionAdaptive,omitempty" tf:"resolution_adaptive,omitempty"`

	// The VOD application ID. For customers who activate VOD service from December 25, 2023, if they want to access resources in a VOD application (whether it's the default application or a newly created one), they must fill in this field with the application ID.
	// The VOD [application](https://intl.cloud.tencent.com/document/product/266/14574) ID. For customers who activate VOD service from December 25, 2023, if they want to access resources in a VOD application (whether it's the default application or a newly created one), they must fill in this field with the application ID.
	SubAppID *float64 `json:"subAppId,omitempty" tf:"sub_app_id,omitempty"`

	// Maximum value of the width (or long side) of a screenshot in px. Value range: 0 and [128, 4,096]. If both width and height are 0, the resolution will be the same as that of the source video; If width is 0, but height is not 0, width will be proportionally scaled; If width is not 0, but height is 0, height will be proportionally scaled; If both width and height are not 0, the custom resolution will be used. Default value: 0.
	// Maximum value of the `width` (or long side) of a screenshot in px. Value range: 0 and [128, 4,096]. If both `width` and `height` are `0`, the resolution will be the same as that of the source video; If `width` is `0`, but `height` is not `0`, width will be proportionally scaled; If `width` is not `0`, but `height` is `0`, `height` will be proportionally scaled; If both `width` and `height` are not `0`, the custom resolution will be used. Default value: `0`.
	Width *float64 `json:"width,omitempty" tf:"width,omitempty"`
}

type SnapshotByTimeOffsetTemplateObservation struct {

	// Template description. Length limit: 256 characters.
	// Template description. Length limit: 256 characters.
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// Creation time of template in ISO date format.
	// Creation time of template in ISO date format.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// Fill refers to the way of processing a screenshot when its aspect ratio is different from that of the source video. The following fill types are supported: stretch: stretch. The screenshot will be stretched frame by frame to match the aspect ratio of the source video, which may make the screenshot shorter or longer; black: fill with black. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with black color blocks. white: fill with white. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with white color blocks. gauss: fill with Gaussian blur. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with Gaussian blur. Default value: black.
	// Fill refers to the way of processing a screenshot when its aspect ratio is different from that of the source video. The following fill types are supported: `stretch`: stretch. The screenshot will be stretched frame by frame to match the aspect ratio of the source video, which may make the screenshot `shorter` or `longer`; `black`: fill with black. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with black color blocks. `white`: fill with white. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with white color blocks. `gauss`: fill with Gaussian blur. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with Gaussian blur. Default value: `black`.
	FillType *string `json:"fillType,omitempty" tf:"fill_type,omitempty"`

	// Image format. Valid values: jpg, png. Default value: jpg.
	// Image format. Valid values: `jpg`, `png`. Default value: `jpg`.
	Format *string `json:"format,omitempty" tf:"format,omitempty"`

	// Maximum value of the height (or short side) of a screenshot in px. Value range: 0 and [128, 4,096]. If both width and height are 0, the resolution will be the same as that of the source video; If width is 0, but height is not 0, width will be proportionally scaled; If width is not 0, but height is 0, height will be proportionally scaled; If both width and height are not 0, the custom resolution will be used. Default value: 0.
	// Maximum value of the `height` (or short side) of a screenshot in px. Value range: 0 and [128, 4,096]. If both `width` and `height` are `0`, the resolution will be the same as that of the source video; If `width` is `0`, but `height` is not `0`, `width` will be proportionally scaled; If `width` is not `0`, but `height` is `0`, `height` will be proportionally scaled; If both `width` and `height` are not `0`, the custom resolution will be used. Default value: `0`.
	Height *float64 `json:"height,omitempty" tf:"height,omitempty"`

	// ID of the resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Name of a time point screen capturing template. Length limit: 64 characters.
	// Name of a time point screen capturing template. Length limit: 64 characters.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Resolution adaption. Valid values: true,false. true: enabled. In this case, width represents the long side of a video, while height the short side; false: disabled. In this case, width represents the width of a video, while height the height. Default value: true.
	// Resolution adaption. Valid values: `true`,`false`. `true`: enabled. In this case, `width` represents the long side of a video, while `height` the short side; `false`: disabled. In this case, `width` represents the width of a video, while `height` the height. Default value: `true`.
	ResolutionAdaptive *bool `json:"resolutionAdaptive,omitempty" tf:"resolution_adaptive,omitempty"`

	// The VOD application ID. For customers who activate VOD service from December 25, 2023, if they want to access resources in a VOD application (whether it's the default application or a newly created one), they must fill in this field with the application ID.
	// The VOD [application](https://intl.cloud.tencent.com/document/product/266/14574) ID. For customers who activate VOD service from December 25, 2023, if they want to access resources in a VOD application (whether it's the default application or a newly created one), they must fill in this field with the application ID.
	SubAppID *float64 `json:"subAppId,omitempty" tf:"sub_app_id,omitempty"`

	// Template type, value range:
	// Template type, value range:
	// - Preset: system preset template;
	// - Custom: user-defined templates.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Last modified time of template in ISO date format.
	// Last modified time of template in ISO date format.
	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time,omitempty"`

	// Maximum value of the width (or long side) of a screenshot in px. Value range: 0 and [128, 4,096]. If both width and height are 0, the resolution will be the same as that of the source video; If width is 0, but height is not 0, width will be proportionally scaled; If width is not 0, but height is 0, height will be proportionally scaled; If both width and height are not 0, the custom resolution will be used. Default value: 0.
	// Maximum value of the `width` (or long side) of a screenshot in px. Value range: 0 and [128, 4,096]. If both `width` and `height` are `0`, the resolution will be the same as that of the source video; If `width` is `0`, but `height` is not `0`, width will be proportionally scaled; If `width` is not `0`, but `height` is `0`, `height` will be proportionally scaled; If both `width` and `height` are not `0`, the custom resolution will be used. Default value: `0`.
	Width *float64 `json:"width,omitempty" tf:"width,omitempty"`
}

type SnapshotByTimeOffsetTemplateParameters struct {

	// Template description. Length limit: 256 characters.
	// Template description. Length limit: 256 characters.
	// +kubebuilder:validation:Optional
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// Fill refers to the way of processing a screenshot when its aspect ratio is different from that of the source video. The following fill types are supported: stretch: stretch. The screenshot will be stretched frame by frame to match the aspect ratio of the source video, which may make the screenshot shorter or longer; black: fill with black. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with black color blocks. white: fill with white. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with white color blocks. gauss: fill with Gaussian blur. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with Gaussian blur. Default value: black.
	// Fill refers to the way of processing a screenshot when its aspect ratio is different from that of the source video. The following fill types are supported: `stretch`: stretch. The screenshot will be stretched frame by frame to match the aspect ratio of the source video, which may make the screenshot `shorter` or `longer`; `black`: fill with black. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with black color blocks. `white`: fill with white. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with white color blocks. `gauss`: fill with Gaussian blur. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with Gaussian blur. Default value: `black`.
	// +kubebuilder:validation:Optional
	FillType *string `json:"fillType,omitempty" tf:"fill_type,omitempty"`

	// Image format. Valid values: jpg, png. Default value: jpg.
	// Image format. Valid values: `jpg`, `png`. Default value: `jpg`.
	// +kubebuilder:validation:Optional
	Format *string `json:"format,omitempty" tf:"format,omitempty"`

	// Maximum value of the height (or short side) of a screenshot in px. Value range: 0 and [128, 4,096]. If both width and height are 0, the resolution will be the same as that of the source video; If width is 0, but height is not 0, width will be proportionally scaled; If width is not 0, but height is 0, height will be proportionally scaled; If both width and height are not 0, the custom resolution will be used. Default value: 0.
	// Maximum value of the `height` (or short side) of a screenshot in px. Value range: 0 and [128, 4,096]. If both `width` and `height` are `0`, the resolution will be the same as that of the source video; If `width` is `0`, but `height` is not `0`, `width` will be proportionally scaled; If `width` is not `0`, but `height` is `0`, `height` will be proportionally scaled; If both `width` and `height` are not `0`, the custom resolution will be used. Default value: `0`.
	// +kubebuilder:validation:Optional
	Height *float64 `json:"height,omitempty" tf:"height,omitempty"`

	// Name of a time point screen capturing template. Length limit: 64 characters.
	// Name of a time point screen capturing template. Length limit: 64 characters.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Resolution adaption. Valid values: true,false. true: enabled. In this case, width represents the long side of a video, while height the short side; false: disabled. In this case, width represents the width of a video, while height the height. Default value: true.
	// Resolution adaption. Valid values: `true`,`false`. `true`: enabled. In this case, `width` represents the long side of a video, while `height` the short side; `false`: disabled. In this case, `width` represents the width of a video, while `height` the height. Default value: `true`.
	// +kubebuilder:validation:Optional
	ResolutionAdaptive *bool `json:"resolutionAdaptive,omitempty" tf:"resolution_adaptive,omitempty"`

	// The VOD application ID. For customers who activate VOD service from December 25, 2023, if they want to access resources in a VOD application (whether it's the default application or a newly created one), they must fill in this field with the application ID.
	// The VOD [application](https://intl.cloud.tencent.com/document/product/266/14574) ID. For customers who activate VOD service from December 25, 2023, if they want to access resources in a VOD application (whether it's the default application or a newly created one), they must fill in this field with the application ID.
	// +kubebuilder:validation:Optional
	SubAppID *float64 `json:"subAppId,omitempty" tf:"sub_app_id,omitempty"`

	// Maximum value of the width (or long side) of a screenshot in px. Value range: 0 and [128, 4,096]. If both width and height are 0, the resolution will be the same as that of the source video; If width is 0, but height is not 0, width will be proportionally scaled; If width is not 0, but height is 0, height will be proportionally scaled; If both width and height are not 0, the custom resolution will be used. Default value: 0.
	// Maximum value of the `width` (or long side) of a screenshot in px. Value range: 0 and [128, 4,096]. If both `width` and `height` are `0`, the resolution will be the same as that of the source video; If `width` is `0`, but `height` is not `0`, width will be proportionally scaled; If `width` is not `0`, but `height` is `0`, `height` will be proportionally scaled; If both `width` and `height` are not `0`, the custom resolution will be used. Default value: `0`.
	// +kubebuilder:validation:Optional
	Width *float64 `json:"width,omitempty" tf:"width,omitempty"`
}

// SnapshotByTimeOffsetTemplateSpec defines the desired state of SnapshotByTimeOffsetTemplate
type SnapshotByTimeOffsetTemplateSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SnapshotByTimeOffsetTemplateParameters `json:"forProvider"`
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
	InitProvider SnapshotByTimeOffsetTemplateInitParameters `json:"initProvider,omitempty"`
}

// SnapshotByTimeOffsetTemplateStatus defines the observed state of SnapshotByTimeOffsetTemplate.
type SnapshotByTimeOffsetTemplateStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SnapshotByTimeOffsetTemplateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SnapshotByTimeOffsetTemplate is the Schema for the SnapshotByTimeOffsetTemplates API. Provide a resource to create a VOD snapshot by time offset template.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloud}
type SnapshotByTimeOffsetTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   SnapshotByTimeOffsetTemplateSpec   `json:"spec"`
	Status SnapshotByTimeOffsetTemplateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SnapshotByTimeOffsetTemplateList contains a list of SnapshotByTimeOffsetTemplates
type SnapshotByTimeOffsetTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SnapshotByTimeOffsetTemplate `json:"items"`
}

// Repository type metadata.
var (
	SnapshotByTimeOffsetTemplate_Kind             = "SnapshotByTimeOffsetTemplate"
	SnapshotByTimeOffsetTemplate_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SnapshotByTimeOffsetTemplate_Kind}.String()
	SnapshotByTimeOffsetTemplate_KindAPIVersion   = SnapshotByTimeOffsetTemplate_Kind + "." + CRDGroupVersion.String()
	SnapshotByTimeOffsetTemplate_GroupVersionKind = CRDGroupVersion.WithKind(SnapshotByTimeOffsetTemplate_Kind)
)

func init() {
	SchemeBuilder.Register(&SnapshotByTimeOffsetTemplate{}, &SnapshotByTimeOffsetTemplateList{})
}