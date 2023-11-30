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

type DdosPolicyObservation struct {
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	PolicyID *string `json:"policyId,omitempty" tf:"policy_id,omitempty"`

	SceneID *string `json:"sceneId,omitempty" tf:"scene_id,omitempty"`

	WatermarkKey []WatermarkKeyObservation `json:"watermarkKey,omitempty" tf:"watermark_key,omitempty"`
}

type DdosPolicyParameters struct {

	// Black IP list.
	// +kubebuilder:validation:Optional
	BlackIps []*string `json:"blackIps,omitempty" tf:"black_ips,omitempty"`

	// Option list of abnormal check of the DDos policy, should set at least one policy.
	// +kubebuilder:validation:Required
	DropOptions []DropOptionsParameters `json:"dropOptions" tf:"drop_options,omitempty"`

	// Message filter options list.
	// +kubebuilder:validation:Optional
	PacketFilters []PacketFiltersParameters `json:"packetFilters,omitempty" tf:"packet_filters,omitempty"`

	// Port limits of abnormal check of the DDos policy.
	// +kubebuilder:validation:Optional
	PortFilters []PortFiltersParameters `json:"portFilters,omitempty" tf:"port_filters,omitempty"`

	// Type of the resource that the DDoS policy works for. Valid values: `bgpip`, `bgp`, `bgp-multip` and `net`.
	// +kubebuilder:validation:Required
	ResourceType *string `json:"resourceType" tf:"resource_type,omitempty"`

	// Watermark policy options, and only support one watermark policy at most.
	// +kubebuilder:validation:Optional
	WatermarkFilters []WatermarkFiltersParameters `json:"watermarkFilters,omitempty" tf:"watermark_filters,omitempty"`

	// White IP list.
	// +kubebuilder:validation:Optional
	WhiteIps []*string `json:"whiteIps,omitempty" tf:"white_ips,omitempty"`
}

type DropOptionsObservation struct {
}

type DropOptionsParameters struct {

	// The number of new connections based on destination IP that trigger suppression of connections. Valid value ranges: (0~4294967295).
	// +kubebuilder:validation:Required
	BadConnThreshold *float64 `json:"badConnThreshold" tf:"bad_conn_threshold,omitempty"`

	// Indicate whether to check null connection or not.
	// +kubebuilder:validation:Required
	CheckSyncConn *bool `json:"checkSyncConn" tf:"check_sync_conn,omitempty"`

	// Connection timeout of abnormal connection check. Valid value ranges: (0~65535).
	// +kubebuilder:validation:Required
	ConnTimeout *float64 `json:"connTimeout" tf:"conn_timeout,omitempty"`

	// The limit of concurrent connections based on destination IP. Valid value ranges: (0~4294967295).
	// +kubebuilder:validation:Required
	DConnLimit *float64 `json:"dConnLimit" tf:"d_conn_limit,omitempty"`

	// The limit of new connections based on destination IP. Valid value ranges: (0~4294967295).
	// +kubebuilder:validation:Required
	DNewLimit *float64 `json:"dNewLimit" tf:"d_new_limit,omitempty"`

	// Indicate whether to drop abroad traffic or not.
	// +kubebuilder:validation:Required
	DropAbroad *bool `json:"dropAbroad" tf:"drop_abroad,omitempty"`

	// Indicate whether to drop ICMP protocol or not.
	// +kubebuilder:validation:Required
	DropIcmp *bool `json:"dropIcmp" tf:"drop_icmp,omitempty"`

	// Indicate whether to drop other protocols(exclude TCP/UDP/ICMP) or not.
	// +kubebuilder:validation:Required
	DropOther *bool `json:"dropOther" tf:"drop_other,omitempty"`

	// Indicate whether to drop TCP protocol or not.
	// +kubebuilder:validation:Required
	DropTCP *bool `json:"dropTcp" tf:"drop_tcp,omitempty"`

	// Indicate to drop UDP protocol or not.
	// +kubebuilder:validation:Required
	DropUDP *bool `json:"dropUdp" tf:"drop_udp,omitempty"`

	// The limit of ICMP traffic rate. Valid value ranges: (0~4294967295)(Mbps).
	// +kubebuilder:validation:Required
	IcmpMbpsLimit *float64 `json:"icmpMbpsLimit" tf:"icmp_mbps_limit,omitempty"`

	// Indicate to enable null connection or not.
	// +kubebuilder:validation:Required
	NullConnEnable *bool `json:"nullConnEnable" tf:"null_conn_enable,omitempty"`

	// The limit of other protocols(exclude TCP/UDP/ICMP) traffic rate. Valid value ranges: (0~4294967295)(Mbps).
	// +kubebuilder:validation:Required
	OtherMbpsLimit *float64 `json:"otherMbpsLimit" tf:"other_mbps_limit,omitempty"`

	// The limit of concurrent connections based on source IP. Valid value ranges: (0~4294967295).
	// +kubebuilder:validation:Required
	SConnLimit *float64 `json:"sConnLimit" tf:"s_conn_limit,omitempty"`

	// The limit of new connections based on source IP. Valid value ranges: (0~4294967295).
	// +kubebuilder:validation:Required
	SNewLimit *float64 `json:"sNewLimit" tf:"s_new_limit,omitempty"`

	// The limit of syn of abnormal connection check. Valid value ranges: (0~100).
	// +kubebuilder:validation:Required
	SynLimit *float64 `json:"synLimit" tf:"syn_limit,omitempty"`

	// The percentage of syn in ack of abnormal connection check. Valid value ranges: (0~100).
	// +kubebuilder:validation:Optional
	SynRate *float64 `json:"synRate,omitempty" tf:"syn_rate,omitempty"`

	// The limit of TCP traffic. Valid value ranges: (0~4294967295)(Mbps).
	// +kubebuilder:validation:Required
	TCPMbpsLimit *float64 `json:"tcpMbpsLimit" tf:"tcp_mbps_limit,omitempty"`

	// The limit of UDP traffic rate. Valid value ranges: (0~4294967295)(Mbps).
	// +kubebuilder:validation:Required
	UDPMbpsLimit *float64 `json:"udpMbpsLimit" tf:"udp_mbps_limit,omitempty"`
}

