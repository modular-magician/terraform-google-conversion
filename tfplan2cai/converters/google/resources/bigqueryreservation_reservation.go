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

const BigqueryReservationReservationAssetType string = "bigqueryreservation.googleapis.com/Reservation"

func resourceConverterBigqueryReservationReservation() ResourceConverter {
	return ResourceConverter{
		AssetType: BigqueryReservationReservationAssetType,
		Convert:   GetBigqueryReservationReservationCaiObject,
	}
}

func GetBigqueryReservationReservationCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]Asset, error) {
	name, err := assetName(d, config, "//bigqueryreservation.googleapis.com/projects/{{project}}/locations/{{location}}/reservations/{{name}}")
	if err != nil {
		return []Asset{}, err
	}
	if obj, err := GetBigqueryReservationReservationApiObject(d, config); err == nil {
		return []Asset{{
			Name: name,
			Type: BigqueryReservationReservationAssetType,
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/bigqueryreservation/v1/rest",
				DiscoveryName:        "Reservation",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []Asset{}, err
	}
}

func GetBigqueryReservationReservationApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	slotCapacityProp, err := expandBigqueryReservationReservationSlotCapacity(d.Get("slot_capacity"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("slot_capacity"); !tpgresource.IsEmptyValue(reflect.ValueOf(slotCapacityProp)) && (ok || !reflect.DeepEqual(v, slotCapacityProp)) {
		obj["slotCapacity"] = slotCapacityProp
	}
	ignoreIdleSlotsProp, err := expandBigqueryReservationReservationIgnoreIdleSlots(d.Get("ignore_idle_slots"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("ignore_idle_slots"); !tpgresource.IsEmptyValue(reflect.ValueOf(ignoreIdleSlotsProp)) && (ok || !reflect.DeepEqual(v, ignoreIdleSlotsProp)) {
		obj["ignoreIdleSlots"] = ignoreIdleSlotsProp
	}
	concurrencyProp, err := expandBigqueryReservationReservationConcurrency(d.Get("concurrency"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("concurrency"); !tpgresource.IsEmptyValue(reflect.ValueOf(concurrencyProp)) && (ok || !reflect.DeepEqual(v, concurrencyProp)) {
		obj["concurrency"] = concurrencyProp
	}
	multiRegionAuxiliaryProp, err := expandBigqueryReservationReservationMultiRegionAuxiliary(d.Get("multi_region_auxiliary"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("multi_region_auxiliary"); !tpgresource.IsEmptyValue(reflect.ValueOf(multiRegionAuxiliaryProp)) && (ok || !reflect.DeepEqual(v, multiRegionAuxiliaryProp)) {
		obj["multiRegionAuxiliary"] = multiRegionAuxiliaryProp
	}
	editionProp, err := expandBigqueryReservationReservationEdition(d.Get("edition"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("edition"); !tpgresource.IsEmptyValue(reflect.ValueOf(editionProp)) && (ok || !reflect.DeepEqual(v, editionProp)) {
		obj["edition"] = editionProp
	}
	autoscaleProp, err := expandBigqueryReservationReservationAutoscale(d.Get("autoscale"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("autoscale"); !tpgresource.IsEmptyValue(reflect.ValueOf(autoscaleProp)) && (ok || !reflect.DeepEqual(v, autoscaleProp)) {
		obj["autoscale"] = autoscaleProp
	}

	return obj, nil
}

func expandBigqueryReservationReservationSlotCapacity(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryReservationReservationIgnoreIdleSlots(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryReservationReservationConcurrency(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryReservationReservationMultiRegionAuxiliary(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryReservationReservationEdition(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryReservationReservationAutoscale(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedCurrentSlots, err := expandBigqueryReservationReservationAutoscaleCurrentSlots(original["current_slots"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCurrentSlots); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["currentSlots"] = transformedCurrentSlots
	}

	transformedMaxSlots, err := expandBigqueryReservationReservationAutoscaleMaxSlots(original["max_slots"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxSlots); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["maxSlots"] = transformedMaxSlots
	}

	return transformed, nil
}

func expandBigqueryReservationReservationAutoscaleCurrentSlots(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryReservationReservationAutoscaleMaxSlots(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
