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

package vertexai

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const VertexAIFeatureGroupFeatureAssetType string = "aiplatform.googleapis.com/FeatureGroupFeature"

func ResourceConverterVertexAIFeatureGroupFeature() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: VertexAIFeatureGroupFeatureAssetType,
		Convert:   GetVertexAIFeatureGroupFeatureCaiObject,
	}
}

func GetVertexAIFeatureGroupFeatureCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//aiplatform.googleapis.com/projects/{{project}}/locations/{{region}}/featureGroups/{{feature_group}}/features/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetVertexAIFeatureGroupFeatureApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: VertexAIFeatureGroupFeatureAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1beta1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/aiplatform/v1beta1/rest",
				DiscoveryName:        "FeatureGroupFeature",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetVertexAIFeatureGroupFeatureApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	descriptionProp, err := expandVertexAIFeatureGroupFeatureDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	versionColumnNameProp, err := expandVertexAIFeatureGroupFeatureVersionColumnName(d.Get("version_column_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("version_column_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(versionColumnNameProp)) && (ok || !reflect.DeepEqual(v, versionColumnNameProp)) {
		obj["versionColumnName"] = versionColumnNameProp
	}
	labelsProp, err := expandVertexAIFeatureGroupFeatureEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	return obj, nil
}

func expandVertexAIFeatureGroupFeatureDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVertexAIFeatureGroupFeatureVersionColumnName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVertexAIFeatureGroupFeatureEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
