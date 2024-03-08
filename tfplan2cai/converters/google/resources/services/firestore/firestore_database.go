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

package firestore

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const FirestoreDatabaseAssetType string = "firestore.googleapis.com/Database"

func ResourceConverterFirestoreDatabase() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: FirestoreDatabaseAssetType,
		Convert:   GetFirestoreDatabaseCaiObject,
	}
}

func GetFirestoreDatabaseCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//firestore.googleapis.com/projects/{{project}}/databases/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetFirestoreDatabaseApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: FirestoreDatabaseAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/firestore/v1/rest",
				DiscoveryName:        "Database",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetFirestoreDatabaseApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandFirestoreDatabaseName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	locationIdProp, err := expandFirestoreDatabaseLocationId(d.Get("location_id"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("location_id"); !tpgresource.IsEmptyValue(reflect.ValueOf(locationIdProp)) && (ok || !reflect.DeepEqual(v, locationIdProp)) {
		obj["locationId"] = locationIdProp
	}
	typeProp, err := expandFirestoreDatabaseType(d.Get("type"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("type"); !tpgresource.IsEmptyValue(reflect.ValueOf(typeProp)) && (ok || !reflect.DeepEqual(v, typeProp)) {
		obj["type"] = typeProp
	}
	concurrencyModeProp, err := expandFirestoreDatabaseConcurrencyMode(d.Get("concurrency_mode"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("concurrency_mode"); !tpgresource.IsEmptyValue(reflect.ValueOf(concurrencyModeProp)) && (ok || !reflect.DeepEqual(v, concurrencyModeProp)) {
		obj["concurrencyMode"] = concurrencyModeProp
	}
	appEngineIntegrationModeProp, err := expandFirestoreDatabaseAppEngineIntegrationMode(d.Get("app_engine_integration_mode"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("app_engine_integration_mode"); !tpgresource.IsEmptyValue(reflect.ValueOf(appEngineIntegrationModeProp)) && (ok || !reflect.DeepEqual(v, appEngineIntegrationModeProp)) {
		obj["appEngineIntegrationMode"] = appEngineIntegrationModeProp
	}
	pointInTimeRecoveryEnablementProp, err := expandFirestoreDatabasePointInTimeRecoveryEnablement(d.Get("point_in_time_recovery_enablement"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("point_in_time_recovery_enablement"); !tpgresource.IsEmptyValue(reflect.ValueOf(pointInTimeRecoveryEnablementProp)) && (ok || !reflect.DeepEqual(v, pointInTimeRecoveryEnablementProp)) {
		obj["pointInTimeRecoveryEnablement"] = pointInTimeRecoveryEnablementProp
	}
	deleteProtectionStateProp, err := expandFirestoreDatabaseDeleteProtectionState(d.Get("delete_protection_state"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("delete_protection_state"); !tpgresource.IsEmptyValue(reflect.ValueOf(deleteProtectionStateProp)) && (ok || !reflect.DeepEqual(v, deleteProtectionStateProp)) {
		obj["deleteProtectionState"] = deleteProtectionStateProp
	}
	etagProp, err := expandFirestoreDatabaseEtag(d.Get("etag"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("etag"); !tpgresource.IsEmptyValue(reflect.ValueOf(etagProp)) && (ok || !reflect.DeepEqual(v, etagProp)) {
		obj["etag"] = etagProp
	}
	cmekConfigProp, err := expandFirestoreDatabaseCmekConfig(d.Get("cmek_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("cmek_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(cmekConfigProp)) && (ok || !reflect.DeepEqual(v, cmekConfigProp)) {
		obj["cmekConfig"] = cmekConfigProp
	}

	return obj, nil
}

func expandFirestoreDatabaseName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return tpgresource.ReplaceVars(d, config, "projects/{{project}}/databases/{{name}}")
}

func expandFirestoreDatabaseLocationId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandFirestoreDatabaseType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandFirestoreDatabaseConcurrencyMode(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandFirestoreDatabaseAppEngineIntegrationMode(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandFirestoreDatabasePointInTimeRecoveryEnablement(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandFirestoreDatabaseDeleteProtectionState(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandFirestoreDatabaseEtag(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandFirestoreDatabaseCmekConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedKmsKeyName, err := expandFirestoreDatabaseCmekConfigKmsKeyName(original["kms_key_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKmsKeyName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["kmsKeyName"] = transformedKmsKeyName
	}

	transformedActiveKeyVersion, err := expandFirestoreDatabaseCmekConfigActiveKeyVersion(original["active_key_version"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedActiveKeyVersion); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["activeKeyVersion"] = transformedActiveKeyVersion
	}

	return transformed, nil
}

func expandFirestoreDatabaseCmekConfigKmsKeyName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandFirestoreDatabaseCmekConfigActiveKeyVersion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
