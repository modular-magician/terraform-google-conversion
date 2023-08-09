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

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgresource"
	transport_tpg "github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/transport"
)

const ComputeInstanceGroupNamedPortAssetType string = "compute.googleapis.com/InstanceGroupNamedPort"

func ResourceConverterComputeInstanceGroupNamedPort() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType: ComputeInstanceGroupNamedPortAssetType,
		Convert:   GetComputeInstanceGroupNamedPortCaiObject,
	}
}

func GetComputeInstanceGroupNamedPortCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	name, err := tpgresource.AssetName(d, config, "//compute.googleapis.com/projects/{{project}}/zones/{{zone}}/instanceGroups/{{group}}")
	if err != nil {
		return []tpgresource.Asset{}, err
	}
	if obj, err := GetComputeInstanceGroupNamedPortApiObject(d, config); err == nil {
		return []tpgresource.Asset{{
			Name: name,
			Type: ComputeInstanceGroupNamedPortAssetType,
			Resource: &tpgresource.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/compute/v1/rest",
				DiscoveryName:        "InstanceGroupNamedPort",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []tpgresource.Asset{}, err
	}
}

func GetComputeInstanceGroupNamedPortApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandComputeInstanceGroupNamedPortName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	portProp, err := expandComputeInstanceGroupNamedPortPort(d.Get("port"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("port"); !tpgresource.IsEmptyValue(reflect.ValueOf(portProp)) && (ok || !reflect.DeepEqual(v, portProp)) {
		obj["port"] = portProp
	}

	return resourceComputeInstanceGroupNamedPortEncoder(d, config, obj)
}

func resourceComputeInstanceGroupNamedPortEncoder(d tpgresource.TerraformResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	config := meta.(*transport_tpg.Config)
	ig, err := tpgresource.ParseInstanceGroupFieldValue(d.Get("group").(string), d, config)
	if err != nil {
		return nil, err
	}

	if err := d.Set("group", ig.Name); err != nil {
		return nil, fmt.Errorf("Error setting group: %s", err)
	}
	if err := d.Set("zone", ig.Zone); err != nil {
		return nil, fmt.Errorf("Error setting zone: %s", err)
	}
	if err := d.Set("project", ig.Project); err != nil {
		return nil, fmt.Errorf("Error setting project: %s", err)
	}

	return obj, nil
}

func expandComputeInstanceGroupNamedPortName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeInstanceGroupNamedPortPort(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
