// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
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

func GetResourceManagerLienCaiObject(d TerraformResourceData, config *Config) (Asset, error) {
	name, err := replaceVars(d, config, "//resourcemanager.googleapis.com/liens?parent={{parent}}")
	if err != nil {
		return Asset{}, err
	}
	if obj, err := GetResourceManagerLienApiObject(d, config); err == nil {
		return Asset{
			Name: name,
			Type: "google.resourcemanager.Lien",
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/resourcemanager/v1/rest",
				DiscoveryName:        "Lien",
				Data:                 obj,
			},
		}, nil
	} else {
		return Asset{}, err
	}
}

func GetResourceManagerLienApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	reasonProp, err := expandResourceManagerLienReason(d.Get("reason"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("reason"); !isEmptyValue(reflect.ValueOf(reasonProp)) && (ok || !reflect.DeepEqual(v, reasonProp)) {
		obj["reason"] = reasonProp
	}
	originProp, err := expandResourceManagerLienOrigin(d.Get("origin"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("origin"); !isEmptyValue(reflect.ValueOf(originProp)) && (ok || !reflect.DeepEqual(v, originProp)) {
		obj["origin"] = originProp
	}
	parentProp, err := expandResourceManagerLienParent(d.Get("parent"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("parent"); !isEmptyValue(reflect.ValueOf(parentProp)) && (ok || !reflect.DeepEqual(v, parentProp)) {
		obj["parent"] = parentProp
	}
	restrictionsProp, err := expandResourceManagerLienRestrictions(d.Get("restrictions"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("restrictions"); !isEmptyValue(reflect.ValueOf(restrictionsProp)) && (ok || !reflect.DeepEqual(v, restrictionsProp)) {
		obj["restrictions"] = restrictionsProp
	}

	return obj, nil
}

func expandResourceManagerLienReason(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandResourceManagerLienOrigin(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandResourceManagerLienParent(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandResourceManagerLienRestrictions(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
