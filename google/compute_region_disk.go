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

import (
	"fmt"
	"log"
	"reflect"
)

func GetComputeRegionDiskCaiObject(d TerraformResourceData, config *Config) (Asset, error) {
	if obj, err := GetComputeRegionDiskApiObject(d, config); err == nil {
		return Asset{
			Name: fmt.Sprintf("//compute.googleapis.com/%s", obj["selfLink"]),
			Type: "google.compute.RegionDisk",
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/compute/v1/rest",
				DiscoveryName:        "RegionDisk",
				Data:                 obj,
			},
		}, nil
	} else {
		return Asset{}, err
	}
}

func GetComputeRegionDiskApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	labelFingerprintProp, err := expandComputeRegionDiskLabelFingerprint(d.Get("label_fingerprint"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("label_fingerprint"); !isEmptyValue(reflect.ValueOf(labelFingerprintProp)) && (ok || !reflect.DeepEqual(v, labelFingerprintProp)) {
		obj["labelFingerprint"] = labelFingerprintProp
	}
	descriptionProp, err := expandComputeRegionDiskDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	labelsProp, err := expandComputeRegionDiskLabels(d.Get("labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	nameProp, err := expandComputeRegionDiskName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	sizeGbProp, err := expandComputeRegionDiskSize(d.Get("size"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("size"); !isEmptyValue(reflect.ValueOf(sizeGbProp)) && (ok || !reflect.DeepEqual(v, sizeGbProp)) {
		obj["sizeGb"] = sizeGbProp
	}
	replicaZonesProp, err := expandComputeRegionDiskReplicaZones(d.Get("replica_zones"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("replica_zones"); !isEmptyValue(reflect.ValueOf(replicaZonesProp)) && (ok || !reflect.DeepEqual(v, replicaZonesProp)) {
		obj["replicaZones"] = replicaZonesProp
	}
	typeProp, err := expandComputeRegionDiskType(d.Get("type"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("type"); !isEmptyValue(reflect.ValueOf(typeProp)) && (ok || !reflect.DeepEqual(v, typeProp)) {
		obj["type"] = typeProp
	}
	regionProp, err := expandComputeRegionDiskRegion(d.Get("region"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("region"); !isEmptyValue(reflect.ValueOf(regionProp)) && (ok || !reflect.DeepEqual(v, regionProp)) {
		obj["region"] = regionProp
	}
	diskEncryptionKeyProp, err := expandComputeRegionDiskDiskEncryptionKey(d.Get("disk_encryption_key"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("disk_encryption_key"); !isEmptyValue(reflect.ValueOf(diskEncryptionKeyProp)) && (ok || !reflect.DeepEqual(v, diskEncryptionKeyProp)) {
		obj["diskEncryptionKey"] = diskEncryptionKeyProp
	}
	sourceSnapshotProp, err := expandComputeRegionDiskSnapshot(d.Get("snapshot"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("snapshot"); !isEmptyValue(reflect.ValueOf(sourceSnapshotProp)) && (ok || !reflect.DeepEqual(v, sourceSnapshotProp)) {
		obj["sourceSnapshot"] = sourceSnapshotProp
	}
	sourceSnapshotEncryptionKeyProp, err := expandComputeRegionDiskSourceSnapshotEncryptionKey(d.Get("source_snapshot_encryption_key"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("source_snapshot_encryption_key"); !isEmptyValue(reflect.ValueOf(sourceSnapshotEncryptionKeyProp)) && (ok || !reflect.DeepEqual(v, sourceSnapshotEncryptionKeyProp)) {
		obj["sourceSnapshotEncryptionKey"] = sourceSnapshotEncryptionKeyProp
	}

	return resourceComputeRegionDiskEncoder(d, config, obj)
}

func resourceComputeRegionDiskEncoder(d TerraformResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return nil, err
	}
	// Get the region
	r, err := getRegion(d, config)
	if err != nil {
		return nil, err
	}
	region, err := config.clientCompute.Regions.Get(project, r).Do()
	if err != nil {
		return nil, err
	}

	if v, ok := d.GetOk("type"); ok {
		log.Printf("[DEBUG] Loading disk type: %s", v.(string))
		diskType, err := readRegionDiskType(config, region, project, v.(string))
		if err != nil {
			return nil, fmt.Errorf(
				"Error loading disk type '%s': %s",
				v.(string), err)
		}

		obj["type"] = diskType.SelfLink
	}

	if v, ok := d.GetOk("image"); ok {
		log.Printf("[DEBUG] Resolving image name: %s", v.(string))
		imageUrl, err := resolveImage(config, project, v.(string))
		if err != nil {
			return nil, fmt.Errorf(
				"Error resolving image name '%s': %s",
				v.(string), err)
		}

		obj["sourceImage"] = imageUrl
		log.Printf("[DEBUG] Image name resolved to: %s", imageUrl)
	}

	return obj, nil
}

func expandComputeRegionDiskLabelFingerprint(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionDiskDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionDiskLabels(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandComputeRegionDiskName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionDiskSize(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionDiskReplicaZones(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		f, err := parseGlobalFieldValue("zones", raw.(string), "project", d, config, true)
		if err != nil {
			return nil, fmt.Errorf("Invalid value for replica_zones: %s", err)
		}
		req = append(req, f.RelativeLink())
	}
	return req, nil
}

func expandComputeRegionDiskType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseRegionalFieldValue("diskTypes", v.(string), "project", "region", "zone", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for type: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeRegionDiskRegion(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseGlobalFieldValue("regions", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for region: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeRegionDiskDiskEncryptionKey(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedRawKey, err := expandComputeRegionDiskDiskEncryptionKeyRawKey(original["raw_key"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRawKey); val.IsValid() && !isEmptyValue(val) {
		transformed["rawKey"] = transformedRawKey
	}

	transformedSha256, err := expandComputeRegionDiskDiskEncryptionKeySha256(original["sha256"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSha256); val.IsValid() && !isEmptyValue(val) {
		transformed["sha256"] = transformedSha256
	}

	return transformed, nil
}

func expandComputeRegionDiskDiskEncryptionKeyRawKey(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionDiskDiskEncryptionKeySha256(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionDiskSnapshot(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseGlobalFieldValue("snapshots", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for snapshot: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeRegionDiskSourceSnapshotEncryptionKey(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedRawKey, err := expandComputeRegionDiskSourceSnapshotEncryptionKeyRawKey(original["raw_key"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRawKey); val.IsValid() && !isEmptyValue(val) {
		transformed["rawKey"] = transformedRawKey
	}

	transformedSha256, err := expandComputeRegionDiskSourceSnapshotEncryptionKeySha256(original["sha256"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSha256); val.IsValid() && !isEmptyValue(val) {
		transformed["sha256"] = transformedSha256
	}

	return transformed, nil
}

func expandComputeRegionDiskSourceSnapshotEncryptionKeyRawKey(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionDiskSourceSnapshotEncryptionKeySha256(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
