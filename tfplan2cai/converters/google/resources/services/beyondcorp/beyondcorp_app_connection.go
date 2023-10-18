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

package beyondcorp

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const BeyondcorpAppConnectionAssetType string = "beyondcorp.googleapis.com/AppConnection"

func ResourceConverterBeyondcorpAppConnection() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: BeyondcorpAppConnectionAssetType,
		Convert:   GetBeyondcorpAppConnectionCaiObject,
	}
}

func GetBeyondcorpAppConnectionCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//beyondcorp.googleapis.com/projects/{{project}}/locations/{{region}}/appConnections/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetBeyondcorpAppConnectionApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: BeyondcorpAppConnectionAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/beyondcorp/v1/rest",
				DiscoveryName:        "AppConnection",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetBeyondcorpAppConnectionApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	displayNameProp, err := expandBeyondcorpAppConnectionDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	typeProp, err := expandBeyondcorpAppConnectionType(d.Get("type"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("type"); !tpgresource.IsEmptyValue(reflect.ValueOf(typeProp)) && (ok || !reflect.DeepEqual(v, typeProp)) {
		obj["type"] = typeProp
	}
	applicationEndpointProp, err := expandBeyondcorpAppConnectionApplicationEndpoint(d.Get("application_endpoint"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("application_endpoint"); !tpgresource.IsEmptyValue(reflect.ValueOf(applicationEndpointProp)) && (ok || !reflect.DeepEqual(v, applicationEndpointProp)) {
		obj["applicationEndpoint"] = applicationEndpointProp
	}
	connectorsProp, err := expandBeyondcorpAppConnectionConnectors(d.Get("connectors"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("connectors"); !tpgresource.IsEmptyValue(reflect.ValueOf(connectorsProp)) && (ok || !reflect.DeepEqual(v, connectorsProp)) {
		obj["connectors"] = connectorsProp
	}
	gatewayProp, err := expandBeyondcorpAppConnectionGateway(d.Get("gateway"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("gateway"); !tpgresource.IsEmptyValue(reflect.ValueOf(gatewayProp)) && (ok || !reflect.DeepEqual(v, gatewayProp)) {
		obj["gateway"] = gatewayProp
	}
	labelsProp, err := expandBeyondcorpAppConnectionEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	return obj, nil
}

func expandBeyondcorpAppConnectionDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBeyondcorpAppConnectionType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBeyondcorpAppConnectionApplicationEndpoint(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedHost, err := expandBeyondcorpAppConnectionApplicationEndpointHost(original["host"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedHost); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["host"] = transformedHost
	}

	transformedPort, err := expandBeyondcorpAppConnectionApplicationEndpointPort(original["port"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPort); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["port"] = transformedPort
	}

	return transformed, nil
}

func expandBeyondcorpAppConnectionApplicationEndpointHost(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBeyondcorpAppConnectionApplicationEndpointPort(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBeyondcorpAppConnectionConnectors(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBeyondcorpAppConnectionGateway(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedAppGateway, err := expandBeyondcorpAppConnectionGatewayAppGateway(original["app_gateway"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAppGateway); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["appGateway"] = transformedAppGateway
	}

	transformedType, err := expandBeyondcorpAppConnectionGatewayType(original["type"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedType); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["type"] = transformedType
	}

	return transformed, nil
}

func expandBeyondcorpAppConnectionGatewayAppGateway(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBeyondcorpAppConnectionGatewayType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBeyondcorpAppConnectionEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
