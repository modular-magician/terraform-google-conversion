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
	"reflect"
)

func GetComputeDiskResourcePolicyAttachmentCaiObject(d TerraformResourceData, config *Config) (Asset, error) {
	name, err := assetName(d, config, "//compute.googleapis.com/projects/{{project}}/zones/{{zone}}/disks/{{disk}}")
	if err != nil {
		return Asset{}, err
	}
	if obj, err := GetComputeDiskResourcePolicyAttachmentApiObject(d, config); err == nil {
		return Asset{
			Name: name,
			Type: "compute.googleapis.com/DiskResourcePolicyAttachment",
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/compute/v1/rest",
				DiscoveryName:        "DiskResourcePolicyAttachment",
				Data:                 obj,
			},
		}, nil
	} else {
		return Asset{}, err
	}
}

func GetComputeDiskResourcePolicyAttachmentApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandComputeDiskResourcePolicyAttachmentName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}

	return resourceComputeDiskResourcePolicyAttachmentEncoder(d, config, obj)
}

func resourceComputeDiskResourcePolicyAttachmentEncoder(d TerraformResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	config := meta.(*Config)
	project, err := getProject(d, config)
	if err != nil {
		return nil, err
	}

	region := getRegionFromZone(d.Get("zone").(string))

	obj["resourcePolicies"] = []interface{}{fmt.Sprintf("projects/%s/regions/%s/resourcePolicies/%s", project, region, obj["name"])}
	delete(obj, "name")
	return obj, nil
}

func expandComputeDiskResourcePolicyAttachmentName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
