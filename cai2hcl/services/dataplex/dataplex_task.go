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

package dataplex

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/cai2hcl/common"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/caiasset"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const DataplexTaskAssetType string = "dataplex.googleapis.com/Task"

const DataplexTaskAssetNameRegex string = "projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/lakes/(?P<lake>[^/]+)/tasks/(?P<task_id>[^/]+)"

type DataplexTaskConverter struct {
	name   string
	schema map[string]*schema.Schema
}

func NewDataplexTaskConverter(name string, schema map[string]*schema.Schema) common.Converter {
	return &DataplexTaskConverter{
		name:   name,
		schema: schema,
	}
}

func (c *DataplexTaskConverter) Convert(assets []*caiasset.Asset) ([]*common.HCLResourceBlock, error) {
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

func (c *DataplexTaskConverter) convertResourceData(asset *caiasset.Asset, config *transport_tpg.Config) (*common.HCLResourceBlock, error) {
	if asset == nil || asset.Resource == nil || asset.Resource.Data == nil {
		return nil, fmt.Errorf("asset resource data is nil")
	}

	assetResourceData := asset.Resource.Data

	hcl, _ := resourceDataplexTaskRead(assetResourceData, config)

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

func resourceDataplexTaskRead(resource map[string]interface{}, config *transport_tpg.Config) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	var resource_data *schema.ResourceData = nil

	result["name"] = flattenDataplexTaskName(resource["name"], resource_data, config)
	result["uid"] = flattenDataplexTaskUid(resource["uid"], resource_data, config)
	result["create_time"] = flattenDataplexTaskCreateTime(resource["createTime"], resource_data, config)
	result["update_time"] = flattenDataplexTaskUpdateTime(resource["updateTime"], resource_data, config)
	result["description"] = flattenDataplexTaskDescription(resource["description"], resource_data, config)
	result["display_name"] = flattenDataplexTaskDisplayName(resource["displayName"], resource_data, config)
	result["state"] = flattenDataplexTaskState(resource["state"], resource_data, config)
	result["labels"] = flattenDataplexTaskLabels(resource["labels"], resource_data, config)
	result["trigger_spec"] = flattenDataplexTaskTriggerSpec(resource["triggerSpec"], resource_data, config)
	result["execution_spec"] = flattenDataplexTaskExecutionSpec(resource["executionSpec"], resource_data, config)
	result["execution_status"] = flattenDataplexTaskExecutionStatus(resource["executionStatus"], resource_data, config)
	result["spark"] = flattenDataplexTaskSpark(resource["spark"], resource_data, config)
	result["notebook"] = flattenDataplexTaskNotebook(resource["notebook"], resource_data, config)
	result["terraform_labels"] = flattenDataplexTaskTerraformLabels(resource["labels"], resource_data, config)
	result["effective_labels"] = flattenDataplexTaskEffectiveLabels(resource["labels"], resource_data, config)

	return result, nil
}

func flattenDataplexTaskName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataplexTaskUid(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataplexTaskCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataplexTaskUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataplexTaskDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataplexTaskDisplayName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataplexTaskState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataplexTaskLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}

	transformed := make(map[string]interface{})
	if l, ok := d.GetOkExists("labels"); ok {
		for k := range l.(map[string]interface{}) {
			transformed[k] = v.(map[string]interface{})[k]
		}
	}

	return transformed
}

