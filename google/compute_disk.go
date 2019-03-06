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
	"regexp"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
)

const (
	computeDiskUserRegexString = "^(?:https://www.googleapis.com/compute/v1/projects/)?(" + ProjectRegex + ")/zones/([-_a-zA-Z0-9]*)/instances/([-_a-zA-Z0-9]*)$"
)

var (
	computeDiskUserRegex = regexp.MustCompile(computeDiskUserRegexString)
)

// Is the new disk size smaller than the old one?
func isDiskShrinkage(old, new, _ interface{}) bool {
	// It's okay to remove size entirely.
	if old == nil || new == nil {
		return false
	}
	return new.(int) < old.(int)
}

// We cannot suppress the diff for the case when family name is not part of the image name since we can't
// make a network call in a DiffSuppressFunc.
func diskImageDiffSuppress(_, old, new string, _ *schema.ResourceData) bool {
	// Understand that this function solves a messy problem ("how do we tell if the diff between two images
	// is 'ForceNew-worthy', without making a network call?") in the best way we can: through a series of special
	// cases and regexes.  If you find yourself here because you are trying to add a new special case,
	// you are probably looking for the diskImageFamilyEquals function and its subfunctions.
	// In order to keep this maintainable, we need to ensure that the positive and negative examples
	// in resource_compute_disk_test.go are as complete as possible.

	// 'old' is read from the API.
	// It always has the format 'https://www.googleapis.com/compute/v1/projects/(%s)/global/images/(%s)'
	matches := resolveImageLink.FindStringSubmatch(old)
	if matches == nil {
		// Image read from the API doesn't have the expected format. In practice, it should never happen
		return false
	}
	oldProject := matches[1]
	oldName := matches[2]

	// Partial or full self link family
	if resolveImageProjectFamily.MatchString(new) {
		// Value matches pattern "projects/{project}/global/images/family/{family-name}$"
		matches := resolveImageProjectFamily.FindStringSubmatch(new)
		newProject := matches[1]
		newFamilyName := matches[2]

		return diskImageProjectNameEquals(oldProject, newProject) && diskImageFamilyEquals(oldName, newFamilyName)
	}

	// Partial or full self link image
	if resolveImageProjectImage.MatchString(new) {
		// Value matches pattern "projects/{project}/global/images/{image-name}$"
		matches := resolveImageProjectImage.FindStringSubmatch(new)
		newProject := matches[1]
		newImageName := matches[2]

		return diskImageProjectNameEquals(oldProject, newProject) && diskImageEquals(oldName, newImageName)
	}

	// Partial link without project family
	if resolveImageGlobalFamily.MatchString(new) {
		// Value is "global/images/family/{family-name}"
		matches := resolveImageGlobalFamily.FindStringSubmatch(new)
		familyName := matches[1]

		return diskImageFamilyEquals(oldName, familyName)
	}

	// Partial link without project image
	if resolveImageGlobalImage.MatchString(new) {
		// Value is "global/images/{image-name}"
		matches := resolveImageGlobalImage.FindStringSubmatch(new)
		imageName := matches[1]

		return diskImageEquals(oldName, imageName)
	}

	// Family shorthand
	if resolveImageFamilyFamily.MatchString(new) {
		// Value is "family/{family-name}"
		matches := resolveImageFamilyFamily.FindStringSubmatch(new)
		familyName := matches[1]

		return diskImageFamilyEquals(oldName, familyName)
	}

	// Shorthand for image or family
	if resolveImageProjectImageShorthand.MatchString(new) {
		// Value is "{project}/{image-name}" or "{project}/{family-name}"
		matches := resolveImageProjectImageShorthand.FindStringSubmatch(new)
		newProject := matches[1]
		newName := matches[2]

		return diskImageProjectNameEquals(oldProject, newProject) &&
			(diskImageEquals(oldName, newName) || diskImageFamilyEquals(oldName, newName))
	}

	// Image or family only
	if diskImageEquals(oldName, new) || diskImageFamilyEquals(oldName, new) {
		// Value is "{image-name}" or "{family-name}"
		return true
	}

	return false
}

func diskImageProjectNameEquals(project1, project2 string) bool {
	// Convert short project name to full name
	// For instance, centos => centos-cloud
	fullProjectName, ok := imageMap[project2]
	if ok {
		project2 = fullProjectName
	}

	return project1 == project2
}

func diskImageEquals(oldImageName, newImageName string) bool {
	return oldImageName == newImageName
}

func diskImageFamilyEquals(imageName, familyName string) bool {
	// Handles the case when the image name includes the family name
	// e.g. image name: debian-9-drawfork-v20180109, family name: debian-9
	if strings.Contains(imageName, familyName) {
		return true
	}

	if suppressCanonicalFamilyDiff(imageName, familyName) {
		return true
	}

	if suppressWindowsSqlFamilyDiff(imageName, familyName) {
		return true
	}

	if suppressWindowsFamilyDiff(imageName, familyName) {
		return true
	}

	return false
}