type PacketFiltersObservation struct {
}

type PacketFiltersParameters struct {

	// Action of port to take. Valid values: `drop`, `drop_black`,`drop_rst`,`drop_black_rst`,`transmit`.`drop`(drop the packet), `drop_black`(drop the packet and black the ip),`drop_rst`(drop the packet and disconnect),`drop_black_rst`(drop the packet, black the ip and disconnect),`transmit`(transmit the packet).
	// +kubebuilder:validation:Optional
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// End port of the destination. Valid value ranges: (0~65535). It must be greater than `d_start_port`.
	// +kubebuilder:validation:Optional
	DEndPort *float64 `json:"dEndPort,omitempty" tf:"d_end_port,omitempty"`

	// Start port of the destination. Valid value ranges: (0~65535).
	// +kubebuilder:validation:Optional
	DStartPort *float64 `json:"dStartPort,omitempty" tf:"d_start_port,omitempty"`

	// The depth of match. Valid value ranges: (0~1500).
	// +kubebuilder:validation:Optional
	Depth *float64 `json:"depth,omitempty" tf:"depth,omitempty"`

	// Indicate whether to include the key word/regular expression or not.
	// +kubebuilder:validation:Optional
	IsInclude *bool `json:"isInclude,omitempty" tf:"is_include,omitempty"`

	// Indicate whether to check load or not, `begin_l5` means to match and `no_match` means not.
	// +kubebuilder:validation:Optional
	MatchBegin *string `json:"matchBegin,omitempty" tf:"match_begin,omitempty"`

	// The key word or regular expression.
	// +kubebuilder:validation:Optional
	MatchStr *string `json:"matchStr,omitempty" tf:"match_str,omitempty"`

	// Match type. Valid values: `sunday` and `pcre`. `sunday` means key word match while `pcre` means regular match.
	// +kubebuilder:validation:Optional
	MatchType *string `json:"matchType,omitempty" tf:"match_type,omitempty"`

	// The offset of match. Valid value ranges: (0~1500).
	// +kubebuilder:validation:Optional
	Offset *float64 `json:"offset,omitempty" tf:"offset,omitempty"`

	// The max length of the packet. Valid value ranges: (0~1500)(Mbps). It must be greater than `pkt_length_min`.
	// +kubebuilder:validation:Optional
	PktLengthMax *float64 `json:"pktLengthMax,omitempty" tf:"pkt_length_max,omitempty"`

	// The minimum length of the packet. Valid value ranges: (0~1500)(Mbps).
	// +kubebuilder:validation:Optional
	PktLengthMin *float64 `json:"pktLengthMin,omitempty" tf:"pkt_length_min,omitempty"`

	// Protocol. Valid values: `tcp`, `udp`, `icmp`, `all`.
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// End port of the source. Valid value ranges: (0~65535). It must be greater than `s_start_port`.
	// +kubebuilder:validation:Optional
	SEndPort *float64 `json:"sEndPort,omitempty" tf:"s_end_port,omitempty"`

	// Start port of the source. Valid value ranges: (0~65535).
	// +kubebuilder:validation:Optional
	SStartPort *float64 `json:"sStartPort,omitempty" tf:"s_start_port,omitempty"`
}

