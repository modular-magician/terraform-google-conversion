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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgresource"
	transport_tpg "github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/transport"
)

const ActiveDirectoryDomainAssetType string = "managedidentities.googleapis.com/Domain"

func ResourceConverterActiveDirectoryDomain() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType: ActiveDirectoryDomainAssetType,
		Convert:   GetActiveDirectoryDomainCaiObject,
	}
}

func GetActiveDirectoryDomainCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	name, err := tpgresource.AssetName(d, config, "//managedidentities.googleapis.com/{{name}}")
	if err != nil {
		return []tpgresource.Asset{}, err
	}
	if obj, err := GetActiveDirectoryDomainApiObject(d, config); err == nil {
		return []tpgresource.Asset{{
			Name: name,
			Type: ActiveDirectoryDomainAssetType,
			Resource: &tpgresource.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/managedidentities/v1/rest",
				DiscoveryName:        "Domain",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []tpgresource.Asset{}, err
	}
}

func GetActiveDirectoryDomainApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	labelsProp, err := expandActiveDirectoryDomainLabels(d.Get("labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	authorizedNetworksProp, err := expandActiveDirectoryDomainAuthorizedNetworks(d.Get("authorized_networks"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("authorized_networks"); !tpgresource.IsEmptyValue(reflect.ValueOf(authorizedNetworksProp)) && (ok || !reflect.DeepEqual(v, authorizedNetworksProp)) {
		obj["authorizedNetworks"] = authorizedNetworksProp
	}
	reservedIpRangeProp, err := expandActiveDirectoryDomainReservedIpRange(d.Get("reserved_ip_range"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("reserved_ip_range"); !tpgresource.IsEmptyValue(reflect.ValueOf(reservedIpRangeProp)) && (ok || !reflect.DeepEqual(v, reservedIpRangeProp)) {
		obj["reservedIpRange"] = reservedIpRangeProp
	}
	locationsProp, err := expandActiveDirectoryDomainLocations(d.Get("locations"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("locations"); !tpgresource.IsEmptyValue(reflect.ValueOf(locationsProp)) && (ok || !reflect.DeepEqual(v, locationsProp)) {
		obj["locations"] = locationsProp
	}
	adminProp, err := expandActiveDirectoryDomainAdmin(d.Get("admin"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("admin"); !tpgresource.IsEmptyValue(reflect.ValueOf(adminProp)) && (ok || !reflect.DeepEqual(v, adminProp)) {
		obj["admin"] = adminProp
	}

	return obj, nil
}

func expandActiveDirectoryDomainLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandActiveDirectoryDomainAuthorizedNetworks(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	return v, nil
}

func expandActiveDirectoryDomainReservedIpRange(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandActiveDirectoryDomainLocations(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandActiveDirectoryDomainAdmin(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
