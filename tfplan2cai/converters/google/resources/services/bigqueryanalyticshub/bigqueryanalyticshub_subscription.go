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

package bigqueryanalyticshub

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const BigqueryAnalyticsHubSubscriptionAssetType string = "analyticshub.googleapis.com/Subscription"

func ResourceConverterBigqueryAnalyticsHubSubscription() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: BigqueryAnalyticsHubSubscriptionAssetType,
		Convert:   GetBigqueryAnalyticsHubSubscriptionCaiObject,
	}
}

func GetBigqueryAnalyticsHubSubscriptionCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//analyticshub.googleapis.com/projects/{{project}}/locations/{{location}}/subscriptions/{{subscription_id}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetBigqueryAnalyticsHubSubscriptionApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: BigqueryAnalyticsHubSubscriptionAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1beta1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/analyticshub/v1beta1/rest",
				DiscoveryName:        "Subscription",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetBigqueryAnalyticsHubSubscriptionApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	destinationDatasetProp, err := expandBigqueryAnalyticsHubSubscriptionDestinationDataset(d.Get("destination_dataset"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("destination_dataset"); !tpgresource.IsEmptyValue(reflect.ValueOf(destinationDatasetProp)) && (ok || !reflect.DeepEqual(v, destinationDatasetProp)) {
		obj["destinationDataset"] = destinationDatasetProp
	}

	return obj, nil
}

func expandBigqueryAnalyticsHubSubscriptionDestinationDataset(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedDatasetReference, err := expandBigqueryAnalyticsHubSubscriptionDestinationDatasetDatasetReference(original["dataset_reference"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDatasetReference); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["datasetReference"] = transformedDatasetReference
	}

	transformedFriendlyName, err := expandBigqueryAnalyticsHubSubscriptionDestinationDatasetFriendlyName(original["friendly_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedFriendlyName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["friendlyName"] = transformedFriendlyName
	}

	transformedDescription, err := expandBigqueryAnalyticsHubSubscriptionDestinationDatasetDescription(original["description"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDescription); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["description"] = transformedDescription
	}

	transformedLocation, err := expandBigqueryAnalyticsHubSubscriptionDestinationDatasetLocation(original["location"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLocation); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["location"] = transformedLocation
	}

	return transformed, nil
}

func expandBigqueryAnalyticsHubSubscriptionDestinationDatasetDatasetReference(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedDatasetId, err := expandBigqueryAnalyticsHubSubscriptionDestinationDatasetDatasetReferenceDatasetId(original["dataset_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDatasetId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["datasetId"] = transformedDatasetId
	}

	transformedProjectId, err := expandBigqueryAnalyticsHubSubscriptionDestinationDatasetDatasetReferenceProjectId(original["project_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedProjectId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["projectId"] = transformedProjectId
	}

	return transformed, nil
}

func expandBigqueryAnalyticsHubSubscriptionDestinationDatasetDatasetReferenceDatasetId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryAnalyticsHubSubscriptionDestinationDatasetDatasetReferenceProjectId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryAnalyticsHubSubscriptionDestinationDatasetFriendlyName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryAnalyticsHubSubscriptionDestinationDatasetDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryAnalyticsHubSubscriptionDestinationDatasetLocation(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
