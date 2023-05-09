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

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgresource"
	transport_tpg "github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/transport"
)

const ApigeeEnvgroupAttachmentAssetType string = "apigee.googleapis.com/EnvgroupAttachment"

func resourceConverterApigeeEnvgroupAttachment() ResourceConverter {
	return ResourceConverter{
		AssetType: ApigeeEnvgroupAttachmentAssetType,
		Convert:   GetApigeeEnvgroupAttachmentCaiObject,
	}
}

func GetApigeeEnvgroupAttachmentCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]Asset, error) {
	name, err := assetName(d, config, "//apigee.googleapis.com/{{envgroup_id}}/attachments/{{name}}")
	if err != nil {
		return []Asset{}, err
	}
	if obj, err := GetApigeeEnvgroupAttachmentApiObject(d, config); err == nil {
		return []Asset{{
			Name: name,
			Type: ApigeeEnvgroupAttachmentAssetType,
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/apigee/v1/rest",
				DiscoveryName:        "EnvgroupAttachment",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []Asset{}, err
	}
}

func GetApigeeEnvgroupAttachmentApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	environmentProp, err := expandApigeeEnvgroupAttachmentEnvironment(d.Get("environment"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("environment"); !tpgresource.IsEmptyValue(reflect.ValueOf(environmentProp)) && (ok || !reflect.DeepEqual(v, environmentProp)) {
		obj["environment"] = environmentProp
	}

	return obj, nil
}

func expandApigeeEnvgroupAttachmentEnvironment(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
