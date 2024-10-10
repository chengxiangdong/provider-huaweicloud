// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	cluster "github.com/huaweicloud/provider-huaweicloud/internal/controller/cce/cluster"
	node "github.com/huaweicloud/provider-huaweicloud/internal/controller/cce/node"
	bucket "github.com/huaweicloud/provider-huaweicloud/internal/controller/obs/bucket"
	bucketacl "github.com/huaweicloud/provider-huaweicloud/internal/controller/obs/bucketacl"
	bucketobject "github.com/huaweicloud/provider-huaweicloud/internal/controller/obs/bucketobject"
	bucketobjectacl "github.com/huaweicloud/provider-huaweicloud/internal/controller/obs/bucketobjectacl"
	providerconfig "github.com/huaweicloud/provider-huaweicloud/internal/controller/providerconfig"
	secgroup "github.com/huaweicloud/provider-huaweicloud/internal/controller/vpc/secgroup"
	secgrouprule "github.com/huaweicloud/provider-huaweicloud/internal/controller/vpc/secgrouprule"
	subnet "github.com/huaweicloud/provider-huaweicloud/internal/controller/vpc/subnet"
	vpc "github.com/huaweicloud/provider-huaweicloud/internal/controller/vpc/vpc"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		cluster.Setup,
		node.Setup,
		bucket.Setup,
		bucketacl.Setup,
		bucketobject.Setup,
		bucketobjectacl.Setup,
		providerconfig.Setup,
		secgroup.Setup,
		secgrouprule.Setup,
		subnet.Setup,
		vpc.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
