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

package datapipeline

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/cai2hcl/common"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/caiasset"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const DataPipelinePipelineAssetType string = "datapipelines.googleapis.com/Pipeline"

const DataPipelinePipelineAssetNameRegex string = "projects/(?P<project>[^/]+)/locations/(?P<region>[^/]+)/pipelines"

type DataPipelinePipelineConverter struct {
	name   string
	schema map[string]*schema.Schema
}

func NewDataPipelinePipelineConverter(name string, schema map[string]*schema.Schema) common.Converter {
	return &DataPipelinePipelineConverter{
		name:   name,
		schema: schema,
	}
}

func (c *DataPipelinePipelineConverter) Convert(assets []*caiasset.Asset) ([]*common.HCLResourceBlock, error) {
	var blocks []*common.HCLResourceBlock
	config := common.NewConfig()

	for _, asset := range assets {
		if asset == nil {
			continue
		}
		if asset.Resource != nil && asset.Resource.Data != nil {
			block, err := c.convertResourceData(asset, config)
			if err != nil {
				return nil, err
			}
			blocks = append(blocks, block)
		}
	}
	return blocks, nil
}

func (c *DataPipelinePipelineConverter) convertResourceData(asset *caiasset.Asset, config *transport_tpg.Config) (*common.HCLResourceBlock, error) {
	if asset == nil || asset.Resource == nil || asset.Resource.Data == nil {
		return nil, fmt.Errorf("asset resource data is nil")
	}

	assetResourceData := asset.Resource.Data

	hcl, _ := resourceDataPipelinePipelineRead(assetResourceData, config)

	ctyVal, err := common.MapToCtyValWithSchema(hcl, c.schema)
	if err != nil {
		return nil, err
	}

	resourceName := assetResourceData["name"].(string)

	return &common.HCLResourceBlock{
		Labels: []string{c.name, resourceName},
		Value:  ctyVal,
	}, nil
}

func resourceDataPipelinePipelineRead(resource map[string]interface{}, config *transport_tpg.Config) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	var resource_data *schema.ResourceData = nil

	result["name"] = flattenDataPipelinePipelineName(resource["name"], resource_data, config)
	result["display_name"] = flattenDataPipelinePipelineDisplayName(resource["displayName"], resource_data, config)
	result["type"] = flattenDataPipelinePipelineType(resource["type"], resource_data, config)
	result["state"] = flattenDataPipelinePipelineState(resource["state"], resource_data, config)
	result["create_time"] = flattenDataPipelinePipelineCreateTime(resource["createTime"], resource_data, config)
	result["last_update_time"] = flattenDataPipelinePipelineLastUpdateTime(resource["lastUpdateTime"], resource_data, config)
	result["workload"] = flattenDataPipelinePipelineWorkload(resource["workload"], resource_data, config)
	result["schedule_info"] = flattenDataPipelinePipelineScheduleInfo(resource["scheduleInfo"], resource_data, config)
	result["job_count"] = flattenDataPipelinePipelineJobCount(resource["jobCount"], resource_data, config)
	result["scheduler_service_account_email"] = flattenDataPipelinePipelineSchedulerServiceAccountEmail(resource["schedulerServiceAccountEmail"], resource_data, config)
	result["pipeline_sources"] = flattenDataPipelinePipelinePipelineSources(resource["pipelineSources"], resource_data, config)

	return result, nil
}

func flattenDataPipelinePipelineName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.NameFromSelfLinkStateFunc(v)
}

