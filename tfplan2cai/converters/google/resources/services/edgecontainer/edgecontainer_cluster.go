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

package edgecontainer

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const EdgecontainerClusterAssetType string = "edgecontainer.googleapis.com/Cluster"

func ResourceConverterEdgecontainerCluster() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: EdgecontainerClusterAssetType,
		Convert:   GetEdgecontainerClusterCaiObject,
	}
}

func GetEdgecontainerClusterCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//edgecontainer.googleapis.com/projects/{{project}}/locations/{{location}}/clusters/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetEdgecontainerClusterApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: EdgecontainerClusterAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/edgecontainer/v1/rest",
				DiscoveryName:        "Cluster",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
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
	controlPlaneProp, err := expandEdgecontainerClusterControlPlane(d.Get("control_plane"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("control_plane"); !tpgresource.IsEmptyValue(reflect.ValueOf(controlPlaneProp)) && (ok || !reflect.DeepEqual(v, controlPlaneProp)) {
		obj["controlPlane"] = controlPlaneProp
	}
	systemAddonsConfigProp, err := expandEdgecontainerClusterSystemAddonsConfig(d.Get("system_addons_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("system_addons_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(systemAddonsConfigProp)) && (ok || !reflect.DeepEqual(v, systemAddonsConfigProp)) {
		obj["systemAddonsConfig"] = systemAddonsConfigProp
	}
	externalLoadBalancerIpv4AddressPoolsProp, err := expandEdgecontainerClusterExternalLoadBalancerIpv4AddressPools(d.Get("external_load_balancer_ipv4_address_pools"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("external_load_balancer_ipv4_address_pools"); !tpgresource.IsEmptyValue(reflect.ValueOf(externalLoadBalancerIpv4AddressPoolsProp)) && (ok || !reflect.DeepEqual(v, externalLoadBalancerIpv4AddressPoolsProp)) {
		obj["externalLoadBalancerIpv4AddressPools"] = externalLoadBalancerIpv4AddressPoolsProp
	}
	controlPlaneEncryptionProp, err := expandEdgecontainerClusterControlPlaneEncryption(d.Get("control_plane_encryption"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("control_plane_encryption"); !tpgresource.IsEmptyValue(reflect.ValueOf(controlPlaneEncryptionProp)) && (ok || !reflect.DeepEqual(v, controlPlaneEncryptionProp)) {
		obj["controlPlaneEncryption"] = controlPlaneEncryptionProp
	}
	targetVersionProp, err := expandEdgecontainerClusterTargetVersion(d.Get("target_version"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("target_version"); !tpgresource.IsEmptyValue(reflect.ValueOf(targetVersionProp)) && (ok || !reflect.DeepEqual(v, targetVersionProp)) {
		obj["targetVersion"] = targetVersionProp
	}
	releaseChannelProp, err := expandEdgecontainerClusterReleaseChannel(d.Get("release_channel"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("release_channel"); !tpgresource.IsEmptyValue(reflect.ValueOf(releaseChannelProp)) && (ok || !reflect.DeepEqual(v, releaseChannelProp)) {
		obj["releaseChannel"] = releaseChannelProp
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

	transformedMembership, err := expandEdgecontainerClusterFleetMembership(original["membership"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMembership); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["membership"] = transformedMembership
	}

	return transformed, nil
}

func expandEdgecontainerClusterFleetProject(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandEdgecontainerClusterFleetMembership(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
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

	transformedClusterIpv6CidrBlocks, err := expandEdgecontainerClusterNetworkingClusterIpv6CidrBlocks(original["cluster_ipv6_cidr_blocks"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedClusterIpv6CidrBlocks); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["clusterIpv6CidrBlocks"] = transformedClusterIpv6CidrBlocks
	}

	transformedServicesIpv6CidrBlocks, err := expandEdgecontainerClusterNetworkingServicesIpv6CidrBlocks(original["services_ipv6_cidr_blocks"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedServicesIpv6CidrBlocks); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["servicesIpv6CidrBlocks"] = transformedServicesIpv6CidrBlocks
	}

	transformedNetworkType, err := expandEdgecontainerClusterNetworkingNetworkType(original["network_type"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNetworkType); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["networkType"] = transformedNetworkType
	}

	return transformed, nil
}

func expandEdgecontainerClusterNetworkingClusterIpv4CidrBlocks(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandEdgecontainerClusterNetworkingServicesIpv4CidrBlocks(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandEdgecontainerClusterNetworkingClusterIpv6CidrBlocks(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandEdgecontainerClusterNetworkingServicesIpv6CidrBlocks(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandEdgecontainerClusterNetworkingNetworkType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
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

func expandEdgecontainerClusterControlPlane(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedRemote, err := expandEdgecontainerClusterControlPlaneRemote(original["remote"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRemote); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["remote"] = transformedRemote
	}

	transformedLocal, err := expandEdgecontainerClusterControlPlaneLocal(original["local"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLocal); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["local"] = transformedLocal
	}

	return transformed, nil
}

func expandEdgecontainerClusterControlPlaneRemote(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedNodeLocation, err := expandEdgecontainerClusterControlPlaneRemoteNodeLocation(original["node_location"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNodeLocation); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["nodeLocation"] = transformedNodeLocation
	}

	return transformed, nil
}

func expandEdgecontainerClusterControlPlaneRemoteNodeLocation(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandEdgecontainerClusterControlPlaneLocal(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedNodeLocation, err := expandEdgecontainerClusterControlPlaneLocalNodeLocation(original["node_location"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNodeLocation); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["nodeLocation"] = transformedNodeLocation
	}

	transformedNodeCount, err := expandEdgecontainerClusterControlPlaneLocalNodeCount(original["node_count"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNodeCount); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["nodeCount"] = transformedNodeCount
	}

	transformedMachineFilter, err := expandEdgecontainerClusterControlPlaneLocalMachineFilter(original["machine_filter"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMachineFilter); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["machineFilter"] = transformedMachineFilter
	}

	transformedSharedDeploymentPolicy, err := expandEdgecontainerClusterControlPlaneLocalSharedDeploymentPolicy(original["shared_deployment_policy"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSharedDeploymentPolicy); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["sharedDeploymentPolicy"] = transformedSharedDeploymentPolicy
	}

	return transformed, nil
}

func expandEdgecontainerClusterControlPlaneLocalNodeLocation(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandEdgecontainerClusterControlPlaneLocalNodeCount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandEdgecontainerClusterControlPlaneLocalMachineFilter(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandEdgecontainerClusterControlPlaneLocalSharedDeploymentPolicy(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandEdgecontainerClusterSystemAddonsConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedIngress, err := expandEdgecontainerClusterSystemAddonsConfigIngress(original["ingress"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedIngress); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["ingress"] = transformedIngress
	}

	return transformed, nil
}

func expandEdgecontainerClusterSystemAddonsConfigIngress(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedDisabled, err := expandEdgecontainerClusterSystemAddonsConfigIngressDisabled(original["disabled"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDisabled); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["disabled"] = transformedDisabled
	}

	transformedIpv4Vip, err := expandEdgecontainerClusterSystemAddonsConfigIngressIpv4Vip(original["ipv4_vip"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedIpv4Vip); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["ipv4Vip"] = transformedIpv4Vip
	}

	return transformed, nil
}

func expandEdgecontainerClusterSystemAddonsConfigIngressDisabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandEdgecontainerClusterSystemAddonsConfigIngressIpv4Vip(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandEdgecontainerClusterExternalLoadBalancerIpv4AddressPools(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandEdgecontainerClusterControlPlaneEncryption(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedKmsKey, err := expandEdgecontainerClusterControlPlaneEncryptionKmsKey(original["kms_key"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKmsKey); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["kmsKey"] = transformedKmsKey
	}

	transformedKmsKeyActiveVersion, err := expandEdgecontainerClusterControlPlaneEncryptionKmsKeyActiveVersion(original["kms_key_active_version"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKmsKeyActiveVersion); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["kmsKeyActiveVersion"] = transformedKmsKeyActiveVersion
	}

	transformedKmsKeyState, err := expandEdgecontainerClusterControlPlaneEncryptionKmsKeyState(original["kms_key_state"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKmsKeyState); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["kmsKeyState"] = transformedKmsKeyState
	}

	transformedKmsStatus, err := expandEdgecontainerClusterControlPlaneEncryptionKmsStatus(original["kms_status"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKmsStatus); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["kmsStatus"] = transformedKmsStatus
	}

	return transformed, nil
}

func expandEdgecontainerClusterControlPlaneEncryptionKmsKey(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandEdgecontainerClusterControlPlaneEncryptionKmsKeyActiveVersion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandEdgecontainerClusterControlPlaneEncryptionKmsKeyState(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandEdgecontainerClusterControlPlaneEncryptionKmsStatus(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedCode, err := expandEdgecontainerClusterControlPlaneEncryptionKmsStatusCode(original["code"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCode); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["code"] = transformedCode
	}

	transformedMessage, err := expandEdgecontainerClusterControlPlaneEncryptionKmsStatusMessage(original["message"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMessage); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["message"] = transformedMessage
	}

	return transformed, nil
}

func expandEdgecontainerClusterControlPlaneEncryptionKmsStatusCode(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandEdgecontainerClusterControlPlaneEncryptionKmsStatusMessage(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandEdgecontainerClusterTargetVersion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandEdgecontainerClusterReleaseChannel(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
