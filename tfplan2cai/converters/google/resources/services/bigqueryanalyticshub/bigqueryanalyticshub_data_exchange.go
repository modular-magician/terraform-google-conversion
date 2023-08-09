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

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const BigqueryAnalyticsHubDataExchangeAssetType string = "analyticshub.googleapis.com/DataExchange"

func ResourceConverterBigqueryAnalyticsHubDataExchange() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType: BigqueryAnalyticsHubDataExchangeAssetType,
		Convert:   GetBigqueryAnalyticsHubDataExchangeCaiObject,
	}
}

func GetBigqueryAnalyticsHubDataExchangeCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	name, err := tpgresource.AssetName(d, config, "//analyticshub.googleapis.com/projects/{{project}}/locations/{{location}}/dataExchanges/{{data_exchange_id}}")
	if err != nil {
		return []tpgresource.Asset{}, err
	}
	if obj, err := GetBigqueryAnalyticsHubDataExchangeApiObject(d, config); err == nil {
		return []tpgresource.Asset{{
			Name: name,
			Type: BigqueryAnalyticsHubDataExchangeAssetType,
			Resource: &tpgresource.AssetResource{
				Version:              "v1beta1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/analyticshub/v1beta1/rest",
				DiscoveryName:        "DataExchange",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []tpgresource.Asset{}, err
	}
}

func GetBigqueryAnalyticsHubDataExchangeApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	displayNameProp, err := expandBigqueryAnalyticsHubDataExchangeDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	descriptionProp, err := expandBigqueryAnalyticsHubDataExchangeDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	primaryContactProp, err := expandBigqueryAnalyticsHubDataExchangePrimaryContact(d.Get("primary_contact"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("primary_contact"); !tpgresource.IsEmptyValue(reflect.ValueOf(primaryContactProp)) && (ok || !reflect.DeepEqual(v, primaryContactProp)) {
		obj["primaryContact"] = primaryContactProp
	}
	documentationProp, err := expandBigqueryAnalyticsHubDataExchangeDocumentation(d.Get("documentation"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("documentation"); !tpgresource.IsEmptyValue(reflect.ValueOf(documentationProp)) && (ok || !reflect.DeepEqual(v, documentationProp)) {
		obj["documentation"] = documentationProp
	}
	iconProp, err := expandBigqueryAnalyticsHubDataExchangeIcon(d.Get("icon"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("icon"); !tpgresource.IsEmptyValue(reflect.ValueOf(iconProp)) && (ok || !reflect.DeepEqual(v, iconProp)) {
		obj["icon"] = iconProp
	}

	return obj, nil
}

func expandBigqueryAnalyticsHubDataExchangeDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryAnalyticsHubDataExchangeDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryAnalyticsHubDataExchangePrimaryContact(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryAnalyticsHubDataExchangeDocumentation(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryAnalyticsHubDataExchangeIcon(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
