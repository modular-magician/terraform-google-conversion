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

const BigqueryDatapolicyDataPolicyAssetType string = "bigquerydatapolicy.googleapis.com/DataPolicy"

func resourceConverterBigqueryDatapolicyDataPolicy() ResourceConverter {
	return ResourceConverter{
		AssetType: BigqueryDatapolicyDataPolicyAssetType,
		Convert:   GetBigqueryDatapolicyDataPolicyCaiObject,
	}
}

func GetBigqueryDatapolicyDataPolicyCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]Asset, error) {
	name, err := assetName(d, config, "//bigquerydatapolicy.googleapis.com/projects/{{project}}/locations/{{location}}/dataPolicies/{{data_policy_id}}")
	if err != nil {
		return []Asset{}, err
	}
	if obj, err := GetBigqueryDatapolicyDataPolicyApiObject(d, config); err == nil {
		return []Asset{{
			Name: name,
			Type: BigqueryDatapolicyDataPolicyAssetType,
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/bigquerydatapolicy/v1/rest",
				DiscoveryName:        "DataPolicy",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []Asset{}, err
	}
}

func GetBigqueryDatapolicyDataPolicyApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	dataPolicyIdProp, err := expandBigqueryDatapolicyDataPolicyDataPolicyId(d.Get("data_policy_id"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("data_policy_id"); !tpgresource.IsEmptyValue(reflect.ValueOf(dataPolicyIdProp)) && (ok || !reflect.DeepEqual(v, dataPolicyIdProp)) {
		obj["dataPolicyId"] = dataPolicyIdProp
	}
	policyTagProp, err := expandBigqueryDatapolicyDataPolicyPolicyTag(d.Get("policy_tag"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("policy_tag"); !tpgresource.IsEmptyValue(reflect.ValueOf(policyTagProp)) && (ok || !reflect.DeepEqual(v, policyTagProp)) {
		obj["policyTag"] = policyTagProp
	}
	dataPolicyTypeProp, err := expandBigqueryDatapolicyDataPolicyDataPolicyType(d.Get("data_policy_type"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("data_policy_type"); !tpgresource.IsEmptyValue(reflect.ValueOf(dataPolicyTypeProp)) && (ok || !reflect.DeepEqual(v, dataPolicyTypeProp)) {
		obj["dataPolicyType"] = dataPolicyTypeProp
	}
	dataMaskingPolicyProp, err := expandBigqueryDatapolicyDataPolicyDataMaskingPolicy(d.Get("data_masking_policy"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("data_masking_policy"); !tpgresource.IsEmptyValue(reflect.ValueOf(dataMaskingPolicyProp)) && (ok || !reflect.DeepEqual(v, dataMaskingPolicyProp)) {
		obj["dataMaskingPolicy"] = dataMaskingPolicyProp
	}

	return obj, nil
}

func expandBigqueryDatapolicyDataPolicyDataPolicyId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryDatapolicyDataPolicyPolicyTag(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryDatapolicyDataPolicyDataPolicyType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryDatapolicyDataPolicyDataMaskingPolicy(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedPredefinedExpression, err := expandBigqueryDatapolicyDataPolicyDataMaskingPolicyPredefinedExpression(original["predefined_expression"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPredefinedExpression); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["predefinedExpression"] = transformedPredefinedExpression
	}

	return transformed, nil
}

func expandBigqueryDatapolicyDataPolicyDataMaskingPolicyPredefinedExpression(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
