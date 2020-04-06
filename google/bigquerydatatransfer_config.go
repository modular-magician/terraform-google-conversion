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

func GetBigqueryDataTransferConfigCaiObject(d TerraformResourceData, config *Config) (Asset, error) {
	name, err := assetName(d, config, "//bigquerydatatransfer.googleapis.com/{{name}}?serviceAccountName={{serviceAccountName}}")
	if err != nil {
		return Asset{}, err
	}
	if obj, err := GetBigqueryDataTransferConfigApiObject(d, config); err == nil {
		return Asset{
			Name: name,
			Type: "bigquerydatatransfer.googleapis.com/Config",
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/bigquerydatatransfer/v1/rest",
				DiscoveryName:        "Config",
				Data:                 obj,
			},
		}, nil
	} else {
		return Asset{}, err
	}
}

func GetBigqueryDataTransferConfigApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	displayNameProp, err := expandBigqueryDataTransferConfigDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	destinationDatasetIdProp, err := expandBigqueryDataTransferConfigDestinationDatasetId(d.Get("destination_dataset_id"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("destination_dataset_id"); !isEmptyValue(reflect.ValueOf(destinationDatasetIdProp)) && (ok || !reflect.DeepEqual(v, destinationDatasetIdProp)) {
		obj["destinationDatasetId"] = destinationDatasetIdProp
	}
	dataSourceIdProp, err := expandBigqueryDataTransferConfigDataSourceId(d.Get("data_source_id"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("data_source_id"); !isEmptyValue(reflect.ValueOf(dataSourceIdProp)) && (ok || !reflect.DeepEqual(v, dataSourceIdProp)) {
		obj["dataSourceId"] = dataSourceIdProp
	}
	scheduleProp, err := expandBigqueryDataTransferConfigSchedule(d.Get("schedule"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("schedule"); !isEmptyValue(reflect.ValueOf(scheduleProp)) && (ok || !reflect.DeepEqual(v, scheduleProp)) {
		obj["schedule"] = scheduleProp
	}
	dataRefreshWindowDaysProp, err := expandBigqueryDataTransferConfigDataRefreshWindowDays(d.Get("data_refresh_window_days"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("data_refresh_window_days"); !isEmptyValue(reflect.ValueOf(dataRefreshWindowDaysProp)) && (ok || !reflect.DeepEqual(v, dataRefreshWindowDaysProp)) {
		obj["dataRefreshWindowDays"] = dataRefreshWindowDaysProp
	}
	disabledProp, err := expandBigqueryDataTransferConfigDisabled(d.Get("disabled"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("disabled"); !isEmptyValue(reflect.ValueOf(disabledProp)) && (ok || !reflect.DeepEqual(v, disabledProp)) {
		obj["disabled"] = disabledProp
	}
	paramsProp, err := expandBigqueryDataTransferConfigParams(d.Get("params"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("params"); !isEmptyValue(reflect.ValueOf(paramsProp)) && (ok || !reflect.DeepEqual(v, paramsProp)) {
		obj["params"] = paramsProp
	}

	return obj, nil
}

func expandBigqueryDataTransferConfigDisplayName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryDataTransferConfigDestinationDatasetId(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryDataTransferConfigDataSourceId(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryDataTransferConfigSchedule(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryDataTransferConfigDataRefreshWindowDays(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryDataTransferConfigDisabled(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryDataTransferConfigParams(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
