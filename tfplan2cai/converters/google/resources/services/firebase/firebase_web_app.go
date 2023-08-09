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

package firebase

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const FirebaseWebAppAssetType string = "firebase.googleapis.com/WebApp"

func ResourceConverterFirebaseWebApp() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType: FirebaseWebAppAssetType,
		Convert:   GetFirebaseWebAppCaiObject,
	}
}

func GetFirebaseWebAppCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	name, err := tpgresource.AssetName(d, config, "//firebase.googleapis.com/projects/{{project}}/webApps/{{app_id}}")
	if err != nil {
		return []tpgresource.Asset{}, err
	}
	if obj, err := GetFirebaseWebAppApiObject(d, config); err == nil {
		return []tpgresource.Asset{{
			Name: name,
			Type: FirebaseWebAppAssetType,
			Resource: &tpgresource.AssetResource{
				Version:              "v1beta1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/firebase/v1beta1/rest",
				DiscoveryName:        "WebApp",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []tpgresource.Asset{}, err
	}
}

func GetFirebaseWebAppApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	displayNameProp, err := expandFirebaseWebAppDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}

	return obj, nil
}

func expandFirebaseWebAppDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
