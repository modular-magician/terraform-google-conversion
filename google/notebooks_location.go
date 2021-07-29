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

func GetNotebooksLocationCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	name, err := assetName(d, config, "//notebooks.googleapis.com/projects/{{project}}/locations/{{name}}")
	if err != nil {
		return []Asset{}, err
	}
	if obj, err := GetNotebooksLocationApiObject(d, config); err == nil {
		return []Asset{{
			Name: name,
			Type: "notebooks.googleapis.com/Location",
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/notebooks/v1/rest",
				DiscoveryName:        "Location",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []Asset{}, err
	}
}

func GetNotebooksLocationApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandNotebooksLocationName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}

	return obj, nil
}

func expandNotebooksLocationName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
