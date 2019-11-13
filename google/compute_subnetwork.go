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
	"fmt"
	"net"
	"reflect"

	"github.com/apparentlymart/go-cidr/cidr"
)

// Whether the IP CIDR change shrinks the block.
func isShrinkageIpCidr(old, new, _ interface{}) bool {
	_, oldCidr, oldErr := net.ParseCIDR(old.(string))
	_, newCidr, newErr := net.ParseCIDR(new.(string))

	if oldErr != nil || newErr != nil {
		// This should never happen. The ValidateFunc on the field ensures it.
		return false
	}

	oldStart, oldEnd := cidr.AddressRange(oldCidr)

	if newCidr.Contains(oldStart) && newCidr.Contains(oldEnd) {
		// This is a CIDR range expansion, no need to ForceNew, we have an update method for it.
		return false
	}

	return true
}

func GetComputeSubnetworkCaiObject(d TerraformResourceData, config *Config) (Asset, error) {
	name, err := assetName(d, config, "//compute.googleapis.com/projects/{{project}}/regions/{{region}}/subnetworks/{{name}}")
	if err != nil {
		return Asset{}, err
	}
	if obj, err := GetComputeSubnetworkApiObject(d, config); err == nil {
		return Asset{
			Name: name,
			Type: "compute.googleapis.com/Subnetwork",
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/compute/v1/rest",
				DiscoveryName:        "Subnetwork",
				Data:                 obj,
			},
		}, nil
	} else {
		return Asset{}, err
	}
}

func GetComputeSubnetworkApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	descriptionProp, err := expandComputeSubnetworkDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	ipCidrRangeProp, err := expandComputeSubnetworkIpCidrRange(d.Get("ip_cidr_range"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("ip_cidr_range"); !isEmptyValue(reflect.ValueOf(ipCidrRangeProp)) && (ok || !reflect.DeepEqual(v, ipCidrRangeProp)) {
		obj["ipCidrRange"] = ipCidrRangeProp
	}
	nameProp, err := expandComputeSubnetworkName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	networkProp, err := expandComputeSubnetworkNetwork(d.Get("network"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("network"); !isEmptyValue(reflect.ValueOf(networkProp)) && (ok || !reflect.DeepEqual(v, networkProp)) {
		obj["network"] = networkProp
	}
	fingerprintProp, err := expandComputeSubnetworkFingerprint(d.Get("fingerprint"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("fingerprint"); !isEmptyValue(reflect.ValueOf(fingerprintProp)) && (ok || !reflect.DeepEqual(v, fingerprintProp)) {
		obj["fingerprint"] = fingerprintProp
	}
	secondaryIpRangesProp, err := expandComputeSubnetworkSecondaryIpRange(d.Get("secondary_ip_range"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("secondary_ip_range"); ok || !reflect.DeepEqual(v, secondaryIpRangesProp) {
		obj["secondaryIpRanges"] = secondaryIpRangesProp
	}
	privateIpGoogleAccessProp, err := expandComputeSubnetworkPrivateIpGoogleAccess(d.Get("private_ip_google_access"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("private_ip_google_access"); !isEmptyValue(reflect.ValueOf(privateIpGoogleAccessProp)) && (ok || !reflect.DeepEqual(v, privateIpGoogleAccessProp)) {
		obj["privateIpGoogleAccess"] = privateIpGoogleAccessProp
	}
	regionProp, err := expandComputeSubnetworkRegion(d.Get("region"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("region"); !isEmptyValue(reflect.ValueOf(regionProp)) && (ok || !reflect.DeepEqual(v, regionProp)) {
		obj["region"] = regionProp
	}
	logConfigProp, err := expandComputeSubnetworkLogConfig(d.Get("log_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("log_config"); ok || !reflect.DeepEqual(v, logConfigProp) {
		obj["logConfig"] = logConfigProp
	}

	return obj, nil
}

func expandComputeSubnetworkDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeSubnetworkIpCidrRange(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeSubnetworkName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeSubnetworkNetwork(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseGlobalFieldValue("networks", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for network: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeSubnetworkFingerprint(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeSubnetworkSecondaryIpRange(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedRangeName, err := expandComputeSubnetworkSecondaryIpRangeRangeName(original["range_name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedRangeName); val.IsValid() && !isEmptyValue(val) {
			transformed["rangeName"] = transformedRangeName
		}

		transformedIpCidrRange, err := expandComputeSubnetworkSecondaryIpRangeIpCidrRange(original["ip_cidr_range"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedIpCidrRange); val.IsValid() && !isEmptyValue(val) {
			transformed["ipCidrRange"] = transformedIpCidrRange
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandComputeSubnetworkSecondaryIpRangeRangeName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeSubnetworkSecondaryIpRangeIpCidrRange(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeSubnetworkPrivateIpGoogleAccess(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeSubnetworkRegion(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseGlobalFieldValue("regions", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for region: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeSubnetworkLogConfig(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	transformed := make(map[string]interface{})
	if len(l) == 0 || l[0] == nil {
		purpose, ok := d.GetOkExists("purpose")

		if ok && purpose.(string) == "INTERNAL_HTTPS_LOAD_BALANCER" {
			// Subnetworks for L7ILB do not accept any values for logConfig
			return nil, nil
		}
		// send enable = false to ensure logging is disabled if there is no config
		transformed["enable"] = false
		return transformed, nil
	}

	raw := l[0]
	original := raw.(map[string]interface{})

	// The log_config block is specified, so logging should be enabled
	transformed["enable"] = true
	transformed["aggregationInterval"] = original["aggregation_interval"]
	transformed["flowSampling"] = original["flow_sampling"]
	transformed["metadata"] = original["metadata"]

	return transformed, nil
}
