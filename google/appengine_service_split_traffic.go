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

func GetAppEngineServiceSplitTrafficCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	name, err := assetName(d, config, "//appengine.googleapis.com/apps/{{project}}/services/{{service}}")
	if err != nil {
		return []Asset{}, err
	}
	if obj, err := GetAppEngineServiceSplitTrafficApiObject(d, config); err == nil {
		return []Asset{{
			Name: name,
			Type: "appengine.googleapis.com/ServiceSplitTraffic",
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/appengine/v1/rest",
				DiscoveryName:        "ServiceSplitTraffic",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []Asset{}, err
	}
}

func GetAppEngineServiceSplitTrafficApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	idProp, err := expandAppEngineServiceSplitTrafficService(d.Get("service"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("service"); !isEmptyValue(reflect.ValueOf(idProp)) && (ok || !reflect.DeepEqual(v, idProp)) {
		obj["id"] = idProp
	}
	splitProp, err := expandAppEngineServiceSplitTrafficSplit(d.Get("split"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("split"); !isEmptyValue(reflect.ValueOf(splitProp)) && (ok || !reflect.DeepEqual(v, splitProp)) {
		obj["split"] = splitProp
	}

	return obj, nil
}

func expandAppEngineServiceSplitTrafficService(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineServiceSplitTrafficSplit(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedShardBy, err := expandAppEngineServiceSplitTrafficSplitShardBy(original["shard_by"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedShardBy); val.IsValid() && !isEmptyValue(val) {
		transformed["shardBy"] = transformedShardBy
	}

	transformedAllocations, err := expandAppEngineServiceSplitTrafficSplitAllocations(original["allocations"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAllocations); val.IsValid() && !isEmptyValue(val) {
		transformed["allocations"] = transformedAllocations
	}

	return transformed, nil
}

func expandAppEngineServiceSplitTrafficSplitShardBy(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineServiceSplitTrafficSplitAllocations(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
