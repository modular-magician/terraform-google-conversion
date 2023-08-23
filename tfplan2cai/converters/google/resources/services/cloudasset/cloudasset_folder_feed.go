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

package cloudasset

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const CloudAssetFolderFeedAssetType string = "cloudasset.googleapis.com/FolderFeed"

func ResourceConverterCloudAssetFolderFeed() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: CloudAssetFolderFeedAssetType,
		Convert:   GetCloudAssetFolderFeedCaiObject,
	}
}

func GetCloudAssetFolderFeedCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//cloudasset.googleapis.com/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetCloudAssetFolderFeedApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: CloudAssetFolderFeedAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/cloudasset/v1/rest",
				DiscoveryName:        "FolderFeed",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetCloudAssetFolderFeedApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	assetNamesProp, err := expandCloudAssetFolderFeedAssetNames(d.Get("asset_names"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("asset_names"); !tpgresource.IsEmptyValue(reflect.ValueOf(assetNamesProp)) && (ok || !reflect.DeepEqual(v, assetNamesProp)) {
		obj["assetNames"] = assetNamesProp
	}
	assetTypesProp, err := expandCloudAssetFolderFeedAssetTypes(d.Get("asset_types"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("asset_types"); !tpgresource.IsEmptyValue(reflect.ValueOf(assetTypesProp)) && (ok || !reflect.DeepEqual(v, assetTypesProp)) {
		obj["assetTypes"] = assetTypesProp
	}
	contentTypeProp, err := expandCloudAssetFolderFeedContentType(d.Get("content_type"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("content_type"); !tpgresource.IsEmptyValue(reflect.ValueOf(contentTypeProp)) && (ok || !reflect.DeepEqual(v, contentTypeProp)) {
		obj["contentType"] = contentTypeProp
	}
	feedOutputConfigProp, err := expandCloudAssetFolderFeedFeedOutputConfig(d.Get("feed_output_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("feed_output_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(feedOutputConfigProp)) && (ok || !reflect.DeepEqual(v, feedOutputConfigProp)) {
		obj["feedOutputConfig"] = feedOutputConfigProp
	}
	conditionProp, err := expandCloudAssetFolderFeedCondition(d.Get("condition"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("condition"); !tpgresource.IsEmptyValue(reflect.ValueOf(conditionProp)) && (ok || !reflect.DeepEqual(v, conditionProp)) {
		obj["condition"] = conditionProp
	}

	return resourceCloudAssetFolderFeedEncoder(d, config, obj)
}

func resourceCloudAssetFolderFeedEncoder(d tpgresource.TerraformResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	// Remove the "folders/" prefix from the folder ID
	if folder, ok := d.GetOkExists("folder"); ok {
		if err := d.Set("folder_id", strings.TrimPrefix(folder.(string), "folders/")); err != nil {
			return nil, fmt.Errorf("Error setting folder_id: %s", err)
		}
	}
	// The feed object must be under the "feed" attribute on the request.
	newObj := make(map[string]interface{})
	newObj["feed"] = obj
	return newObj, nil
}

func expandCloudAssetFolderFeedAssetNames(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudAssetFolderFeedAssetTypes(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudAssetFolderFeedContentType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudAssetFolderFeedFeedOutputConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedPubsubDestination, err := expandCloudAssetFolderFeedFeedOutputConfigPubsubDestination(original["pubsub_destination"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPubsubDestination); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["pubsubDestination"] = transformedPubsubDestination
	}

	return transformed, nil
}

func expandCloudAssetFolderFeedFeedOutputConfigPubsubDestination(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedTopic, err := expandCloudAssetFolderFeedFeedOutputConfigPubsubDestinationTopic(original["topic"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTopic); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["topic"] = transformedTopic
	}

	return transformed, nil
}

func expandCloudAssetFolderFeedFeedOutputConfigPubsubDestinationTopic(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudAssetFolderFeedCondition(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedExpression, err := expandCloudAssetFolderFeedConditionExpression(original["expression"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedExpression); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["expression"] = transformedExpression
	}

	transformedTitle, err := expandCloudAssetFolderFeedConditionTitle(original["title"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTitle); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["title"] = transformedTitle
	}

	transformedDescription, err := expandCloudAssetFolderFeedConditionDescription(original["description"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDescription); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["description"] = transformedDescription
	}

	transformedLocation, err := expandCloudAssetFolderFeedConditionLocation(original["location"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLocation); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["location"] = transformedLocation
	}

	return transformed, nil
}

func expandCloudAssetFolderFeedConditionExpression(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudAssetFolderFeedConditionTitle(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudAssetFolderFeedConditionDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudAssetFolderFeedConditionLocation(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
