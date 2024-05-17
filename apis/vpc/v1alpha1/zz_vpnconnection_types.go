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

type SecurityGroupPolicyInitParameters struct {

	// Local cidr block.
	// Local cidr block.
	LocalCidrBlock *string `json:"localCidrBlock,omitempty" tf:"local_cidr_block,omitempty"`

	// Remote cidr block list.
	// Remote cidr block list.
	RemoteCidrBlock []*string `json:"remoteCidrBlock,omitempty" tf:"remote_cidr_block,omitempty"`
}

type SecurityGroupPolicyObservation struct {

	// Local cidr block.
	// Local cidr block.
	LocalCidrBlock *string `json:"localCidrBlock,omitempty" tf:"local_cidr_block,omitempty"`

	// Remote cidr block list.
	// Remote cidr block list.
	RemoteCidrBlock []*string `json:"remoteCidrBlock,omitempty" tf:"remote_cidr_block,omitempty"`
}

type SecurityGroupPolicyParameters struct {

	// Local cidr block.
	// Local cidr block.
	// +kubebuilder:validation:Optional
	LocalCidrBlock *string `json:"localCidrBlock" tf:"local_cidr_block,omitempty"`

	// Remote cidr block list.
	// Remote cidr block list.
	// +kubebuilder:validation:Optional
	RemoteCidrBlock []*string `json:"remoteCidrBlock" tf:"remote_cidr_block,omitempty"`
}

