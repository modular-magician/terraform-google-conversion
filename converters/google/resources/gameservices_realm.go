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

const GameServicesRealmAssetType string = "gameservices.googleapis.com/Realm"

func resourceConverterGameServicesRealm() ResourceConverter {
	return ResourceConverter{
		AssetType: GameServicesRealmAssetType,
		Convert:   GetGameServicesRealmCaiObject,
	}
}

func GetGameServicesRealmCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	name, err := assetName(d, config, "//gameservices.googleapis.com/projects/{{project}}/locations/{{location}}/realms/{{realm_id}}")
	if err != nil {
		return []Asset{}, err
	}
	if obj, err := GetGameServicesRealmApiObject(d, config); err == nil {
		return []Asset{{
			Name: name,
			Type: GameServicesRealmAssetType,
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/gameservices/v1/rest",
				DiscoveryName:        "Realm",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []Asset{}, err
	}
}

func GetGameServicesRealmApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	labelsProp, err := expandGameServicesRealmLabels(d.Get("labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	timeZoneProp, err := expandGameServicesRealmTimeZone(d.Get("time_zone"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("time_zone"); !isEmptyValue(reflect.ValueOf(timeZoneProp)) && (ok || !reflect.DeepEqual(v, timeZoneProp)) {
		obj["timeZone"] = timeZoneProp
	}
	descriptionProp, err := expandGameServicesRealmDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}

	return obj, nil
}

func expandGameServicesRealmLabels(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandGameServicesRealmTimeZone(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandGameServicesRealmDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
