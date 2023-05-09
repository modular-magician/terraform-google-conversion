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

const TagsTagBindingAssetType string = "cloudresourcemanager.googleapis.com/TagBinding"

func resourceConverterTagsTagBinding() ResourceConverter {
	return ResourceConverter{
		AssetType: TagsTagBindingAssetType,
		Convert:   GetTagsTagBindingCaiObject,
	}
}

func GetTagsTagBindingCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]Asset, error) {
	name, err := assetName(d, config, "//cloudresourcemanager.googleapis.com/tagBindings/?parent={{parent}}&pageSize=300")
	if err != nil {
		return []Asset{}, err
	}
	if obj, err := GetTagsTagBindingApiObject(d, config); err == nil {
		return []Asset{{
			Name: name,
			Type: TagsTagBindingAssetType,
			Resource: &AssetResource{
				Version:              "v3",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/cloudresourcemanager/v3/rest",
				DiscoveryName:        "TagBinding",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []Asset{}, err
	}
}

func GetTagsTagBindingApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	parentProp, err := expandTagsTagBindingParent(d.Get("parent"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("parent"); !tpgresource.IsEmptyValue(reflect.ValueOf(parentProp)) && (ok || !reflect.DeepEqual(v, parentProp)) {
		obj["parent"] = parentProp
	}
	tagValueProp, err := expandTagsTagBindingTagValue(d.Get("tag_value"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("tag_value"); !tpgresource.IsEmptyValue(reflect.ValueOf(tagValueProp)) && (ok || !reflect.DeepEqual(v, tagValueProp)) {
		obj["tagValue"] = tagValueProp
	}

	return obj, nil
}

func expandTagsTagBindingParent(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandTagsTagBindingTagValue(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
