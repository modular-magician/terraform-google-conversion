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

package compute

import (
	"fmt"
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const ComputeStoragePoolAssetType string = "compute.googleapis.com/StoragePool"

func ResourceConverterComputeStoragePool() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: ComputeStoragePoolAssetType,
		Convert:   GetComputeStoragePoolCaiObject,
	}
}

func GetComputeStoragePoolCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//compute.googleapis.com/projects/{{project}}/zones/{{zone}}/storagePools/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetComputeStoragePoolApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: ComputeStoragePoolAssetType,
			Resource: &cai.AssetResource{
				Version:              "beta",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/compute/beta/rest",
				DiscoveryName:        "StoragePool",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetComputeStoragePoolApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandComputeStoragePoolName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	descriptionProp, err := expandComputeStoragePoolDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	poolProvisionedCapacityGbProp, err := expandComputeStoragePoolPoolProvisionedCapacityGb(d.Get("pool_provisioned_capacity_gb"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("pool_provisioned_capacity_gb"); !tpgresource.IsEmptyValue(reflect.ValueOf(poolProvisionedCapacityGbProp)) && (ok || !reflect.DeepEqual(v, poolProvisionedCapacityGbProp)) {
		obj["poolProvisionedCapacityGb"] = poolProvisionedCapacityGbProp
	}
	poolProvisionedIopsProp, err := expandComputeStoragePoolPoolProvisionedIops(d.Get("pool_provisioned_iops"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("pool_provisioned_iops"); !tpgresource.IsEmptyValue(reflect.ValueOf(poolProvisionedIopsProp)) && (ok || !reflect.DeepEqual(v, poolProvisionedIopsProp)) {
		obj["poolProvisionedIops"] = poolProvisionedIopsProp
	}
	poolProvisionedThroughputProp, err := expandComputeStoragePoolPoolProvisionedThroughput(d.Get("pool_provisioned_throughput"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("pool_provisioned_throughput"); !tpgresource.IsEmptyValue(reflect.ValueOf(poolProvisionedThroughputProp)) && (ok || !reflect.DeepEqual(v, poolProvisionedThroughputProp)) {
		obj["poolProvisionedThroughput"] = poolProvisionedThroughputProp
	}
	storagePoolTypeProp, err := expandComputeStoragePoolStoragePoolType(d.Get("storage_pool_type"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("storage_pool_type"); !tpgresource.IsEmptyValue(reflect.ValueOf(storagePoolTypeProp)) && (ok || !reflect.DeepEqual(v, storagePoolTypeProp)) {
		obj["storagePoolType"] = storagePoolTypeProp
	}
	capacityProvisioningTypeProp, err := expandComputeStoragePoolCapacityProvisioningType(d.Get("capacity_provisioning_type"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("capacity_provisioning_type"); !tpgresource.IsEmptyValue(reflect.ValueOf(capacityProvisioningTypeProp)) && (ok || !reflect.DeepEqual(v, capacityProvisioningTypeProp)) {
		obj["capacityProvisioningType"] = capacityProvisioningTypeProp
	}
	performanceProvisioningTypeProp, err := expandComputeStoragePoolPerformanceProvisioningType(d.Get("performance_provisioning_type"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("performance_provisioning_type"); !tpgresource.IsEmptyValue(reflect.ValueOf(performanceProvisioningTypeProp)) && (ok || !reflect.DeepEqual(v, performanceProvisioningTypeProp)) {
		obj["performanceProvisioningType"] = performanceProvisioningTypeProp
	}
	zoneProp, err := expandComputeStoragePoolZone(d.Get("zone"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("zone"); !tpgresource.IsEmptyValue(reflect.ValueOf(zoneProp)) && (ok || !reflect.DeepEqual(v, zoneProp)) {
		obj["zone"] = zoneProp
	}

	return obj, nil
}

func expandComputeStoragePoolName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeStoragePoolDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeStoragePoolPoolProvisionedCapacityGb(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeStoragePoolPoolProvisionedIops(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeStoragePoolPoolProvisionedThroughput(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeStoragePoolStoragePoolType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	f, err := tpgresource.ParseZonalFieldValue("storagePoolTypes", v.(string), "project", "zone", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for storage_pool_type: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeStoragePoolCapacityProvisioningType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeStoragePoolPerformanceProvisioningType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeStoragePoolZone(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	f, err := tpgresource.ParseGlobalFieldValue("zones", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for zone: %s", err)
	}
	return f.RelativeLink(), nil
}
