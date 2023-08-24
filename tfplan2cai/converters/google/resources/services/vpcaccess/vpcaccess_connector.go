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

package vpcaccess

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const VPCAccessConnectorAssetType string = "vpcaccess.googleapis.com/Connector"

func ResourceConverterVPCAccessConnector() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: VPCAccessConnectorAssetType,
		Convert:   GetVPCAccessConnectorCaiObject,
	}
}

func GetVPCAccessConnectorCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//vpcaccess.googleapis.com/projects/{{project}}/locations/{{region}}/connectors/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetVPCAccessConnectorApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: VPCAccessConnectorAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1beta1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/vpcaccess/v1beta1/rest",
				DiscoveryName:        "Connector",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetVPCAccessConnectorApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandVPCAccessConnectorName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	networkProp, err := expandVPCAccessConnectorNetwork(d.Get("network"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("network"); !tpgresource.IsEmptyValue(reflect.ValueOf(networkProp)) && (ok || !reflect.DeepEqual(v, networkProp)) {
		obj["network"] = networkProp
	}
	ipCidrRangeProp, err := expandVPCAccessConnectorIpCidrRange(d.Get("ip_cidr_range"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("ip_cidr_range"); !tpgresource.IsEmptyValue(reflect.ValueOf(ipCidrRangeProp)) && (ok || !reflect.DeepEqual(v, ipCidrRangeProp)) {
		obj["ipCidrRange"] = ipCidrRangeProp
	}
	machineTypeProp, err := expandVPCAccessConnectorMachineType(d.Get("machine_type"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("machine_type"); !tpgresource.IsEmptyValue(reflect.ValueOf(machineTypeProp)) && (ok || !reflect.DeepEqual(v, machineTypeProp)) {
		obj["machineType"] = machineTypeProp
	}
	minThroughputProp, err := expandVPCAccessConnectorMinThroughput(d.Get("min_throughput"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("min_throughput"); !tpgresource.IsEmptyValue(reflect.ValueOf(minThroughputProp)) && (ok || !reflect.DeepEqual(v, minThroughputProp)) {
		obj["minThroughput"] = minThroughputProp
	}
	minInstancesProp, err := expandVPCAccessConnectorMinInstances(d.Get("min_instances"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("min_instances"); !tpgresource.IsEmptyValue(reflect.ValueOf(minInstancesProp)) && (ok || !reflect.DeepEqual(v, minInstancesProp)) {
		obj["minInstances"] = minInstancesProp
	}
	maxInstancesProp, err := expandVPCAccessConnectorMaxInstances(d.Get("max_instances"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("max_instances"); !tpgresource.IsEmptyValue(reflect.ValueOf(maxInstancesProp)) && (ok || !reflect.DeepEqual(v, maxInstancesProp)) {
		obj["maxInstances"] = maxInstancesProp
	}
	maxThroughputProp, err := expandVPCAccessConnectorMaxThroughput(d.Get("max_throughput"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("max_throughput"); !tpgresource.IsEmptyValue(reflect.ValueOf(maxThroughputProp)) && (ok || !reflect.DeepEqual(v, maxThroughputProp)) {
		obj["maxThroughput"] = maxThroughputProp
	}
	subnetProp, err := expandVPCAccessConnectorSubnet(d.Get("subnet"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("subnet"); !tpgresource.IsEmptyValue(reflect.ValueOf(subnetProp)) && (ok || !reflect.DeepEqual(v, subnetProp)) {
		obj["subnet"] = subnetProp
	}

	return resourceVPCAccessConnectorEncoder(d, config, obj)
}

func resourceVPCAccessConnectorEncoder(d tpgresource.TerraformResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	delete(obj, "name")
	return obj, nil
}

func expandVPCAccessConnectorName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVPCAccessConnectorNetwork(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return tpgresource.GetResourceNameFromSelfLink(v.(string)), nil
}

func expandVPCAccessConnectorIpCidrRange(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVPCAccessConnectorMachineType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVPCAccessConnectorMinThroughput(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVPCAccessConnectorMinInstances(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVPCAccessConnectorMaxInstances(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVPCAccessConnectorMaxThroughput(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVPCAccessConnectorSubnet(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedName, err := expandVPCAccessConnectorSubnetName(original["name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["name"] = transformedName
	}

	transformedProjectId, err := expandVPCAccessConnectorSubnetProjectId(original["project_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedProjectId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["projectId"] = transformedProjectId
	}

	return transformed, nil
}

func expandVPCAccessConnectorSubnetName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVPCAccessConnectorSubnetProjectId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
