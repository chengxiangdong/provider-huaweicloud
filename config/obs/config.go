package obs

import (
	"context"
	"fmt"

	"github.com/crossplane/upjet/pkg/config"
)

const shortGroupObs = "obs"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("huaweicloud_obs_bucket", func(r *config.Resource) {
		r.ShortGroup = shortGroupObs
		r.ExternalName = config.IdentifierFromProvider
	})

	p.AddResourceConfigurator("huaweicloud_obs_bucket_object", func(r *config.Resource) {
		r.ShortGroup = shortGroupObs
		r.ExternalName = config.IdentifierFromProvider
		r.ExternalName.GetIDFn = getFullyQualifiedIDfunc

		r.References["bucket"] = config.Reference{
			TerraformName: "huaweicloud_obs_bucket",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractParamPath("bucket",true)`,
		}
	})
}

func getFullyQualifiedIDfunc(_ context.Context, _ string, _ map[string]any, providerConfig map[string]any) (string, error) {
	bucket, ok := providerConfig["bucket"]
	if !ok {
		return "", fmt.Errorf("not found attribute bucket")
	}
	key, ok := providerConfig["key"]
	if !ok {
		return "", fmt.Errorf("not found attribute key")
	}
	return fmt.Sprintf("%s/%s", bucket, key), nil
}
