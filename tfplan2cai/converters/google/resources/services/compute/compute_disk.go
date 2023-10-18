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

package compute

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

// diffsupress for beta and to check change in source_disk attribute
func sourceDiskDiffSupress(_, old, new string, _ *schema.ResourceData) bool {
	s1 := strings.TrimPrefix(old, "https://www.googleapis.com/compute/beta")
	s2 := strings.TrimPrefix(new, "https://www.googleapis.com/compute/v1")
	if strings.HasSuffix(s1, s2) {
		return true
	}
	return false
}

// Is the new disk size smaller than the old one?
func IsDiskShrinkage(_ context.Context, old, new, _ interface{}) bool {
	// It's okay to remove size entirely.
	if old == nil || new == nil {
		return false
	}
	return new.(int) < old.(int)
}

// We cannot suppress the diff for the case when family name is not part of the image name since we can't
// make a network call in a DiffSuppressFunc.
func DiskImageDiffSuppress(_, old, new string, _ *schema.ResourceData) bool {
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
	fullProjectName, ok := ImageMap[project2]
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
	// e.g. image name: debian-11-bullseye-v20220719, family name: debian-11

	// First condition is to check if image contains arm64 because of case like:
	// image name: opensuse-leap-15-4-v20220713-arm64, family name: opensuse-leap (should not be evaluated during handling of amd64 cases)
	// In second condition, we have to check for amd64 because of cases like:
	// image name: ubuntu-2210-kinetic-amd64-v20221022, family name: ubuntu-2210 (should not suppress)
	if !strings.Contains(imageName, "-arm64") && strings.Contains(imageName, strings.TrimSuffix(familyName, "-amd64")) {
		if strings.Contains(imageName, "-amd64") {
			return strings.HasSuffix(familyName, "-amd64")
		} else {
			return !strings.HasSuffix(familyName, "-amd64")
		}
	}

	// We have to check for arm64 because of cases like:
	// image name: opensuse-leap-15-4-v20220713-arm64, family name: opensuse-leap (should not suppress)
	if strings.Contains(imageName, strings.TrimSuffix(familyName, "-arm64")) {
		if strings.Contains(imageName, "-arm64") {
			return strings.HasSuffix(familyName, "-arm64")
		} else {
			return !strings.HasSuffix(familyName, "-arm64")
		}
	}

	if suppressCanonicalFamilyDiff(imageName, familyName) {
		return true
	}

	if suppressCosFamilyDiff(imageName, familyName) {
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
	if len(parts) == 4 {
		var f string
		if parts[3] == "" {
			f = fmt.Sprintf("ubuntu-%s%s-lts", parts[1], parts[2])
		} else {
			f = fmt.Sprintf("ubuntu-%s%s-lts-%s", parts[1], parts[2], parts[3])
		}
		if f == familyName {
			return true
		}
	}

	return false
}

// e.g. image: cos-NN-*, family: cos-NN-lts
func suppressCosFamilyDiff(imageName, familyName string) bool {
	parts := cosLtsImage.FindStringSubmatch(imageName)
	if len(parts) == 2 {
		f := fmt.Sprintf("cos-%s-lts", parts[1])
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

	return strings.Contains(updatedImageName, updatedFamilyString)
}

const ComputeDiskAssetType string = "compute.googleapis.com/Disk"

func ResourceConverterComputeDisk() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: ComputeDiskAssetType,
		Convert:   GetComputeDiskCaiObject,
	}
}

func GetComputeDiskCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//compute.googleapis.com/projects/{{project}}/zones/{{zone}}/disks/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetComputeDiskApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: ComputeDiskAssetType,
			Resource: &cai.AssetResource{
				Version:              "beta",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/compute/beta/rest",
				DiscoveryName:        "Disk",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetComputeDiskApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	labelFingerprintProp, err := expandComputeDiskLabelFingerprint(d.Get("label_fingerprint"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("label_fingerprint"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelFingerprintProp)) && (ok || !reflect.DeepEqual(v, labelFingerprintProp)) {
		obj["labelFingerprint"] = labelFingerprintProp
	}
	descriptionProp, err := expandComputeDiskDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	nameProp, err := expandComputeDiskName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	sizeGbProp, err := expandComputeDiskSize(d.Get("size"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("size"); !tpgresource.IsEmptyValue(reflect.ValueOf(sizeGbProp)) && (ok || !reflect.DeepEqual(v, sizeGbProp)) {
		obj["sizeGb"] = sizeGbProp
	}
	physicalBlockSizeBytesProp, err := expandComputeDiskPhysicalBlockSizeBytes(d.Get("physical_block_size_bytes"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("physical_block_size_bytes"); !tpgresource.IsEmptyValue(reflect.ValueOf(physicalBlockSizeBytesProp)) && (ok || !reflect.DeepEqual(v, physicalBlockSizeBytesProp)) {
		obj["physicalBlockSizeBytes"] = physicalBlockSizeBytesProp
	}
	sourceDiskProp, err := expandComputeDiskSourceDisk(d.Get("source_disk"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("source_disk"); !tpgresource.IsEmptyValue(reflect.ValueOf(sourceDiskProp)) && (ok || !reflect.DeepEqual(v, sourceDiskProp)) {
		obj["sourceDisk"] = sourceDiskProp
	}
	typeProp, err := expandComputeDiskType(d.Get("type"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("type"); !tpgresource.IsEmptyValue(reflect.ValueOf(typeProp)) && (ok || !reflect.DeepEqual(v, typeProp)) {
		obj["type"] = typeProp
	}
	sourceImageProp, err := expandComputeDiskImage(d.Get("image"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("image"); !tpgresource.IsEmptyValue(reflect.ValueOf(sourceImageProp)) && (ok || !reflect.DeepEqual(v, sourceImageProp)) {
		obj["sourceImage"] = sourceImageProp
	}
	resourcePoliciesProp, err := expandComputeDiskResourcePolicies(d.Get("resource_policies"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("resource_policies"); !tpgresource.IsEmptyValue(reflect.ValueOf(resourcePoliciesProp)) && (ok || !reflect.DeepEqual(v, resourcePoliciesProp)) {
		obj["resourcePolicies"] = resourcePoliciesProp
	}
	enableConfidentialComputeProp, err := expandComputeDiskEnableConfidentialCompute(d.Get("enable_confidential_compute"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("enable_confidential_compute"); !tpgresource.IsEmptyValue(reflect.ValueOf(enableConfidentialComputeProp)) && (ok || !reflect.DeepEqual(v, enableConfidentialComputeProp)) {
		obj["enableConfidentialCompute"] = enableConfidentialComputeProp
	}
	multiWriterProp, err := expandComputeDiskMultiWriter(d.Get("multi_writer"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("multi_writer"); !tpgresource.IsEmptyValue(reflect.ValueOf(multiWriterProp)) && (ok || !reflect.DeepEqual(v, multiWriterProp)) {
		obj["multiWriter"] = multiWriterProp
	}
	provisionedIopsProp, err := expandComputeDiskProvisionedIops(d.Get("provisioned_iops"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("provisioned_iops"); !tpgresource.IsEmptyValue(reflect.ValueOf(provisionedIopsProp)) && (ok || !reflect.DeepEqual(v, provisionedIopsProp)) {
		obj["provisionedIops"] = provisionedIopsProp
	}
	provisionedThroughputProp, err := expandComputeDiskProvisionedThroughput(d.Get("provisioned_throughput"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("provisioned_throughput"); !tpgresource.IsEmptyValue(reflect.ValueOf(provisionedThroughputProp)) && (ok || !reflect.DeepEqual(v, provisionedThroughputProp)) {
		obj["provisionedThroughput"] = provisionedThroughputProp
	}
	asyncPrimaryDiskProp, err := expandComputeDiskAsyncPrimaryDisk(d.Get("async_primary_disk"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("async_primary_disk"); !tpgresource.IsEmptyValue(reflect.ValueOf(asyncPrimaryDiskProp)) && (ok || !reflect.DeepEqual(v, asyncPrimaryDiskProp)) {
		obj["asyncPrimaryDisk"] = asyncPrimaryDiskProp
	}
	guestOsFeaturesProp, err := expandComputeDiskGuestOsFeatures(d.Get("guest_os_features"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("guest_os_features"); !tpgresource.IsEmptyValue(reflect.ValueOf(guestOsFeaturesProp)) && (ok || !reflect.DeepEqual(v, guestOsFeaturesProp)) {
		obj["guestOsFeatures"] = guestOsFeaturesProp
	}
	licensesProp, err := expandComputeDiskLicenses(d.Get("licenses"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("licenses"); !tpgresource.IsEmptyValue(reflect.ValueOf(licensesProp)) && (ok || !reflect.DeepEqual(v, licensesProp)) {
		obj["licenses"] = licensesProp
	}
	labelsProp, err := expandComputeDiskEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	zoneProp, err := expandComputeDiskZone(d.Get("zone"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("zone"); !tpgresource.IsEmptyValue(reflect.ValueOf(zoneProp)) && (ok || !reflect.DeepEqual(v, zoneProp)) {
		obj["zone"] = zoneProp
	}
	sourceImageEncryptionKeyProp, err := expandComputeDiskSourceImageEncryptionKey(d.Get("source_image_encryption_key"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("source_image_encryption_key"); !tpgresource.IsEmptyValue(reflect.ValueOf(sourceImageEncryptionKeyProp)) && (ok || !reflect.DeepEqual(v, sourceImageEncryptionKeyProp)) {
		obj["sourceImageEncryptionKey"] = sourceImageEncryptionKeyProp
	}
	diskEncryptionKeyProp, err := expandComputeDiskDiskEncryptionKey(d.Get("disk_encryption_key"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("disk_encryption_key"); !tpgresource.IsEmptyValue(reflect.ValueOf(diskEncryptionKeyProp)) && (ok || !reflect.DeepEqual(v, diskEncryptionKeyProp)) {
		obj["diskEncryptionKey"] = diskEncryptionKeyProp
	}
	sourceSnapshotProp, err := expandComputeDiskSnapshot(d.Get("snapshot"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("snapshot"); !tpgresource.IsEmptyValue(reflect.ValueOf(sourceSnapshotProp)) && (ok || !reflect.DeepEqual(v, sourceSnapshotProp)) {
		obj["sourceSnapshot"] = sourceSnapshotProp
	}
	sourceSnapshotEncryptionKeyProp, err := expandComputeDiskSourceSnapshotEncryptionKey(d.Get("source_snapshot_encryption_key"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("source_snapshot_encryption_key"); !tpgresource.IsEmptyValue(reflect.ValueOf(sourceSnapshotEncryptionKeyProp)) && (ok || !reflect.DeepEqual(v, sourceSnapshotEncryptionKeyProp)) {
		obj["sourceSnapshotEncryptionKey"] = sourceSnapshotEncryptionKeyProp
	}

	return resourceComputeDiskEncoder(d, config, obj)
}

func resourceComputeDiskEncoder(d tpgresource.TerraformResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	config := meta.(*transport_tpg.Config)

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return nil, err
	}

	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return nil, err
	}

	if v, ok := d.GetOk("type"); ok {
		log.Printf("[DEBUG] Loading disk type: %s", v.(string))
		diskType, err := readDiskType(config, d, v.(string))
		if err != nil {
			return nil, fmt.Errorf(
				"Error loading disk type '%s': %s",
				v.(string), err)
		}

		obj["type"] = diskType.RelativeLink()
	}

	if v, ok := d.GetOk("image"); ok {
		log.Printf("[DEBUG] Resolving image name: %s", v.(string))
		imageUrl, err := ResolveImage(config, project, v.(string), userAgent)
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

func expandComputeDiskLabelFingerprint(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeDiskDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeDiskName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeDiskSize(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeDiskPhysicalBlockSizeBytes(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeDiskSourceDisk(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeDiskType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	f, err := tpgresource.ParseZonalFieldValue("diskTypes", v.(string), "project", "zone", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for type: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeDiskImage(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeDiskResourcePolicies(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			return nil, fmt.Errorf("Invalid value for resource_policies: nil")
		}
		f, err := tpgresource.ParseRegionalFieldValue("resourcePolicies", raw.(string), "project", "region", "zone", d, config, true)
		if err != nil {
			return nil, fmt.Errorf("Invalid value for resource_policies: %s", err)
		}
		req = append(req, f.RelativeLink())
	}
	return req, nil
}

func expandComputeDiskEnableConfidentialCompute(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeDiskMultiWriter(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeDiskProvisionedIops(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeDiskProvisionedThroughput(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeDiskAsyncPrimaryDisk(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedDisk, err := expandComputeDiskAsyncPrimaryDiskDisk(original["disk"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDisk); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["disk"] = transformedDisk
	}

	return transformed, nil
}

func expandComputeDiskAsyncPrimaryDiskDisk(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeDiskGuestOsFeatures(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedType, err := expandComputeDiskGuestOsFeaturesType(original["type"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedType); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["type"] = transformedType
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandComputeDiskGuestOsFeaturesType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeDiskLicenses(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
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

func expandComputeDiskEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandComputeDiskZone(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	f, err := tpgresource.ParseGlobalFieldValue("zones", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for zone: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeDiskSourceImageEncryptionKey(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
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
	} else if val := reflect.ValueOf(transformedRawKey); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["rawKey"] = transformedRawKey
	}

	transformedKmsKeySelfLink, err := expandComputeDiskSourceImageEncryptionKeyKmsKeySelfLink(original["kms_key_self_link"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKmsKeySelfLink); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["kmsKeyName"] = transformedKmsKeySelfLink
	}

	transformedKmsKeyServiceAccount, err := expandComputeDiskSourceImageEncryptionKeyKmsKeyServiceAccount(original["kms_key_service_account"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKmsKeyServiceAccount); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["kmsKeyServiceAccount"] = transformedKmsKeyServiceAccount
	}

	return transformed, nil
}

func expandComputeDiskSourceImageEncryptionKeyRawKey(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeDiskSourceImageEncryptionKeyKmsKeySelfLink(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeDiskSourceImageEncryptionKeyKmsKeyServiceAccount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeDiskDiskEncryptionKey(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
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
	} else if val := reflect.ValueOf(transformedRawKey); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["rawKey"] = transformedRawKey
	}

	transformedRsaEncryptedKey, err := expandComputeDiskDiskEncryptionKeyRsaEncryptedKey(original["rsa_encrypted_key"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRsaEncryptedKey); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["rsaEncryptedKey"] = transformedRsaEncryptedKey
	}

	transformedKmsKeySelfLink, err := expandComputeDiskDiskEncryptionKeyKmsKeySelfLink(original["kms_key_self_link"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKmsKeySelfLink); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["kmsKeyName"] = transformedKmsKeySelfLink
	}

	transformedKmsKeyServiceAccount, err := expandComputeDiskDiskEncryptionKeyKmsKeyServiceAccount(original["kms_key_service_account"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKmsKeyServiceAccount); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["kmsKeyServiceAccount"] = transformedKmsKeyServiceAccount
	}

	return transformed, nil
}

func expandComputeDiskDiskEncryptionKeyRawKey(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeDiskDiskEncryptionKeyRsaEncryptedKey(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeDiskDiskEncryptionKeyKmsKeySelfLink(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeDiskDiskEncryptionKeyKmsKeyServiceAccount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeDiskSnapshot(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	f, err := tpgresource.ParseGlobalFieldValue("snapshots", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for snapshot: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeDiskSourceSnapshotEncryptionKey(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
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
	} else if val := reflect.ValueOf(transformedRawKey); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["rawKey"] = transformedRawKey
	}

	transformedKmsKeySelfLink, err := expandComputeDiskSourceSnapshotEncryptionKeyKmsKeySelfLink(original["kms_key_self_link"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKmsKeySelfLink); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["kmsKeyName"] = transformedKmsKeySelfLink
	}

	transformedKmsKeyServiceAccount, err := expandComputeDiskSourceSnapshotEncryptionKeyKmsKeyServiceAccount(original["kms_key_service_account"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKmsKeyServiceAccount); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["kmsKeyServiceAccount"] = transformedKmsKeyServiceAccount
	}

	return transformed, nil
}

func expandComputeDiskSourceSnapshotEncryptionKeyRawKey(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeDiskSourceSnapshotEncryptionKeyKmsKeySelfLink(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeDiskSourceSnapshotEncryptionKeyKmsKeyServiceAccount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
