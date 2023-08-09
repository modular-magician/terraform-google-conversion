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

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const ComputeHaVpnGatewayAssetType string = "compute.googleapis.com/HaVpnGateway"

func ResourceConverterComputeHaVpnGateway() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType: ComputeHaVpnGatewayAssetType,
		Convert:   GetComputeHaVpnGatewayCaiObject,
	}
}

func GetComputeHaVpnGatewayCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	name, err := tpgresource.AssetName(d, config, "//compute.googleapis.com/projects/{{project}}/regions/{{region}}/vpnGateways/{{name}}")
	if err != nil {
		return []tpgresource.Asset{}, err
	}
	if obj, err := GetComputeHaVpnGatewayApiObject(d, config); err == nil {
		return []tpgresource.Asset{{
			Name: name,
			Type: ComputeHaVpnGatewayAssetType,
			Resource: &tpgresource.AssetResource{
				Version:              "beta",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/compute/beta/rest",
				DiscoveryName:        "HaVpnGateway",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []tpgresource.Asset{}, err
	}
}

func GetComputeHaVpnGatewayApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	descriptionProp, err := expandComputeHaVpnGatewayDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	nameProp, err := expandComputeHaVpnGatewayName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	networkProp, err := expandComputeHaVpnGatewayNetwork(d.Get("network"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("network"); !tpgresource.IsEmptyValue(reflect.ValueOf(networkProp)) && (ok || !reflect.DeepEqual(v, networkProp)) {
		obj["network"] = networkProp
	}
	stackTypeProp, err := expandComputeHaVpnGatewayStackType(d.Get("stack_type"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("stack_type"); !tpgresource.IsEmptyValue(reflect.ValueOf(stackTypeProp)) && (ok || !reflect.DeepEqual(v, stackTypeProp)) {
		obj["stackType"] = stackTypeProp
	}
	vpnInterfacesProp, err := expandComputeHaVpnGatewayVpnInterfaces(d.Get("vpn_interfaces"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("vpn_interfaces"); !tpgresource.IsEmptyValue(reflect.ValueOf(vpnInterfacesProp)) && (ok || !reflect.DeepEqual(v, vpnInterfacesProp)) {
		obj["vpnInterfaces"] = vpnInterfacesProp
	}
	regionProp, err := expandComputeHaVpnGatewayRegion(d.Get("region"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("region"); !tpgresource.IsEmptyValue(reflect.ValueOf(regionProp)) && (ok || !reflect.DeepEqual(v, regionProp)) {
		obj["region"] = regionProp
	}

	return obj, nil
}

func expandComputeHaVpnGatewayDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeHaVpnGatewayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeHaVpnGatewayNetwork(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	f, err := tpgresource.ParseGlobalFieldValue("networks", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for network: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeHaVpnGatewayStackType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeHaVpnGatewayVpnInterfaces(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedId, err := expandComputeHaVpnGatewayVpnInterfacesId(original["id"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["id"] = transformedId
		}

		transformedIpAddress, err := expandComputeHaVpnGatewayVpnInterfacesIpAddress(original["ip_address"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedIpAddress); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["ipAddress"] = transformedIpAddress
		}

		transformedInterconnectAttachment, err := expandComputeHaVpnGatewayVpnInterfacesInterconnectAttachment(original["interconnect_attachment"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedInterconnectAttachment); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["interconnectAttachment"] = transformedInterconnectAttachment
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandComputeHaVpnGatewayVpnInterfacesId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeHaVpnGatewayVpnInterfacesIpAddress(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeHaVpnGatewayVpnInterfacesInterconnectAttachment(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	f, err := tpgresource.ParseRegionalFieldValue("interconnectAttachments", v.(string), "project", "region", "zone", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for interconnect_attachment: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeHaVpnGatewayRegion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	f, err := tpgresource.ParseGlobalFieldValue("regions", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for region: %s", err)
	}
	return f.RelativeLink(), nil
}
