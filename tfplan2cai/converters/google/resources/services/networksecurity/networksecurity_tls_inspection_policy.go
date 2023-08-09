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

package networksecurity

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const NetworkSecurityTlsInspectionPolicyAssetType string = "networksecurity.googleapis.com/TlsInspectionPolicy"

func ResourceConverterNetworkSecurityTlsInspectionPolicy() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType: NetworkSecurityTlsInspectionPolicyAssetType,
		Convert:   GetNetworkSecurityTlsInspectionPolicyCaiObject,
	}
}

func GetNetworkSecurityTlsInspectionPolicyCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	name, err := tpgresource.AssetName(d, config, "//networksecurity.googleapis.com/projects/{{project}}/locations/{{location}}/tlsInspectionPolicies/{{name}}")
	if err != nil {
		return []tpgresource.Asset{}, err
	}
	if obj, err := GetNetworkSecurityTlsInspectionPolicyApiObject(d, config); err == nil {
		return []tpgresource.Asset{{
			Name: name,
			Type: NetworkSecurityTlsInspectionPolicyAssetType,
			Resource: &tpgresource.AssetResource{
				Version:              "v1beta1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/networksecurity/v1beta1/rest",
				DiscoveryName:        "TlsInspectionPolicy",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []tpgresource.Asset{}, err
	}
}

func GetNetworkSecurityTlsInspectionPolicyApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	descriptionProp, err := expandNetworkSecurityTlsInspectionPolicyDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	caPoolProp, err := expandNetworkSecurityTlsInspectionPolicyCaPool(d.Get("ca_pool"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("ca_pool"); !tpgresource.IsEmptyValue(reflect.ValueOf(caPoolProp)) && (ok || !reflect.DeepEqual(v, caPoolProp)) {
		obj["caPool"] = caPoolProp
	}
	excludePublicCaSetProp, err := expandNetworkSecurityTlsInspectionPolicyExcludePublicCaSet(d.Get("exclude_public_ca_set"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("exclude_public_ca_set"); !tpgresource.IsEmptyValue(reflect.ValueOf(excludePublicCaSetProp)) && (ok || !reflect.DeepEqual(v, excludePublicCaSetProp)) {
		obj["excludePublicCaSet"] = excludePublicCaSetProp
	}

	return obj, nil
}

func expandNetworkSecurityTlsInspectionPolicyDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkSecurityTlsInspectionPolicyCaPool(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkSecurityTlsInspectionPolicyExcludePublicCaSet(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
