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

type ApplicationProxyRuleInitParameters struct {

	// Passes the client IP. Default value is OFF. When Proto is TCP, valid values: TOA: Pass the client IP via TOA; PPV1: Pass the client IP via Proxy Protocol V1; PPV2: Pass the client IP via Proxy Protocol V2; OFF: Do not pass the client IP. When Proto=UDP, valid values: PPV2: Pass the client IP via Proxy Protocol V2; OFF: Do not pass the client IP.
	// Passes the client IP. Default value is `OFF`. When Proto is TCP, valid values: `TOA`: Pass the client IP via TOA; `PPV1`: Pass the client IP via Proxy Protocol V1; `PPV2`: Pass the client IP via Proxy Protocol V2; `OFF`: Do not pass the client IP. When Proto=UDP, valid values: `PPV2`: Pass the client IP via Proxy Protocol V2; `OFF`: Do not pass the client IP.
	ForwardClientIP *string `json:"forwardClientIp,omitempty" tf:"forward_client_ip,omitempty"`

	// Origin port, supported formats: single port: 80; Port segment: 81-90, 81 to 90 ports.
	// Origin port, supported formats: single port: 80; Port segment: 81-90, 81 to 90 ports.
	OriginPort *string `json:"originPort,omitempty" tf:"origin_port,omitempty"`

	// Origin server type. Valid values: custom: Specified origins; origins: An origin group.
	// Origin server type. Valid values: `custom`: Specified origins; `origins`: An origin group.
	OriginType *string `json:"originType,omitempty" tf:"origin_type,omitempty"`

	// Origin site information: When OriginType is custom, it indicates one or more origin sites, such as ['8.8.8.8', '9.9.9.9'] or OriginValue=['test.com']; When OriginType is origins, there is required to be one and only one element, representing the origin site group ID, such as ['origin-537f5b41-162a-11ed-abaa-525400c5da15'].
	// Origin site information: When `OriginType` is `custom`, it indicates one or more origin sites, such as `['8.8.8.8', '9.9.9.9']` or `OriginValue=['test.com']`; When `OriginType` is `origins`, there is required to be one and only one element, representing the origin site group ID, such as `['origin-537f5b41-162a-11ed-abaa-525400c5da15']`.
	OriginValue []*string `json:"originValue,omitempty" tf:"origin_value,omitempty"`

	// Valid values: 80 means port 80; 81-90 means port range 81-90.
	// Valid values: `80` means port 80; `81-90` means port range 81-90.
	Port []*string `json:"port,omitempty" tf:"port,omitempty"`

	// Protocol. Valid values: TCP, UDP.
	// Protocol. Valid values: `TCP`, `UDP`.
	Proto *string `json:"proto,omitempty" tf:"proto,omitempty"`

	// Proxy ID.
	// Proxy ID.
	ProxyID *string `json:"proxyId,omitempty" tf:"proxy_id,omitempty"`

	// Specifies whether to enable session persistence. Default value is false.
	// Specifies whether to enable session persistence. Default value is false.
	SessionPersist *bool `json:"sessionPersist,omitempty" tf:"session_persist,omitempty"`

	// Status, the values are: online: enabled; offline: deactivated; progress: being deployed; stopping: being deactivated; fail: deployment failure/deactivation failure.
	// Status, the values are: `online`: enabled; `offline`: deactivated; `progress`: being deployed; `stopping`: being deactivated; `fail`: deployment failure/deactivation failure.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type ApplicationProxyRuleObservation struct {

	// Passes the client IP. Default value is OFF. When Proto is TCP, valid values: TOA: Pass the client IP via TOA; PPV1: Pass the client IP via Proxy Protocol V1; PPV2: Pass the client IP via Proxy Protocol V2; OFF: Do not pass the client IP. When Proto=UDP, valid values: PPV2: Pass the client IP via Proxy Protocol V2; OFF: Do not pass the client IP.
	// Passes the client IP. Default value is `OFF`. When Proto is TCP, valid values: `TOA`: Pass the client IP via TOA; `PPV1`: Pass the client IP via Proxy Protocol V1; `PPV2`: Pass the client IP via Proxy Protocol V2; `OFF`: Do not pass the client IP. When Proto=UDP, valid values: `PPV2`: Pass the client IP via Proxy Protocol V2; `OFF`: Do not pass the client IP.
	ForwardClientIP *string `json:"forwardClientIp,omitempty" tf:"forward_client_ip,omitempty"`

	// ID of the resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Origin port, supported formats: single port: 80; Port segment: 81-90, 81 to 90 ports.
	// Origin port, supported formats: single port: 80; Port segment: 81-90, 81 to 90 ports.
	OriginPort *string `json:"originPort,omitempty" tf:"origin_port,omitempty"`

	// Origin server type. Valid values: custom: Specified origins; origins: An origin group.
	// Origin server type. Valid values: `custom`: Specified origins; `origins`: An origin group.
	OriginType *string `json:"originType,omitempty" tf:"origin_type,omitempty"`

	// Origin site information: When OriginType is custom, it indicates one or more origin sites, such as ['8.8.8.8', '9.9.9.9'] or OriginValue=['test.com']; When OriginType is origins, there is required to be one and only one element, representing the origin site group ID, such as ['origin-537f5b41-162a-11ed-abaa-525400c5da15'].
	// Origin site information: When `OriginType` is `custom`, it indicates one or more origin sites, such as `['8.8.8.8', '9.9.9.9']` or `OriginValue=['test.com']`; When `OriginType` is `origins`, there is required to be one and only one element, representing the origin site group ID, such as `['origin-537f5b41-162a-11ed-abaa-525400c5da15']`.
	OriginValue []*string `json:"originValue,omitempty" tf:"origin_value,omitempty"`

	// Valid values: 80 means port 80; 81-90 means port range 81-90.
	// Valid values: `80` means port 80; `81-90` means port range 81-90.
	Port []*string `json:"port,omitempty" tf:"port,omitempty"`

	// Protocol. Valid values: TCP, UDP.
	// Protocol. Valid values: `TCP`, `UDP`.
	Proto *string `json:"proto,omitempty" tf:"proto,omitempty"`

	// Proxy ID.
	// Proxy ID.
	ProxyID *string `json:"proxyId,omitempty" tf:"proxy_id,omitempty"`

	// Rule ID.
	// Rule ID.
	RuleID *string `json:"ruleId,omitempty" tf:"rule_id,omitempty"`

	// Specifies whether to enable session persistence. Default value is false.
	// Specifies whether to enable session persistence. Default value is false.
	SessionPersist *bool `json:"sessionPersist,omitempty" tf:"session_persist,omitempty"`

	// Status, the values are: online: enabled; offline: deactivated; progress: being deployed; stopping: being deactivated; fail: deployment failure/deactivation failure.
	// Status, the values are: `online`: enabled; `offline`: deactivated; `progress`: being deployed; `stopping`: being deactivated; `fail`: deployment failure/deactivation failure.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Site ID.
	// Site ID.
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`
}

type ApplicationProxyRuleParameters struct {

	// Passes the client IP. Default value is OFF. When Proto is TCP, valid values: TOA: Pass the client IP via TOA; PPV1: Pass the client IP via Proxy Protocol V1; PPV2: Pass the client IP via Proxy Protocol V2; OFF: Do not pass the client IP. When Proto=UDP, valid values: PPV2: Pass the client IP via Proxy Protocol V2; OFF: Do not pass the client IP.
	// Passes the client IP. Default value is `OFF`. When Proto is TCP, valid values: `TOA`: Pass the client IP via TOA; `PPV1`: Pass the client IP via Proxy Protocol V1; `PPV2`: Pass the client IP via Proxy Protocol V2; `OFF`: Do not pass the client IP. When Proto=UDP, valid values: `PPV2`: Pass the client IP via Proxy Protocol V2; `OFF`: Do not pass the client IP.
	// +kubebuilder:validation:Optional
	ForwardClientIP *string `json:"forwardClientIp,omitempty" tf:"forward_client_ip,omitempty"`

	// Origin port, supported formats: single port: 80; Port segment: 81-90, 81 to 90 ports.
	// Origin port, supported formats: single port: 80; Port segment: 81-90, 81 to 90 ports.
	// +kubebuilder:validation:Optional
	OriginPort *string `json:"originPort,omitempty" tf:"origin_port,omitempty"`

	// Origin server type. Valid values: custom: Specified origins; origins: An origin group.
	// Origin server type. Valid values: `custom`: Specified origins; `origins`: An origin group.
	// +kubebuilder:validation:Optional
	OriginType *string `json:"originType,omitempty" tf:"origin_type,omitempty"`

	// Origin site information: When OriginType is custom, it indicates one or more origin sites, such as ['8.8.8.8', '9.9.9.9'] or OriginValue=['test.com']; When OriginType is origins, there is required to be one and only one element, representing the origin site group ID, such as ['origin-537f5b41-162a-11ed-abaa-525400c5da15'].
	// Origin site information: When `OriginType` is `custom`, it indicates one or more origin sites, such as `['8.8.8.8', '9.9.9.9']` or `OriginValue=['test.com']`; When `OriginType` is `origins`, there is required to be one and only one element, representing the origin site group ID, such as `['origin-537f5b41-162a-11ed-abaa-525400c5da15']`.
	// +kubebuilder:validation:Optional
	OriginValue []*string `json:"originValue,omitempty" tf:"origin_value,omitempty"`

	// Valid values: 80 means port 80; 81-90 means port range 81-90.
	// Valid values: `80` means port 80; `81-90` means port range 81-90.
	// +kubebuilder:validation:Optional
	Port []*string `json:"port,omitempty" tf:"port,omitempty"`

	// Protocol. Valid values: TCP, UDP.
	// Protocol. Valid values: `TCP`, `UDP`.
	// +kubebuilder:validation:Optional
	Proto *string `json:"proto,omitempty" tf:"proto,omitempty"`

	// Proxy ID.
	// Proxy ID.
	// +kubebuilder:validation:Optional
	ProxyID *string `json:"proxyId,omitempty" tf:"proxy_id,omitempty"`

	// Specifies whether to enable session persistence. Default value is false.
	// Specifies whether to enable session persistence. Default value is false.
	// +kubebuilder:validation:Optional
	SessionPersist *bool `json:"sessionPersist,omitempty" tf:"session_persist,omitempty"`

	// Status, the values are: online: enabled; offline: deactivated; progress: being deployed; stopping: being deactivated; fail: deployment failure/deactivation failure.
	// Status, the values are: `online`: enabled; `offline`: deactivated; `progress`: being deployed; `stopping`: being deactivated; `fail`: deployment failure/deactivation failure.
	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Site ID.
	// Site ID.
	// +crossplane:generate:reference:type=Zone
	// +kubebuilder:validation:Optional
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`

	// Reference to a Zone to populate zoneId.
	// +kubebuilder:validation:Optional
	ZoneIDRef *v1.Reference `json:"zoneIdRef,omitempty" tf:"-"`

	// Selector for a Zone to populate zoneId.
	// +kubebuilder:validation:Optional
	ZoneIDSelector *v1.Selector `json:"zoneIdSelector,omitempty" tf:"-"`
}

