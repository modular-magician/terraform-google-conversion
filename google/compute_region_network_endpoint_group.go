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
	"reflect"
)

func GetComputeRegionNetworkEndpointGroupCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	name, err := assetName(d, config, "//compute.googleapis.com/projects/{{project}}/regions/{{region}}/networkEndpointGroups/{{name}}")
	if err != nil {
		return []Asset{}, err
	}
	if obj, err := GetComputeRegionNetworkEndpointGroupApiObject(d, config); err == nil {
		return []Asset{{
			Name: name,
			Type: "compute.googleapis.com/RegionNetworkEndpointGroup",
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/compute/v1/rest",
				DiscoveryName:        "RegionNetworkEndpointGroup",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []Asset{}, err
	}
}

func GetComputeRegionNetworkEndpointGroupApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandComputeRegionNetworkEndpointGroupName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	descriptionProp, err := expandComputeRegionNetworkEndpointGroupDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	networkEndpointTypeProp, err := expandComputeRegionNetworkEndpointGroupNetworkEndpointType(d.Get("network_endpoint_type"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("network_endpoint_type"); !isEmptyValue(reflect.ValueOf(networkEndpointTypeProp)) && (ok || !reflect.DeepEqual(v, networkEndpointTypeProp)) {
		obj["networkEndpointType"] = networkEndpointTypeProp
	}
	cloudRunProp, err := expandComputeRegionNetworkEndpointGroupCloudRun(d.Get("cloud_run"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("cloud_run"); !isEmptyValue(reflect.ValueOf(cloudRunProp)) && (ok || !reflect.DeepEqual(v, cloudRunProp)) {
		obj["cloudRun"] = cloudRunProp
	}
	appEngineProp, err := expandComputeRegionNetworkEndpointGroupAppEngine(d.Get("app_engine"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("app_engine"); !isEmptyValue(reflect.ValueOf(appEngineProp)) && (ok || !reflect.DeepEqual(v, appEngineProp)) {
		obj["appEngine"] = appEngineProp
	}
	cloudFunctionProp, err := expandComputeRegionNetworkEndpointGroupCloudFunction(d.Get("cloud_function"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("cloud_function"); !isEmptyValue(reflect.ValueOf(cloudFunctionProp)) && (ok || !reflect.DeepEqual(v, cloudFunctionProp)) {
		obj["cloudFunction"] = cloudFunctionProp
	}
	regionProp, err := expandComputeRegionNetworkEndpointGroupRegion(d.Get("region"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("region"); !isEmptyValue(reflect.ValueOf(regionProp)) && (ok || !reflect.DeepEqual(v, regionProp)) {
		obj["region"] = regionProp
	}

	return obj, nil
}

func expandComputeRegionNetworkEndpointGroupName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionNetworkEndpointGroupDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionNetworkEndpointGroupNetworkEndpointType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionNetworkEndpointGroupCloudRun(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedService, err := expandComputeRegionNetworkEndpointGroupCloudRunService(original["service"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedService); val.IsValid() && !isEmptyValue(val) {
		transformed["service"] = transformedService
	}

	transformedTag, err := expandComputeRegionNetworkEndpointGroupCloudRunTag(original["tag"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTag); val.IsValid() && !isEmptyValue(val) {
		transformed["tag"] = transformedTag
	}

	transformedUrlMask, err := expandComputeRegionNetworkEndpointGroupCloudRunUrlMask(original["url_mask"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUrlMask); val.IsValid() && !isEmptyValue(val) {
		transformed["urlMask"] = transformedUrlMask
	}

	return transformed, nil
}

func expandComputeRegionNetworkEndpointGroupCloudRunService(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionNetworkEndpointGroupCloudRunTag(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionNetworkEndpointGroupCloudRunUrlMask(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionNetworkEndpointGroupAppEngine(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 {
		return nil, nil
	}

	if l[0] == nil {
		transformed := make(map[string]interface{})
		return transformed, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedService, err := expandComputeRegionNetworkEndpointGroupAppEngineService(original["service"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedService); val.IsValid() && !isEmptyValue(val) {
		transformed["service"] = transformedService
	}

	transformedVersion, err := expandComputeRegionNetworkEndpointGroupAppEngineVersion(original["version"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedVersion); val.IsValid() && !isEmptyValue(val) {
		transformed["version"] = transformedVersion
	}

	transformedUrlMask, err := expandComputeRegionNetworkEndpointGroupAppEngineUrlMask(original["url_mask"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUrlMask); val.IsValid() && !isEmptyValue(val) {
		transformed["urlMask"] = transformedUrlMask
	}

	return transformed, nil
}

func expandComputeRegionNetworkEndpointGroupAppEngineService(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionNetworkEndpointGroupAppEngineVersion(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionNetworkEndpointGroupAppEngineUrlMask(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionNetworkEndpointGroupCloudFunction(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedFunction, err := expandComputeRegionNetworkEndpointGroupCloudFunctionFunction(original["function"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedFunction); val.IsValid() && !isEmptyValue(val) {
		transformed["function"] = transformedFunction
	}

	transformedUrlMask, err := expandComputeRegionNetworkEndpointGroupCloudFunctionUrlMask(original["url_mask"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUrlMask); val.IsValid() && !isEmptyValue(val) {
		transformed["urlMask"] = transformedUrlMask
	}

	return transformed, nil
}

func expandComputeRegionNetworkEndpointGroupCloudFunctionFunction(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionNetworkEndpointGroupCloudFunctionUrlMask(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionNetworkEndpointGroupRegion(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseGlobalFieldValue("regions", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for region: %s", err)
	}
	return f.RelativeLink(), nil
}
