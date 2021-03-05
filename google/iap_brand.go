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

func GetIapBrandCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	name, err := assetName(d, config, "//iap.googleapis.com/{{name}}")
	if err != nil {
		return []Asset{}, err
	}
	if obj, err := GetIapBrandApiObject(d, config); err == nil {
		return []Asset{{
			Name: name,
			Type: "iap.googleapis.com/Brand",
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/iap/v1/rest",
				DiscoveryName:        "Brand",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []Asset{}, err
	}
}

func GetIapBrandApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	supportEmailProp, err := expandIapBrandSupportEmail(d.Get("support_email"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("support_email"); !isEmptyValue(reflect.ValueOf(supportEmailProp)) && (ok || !reflect.DeepEqual(v, supportEmailProp)) {
		obj["supportEmail"] = supportEmailProp
	}
	applicationTitleProp, err := expandIapBrandApplicationTitle(d.Get("application_title"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("application_title"); !isEmptyValue(reflect.ValueOf(applicationTitleProp)) && (ok || !reflect.DeepEqual(v, applicationTitleProp)) {
		obj["applicationTitle"] = applicationTitleProp
	}

	return obj, nil
}

func expandIapBrandSupportEmail(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandIapBrandApplicationTitle(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
