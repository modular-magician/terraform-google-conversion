// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Nathan is editing this to generate diffs in lots of files.
//     He won't submit this change.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"reflect"
)

func GetComputeAddressCaiObject(d TerraformResourceData, config *Config) (Asset, error) {
	name, err := assetName(d, config, "//compute.googleapis.com/projects/{{project}}/regions/{{region}}/addresses/{{name}}")
	if err != nil {
		return Asset{}, err
	}
	if obj, err := GetComputeAddressApiObject(d, config); err == nil {
		return Asset{
			Name: name,
			Type: "compute.googleapis.com/Address",
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/compute/v1/rest",
				DiscoveryName:        "Address",
				Data:                 obj,
			},
		}, nil
	} else {
		return Asset{}, err
	}
}

func GetComputeAddressApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	addressProp, err := expandComputeAddressAddress(d.Get("address"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("address"); !isEmptyValue(reflect.ValueOf(addressProp)) && (ok || !reflect.DeepEqual(v, addressProp)) {
		obj["address"] = addressProp
	}
	addressTypeProp, err := expandComputeAddressAddressType(d.Get("address_type"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("address_type"); !isEmptyValue(reflect.ValueOf(addressTypeProp)) && (ok || !reflect.DeepEqual(v, addressTypeProp)) {
		obj["addressType"] = addressTypeProp
	}
	descriptionProp, err := expandComputeAddressDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	nameProp, err := expandComputeAddressName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	purposeProp, err := expandComputeAddressPurpose(d.Get("purpose"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("purpose"); !isEmptyValue(reflect.ValueOf(purposeProp)) && (ok || !reflect.DeepEqual(v, purposeProp)) {
		obj["purpose"] = purposeProp
	}
	networkTierProp, err := expandComputeAddressNetworkTier(d.Get("network_tier"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("network_tier"); !isEmptyValue(reflect.ValueOf(networkTierProp)) && (ok || !reflect.DeepEqual(v, networkTierProp)) {
		obj["networkTier"] = networkTierProp
	}
	subnetworkProp, err := expandComputeAddressSubnetwork(d.Get("subnetwork"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("subnetwork"); !isEmptyValue(reflect.ValueOf(subnetworkProp)) && (ok || !reflect.DeepEqual(v, subnetworkProp)) {
		obj["subnetwork"] = subnetworkProp
	}
	regionProp, err := expandComputeAddressRegion(d.Get("region"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("region"); !isEmptyValue(reflect.ValueOf(regionProp)) && (ok || !reflect.DeepEqual(v, regionProp)) {
		obj["region"] = regionProp
	}

	return obj, nil
}

func expandComputeAddressAddress(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeAddressAddressType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeAddressDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeAddressName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeAddressPurpose(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeAddressNetworkTier(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeAddressSubnetwork(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseRegionalFieldValue("subnetworks", v.(string), "project", "region", "zone", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for subnetwork: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeAddressRegion(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseGlobalFieldValue("regions", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for region: %s", err)
	}
	return f.RelativeLink(), nil
}
