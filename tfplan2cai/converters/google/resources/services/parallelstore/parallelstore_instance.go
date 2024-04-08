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

package parallelstore

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const ParallelstoreInstanceAssetType string = "parallelstore.googleapis.com/Instance"

func ResourceConverterParallelstoreInstance() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: ParallelstoreInstanceAssetType,
		Convert:   GetParallelstoreInstanceCaiObject,
	}
}

func GetParallelstoreInstanceCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//parallelstore.googleapis.com/projects/{{project}}/locations/{{location}}/instances/{{instance_id}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetParallelstoreInstanceApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: ParallelstoreInstanceAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1beta",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/parallelstore/v1beta/rest",
				DiscoveryName:        "Instance",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetParallelstoreInstanceApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	descriptionProp, err := expandParallelstoreInstanceDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	capacityGibProp, err := expandParallelstoreInstanceCapacityGib(d.Get("capacity_gib"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("capacity_gib"); !tpgresource.IsEmptyValue(reflect.ValueOf(capacityGibProp)) && (ok || !reflect.DeepEqual(v, capacityGibProp)) {
		obj["capacityGib"] = capacityGibProp
	}
	networkProp, err := expandParallelstoreInstanceNetwork(d.Get("network"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("network"); !tpgresource.IsEmptyValue(reflect.ValueOf(networkProp)) && (ok || !reflect.DeepEqual(v, networkProp)) {
		obj["network"] = networkProp
	}
	reservedIpRangeProp, err := expandParallelstoreInstanceReservedIpRange(d.Get("reserved_ip_range"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("reserved_ip_range"); !tpgresource.IsEmptyValue(reflect.ValueOf(reservedIpRangeProp)) && (ok || !reflect.DeepEqual(v, reservedIpRangeProp)) {
		obj["reservedIpRange"] = reservedIpRangeProp
	}
	labelsProp, err := expandParallelstoreInstanceEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	return obj, nil
}

func expandParallelstoreInstanceDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandParallelstoreInstanceCapacityGib(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandParallelstoreInstanceNetwork(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandParallelstoreInstanceReservedIpRange(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandParallelstoreInstanceEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
