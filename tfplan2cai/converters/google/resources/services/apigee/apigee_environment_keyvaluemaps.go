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

package apigee

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const ApigeeEnvironmentKeyvaluemapsAssetType string = "apigee.googleapis.com/EnvironmentKeyvaluemaps"

func ResourceConverterApigeeEnvironmentKeyvaluemaps() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: ApigeeEnvironmentKeyvaluemapsAssetType,
		Convert:   GetApigeeEnvironmentKeyvaluemapsCaiObject,
	}
}

func GetApigeeEnvironmentKeyvaluemapsCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//apigee.googleapis.com/{{env_id}}/keyvaluemaps/{{name}}/entries")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetApigeeEnvironmentKeyvaluemapsApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: ApigeeEnvironmentKeyvaluemapsAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/apigee/v1/rest",
				DiscoveryName:        "EnvironmentKeyvaluemaps",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetApigeeEnvironmentKeyvaluemapsApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandApigeeEnvironmentKeyvaluemapsName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	encryptedProp, err := expandApigeeEnvironmentKeyvaluemapsEncrypted(d.Get("encrypted"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("encrypted"); !tpgresource.IsEmptyValue(reflect.ValueOf(encryptedProp)) && (ok || !reflect.DeepEqual(v, encryptedProp)) {
		obj["encrypted"] = encryptedProp
	}

	return obj, nil
}

func expandApigeeEnvironmentKeyvaluemapsName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeEnvironmentKeyvaluemapsEncrypted(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
