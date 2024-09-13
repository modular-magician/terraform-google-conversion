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

package discoveryengine

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const DiscoveryEngineTargetSiteAssetType string = "{{location}}-discoveryengine.googleapis.com/TargetSite"

func ResourceConverterDiscoveryEngineTargetSite() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: DiscoveryEngineTargetSiteAssetType,
		Convert:   GetDiscoveryEngineTargetSiteCaiObject,
	}
}

func GetDiscoveryEngineTargetSiteCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//{{location}}-discoveryengine.googleapis.com/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetDiscoveryEngineTargetSiteApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: DiscoveryEngineTargetSiteAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/{{location}}-discoveryengine/v1/rest",
				DiscoveryName:        "TargetSite",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetDiscoveryEngineTargetSiteApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	providedUriPatternProp, err := expandDiscoveryEngineTargetSiteProvidedUriPattern(d.Get("provided_uri_pattern"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("provided_uri_pattern"); !tpgresource.IsEmptyValue(reflect.ValueOf(providedUriPatternProp)) && (ok || !reflect.DeepEqual(v, providedUriPatternProp)) {
		obj["providedUriPattern"] = providedUriPatternProp
	}
	typeProp, err := expandDiscoveryEngineTargetSiteType(d.Get("type"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("type"); !tpgresource.IsEmptyValue(reflect.ValueOf(typeProp)) && (ok || !reflect.DeepEqual(v, typeProp)) {
		obj["type"] = typeProp
	}
	exactMatchProp, err := expandDiscoveryEngineTargetSiteExactMatch(d.Get("exact_match"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("exact_match"); !tpgresource.IsEmptyValue(reflect.ValueOf(exactMatchProp)) && (ok || !reflect.DeepEqual(v, exactMatchProp)) {
		obj["exactMatch"] = exactMatchProp
	}

	return obj, nil
}

func expandDiscoveryEngineTargetSiteProvidedUriPattern(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDiscoveryEngineTargetSiteType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDiscoveryEngineTargetSiteExactMatch(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
