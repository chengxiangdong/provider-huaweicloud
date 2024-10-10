package cce

import (
	"github.com/crossplane/upjet/pkg/config"
)

const shortGroupCce = "cce"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("huaweicloud_cce_cluster", func(r *config.Resource) {
		r.ShortGroup = shortGroupCce

		r.References["vpc_id"] = config.Reference{
			TerraformName: "huaweicloud_vpc",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}

		r.References["subnet_id"] = config.Reference{
			TerraformName: "huaweicloud_vpc_subnet",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_cce_node", func(r *config.Resource) {
		r.ShortGroup = shortGroupCce
		r.References["cluster_id"] = config.Reference{
			TerraformName: "huaweicloud_cce_cluster",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})
}
