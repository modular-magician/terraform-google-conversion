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

package google

import (
	"bytes"
	"context"
	"fmt"
	"reflect"
	"sort"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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

	return hashcode(buf.String())
}

func compareCaseInsensitive(k, old, new string, d *schema.ResourceData) bool {
	return strings.ToLower(old) == strings.ToLower(new)
}

func diffSuppressEnableLogging(k, old, new string, d *schema.ResourceData) bool {
	if k == "log_config.#" {
		if new == "0" && d.Get("enable_logging").(bool) {
			return true
		}
	}

	return false
}

func resourceComputeFirewallEnableLoggingCustomizeDiff(_ context.Context, diff *schema.ResourceDiff, v interface{}) error {
	enableLogging, enableExists := diff.GetOkExists("enable_logging")
	if !enableExists {
		return nil
	}

	logConfigExists := diff.Get("log_config.#").(int) != 0
	if logConfigExists && enableLogging == false {
		return fmt.Errorf("log_config cannot be defined when enable_logging is false")
	}

	return nil
}

func diffSuppressSourceRanges(k, old, new string, d *schema.ResourceData) bool {
	if k == "source_ranges.#" {
		if old == "1" && new == "0" {
			// Allow diffing on the individual element if we are going from 1 -> 0
			// this allows for diff suppress on ["0.0.0.0/0"] -> []
			return true
		}
		return old == new
	}
	kLength := "source_ranges.#"
	oldLength, newLength := d.GetChange(kLength)
	oldInt, ok := oldLength.(int)

	if !ok {
		return false
	}

	newInt, ok := newLength.(int)
	if !ok {
		return false
	}

	// Diff suppress only should suppress removing the default range
	// This should probably be newInt == 0, but due to Terraform core internals
	// (bug?) values found via GetChange may not have the correct new value
	// in some circumstances
	if oldInt == 1 && newInt == 1 {
		if old == "0.0.0.0/0" && new == "" {
			return true
		}
	}
	return old == new
}

const ComputeFirewallAssetType string = "compute.googleapis.com/Firewall"

func resourceConverterComputeFirewall() ResourceConverter {
	return ResourceConverter{
		AssetType: ComputeFirewallAssetType,
		Convert:   GetComputeFirewallCaiObject,
	}
}

func GetComputeFirewallCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	name, err := assetName(d, config, "//compute.googleapis.com/projects/{{project}}/global/firewalls/{{name}}")
	if err != nil {
		return []Asset{}, err
	}
	if obj, err := GetComputeFirewallApiObject(d, config); err == nil {
		return []Asset{{
			Name: name,
			Type: ComputeFirewallAssetType,
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/compute/v1/rest",
				DiscoveryName:        "Firewall",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []Asset{}, err
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
	logConfigProp, err := expandComputeFirewallLogConfig(d.Get("log_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("log_config"); ok || !reflect.DeepEqual(v, logConfigProp) {
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
	l := v.([]interface{})
	transformed := make(map[string]interface{})

	if len(l) == 0 || l[0] == nil {
		// send enable = enable_logging value to ensure correct logging status if there is no config
		transformed["enable"] = d.Get("enable_logging").(bool)
		return transformed, nil
	}

	raw := l[0]
	original := raw.(map[string]interface{})

	// The log_config block is specified, so logging should be enabled
	transformed["enable"] = true
	transformed["metadata"] = original["metadata"]

	return transformed, nil
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
