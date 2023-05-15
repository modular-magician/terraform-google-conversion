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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgresource"
	transport_tpg "github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/transport"
)

const ComputeImageAssetType string = "compute.googleapis.com/Image"

func resourceConverterComputeImage() ResourceConverter {
	return ResourceConverter{
		AssetType: ComputeImageAssetType,
		Convert:   GetComputeImageCaiObject,
	}
}

func GetComputeImageCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]Asset, error) {
	name, err := assetName(d, config, "//compute.googleapis.com/projects/{{project}}/global/images/{{name}}")
	if err != nil {
		return []Asset{}, err
	}
	if obj, err := GetComputeImageApiObject(d, config); err == nil {
		return []Asset{{
			Name: name,
			Type: ComputeImageAssetType,
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/compute/v1/rest",
				DiscoveryName:        "Image",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []Asset{}, err
	}
}

func GetComputeImageApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	descriptionProp, err := expandComputeImageDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	storageLocationsProp, err := expandComputeImageStorageLocations(d.Get("storage_locations"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("storage_locations"); !tpgresource.IsEmptyValue(reflect.ValueOf(storageLocationsProp)) && (ok || !reflect.DeepEqual(v, storageLocationsProp)) {
		obj["storageLocations"] = storageLocationsProp
	}
	diskSizeGbProp, err := expandComputeImageDiskSizeGb(d.Get("disk_size_gb"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("disk_size_gb"); !tpgresource.IsEmptyValue(reflect.ValueOf(diskSizeGbProp)) && (ok || !reflect.DeepEqual(v, diskSizeGbProp)) {
		obj["diskSizeGb"] = diskSizeGbProp
	}
	familyProp, err := expandComputeImageFamily(d.Get("family"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("family"); !tpgresource.IsEmptyValue(reflect.ValueOf(familyProp)) && (ok || !reflect.DeepEqual(v, familyProp)) {
		obj["family"] = familyProp
	}
	guestOsFeaturesProp, err := expandComputeImageGuestOsFeatures(d.Get("guest_os_features"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("guest_os_features"); !tpgresource.IsEmptyValue(reflect.ValueOf(guestOsFeaturesProp)) && (ok || !reflect.DeepEqual(v, guestOsFeaturesProp)) {
		obj["guestOsFeatures"] = guestOsFeaturesProp
	}
	imageEncryptionKeyProp, err := expandComputeImageImageEncryptionKey(d.Get("image_encryption_key"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("image_encryption_key"); !tpgresource.IsEmptyValue(reflect.ValueOf(imageEncryptionKeyProp)) && (ok || !reflect.DeepEqual(v, imageEncryptionKeyProp)) {
		obj["imageEncryptionKey"] = imageEncryptionKeyProp
	}
	labelsProp, err := expandComputeImageLabels(d.Get("labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	labelFingerprintProp, err := expandComputeImageLabelFingerprint(d.Get("label_fingerprint"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("label_fingerprint"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelFingerprintProp)) && (ok || !reflect.DeepEqual(v, labelFingerprintProp)) {
		obj["labelFingerprint"] = labelFingerprintProp
	}
	licensesProp, err := expandComputeImageLicenses(d.Get("licenses"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("licenses"); !tpgresource.IsEmptyValue(reflect.ValueOf(licensesProp)) && (ok || !reflect.DeepEqual(v, licensesProp)) {
		obj["licenses"] = licensesProp
	}
	nameProp, err := expandComputeImageName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	rawDiskProp, err := expandComputeImageRawDisk(d.Get("raw_disk"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("raw_disk"); !tpgresource.IsEmptyValue(reflect.ValueOf(rawDiskProp)) && (ok || !reflect.DeepEqual(v, rawDiskProp)) {
		obj["rawDisk"] = rawDiskProp
	}
	sourceDiskProp, err := expandComputeImageSourceDisk(d.Get("source_disk"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("source_disk"); !tpgresource.IsEmptyValue(reflect.ValueOf(sourceDiskProp)) && (ok || !reflect.DeepEqual(v, sourceDiskProp)) {
		obj["sourceDisk"] = sourceDiskProp
	}
	sourceImageProp, err := expandComputeImageSourceImage(d.Get("source_image"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("source_image"); !tpgresource.IsEmptyValue(reflect.ValueOf(sourceImageProp)) && (ok || !reflect.DeepEqual(v, sourceImageProp)) {
		obj["sourceImage"] = sourceImageProp
	}
	sourceSnapshotProp, err := expandComputeImageSourceSnapshot(d.Get("source_snapshot"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("source_snapshot"); !tpgresource.IsEmptyValue(reflect.ValueOf(sourceSnapshotProp)) && (ok || !reflect.DeepEqual(v, sourceSnapshotProp)) {
		obj["sourceSnapshot"] = sourceSnapshotProp
	}

	return obj, nil
}

func expandComputeImageDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeImageStorageLocations(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeImageDiskSizeGb(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeImageFamily(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeImageGuestOsFeatures(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedType, err := expandComputeImageGuestOsFeaturesType(original["type"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedType); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["type"] = transformedType
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandComputeImageGuestOsFeaturesType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeImageImageEncryptionKey(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedKmsKeySelfLink, err := expandComputeImageImageEncryptionKeyKmsKeySelfLink(original["kms_key_self_link"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKmsKeySelfLink); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["kmsKeyName"] = transformedKmsKeySelfLink
	}

	transformedKmsKeyServiceAccount, err := expandComputeImageImageEncryptionKeyKmsKeyServiceAccount(original["kms_key_service_account"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKmsKeyServiceAccount); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["kmsKeyServiceAccount"] = transformedKmsKeyServiceAccount
	}

	return transformed, nil
}

func expandComputeImageImageEncryptionKeyKmsKeySelfLink(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeImageImageEncryptionKeyKmsKeyServiceAccount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeImageLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandComputeImageLabelFingerprint(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeImageLicenses(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			return nil, fmt.Errorf("Invalid value for licenses: nil")
		}
		f, err := tpgresource.ParseGlobalFieldValue("licenses", raw.(string), "project", d, config, true)
		if err != nil {
			return nil, fmt.Errorf("Invalid value for licenses: %s", err)
		}
		req = append(req, f.RelativeLink())
	}
	return req, nil
}

func expandComputeImageName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeImageRawDisk(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedContainerType, err := expandComputeImageRawDiskContainerType(original["container_type"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedContainerType); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["containerType"] = transformedContainerType
	}

	transformedSha1, err := expandComputeImageRawDiskSha1(original["sha1"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSha1); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["sha1Checksum"] = transformedSha1
	}

	transformedSource, err := expandComputeImageRawDiskSource(original["source"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSource); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["source"] = transformedSource
	}

	return transformed, nil
}

func expandComputeImageRawDiskContainerType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeImageRawDiskSha1(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeImageRawDiskSource(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeImageSourceDisk(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	f, err := tpgresource.ParseZonalFieldValue("disks", v.(string), "project", "zone", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for source_disk: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeImageSourceImage(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	f, err := tpgresource.ParseGlobalFieldValue("images", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for source_image: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeImageSourceSnapshot(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	f, err := tpgresource.ParseGlobalFieldValue("snapshots", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for source_snapshot: %s", err)
	}
	return f.RelativeLink(), nil
}