type VPNConnectionInitParameters struct {

	// The action after DPD timeout. Valid values: clear (disconnect) and restart (try again). It is valid when DpdEnable is 1.
	// The action after DPD timeout. Valid values: clear (disconnect) and restart (try again). It is valid when DpdEnable is 1.
	DpdAction *string `json:"dpdAction,omitempty" tf:"dpd_action,omitempty"`

	// Specifies whether to enable DPD. Valid values: 0 (disable) and 1 (enable).
	// Specifies whether to enable DPD. Valid values: 0 (disable) and 1 (enable).
	DpdEnable *float64 `json:"dpdEnable,omitempty" tf:"dpd_enable,omitempty"`

	// DPD timeout period.Valid value ranges: [30~60], Default: 30; unit: second. If the request is not responded within this period, the peer end is considered not exists. This parameter is valid when the value of DpdEnable is 1.
	// DPD timeout period.Valid value ranges: [30~60], Default: 30; unit: second. If the request is not responded within this period, the peer end is considered not exists. This parameter is valid when the value of DpdEnable is 1.
	DpdTimeout *float64 `json:"dpdTimeout,omitempty" tf:"dpd_timeout,omitempty"`

	// Whether intra-tunnel health checks are supported.
	// Whether intra-tunnel health checks are supported.
	EnableHealthCheck *bool `json:"enableHealthCheck,omitempty" tf:"enable_health_check,omitempty"`

	// Health check the address of this terminal.
	// Health check the address of this terminal.
	HealthCheckLocalIP *string `json:"healthCheckLocalIp,omitempty" tf:"health_check_local_ip,omitempty"`

	// Health check peer address.
	// Health check peer address.
	HealthCheckRemoteIP *string `json:"healthCheckRemoteIp,omitempty" tf:"health_check_remote_ip,omitempty"`

	// DH group name of the IKE operation specification. Valid values: GROUP1, GROUP2, GROUP5, GROUP14, GROUP24. Default value is GROUP1.
	// DH group name of the IKE operation specification. Valid values: `GROUP1`, `GROUP2`, `GROUP5`, `GROUP14`, `GROUP24`. Default value is `GROUP1`.
	IkeDhGroupName *string `json:"ikeDhGroupName,omitempty" tf:"ike_dh_group_name,omitempty"`

	// Exchange mode of the IKE operation specification. Valid values: AGGRESSIVE, MAIN. Default value is MAIN.
	// Exchange mode of the IKE operation specification. Valid values: `AGGRESSIVE`, `MAIN`. Default value is `MAIN`.
	IkeExchangeMode *string `json:"ikeExchangeMode,omitempty" tf:"ike_exchange_mode,omitempty"`

	// Local address of IKE operation specification, valid when ike_local_identity is ADDRESS, generally the value is public_ip_address of the related VPN gateway.
	// Local address of IKE operation specification, valid when ike_local_identity is `ADDRESS`, generally the value is `public_ip_address` of the related VPN gateway.
	IkeLocalAddress *string `json:"ikeLocalAddress,omitempty" tf:"ike_local_address,omitempty"`

	// Local FQDN name of the IKE operation specification.
	// Local FQDN name of the IKE operation specification.
	IkeLocalFqdnName *string `json:"ikeLocalFqdnName,omitempty" tf:"ike_local_fqdn_name,omitempty"`

	// Local identity way of IKE operation specification. Valid values: ADDRESS, FQDN. Default value is ADDRESS.
	// Local identity way of IKE operation specification. Valid values: `ADDRESS`, `FQDN`. Default value is `ADDRESS`.
	IkeLocalIdentity *string `json:"ikeLocalIdentity,omitempty" tf:"ike_local_identity,omitempty"`

	// Proto authenticate algorithm of the IKE operation specification. Valid values: MD5, SHA, SHA-256. Default Value is MD5.
	// Proto authenticate algorithm of the IKE operation specification. Valid values: `MD5`, `SHA`, `SHA-256`. Default Value is `MD5`.
	IkeProtoAuthenAlgorithm *string `json:"ikeProtoAuthenAlgorithm,omitempty" tf:"ike_proto_authen_algorithm,omitempty"`

	// Proto encrypt algorithm of the IKE operation specification. Valid values: 3DES-CBC, AES-CBC-128, AES-CBC-192, AES-CBC-256, DES-CBC, SM4, AES128GCM128, AES192GCM128, AES256GCM128,AES128GCM128, AES192GCM128, AES256GCM128. Default value is 3DES-CBC.
	// Proto encrypt algorithm of the IKE operation specification. Valid values: `3DES-CBC`, `AES-CBC-128`, `AES-CBC-192`, `AES-CBC-256`, `DES-CBC`, `SM4`, `AES128GCM128`, `AES192GCM128`, `AES256GCM128`,`AES128GCM128`, `AES192GCM128`, `AES256GCM128`. Default value is `3DES-CBC`.
	IkeProtoEncryAlgorithm *string `json:"ikeProtoEncryAlgorithm,omitempty" tf:"ike_proto_encry_algorithm,omitempty"`

	// Remote address of IKE operation specification, valid when ike_remote_identity is ADDRESS, generally the value is public_ip_address of the related customer gateway.
	// Remote address of IKE operation specification, valid when ike_remote_identity is `ADDRESS`, generally the value is `public_ip_address` of the related customer gateway.
	IkeRemoteAddress *string `json:"ikeRemoteAddress,omitempty" tf:"ike_remote_address,omitempty"`

	// Remote FQDN name of the IKE operation specification.
	// Remote FQDN name of the IKE operation specification.
	IkeRemoteFqdnName *string `json:"ikeRemoteFqdnName,omitempty" tf:"ike_remote_fqdn_name,omitempty"`

	// Remote identity way of IKE operation specification. Valid values: ADDRESS, FQDN. Default value is ADDRESS.
	// Remote identity way of IKE operation specification. Valid values: `ADDRESS`, `FQDN`. Default value is `ADDRESS`.
	IkeRemoteIdentity *string `json:"ikeRemoteIdentity,omitempty" tf:"ike_remote_identity,omitempty"`

	// SA lifetime of the IKE operation specification, unit is second. The value ranges from 60 to 604800. Default value is 86400 seconds.
	// SA lifetime of the IKE operation specification, unit is `second`. The value ranges from 60 to 604800. Default value is 86400 seconds.
	IkeSaLifetimeSeconds *float64 `json:"ikeSaLifetimeSeconds,omitempty" tf:"ike_sa_lifetime_seconds,omitempty"`

	// Version of the IKE operation specification, values: IKEV1, IKEV2. Default value is IKEV1.
	// Version of the IKE operation specification, values: `IKEV1`, `IKEV2`. Default value is `IKEV1`.
	IkeVersion *string `json:"ikeVersion,omitempty" tf:"ike_version,omitempty"`

	// Encrypt algorithm of the IPSEC operation specification. Valid values: 3DES-CBC, AES-CBC-128, AES-CBC-192, AES-CBC-256, DES-CBC, SM4, NULL, AES128GCM128, AES192GCM128, AES256GCM128. Default value is 3DES-CBC.
	// Encrypt algorithm of the IPSEC operation specification. Valid values: `3DES-CBC`, `AES-CBC-128`, `AES-CBC-192`, `AES-CBC-256`, `DES-CBC`, `SM4`, `NULL`, `AES128GCM128`, `AES192GCM128`, `AES256GCM128`. Default value is `3DES-CBC`.
	IpsecEncryptAlgorithm *string `json:"ipsecEncryptAlgorithm,omitempty" tf:"ipsec_encrypt_algorithm,omitempty"`

	// Integrity algorithm of the IPSEC operation specification. Valid values: SHA1, MD5, SHA-256. Default value is MD5.
	// Integrity algorithm of the IPSEC operation specification. Valid values: `SHA1`, `MD5`, `SHA-256`. Default value is `MD5`.
	IpsecIntegrityAlgorithm *string `json:"ipsecIntegrityAlgorithm,omitempty" tf:"ipsec_integrity_algorithm,omitempty"`

	// PFS DH group. Valid value: DH-GROUP1, DH-GROUP2, DH-GROUP5, DH-GROUP14, DH-GROUP24, NULL. Default value is NULL.
	// PFS DH group. Valid value: `DH-GROUP1`, `DH-GROUP2`, `DH-GROUP5`, `DH-GROUP14`, `DH-GROUP24`, `NULL`. Default value is `NULL`.
	IpsecPfsDhGroup *string `json:"ipsecPfsDhGroup,omitempty" tf:"ipsec_pfs_dh_group,omitempty"`

	// SA lifetime of the IPSEC operation specification, unit is second. Valid value ranges: [180~604800]. Default value is 3600 seconds.
	// SA lifetime of the IPSEC operation specification, unit is second. Valid value ranges: [180~604800]. Default value is 3600 seconds.
	IpsecSaLifetimeSeconds *float64 `json:"ipsecSaLifetimeSeconds,omitempty" tf:"ipsec_sa_lifetime_seconds,omitempty"`

	// SA lifetime of the IPSEC operation specification, unit is KB. The value should not be less then 2560. Default value is 1843200.
	// SA lifetime of the IPSEC operation specification, unit is KB. The value should not be less then 2560. Default value is 1843200.
	IpsecSaLifetimeTraffic *float64 `json:"ipsecSaLifetimeTraffic,omitempty" tf:"ipsec_sa_lifetime_traffic,omitempty"`

	// Name of the VPN connection. The length of character is limited to 1-60.
	// Name of the VPN connection. The length of character is limited to 1-60.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Pre-shared key of the VPN connection.
	// Pre-shared key of the VPN connection.
	PreShareKey *string `json:"preShareKey,omitempty" tf:"pre_share_key,omitempty"`

	// Route type of the VPN connection. Valid value: STATIC, StaticRoute, Policy.
	// Route type of the VPN connection. Valid value: `STATIC`, `StaticRoute`, `Policy`.
	RouteType *string `json:"routeType,omitempty" tf:"route_type,omitempty"`

	// SPD policy group, for example: {"10.0.0.5/24":["172.123.10.5/16"]}, 10.0.0.5/24 is the vpc intranet segment, and 172.123.10.5/16 is the IDC network segment. Users specify which network segments in the VPC can communicate with which network segments in your IDC.
	// SPD policy group, for example: {"10.0.0.5/24":["172.123.10.5/16"]}, 10.0.0.5/24 is the vpc intranet segment, and 172.123.10.5/16 is the IDC network segment. Users specify which network segments in the VPC can communicate with which network segments in your IDC.
	SecurityGroupPolicy []SecurityGroupPolicyInitParameters `json:"securityGroupPolicy,omitempty" tf:"security_group_policy,omitempty"`

	// A list of tags used to associate different resources.
	// A list of tags used to associate different resources.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type VPNConnectionObservation struct {

	// Create time of the VPN connection.
	// Create time of the VPN connection.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// ID of the customer gateway.
	// ID of the customer gateway.
	CustomerGatewayID *string `json:"customerGatewayId,omitempty" tf:"customer_gateway_id,omitempty"`

	// The action after DPD timeout. Valid values: clear (disconnect) and restart (try again). It is valid when DpdEnable is 1.
	// The action after DPD timeout. Valid values: clear (disconnect) and restart (try again). It is valid when DpdEnable is 1.
	DpdAction *string `json:"dpdAction,omitempty" tf:"dpd_action,omitempty"`

	// Specifies whether to enable DPD. Valid values: 0 (disable) and 1 (enable).
	// Specifies whether to enable DPD. Valid values: 0 (disable) and 1 (enable).
	DpdEnable *float64 `json:"dpdEnable,omitempty" tf:"dpd_enable,omitempty"`

	// DPD timeout period.Valid value ranges: [30~60], Default: 30; unit: second. If the request is not responded within this period, the peer end is considered not exists. This parameter is valid when the value of DpdEnable is 1.
	// DPD timeout period.Valid value ranges: [30~60], Default: 30; unit: second. If the request is not responded within this period, the peer end is considered not exists. This parameter is valid when the value of DpdEnable is 1.
	DpdTimeout *float64 `json:"dpdTimeout,omitempty" tf:"dpd_timeout,omitempty"`

	// Whether intra-tunnel health checks are supported.
	// Whether intra-tunnel health checks are supported.
	EnableHealthCheck *bool `json:"enableHealthCheck,omitempty" tf:"enable_health_check,omitempty"`

	// Encrypt proto of the VPN connection.
	// Encrypt proto of the VPN connection.
	EncryptProto *string `json:"encryptProto,omitempty" tf:"encrypt_proto,omitempty"`

	// Health check the address of this terminal.
	// Health check the address of this terminal.
	HealthCheckLocalIP *string `json:"healthCheckLocalIp,omitempty" tf:"health_check_local_ip,omitempty"`

	// Health check peer address.
	// Health check peer address.
	HealthCheckRemoteIP *string `json:"healthCheckRemoteIp,omitempty" tf:"health_check_remote_ip,omitempty"`

	// ID of the resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// DH group name of the IKE operation specification. Valid values: GROUP1, GROUP2, GROUP5, GROUP14, GROUP24. Default value is GROUP1.
	// DH group name of the IKE operation specification. Valid values: `GROUP1`, `GROUP2`, `GROUP5`, `GROUP14`, `GROUP24`. Default value is `GROUP1`.
	IkeDhGroupName *string `json:"ikeDhGroupName,omitempty" tf:"ike_dh_group_name,omitempty"`

	// Exchange mode of the IKE operation specification. Valid values: AGGRESSIVE, MAIN. Default value is MAIN.
	// Exchange mode of the IKE operation specification. Valid values: `AGGRESSIVE`, `MAIN`. Default value is `MAIN`.
	IkeExchangeMode *string `json:"ikeExchangeMode,omitempty" tf:"ike_exchange_mode,omitempty"`

	// Local address of IKE operation specification, valid when ike_local_identity is ADDRESS, generally the value is public_ip_address of the related VPN gateway.
	// Local address of IKE operation specification, valid when ike_local_identity is `ADDRESS`, generally the value is `public_ip_address` of the related VPN gateway.
	IkeLocalAddress *string `json:"ikeLocalAddress,omitempty" tf:"ike_local_address,omitempty"`

	// Local FQDN name of the IKE operation specification.
	// Local FQDN name of the IKE operation specification.
	IkeLocalFqdnName *string `json:"ikeLocalFqdnName,omitempty" tf:"ike_local_fqdn_name,omitempty"`

	// Local identity way of IKE operation specification. Valid values: ADDRESS, FQDN. Default value is ADDRESS.
	// Local identity way of IKE operation specification. Valid values: `ADDRESS`, `FQDN`. Default value is `ADDRESS`.
	IkeLocalIdentity *string `json:"ikeLocalIdentity,omitempty" tf:"ike_local_identity,omitempty"`

	// Proto authenticate algorithm of the IKE operation specification. Valid values: MD5, SHA, SHA-256. Default Value is MD5.
	// Proto authenticate algorithm of the IKE operation specification. Valid values: `MD5`, `SHA`, `SHA-256`. Default Value is `MD5`.
	IkeProtoAuthenAlgorithm *string `json:"ikeProtoAuthenAlgorithm,omitempty" tf:"ike_proto_authen_algorithm,omitempty"`

	// Proto encrypt algorithm of the IKE operation specification. Valid values: 3DES-CBC, AES-CBC-128, AES-CBC-192, AES-CBC-256, DES-CBC, SM4, AES128GCM128, AES192GCM128, AES256GCM128,AES128GCM128, AES192GCM128, AES256GCM128. Default value is 3DES-CBC.
	// Proto encrypt algorithm of the IKE operation specification. Valid values: `3DES-CBC`, `AES-CBC-128`, `AES-CBC-192`, `AES-CBC-256`, `DES-CBC`, `SM4`, `AES128GCM128`, `AES192GCM128`, `AES256GCM128`,`AES128GCM128`, `AES192GCM128`, `AES256GCM128`. Default value is `3DES-CBC`.
	IkeProtoEncryAlgorithm *string `json:"ikeProtoEncryAlgorithm,omitempty" tf:"ike_proto_encry_algorithm,omitempty"`

	// Remote address of IKE operation specification, valid when ike_remote_identity is ADDRESS, generally the value is public_ip_address of the related customer gateway.
	// Remote address of IKE operation specification, valid when ike_remote_identity is `ADDRESS`, generally the value is `public_ip_address` of the related customer gateway.
	IkeRemoteAddress *string `json:"ikeRemoteAddress,omitempty" tf:"ike_remote_address,omitempty"`

	// Remote FQDN name of the IKE operation specification.
	// Remote FQDN name of the IKE operation specification.
	IkeRemoteFqdnName *string `json:"ikeRemoteFqdnName,omitempty" tf:"ike_remote_fqdn_name,omitempty"`

	// Remote identity way of IKE operation specification. Valid values: ADDRESS, FQDN. Default value is ADDRESS.
	// Remote identity way of IKE operation specification. Valid values: `ADDRESS`, `FQDN`. Default value is `ADDRESS`.
	IkeRemoteIdentity *string `json:"ikeRemoteIdentity,omitempty" tf:"ike_remote_identity,omitempty"`

	// SA lifetime of the IKE operation specification, unit is second. The value ranges from 60 to 604800. Default value is 86400 seconds.
	// SA lifetime of the IKE operation specification, unit is `second`. The value ranges from 60 to 604800. Default value is 86400 seconds.
	IkeSaLifetimeSeconds *float64 `json:"ikeSaLifetimeSeconds,omitempty" tf:"ike_sa_lifetime_seconds,omitempty"`

	// Version of the IKE operation specification, values: IKEV1, IKEV2. Default value is IKEV1.
	// Version of the IKE operation specification, values: `IKEV1`, `IKEV2`. Default value is `IKEV1`.
	IkeVersion *string `json:"ikeVersion,omitempty" tf:"ike_version,omitempty"`

	// Encrypt algorithm of the IPSEC operation specification. Valid values: 3DES-CBC, AES-CBC-128, AES-CBC-192, AES-CBC-256, DES-CBC, SM4, NULL, AES128GCM128, AES192GCM128, AES256GCM128. Default value is 3DES-CBC.
	// Encrypt algorithm of the IPSEC operation specification. Valid values: `3DES-CBC`, `AES-CBC-128`, `AES-CBC-192`, `AES-CBC-256`, `DES-CBC`, `SM4`, `NULL`, `AES128GCM128`, `AES192GCM128`, `AES256GCM128`. Default value is `3DES-CBC`.
	IpsecEncryptAlgorithm *string `json:"ipsecEncryptAlgorithm,omitempty" tf:"ipsec_encrypt_algorithm,omitempty"`

	// Integrity algorithm of the IPSEC operation specification. Valid values: SHA1, MD5, SHA-256. Default value is MD5.
	// Integrity algorithm of the IPSEC operation specification. Valid values: `SHA1`, `MD5`, `SHA-256`. Default value is `MD5`.
	IpsecIntegrityAlgorithm *string `json:"ipsecIntegrityAlgorithm,omitempty" tf:"ipsec_integrity_algorithm,omitempty"`

	// PFS DH group. Valid value: DH-GROUP1, DH-GROUP2, DH-GROUP5, DH-GROUP14, DH-GROUP24, NULL. Default value is NULL.
	// PFS DH group. Valid value: `DH-GROUP1`, `DH-GROUP2`, `DH-GROUP5`, `DH-GROUP14`, `DH-GROUP24`, `NULL`. Default value is `NULL`.
	IpsecPfsDhGroup *string `json:"ipsecPfsDhGroup,omitempty" tf:"ipsec_pfs_dh_group,omitempty"`

	// SA lifetime of the IPSEC operation specification, unit is second. Valid value ranges: [180~604800]. Default value is 3600 seconds.
	// SA lifetime of the IPSEC operation specification, unit is second. Valid value ranges: [180~604800]. Default value is 3600 seconds.
	IpsecSaLifetimeSeconds *float64 `json:"ipsecSaLifetimeSeconds,omitempty" tf:"ipsec_sa_lifetime_seconds,omitempty"`

	// SA lifetime of the IPSEC operation specification, unit is KB. The value should not be less then 2560. Default value is 1843200.
	// SA lifetime of the IPSEC operation specification, unit is KB. The value should not be less then 2560. Default value is 1843200.
	IpsecSaLifetimeTraffic *float64 `json:"ipsecSaLifetimeTraffic,omitempty" tf:"ipsec_sa_lifetime_traffic,omitempty"`

	// Indicate whether is ccn type. Modification of this field only impacts force new logic of vpc_id. If is_ccn_type is true, modification of vpc_id will be ignored.
	// Indicate whether is ccn type. Modification of this field only impacts force new logic of `vpc_id`. If `is_ccn_type` is true, modification of `vpc_id` will be ignored.
	IsCcnType *bool `json:"isCcnType,omitempty" tf:"is_ccn_type,omitempty"`

	// Name of the VPN connection. The length of character is limited to 1-60.
	// Name of the VPN connection. The length of character is limited to 1-60.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Net status of the VPN connection. Valid value: AVAILABLE.
	// Net status of the VPN connection. Valid value: `AVAILABLE`.
	NetStatus *string `json:"netStatus,omitempty" tf:"net_status,omitempty"`

	// Pre-shared key of the VPN connection.
	// Pre-shared key of the VPN connection.
	PreShareKey *string `json:"preShareKey,omitempty" tf:"pre_share_key,omitempty"`

	// Route type of the VPN connection. Valid value: STATIC, StaticRoute, Policy.
	// Route type of the VPN connection. Valid value: `STATIC`, `StaticRoute`, `Policy`.
	RouteType *string `json:"routeType,omitempty" tf:"route_type,omitempty"`

	// SPD policy group, for example: {"10.0.0.5/24":["172.123.10.5/16"]}, 10.0.0.5/24 is the vpc intranet segment, and 172.123.10.5/16 is the IDC network segment. Users specify which network segments in the VPC can communicate with which network segments in your IDC.
	// SPD policy group, for example: {"10.0.0.5/24":["172.123.10.5/16"]}, 10.0.0.5/24 is the vpc intranet segment, and 172.123.10.5/16 is the IDC network segment. Users specify which network segments in the VPC can communicate with which network segments in your IDC.
	SecurityGroupPolicy []SecurityGroupPolicyObservation `json:"securityGroupPolicy,omitempty" tf:"security_group_policy,omitempty"`

	// State of the connection. Valid value: PENDING, AVAILABLE, DELETING.
	// State of the connection. Valid value: `PENDING`, `AVAILABLE`, `DELETING`.
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// A list of tags used to associate different resources.
	// A list of tags used to associate different resources.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// ID of the VPC. Required if vpn gateway is not in CCN type, and doesn't make sense for CCN vpn gateway.
	// ID of the VPC. Required if vpn gateway is not in `CCN` type, and doesn't make sense for `CCN` vpn gateway.
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// ID of the VPN gateway.
	// ID of the VPN gateway.
	VPNGatewayID *string `json:"vpnGatewayId,omitempty" tf:"vpn_gateway_id,omitempty"`

	// Vpn proto of the VPN connection.
	// Vpn proto of the VPN connection.
	VPNProto *string `json:"vpnProto,omitempty" tf:"vpn_proto,omitempty"`
}

