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

package cos

import "github.com/crossplane/upjet/pkg/config"

const shortGroupCos = "cos"

// Configure configures the cos group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("tencentcloud_cos_bucket", func(r *config.Resource) {
		r.ShortGroup = shortGroupCos
		r.Kind = "Bucket"
	})

	p.AddResourceConfigurator("tencentcloud_cos_bucket_policy", func(r *config.Resource) {
		r.ShortGroup = shortGroupCos
		r.Kind = "BucketPolicy"
		r.References["bucket"] = config.Reference{
			Type: "Bucket",
		}
	})

	p.AddResourceConfigurator("tencentcloud_cos_bucket_object", func(r *config.Resource) {
		r.ShortGroup = shortGroupCos
		r.Kind = "BucketObject"
		r.References["bucket"] = config.Reference{
			Type: "Bucket",
		}
	})

	p.AddResourceConfigurator("tencentcloud_cos_bucket_domain_certificate_attachment", func(r *config.Resource) {
		r.ShortGroup = shortGroupCos
		r.Kind = "BucketDomainCertificateAttachment"
		r.References["bucket"] = config.Reference{
			Type: "Bucket",
		}
	})
}