// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//     ***     DIFF TEST DIFF TEST    ***
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

func GetAccessContextManagerAccessPolicyCaiObject(d TerraformResourceData, config *Config) (Asset, error) {
	name, err := assetName(d, config, "//accesscontextmanager.googleapis.com/accessPolicies/{{name}}")
	if err != nil {
		return Asset{}, err
	}
	if obj, err := GetAccessContextManagerAccessPolicyApiObject(d, config); err == nil {
		return Asset{
			Name: name,
			Type: "accesscontextmanager.googleapis.com/AccessPolicy",
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/accesscontextmanager/v1/rest",
				DiscoveryName:        "AccessPolicy",
				Data:                 obj,
			},
		}, nil
	} else {
		return Asset{}, err
	}
}

func GetAccessContextManagerAccessPolicyApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	parentProp, err := expandAccessContextManagerAccessPolicyParent(d.Get("parent"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("parent"); !isEmptyValue(reflect.ValueOf(parentProp)) && (ok || !reflect.DeepEqual(v, parentProp)) {
		obj["parent"] = parentProp
	}
	titleProp, err := expandAccessContextManagerAccessPolicyTitle(d.Get("title"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("title"); !isEmptyValue(reflect.ValueOf(titleProp)) && (ok || !reflect.DeepEqual(v, titleProp)) {
		obj["title"] = titleProp
	}

	return obj, nil
}

func expandAccessContextManagerAccessPolicyParent(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerAccessPolicyTitle(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
