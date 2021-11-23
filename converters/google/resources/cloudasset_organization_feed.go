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
	"fmt"
	"reflect"
	"strings"
)

const CloudAssetOrganizationFeedAssetType string = "cloudasset.googleapis.com/OrganizationFeed"

func resourceConverterCloudAssetOrganizationFeed() ResourceConverter {
	return ResourceConverter{
		AssetType: CloudAssetOrganizationFeedAssetType,
		Convert:   GetCloudAssetOrganizationFeedCaiObject,
	}
}

func GetCloudAssetOrganizationFeedCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	name, err := assetName(d, config, "//cloudasset.googleapis.com/{{name}}")
	if err != nil {
		return []Asset{}, err
	}
	if obj, err := GetCloudAssetOrganizationFeedApiObject(d, config); err == nil {
		return []Asset{{
			Name: name,
			Type: CloudAssetOrganizationFeedAssetType,
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/cloudasset/v1/rest",
				DiscoveryName:        "OrganizationFeed",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []Asset{}, err
	}
}

func GetCloudAssetOrganizationFeedApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	assetNamesProp, err := expandCloudAssetOrganizationFeedAssetNames(d.Get("asset_names"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("asset_names"); !isEmptyValue(reflect.ValueOf(assetNamesProp)) && (ok || !reflect.DeepEqual(v, assetNamesProp)) {
		obj["assetNames"] = assetNamesProp
	}
	assetTypesProp, err := expandCloudAssetOrganizationFeedAssetTypes(d.Get("asset_types"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("asset_types"); !isEmptyValue(reflect.ValueOf(assetTypesProp)) && (ok || !reflect.DeepEqual(v, assetTypesProp)) {
		obj["assetTypes"] = assetTypesProp
	}
	contentTypeProp, err := expandCloudAssetOrganizationFeedContentType(d.Get("content_type"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("content_type"); !isEmptyValue(reflect.ValueOf(contentTypeProp)) && (ok || !reflect.DeepEqual(v, contentTypeProp)) {
		obj["contentType"] = contentTypeProp
	}
	feedOutputConfigProp, err := expandCloudAssetOrganizationFeedFeedOutputConfig(d.Get("feed_output_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("feed_output_config"); !isEmptyValue(reflect.ValueOf(feedOutputConfigProp)) && (ok || !reflect.DeepEqual(v, feedOutputConfigProp)) {
		obj["feedOutputConfig"] = feedOutputConfigProp
	}
	conditionProp, err := expandCloudAssetOrganizationFeedCondition(d.Get("condition"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("condition"); !isEmptyValue(reflect.ValueOf(conditionProp)) && (ok || !reflect.DeepEqual(v, conditionProp)) {
		obj["condition"] = conditionProp
	}

	return resourceCloudAssetOrganizationFeedEncoder(d, config, obj)
}

func resourceCloudAssetOrganizationFeedEncoder(d TerraformResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
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

func expandCloudAssetOrganizationFeedAssetNames(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudAssetOrganizationFeedAssetTypes(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudAssetOrganizationFeedContentType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudAssetOrganizationFeedFeedOutputConfig(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedPubsubDestination, err := expandCloudAssetOrganizationFeedFeedOutputConfigPubsubDestination(original["pubsub_destination"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPubsubDestination); val.IsValid() && !isEmptyValue(val) {
		transformed["pubsubDestination"] = transformedPubsubDestination
	}

	return transformed, nil
}

func expandCloudAssetOrganizationFeedFeedOutputConfigPubsubDestination(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedTopic, err := expandCloudAssetOrganizationFeedFeedOutputConfigPubsubDestinationTopic(original["topic"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTopic); val.IsValid() && !isEmptyValue(val) {
		transformed["topic"] = transformedTopic
	}

	return transformed, nil
}

func expandCloudAssetOrganizationFeedFeedOutputConfigPubsubDestinationTopic(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudAssetOrganizationFeedCondition(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedExpression, err := expandCloudAssetOrganizationFeedConditionExpression(original["expression"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedExpression); val.IsValid() && !isEmptyValue(val) {
		transformed["expression"] = transformedExpression
	}

	transformedTitle, err := expandCloudAssetOrganizationFeedConditionTitle(original["title"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTitle); val.IsValid() && !isEmptyValue(val) {
		transformed["title"] = transformedTitle
	}

	transformedDescription, err := expandCloudAssetOrganizationFeedConditionDescription(original["description"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDescription); val.IsValid() && !isEmptyValue(val) {
		transformed["description"] = transformedDescription
	}

	transformedLocation, err := expandCloudAssetOrganizationFeedConditionLocation(original["location"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLocation); val.IsValid() && !isEmptyValue(val) {
		transformed["location"] = transformedLocation
	}

	return transformed, nil
}

func expandCloudAssetOrganizationFeedConditionExpression(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudAssetOrganizationFeedConditionTitle(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudAssetOrganizationFeedConditionDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudAssetOrganizationFeedConditionLocation(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
