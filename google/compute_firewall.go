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
	"bytes"
	"fmt"
	"reflect"
	"sort"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceComputeFirewallRuleHash(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	buf.WriteString(fmt.Sprintf("%s-", strings.ToLower(m["protocol"].(string))))

	// We need to make sure to sort the strings below so that we always
	// generate the same hash code no matter what is in the set.
	if v, ok := m["ports"]; ok && v != nil {
		s := convertStringArr(v.([]interface{}))
		sort.Strings(s)

		for _, v := range s {
			buf.WriteString(fmt.Sprintf("%s-", v))
		}
	}

	return hashcode.String(buf.String())
}

func compareCaseInsensitive(k, old, new string, d *schema.ResourceData) bool {
	return strings.ToLower(old) == strings.ToLower(new)
}

func GetComputeFirewallCaiObject(d TerraformResourceData, config *Config) (Asset, error) {
	name, err := assetName(d, config, "//compute.googleapis.com/projects/{{project}}/global/firewalls/{{name}}")
	if err != nil {
		return Asset{}, err
	}
	if obj, err := GetComputeFirewallApiObject(d, config); err == nil {
		return Asset{
			Name: name,
			Type: "compute.googleapis.com/Firewall",
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/compute/v1/rest",
				DiscoveryName:        "Firewall",
				Data:                 obj,
			},
		}, nil
	} else {
		return Asset{}, err
	}
}

func GetComputeFirewallApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	allowedProp, err := expandComputeFirewallAllow(d.Get("allow"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("allow"); !isEmptyValue(reflect.ValueOf(allowedProp)) && (ok || !reflect.DeepEqual(v, allowedProp)) {
		obj["allowed"] = allowedProp
	}
	deniedProp, err := expandComputeFirewallDeny(d.Get("deny"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("deny"); !isEmptyValue(reflect.ValueOf(deniedProp)) && (ok || !reflect.DeepEqual(v, deniedProp)) {
		obj["denied"] = deniedProp
	}
	descriptionProp, err := expandComputeFirewallDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	destinationRangesProp, err := expandComputeFirewallDestinationRanges(d.Get("destination_ranges"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("destination_ranges"); !isEmptyValue(reflect.ValueOf(destinationRangesProp)) && (ok || !reflect.DeepEqual(v, destinationRangesProp)) {
		obj["destinationRanges"] = destinationRangesProp
	}
	directionProp, err := expandComputeFirewallDirection(d.Get("direction"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("direction"); !isEmptyValue(reflect.ValueOf(directionProp)) && (ok || !reflect.DeepEqual(v, directionProp)) {
		obj["direction"] = directionProp
	}
	disabledProp, err := expandComputeFirewallDisabled(d.Get("disabled"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("disabled"); ok || !reflect.DeepEqual(v, disabledProp) {
		obj["disabled"] = disabledProp
	}
	logConfigProp, err := expandComputeFirewallLogConfig(nil, d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("log_config"); !isEmptyValue(reflect.ValueOf(logConfigProp)) && (ok || !reflect.DeepEqual(v, logConfigProp)) {
		obj["logConfig"] = logConfigProp
	}
	nameProp, err := expandComputeFirewallName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	networkProp, err := expandComputeFirewallNetwork(d.Get("network"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("network"); !isEmptyValue(reflect.ValueOf(networkProp)) && (ok || !reflect.DeepEqual(v, networkProp)) {
		obj["network"] = networkProp
	}
	priorityProp, err := expandComputeFirewallPriority(d.Get("priority"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("priority"); ok || !reflect.DeepEqual(v, priorityProp) {
		obj["priority"] = priorityProp
	}
	sourceRangesProp, err := expandComputeFirewallSourceRanges(d.Get("source_ranges"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("source_ranges"); !isEmptyValue(reflect.ValueOf(sourceRangesProp)) && (ok || !reflect.DeepEqual(v, sourceRangesProp)) {
		obj["sourceRanges"] = sourceRangesProp
	}
	sourceServiceAccountsProp, err := expandComputeFirewallSourceServiceAccounts(d.Get("source_service_accounts"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("source_service_accounts"); !isEmptyValue(reflect.ValueOf(sourceServiceAccountsProp)) && (ok || !reflect.DeepEqual(v, sourceServiceAccountsProp)) {
		obj["sourceServiceAccounts"] = sourceServiceAccountsProp
	}
	sourceTagsProp, err := expandComputeFirewallSourceTags(d.Get("source_tags"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("source_tags"); !isEmptyValue(reflect.ValueOf(sourceTagsProp)) && (ok || !reflect.DeepEqual(v, sourceTagsProp)) {
		obj["sourceTags"] = sourceTagsProp
	}
	targetServiceAccountsProp, err := expandComputeFirewallTargetServiceAccounts(d.Get("target_service_accounts"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("target_service_accounts"); !isEmptyValue(reflect.ValueOf(targetServiceAccountsProp)) && (ok || !reflect.DeepEqual(v, targetServiceAccountsProp)) {
		obj["targetServiceAccounts"] = targetServiceAccountsProp
	}
	targetTagsProp, err := expandComputeFirewallTargetTags(d.Get("target_tags"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("target_tags"); !isEmptyValue(reflect.ValueOf(targetTagsProp)) && (ok || !reflect.DeepEqual(v, targetTagsProp)) {
		obj["targetTags"] = targetTagsProp
	}

	return obj, nil
}

func expandComputeFirewallAllow(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedProtocol, err := expandComputeFirewallAllowProtocol(original["protocol"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedProtocol); val.IsValid() && !isEmptyValue(val) {
			transformed["IPProtocol"] = transformedProtocol
		}

		transformedPorts, err := expandComputeFirewallAllowPorts(original["ports"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedPorts); val.IsValid() && !isEmptyValue(val) {
			transformed["ports"] = transformedPorts
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandComputeFirewallAllowProtocol(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeFirewallAllowPorts(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeFirewallDeny(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedProtocol, err := expandComputeFirewallDenyProtocol(original["protocol"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedProtocol); val.IsValid() && !isEmptyValue(val) {
			transformed["IPProtocol"] = transformedProtocol
		}

		transformedPorts, err := expandComputeFirewallDenyPorts(original["ports"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedPorts); val.IsValid() && !isEmptyValue(val) {
			transformed["ports"] = transformedPorts
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandComputeFirewallDenyProtocol(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeFirewallDenyPorts(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeFirewallDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeFirewallDestinationRanges(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	return v, nil
}

func expandComputeFirewallDirection(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeFirewallDisabled(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeFirewallLogConfig(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	transformed := make(map[string]interface{})
	transformedEnableLogging, err := expandComputeFirewallLogConfigEnableLogging(d.Get("enable_logging"), d, config)
	if err != nil {
		return nil, err
	} else {
		transformed["enable"] = transformedEnableLogging
	}

	return transformed, nil
}

func expandComputeFirewallLogConfigEnableLogging(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeFirewallName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeFirewallNetwork(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseGlobalFieldValue("networks", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for network: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeFirewallPriority(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeFirewallSourceRanges(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	return v, nil
}

func expandComputeFirewallSourceServiceAccounts(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	return v, nil
}

func expandComputeFirewallSourceTags(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	return v, nil
}

func expandComputeFirewallTargetServiceAccounts(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	return v, nil
}

func expandComputeFirewallTargetTags(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	return v, nil
}