type PortFiltersObservation struct {
}

type PortFiltersParameters struct {

	// Action of port to take. Valid values: `drop`, `transmit`.
	// +kubebuilder:validation:Optional
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// End port. Valid value ranges: (0~65535). It must be greater than `start_port`.
	// +kubebuilder:validation:Optional
	EndPort *float64 `json:"endPort,omitempty" tf:"end_port,omitempty"`

	// The type of forbidden port. Valid values: `0`, `1`, `2`. `0` for destination ports make effect, `1` for source ports make effect. `2` for both destination and source ports.
	// +kubebuilder:validation:Optional
	Kind *float64 `json:"kind,omitempty" tf:"kind,omitempty"`

	// Protocol. Valid values are `tcp`, `udp`, `icmp`, `all`.
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Start port. Valid value ranges: (0~65535).
	// +kubebuilder:validation:Optional
	StartPort *float64 `json:"startPort,omitempty" tf:"start_port,omitempty"`
}

type WatermarkFiltersObservation struct {
}

type WatermarkFiltersParameters struct {

	// Indicate whether to auto-remove the watermark or not.
	// +kubebuilder:validation:Optional
	AutoRemove *bool `json:"autoRemove,omitempty" tf:"auto_remove,omitempty"`

	// The offset of watermark. Valid value ranges: (0~1500).
	// +kubebuilder:validation:Optional
	Offset *float64 `json:"offset,omitempty" tf:"offset,omitempty"`

	// Indicate whether to open watermark or not. It muse be set `true` when any field of watermark was set.
	// +kubebuilder:validation:Optional
	OpenSwitch *bool `json:"openSwitch,omitempty" tf:"open_switch,omitempty"`

	// Port range of TCP, the format is like `2000-3000`.
	// +kubebuilder:validation:Optional
	TCPPortList []*string `json:"tcpPortList,omitempty" tf:"tcp_port_list,omitempty"`

	// Port range of TCP, the format is like `2000-3000`.
	// +kubebuilder:validation:Optional
	UDPPortList []*string `json:"udpPortList,omitempty" tf:"udp_port_list,omitempty"`
}

type WatermarkKeyObservation struct {
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	OpenSwitch *bool `json:"openSwitch,omitempty" tf:"open_switch,omitempty"`
}

type WatermarkKeyParameters struct {
}

// DdosPolicySpec defines the desired state of DdosPolicy
type DdosPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DdosPolicyParameters `json:"forProvider"`
}

// DdosPolicyStatus defines the observed state of DdosPolicy.
type DdosPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DdosPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DdosPolicy is the Schema for the DdosPolicys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloudjet}
type DdosPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DdosPolicySpec   `json:"spec"`
	Status            DdosPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DdosPolicyList contains a list of DdosPolicys
type DdosPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DdosPolicy `json:"items"`
}

// Repository type metadata.
var (
	DdosPolicy_Kind             = "DdosPolicy"
	DdosPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DdosPolicy_Kind}.String()
	DdosPolicy_KindAPIVersion   = DdosPolicy_Kind + "." + CRDGroupVersion.String()
	DdosPolicy_GroupVersionKind = CRDGroupVersion.WithKind(DdosPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&DdosPolicy{}, &DdosPolicyList{})
}