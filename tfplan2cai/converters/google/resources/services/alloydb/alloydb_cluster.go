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

package alloydb

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const AlloydbClusterAssetType string = "alloydb.googleapis.com/Cluster"

func ResourceConverterAlloydbCluster() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: AlloydbClusterAssetType,
		Convert:   GetAlloydbClusterCaiObject,
	}
}

func GetAlloydbClusterCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//alloydb.googleapis.com/projects/{{project}}/locations/{{location}}/clusters/{{cluster_id}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetAlloydbClusterApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: AlloydbClusterAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1beta",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/alloydb/v1beta/rest",
				DiscoveryName:        "Cluster",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetAlloydbClusterApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	encryptionConfigProp, err := expandAlloydbClusterEncryptionConfig(d.Get("encryption_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("encryption_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(encryptionConfigProp)) && (ok || !reflect.DeepEqual(v, encryptionConfigProp)) {
		obj["encryptionConfig"] = encryptionConfigProp
	}
	networkProp, err := expandAlloydbClusterNetwork(d.Get("network"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("network"); !tpgresource.IsEmptyValue(reflect.ValueOf(networkProp)) && (ok || !reflect.DeepEqual(v, networkProp)) {
		obj["network"] = networkProp
	}
	networkConfigProp, err := expandAlloydbClusterNetworkConfig(d.Get("network_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("network_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(networkConfigProp)) && (ok || !reflect.DeepEqual(v, networkConfigProp)) {
		obj["networkConfig"] = networkConfigProp
	}
	displayNameProp, err := expandAlloydbClusterDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	etagProp, err := expandAlloydbClusterEtag(d.Get("etag"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("etag"); !tpgresource.IsEmptyValue(reflect.ValueOf(etagProp)) && (ok || !reflect.DeepEqual(v, etagProp)) {
		obj["etag"] = etagProp
	}
	annotationsProp, err := expandAlloydbClusterAnnotations(d.Get("annotations"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("annotations"); !tpgresource.IsEmptyValue(reflect.ValueOf(annotationsProp)) && (ok || !reflect.DeepEqual(v, annotationsProp)) {
		obj["annotations"] = annotationsProp
	}
	initialUserProp, err := expandAlloydbClusterInitialUser(d.Get("initial_user"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("initial_user"); !tpgresource.IsEmptyValue(reflect.ValueOf(initialUserProp)) && (ok || !reflect.DeepEqual(v, initialUserProp)) {
		obj["initialUser"] = initialUserProp
	}
	restoreBackupSourceProp, err := expandAlloydbClusterRestoreBackupSource(d.Get("restore_backup_source"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("restore_backup_source"); !tpgresource.IsEmptyValue(reflect.ValueOf(restoreBackupSourceProp)) && (ok || !reflect.DeepEqual(v, restoreBackupSourceProp)) {
		obj["restoreBackupSource"] = restoreBackupSourceProp
	}
	restoreContinuousBackupSourceProp, err := expandAlloydbClusterRestoreContinuousBackupSource(d.Get("restore_continuous_backup_source"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("restore_continuous_backup_source"); !tpgresource.IsEmptyValue(reflect.ValueOf(restoreContinuousBackupSourceProp)) && (ok || !reflect.DeepEqual(v, restoreContinuousBackupSourceProp)) {
		obj["restoreContinuousBackupSource"] = restoreContinuousBackupSourceProp
	}
	continuousBackupConfigProp, err := expandAlloydbClusterContinuousBackupConfig(d.Get("continuous_backup_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("continuous_backup_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(continuousBackupConfigProp)) && (ok || !reflect.DeepEqual(v, continuousBackupConfigProp)) {
		obj["continuousBackupConfig"] = continuousBackupConfigProp
	}
	automatedBackupPolicyProp, err := expandAlloydbClusterAutomatedBackupPolicy(d.Get("automated_backup_policy"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("automated_backup_policy"); !tpgresource.IsEmptyValue(reflect.ValueOf(automatedBackupPolicyProp)) && (ok || !reflect.DeepEqual(v, automatedBackupPolicyProp)) {
		obj["automatedBackupPolicy"] = automatedBackupPolicyProp
	}
	labelsProp, err := expandAlloydbClusterEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	return obj, nil
}

func expandAlloydbClusterEncryptionConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedKmsKeyName, err := expandAlloydbClusterEncryptionConfigKmsKeyName(original["kms_key_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKmsKeyName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["kmsKeyName"] = transformedKmsKeyName
	}

	return transformed, nil
}

func expandAlloydbClusterEncryptionConfigKmsKeyName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAlloydbClusterNetwork(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAlloydbClusterNetworkConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedNetwork, err := expandAlloydbClusterNetworkConfigNetwork(original["network"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNetwork); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["network"] = transformedNetwork
	}

	transformedAllocatedIpRange, err := expandAlloydbClusterNetworkConfigAllocatedIpRange(original["allocated_ip_range"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAllocatedIpRange); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["allocatedIpRange"] = transformedAllocatedIpRange
	}

	return transformed, nil
}

func expandAlloydbClusterNetworkConfigNetwork(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAlloydbClusterNetworkConfigAllocatedIpRange(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAlloydbClusterDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAlloydbClusterEtag(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAlloydbClusterAnnotations(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandAlloydbClusterInitialUser(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedUser, err := expandAlloydbClusterInitialUserUser(original["user"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUser); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["user"] = transformedUser
	}

	transformedPassword, err := expandAlloydbClusterInitialUserPassword(original["password"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPassword); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["password"] = transformedPassword
	}

	return transformed, nil
}

func expandAlloydbClusterInitialUserUser(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAlloydbClusterInitialUserPassword(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAlloydbClusterRestoreBackupSource(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedBackupName, err := expandAlloydbClusterRestoreBackupSourceBackupName(original["backup_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedBackupName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["backupName"] = transformedBackupName
	}

	return transformed, nil
}

func expandAlloydbClusterRestoreBackupSourceBackupName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAlloydbClusterRestoreContinuousBackupSource(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedCluster, err := expandAlloydbClusterRestoreContinuousBackupSourceCluster(original["cluster"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCluster); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["cluster"] = transformedCluster
	}

	transformedPointInTime, err := expandAlloydbClusterRestoreContinuousBackupSourcePointInTime(original["point_in_time"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPointInTime); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["pointInTime"] = transformedPointInTime
	}

	return transformed, nil
}

func expandAlloydbClusterRestoreContinuousBackupSourceCluster(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAlloydbClusterRestoreContinuousBackupSourcePointInTime(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAlloydbClusterContinuousBackupConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedEnabled, err := expandAlloydbClusterContinuousBackupConfigEnabled(original["enabled"], d, config)
	if err != nil {
		return nil, err
	} else {
		transformed["enabled"] = transformedEnabled
	}

	transformedRecoveryWindowDays, err := expandAlloydbClusterContinuousBackupConfigRecoveryWindowDays(original["recovery_window_days"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRecoveryWindowDays); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["recoveryWindowDays"] = transformedRecoveryWindowDays
	}

	transformedEncryptionConfig, err := expandAlloydbClusterContinuousBackupConfigEncryptionConfig(original["encryption_config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEncryptionConfig); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["encryptionConfig"] = transformedEncryptionConfig
	}

	return transformed, nil
}

func expandAlloydbClusterContinuousBackupConfigEnabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAlloydbClusterContinuousBackupConfigRecoveryWindowDays(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAlloydbClusterContinuousBackupConfigEncryptionConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedKmsKeyName, err := expandAlloydbClusterContinuousBackupConfigEncryptionConfigKmsKeyName(original["kms_key_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKmsKeyName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["kmsKeyName"] = transformedKmsKeyName
	}

	return transformed, nil
}

func expandAlloydbClusterContinuousBackupConfigEncryptionConfigKmsKeyName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAlloydbClusterAutomatedBackupPolicy(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedBackupWindow, err := expandAlloydbClusterAutomatedBackupPolicyBackupWindow(original["backup_window"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedBackupWindow); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["backupWindow"] = transformedBackupWindow
	}

	transformedLocation, err := expandAlloydbClusterAutomatedBackupPolicyLocation(original["location"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLocation); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["location"] = transformedLocation
	}

	transformedLabels, err := expandAlloydbClusterAutomatedBackupPolicyLabels(original["labels"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLabels); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["labels"] = transformedLabels
	}

	transformedEncryptionConfig, err := expandAlloydbClusterAutomatedBackupPolicyEncryptionConfig(original["encryption_config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEncryptionConfig); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["encryptionConfig"] = transformedEncryptionConfig
	}

	transformedWeeklySchedule, err := expandAlloydbClusterAutomatedBackupPolicyWeeklySchedule(original["weekly_schedule"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedWeeklySchedule); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["weeklySchedule"] = transformedWeeklySchedule
	}

	transformedTimeBasedRetention, err := expandAlloydbClusterAutomatedBackupPolicyTimeBasedRetention(original["time_based_retention"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTimeBasedRetention); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["timeBasedRetention"] = transformedTimeBasedRetention
	}

	transformedQuantityBasedRetention, err := expandAlloydbClusterAutomatedBackupPolicyQuantityBasedRetention(original["quantity_based_retention"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedQuantityBasedRetention); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["quantityBasedRetention"] = transformedQuantityBasedRetention
	}

	transformedEnabled, err := expandAlloydbClusterAutomatedBackupPolicyEnabled(original["enabled"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEnabled); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["enabled"] = transformedEnabled
	}

	return transformed, nil
}

func expandAlloydbClusterAutomatedBackupPolicyBackupWindow(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAlloydbClusterAutomatedBackupPolicyLocation(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAlloydbClusterAutomatedBackupPolicyLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandAlloydbClusterAutomatedBackupPolicyEncryptionConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedKmsKeyName, err := expandAlloydbClusterAutomatedBackupPolicyEncryptionConfigKmsKeyName(original["kms_key_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKmsKeyName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["kmsKeyName"] = transformedKmsKeyName
	}

	return transformed, nil
}

func expandAlloydbClusterAutomatedBackupPolicyEncryptionConfigKmsKeyName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAlloydbClusterAutomatedBackupPolicyWeeklySchedule(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedDaysOfWeek, err := expandAlloydbClusterAutomatedBackupPolicyWeeklyScheduleDaysOfWeek(original["days_of_week"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDaysOfWeek); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["daysOfWeek"] = transformedDaysOfWeek
	}

	transformedStartTimes, err := expandAlloydbClusterAutomatedBackupPolicyWeeklyScheduleStartTimes(original["start_times"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedStartTimes); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["startTimes"] = transformedStartTimes
	}

	return transformed, nil
}

func expandAlloydbClusterAutomatedBackupPolicyWeeklyScheduleDaysOfWeek(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAlloydbClusterAutomatedBackupPolicyWeeklyScheduleStartTimes(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedHours, err := expandAlloydbClusterAutomatedBackupPolicyWeeklyScheduleStartTimesHours(original["hours"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedHours); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["hours"] = transformedHours
		}

		transformedMinutes, err := expandAlloydbClusterAutomatedBackupPolicyWeeklyScheduleStartTimesMinutes(original["minutes"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedMinutes); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["minutes"] = transformedMinutes
		}

		transformedSeconds, err := expandAlloydbClusterAutomatedBackupPolicyWeeklyScheduleStartTimesSeconds(original["seconds"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedSeconds); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["seconds"] = transformedSeconds
		}

		transformedNanos, err := expandAlloydbClusterAutomatedBackupPolicyWeeklyScheduleStartTimesNanos(original["nanos"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedNanos); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["nanos"] = transformedNanos
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandAlloydbClusterAutomatedBackupPolicyWeeklyScheduleStartTimesHours(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAlloydbClusterAutomatedBackupPolicyWeeklyScheduleStartTimesMinutes(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAlloydbClusterAutomatedBackupPolicyWeeklyScheduleStartTimesSeconds(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAlloydbClusterAutomatedBackupPolicyWeeklyScheduleStartTimesNanos(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAlloydbClusterAutomatedBackupPolicyTimeBasedRetention(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedRetentionPeriod, err := expandAlloydbClusterAutomatedBackupPolicyTimeBasedRetentionRetentionPeriod(original["retention_period"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRetentionPeriod); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["retentionPeriod"] = transformedRetentionPeriod
	}

	return transformed, nil
}

func expandAlloydbClusterAutomatedBackupPolicyTimeBasedRetentionRetentionPeriod(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAlloydbClusterAutomatedBackupPolicyQuantityBasedRetention(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedCount, err := expandAlloydbClusterAutomatedBackupPolicyQuantityBasedRetentionCount(original["count"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCount); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["count"] = transformedCount
	}

	return transformed, nil
}

func expandAlloydbClusterAutomatedBackupPolicyQuantityBasedRetentionCount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAlloydbClusterAutomatedBackupPolicyEnabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAlloydbClusterEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
