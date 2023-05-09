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
	"fmt"
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgresource"
	transport_tpg "github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/transport"
)

const ComputeSnapshotAssetType string = "compute.googleapis.com/Snapshot"

func resourceConverterComputeSnapshot() ResourceConverter {
	return ResourceConverter{
		AssetType: ComputeSnapshotAssetType,
		Convert:   GetComputeSnapshotCaiObject,
	}
}

func GetComputeSnapshotCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]Asset, error) {
	name, err := assetName(d, config, "//compute.googleapis.com/projects/{{project}}/global/snapshots/{{name}}")
	if err != nil {
		return []Asset{}, err
	}
	if obj, err := GetComputeSnapshotApiObject(d, config); err == nil {
		return []Asset{{
			Name: name,
			Type: ComputeSnapshotAssetType,
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/compute/v1/rest",
				DiscoveryName:        "Snapshot",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []Asset{}, err
	}
}

func GetComputeSnapshotApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	chainNameProp, err := expandComputeSnapshotChainName(d.Get("chain_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("chain_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(chainNameProp)) && (ok || !reflect.DeepEqual(v, chainNameProp)) {
		obj["chainName"] = chainNameProp
	}
	nameProp, err := expandComputeSnapshotName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	descriptionProp, err := expandComputeSnapshotDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	storageLocationsProp, err := expandComputeSnapshotStorageLocations(d.Get("storage_locations"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("storage_locations"); !tpgresource.IsEmptyValue(reflect.ValueOf(storageLocationsProp)) && (ok || !reflect.DeepEqual(v, storageLocationsProp)) {
		obj["storageLocations"] = storageLocationsProp
	}
	labelsProp, err := expandComputeSnapshotLabels(d.Get("labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	labelFingerprintProp, err := expandComputeSnapshotLabelFingerprint(d.Get("label_fingerprint"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("label_fingerprint"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelFingerprintProp)) && (ok || !reflect.DeepEqual(v, labelFingerprintProp)) {
		obj["labelFingerprint"] = labelFingerprintProp
	}
	sourceDiskProp, err := expandComputeSnapshotSourceDisk(d.Get("source_disk"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("source_disk"); !tpgresource.IsEmptyValue(reflect.ValueOf(sourceDiskProp)) && (ok || !reflect.DeepEqual(v, sourceDiskProp)) {
		obj["sourceDisk"] = sourceDiskProp
	}
	zoneProp, err := expandComputeSnapshotZone(d.Get("zone"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("zone"); !tpgresource.IsEmptyValue(reflect.ValueOf(zoneProp)) && (ok || !reflect.DeepEqual(v, zoneProp)) {
		obj["zone"] = zoneProp
	}
	snapshotEncryptionKeyProp, err := expandComputeSnapshotSnapshotEncryptionKey(d.Get("snapshot_encryption_key"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("snapshot_encryption_key"); !tpgresource.IsEmptyValue(reflect.ValueOf(snapshotEncryptionKeyProp)) && (ok || !reflect.DeepEqual(v, snapshotEncryptionKeyProp)) {
		obj["snapshotEncryptionKey"] = snapshotEncryptionKeyProp
	}
	sourceDiskEncryptionKeyProp, err := expandComputeSnapshotSourceDiskEncryptionKey(d.Get("source_disk_encryption_key"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("source_disk_encryption_key"); !tpgresource.IsEmptyValue(reflect.ValueOf(sourceDiskEncryptionKeyProp)) && (ok || !reflect.DeepEqual(v, sourceDiskEncryptionKeyProp)) {
		obj["sourceDiskEncryptionKey"] = sourceDiskEncryptionKeyProp
	}

	return obj, nil
}

func expandComputeSnapshotChainName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeSnapshotName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeSnapshotDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeSnapshotStorageLocations(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeSnapshotLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandComputeSnapshotLabelFingerprint(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeSnapshotSourceDisk(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	f, err := tpgresource.ParseZonalFieldValue("disks", v.(string), "project", "zone", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for source_disk: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeSnapshotZone(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	f, err := tpgresource.ParseGlobalFieldValue("zones", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for zone: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeSnapshotSnapshotEncryptionKey(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedRawKey, err := expandComputeSnapshotSnapshotEncryptionKeyRawKey(original["raw_key"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRawKey); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["rawKey"] = transformedRawKey
	}

	transformedSha256, err := expandComputeSnapshotSnapshotEncryptionKeySha256(original["sha256"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSha256); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["sha256"] = transformedSha256
	}

	transformedKmsKeySelfLink, err := expandComputeSnapshotSnapshotEncryptionKeyKmsKeySelfLink(original["kms_key_self_link"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKmsKeySelfLink); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["kmsKeyName"] = transformedKmsKeySelfLink
	}

	transformedKmsKeyServiceAccount, err := expandComputeSnapshotSnapshotEncryptionKeyKmsKeyServiceAccount(original["kms_key_service_account"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKmsKeyServiceAccount); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["kmsKeyServiceAccount"] = transformedKmsKeyServiceAccount
	}

	return transformed, nil
}

func expandComputeSnapshotSnapshotEncryptionKeyRawKey(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeSnapshotSnapshotEncryptionKeySha256(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeSnapshotSnapshotEncryptionKeyKmsKeySelfLink(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeSnapshotSnapshotEncryptionKeyKmsKeyServiceAccount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeSnapshotSourceDiskEncryptionKey(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedRawKey, err := expandComputeSnapshotSourceDiskEncryptionKeyRawKey(original["raw_key"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRawKey); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["rawKey"] = transformedRawKey
	}

	transformedKmsKeyServiceAccount, err := expandComputeSnapshotSourceDiskEncryptionKeyKmsKeyServiceAccount(original["kms_key_service_account"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKmsKeyServiceAccount); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["kmsKeyServiceAccount"] = transformedKmsKeyServiceAccount
	}

	return transformed, nil
}

func expandComputeSnapshotSourceDiskEncryptionKeyRawKey(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeSnapshotSourceDiskEncryptionKeyKmsKeyServiceAccount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
