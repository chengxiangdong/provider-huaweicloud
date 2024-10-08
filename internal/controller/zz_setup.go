package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	providerconfig "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/providerconfig"
	addresstemplate "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vpc/addresstemplate"
	addresstemplategroup "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vpc/addresstemplategroup"
	dnat "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vpc/dnat"
	eip "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vpc/eip"
	eipassociation "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vpc/eipassociation"
	havip "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vpc/havip"
	natgateway "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vpc/natgateway"
	natgatewaysnat "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vpc/natgatewaysnat"
	protocoltemplate "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vpc/protocoltemplate"
	protocoltemplategroup "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vpc/protocoltemplategroup"
	routeentry "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vpc/routeentry"
	routetable "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vpc/routetable"
	routetableentry "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vpc/routetableentry"
	securitygroup "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vpc/securitygroup"
	securitygroupliterule "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vpc/securitygroupliterule"
	securitygrouprule "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vpc/securitygrouprule"
	subnet "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vpc/subnet"
	vpc "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vpc/vpc"
	vpcacl "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vpc/vpcacl"
	vpcbandwidthpackage "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vpc/vpcbandwidthpackage"
	vpcbandwidthpackageattachment "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vpc/vpcbandwidthpackageattachment"
	vpnconnection "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vpc/vpnconnection"
	vpncustomergateway "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vpc/vpncustomergateway"
	vpngateway "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vpc/vpngateway"
	vpngatewayroute "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vpc/vpngatewayroute"
	vpnsslclient "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vpc/vpnsslclient"
	vpnsslserver "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vpc/vpnsslserver"
	vpvaclattachment "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vpc/vpvaclattachment"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		providerconfig.Setup,
		addresstemplate.Setup,
		addresstemplategroup.Setup,
		dnat.Setup,
		eip.Setup,
		eipassociation.Setup,
		havip.Setup,
		natgateway.Setup,
		natgatewaysnat.Setup,
		protocoltemplate.Setup,
		protocoltemplategroup.Setup,
		routeentry.Setup,
		routetable.Setup,
		routetableentry.Setup,
		securitygroup.Setup,
		securitygroupliterule.Setup,
		securitygrouprule.Setup,
		subnet.Setup,
		vpc.Setup,
		vpcacl.Setup,
		vpcbandwidthpackage.Setup,
		vpcbandwidthpackageattachment.Setup,
		vpnconnection.Setup,
		vpncustomergateway.Setup,
		vpngateway.Setup,
		vpngatewayroute.Setup,
		vpnsslclient.Setup,
		vpnsslserver.Setup,
		vpvaclattachment.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