// e.g. image: ubuntu-1404-trusty-v20180122, family: ubuntu-1404-lts
func suppressCanonicalFamilyDiff(imageName, familyName string) bool {
	parts := canonicalUbuntuLtsImage.FindStringSubmatch(imageName)
	if len(parts) == 3 {
		f := fmt.Sprintf("ubuntu-%s%s-lts", parts[1], parts[2])
		if f == familyName {
			return true
		}
	}

	return false
}

// e.g. image: sql-2017-standard-windows-2016-dc-v20180109, family: sql-std-2017-win-2016
// e.g. image: sql-2017-express-windows-2012-r2-dc-v20180109, family: sql-exp-2017-win-2012-r2
func suppressWindowsSqlFamilyDiff(imageName, familyName string) bool {
	parts := windowsSqlImage.FindStringSubmatch(imageName)
	if len(parts) == 5 {
		edition := parts[2] // enterprise, standard or web.
		sqlVersion := parts[1]
		windowsVersion := parts[3]

		// Translate edition
		switch edition {
		case "enterprise":
			edition = "ent"
		case "standard":
			edition = "std"
		case "express":
			edition = "exp"
		}

		var f string
		if revision := parts[4]; revision != "" {
			// With revision
			f = fmt.Sprintf("sql-%s-%s-win-%s-r%s", edition, sqlVersion, windowsVersion, revision)
		} else {
			// No revision
			f = fmt.Sprintf("sql-%s-%s-win-%s", edition, sqlVersion, windowsVersion)
		}

		if f == familyName {
			return true
		}
	}

	return false
}

// e.g. image: windows-server-1709-dc-core-v20180109, family: windows-1709-core
// e.g. image: windows-server-1709-dc-core-for-containers-v20180109, family: "windows-1709-core-for-containers
func suppressWindowsFamilyDiff(imageName, familyName string) bool {
	updatedFamilyString := strings.Replace(familyName, "windows-", "windows-server-", 1)
	updatedImageName := strings.Replace(imageName, "-dc-", "-", 1)

	if strings.Contains(updatedImageName, updatedFamilyString) {
		return true
	}

	return false
}

func GetComputeDiskCaiObject(d TerraformResourceData, config *Config) (Asset, error) {
	if obj, err := GetComputeDiskApiObject(d, config); err == nil {
		return Asset{
			Name: fmt.Sprintf("//compute.googleapis.com/%s", obj["selfLink"]),
			Type: "google.compute.Disk",
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/compute/v1/rest",
				DiscoveryName:        "Disk",
				Data:                 obj,
			},
		}, nil
	} else {
		return Asset{}, err
	}
}

func GetComputeDiskApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	labelFingerprintProp, err := expandComputeDiskLabelFingerprint(d.Get("label_fingerprint"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("label_fingerprint"); !isEmptyValue(reflect.ValueOf(labelFingerprintProp)) && (ok || !reflect.DeepEqual(v, labelFingerprintProp)) {
		obj["labelFingerprint"] = labelFingerprintProp
	}
	descriptionProp, err := expandComputeDiskDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	labelsProp, err := expandComputeDiskLabels(d.Get("labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	nameProp, err := expandComputeDiskName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	sizeGbProp, err := expandComputeDiskSize(d.Get("size"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("size"); !isEmptyValue(reflect.ValueOf(sizeGbProp)) && (ok || !reflect.DeepEqual(v, sizeGbProp)) {
		obj["sizeGb"] = sizeGbProp
	}
	typeProp, err := expandComputeDiskType(d.Get("type"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("type"); !isEmptyValue(reflect.ValueOf(typeProp)) && (ok || !reflect.DeepEqual(v, typeProp)) {
		obj["type"] = typeProp
	}
	sourceImageProp, err := expandComputeDiskImage(d.Get("image"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("image"); !isEmptyValue(reflect.ValueOf(sourceImageProp)) && (ok || !reflect.DeepEqual(v, sourceImageProp)) {
		obj["sourceImage"] = sourceImageProp
	}
	zoneProp, err := expandComputeDiskZone(d.Get("zone"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("zone"); !isEmptyValue(reflect.ValueOf(zoneProp)) && (ok || !reflect.DeepEqual(v, zoneProp)) {
		obj["zone"] = zoneProp
	}
	sourceImageEncryptionKeyProp, err := expandComputeDiskSourceImageEncryptionKey(d.Get("source_image_encryption_key"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("source_image_encryption_key"); !isEmptyValue(reflect.ValueOf(sourceImageEncryptionKeyProp)) && (ok || !reflect.DeepEqual(v, sourceImageEncryptionKeyProp)) {
		obj["sourceImageEncryptionKey"] = sourceImageEncryptionKeyProp
	}
	diskEncryptionKeyProp, err := expandComputeDiskDiskEncryptionKey(d.Get("disk_encryption_key"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("disk_encryption_key"); !isEmptyValue(reflect.ValueOf(diskEncryptionKeyProp)) && (ok || !reflect.DeepEqual(v, diskEncryptionKeyProp)) {
		obj["diskEncryptionKey"] = diskEncryptionKeyProp
	}
	sourceSnapshotProp, err := expandComputeDiskSnapshot(d.Get("snapshot"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("snapshot"); !isEmptyValue(reflect.ValueOf(sourceSnapshotProp)) && (ok || !reflect.DeepEqual(v, sourceSnapshotProp)) {
		obj["sourceSnapshot"] = sourceSnapshotProp
	}
	sourceSnapshotEncryptionKeyProp, err := expandComputeDiskSourceSnapshotEncryptionKey(d.Get("source_snapshot_encryption_key"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("source_snapshot_encryption_key"); !isEmptyValue(reflect.ValueOf(sourceSnapshotEncryptionKeyProp)) && (ok || !reflect.DeepEqual(v, sourceSnapshotEncryptionKeyProp)) {
		obj["sourceSnapshotEncryptionKey"] = sourceSnapshotEncryptionKeyProp
	}

	return resourceComputeDiskEncoder(d, config, obj)
}

func resourceComputeDiskEncoder(d TerraformResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return nil, err
	}
	// Get the zone
	z, err := getZone(d, config)
	if err != nil {
		return nil, err
	}
	zone, err := config.clientCompute.Zones.Get(project, z).Do()
	if err != nil {
		return nil, err
	}

	if v, ok := d.GetOk("type"); ok {
		log.Printf("[DEBUG] Loading disk type: %s", v.(string))
		diskType, err := readDiskType(config, zone, project, v.(string))
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

func expandComputeDiskLabelFingerprint(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeDiskDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeDiskLabels(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandComputeDiskName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeDiskSize(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeDiskType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseZonalFieldValue("diskTypes", v.(string), "project", "zone", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for type: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeDiskImage(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeDiskZone(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseGlobalFieldValue("zones", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for zone: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeDiskSourceImageEncryptionKey(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedRawKey, err := expandComputeDiskSourceImageEncryptionKeyRawKey(original["raw_key"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRawKey); val.IsValid() && !isEmptyValue(val) {
		transformed["rawKey"] = transformedRawKey
	}

	transformedSha256, err := expandComputeDiskSourceImageEncryptionKeySha256(original["sha256"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSha256); val.IsValid() && !isEmptyValue(val) {
		transformed["sha256"] = transformedSha256
	}

	transformedKmsKeySelfLink, err := expandComputeDiskSourceImageEncryptionKeyKmsKeySelfLink(original["kms_key_self_link"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKmsKeySelfLink); val.IsValid() && !isEmptyValue(val) {
		transformed["kmsKeyName"] = transformedKmsKeySelfLink
	}

	return transformed, nil
}

func expandComputeDiskSourceImageEncryptionKeyRawKey(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeDiskSourceImageEncryptionKeySha256(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeDiskSourceImageEncryptionKeyKmsKeySelfLink(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeDiskDiskEncryptionKey(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedRawKey, err := expandComputeDiskDiskEncryptionKeyRawKey(original["raw_key"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRawKey); val.IsValid() && !isEmptyValue(val) {
		transformed["rawKey"] = transformedRawKey
	}

	transformedSha256, err := expandComputeDiskDiskEncryptionKeySha256(original["sha256"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSha256); val.IsValid() && !isEmptyValue(val) {
		transformed["sha256"] = transformedSha256
	}

	transformedKmsKeySelfLink, err := expandComputeDiskDiskEncryptionKeyKmsKeySelfLink(original["kms_key_self_link"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKmsKeySelfLink); val.IsValid() && !isEmptyValue(val) {
		transformed["kmsKeyName"] = transformedKmsKeySelfLink
	}

	return transformed, nil
}

func expandComputeDiskDiskEncryptionKeyRawKey(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeDiskDiskEncryptionKeySha256(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeDiskDiskEncryptionKeyKmsKeySelfLink(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeDiskSnapshot(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseGlobalFieldValue("snapshots", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for snapshot: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeDiskSourceSnapshotEncryptionKey(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedRawKey, err := expandComputeDiskSourceSnapshotEncryptionKeyRawKey(original["raw_key"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRawKey); val.IsValid() && !isEmptyValue(val) {
		transformed["rawKey"] = transformedRawKey
	}

	transformedKmsKeySelfLink, err := expandComputeDiskSourceSnapshotEncryptionKeyKmsKeySelfLink(original["kms_key_self_link"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKmsKeySelfLink); val.IsValid() && !isEmptyValue(val) {
		transformed["kmsKeyName"] = transformedKmsKeySelfLink
	}

	transformedSha256, err := expandComputeDiskSourceSnapshotEncryptionKeySha256(original["sha256"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSha256); val.IsValid() && !isEmptyValue(val) {
		transformed["sha256"] = transformedSha256
	}

	return transformed, nil
}

func expandComputeDiskSourceSnapshotEncryptionKeyRawKey(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeDiskSourceSnapshotEncryptionKeyKmsKeySelfLink(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeDiskSourceSnapshotEncryptionKeySha256(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
