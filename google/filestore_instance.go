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

import "reflect"

func GetFilestoreInstanceCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	name, err := assetName(d, config, "//filestore.googleapis.com/projects/{{project}}/locations/{{zone}}/instances/{{name}}")
	if err != nil {
		return []Asset{}, err
	}
	if obj, err := GetFilestoreInstanceApiObject(d, config); err == nil {
		return []Asset{{
			Name: name,
			Type: "filestore.googleapis.com/Instance",
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/filestore/v1/rest",
				DiscoveryName:        "Instance",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []Asset{}, err
	}
}

func GetFilestoreInstanceApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	descriptionProp, err := expandFilestoreInstanceDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	tierProp, err := expandFilestoreInstanceTier(d.Get("tier"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("tier"); !isEmptyValue(reflect.ValueOf(tierProp)) && (ok || !reflect.DeepEqual(v, tierProp)) {
		obj["tier"] = tierProp
	}
	labelsProp, err := expandFilestoreInstanceLabels(d.Get("labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	fileSharesProp, err := expandFilestoreInstanceFileShares(d.Get("file_shares"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("file_shares"); !isEmptyValue(reflect.ValueOf(fileSharesProp)) && (ok || !reflect.DeepEqual(v, fileSharesProp)) {
		obj["fileShares"] = fileSharesProp
	}
	networksProp, err := expandFilestoreInstanceNetworks(d.Get("networks"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("networks"); !isEmptyValue(reflect.ValueOf(networksProp)) && (ok || !reflect.DeepEqual(v, networksProp)) {
		obj["networks"] = networksProp
	}

	return obj, nil
}

func expandFilestoreInstanceDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandFilestoreInstanceTier(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandFilestoreInstanceLabels(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandFilestoreInstanceFileShares(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedName, err := expandFilestoreInstanceFileSharesName(original["name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedName); val.IsValid() && !isEmptyValue(val) {
			transformed["name"] = transformedName
		}

		transformedCapacityGb, err := expandFilestoreInstanceFileSharesCapacityGb(original["capacity_gb"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedCapacityGb); val.IsValid() && !isEmptyValue(val) {
			transformed["capacityGb"] = transformedCapacityGb
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandFilestoreInstanceFileSharesName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandFilestoreInstanceFileSharesCapacityGb(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandFilestoreInstanceNetworks(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedNetwork, err := expandFilestoreInstanceNetworksNetwork(original["network"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedNetwork); val.IsValid() && !isEmptyValue(val) {
			transformed["network"] = transformedNetwork
		}

		transformedModes, err := expandFilestoreInstanceNetworksModes(original["modes"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedModes); val.IsValid() && !isEmptyValue(val) {
			transformed["modes"] = transformedModes
		}

		transformedReservedIpRange, err := expandFilestoreInstanceNetworksReservedIpRange(original["reserved_ip_range"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedReservedIpRange); val.IsValid() && !isEmptyValue(val) {
			transformed["reservedIpRange"] = transformedReservedIpRange
		}

		transformedIpAddresses, err := expandFilestoreInstanceNetworksIpAddresses(original["ip_addresses"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedIpAddresses); val.IsValid() && !isEmptyValue(val) {
			transformed["ipAddresses"] = transformedIpAddresses
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandFilestoreInstanceNetworksNetwork(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandFilestoreInstanceNetworksModes(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandFilestoreInstanceNetworksReservedIpRange(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandFilestoreInstanceNetworksIpAddresses(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
