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

package activedirectory

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const ActiveDirectoryPeeringAssetType string = "managedidentities.googleapis.com/Peering"

func ResourceConverterActiveDirectoryPeering() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: ActiveDirectoryPeeringAssetType,
		Convert:   GetActiveDirectoryPeeringCaiObject,
	}
}

func GetActiveDirectoryPeeringCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//managedidentities.googleapis.com/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetActiveDirectoryPeeringApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: ActiveDirectoryPeeringAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1beta1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/managedidentities/v1beta1/rest",
				DiscoveryName:        "Peering",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetActiveDirectoryPeeringApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	authorizedNetworkProp, err := expandActiveDirectoryPeeringAuthorizedNetwork(d.Get("authorized_network"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("authorized_network"); !tpgresource.IsEmptyValue(reflect.ValueOf(authorizedNetworkProp)) && (ok || !reflect.DeepEqual(v, authorizedNetworkProp)) {
		obj["authorizedNetwork"] = authorizedNetworkProp
	}
	domainResourceProp, err := expandActiveDirectoryPeeringDomainResource(d.Get("domain_resource"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("domain_resource"); !tpgresource.IsEmptyValue(reflect.ValueOf(domainResourceProp)) && (ok || !reflect.DeepEqual(v, domainResourceProp)) {
		obj["domainResource"] = domainResourceProp
	}
	statusMessageProp, err := expandActiveDirectoryPeeringStatusMessage(d.Get("status_message"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("status_message"); !tpgresource.IsEmptyValue(reflect.ValueOf(statusMessageProp)) && (ok || !reflect.DeepEqual(v, statusMessageProp)) {
		obj["statusMessage"] = statusMessageProp
	}
	labelsProp, err := expandActiveDirectoryPeeringEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	return obj, nil
}

func expandActiveDirectoryPeeringAuthorizedNetwork(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandActiveDirectoryPeeringDomainResource(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandActiveDirectoryPeeringStatusMessage(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandActiveDirectoryPeeringEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