// ApplicationProxyRuleSpec defines the desired state of ApplicationProxyRule
type ApplicationProxyRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ApplicationProxyRuleParameters `json:"forProvider"`
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
	InitProvider ApplicationProxyRuleInitParameters `json:"initProvider,omitempty"`
}

// ApplicationProxyRuleStatus defines the observed state of ApplicationProxyRule.
type ApplicationProxyRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ApplicationProxyRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ApplicationProxyRule is the Schema for the ApplicationProxyRules API. Provides a resource to create a teo application_proxy_rule
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloud}
type ApplicationProxyRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.originPort) || (has(self.initProvider) && has(self.initProvider.originPort))",message="spec.forProvider.originPort is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.originType) || (has(self.initProvider) && has(self.initProvider.originType))",message="spec.forProvider.originType is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.originValue) || (has(self.initProvider) && has(self.initProvider.originValue))",message="spec.forProvider.originValue is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.port) || (has(self.initProvider) && has(self.initProvider.port))",message="spec.forProvider.port is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.proto) || (has(self.initProvider) && has(self.initProvider.proto))",message="spec.forProvider.proto is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.proxyId) || (has(self.initProvider) && has(self.initProvider.proxyId))",message="spec.forProvider.proxyId is a required parameter"
	Spec   ApplicationProxyRuleSpec   `json:"spec"`
	Status ApplicationProxyRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApplicationProxyRuleList contains a list of ApplicationProxyRules
type ApplicationProxyRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApplicationProxyRule `json:"items"`
}

// Repository type metadata.
var (
	ApplicationProxyRule_Kind             = "ApplicationProxyRule"
	ApplicationProxyRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ApplicationProxyRule_Kind}.String()
	ApplicationProxyRule_KindAPIVersion   = ApplicationProxyRule_Kind + "." + CRDGroupVersion.String()
	ApplicationProxyRule_GroupVersionKind = CRDGroupVersion.WithKind(ApplicationProxyRule_Kind)
)

func init() {
	SchemeBuilder.Register(&ApplicationProxyRule{}, &ApplicationProxyRuleList{})
}