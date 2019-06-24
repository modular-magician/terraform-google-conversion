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

	"github.com/hashicorp/terraform/helper/schema"
)

func GetDnsManagedZoneCaiObject(d TerraformResourceData, config *Config) (Asset, error) {
	name, err := assetName(d, config, "//dns.googleapis.com/projects/{{project}}/managedZones/{{name}}")
	if err != nil {
		return Asset{}, err
	}
	if obj, err := GetDnsManagedZoneApiObject(d, config); err == nil {
		return Asset{
			Name: name,
			Type: "dns.googleapis.com/ManagedZone",
			Resource: &AssetResource{
				Version:              "managedZones",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/dns/managedZones/rest",
				DiscoveryName:        "ManagedZone",
				Data:                 obj,
			},
		}, nil
	} else {
		return Asset{}, err
	}
}

func GetDnsManagedZoneApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	descriptionProp, err := expandDnsManagedZoneDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	dnsNameProp, err := expandDnsManagedZoneDnsName(d.Get("dns_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("dns_name"); !isEmptyValue(reflect.ValueOf(dnsNameProp)) && (ok || !reflect.DeepEqual(v, dnsNameProp)) {
		obj["dnsName"] = dnsNameProp
	}
	dnssecConfigProp, err := expandDnsManagedZoneDnssecConfig(d.Get("dnssec_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("dnssec_config"); !isEmptyValue(reflect.ValueOf(dnssecConfigProp)) && (ok || !reflect.DeepEqual(v, dnssecConfigProp)) {
		obj["dnssecConfig"] = dnssecConfigProp
	}
	nameProp, err := expandDnsManagedZoneName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	labelsProp, err := expandDnsManagedZoneLabels(d.Get("labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	visibilityProp, err := expandDnsManagedZoneVisibility(d.Get("visibility"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("visibility"); !isEmptyValue(reflect.ValueOf(visibilityProp)) && (ok || !reflect.DeepEqual(v, visibilityProp)) {
		obj["visibility"] = visibilityProp
	}
	privateVisibilityConfigProp, err := expandDnsManagedZonePrivateVisibilityConfig(d.Get("private_visibility_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("private_visibility_config"); !isEmptyValue(reflect.ValueOf(privateVisibilityConfigProp)) && (ok || !reflect.DeepEqual(v, privateVisibilityConfigProp)) {
		obj["privateVisibilityConfig"] = privateVisibilityConfigProp
	}

	return obj, nil
}

func expandDnsManagedZoneDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDnsManagedZoneDnsName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDnsManagedZoneDnssecConfig(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedKind, err := expandDnsManagedZoneDnssecConfigKind(original["kind"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKind); val.IsValid() && !isEmptyValue(val) {
		transformed["kind"] = transformedKind
	}

	transformedNonExistence, err := expandDnsManagedZoneDnssecConfigNonExistence(original["non_existence"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNonExistence); val.IsValid() && !isEmptyValue(val) {
		transformed["nonExistence"] = transformedNonExistence
	}

	transformedState, err := expandDnsManagedZoneDnssecConfigState(original["state"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedState); val.IsValid() && !isEmptyValue(val) {
		transformed["state"] = transformedState
	}

	transformedDefaultKeySpecs, err := expandDnsManagedZoneDnssecConfigDefaultKeySpecs(original["default_key_specs"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDefaultKeySpecs); val.IsValid() && !isEmptyValue(val) {
		transformed["defaultKeySpecs"] = transformedDefaultKeySpecs
	}

	return transformed, nil
}

func expandDnsManagedZoneDnssecConfigKind(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDnsManagedZoneDnssecConfigNonExistence(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDnsManagedZoneDnssecConfigState(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDnsManagedZoneDnssecConfigDefaultKeySpecs(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedAlgorithm, err := expandDnsManagedZoneDnssecConfigDefaultKeySpecsAlgorithm(original["algorithm"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedAlgorithm); val.IsValid() && !isEmptyValue(val) {
			transformed["algorithm"] = transformedAlgorithm
		}

		transformedKeyLength, err := expandDnsManagedZoneDnssecConfigDefaultKeySpecsKeyLength(original["key_length"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedKeyLength); val.IsValid() && !isEmptyValue(val) {
			transformed["keyLength"] = transformedKeyLength
		}

		transformedKeyType, err := expandDnsManagedZoneDnssecConfigDefaultKeySpecsKeyType(original["key_type"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedKeyType); val.IsValid() && !isEmptyValue(val) {
			transformed["keyType"] = transformedKeyType
		}

		transformedKind, err := expandDnsManagedZoneDnssecConfigDefaultKeySpecsKind(original["kind"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedKind); val.IsValid() && !isEmptyValue(val) {
			transformed["kind"] = transformedKind
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandDnsManagedZoneDnssecConfigDefaultKeySpecsAlgorithm(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDnsManagedZoneDnssecConfigDefaultKeySpecsKeyLength(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDnsManagedZoneDnssecConfigDefaultKeySpecsKeyType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDnsManagedZoneDnssecConfigDefaultKeySpecsKind(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDnsManagedZoneName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDnsManagedZoneLabels(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandDnsManagedZoneVisibility(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDnsManagedZonePrivateVisibilityConfig(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedNetworks, err := expandDnsManagedZonePrivateVisibilityConfigNetworks(original["networks"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNetworks); val.IsValid() && !isEmptyValue(val) {
		transformed["networks"] = transformedNetworks
	}

	return transformed, nil
}

func expandDnsManagedZonePrivateVisibilityConfigNetworks(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedNetworkUrl, err := expandDnsManagedZonePrivateVisibilityConfigNetworksNetworkUrl(original["network_url"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedNetworkUrl); val.IsValid() && !isEmptyValue(val) {
			transformed["networkUrl"] = transformedNetworkUrl
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandDnsManagedZonePrivateVisibilityConfigNetworksNetworkUrl(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