func flattenDataPipelinePipelineDisplayName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataPipelinePipelineType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataPipelinePipelineState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataPipelinePipelineCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataPipelinePipelineLastUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataPipelinePipelineWorkload(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["dataflow_launch_template_request"] =
		flattenDataPipelinePipelineWorkloadDataflowLaunchTemplateRequest(original["dataflowLaunchTemplateRequest"], d, config)
	transformed["dataflow_flex_template_request"] =
		flattenDataPipelinePipelineWorkloadDataflowFlexTemplateRequest(original["dataflowFlexTemplateRequest"], d, config)
	return []interface{}{transformed}
}
func flattenDataPipelinePipelineWorkloadDataflowLaunchTemplateRequest(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["project_id"] =
		flattenDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestProjectId(original["projectId"], d, config)
	transformed["validate_only"] =
		flattenDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestValidateOnly(original["validateOnly"], d, config)
	transformed["launch_parameters"] =
		flattenDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestLaunchParameters(original["launchParameters"], d, config)
	transformed["location"] =
		flattenDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestLocation(original["location"], d, config)
	transformed["gcs_path"] =
		flattenDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestGcsPath(original["gcsPath"], d, config)
	return []interface{}{transformed}
}
func flattenDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestProjectId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestValidateOnly(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestLaunchParameters(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["job_name"] =
		flattenDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestLaunchParametersJobName(original["jobName"], d, config)
	transformed["parameters"] =
		flattenDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestLaunchParametersParameters(original["parameters"], d, config)
	transformed["environment"] =
		flattenDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestLaunchParametersEnvironment(original["environment"], d, config)
	transformed["update"] =
		flattenDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestLaunchParametersUpdate(original["update"], d, config)
	transformed["transform_name_mapping"] =
		flattenDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestLaunchParametersTransformNameMapping(original["transformNameMapping"], d, config)
	return []interface{}{transformed}
}
func flattenDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestLaunchParametersJobName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestLaunchParametersParameters(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestLaunchParametersEnvironment(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["num_workers"] =
		flattenDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestLaunchParametersEnvironmentNumWorkers(original["numWorkers"], d, config)
	transformed["max_workers"] =
		flattenDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestLaunchParametersEnvironmentMaxWorkers(original["maxWorkers"], d, config)
	transformed["zone"] =
		flattenDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestLaunchParametersEnvironmentZone(original["zone"], d, config)
	transformed["service_account_email"] =
		flattenDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestLaunchParametersEnvironmentServiceAccountEmail(original["serviceAccountEmail"], d, config)
	transformed["temp_location"] =
		flattenDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestLaunchParametersEnvironmentTempLocation(original["tempLocation"], d, config)
	transformed["bypass_temp_dir_validation"] =
		flattenDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestLaunchParametersEnvironmentBypassTempDirValidation(original["bypassTempDirValidation"], d, config)
	transformed["machine_type"] =
		flattenDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestLaunchParametersEnvironmentMachineType(original["machineType"], d, config)
	transformed["additional_experiments"] =
		flattenDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestLaunchParametersEnvironmentAdditionalExperiments(original["additionalExperiments"], d, config)
	transformed["network"] =
		flattenDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestLaunchParametersEnvironmentNetwork(original["network"], d, config)
	transformed["subnetwork"] =
		flattenDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestLaunchParametersEnvironmentSubnetwork(original["subnetwork"], d, config)
	transformed["additional_user_labels"] =
		flattenDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestLaunchParametersEnvironmentAdditionalUserLabels(original["additionalUserLabels"], d, config)
	transformed["kms_key_name"] =
		flattenDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestLaunchParametersEnvironmentKmsKeyName(original["kmsKeyName"], d, config)
	transformed["ip_configuration"] =
		flattenDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestLaunchParametersEnvironmentIpConfiguration(original["ipConfiguration"], d, config)
	transformed["worker_region"] =
		flattenDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestLaunchParametersEnvironmentWorkerRegion(original["workerRegion"], d, config)
	transformed["worker_zone"] =
		flattenDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestLaunchParametersEnvironmentWorkerZone(original["workerZone"], d, config)
	transformed["enable_streaming_engine"] =
		flattenDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestLaunchParametersEnvironmentEnableStreamingEngine(original["enableStreamingEngine"], d, config)
	return []interface{}{transformed}
}
func flattenDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestLaunchParametersEnvironmentNumWorkers(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := tpgresource.StringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestLaunchParametersEnvironmentMaxWorkers(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := tpgresource.StringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestLaunchParametersEnvironmentZone(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestLaunchParametersEnvironmentServiceAccountEmail(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestLaunchParametersEnvironmentTempLocation(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestLaunchParametersEnvironmentBypassTempDirValidation(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestLaunchParametersEnvironmentMachineType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestLaunchParametersEnvironmentAdditionalExperiments(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestLaunchParametersEnvironmentNetwork(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestLaunchParametersEnvironmentSubnetwork(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestLaunchParametersEnvironmentAdditionalUserLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestLaunchParametersEnvironmentKmsKeyName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestLaunchParametersEnvironmentIpConfiguration(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestLaunchParametersEnvironmentWorkerRegion(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestLaunchParametersEnvironmentWorkerZone(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestLaunchParametersEnvironmentEnableStreamingEngine(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestLaunchParametersUpdate(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestLaunchParametersTransformNameMapping(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestLocation(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestGcsPath(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataPipelinePipelineWorkloadDataflowFlexTemplateRequest(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["project_id"] =
		flattenDataPipelinePipelineWorkloadDataflowFlexTemplateRequestProjectId(original["projectId"], d, config)
	transformed["launch_parameter"] =
		flattenDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameter(original["launchParameter"], d, config)
	transformed["location"] =
		flattenDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLocation(original["location"], d, config)
	transformed["validate_only"] =
		flattenDataPipelinePipelineWorkloadDataflowFlexTemplateRequestValidateOnly(original["validateOnly"], d, config)
	return []interface{}{transformed}
}
func flattenDataPipelinePipelineWorkloadDataflowFlexTemplateRequestProjectId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameter(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["job_name"] =
		flattenDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterJobName(original["jobName"], d, config)
	transformed["parameters"] =
		flattenDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterParameters(original["parameters"], d, config)
	transformed["launch_options"] =
		flattenDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterLaunchOptions(original["launchOptions"], d, config)
	transformed["environment"] =
		flattenDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterEnvironment(original["environment"], d, config)
	transformed["update"] =
		flattenDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterUpdate(original["update"], d, config)
	transformed["transform_name_mappings"] =
		flattenDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterTransformNameMappings(original["transformNameMappings"], d, config)
	transformed["container_spec_gcs_path"] =
		flattenDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterContainerSpecGcsPath(original["containerSpecGcsPath"], d, config)
	return []interface{}{transformed}
}
func flattenDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterJobName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterParameters(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterLaunchOptions(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterEnvironment(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["num_workers"] =
		flattenDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterEnvironmentNumWorkers(original["numWorkers"], d, config)
	transformed["max_workers"] =
		flattenDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterEnvironmentMaxWorkers(original["maxWorkers"], d, config)
	transformed["zone"] =
		flattenDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterEnvironmentZone(original["zone"], d, config)
	transformed["service_account_email"] =
		flattenDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterEnvironmentServiceAccountEmail(original["serviceAccountEmail"], d, config)
	transformed["temp_location"] =
		flattenDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterEnvironmentTempLocation(original["tempLocation"], d, config)
	transformed["machine_type"] =
		flattenDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterEnvironmentMachineType(original["machineType"], d, config)
	transformed["additional_experiments"] =
		flattenDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterEnvironmentAdditionalExperiments(original["additionalExperiments"], d, config)
	transformed["network"] =
		flattenDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterEnvironmentNetwork(original["network"], d, config)
	transformed["subnetwork"] =
		flattenDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterEnvironmentSubnetwork(original["subnetwork"], d, config)
	transformed["additional_user_labels"] =
		flattenDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterEnvironmentAdditionalUserLabels(original["additionalUserLabels"], d, config)
	transformed["kms_key_name"] =
		flattenDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterEnvironmentKmsKeyName(original["kmsKeyName"], d, config)
	transformed["ip_configuration"] =
		flattenDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterEnvironmentIpConfiguration(original["ipConfiguration"], d, config)
	transformed["worker_region"] =
		flattenDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterEnvironmentWorkerRegion(original["workerRegion"], d, config)
	transformed["worker_zone"] =
		flattenDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterEnvironmentWorkerZone(original["workerZone"], d, config)
	transformed["enable_streaming_engine"] =
		flattenDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterEnvironmentEnableStreamingEngine(original["enableStreamingEngine"], d, config)
	transformed["flexrs_goal"] =
		flattenDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterEnvironmentFlexrsGoal(original["flexrsGoal"], d, config)
	return []interface{}{transformed}
}
func flattenDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterEnvironmentNumWorkers(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := tpgresource.StringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterEnvironmentMaxWorkers(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := tpgresource.StringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterEnvironmentZone(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterEnvironmentServiceAccountEmail(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterEnvironmentTempLocation(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterEnvironmentMachineType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterEnvironmentAdditionalExperiments(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterEnvironmentNetwork(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterEnvironmentSubnetwork(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterEnvironmentAdditionalUserLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterEnvironmentKmsKeyName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterEnvironmentIpConfiguration(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterEnvironmentWorkerRegion(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterEnvironmentWorkerZone(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterEnvironmentEnableStreamingEngine(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterEnvironmentFlexrsGoal(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterUpdate(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterTransformNameMappings(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterContainerSpecGcsPath(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLocation(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataPipelinePipelineWorkloadDataflowFlexTemplateRequestValidateOnly(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataPipelinePipelineScheduleInfo(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["schedule"] =
		flattenDataPipelinePipelineScheduleInfoSchedule(original["schedule"], d, config)
	transformed["time_zone"] =
		flattenDataPipelinePipelineScheduleInfoTimeZone(original["timeZone"], d, config)
	transformed["next_job_time"] =
		flattenDataPipelinePipelineScheduleInfoNextJobTime(original["nextJobTime"], d, config)
	return []interface{}{transformed}
}
func flattenDataPipelinePipelineScheduleInfoSchedule(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataPipelinePipelineScheduleInfoTimeZone(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataPipelinePipelineScheduleInfoNextJobTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataPipelinePipelineJobCount(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := tpgresource.StringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenDataPipelinePipelineSchedulerServiceAccountEmail(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataPipelinePipelinePipelineSources(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}
