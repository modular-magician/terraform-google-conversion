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

package networkservices

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgresource"
	transport_tpg "github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/transport"
)

const NetworkServicesMeshAssetType string = "networkservices.googleapis.com/Mesh"

func ResourceConverterNetworkServicesMesh() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType: NetworkServicesMeshAssetType,
		Convert:   GetNetworkServicesMeshCaiObject,
	}
}

func GetNetworkServicesMeshCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	name, err := tpgresource.AssetName(d, config, "//networkservices.googleapis.com/projects/{{project}}/locations/global/meshes/{{name}}")
	if err != nil {
		return []tpgresource.Asset{}, err
	}
	if obj, err := GetNetworkServicesMeshApiObject(d, config); err == nil {
		return []tpgresource.Asset{{
			Name: name,
			Type: NetworkServicesMeshAssetType,
			Resource: &tpgresource.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/networkservices/v1/rest",
				DiscoveryName:        "Mesh",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []tpgresource.Asset{}, err
	}
}

func GetNetworkServicesMeshApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	labelsProp, err := expandNetworkServicesMeshLabels(d.Get("labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	descriptionProp, err := expandNetworkServicesMeshDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	interceptionPortProp, err := expandNetworkServicesMeshInterceptionPort(d.Get("interception_port"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("interception_port"); !tpgresource.IsEmptyValue(reflect.ValueOf(interceptionPortProp)) && (ok || !reflect.DeepEqual(v, interceptionPortProp)) {
		obj["interceptionPort"] = interceptionPortProp
	}

	return obj, nil
}

func expandNetworkServicesMeshLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandNetworkServicesMeshDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesMeshInterceptionPort(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
