// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Nathan is editing this to generate diffs in lots of files.
//     He won't submit this change.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import "reflect"

func GetKMSKeyRingCaiObject(d TerraformResourceData, config *Config) (Asset, error) {
	name, err := assetName(d, config, "//kms.googleapis.com/projects/{{project}}/locations/{{location}}/keyRings/{{name}}")
	if err != nil {
		return Asset{}, err
	}
	if obj, err := GetKMSKeyRingApiObject(d, config); err == nil {
		return Asset{
			Name: name,
			Type: "kms.googleapis.com/KeyRing",
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/kms/v1/rest",
				DiscoveryName:        "KeyRing",
				Data:                 obj,
			},
		}, nil
	} else {
		return Asset{}, err
	}
}

func GetKMSKeyRingApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandKMSKeyRingName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	locationProp, err := expandKMSKeyRingLocation(d.Get("location"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("location"); !isEmptyValue(reflect.ValueOf(locationProp)) && (ok || !reflect.DeepEqual(v, locationProp)) {
		obj["location"] = locationProp
	}

	return resourceKMSKeyRingEncoder(d, config, obj)
}

func resourceKMSKeyRingEncoder(d TerraformResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	return nil, nil
}

func expandKMSKeyRingName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandKMSKeyRingLocation(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
