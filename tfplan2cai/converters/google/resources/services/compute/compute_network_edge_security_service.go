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

package compute

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgresource"
	transport_tpg "github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/transport"
)

const ComputeNetworkEdgeSecurityServiceAssetType string = "compute.googleapis.com/NetworkEdgeSecurityService"

func ResourceConverterComputeNetworkEdgeSecurityService() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType: ComputeNetworkEdgeSecurityServiceAssetType,
		Convert:   GetComputeNetworkEdgeSecurityServiceCaiObject,
	}
}

func GetComputeNetworkEdgeSecurityServiceCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	name, err := tpgresource.AssetName(d, config, "//compute.googleapis.com/projects/{{project}}/regions/{{region}}/networkEdgeSecurityServices/{{name}}")
	if err != nil {
		return []tpgresource.Asset{}, err
	}
	if obj, err := GetComputeNetworkEdgeSecurityServiceApiObject(d, config); err == nil {
		return []tpgresource.Asset{{
			Name: name,
			Type: ComputeNetworkEdgeSecurityServiceAssetType,
			Resource: &tpgresource.AssetResource{
				Version:              "beta",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/compute/beta/rest",
				DiscoveryName:        "NetworkEdgeSecurityService",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []tpgresource.Asset{}, err
	}
}

func GetComputeNetworkEdgeSecurityServiceApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandComputeNetworkEdgeSecurityServiceName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	descriptionProp, err := expandComputeNetworkEdgeSecurityServiceDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	fingerprintProp, err := expandComputeNetworkEdgeSecurityServiceFingerprint(d.Get("fingerprint"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("fingerprint"); !tpgresource.IsEmptyValue(reflect.ValueOf(fingerprintProp)) && (ok || !reflect.DeepEqual(v, fingerprintProp)) {
		obj["fingerprint"] = fingerprintProp
	}
	securityPolicyProp, err := expandComputeNetworkEdgeSecurityServiceSecurityPolicy(d.Get("security_policy"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("security_policy"); !tpgresource.IsEmptyValue(reflect.ValueOf(securityPolicyProp)) && (ok || !reflect.DeepEqual(v, securityPolicyProp)) {
		obj["securityPolicy"] = securityPolicyProp
	}

	return obj, nil
}

func expandComputeNetworkEdgeSecurityServiceName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkEdgeSecurityServiceDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkEdgeSecurityServiceFingerprint(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkEdgeSecurityServiceSecurityPolicy(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
