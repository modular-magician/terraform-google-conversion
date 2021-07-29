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

import "reflect"

func GetAppEngineServiceNetworkSettingsCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	name, err := assetName(d, config, "//appengine.googleapis.com/apps/{{project}}/services/{{service}}")
	if err != nil {
		return []Asset{}, err
	}
	if obj, err := GetAppEngineServiceNetworkSettingsApiObject(d, config); err == nil {
		return []Asset{{
			Name: name,
			Type: "appengine.googleapis.com/ServiceNetworkSettings",
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/appengine/v1/rest",
				DiscoveryName:        "ServiceNetworkSettings",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []Asset{}, err
	}
}

func GetAppEngineServiceNetworkSettingsApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	idProp, err := expandAppEngineServiceNetworkSettingsService(d.Get("service"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("service"); !isEmptyValue(reflect.ValueOf(idProp)) && (ok || !reflect.DeepEqual(v, idProp)) {
		obj["id"] = idProp
	}
	networkSettingsProp, err := expandAppEngineServiceNetworkSettingsNetworkSettings(d.Get("network_settings"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("network_settings"); !isEmptyValue(reflect.ValueOf(networkSettingsProp)) && (ok || !reflect.DeepEqual(v, networkSettingsProp)) {
		obj["networkSettings"] = networkSettingsProp
	}

	return obj, nil
}

func expandAppEngineServiceNetworkSettingsService(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineServiceNetworkSettingsNetworkSettings(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedIngressTrafficAllowed, err := expandAppEngineServiceNetworkSettingsNetworkSettingsIngressTrafficAllowed(original["ingress_traffic_allowed"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedIngressTrafficAllowed); val.IsValid() && !isEmptyValue(val) {
		transformed["ingressTrafficAllowed"] = transformedIngressTrafficAllowed
	}

	return transformed, nil
}

func expandAppEngineServiceNetworkSettingsNetworkSettingsIngressTrafficAllowed(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
