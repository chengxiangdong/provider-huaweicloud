// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	secgroup "github.com/huaweicloud/provider-huaweicloud/internal/controller/networking/secgroup"
	bucket "github.com/huaweicloud/provider-huaweicloud/internal/controller/obs/bucket"
	bucketobject "github.com/huaweicloud/provider-huaweicloud/internal/controller/obs/bucketobject"
	providerconfig "github.com/huaweicloud/provider-huaweicloud/internal/controller/providerconfig"
	subnet "github.com/huaweicloud/provider-huaweicloud/internal/controller/vpc/subnet"
	vpc "github.com/huaweicloud/provider-huaweicloud/internal/controller/vpc/vpc"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		secgroup.Setup,
		bucket.Setup,
		bucketobject.Setup,
		providerconfig.Setup,
		subnet.Setup,
		vpc.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
