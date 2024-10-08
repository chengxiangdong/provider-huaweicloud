/*
Copyright 2022 Upbound Inc.
*/

package config

import "github.com/crossplane/upjet/pkg/config"

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	// Import requires using a randomly generated ID from provider: nl-2e21sda
	// vpc
	"tencentcloud_vpc":                              config.IdentifierFromProvider,
	"tencentcloud_vpc_acl":                          config.IdentifierFromProvider,
	"tencentcloud_vpc_acl_attachment":               config.IdentifierFromProvider,
	"tencentcloud_subnet":                           config.IdentifierFromProvider,
	"tencentcloud_eip":                              config.IdentifierFromProvider,
	"tencentcloud_eip_association":                  config.IdentifierFromProvider,
	"tencentcloud_ha_vip":                           config.IdentifierFromProvider,
	"tencentcloud_security_group":                   config.IdentifierFromProvider,
	"tencentcloud_security_group_rule":              config.IdentifierFromProvider,
	"tencentcloud_security_group_lite_rule":         config.IdentifierFromProvider,
	"tencentcloud_address_template":                 config.IdentifierFromProvider,
	"tencentcloud_address_template_group":           config.IdentifierFromProvider,
	"tencentcloud_protocol_template":                config.IdentifierFromProvider,
	"tencentcloud_protocol_template_group":          config.IdentifierFromProvider,
	"tencentcloud_route_table":                      config.IdentifierFromProvider,
	"tencentcloud_route_table_entry":                config.IdentifierFromProvider,
	"tencentcloud_route_entry":                      config.IdentifierFromProvider,
	"tencentcloud_nat_gateway":                      config.IdentifierFromProvider,
	"tencentcloud_nat_gateway_snat":                 config.IdentifierFromProvider,
	"tencentcloud_dnat":                             config.IdentifierFromProvider,
	"tencentcloud_vpc_bandwidth_package":            config.IdentifierFromProvider,
	"tencentcloud_vpc_bandwidth_package_attachment": config.IdentifierFromProvider,
	"tencentcloud_vpn_gateway":                      config.IdentifierFromProvider,
	"tencentcloud_vpn_ssl_client":                   config.IdentifierFromProvider,
	"tencentcloud_vpn_customer_gateway":             config.IdentifierFromProvider,
	"tencentcloud_vpn_connection":                   config.IdentifierFromProvider,
	"tencentcloud_vpn_gateway_route":                config.IdentifierFromProvider,
	"tencentcloud_vpn_ssl_server":                   config.IdentifierFromProvider,
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
