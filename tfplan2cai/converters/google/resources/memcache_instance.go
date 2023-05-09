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
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgresource"
	transport_tpg "github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/transport"
)

const MemcacheInstanceAssetType string = "memcache.googleapis.com/Instance"

func resourceConverterMemcacheInstance() ResourceConverter {
	return ResourceConverter{
		AssetType: MemcacheInstanceAssetType,
		Convert:   GetMemcacheInstanceCaiObject,
	}
}

func GetMemcacheInstanceCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]Asset, error) {
	name, err := assetName(d, config, "//memcache.googleapis.com/projects/{{project}}/locations/{{region}}/instances/{{name}}")
	if err != nil {
		return []Asset{}, err
	}
	if obj, err := GetMemcacheInstanceApiObject(d, config); err == nil {
		return []Asset{{
			Name: name,
			Type: MemcacheInstanceAssetType,
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/memcache/v1/rest",
				DiscoveryName:        "Instance",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []Asset{}, err
	}
}

func GetMemcacheInstanceApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	displayNameProp, err := expandMemcacheInstanceDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	labelsProp, err := expandMemcacheInstanceLabels(d.Get("labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	zonesProp, err := expandMemcacheInstanceZones(d.Get("zones"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("zones"); !tpgresource.IsEmptyValue(reflect.ValueOf(zonesProp)) && (ok || !reflect.DeepEqual(v, zonesProp)) {
		obj["zones"] = zonesProp
	}
	authorizedNetworkProp, err := expandMemcacheInstanceAuthorizedNetwork(d.Get("authorized_network"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("authorized_network"); !tpgresource.IsEmptyValue(reflect.ValueOf(authorizedNetworkProp)) && (ok || !reflect.DeepEqual(v, authorizedNetworkProp)) {
		obj["authorizedNetwork"] = authorizedNetworkProp
	}
	nodeCountProp, err := expandMemcacheInstanceNodeCount(d.Get("node_count"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("node_count"); !tpgresource.IsEmptyValue(reflect.ValueOf(nodeCountProp)) && (ok || !reflect.DeepEqual(v, nodeCountProp)) {
		obj["nodeCount"] = nodeCountProp
	}
	memcacheVersionProp, err := expandMemcacheInstanceMemcacheVersion(d.Get("memcache_version"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("memcache_version"); !tpgresource.IsEmptyValue(reflect.ValueOf(memcacheVersionProp)) && (ok || !reflect.DeepEqual(v, memcacheVersionProp)) {
		obj["memcacheVersion"] = memcacheVersionProp
	}
	nodeConfigProp, err := expandMemcacheInstanceNodeConfig(d.Get("node_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("node_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(nodeConfigProp)) && (ok || !reflect.DeepEqual(v, nodeConfigProp)) {
		obj["nodeConfig"] = nodeConfigProp
	}
	parametersProp, err := expandMemcacheInstanceMemcacheParameters(d.Get("memcache_parameters"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("memcache_parameters"); !tpgresource.IsEmptyValue(reflect.ValueOf(parametersProp)) && (ok || !reflect.DeepEqual(v, parametersProp)) {
		obj["parameters"] = parametersProp
	}
	maintenancePolicyProp, err := expandMemcacheInstanceMaintenancePolicy(d.Get("maintenance_policy"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("maintenance_policy"); !tpgresource.IsEmptyValue(reflect.ValueOf(maintenancePolicyProp)) && (ok || !reflect.DeepEqual(v, maintenancePolicyProp)) {
		obj["maintenancePolicy"] = maintenancePolicyProp
	}

	return obj, nil
}

func expandMemcacheInstanceDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMemcacheInstanceLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandMemcacheInstanceZones(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	return v, nil
}

func expandMemcacheInstanceAuthorizedNetwork(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMemcacheInstanceNodeCount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMemcacheInstanceMemcacheVersion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMemcacheInstanceNodeConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedCpuCount, err := expandMemcacheInstanceNodeConfigCpuCount(original["cpu_count"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCpuCount); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["cpuCount"] = transformedCpuCount
	}

	transformedMemorySizeMb, err := expandMemcacheInstanceNodeConfigMemorySizeMb(original["memory_size_mb"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMemorySizeMb); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["memorySizeMb"] = transformedMemorySizeMb
	}

	return transformed, nil
}

func expandMemcacheInstanceNodeConfigCpuCount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMemcacheInstanceNodeConfigMemorySizeMb(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMemcacheInstanceMemcacheParameters(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedId, err := expandMemcacheInstanceMemcacheParametersId(original["id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["id"] = transformedId
	}

	transformedParams, err := expandMemcacheInstanceMemcacheParametersParams(original["params"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedParams); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["params"] = transformedParams
	}

	return transformed, nil
}

func expandMemcacheInstanceMemcacheParametersId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMemcacheInstanceMemcacheParametersParams(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandMemcacheInstanceMaintenancePolicy(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedCreateTime, err := expandMemcacheInstanceMaintenancePolicyCreateTime(original["create_time"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCreateTime); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["createTime"] = transformedCreateTime
	}

	transformedUpdateTime, err := expandMemcacheInstanceMaintenancePolicyUpdateTime(original["update_time"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUpdateTime); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["updateTime"] = transformedUpdateTime
	}

	transformedDescription, err := expandMemcacheInstanceMaintenancePolicyDescription(original["description"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDescription); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["description"] = transformedDescription
	}

	transformedWeeklyMaintenanceWindow, err := expandMemcacheInstanceMaintenancePolicyWeeklyMaintenanceWindow(original["weekly_maintenance_window"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedWeeklyMaintenanceWindow); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["weeklyMaintenanceWindow"] = transformedWeeklyMaintenanceWindow
	}

	return transformed, nil
}

func expandMemcacheInstanceMaintenancePolicyCreateTime(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMemcacheInstanceMaintenancePolicyUpdateTime(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMemcacheInstanceMaintenancePolicyDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMemcacheInstanceMaintenancePolicyWeeklyMaintenanceWindow(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedDay, err := expandMemcacheInstanceMaintenancePolicyWeeklyMaintenanceWindowDay(original["day"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDay); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["day"] = transformedDay
		}

		transformedDuration, err := expandMemcacheInstanceMaintenancePolicyWeeklyMaintenanceWindowDuration(original["duration"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDuration); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["duration"] = transformedDuration
		}

		transformedStartTime, err := expandMemcacheInstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime(original["start_time"], d, config)
		if err != nil {
			return nil, err
		} else {
			transformed["startTime"] = transformedStartTime
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandMemcacheInstanceMaintenancePolicyWeeklyMaintenanceWindowDay(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMemcacheInstanceMaintenancePolicyWeeklyMaintenanceWindowDuration(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMemcacheInstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
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

	transformedHours, err := expandMemcacheInstanceMaintenancePolicyWeeklyMaintenanceWindowStartTimeHours(original["hours"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedHours); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["hours"] = transformedHours
	}

	transformedMinutes, err := expandMemcacheInstanceMaintenancePolicyWeeklyMaintenanceWindowStartTimeMinutes(original["minutes"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMinutes); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["minutes"] = transformedMinutes
	}

	transformedSeconds, err := expandMemcacheInstanceMaintenancePolicyWeeklyMaintenanceWindowStartTimeSeconds(original["seconds"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSeconds); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["seconds"] = transformedSeconds
	}

	transformedNanos, err := expandMemcacheInstanceMaintenancePolicyWeeklyMaintenanceWindowStartTimeNanos(original["nanos"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNanos); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["nanos"] = transformedNanos
	}

	return transformed, nil
}

func expandMemcacheInstanceMaintenancePolicyWeeklyMaintenanceWindowStartTimeHours(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMemcacheInstanceMaintenancePolicyWeeklyMaintenanceWindowStartTimeMinutes(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMemcacheInstanceMaintenancePolicyWeeklyMaintenanceWindowStartTimeSeconds(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMemcacheInstanceMaintenancePolicyWeeklyMaintenanceWindowStartTimeNanos(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
