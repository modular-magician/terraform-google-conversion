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

package vmwareengine

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const VmwareengineExternalAccessRuleAssetType string = "vmwareengine.googleapis.com/ExternalAccessRule"

func ResourceConverterVmwareengineExternalAccessRule() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: VmwareengineExternalAccessRuleAssetType,
		Convert:   GetVmwareengineExternalAccessRuleCaiObject,
	}
}

func GetVmwareengineExternalAccessRuleCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//vmwareengine.googleapis.com/{{parent}}/externalAccessRules/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetVmwareengineExternalAccessRuleApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: VmwareengineExternalAccessRuleAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/vmwareengine/v1/rest",
				DiscoveryName:        "ExternalAccessRule",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetVmwareengineExternalAccessRuleApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	descriptionProp, err := expandVmwareengineExternalAccessRuleDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	priorityProp, err := expandVmwareengineExternalAccessRulePriority(d.Get("priority"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("priority"); !tpgresource.IsEmptyValue(reflect.ValueOf(priorityProp)) && (ok || !reflect.DeepEqual(v, priorityProp)) {
		obj["priority"] = priorityProp
	}
	actionProp, err := expandVmwareengineExternalAccessRuleAction(d.Get("action"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("action"); !tpgresource.IsEmptyValue(reflect.ValueOf(actionProp)) && (ok || !reflect.DeepEqual(v, actionProp)) {
		obj["action"] = actionProp
	}
	ipProtocolProp, err := expandVmwareengineExternalAccessRuleIpProtocol(d.Get("ip_protocol"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("ip_protocol"); !tpgresource.IsEmptyValue(reflect.ValueOf(ipProtocolProp)) && (ok || !reflect.DeepEqual(v, ipProtocolProp)) {
		obj["ipProtocol"] = ipProtocolProp
	}
	sourceIpRangesProp, err := expandVmwareengineExternalAccessRuleSourceIpRanges(d.Get("source_ip_ranges"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("source_ip_ranges"); !tpgresource.IsEmptyValue(reflect.ValueOf(sourceIpRangesProp)) && (ok || !reflect.DeepEqual(v, sourceIpRangesProp)) {
		obj["sourceIpRanges"] = sourceIpRangesProp
	}
	sourcePortsProp, err := expandVmwareengineExternalAccessRuleSourcePorts(d.Get("source_ports"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("source_ports"); !tpgresource.IsEmptyValue(reflect.ValueOf(sourcePortsProp)) && (ok || !reflect.DeepEqual(v, sourcePortsProp)) {
		obj["sourcePorts"] = sourcePortsProp
	}
	destinationIpRangesProp, err := expandVmwareengineExternalAccessRuleDestinationIpRanges(d.Get("destination_ip_ranges"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("destination_ip_ranges"); !tpgresource.IsEmptyValue(reflect.ValueOf(destinationIpRangesProp)) && (ok || !reflect.DeepEqual(v, destinationIpRangesProp)) {
		obj["destinationIpRanges"] = destinationIpRangesProp
	}
	destinationPortsProp, err := expandVmwareengineExternalAccessRuleDestinationPorts(d.Get("destination_ports"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("destination_ports"); !tpgresource.IsEmptyValue(reflect.ValueOf(destinationPortsProp)) && (ok || !reflect.DeepEqual(v, destinationPortsProp)) {
		obj["destinationPorts"] = destinationPortsProp
	}

	return obj, nil
}

func expandVmwareengineExternalAccessRuleDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVmwareengineExternalAccessRulePriority(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVmwareengineExternalAccessRuleAction(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVmwareengineExternalAccessRuleIpProtocol(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVmwareengineExternalAccessRuleSourceIpRanges(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedIpAddress, err := expandVmwareengineExternalAccessRuleSourceIpRangesIpAddress(original["ip_address"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedIpAddress); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["ipAddress"] = transformedIpAddress
		}

		transformedIpAddressRange, err := expandVmwareengineExternalAccessRuleSourceIpRangesIpAddressRange(original["ip_address_range"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedIpAddressRange); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["ipAddressRange"] = transformedIpAddressRange
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandVmwareengineExternalAccessRuleSourceIpRangesIpAddress(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVmwareengineExternalAccessRuleSourceIpRangesIpAddressRange(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVmwareengineExternalAccessRuleSourcePorts(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVmwareengineExternalAccessRuleDestinationIpRanges(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedIpAddressRange, err := expandVmwareengineExternalAccessRuleDestinationIpRangesIpAddressRange(original["ip_address_range"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedIpAddressRange); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["ipAddressRange"] = transformedIpAddressRange
		}

		transformedExternalAddress, err := expandVmwareengineExternalAccessRuleDestinationIpRangesExternalAddress(original["external_address"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedExternalAddress); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["externalAddress"] = transformedExternalAddress
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandVmwareengineExternalAccessRuleDestinationIpRangesIpAddressRange(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVmwareengineExternalAccessRuleDestinationIpRangesExternalAddress(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVmwareengineExternalAccessRuleDestinationPorts(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}