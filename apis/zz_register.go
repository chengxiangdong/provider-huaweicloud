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

// Package apis contains Kubernetes API for the provider.
package apis

import (
	"k8s.io/apimachinery/pkg/runtime"

	v1alpha1 "github.com/crossplane-contrib/provider-tencentcloud/apis/apigateway/v1alpha1"
	v1alpha1as "github.com/crossplane-contrib/provider-tencentcloud/apis/as/v1alpha1"
	v1alpha1audit "github.com/crossplane-contrib/provider-tencentcloud/apis/audit/v1alpha1"
	v1alpha1cam "github.com/crossplane-contrib/provider-tencentcloud/apis/cam/v1alpha1"
	v1alpha1cat "github.com/crossplane-contrib/provider-tencentcloud/apis/cat/v1alpha1"
	v1alpha1cbs "github.com/crossplane-contrib/provider-tencentcloud/apis/cbs/v1alpha1"
	v1alpha1ccn "github.com/crossplane-contrib/provider-tencentcloud/apis/ccn/v1alpha1"
	v1alpha1cdh "github.com/crossplane-contrib/provider-tencentcloud/apis/cdh/v1alpha1"
	v1alpha1cdn "github.com/crossplane-contrib/provider-tencentcloud/apis/cdn/v1alpha1"
	v1alpha1cfs "github.com/crossplane-contrib/provider-tencentcloud/apis/cfs/v1alpha1"
	v1alpha1clb "github.com/crossplane-contrib/provider-tencentcloud/apis/clb/v1alpha1"
	v1alpha1cls "github.com/crossplane-contrib/provider-tencentcloud/apis/cls/v1alpha1"
	v1alpha1containercluster "github.com/crossplane-contrib/provider-tencentcloud/apis/containercluster/v1alpha1"
	v1alpha1cos "github.com/crossplane-contrib/provider-tencentcloud/apis/cos/v1alpha1"
	v1alpha1cvm "github.com/crossplane-contrib/provider-tencentcloud/apis/cvm/v1alpha1"
	v1alpha1cynosdb "github.com/crossplane-contrib/provider-tencentcloud/apis/cynosdb/v1alpha1"
	v1alpha1dayu "github.com/crossplane-contrib/provider-tencentcloud/apis/dayu/v1alpha1"
	v1alpha1dc "github.com/crossplane-contrib/provider-tencentcloud/apis/dc/v1alpha1"
	v1alpha1dcdb "github.com/crossplane-contrib/provider-tencentcloud/apis/dcdb/v1alpha1"
	v1alpha1dnspod "github.com/crossplane-contrib/provider-tencentcloud/apis/dnspod/v1alpha1"
	v1alpha1elasticsearch "github.com/crossplane-contrib/provider-tencentcloud/apis/elasticsearch/v1alpha1"
	v1alpha1emr "github.com/crossplane-contrib/provider-tencentcloud/apis/emr/v1alpha1"
	v1alpha1eni "github.com/crossplane-contrib/provider-tencentcloud/apis/eni/v1alpha1"
	v1alpha1gaap "github.com/crossplane-contrib/provider-tencentcloud/apis/gaap/v1alpha1"
	v1alpha1kafka "github.com/crossplane-contrib/provider-tencentcloud/apis/kafka/v1alpha1"
	v1alpha1kms "github.com/crossplane-contrib/provider-tencentcloud/apis/kms/v1alpha1"
	v1alpha1lighthouse "github.com/crossplane-contrib/provider-tencentcloud/apis/lighthouse/v1alpha1"
	v1alpha1mariadb "github.com/crossplane-contrib/provider-tencentcloud/apis/mariadb/v1alpha1"
	v1alpha1mongodb "github.com/crossplane-contrib/provider-tencentcloud/apis/mongodb/v1alpha1"
	v1alpha1monitor "github.com/crossplane-contrib/provider-tencentcloud/apis/monitor/v1alpha1"
	v1alpha1mysql "github.com/crossplane-contrib/provider-tencentcloud/apis/mysql/v1alpha1"
	v1alpha1postgresql "github.com/crossplane-contrib/provider-tencentcloud/apis/postgresql/v1alpha1"
	v1alpha1privatedns "github.com/crossplane-contrib/provider-tencentcloud/apis/privatedns/v1alpha1"
	v1alpha1redis "github.com/crossplane-contrib/provider-tencentcloud/apis/redis/v1alpha1"
	v1alpha1scf "github.com/crossplane-contrib/provider-tencentcloud/apis/scf/v1alpha1"
	v1alpha1ses "github.com/crossplane-contrib/provider-tencentcloud/apis/ses/v1alpha1"
	v1alpha1sms "github.com/crossplane-contrib/provider-tencentcloud/apis/sms/v1alpha1"
	v1alpha1sqlserver "github.com/crossplane-contrib/provider-tencentcloud/apis/sqlserver/v1alpha1"
	v1alpha1ssl "github.com/crossplane-contrib/provider-tencentcloud/apis/ssl/v1alpha1"
	v1alpha1ssm "github.com/crossplane-contrib/provider-tencentcloud/apis/ssm/v1alpha1"
	v1alpha1tcaplus "github.com/crossplane-contrib/provider-tencentcloud/apis/tcaplus/v1alpha1"
	v1alpha1tcm "github.com/crossplane-contrib/provider-tencentcloud/apis/tcm/v1alpha1"
	v1alpha1tcr "github.com/crossplane-contrib/provider-tencentcloud/apis/tcr/v1alpha1"
	v1alpha1tdmq "github.com/crossplane-contrib/provider-tencentcloud/apis/tdmq/v1alpha1"
	v1alpha1tem "github.com/crossplane-contrib/provider-tencentcloud/apis/tem/v1alpha1"
	v1alpha1teo "github.com/crossplane-contrib/provider-tencentcloud/apis/teo/v1alpha1"
	v1alpha1tke "github.com/crossplane-contrib/provider-tencentcloud/apis/tke/v1alpha1"
	v1alpha1apis "github.com/crossplane-contrib/provider-tencentcloud/apis/v1alpha1"
	v1alpha1vod "github.com/crossplane-contrib/provider-tencentcloud/apis/vod/v1alpha1"
	v1alpha1vpc "github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes,
		v1alpha1.SchemeBuilder.AddToScheme,
		v1alpha1as.SchemeBuilder.AddToScheme,
		v1alpha1audit.SchemeBuilder.AddToScheme,
		v1alpha1cam.SchemeBuilder.AddToScheme,
		v1alpha1cat.SchemeBuilder.AddToScheme,
		v1alpha1cbs.SchemeBuilder.AddToScheme,
		v1alpha1ccn.SchemeBuilder.AddToScheme,
		v1alpha1cdh.SchemeBuilder.AddToScheme,
		v1alpha1cdn.SchemeBuilder.AddToScheme,
		v1alpha1cfs.SchemeBuilder.AddToScheme,
		v1alpha1clb.SchemeBuilder.AddToScheme,
		v1alpha1cls.SchemeBuilder.AddToScheme,
		v1alpha1containercluster.SchemeBuilder.AddToScheme,
		v1alpha1cos.SchemeBuilder.AddToScheme,
		v1alpha1cvm.SchemeBuilder.AddToScheme,
		v1alpha1cynosdb.SchemeBuilder.AddToScheme,
		v1alpha1dayu.SchemeBuilder.AddToScheme,
		v1alpha1dc.SchemeBuilder.AddToScheme,
		v1alpha1dcdb.SchemeBuilder.AddToScheme,
		v1alpha1dnspod.SchemeBuilder.AddToScheme,
		v1alpha1elasticsearch.SchemeBuilder.AddToScheme,
		v1alpha1emr.SchemeBuilder.AddToScheme,
		v1alpha1eni.SchemeBuilder.AddToScheme,
		v1alpha1gaap.SchemeBuilder.AddToScheme,
		v1alpha1kafka.SchemeBuilder.AddToScheme,
		v1alpha1kms.SchemeBuilder.AddToScheme,
		v1alpha1lighthouse.SchemeBuilder.AddToScheme,
		v1alpha1mariadb.SchemeBuilder.AddToScheme,
		v1alpha1mongodb.SchemeBuilder.AddToScheme,
		v1alpha1monitor.SchemeBuilder.AddToScheme,
		v1alpha1mysql.SchemeBuilder.AddToScheme,
		v1alpha1postgresql.SchemeBuilder.AddToScheme,
		v1alpha1privatedns.SchemeBuilder.AddToScheme,
		v1alpha1redis.SchemeBuilder.AddToScheme,
		v1alpha1scf.SchemeBuilder.AddToScheme,
		v1alpha1ses.SchemeBuilder.AddToScheme,
		v1alpha1sms.SchemeBuilder.AddToScheme,
		v1alpha1sqlserver.SchemeBuilder.AddToScheme,
		v1alpha1ssl.SchemeBuilder.AddToScheme,
		v1alpha1ssm.SchemeBuilder.AddToScheme,
		v1alpha1tcaplus.SchemeBuilder.AddToScheme,
		v1alpha1tcm.SchemeBuilder.AddToScheme,
		v1alpha1tcr.SchemeBuilder.AddToScheme,
		v1alpha1tdmq.SchemeBuilder.AddToScheme,
		v1alpha1tem.SchemeBuilder.AddToScheme,
		v1alpha1teo.SchemeBuilder.AddToScheme,
		v1alpha1tke.SchemeBuilder.AddToScheme,
		v1alpha1apis.SchemeBuilder.AddToScheme,
		v1alpha1vod.SchemeBuilder.AddToScheme,
		v1alpha1vpc.SchemeBuilder.AddToScheme,
	)
}

// AddToSchemes may be used to add all resources defined in the project to a Scheme
var AddToSchemes runtime.SchemeBuilder

// AddToScheme adds all Resources to the Scheme
func AddToScheme(s *runtime.Scheme) error {
	return AddToSchemes.AddToScheme(s)
}