func flattenDataplexTaskTriggerSpec(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["type"] =
		flattenDataplexTaskTriggerSpecType(original["type"], d, config)
	transformed["start_time"] =
		flattenDataplexTaskTriggerSpecStartTime(original["startTime"], d, config)
	transformed["disabled"] =
		flattenDataplexTaskTriggerSpecDisabled(original["disabled"], d, config)
	transformed["max_retries"] =
		flattenDataplexTaskTriggerSpecMaxRetries(original["maxRetries"], d, config)
	transformed["schedule"] =
		flattenDataplexTaskTriggerSpecSchedule(original["schedule"], d, config)
	return []interface{}{transformed}
}
func flattenDataplexTaskTriggerSpecType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataplexTaskTriggerSpecStartTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataplexTaskTriggerSpecDisabled(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataplexTaskTriggerSpecMaxRetries(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenDataplexTaskTriggerSpecSchedule(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataplexTaskExecutionSpec(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["args"] =
		flattenDataplexTaskExecutionSpecArgs(original["args"], d, config)
	transformed["service_account"] =
		flattenDataplexTaskExecutionSpecServiceAccount(original["serviceAccount"], d, config)
	transformed["project"] =
		flattenDataplexTaskExecutionSpecProject(original["project"], d, config)
	transformed["max_job_execution_lifetime"] =
		flattenDataplexTaskExecutionSpecMaxJobExecutionLifetime(original["maxJobExecutionLifetime"], d, config)
	transformed["kms_key"] =
		flattenDataplexTaskExecutionSpecKmsKey(original["kmsKey"], d, config)
	return []interface{}{transformed}
}
func flattenDataplexTaskExecutionSpecArgs(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataplexTaskExecutionSpecServiceAccount(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataplexTaskExecutionSpecProject(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataplexTaskExecutionSpecMaxJobExecutionLifetime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataplexTaskExecutionSpecKmsKey(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataplexTaskExecutionStatus(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["update_time"] =
		flattenDataplexTaskExecutionStatusUpdateTime(original["updateTime"], d, config)
	transformed["latest_job"] =
		flattenDataplexTaskExecutionStatusLatestJob(original["latestJob"], d, config)
	return []interface{}{transformed}
}
func flattenDataplexTaskExecutionStatusUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataplexTaskExecutionStatusLatestJob(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["name"] =
		flattenDataplexTaskExecutionStatusLatestJobName(original["name"], d, config)
	transformed["uid"] =
		flattenDataplexTaskExecutionStatusLatestJobUid(original["uid"], d, config)
	transformed["start_time"] =
		flattenDataplexTaskExecutionStatusLatestJobStartTime(original["startTime"], d, config)
	transformed["end_time"] =
		flattenDataplexTaskExecutionStatusLatestJobEndTime(original["endTime"], d, config)
	transformed["state"] =
		flattenDataplexTaskExecutionStatusLatestJobState(original["state"], d, config)
	transformed["retry_count"] =
		flattenDataplexTaskExecutionStatusLatestJobRetryCount(original["retryCount"], d, config)
	transformed["service"] =
		flattenDataplexTaskExecutionStatusLatestJobService(original["service"], d, config)
	transformed["service_job"] =
		flattenDataplexTaskExecutionStatusLatestJobServiceJob(original["serviceJob"], d, config)
	transformed["message"] =
		flattenDataplexTaskExecutionStatusLatestJobMessage(original["message"], d, config)
	return []interface{}{transformed}
}
func flattenDataplexTaskExecutionStatusLatestJobName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataplexTaskExecutionStatusLatestJobUid(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataplexTaskExecutionStatusLatestJobStartTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataplexTaskExecutionStatusLatestJobEndTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataplexTaskExecutionStatusLatestJobState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataplexTaskExecutionStatusLatestJobRetryCount(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenDataplexTaskExecutionStatusLatestJobService(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataplexTaskExecutionStatusLatestJobServiceJob(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataplexTaskExecutionStatusLatestJobMessage(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataplexTaskSpark(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["file_uris"] =
		flattenDataplexTaskSparkFileUris(original["fileUris"], d, config)
	transformed["archive_uris"] =
		flattenDataplexTaskSparkArchiveUris(original["archiveUris"], d, config)
	transformed["infrastructure_spec"] =
		flattenDataplexTaskSparkInfrastructureSpec(original["infrastructureSpec"], d, config)
	transformed["main_jar_file_uri"] =
		flattenDataplexTaskSparkMainJarFileUri(original["mainJarFileUri"], d, config)
	transformed["main_class"] =
		flattenDataplexTaskSparkMainClass(original["mainClass"], d, config)
	transformed["python_script_file"] =
		flattenDataplexTaskSparkPythonScriptFile(original["pythonScriptFile"], d, config)
	transformed["sql_script_file"] =
		flattenDataplexTaskSparkSqlScriptFile(original["sqlScriptFile"], d, config)
	transformed["sql_script"] =
		flattenDataplexTaskSparkSqlScript(original["sqlScript"], d, config)
	return []interface{}{transformed}
}
func flattenDataplexTaskSparkFileUris(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataplexTaskSparkArchiveUris(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataplexTaskSparkInfrastructureSpec(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["batch"] =
		flattenDataplexTaskSparkInfrastructureSpecBatch(original["batch"], d, config)
	transformed["container_image"] =
		flattenDataplexTaskSparkInfrastructureSpecContainerImage(original["containerImage"], d, config)
	transformed["vpc_network"] =
		flattenDataplexTaskSparkInfrastructureSpecVpcNetwork(original["vpcNetwork"], d, config)
	return []interface{}{transformed}
}
func flattenDataplexTaskSparkInfrastructureSpecBatch(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["executors_count"] =
		flattenDataplexTaskSparkInfrastructureSpecBatchExecutorsCount(original["executorsCount"], d, config)
	transformed["max_executors_count"] =
		flattenDataplexTaskSparkInfrastructureSpecBatchMaxExecutorsCount(original["maxExecutorsCount"], d, config)
	return []interface{}{transformed}
}
func flattenDataplexTaskSparkInfrastructureSpecBatchExecutorsCount(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenDataplexTaskSparkInfrastructureSpecBatchMaxExecutorsCount(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenDataplexTaskSparkInfrastructureSpecContainerImage(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["image"] =
		flattenDataplexTaskSparkInfrastructureSpecContainerImageImage(original["image"], d, config)
	transformed["java_jars"] =
		flattenDataplexTaskSparkInfrastructureSpecContainerImageJavaJars(original["javaJars"], d, config)
	transformed["python_packages"] =
		flattenDataplexTaskSparkInfrastructureSpecContainerImagePythonPackages(original["pythonPackages"], d, config)
	transformed["properties"] =
		flattenDataplexTaskSparkInfrastructureSpecContainerImageProperties(original["properties"], d, config)
	return []interface{}{transformed}
}
func flattenDataplexTaskSparkInfrastructureSpecContainerImageImage(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataplexTaskSparkInfrastructureSpecContainerImageJavaJars(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataplexTaskSparkInfrastructureSpecContainerImagePythonPackages(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataplexTaskSparkInfrastructureSpecContainerImageProperties(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataplexTaskSparkInfrastructureSpecVpcNetwork(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["network_tags"] =
		flattenDataplexTaskSparkInfrastructureSpecVpcNetworkNetworkTags(original["networkTags"], d, config)
	transformed["network"] =
		flattenDataplexTaskSparkInfrastructureSpecVpcNetworkNetwork(original["network"], d, config)
	transformed["sub_network"] =
		flattenDataplexTaskSparkInfrastructureSpecVpcNetworkSubNetwork(original["subNetwork"], d, config)
	return []interface{}{transformed}
}
func flattenDataplexTaskSparkInfrastructureSpecVpcNetworkNetworkTags(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataplexTaskSparkInfrastructureSpecVpcNetworkNetwork(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataplexTaskSparkInfrastructureSpecVpcNetworkSubNetwork(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataplexTaskSparkMainJarFileUri(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataplexTaskSparkMainClass(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataplexTaskSparkPythonScriptFile(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataplexTaskSparkSqlScriptFile(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataplexTaskSparkSqlScript(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataplexTaskNotebook(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["notebook"] =
		flattenDataplexTaskNotebookNotebook(original["notebook"], d, config)
	transformed["infrastructure_spec"] =
		flattenDataplexTaskNotebookInfrastructureSpec(original["infrastructureSpec"], d, config)
	transformed["file_uris"] =
		flattenDataplexTaskNotebookFileUris(original["fileUris"], d, config)
	transformed["archive_uris"] =
		flattenDataplexTaskNotebookArchiveUris(original["archiveUris"], d, config)
	return []interface{}{transformed}
}
func flattenDataplexTaskNotebookNotebook(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataplexTaskNotebookInfrastructureSpec(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["batch"] =
		flattenDataplexTaskNotebookInfrastructureSpecBatch(original["batch"], d, config)
	transformed["container_image"] =
		flattenDataplexTaskNotebookInfrastructureSpecContainerImage(original["containerImage"], d, config)
	transformed["vpc_network"] =
		flattenDataplexTaskNotebookInfrastructureSpecVpcNetwork(original["vpcNetwork"], d, config)
	return []interface{}{transformed}
}
func flattenDataplexTaskNotebookInfrastructureSpecBatch(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["executors_count"] =
		flattenDataplexTaskNotebookInfrastructureSpecBatchExecutorsCount(original["executorsCount"], d, config)
	transformed["max_executors_count"] =
		flattenDataplexTaskNotebookInfrastructureSpecBatchMaxExecutorsCount(original["maxExecutorsCount"], d, config)
	return []interface{}{transformed}
}
func flattenDataplexTaskNotebookInfrastructureSpecBatchExecutorsCount(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenDataplexTaskNotebookInfrastructureSpecBatchMaxExecutorsCount(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenDataplexTaskNotebookInfrastructureSpecContainerImage(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["image"] =
		flattenDataplexTaskNotebookInfrastructureSpecContainerImageImage(original["image"], d, config)
	transformed["java_jars"] =
		flattenDataplexTaskNotebookInfrastructureSpecContainerImageJavaJars(original["javaJars"], d, config)
	transformed["python_packages"] =
		flattenDataplexTaskNotebookInfrastructureSpecContainerImagePythonPackages(original["pythonPackages"], d, config)
	transformed["properties"] =
		flattenDataplexTaskNotebookInfrastructureSpecContainerImageProperties(original["properties"], d, config)
	return []interface{}{transformed}
}
func flattenDataplexTaskNotebookInfrastructureSpecContainerImageImage(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataplexTaskNotebookInfrastructureSpecContainerImageJavaJars(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataplexTaskNotebookInfrastructureSpecContainerImagePythonPackages(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataplexTaskNotebookInfrastructureSpecContainerImageProperties(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataplexTaskNotebookInfrastructureSpecVpcNetwork(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["network_tags"] =
		flattenDataplexTaskNotebookInfrastructureSpecVpcNetworkNetworkTags(original["networkTags"], d, config)
	transformed["network"] =
		flattenDataplexTaskNotebookInfrastructureSpecVpcNetworkNetwork(original["network"], d, config)
	transformed["sub_network"] =
		flattenDataplexTaskNotebookInfrastructureSpecVpcNetworkSubNetwork(original["subNetwork"], d, config)
	return []interface{}{transformed}
}
func flattenDataplexTaskNotebookInfrastructureSpecVpcNetworkNetworkTags(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataplexTaskNotebookInfrastructureSpecVpcNetworkNetwork(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataplexTaskNotebookInfrastructureSpecVpcNetworkSubNetwork(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataplexTaskNotebookFileUris(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataplexTaskNotebookArchiveUris(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataplexTaskTerraformLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}

	transformed := make(map[string]interface{})
	if l, ok := d.GetOkExists("terraform_labels"); ok {
		for k := range l.(map[string]interface{}) {
			transformed[k] = v.(map[string]interface{})[k]
		}
	}

	return transformed
}

func flattenDataplexTaskEffectiveLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}
