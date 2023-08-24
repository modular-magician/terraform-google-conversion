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

package apigateway

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const ApiGatewayApiAssetType string = "apigateway.googleapis.com/Api"

func ResourceConverterApiGatewayApi() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: ApiGatewayApiAssetType,
		Convert:   GetApiGatewayApiCaiObject,
	}
}

func GetApiGatewayApiCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//apigateway.googleapis.com/projects/{{project}}/locations/global/apis/{{api_id}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetApiGatewayApiApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: ApiGatewayApiAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1beta",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/apigateway/v1beta/rest",
				DiscoveryName:        "Api",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetApiGatewayApiApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	displayNameProp, err := expandApiGatewayApiDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	managedServiceProp, err := expandApiGatewayApiManagedService(d.Get("managed_service"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("managed_service"); !tpgresource.IsEmptyValue(reflect.ValueOf(managedServiceProp)) && (ok || !reflect.DeepEqual(v, managedServiceProp)) {
		obj["managedService"] = managedServiceProp
	}
	labelsProp, err := expandApiGatewayApiLabels(d.Get("labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	return obj, nil
}

func expandApiGatewayApiDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApiGatewayApiManagedService(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApiGatewayApiLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
