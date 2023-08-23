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

const BigqueryAnalyticsHubListingAssetType string = "analyticshub.googleapis.com/Listing"

func ResourceConverterBigqueryAnalyticsHubListing() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: BigqueryAnalyticsHubListingAssetType,
		Convert:   GetBigqueryAnalyticsHubListingCaiObject,
	}
}

func GetBigqueryAnalyticsHubListingCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//analyticshub.googleapis.com/projects/{{project}}/locations/{{location}}/dataExchanges/{{data_exchange_id}}/listings/{{listing_id}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetBigqueryAnalyticsHubListingApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: BigqueryAnalyticsHubListingAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1beta1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/analyticshub/v1beta1/rest",
				DiscoveryName:        "Listing",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetBigqueryAnalyticsHubListingApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	displayNameProp, err := expandBigqueryAnalyticsHubListingDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	descriptionProp, err := expandBigqueryAnalyticsHubListingDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	primaryContactProp, err := expandBigqueryAnalyticsHubListingPrimaryContact(d.Get("primary_contact"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("primary_contact"); !tpgresource.IsEmptyValue(reflect.ValueOf(primaryContactProp)) && (ok || !reflect.DeepEqual(v, primaryContactProp)) {
		obj["primaryContact"] = primaryContactProp
	}
	documentationProp, err := expandBigqueryAnalyticsHubListingDocumentation(d.Get("documentation"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("documentation"); !tpgresource.IsEmptyValue(reflect.ValueOf(documentationProp)) && (ok || !reflect.DeepEqual(v, documentationProp)) {
		obj["documentation"] = documentationProp
	}
	iconProp, err := expandBigqueryAnalyticsHubListingIcon(d.Get("icon"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("icon"); !tpgresource.IsEmptyValue(reflect.ValueOf(iconProp)) && (ok || !reflect.DeepEqual(v, iconProp)) {
		obj["icon"] = iconProp
	}
	requestAccessProp, err := expandBigqueryAnalyticsHubListingRequestAccess(d.Get("request_access"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("request_access"); !tpgresource.IsEmptyValue(reflect.ValueOf(requestAccessProp)) && (ok || !reflect.DeepEqual(v, requestAccessProp)) {
		obj["requestAccess"] = requestAccessProp
	}
	dataProviderProp, err := expandBigqueryAnalyticsHubListingDataProvider(d.Get("data_provider"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("data_provider"); !tpgresource.IsEmptyValue(reflect.ValueOf(dataProviderProp)) && (ok || !reflect.DeepEqual(v, dataProviderProp)) {
		obj["dataProvider"] = dataProviderProp
	}
	publisherProp, err := expandBigqueryAnalyticsHubListingPublisher(d.Get("publisher"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("publisher"); !tpgresource.IsEmptyValue(reflect.ValueOf(publisherProp)) && (ok || !reflect.DeepEqual(v, publisherProp)) {
		obj["publisher"] = publisherProp
	}
	categoriesProp, err := expandBigqueryAnalyticsHubListingCategories(d.Get("categories"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("categories"); !tpgresource.IsEmptyValue(reflect.ValueOf(categoriesProp)) && (ok || !reflect.DeepEqual(v, categoriesProp)) {
		obj["categories"] = categoriesProp
	}
	bigqueryDatasetProp, err := expandBigqueryAnalyticsHubListingBigqueryDataset(d.Get("bigquery_dataset"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("bigquery_dataset"); !tpgresource.IsEmptyValue(reflect.ValueOf(bigqueryDatasetProp)) && (ok || !reflect.DeepEqual(v, bigqueryDatasetProp)) {
		obj["bigqueryDataset"] = bigqueryDatasetProp
	}

	return obj, nil
}

func expandBigqueryAnalyticsHubListingDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryAnalyticsHubListingDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryAnalyticsHubListingPrimaryContact(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryAnalyticsHubListingDocumentation(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryAnalyticsHubListingIcon(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryAnalyticsHubListingRequestAccess(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryAnalyticsHubListingDataProvider(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedName, err := expandBigqueryAnalyticsHubListingDataProviderName(original["name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["name"] = transformedName
	}

	transformedPrimaryContact, err := expandBigqueryAnalyticsHubListingDataProviderPrimaryContact(original["primary_contact"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPrimaryContact); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["primaryContact"] = transformedPrimaryContact
	}

	return transformed, nil
}

func expandBigqueryAnalyticsHubListingDataProviderName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryAnalyticsHubListingDataProviderPrimaryContact(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryAnalyticsHubListingPublisher(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedName, err := expandBigqueryAnalyticsHubListingPublisherName(original["name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["name"] = transformedName
	}

	transformedPrimaryContact, err := expandBigqueryAnalyticsHubListingPublisherPrimaryContact(original["primary_contact"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPrimaryContact); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["primaryContact"] = transformedPrimaryContact
	}

	return transformed, nil
}

func expandBigqueryAnalyticsHubListingPublisherName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryAnalyticsHubListingPublisherPrimaryContact(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryAnalyticsHubListingCategories(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryAnalyticsHubListingBigqueryDataset(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedDataset, err := expandBigqueryAnalyticsHubListingBigqueryDatasetDataset(original["dataset"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDataset); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["dataset"] = transformedDataset
	}

	return transformed, nil
}

func expandBigqueryAnalyticsHubListingBigqueryDatasetDataset(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
