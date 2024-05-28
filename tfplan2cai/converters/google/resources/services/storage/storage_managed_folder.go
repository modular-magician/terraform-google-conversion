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

package storage

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const StorageManagedFolderAssetType string = "storage.googleapis.com/ManagedFolder"

func ResourceConverterStorageManagedFolder() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: StorageManagedFolderAssetType,
		Convert:   GetStorageManagedFolderCaiObject,
	}
}

func GetStorageManagedFolderCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//storage.googleapis.com/b/{{bucket}}/managedFolders/{{%name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetStorageManagedFolderApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: StorageManagedFolderAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/storage/v1/rest",
				DiscoveryName:        "ManagedFolder",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetStorageManagedFolderApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandStorageManagedFolderName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}

	return obj, nil
}

func expandStorageManagedFolderName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
