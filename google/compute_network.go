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

func GetComputeNetworkCaiObject(d TerraformResourceData, config *Config) (Asset, error) {
	name, err := assetName(d, config, "//compute.googleapis.com/projects/{{project}}/global/networks/{{name}}")
	if err != nil {
		return Asset{}, err
	}
	if obj, err := GetComputeNetworkApiObject(d, config); err == nil {
		return Asset{
			Name: name,
			Type: "compute.googleapis.com/Network",
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/compute/v1/rest",
				DiscoveryName:        "Network",
				Data:                 obj,
			},
		}, nil
	} else {
		return Asset{}, err
	}
}

func GetComputeNetworkApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	descriptionProp, err := expandComputeNetworkDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	IPv4RangeProp, err := expandComputeNetworkIpv4_range(d.Get("ipv4_range"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("ipv4_range"); !isEmptyValue(reflect.ValueOf(IPv4RangeProp)) && (ok || !reflect.DeepEqual(v, IPv4RangeProp)) {
		obj["IPv4Range"] = IPv4RangeProp
	}
	nameProp, err := expandComputeNetworkName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	autoCreateSubnetworksProp, err := expandComputeNetworkAutoCreateSubnetworks(d.Get("auto_create_subnetworks"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("auto_create_subnetworks"); !isEmptyValue(reflect.ValueOf(autoCreateSubnetworksProp)) && (ok || !reflect.DeepEqual(v, autoCreateSubnetworksProp)) {
		obj["autoCreateSubnetworks"] = autoCreateSubnetworksProp
	}
	routingConfigProp, err := expandComputeNetworkRoutingConfig(d, config)
	if err != nil {
		return nil, err
	} else if !isEmptyValue(reflect.ValueOf(routingConfigProp)) {
		obj["routingConfig"] = routingConfigProp
	}

	return resourceComputeNetworkEncoder(d, config, obj)
}

func resourceComputeNetworkEncoder(d TerraformResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	if _, ok := d.GetOk("ipv4_range"); !ok {
		obj["autoCreateSubnetworks"] = d.Get("auto_create_subnetworks")
	}

	return obj, nil
}

func expandComputeNetworkDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkIpv4_range(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkAutoCreateSubnetworks(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkRoutingConfig(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	transformed := make(map[string]interface{})
	transformedRoutingMode, err := expandComputeNetworkRoutingConfigRoutingMode(d.Get("routing_mode"), d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRoutingMode); val.IsValid() && !isEmptyValue(val) {
		transformed["routingMode"] = transformedRoutingMode
	}

	return transformed, nil
}

func expandComputeNetworkRoutingConfigRoutingMode(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