type VPNConnectionParameters struct {

	// ID of the customer gateway.
	// ID of the customer gateway.
	// +crossplane:generate:reference:type=VPNCustomerGateway
	// +kubebuilder:validation:Optional
	CustomerGatewayID *string `json:"customerGatewayId,omitempty" tf:"customer_gateway_id,omitempty"`

	// Reference to a VPNCustomerGateway to populate customerGatewayId.
	// +kubebuilder:validation:Optional
	CustomerGatewayIDRef *v1.Reference `json:"customerGatewayIdRef,omitempty" tf:"-"`

	// Selector for a VPNCustomerGateway to populate customerGatewayId.
	// +kubebuilder:validation:Optional
	CustomerGatewayIDSelector *v1.Selector `json:"customerGatewayIdSelector,omitempty" tf:"-"`

	// The action after DPD timeout. Valid values: clear (disconnect) and restart (try again). It is valid when DpdEnable is 1.
	// The action after DPD timeout. Valid values: clear (disconnect) and restart (try again). It is valid when DpdEnable is 1.
	// +kubebuilder:validation:Optional
	DpdAction *string `json:"dpdAction,omitempty" tf:"dpd_action,omitempty"`

	// Specifies whether to enable DPD. Valid values: 0 (disable) and 1 (enable).
	// Specifies whether to enable DPD. Valid values: 0 (disable) and 1 (enable).
	// +kubebuilder:validation:Optional
	DpdEnable *float64 `json:"dpdEnable,omitempty" tf:"dpd_enable,omitempty"`

	// DPD timeout period.Valid value ranges: [30~60], Default: 30; unit: second. If the request is not responded within this period, the peer end is considered not exists. This parameter is valid when the value of DpdEnable is 1.
	// DPD timeout period.Valid value ranges: [30~60], Default: 30; unit: second. If the request is not responded within this period, the peer end is considered not exists. This parameter is valid when the value of DpdEnable is 1.
	// +kubebuilder:validation:Optional
	DpdTimeout *float64 `json:"dpdTimeout,omitempty" tf:"dpd_timeout,omitempty"`

	// Whether intra-tunnel health checks are supported.
	// Whether intra-tunnel health checks are supported.
	// +kubebuilder:validation:Optional
	EnableHealthCheck *bool `json:"enableHealthCheck,omitempty" tf:"enable_health_check,omitempty"`

	// Health check the address of this terminal.
	// Health check the address of this terminal.
	// +kubebuilder:validation:Optional
	HealthCheckLocalIP *string `json:"healthCheckLocalIp,omitempty" tf:"health_check_local_ip,omitempty"`

	// Health check peer address.
	// Health check peer address.
	// +kubebuilder:validation:Optional
	HealthCheckRemoteIP *string `json:"healthCheckRemoteIp,omitempty" tf:"health_check_remote_ip,omitempty"`

	// DH group name of the IKE operation specification. Valid values: GROUP1, GROUP2, GROUP5, GROUP14, GROUP24. Default value is GROUP1.
	// DH group name of the IKE operation specification. Valid values: `GROUP1`, `GROUP2`, `GROUP5`, `GROUP14`, `GROUP24`. Default value is `GROUP1`.
	// +kubebuilder:validation:Optional
	IkeDhGroupName *string `json:"ikeDhGroupName,omitempty" tf:"ike_dh_group_name,omitempty"`

	// Exchange mode of the IKE operation specification. Valid values: AGGRESSIVE, MAIN. Default value is MAIN.
	// Exchange mode of the IKE operation specification. Valid values: `AGGRESSIVE`, `MAIN`. Default value is `MAIN`.
	// +kubebuilder:validation:Optional
	IkeExchangeMode *string `json:"ikeExchangeMode,omitempty" tf:"ike_exchange_mode,omitempty"`

	// Local address of IKE operation specification, valid when ike_local_identity is ADDRESS, generally the value is public_ip_address of the related VPN gateway.
	// Local address of IKE operation specification, valid when ike_local_identity is `ADDRESS`, generally the value is `public_ip_address` of the related VPN gateway.
	// +kubebuilder:validation:Optional
	IkeLocalAddress *string `json:"ikeLocalAddress,omitempty" tf:"ike_local_address,omitempty"`

	// Local FQDN name of the IKE operation specification.
	// Local FQDN name of the IKE operation specification.
	// +kubebuilder:validation:Optional
	IkeLocalFqdnName *string `json:"ikeLocalFqdnName,omitempty" tf:"ike_local_fqdn_name,omitempty"`

	// Local identity way of IKE operation specification. Valid values: ADDRESS, FQDN. Default value is ADDRESS.
	// Local identity way of IKE operation specification. Valid values: `ADDRESS`, `FQDN`. Default value is `ADDRESS`.
	// +kubebuilder:validation:Optional
	IkeLocalIdentity *string `json:"ikeLocalIdentity,omitempty" tf:"ike_local_identity,omitempty"`

	// Proto authenticate algorithm of the IKE operation specification. Valid values: MD5, SHA, SHA-256. Default Value is MD5.
	// Proto authenticate algorithm of the IKE operation specification. Valid values: `MD5`, `SHA`, `SHA-256`. Default Value is `MD5`.
	// +kubebuilder:validation:Optional
	IkeProtoAuthenAlgorithm *string `json:"ikeProtoAuthenAlgorithm,omitempty" tf:"ike_proto_authen_algorithm,omitempty"`

	// Proto encrypt algorithm of the IKE operation specification. Valid values: 3DES-CBC, AES-CBC-128, AES-CBC-192, AES-CBC-256, DES-CBC, SM4, AES128GCM128, AES192GCM128, AES256GCM128,AES128GCM128, AES192GCM128, AES256GCM128. Default value is 3DES-CBC.
	// Proto encrypt algorithm of the IKE operation specification. Valid values: `3DES-CBC`, `AES-CBC-128`, `AES-CBC-192`, `AES-CBC-256`, `DES-CBC`, `SM4`, `AES128GCM128`, `AES192GCM128`, `AES256GCM128`,`AES128GCM128`, `AES192GCM128`, `AES256GCM128`. Default value is `3DES-CBC`.
	// +kubebuilder:validation:Optional
	IkeProtoEncryAlgorithm *string `json:"ikeProtoEncryAlgorithm,omitempty" tf:"ike_proto_encry_algorithm,omitempty"`

	// Remote address of IKE operation specification, valid when ike_remote_identity is ADDRESS, generally the value is public_ip_address of the related customer gateway.
	// Remote address of IKE operation specification, valid when ike_remote_identity is `ADDRESS`, generally the value is `public_ip_address` of the related customer gateway.
	// +kubebuilder:validation:Optional
	IkeRemoteAddress *string `json:"ikeRemoteAddress,omitempty" tf:"ike_remote_address,omitempty"`

	// Remote FQDN name of the IKE operation specification.
	// Remote FQDN name of the IKE operation specification.
	// +kubebuilder:validation:Optional
	IkeRemoteFqdnName *string `json:"ikeRemoteFqdnName,omitempty" tf:"ike_remote_fqdn_name,omitempty"`

	// Remote identity way of IKE operation specification. Valid values: ADDRESS, FQDN. Default value is ADDRESS.
	// Remote identity way of IKE operation specification. Valid values: `ADDRESS`, `FQDN`. Default value is `ADDRESS`.
	// +kubebuilder:validation:Optional
	IkeRemoteIdentity *string `json:"ikeRemoteIdentity,omitempty" tf:"ike_remote_identity,omitempty"`

	// SA lifetime of the IKE operation specification, unit is second. The value ranges from 60 to 604800. Default value is 86400 seconds.
	// SA lifetime of the IKE operation specification, unit is `second`. The value ranges from 60 to 604800. Default value is 86400 seconds.
	// +kubebuilder:validation:Optional
	IkeSaLifetimeSeconds *float64 `json:"ikeSaLifetimeSeconds,omitempty" tf:"ike_sa_lifetime_seconds,omitempty"`

	// Version of the IKE operation specification, values: IKEV1, IKEV2. Default value is IKEV1.
	// Version of the IKE operation specification, values: `IKEV1`, `IKEV2`. Default value is `IKEV1`.
	// +kubebuilder:validation:Optional
	IkeVersion *string `json:"ikeVersion,omitempty" tf:"ike_version,omitempty"`

	// Encrypt algorithm of the IPSEC operation specification. Valid values: 3DES-CBC, AES-CBC-128, AES-CBC-192, AES-CBC-256, DES-CBC, SM4, NULL, AES128GCM128, AES192GCM128, AES256GCM128. Default value is 3DES-CBC.
	// Encrypt algorithm of the IPSEC operation specification. Valid values: `3DES-CBC`, `AES-CBC-128`, `AES-CBC-192`, `AES-CBC-256`, `DES-CBC`, `SM4`, `NULL`, `AES128GCM128`, `AES192GCM128`, `AES256GCM128`. Default value is `3DES-CBC`.
	// +kubebuilder:validation:Optional
	IpsecEncryptAlgorithm *string `json:"ipsecEncryptAlgorithm,omitempty" tf:"ipsec_encrypt_algorithm,omitempty"`

	// Integrity algorithm of the IPSEC operation specification. Valid values: SHA1, MD5, SHA-256. Default value is MD5.
	// Integrity algorithm of the IPSEC operation specification. Valid values: `SHA1`, `MD5`, `SHA-256`. Default value is `MD5`.
	// +kubebuilder:validation:Optional
	IpsecIntegrityAlgorithm *string `json:"ipsecIntegrityAlgorithm,omitempty" tf:"ipsec_integrity_algorithm,omitempty"`

	// PFS DH group. Valid value: DH-GROUP1, DH-GROUP2, DH-GROUP5, DH-GROUP14, DH-GROUP24, NULL. Default value is NULL.
	// PFS DH group. Valid value: `DH-GROUP1`, `DH-GROUP2`, `DH-GROUP5`, `DH-GROUP14`, `DH-GROUP24`, `NULL`. Default value is `NULL`.
	// +kubebuilder:validation:Optional
	IpsecPfsDhGroup *string `json:"ipsecPfsDhGroup,omitempty" tf:"ipsec_pfs_dh_group,omitempty"`

	// SA lifetime of the IPSEC operation specification, unit is second. Valid value ranges: [180~604800]. Default value is 3600 seconds.
	// SA lifetime of the IPSEC operation specification, unit is second. Valid value ranges: [180~604800]. Default value is 3600 seconds.
	// +kubebuilder:validation:Optional
	IpsecSaLifetimeSeconds *float64 `json:"ipsecSaLifetimeSeconds,omitempty" tf:"ipsec_sa_lifetime_seconds,omitempty"`

	// SA lifetime of the IPSEC operation specification, unit is KB. The value should not be less then 2560. Default value is 1843200.
	// SA lifetime of the IPSEC operation specification, unit is KB. The value should not be less then 2560. Default value is 1843200.
	// +kubebuilder:validation:Optional
	IpsecSaLifetimeTraffic *float64 `json:"ipsecSaLifetimeTraffic,omitempty" tf:"ipsec_sa_lifetime_traffic,omitempty"`

	// Name of the VPN connection. The length of character is limited to 1-60.
	// Name of the VPN connection. The length of character is limited to 1-60.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Pre-shared key of the VPN connection.
	// Pre-shared key of the VPN connection.
	// +kubebuilder:validation:Optional
	PreShareKey *string `json:"preShareKey,omitempty" tf:"pre_share_key,omitempty"`

	// Route type of the VPN connection. Valid value: STATIC, StaticRoute, Policy.
	// Route type of the VPN connection. Valid value: `STATIC`, `StaticRoute`, `Policy`.
	// +kubebuilder:validation:Optional
	RouteType *string `json:"routeType,omitempty" tf:"route_type,omitempty"`

	// SPD policy group, for example: {"10.0.0.5/24":["172.123.10.5/16"]}, 10.0.0.5/24 is the vpc intranet segment, and 172.123.10.5/16 is the IDC network segment. Users specify which network segments in the VPC can communicate with which network segments in your IDC.
	// SPD policy group, for example: {"10.0.0.5/24":["172.123.10.5/16"]}, 10.0.0.5/24 is the vpc intranet segment, and 172.123.10.5/16 is the IDC network segment. Users specify which network segments in the VPC can communicate with which network segments in your IDC.
	// +kubebuilder:validation:Optional
	SecurityGroupPolicy []SecurityGroupPolicyParameters `json:"securityGroupPolicy,omitempty" tf:"security_group_policy,omitempty"`

	// A list of tags used to associate different resources.
	// A list of tags used to associate different resources.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// ID of the VPC. Required if vpn gateway is not in CCN type, and doesn't make sense for CCN vpn gateway.
	// ID of the VPC. Required if vpn gateway is not in `CCN` type, and doesn't make sense for `CCN` vpn gateway.
	// +crossplane:generate:reference:type=VPC
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Reference to a VPC to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// Selector for a VPC to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`

	// ID of the VPN gateway.
	// ID of the VPN gateway.
	// +crossplane:generate:reference:type=VPNGateway
	// +kubebuilder:validation:Optional
	VPNGatewayID *string `json:"vpnGatewayId,omitempty" tf:"vpn_gateway_id,omitempty"`

	// Reference to a VPNGateway to populate vpnGatewayId.
	// +kubebuilder:validation:Optional
	VPNGatewayIDRef *v1.Reference `json:"vpnGatewayIdRef,omitempty" tf:"-"`

	// Selector for a VPNGateway to populate vpnGatewayId.
	// +kubebuilder:validation:Optional
	VPNGatewayIDSelector *v1.Selector `json:"vpnGatewayIdSelector,omitempty" tf:"-"`
}

// VPNConnectionSpec defines the desired state of VPNConnection
type VPNConnectionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VPNConnectionParameters `json:"forProvider"`
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
	InitProvider VPNConnectionInitParameters `json:"initProvider,omitempty"`
}

// VPNConnectionStatus defines the observed state of VPNConnection.
type VPNConnectionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VPNConnectionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VPNConnection is the Schema for the VPNConnections API. Provides a resource to create a VPN connection.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloud}
type VPNConnection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.preShareKey) || (has(self.initProvider) && has(self.initProvider.preShareKey))",message="spec.forProvider.preShareKey is a required parameter"
	Spec   VPNConnectionSpec   `json:"spec"`
	Status VPNConnectionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VPNConnectionList contains a list of VPNConnections
type VPNConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VPNConnection `json:"items"`
}

// Repository type metadata.
var (
	VPNConnection_Kind             = "VPNConnection"
	VPNConnection_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VPNConnection_Kind}.String()
	VPNConnection_KindAPIVersion   = VPNConnection_Kind + "." + CRDGroupVersion.String()
	VPNConnection_GroupVersionKind = CRDGroupVersion.WithKind(VPNConnection_Kind)
)

func init() {
	SchemeBuilder.Register(&VPNConnection{}, &VPNConnectionList{})
}