package vpc

import (
	"github.com/crossplane/upjet/pkg/config"
)

const shortGroupVpc = "vpc"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("huaweicloud_vpc", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "vpc"
		r.ShortGroup = shortGroupVpc
	})

	p.AddResourceConfigurator("huaweicloud_vpc_subnet", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		r.ShortGroup = shortGroupVpc

		r.References["vpc_id"] = config.Reference{
			TerraformName: "huaweicloud_vpc",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractParamPath("id",true)`,
		}
	})
}
