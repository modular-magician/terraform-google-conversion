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

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const VmwareengineNetworkPolicyAssetType string = "vmwareengine.googleapis.com/NetworkPolicy"

func ResourceConverterVmwareengineNetworkPolicy() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: VmwareengineNetworkPolicyAssetType,
		Convert:   GetVmwareengineNetworkPolicyCaiObject,
	}
}

func GetVmwareengineNetworkPolicyCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//vmwareengine.googleapis.com/projects/{{project}}/locations/{{location}}/networkPolicies/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetVmwareengineNetworkPolicyApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: VmwareengineNetworkPolicyAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/vmwareengine/v1/rest",
				DiscoveryName:        "NetworkPolicy",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetVmwareengineNetworkPolicyApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	edgeServicesCidrProp, err := expandVmwareengineNetworkPolicyEdgeServicesCidr(d.Get("edge_services_cidr"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("edge_services_cidr"); !tpgresource.IsEmptyValue(reflect.ValueOf(edgeServicesCidrProp)) && (ok || !reflect.DeepEqual(v, edgeServicesCidrProp)) {
		obj["edgeServicesCidr"] = edgeServicesCidrProp
	}
	descriptionProp, err := expandVmwareengineNetworkPolicyDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	vmwareEngineNetworkProp, err := expandVmwareengineNetworkPolicyVmwareEngineNetwork(d.Get("vmware_engine_network"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("vmware_engine_network"); !tpgresource.IsEmptyValue(reflect.ValueOf(vmwareEngineNetworkProp)) && (ok || !reflect.DeepEqual(v, vmwareEngineNetworkProp)) {
		obj["vmwareEngineNetwork"] = vmwareEngineNetworkProp
	}
	internetAccessProp, err := expandVmwareengineNetworkPolicyInternetAccess(d.Get("internet_access"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("internet_access"); !tpgresource.IsEmptyValue(reflect.ValueOf(internetAccessProp)) && (ok || !reflect.DeepEqual(v, internetAccessProp)) {
		obj["internetAccess"] = internetAccessProp
	}
	externalIpProp, err := expandVmwareengineNetworkPolicyExternalIp(d.Get("external_ip"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("external_ip"); !tpgresource.IsEmptyValue(reflect.ValueOf(externalIpProp)) && (ok || !reflect.DeepEqual(v, externalIpProp)) {
		obj["externalIp"] = externalIpProp
	}

	return obj, nil
}

func expandVmwareengineNetworkPolicyEdgeServicesCidr(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVmwareengineNetworkPolicyDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVmwareengineNetworkPolicyVmwareEngineNetwork(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVmwareengineNetworkPolicyInternetAccess(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedEnabled, err := expandVmwareengineNetworkPolicyInternetAccessEnabled(original["enabled"], d, config)
	if err != nil {
		return nil, err
	} else {
		transformed["enabled"] = transformedEnabled
	}

	transformedState, err := expandVmwareengineNetworkPolicyInternetAccessState(original["state"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedState); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["state"] = transformedState
	}

	return transformed, nil
}

func expandVmwareengineNetworkPolicyInternetAccessEnabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVmwareengineNetworkPolicyInternetAccessState(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVmwareengineNetworkPolicyExternalIp(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedEnabled, err := expandVmwareengineNetworkPolicyExternalIpEnabled(original["enabled"], d, config)
	if err != nil {
		return nil, err
	} else {
		transformed["enabled"] = transformedEnabled
	}

	transformedState, err := expandVmwareengineNetworkPolicyExternalIpState(original["state"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedState); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["state"] = transformedState
	}

	return transformed, nil
}

func expandVmwareengineNetworkPolicyExternalIpEnabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVmwareengineNetworkPolicyExternalIpState(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
