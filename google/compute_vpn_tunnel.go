// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"bytes"
	"fmt"
	"net"
	"reflect"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// validatePeerAddr returns false if a tunnel's peer_ip property
// is invalid. Currently, only addresses that collide with RFC
// 5735 (https://tools.ietf.org/html/rfc5735) fail validation.
func validatePeerAddr(i interface{}, val string) ([]string, []error) {
	ip := net.ParseIP(i.(string))
	if ip == nil {
		return nil, []error{fmt.Errorf("could not parse %q to IP address", val)}
	}
	for _, test := range invalidPeerAddrs {
		if bytes.Compare(ip, test.from) >= 0 && bytes.Compare(ip, test.to) <= 0 {
			return nil, []error{fmt.Errorf("address is invalid (is between %q and %q, conflicting with RFC5735)", test.from, test.to)}
		}
	}
	return nil, nil
}

// invalidPeerAddrs is a collection of IP address ranges that represent
// a conflict with RFC 5735 (https://tools.ietf.org/html/rfc5735#page-3).
// CIDR range notations in the RFC were converted to a (from, to) pair
// for easy checking with bytes.Compare.
var invalidPeerAddrs = []struct {
	from net.IP
	to   net.IP
}{
	{
		from: net.ParseIP("0.0.0.0"),
		to:   net.ParseIP("0.255.255.255"),
	},
	{
		from: net.ParseIP("10.0.0.0"),
		to:   net.ParseIP("10.255.255.255"),
	},
	{
		from: net.ParseIP("127.0.0.0"),
		to:   net.ParseIP("127.255.255.255"),
	},
	{
		from: net.ParseIP("169.254.0.0"),
		to:   net.ParseIP("169.254.255.255"),
	},
	{
		from: net.ParseIP("172.16.0.0"),
		to:   net.ParseIP("172.31.255.255"),
	},
	{
		from: net.ParseIP("192.0.0.0"),
		to:   net.ParseIP("192.0.0.255"),
	},
	{
		from: net.ParseIP("192.0.2.0"),
		to:   net.ParseIP("192.0.2.255"),
	},
	{
		from: net.ParseIP("192.88.99.0"),
		to:   net.ParseIP("192.88.99.255"),
	},
	{
		from: net.ParseIP("192.168.0.0"),
		to:   net.ParseIP("192.168.255.255"),
	},
	{
		from: net.ParseIP("198.18.0.0"),
		to:   net.ParseIP("198.19.255.255"),
	},
	{
		from: net.ParseIP("198.51.100.0"),
		to:   net.ParseIP("198.51.100.255"),
	},
	{
		from: net.ParseIP("203.0.113.0"),
		to:   net.ParseIP("203.0.113.255"),
	},
	{
		from: net.ParseIP("224.0.0.0"),
		to:   net.ParseIP("239.255.255.255"),
	},
	{
		from: net.ParseIP("240.0.0.0"),
		to:   net.ParseIP("255.255.255.255"),
	},
	{
		from: net.ParseIP("255.255.255.255"),
		to:   net.ParseIP("255.255.255.255"),
	},
}

func getVpnTunnelLink(config *Config, project, region, tunnel, userAgent string) (string, error) {
	if !strings.Contains(tunnel, "/") {
		// Tunnel value provided is just the name, lookup the tunnel SelfLink
		tunnelData, err := config.NewComputeClient(userAgent).VpnTunnels.Get(
			project, region, tunnel).Do()
		if err != nil {
			return "", fmt.Errorf("Error reading tunnel: %s", err)
		}
		tunnel = tunnelData.SelfLink
	}

	return tunnel, nil

}

func GetComputeVpnTunnelCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	name, err := assetName(d, config, "//compute.googleapis.com/projects/{{project}}/regions/{{region}}/vpnTunnels/{{name}}")
	if err != nil {
		return []Asset{}, err
	}
	if obj, err := GetComputeVpnTunnelApiObject(d, config); err == nil {
		return []Asset{{
			Name: name,
			Type: "compute.googleapis.com/VpnTunnel",
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/compute/v1/rest",
				DiscoveryName:        "VpnTunnel",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []Asset{}, err
	}
}

func GetComputeVpnTunnelApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandComputeVpnTunnelName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	descriptionProp, err := expandComputeVpnTunnelDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	targetVpnGatewayProp, err := expandComputeVpnTunnelTargetVpnGateway(d.Get("target_vpn_gateway"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("target_vpn_gateway"); !isEmptyValue(reflect.ValueOf(targetVpnGatewayProp)) && (ok || !reflect.DeepEqual(v, targetVpnGatewayProp)) {
		obj["targetVpnGateway"] = targetVpnGatewayProp
	}
	vpnGatewayProp, err := expandComputeVpnTunnelVpnGateway(d.Get("vpn_gateway"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("vpn_gateway"); !isEmptyValue(reflect.ValueOf(vpnGatewayProp)) && (ok || !reflect.DeepEqual(v, vpnGatewayProp)) {
		obj["vpnGateway"] = vpnGatewayProp
	}
	vpnGatewayInterfaceProp, err := expandComputeVpnTunnelVpnGatewayInterface(d.Get("vpn_gateway_interface"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("vpn_gateway_interface"); ok || !reflect.DeepEqual(v, vpnGatewayInterfaceProp) {
		obj["vpnGatewayInterface"] = vpnGatewayInterfaceProp
	}
	peerExternalGatewayProp, err := expandComputeVpnTunnelPeerExternalGateway(d.Get("peer_external_gateway"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("peer_external_gateway"); !isEmptyValue(reflect.ValueOf(peerExternalGatewayProp)) && (ok || !reflect.DeepEqual(v, peerExternalGatewayProp)) {
		obj["peerExternalGateway"] = peerExternalGatewayProp
	}
	peerExternalGatewayInterfaceProp, err := expandComputeVpnTunnelPeerExternalGatewayInterface(d.Get("peer_external_gateway_interface"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("peer_external_gateway_interface"); ok || !reflect.DeepEqual(v, peerExternalGatewayInterfaceProp) {
		obj["peerExternalGatewayInterface"] = peerExternalGatewayInterfaceProp
	}
	peerGcpGatewayProp, err := expandComputeVpnTunnelPeerGcpGateway(d.Get("peer_gcp_gateway"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("peer_gcp_gateway"); !isEmptyValue(reflect.ValueOf(peerGcpGatewayProp)) && (ok || !reflect.DeepEqual(v, peerGcpGatewayProp)) {
		obj["peerGcpGateway"] = peerGcpGatewayProp
	}
	routerProp, err := expandComputeVpnTunnelRouter(d.Get("router"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("router"); !isEmptyValue(reflect.ValueOf(routerProp)) && (ok || !reflect.DeepEqual(v, routerProp)) {
		obj["router"] = routerProp
	}
	peerIpProp, err := expandComputeVpnTunnelPeerIp(d.Get("peer_ip"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("peer_ip"); !isEmptyValue(reflect.ValueOf(peerIpProp)) && (ok || !reflect.DeepEqual(v, peerIpProp)) {
		obj["peerIp"] = peerIpProp
	}
	sharedSecretProp, err := expandComputeVpnTunnelSharedSecret(d.Get("shared_secret"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("shared_secret"); !isEmptyValue(reflect.ValueOf(sharedSecretProp)) && (ok || !reflect.DeepEqual(v, sharedSecretProp)) {
		obj["sharedSecret"] = sharedSecretProp
	}
	ikeVersionProp, err := expandComputeVpnTunnelIkeVersion(d.Get("ike_version"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("ike_version"); !isEmptyValue(reflect.ValueOf(ikeVersionProp)) && (ok || !reflect.DeepEqual(v, ikeVersionProp)) {
		obj["ikeVersion"] = ikeVersionProp
	}
	localTrafficSelectorProp, err := expandComputeVpnTunnelLocalTrafficSelector(d.Get("local_traffic_selector"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("local_traffic_selector"); !isEmptyValue(reflect.ValueOf(localTrafficSelectorProp)) && (ok || !reflect.DeepEqual(v, localTrafficSelectorProp)) {
		obj["localTrafficSelector"] = localTrafficSelectorProp
	}
	remoteTrafficSelectorProp, err := expandComputeVpnTunnelRemoteTrafficSelector(d.Get("remote_traffic_selector"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("remote_traffic_selector"); !isEmptyValue(reflect.ValueOf(remoteTrafficSelectorProp)) && (ok || !reflect.DeepEqual(v, remoteTrafficSelectorProp)) {
		obj["remoteTrafficSelector"] = remoteTrafficSelectorProp
	}
	regionProp, err := expandComputeVpnTunnelRegion(d.Get("region"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("region"); !isEmptyValue(reflect.ValueOf(regionProp)) && (ok || !reflect.DeepEqual(v, regionProp)) {
		obj["region"] = regionProp
	}

	return resourceComputeVpnTunnelEncoder(d, config, obj)
}

func resourceComputeVpnTunnelEncoder(d TerraformResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	config := meta.(*Config)
	f, err := parseRegionalFieldValue("targetVpnGateways", d.Get("target_vpn_gateway").(string), "project", "region", "zone", d, config, true)
	if err != nil {
		return nil, err
	}
	if _, ok := d.GetOk("project"); !ok {
		if err := d.Set("project", f.Project); err != nil {
			return nil, fmt.Errorf("Error setting project: %s", err)
		}
	}
	if _, ok := d.GetOk("region"); !ok {
		if err := d.Set("region", f.Region); err != nil {
			return nil, fmt.Errorf("Error setting region: %s", err)
		}
	}
	return obj, nil
}

func expandComputeVpnTunnelName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeVpnTunnelDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeVpnTunnelTargetVpnGateway(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseRegionalFieldValue("targetVpnGateways", v.(string), "project", "region", "zone", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for target_vpn_gateway: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeVpnTunnelVpnGateway(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseRegionalFieldValue("vpnGateways", v.(string), "project", "region", "zone", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for vpn_gateway: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeVpnTunnelVpnGatewayInterface(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeVpnTunnelPeerExternalGateway(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseGlobalFieldValue("externalVpnGateways", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for peer_external_gateway: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeVpnTunnelPeerExternalGatewayInterface(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeVpnTunnelPeerGcpGateway(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseRegionalFieldValue("vpnGateways", v.(string), "project", "region", "zone", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for peer_gcp_gateway: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeVpnTunnelRouter(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	if v == nil || v.(string) == "" {
		return "", nil
	}
	f, err := parseRegionalFieldValue("routers", v.(string), "project", "region", "zone", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for router: %s", err)
	}

	url, err := replaceVars(d, config, "{{ComputeBasePath}}"+f.RelativeLink())
	if err != nil {
		return nil, err
	}

	return url, nil
}

func expandComputeVpnTunnelPeerIp(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeVpnTunnelSharedSecret(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeVpnTunnelIkeVersion(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeVpnTunnelLocalTrafficSelector(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	return v, nil
}

func expandComputeVpnTunnelRemoteTrafficSelector(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	return v, nil
}

func expandComputeVpnTunnelRegion(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseGlobalFieldValue("regions", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for region: %s", err)
	}
	return f.RelativeLink(), nil
}
