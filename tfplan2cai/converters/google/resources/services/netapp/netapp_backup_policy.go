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

package netapp

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const NetappBackupPolicyAssetType string = "netapp.googleapis.com/BackupPolicy"

func ResourceConverterNetappBackupPolicy() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: NetappBackupPolicyAssetType,
		Convert:   GetNetappBackupPolicyCaiObject,
	}
}

func GetNetappBackupPolicyCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//netapp.googleapis.com/projects/{{project}}/locations/{{location}}/backupPolicies/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetNetappBackupPolicyApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: NetappBackupPolicyAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1beta1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/netapp/v1beta1/rest",
				DiscoveryName:        "BackupPolicy",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetNetappBackupPolicyApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	dailyBackupLimitProp, err := expandNetappBackupPolicyDailyBackupLimit(d.Get("daily_backup_limit"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("daily_backup_limit"); !tpgresource.IsEmptyValue(reflect.ValueOf(dailyBackupLimitProp)) && (ok || !reflect.DeepEqual(v, dailyBackupLimitProp)) {
		obj["dailyBackupLimit"] = dailyBackupLimitProp
	}
	weeklyBackupLimitProp, err := expandNetappBackupPolicyWeeklyBackupLimit(d.Get("weekly_backup_limit"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("weekly_backup_limit"); !tpgresource.IsEmptyValue(reflect.ValueOf(weeklyBackupLimitProp)) && (ok || !reflect.DeepEqual(v, weeklyBackupLimitProp)) {
		obj["weeklyBackupLimit"] = weeklyBackupLimitProp
	}
	monthlyBackupLimitProp, err := expandNetappBackupPolicyMonthlyBackupLimit(d.Get("monthly_backup_limit"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("monthly_backup_limit"); !tpgresource.IsEmptyValue(reflect.ValueOf(monthlyBackupLimitProp)) && (ok || !reflect.DeepEqual(v, monthlyBackupLimitProp)) {
		obj["monthlyBackupLimit"] = monthlyBackupLimitProp
	}
	descriptionProp, err := expandNetappBackupPolicyDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	enabledProp, err := expandNetappBackupPolicyEnabled(d.Get("enabled"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("enabled"); ok || !reflect.DeepEqual(v, enabledProp) {
		obj["enabled"] = enabledProp
	}
	labelsProp, err := expandNetappBackupPolicyEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	return obj, nil
}

func expandNetappBackupPolicyDailyBackupLimit(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetappBackupPolicyWeeklyBackupLimit(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetappBackupPolicyMonthlyBackupLimit(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetappBackupPolicyDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetappBackupPolicyEnabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetappBackupPolicyEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
