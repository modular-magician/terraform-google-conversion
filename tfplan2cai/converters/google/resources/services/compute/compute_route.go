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

package compute

import (
	"fmt"
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgresource"
	transport_tpg "github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/transport"
)

const ComputeRouteAssetType string = "compute.googleapis.com/Route"

func ResourceConverterComputeRoute() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType: ComputeRouteAssetType,
		Convert:   GetComputeRouteCaiObject,
	}
}

func GetComputeRouteCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	name, err := tpgresource.AssetName(d, config, "//compute.googleapis.com/projects/{{project}}/global/routes/{{name}}")
	if err != nil {
		return []tpgresource.Asset{}, err
	}
	if obj, err := GetComputeRouteApiObject(d, config); err == nil {
		return []tpgresource.Asset{{
			Name: name,
			Type: ComputeRouteAssetType,
			Resource: &tpgresource.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/compute/v1/rest",
				DiscoveryName:        "Route",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []tpgresource.Asset{}, err
	}
}

func GetComputeRouteApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	destRangeProp, err := expandComputeRouteDestRange(d.Get("dest_range"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("dest_range"); !tpgresource.IsEmptyValue(reflect.ValueOf(destRangeProp)) && (ok || !reflect.DeepEqual(v, destRangeProp)) {
		obj["destRange"] = destRangeProp
	}
	descriptionProp, err := expandComputeRouteDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	nameProp, err := expandComputeRouteName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	networkProp, err := expandComputeRouteNetwork(d.Get("network"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("network"); !tpgresource.IsEmptyValue(reflect.ValueOf(networkProp)) && (ok || !reflect.DeepEqual(v, networkProp)) {
		obj["network"] = networkProp
	}
	priorityProp, err := expandComputeRoutePriority(d.Get("priority"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("priority"); ok || !reflect.DeepEqual(v, priorityProp) {
		obj["priority"] = priorityProp
	}
	tagsProp, err := expandComputeRouteTags(d.Get("tags"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("tags"); !tpgresource.IsEmptyValue(reflect.ValueOf(tagsProp)) && (ok || !reflect.DeepEqual(v, tagsProp)) {
		obj["tags"] = tagsProp
	}
	nextHopGatewayProp, err := expandComputeRouteNextHopGateway(d.Get("next_hop_gateway"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("next_hop_gateway"); !tpgresource.IsEmptyValue(reflect.ValueOf(nextHopGatewayProp)) && (ok || !reflect.DeepEqual(v, nextHopGatewayProp)) {
		obj["nextHopGateway"] = nextHopGatewayProp
	}
	nextHopInstanceProp, err := expandComputeRouteNextHopInstance(d.Get("next_hop_instance"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("next_hop_instance"); !tpgresource.IsEmptyValue(reflect.ValueOf(nextHopInstanceProp)) && (ok || !reflect.DeepEqual(v, nextHopInstanceProp)) {
		obj["nextHopInstance"] = nextHopInstanceProp
	}
	nextHopIpProp, err := expandComputeRouteNextHopIp(d.Get("next_hop_ip"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("next_hop_ip"); !tpgresource.IsEmptyValue(reflect.ValueOf(nextHopIpProp)) && (ok || !reflect.DeepEqual(v, nextHopIpProp)) {
		obj["nextHopIp"] = nextHopIpProp
	}
	nextHopVpnTunnelProp, err := expandComputeRouteNextHopVpnTunnel(d.Get("next_hop_vpn_tunnel"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("next_hop_vpn_tunnel"); !tpgresource.IsEmptyValue(reflect.ValueOf(nextHopVpnTunnelProp)) && (ok || !reflect.DeepEqual(v, nextHopVpnTunnelProp)) {
		obj["nextHopVpnTunnel"] = nextHopVpnTunnelProp
	}
	nextHopIlbProp, err := expandComputeRouteNextHopIlb(d.Get("next_hop_ilb"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("next_hop_ilb"); !tpgresource.IsEmptyValue(reflect.ValueOf(nextHopIlbProp)) && (ok || !reflect.DeepEqual(v, nextHopIlbProp)) {
		obj["nextHopIlb"] = nextHopIlbProp
	}

	return obj, nil
}

func expandComputeRouteDestRange(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRouteDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRouteName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRouteNetwork(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	f, err := tpgresource.ParseGlobalFieldValue("networks", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for network: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeRoutePriority(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRouteTags(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v.(*schema.Set).List(), nil
}

func expandComputeRouteNextHopGateway(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	if v == "default-internet-gateway" {
		return tpgresource.ReplaceVars(d, config, "projects/{{project}}/global/gateways/default-internet-gateway")
	} else {
		return v, nil
	}
}

func expandComputeRouteNextHopInstance(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	if v == "" {
		return v, nil
	}
	val, err := tpgresource.ParseZonalFieldValue("instances", v.(string), "project", "next_hop_instance_zone", d, config, true)
	if err != nil {
		return nil, err
	}

	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return nil, err
	}

	nextInstance, err := config.NewComputeClient(userAgent).Instances.Get(val.Project, val.Zone, val.Name).Do()
	if err != nil {
		return nil, err
	}
	return nextInstance.SelfLink, nil
}

func expandComputeRouteNextHopIp(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRouteNextHopVpnTunnel(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	f, err := tpgresource.ParseRegionalFieldValue("vpnTunnels", v.(string), "project", "region", "zone", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for next_hop_vpn_tunnel: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeRouteNextHopIlb(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
