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

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgresource"
	transport_tpg "github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/transport"
)

const EdgecontainerClusterAssetType string = "edgecontainer.googleapis.com/Cluster"

func resourceConverterEdgecontainerCluster() ResourceConverter {
	return ResourceConverter{
		AssetType: EdgecontainerClusterAssetType,
		Convert:   GetEdgecontainerClusterCaiObject,
	}
}

func GetEdgecontainerClusterCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]Asset, error) {
	name, err := assetName(d, config, "//edgecontainer.googleapis.com/projects/{{project}}/locations/{{location}}/clusters/{{name}}")
	if err != nil {
		return []Asset{}, err
	}
	if obj, err := GetEdgecontainerClusterApiObject(d, config); err == nil {
		return []Asset{{
			Name: name,
			Type: EdgecontainerClusterAssetType,
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/edgecontainer/v1/rest",
				DiscoveryName:        "Cluster",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []Asset{}, err
	}
}

func GetEdgecontainerClusterApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	labelsProp, err := expandEdgecontainerClusterLabels(d.Get("labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	fleetProp, err := expandEdgecontainerClusterFleet(d.Get("fleet"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("fleet"); !tpgresource.IsEmptyValue(reflect.ValueOf(fleetProp)) && (ok || !reflect.DeepEqual(v, fleetProp)) {
		obj["fleet"] = fleetProp
	}
	networkingProp, err := expandEdgecontainerClusterNetworking(d.Get("networking"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("networking"); !tpgresource.IsEmptyValue(reflect.ValueOf(networkingProp)) && (ok || !reflect.DeepEqual(v, networkingProp)) {
		obj["networking"] = networkingProp
	}
	authorizationProp, err := expandEdgecontainerClusterAuthorization(d.Get("authorization"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("authorization"); !tpgresource.IsEmptyValue(reflect.ValueOf(authorizationProp)) && (ok || !reflect.DeepEqual(v, authorizationProp)) {
		obj["authorization"] = authorizationProp
	}
	defaultMaxPodsPerNodeProp, err := expandEdgecontainerClusterDefaultMaxPodsPerNode(d.Get("default_max_pods_per_node"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("default_max_pods_per_node"); !tpgresource.IsEmptyValue(reflect.ValueOf(defaultMaxPodsPerNodeProp)) && (ok || !reflect.DeepEqual(v, defaultMaxPodsPerNodeProp)) {
		obj["defaultMaxPodsPerNode"] = defaultMaxPodsPerNodeProp
	}
	maintenancePolicyProp, err := expandEdgecontainerClusterMaintenancePolicy(d.Get("maintenance_policy"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("maintenance_policy"); !tpgresource.IsEmptyValue(reflect.ValueOf(maintenancePolicyProp)) && (ok || !reflect.DeepEqual(v, maintenancePolicyProp)) {
		obj["maintenancePolicy"] = maintenancePolicyProp
	}

	return obj, nil
}

func expandEdgecontainerClusterLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandEdgecontainerClusterFleet(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedProject, err := expandEdgecontainerClusterFleetProject(original["project"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedProject); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["project"] = transformedProject
	}

	return transformed, nil
}

func expandEdgecontainerClusterFleetProject(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandEdgecontainerClusterNetworking(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedClusterIpv4CidrBlocks, err := expandEdgecontainerClusterNetworkingClusterIpv4CidrBlocks(original["cluster_ipv4_cidr_blocks"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedClusterIpv4CidrBlocks); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["clusterIpv4CidrBlocks"] = transformedClusterIpv4CidrBlocks
	}

	transformedServicesIpv4CidrBlocks, err := expandEdgecontainerClusterNetworkingServicesIpv4CidrBlocks(original["services_ipv4_cidr_blocks"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedServicesIpv4CidrBlocks); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["servicesIpv4CidrBlocks"] = transformedServicesIpv4CidrBlocks
	}

	return transformed, nil
}

func expandEdgecontainerClusterNetworkingClusterIpv4CidrBlocks(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandEdgecontainerClusterNetworkingServicesIpv4CidrBlocks(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandEdgecontainerClusterAuthorization(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedAdminUsers, err := expandEdgecontainerClusterAuthorizationAdminUsers(original["admin_users"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAdminUsers); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["adminUsers"] = transformedAdminUsers
	}

	return transformed, nil
}

func expandEdgecontainerClusterAuthorizationAdminUsers(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedUsername, err := expandEdgecontainerClusterAuthorizationAdminUsersUsername(original["username"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUsername); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["username"] = transformedUsername
	}

	return transformed, nil
}

func expandEdgecontainerClusterAuthorizationAdminUsersUsername(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandEdgecontainerClusterDefaultMaxPodsPerNode(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandEdgecontainerClusterMaintenancePolicy(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedWindow, err := expandEdgecontainerClusterMaintenancePolicyWindow(original["window"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedWindow); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["window"] = transformedWindow
	}

	return transformed, nil
}

func expandEdgecontainerClusterMaintenancePolicyWindow(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedRecurringWindow, err := expandEdgecontainerClusterMaintenancePolicyWindowRecurringWindow(original["recurring_window"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRecurringWindow); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["recurringWindow"] = transformedRecurringWindow
	}

	return transformed, nil
}

func expandEdgecontainerClusterMaintenancePolicyWindowRecurringWindow(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedWindow, err := expandEdgecontainerClusterMaintenancePolicyWindowRecurringWindowWindow(original["window"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedWindow); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["window"] = transformedWindow
	}

	transformedRecurrence, err := expandEdgecontainerClusterMaintenancePolicyWindowRecurringWindowRecurrence(original["recurrence"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRecurrence); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["recurrence"] = transformedRecurrence
	}

	return transformed, nil
}

func expandEdgecontainerClusterMaintenancePolicyWindowRecurringWindowWindow(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedStartTime, err := expandEdgecontainerClusterMaintenancePolicyWindowRecurringWindowWindowStartTime(original["start_time"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedStartTime); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["startTime"] = transformedStartTime
	}

	transformedEndTime, err := expandEdgecontainerClusterMaintenancePolicyWindowRecurringWindowWindowEndTime(original["end_time"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEndTime); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["endTime"] = transformedEndTime
	}

	return transformed, nil
}

func expandEdgecontainerClusterMaintenancePolicyWindowRecurringWindowWindowStartTime(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandEdgecontainerClusterMaintenancePolicyWindowRecurringWindowWindowEndTime(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandEdgecontainerClusterMaintenancePolicyWindowRecurringWindowRecurrence(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
