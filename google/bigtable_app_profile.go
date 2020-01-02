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

func GetBigtableAppProfileCaiObject(d TerraformResourceData, config *Config) (Asset, error) {
	name, err := assetName(d, config, "//bigtable.googleapis.com/projects/{{project}}/instances/{{instance}}/appProfiles/{{app_profile_id}}")
	if err != nil {
		return Asset{}, err
	}
	if obj, err := GetBigtableAppProfileApiObject(d, config); err == nil {
		return Asset{
			Name: name,
			Type: "bigtable.googleapis.com/AppProfile",
			Resource: &AssetResource{
				Version:              "v2",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/bigtable/v2/rest",
				DiscoveryName:        "AppProfile",
				Data:                 obj,
			},
		}, nil
	} else {
		return Asset{}, err
	}
}

func GetBigtableAppProfileApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	descriptionProp, err := expandBigtableAppProfileDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	multiClusterRoutingUseAnyProp, err := expandBigtableAppProfileMultiClusterRoutingUseAny(d.Get("multi_cluster_routing_use_any"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("multi_cluster_routing_use_any"); !isEmptyValue(reflect.ValueOf(multiClusterRoutingUseAnyProp)) && (ok || !reflect.DeepEqual(v, multiClusterRoutingUseAnyProp)) {
		obj["multiClusterRoutingUseAny"] = multiClusterRoutingUseAnyProp
	}
	singleClusterRoutingProp, err := expandBigtableAppProfileSingleClusterRouting(d.Get("single_cluster_routing"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("single_cluster_routing"); !isEmptyValue(reflect.ValueOf(singleClusterRoutingProp)) && (ok || !reflect.DeepEqual(v, singleClusterRoutingProp)) {
		obj["singleClusterRouting"] = singleClusterRoutingProp
	}

	return obj, nil
}

func expandBigtableAppProfileDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBigtableAppProfileMultiClusterRoutingUseAny(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	if v == nil || !v.(bool) {
		return nil, nil
	}

	return map[string]interface{}{}, nil
}

func expandBigtableAppProfileSingleClusterRouting(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedClusterId, err := expandBigtableAppProfileSingleClusterRoutingClusterId(original["cluster_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedClusterId); val.IsValid() && !isEmptyValue(val) {
		transformed["clusterId"] = transformedClusterId
	}

	transformedAllowTransactionalWrites, err := expandBigtableAppProfileSingleClusterRoutingAllowTransactionalWrites(original["allow_transactional_writes"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAllowTransactionalWrites); val.IsValid() && !isEmptyValue(val) {
		transformed["allowTransactionalWrites"] = transformedAllowTransactionalWrites
	}

	return transformed, nil
}

func expandBigtableAppProfileSingleClusterRoutingClusterId(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBigtableAppProfileSingleClusterRoutingAllowTransactionalWrites(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
