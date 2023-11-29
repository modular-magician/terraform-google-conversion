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

package kms

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const KMSKeyRingImportJobAssetType string = "cloudkms.googleapis.com/KeyRingImportJob"

func ResourceConverterKMSKeyRingImportJob() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: KMSKeyRingImportJobAssetType,
		Convert:   GetKMSKeyRingImportJobCaiObject,
	}
}

func GetKMSKeyRingImportJobCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//cloudkms.googleapis.com/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetKMSKeyRingImportJobApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: KMSKeyRingImportJobAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/cloudkms/v1/rest",
				DiscoveryName:        "KeyRingImportJob",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetKMSKeyRingImportJobApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	importMethodProp, err := expandKMSKeyRingImportJobImportMethod(d.Get("import_method"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("import_method"); !tpgresource.IsEmptyValue(reflect.ValueOf(importMethodProp)) && (ok || !reflect.DeepEqual(v, importMethodProp)) {
		obj["importMethod"] = importMethodProp
	}
	protectionLevelProp, err := expandKMSKeyRingImportJobProtectionLevel(d.Get("protection_level"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("protection_level"); !tpgresource.IsEmptyValue(reflect.ValueOf(protectionLevelProp)) && (ok || !reflect.DeepEqual(v, protectionLevelProp)) {
		obj["protectionLevel"] = protectionLevelProp
	}

	return obj, nil
}

func expandKMSKeyRingImportJobImportMethod(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandKMSKeyRingImportJobProtectionLevel(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
