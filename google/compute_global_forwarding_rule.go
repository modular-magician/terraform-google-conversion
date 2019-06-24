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

func GetComputeGlobalForwardingRuleCaiObject(d TerraformResourceData, config *Config) (Asset, error) {
	name, err := assetName(d, config, "//compute.googleapis.com/projects/{{project}}/global/forwardingRules/{{name}}")
	if err != nil {
		return Asset{}, err
	}
	if obj, err := GetComputeGlobalForwardingRuleApiObject(d, config); err == nil {
		return Asset{
			Name: name,
			Type: "compute.googleapis.com/GlobalForwardingRule",
			Resource: &AssetResource{
				Version:              "ga",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/compute/ga/rest",
				DiscoveryName:        "GlobalForwardingRule",
				Data:                 obj,
			},
		}, nil
	} else {
		return Asset{}, err
	}
}

func GetComputeGlobalForwardingRuleApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	descriptionProp, err := expandComputeGlobalForwardingRuleDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	IPAddressProp, err := expandComputeGlobalForwardingRuleIPAddress(d.Get("ip_address"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("ip_address"); !isEmptyValue(reflect.ValueOf(IPAddressProp)) && (ok || !reflect.DeepEqual(v, IPAddressProp)) {
		obj["IPAddress"] = IPAddressProp
	}
	IPProtocolProp, err := expandComputeGlobalForwardingRuleIPProtocol(d.Get("ip_protocol"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("ip_protocol"); !isEmptyValue(reflect.ValueOf(IPProtocolProp)) && (ok || !reflect.DeepEqual(v, IPProtocolProp)) {
		obj["IPProtocol"] = IPProtocolProp
	}
	ipVersionProp, err := expandComputeGlobalForwardingRuleIpVersion(d.Get("ip_version"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("ip_version"); !isEmptyValue(reflect.ValueOf(ipVersionProp)) && (ok || !reflect.DeepEqual(v, ipVersionProp)) {
		obj["ipVersion"] = ipVersionProp
	}
	loadBalancingSchemeProp, err := expandComputeGlobalForwardingRuleLoadBalancingScheme(d.Get("load_balancing_scheme"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("load_balancing_scheme"); !isEmptyValue(reflect.ValueOf(loadBalancingSchemeProp)) && (ok || !reflect.DeepEqual(v, loadBalancingSchemeProp)) {
		obj["loadBalancingScheme"] = loadBalancingSchemeProp
	}
	nameProp, err := expandComputeGlobalForwardingRuleName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	portRangeProp, err := expandComputeGlobalForwardingRulePortRange(d.Get("port_range"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("port_range"); !isEmptyValue(reflect.ValueOf(portRangeProp)) && (ok || !reflect.DeepEqual(v, portRangeProp)) {
		obj["portRange"] = portRangeProp
	}
	targetProp, err := expandComputeGlobalForwardingRuleTarget(d.Get("target"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("target"); !isEmptyValue(reflect.ValueOf(targetProp)) && (ok || !reflect.DeepEqual(v, targetProp)) {
		obj["target"] = targetProp
	}

	return obj, nil
}

func expandComputeGlobalForwardingRuleDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeGlobalForwardingRuleIPAddress(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeGlobalForwardingRuleIPProtocol(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeGlobalForwardingRuleIpVersion(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeGlobalForwardingRuleLoadBalancingScheme(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeGlobalForwardingRuleName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeGlobalForwardingRulePortRange(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeGlobalForwardingRuleTarget(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
