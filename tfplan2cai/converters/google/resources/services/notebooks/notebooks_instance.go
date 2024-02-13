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

package notebooks

import (
	"fmt"
	"reflect"
	"sort"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

var NotebooksInstanceProvidedScopes = []string{
	"https://www.googleapis.com/auth/cloud-platform",
	"https://www.googleapis.com/auth/userinfo.email",
}

func NotebooksInstanceScopesDiffSuppress(_, _, _ string, d *schema.ResourceData) bool {
	old, new := d.GetChange("service_account_scopes")
	oldValue := old.([]interface{})
	newValue := new.([]interface{})
	oldValueList := []string{}
	newValueList := []string{}

	for _, item := range oldValue {
		oldValueList = append(oldValueList, item.(string))
	}

	for _, item := range newValue {
		newValueList = append(newValueList, item.(string))
	}
	newValueList = append(newValueList, NotebooksInstanceProvidedScopes...)

	sort.Strings(oldValueList)
	sort.Strings(newValueList)
	if reflect.DeepEqual(oldValueList, newValueList) {
		return true
	}
	return false
}

func NotebooksInstanceKmsDiffSuppress(_, old, new string, _ *schema.ResourceData) bool {
	if strings.HasPrefix(old, new) {
		return true
	}
	return false
}

func modifyNotebooksInstanceState(config *transport_tpg.Config, d *schema.ResourceData, project string, billingProject string, userAgent string, state string) (map[string]interface{}, error) {
	url, err := tpgresource.ReplaceVars(d, config, "{{NotebooksBasePath}}projects/{{project}}/locations/{{location}}/instances/{{name}}:"+state)
	if err != nil {
		return nil, err
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "POST",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
	})
	if err != nil {
		return nil, fmt.Errorf("Unable to %q google_notebooks_instance %q: %s", state, d.Id(), err)
	}
	return res, nil
}

const NotebooksInstanceAssetType string = "notebooks.googleapis.com/Instance"

func ResourceConverterNotebooksInstance() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: NotebooksInstanceAssetType,
		Convert:   GetNotebooksInstanceCaiObject,
	}
}

func GetNotebooksInstanceCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//notebooks.googleapis.com/projects/{{project}}/locations/{{location}}/instances/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetNotebooksInstanceApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: NotebooksInstanceAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/notebooks/v1/rest",
				DiscoveryName:        "Instance",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetNotebooksInstanceApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	machineTypeProp, err := expandNotebooksInstanceMachineType(d.Get("machine_type"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("machine_type"); !tpgresource.IsEmptyValue(reflect.ValueOf(machineTypeProp)) && (ok || !reflect.DeepEqual(v, machineTypeProp)) {
		obj["machineType"] = machineTypeProp
	}
	postStartupScriptProp, err := expandNotebooksInstancePostStartupScript(d.Get("post_startup_script"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("post_startup_script"); !tpgresource.IsEmptyValue(reflect.ValueOf(postStartupScriptProp)) && (ok || !reflect.DeepEqual(v, postStartupScriptProp)) {
		obj["postStartupScript"] = postStartupScriptProp
	}
	instanceOwnersProp, err := expandNotebooksInstanceInstanceOwners(d.Get("instance_owners"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("instance_owners"); !tpgresource.IsEmptyValue(reflect.ValueOf(instanceOwnersProp)) && (ok || !reflect.DeepEqual(v, instanceOwnersProp)) {
		obj["instanceOwners"] = instanceOwnersProp
	}
	serviceAccountProp, err := expandNotebooksInstanceServiceAccount(d.Get("service_account"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("service_account"); !tpgresource.IsEmptyValue(reflect.ValueOf(serviceAccountProp)) && (ok || !reflect.DeepEqual(v, serviceAccountProp)) {
		obj["serviceAccount"] = serviceAccountProp
	}
	serviceAccountScopesProp, err := expandNotebooksInstanceServiceAccountScopes(d.Get("service_account_scopes"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("service_account_scopes"); !tpgresource.IsEmptyValue(reflect.ValueOf(serviceAccountScopesProp)) && (ok || !reflect.DeepEqual(v, serviceAccountScopesProp)) {
		obj["serviceAccountScopes"] = serviceAccountScopesProp
	}
	acceleratorConfigProp, err := expandNotebooksInstanceAcceleratorConfig(d.Get("accelerator_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("accelerator_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(acceleratorConfigProp)) && (ok || !reflect.DeepEqual(v, acceleratorConfigProp)) {
		obj["acceleratorConfig"] = acceleratorConfigProp
	}
	shieldedInstanceConfigProp, err := expandNotebooksInstanceShieldedInstanceConfig(d.Get("shielded_instance_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("shielded_instance_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(shieldedInstanceConfigProp)) && (ok || !reflect.DeepEqual(v, shieldedInstanceConfigProp)) {
		obj["shieldedInstanceConfig"] = shieldedInstanceConfigProp
	}
	nicTypeProp, err := expandNotebooksInstanceNicType(d.Get("nic_type"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("nic_type"); !tpgresource.IsEmptyValue(reflect.ValueOf(nicTypeProp)) && (ok || !reflect.DeepEqual(v, nicTypeProp)) {
		obj["nicType"] = nicTypeProp
	}
	reservationAffinityProp, err := expandNotebooksInstanceReservationAffinity(d.Get("reservation_affinity"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("reservation_affinity"); !tpgresource.IsEmptyValue(reflect.ValueOf(reservationAffinityProp)) && (ok || !reflect.DeepEqual(v, reservationAffinityProp)) {
		obj["reservationAffinity"] = reservationAffinityProp
	}
	installGpuDriverProp, err := expandNotebooksInstanceInstallGpuDriver(d.Get("install_gpu_driver"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("install_gpu_driver"); !tpgresource.IsEmptyValue(reflect.ValueOf(installGpuDriverProp)) && (ok || !reflect.DeepEqual(v, installGpuDriverProp)) {
		obj["installGpuDriver"] = installGpuDriverProp
	}
	customGpuDriverPathProp, err := expandNotebooksInstanceCustomGpuDriverPath(d.Get("custom_gpu_driver_path"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("custom_gpu_driver_path"); !tpgresource.IsEmptyValue(reflect.ValueOf(customGpuDriverPathProp)) && (ok || !reflect.DeepEqual(v, customGpuDriverPathProp)) {
		obj["customGpuDriverPath"] = customGpuDriverPathProp
	}
	bootDiskTypeProp, err := expandNotebooksInstanceBootDiskType(d.Get("boot_disk_type"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("boot_disk_type"); !tpgresource.IsEmptyValue(reflect.ValueOf(bootDiskTypeProp)) && (ok || !reflect.DeepEqual(v, bootDiskTypeProp)) {
		obj["bootDiskType"] = bootDiskTypeProp
	}
	bootDiskSizeGbProp, err := expandNotebooksInstanceBootDiskSizeGb(d.Get("boot_disk_size_gb"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("boot_disk_size_gb"); !tpgresource.IsEmptyValue(reflect.ValueOf(bootDiskSizeGbProp)) && (ok || !reflect.DeepEqual(v, bootDiskSizeGbProp)) {
		obj["bootDiskSizeGb"] = bootDiskSizeGbProp
	}
	dataDiskTypeProp, err := expandNotebooksInstanceDataDiskType(d.Get("data_disk_type"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("data_disk_type"); !tpgresource.IsEmptyValue(reflect.ValueOf(dataDiskTypeProp)) && (ok || !reflect.DeepEqual(v, dataDiskTypeProp)) {
		obj["dataDiskType"] = dataDiskTypeProp
	}
	dataDiskSizeGbProp, err := expandNotebooksInstanceDataDiskSizeGb(d.Get("data_disk_size_gb"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("data_disk_size_gb"); !tpgresource.IsEmptyValue(reflect.ValueOf(dataDiskSizeGbProp)) && (ok || !reflect.DeepEqual(v, dataDiskSizeGbProp)) {
		obj["dataDiskSizeGb"] = dataDiskSizeGbProp
	}
	noRemoveDataDiskProp, err := expandNotebooksInstanceNoRemoveDataDisk(d.Get("no_remove_data_disk"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("no_remove_data_disk"); !tpgresource.IsEmptyValue(reflect.ValueOf(noRemoveDataDiskProp)) && (ok || !reflect.DeepEqual(v, noRemoveDataDiskProp)) {
		obj["noRemoveDataDisk"] = noRemoveDataDiskProp
	}
	diskEncryptionProp, err := expandNotebooksInstanceDiskEncryption(d.Get("disk_encryption"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("disk_encryption"); !tpgresource.IsEmptyValue(reflect.ValueOf(diskEncryptionProp)) && (ok || !reflect.DeepEqual(v, diskEncryptionProp)) {
		obj["diskEncryption"] = diskEncryptionProp
	}
	kmsKeyProp, err := expandNotebooksInstanceKmsKey(d.Get("kms_key"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("kms_key"); !tpgresource.IsEmptyValue(reflect.ValueOf(kmsKeyProp)) && (ok || !reflect.DeepEqual(v, kmsKeyProp)) {
		obj["kmsKey"] = kmsKeyProp
	}
	noPublicIpProp, err := expandNotebooksInstanceNoPublicIp(d.Get("no_public_ip"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("no_public_ip"); !tpgresource.IsEmptyValue(reflect.ValueOf(noPublicIpProp)) && (ok || !reflect.DeepEqual(v, noPublicIpProp)) {
		obj["noPublicIp"] = noPublicIpProp
	}
	noProxyAccessProp, err := expandNotebooksInstanceNoProxyAccess(d.Get("no_proxy_access"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("no_proxy_access"); !tpgresource.IsEmptyValue(reflect.ValueOf(noProxyAccessProp)) && (ok || !reflect.DeepEqual(v, noProxyAccessProp)) {
		obj["noProxyAccess"] = noProxyAccessProp
	}
	networkProp, err := expandNotebooksInstanceNetwork(d.Get("network"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("network"); !tpgresource.IsEmptyValue(reflect.ValueOf(networkProp)) && (ok || !reflect.DeepEqual(v, networkProp)) {
		obj["network"] = networkProp
	}
	subnetProp, err := expandNotebooksInstanceSubnet(d.Get("subnet"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("subnet"); !tpgresource.IsEmptyValue(reflect.ValueOf(subnetProp)) && (ok || !reflect.DeepEqual(v, subnetProp)) {
		obj["subnet"] = subnetProp
	}
	tagsProp, err := expandNotebooksInstanceTags(d.Get("tags"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("tags"); !tpgresource.IsEmptyValue(reflect.ValueOf(tagsProp)) && (ok || !reflect.DeepEqual(v, tagsProp)) {
		obj["tags"] = tagsProp
	}
	metadataProp, err := expandNotebooksInstanceMetadata(d.Get("metadata"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("metadata"); !tpgresource.IsEmptyValue(reflect.ValueOf(metadataProp)) && (ok || !reflect.DeepEqual(v, metadataProp)) {
		obj["metadata"] = metadataProp
	}
	vmImageProp, err := expandNotebooksInstanceVmImage(d.Get("vm_image"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("vm_image"); !tpgresource.IsEmptyValue(reflect.ValueOf(vmImageProp)) && (ok || !reflect.DeepEqual(v, vmImageProp)) {
		obj["vmImage"] = vmImageProp
	}
	containerImageProp, err := expandNotebooksInstanceContainerImage(d.Get("container_image"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("container_image"); !tpgresource.IsEmptyValue(reflect.ValueOf(containerImageProp)) && (ok || !reflect.DeepEqual(v, containerImageProp)) {
		obj["containerImage"] = containerImageProp
	}
	labelsProp, err := expandNotebooksInstanceEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	return obj, nil
}

func expandNotebooksInstanceMachineType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNotebooksInstancePostStartupScript(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNotebooksInstanceInstanceOwners(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNotebooksInstanceServiceAccount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNotebooksInstanceServiceAccountScopes(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNotebooksInstanceAcceleratorConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedType, err := expandNotebooksInstanceAcceleratorConfigType(original["type"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedType); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["type"] = transformedType
	}

	transformedCoreCount, err := expandNotebooksInstanceAcceleratorConfigCoreCount(original["core_count"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCoreCount); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["coreCount"] = transformedCoreCount
	}

	return transformed, nil
}

func expandNotebooksInstanceAcceleratorConfigType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNotebooksInstanceAcceleratorConfigCoreCount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNotebooksInstanceShieldedInstanceConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedEnableIntegrityMonitoring, err := expandNotebooksInstanceShieldedInstanceConfigEnableIntegrityMonitoring(original["enable_integrity_monitoring"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEnableIntegrityMonitoring); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["enableIntegrityMonitoring"] = transformedEnableIntegrityMonitoring
	}

	transformedEnableSecureBoot, err := expandNotebooksInstanceShieldedInstanceConfigEnableSecureBoot(original["enable_secure_boot"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEnableSecureBoot); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["enableSecureBoot"] = transformedEnableSecureBoot
	}

	transformedEnableVtpm, err := expandNotebooksInstanceShieldedInstanceConfigEnableVtpm(original["enable_vtpm"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEnableVtpm); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["enableVtpm"] = transformedEnableVtpm
	}

	return transformed, nil
}

func expandNotebooksInstanceShieldedInstanceConfigEnableIntegrityMonitoring(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNotebooksInstanceShieldedInstanceConfigEnableSecureBoot(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNotebooksInstanceShieldedInstanceConfigEnableVtpm(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNotebooksInstanceNicType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNotebooksInstanceReservationAffinity(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedConsumeReservationType, err := expandNotebooksInstanceReservationAffinityConsumeReservationType(original["consume_reservation_type"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedConsumeReservationType); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["consumeReservationType"] = transformedConsumeReservationType
	}

	transformedKey, err := expandNotebooksInstanceReservationAffinityKey(original["key"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKey); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["key"] = transformedKey
	}

	transformedValues, err := expandNotebooksInstanceReservationAffinityValues(original["values"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedValues); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["values"] = transformedValues
	}

	return transformed, nil
}

func expandNotebooksInstanceReservationAffinityConsumeReservationType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNotebooksInstanceReservationAffinityKey(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNotebooksInstanceReservationAffinityValues(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNotebooksInstanceInstallGpuDriver(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNotebooksInstanceCustomGpuDriverPath(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNotebooksInstanceBootDiskType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNotebooksInstanceBootDiskSizeGb(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNotebooksInstanceDataDiskType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNotebooksInstanceDataDiskSizeGb(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNotebooksInstanceNoRemoveDataDisk(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNotebooksInstanceDiskEncryption(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNotebooksInstanceKmsKey(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNotebooksInstanceNoPublicIp(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNotebooksInstanceNoProxyAccess(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNotebooksInstanceNetwork(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNotebooksInstanceSubnet(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNotebooksInstanceTags(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNotebooksInstanceMetadata(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandNotebooksInstanceVmImage(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedProject, err := expandNotebooksInstanceVmImageProject(original["project"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedProject); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["project"] = transformedProject
	}

	transformedImageFamily, err := expandNotebooksInstanceVmImageImageFamily(original["image_family"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedImageFamily); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["imageFamily"] = transformedImageFamily
	}

	transformedImageName, err := expandNotebooksInstanceVmImageImageName(original["image_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedImageName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["imageName"] = transformedImageName
	}

	return transformed, nil
}

func expandNotebooksInstanceVmImageProject(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNotebooksInstanceVmImageImageFamily(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNotebooksInstanceVmImageImageName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNotebooksInstanceContainerImage(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedRepository, err := expandNotebooksInstanceContainerImageRepository(original["repository"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRepository); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["repository"] = transformedRepository
	}

	transformedTag, err := expandNotebooksInstanceContainerImageTag(original["tag"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTag); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["tag"] = transformedTag
	}

	return transformed, nil
}

func expandNotebooksInstanceContainerImageRepository(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNotebooksInstanceContainerImageTag(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNotebooksInstanceEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
