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

package networkmanagement

import (
	"fmt"
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const NetworkManagementConnectivityTestAssetType string = "networkmanagement.googleapis.com/ConnectivityTest"

func ResourceConverterNetworkManagementConnectivityTest() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: NetworkManagementConnectivityTestAssetType,
		Convert:   GetNetworkManagementConnectivityTestCaiObject,
	}
}

func GetNetworkManagementConnectivityTestCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//networkmanagement.googleapis.com/projects/{{project}}/locations/global/connectivityTests/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetNetworkManagementConnectivityTestApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: NetworkManagementConnectivityTestAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/networkmanagement/v1/rest",
				DiscoveryName:        "ConnectivityTest",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetNetworkManagementConnectivityTestApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandNetworkManagementConnectivityTestName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	descriptionProp, err := expandNetworkManagementConnectivityTestDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	sourceProp, err := expandNetworkManagementConnectivityTestSource(d.Get("source"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("source"); !tpgresource.IsEmptyValue(reflect.ValueOf(sourceProp)) && (ok || !reflect.DeepEqual(v, sourceProp)) {
		obj["source"] = sourceProp
	}
	destinationProp, err := expandNetworkManagementConnectivityTestDestination(d.Get("destination"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("destination"); !tpgresource.IsEmptyValue(reflect.ValueOf(destinationProp)) && (ok || !reflect.DeepEqual(v, destinationProp)) {
		obj["destination"] = destinationProp
	}
	protocolProp, err := expandNetworkManagementConnectivityTestProtocol(d.Get("protocol"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("protocol"); !tpgresource.IsEmptyValue(reflect.ValueOf(protocolProp)) && (ok || !reflect.DeepEqual(v, protocolProp)) {
		obj["protocol"] = protocolProp
	}
	relatedProjectsProp, err := expandNetworkManagementConnectivityTestRelatedProjects(d.Get("related_projects"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("related_projects"); !tpgresource.IsEmptyValue(reflect.ValueOf(relatedProjectsProp)) && (ok || !reflect.DeepEqual(v, relatedProjectsProp)) {
		obj["relatedProjects"] = relatedProjectsProp
	}
	labelsProp, err := expandNetworkManagementConnectivityTestEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	return obj, nil
}

func expandNetworkManagementConnectivityTestName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	// projects/X/tests/Y - note not "connectivityTests"
	f, err := tpgresource.ParseGlobalFieldValue("tests", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for zone: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandNetworkManagementConnectivityTestDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkManagementConnectivityTestSource(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedIpAddress, err := expandNetworkManagementConnectivityTestSourceIpAddress(original["ip_address"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedIpAddress); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["ipAddress"] = transformedIpAddress
	}

	transformedPort, err := expandNetworkManagementConnectivityTestSourcePort(original["port"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPort); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["port"] = transformedPort
	}

	transformedInstance, err := expandNetworkManagementConnectivityTestSourceInstance(original["instance"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedInstance); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["instance"] = transformedInstance
	}

	transformedNetwork, err := expandNetworkManagementConnectivityTestSourceNetwork(original["network"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNetwork); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["network"] = transformedNetwork
	}

	transformedNetworkType, err := expandNetworkManagementConnectivityTestSourceNetworkType(original["network_type"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNetworkType); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["networkType"] = transformedNetworkType
	}

	transformedProjectId, err := expandNetworkManagementConnectivityTestSourceProjectId(original["project_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedProjectId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["projectId"] = transformedProjectId
	}

	return transformed, nil
}

func expandNetworkManagementConnectivityTestSourceIpAddress(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkManagementConnectivityTestSourcePort(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkManagementConnectivityTestSourceInstance(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkManagementConnectivityTestSourceNetwork(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkManagementConnectivityTestSourceNetworkType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkManagementConnectivityTestSourceProjectId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkManagementConnectivityTestDestination(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedIpAddress, err := expandNetworkManagementConnectivityTestDestinationIpAddress(original["ip_address"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedIpAddress); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["ipAddress"] = transformedIpAddress
	}

	transformedPort, err := expandNetworkManagementConnectivityTestDestinationPort(original["port"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPort); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["port"] = transformedPort
	}

	transformedInstance, err := expandNetworkManagementConnectivityTestDestinationInstance(original["instance"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedInstance); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["instance"] = transformedInstance
	}

	transformedNetwork, err := expandNetworkManagementConnectivityTestDestinationNetwork(original["network"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNetwork); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["network"] = transformedNetwork
	}

	transformedProjectId, err := expandNetworkManagementConnectivityTestDestinationProjectId(original["project_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedProjectId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["projectId"] = transformedProjectId
	}

	return transformed, nil
}

func expandNetworkManagementConnectivityTestDestinationIpAddress(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkManagementConnectivityTestDestinationPort(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkManagementConnectivityTestDestinationInstance(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkManagementConnectivityTestDestinationNetwork(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkManagementConnectivityTestDestinationProjectId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkManagementConnectivityTestProtocol(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkManagementConnectivityTestRelatedProjects(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkManagementConnectivityTestEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
