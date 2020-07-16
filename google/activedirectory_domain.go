// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
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

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func GetActiveDirectoryDomainCaiObject(d TerraformResourceData, config *Config) (Asset, error) {
	name, err := assetName(d, config, "//activedirectory.googleapis.com/{{name}}")
	if err != nil {
		return Asset{}, err
	}
	if obj, err := GetActiveDirectoryDomainApiObject(d, config); err == nil {
		return Asset{
			Name: name,
			Type: "activedirectory.googleapis.com/Domain",
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/activedirectory/v1/rest",
				DiscoveryName:        "Domain",
				Data:                 obj,
			},
		}, nil
	} else {
		return Asset{}, err
	}
}

func GetActiveDirectoryDomainApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	labelsProp, err := expandActiveDirectoryDomainLabels(d.Get("labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	authorizedNetworksProp, err := expandActiveDirectoryDomainAuthorizedNetworks(d.Get("authorized_networks"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("authorized_networks"); !isEmptyValue(reflect.ValueOf(authorizedNetworksProp)) && (ok || !reflect.DeepEqual(v, authorizedNetworksProp)) {
		obj["authorizedNetworks"] = authorizedNetworksProp
	}
	reservedIpRangeProp, err := expandActiveDirectoryDomainReservedIpRange(d.Get("reserved_ip_range"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("reserved_ip_range"); !isEmptyValue(reflect.ValueOf(reservedIpRangeProp)) && (ok || !reflect.DeepEqual(v, reservedIpRangeProp)) {
		obj["reservedIpRange"] = reservedIpRangeProp
	}
	locationsProp, err := expandActiveDirectoryDomainLocations(d.Get("locations"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("locations"); !isEmptyValue(reflect.ValueOf(locationsProp)) && (ok || !reflect.DeepEqual(v, locationsProp)) {
		obj["locations"] = locationsProp
	}
	adminProp, err := expandActiveDirectoryDomainAdmin(d.Get("admin"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("admin"); !isEmptyValue(reflect.ValueOf(adminProp)) && (ok || !reflect.DeepEqual(v, adminProp)) {
		obj["admin"] = adminProp
	}

	return obj, nil
}

func expandActiveDirectoryDomainLabels(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandActiveDirectoryDomainAuthorizedNetworks(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	return v, nil
}

func expandActiveDirectoryDomainReservedIpRange(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandActiveDirectoryDomainLocations(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandActiveDirectoryDomainAdmin(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
