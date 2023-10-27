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

package monitoring

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/cai2hcl/common"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/caiasset"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

// API does not return a value for REDUCE_NONE
func crossSeriesReducerDiffSuppress(k, old, new string, d *schema.ResourceData) bool {
	return (new == "" && old == "REDUCE_NONE") || (new == "REDUCE_NONE" && old == "")
}

const MonitoringAlertPolicyAssetType string = "monitoring.googleapis.com/AlertPolicy"

const MonitoringAlertPolicyAssetNameRegex string = "v3/projects/(?P<project>[^/]+)/alertPolicies"

type MonitoringAlertPolicyConverter struct {
	name   string
	schema map[string]*schema.Schema
}

func NewMonitoringAlertPolicyConverter(name string, schema map[string]*schema.Schema) common.Converter {
	return &MonitoringAlertPolicyConverter{
		name:   name,
		schema: schema,
	}
}

func (c *MonitoringAlertPolicyConverter) Convert(assets []*caiasset.Asset) ([]*common.HCLResourceBlock, error) {
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

func (c *MonitoringAlertPolicyConverter) convertResourceData(asset *caiasset.Asset, config *transport_tpg.Config) (*common.HCLResourceBlock, error) {
	if asset == nil || asset.Resource == nil || asset.Resource.Data == nil {
		return nil, fmt.Errorf("asset resource data is nil")
	}

	assetResourceData := asset.Resource.Data

	hcl, _ := resourceMonitoringAlertPolicyRead(assetResourceData, config)

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

func resourceMonitoringAlertPolicyRead(resource map[string]interface{}, config *transport_tpg.Config) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	var resource_data *schema.ResourceData = nil

	result["name"] = flattenMonitoringAlertPolicyName(resource["name"], resource_data, config)
	result["display_name"] = flattenMonitoringAlertPolicyDisplayName(resource["displayName"], resource_data, config)
	result["combiner"] = flattenMonitoringAlertPolicyCombiner(resource["combiner"], resource_data, config)
	result["creation_record"] = flattenMonitoringAlertPolicyCreationRecord(resource["creationRecord"], resource_data, config)
	result["enabled"] = flattenMonitoringAlertPolicyEnabled(resource["enabled"], resource_data, config)
	result["conditions"] = flattenMonitoringAlertPolicyConditions(resource["conditions"], resource_data, config)
	result["notification_channels"] = flattenMonitoringAlertPolicyNotificationChannels(resource["notificationChannels"], resource_data, config)
	result["alert_strategy"] = flattenMonitoringAlertPolicyAlertStrategy(resource["alertStrategy"], resource_data, config)
	result["user_labels"] = flattenMonitoringAlertPolicyUserLabels(resource["userLabels"], resource_data, config)
	result["documentation"] = flattenMonitoringAlertPolicyDocumentation(resource["documentation"], resource_data, config)

	return result, nil
}

func flattenMonitoringAlertPolicyName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenMonitoringAlertPolicyDisplayName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenMonitoringAlertPolicyCombiner(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenMonitoringAlertPolicyCreationRecord(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["mutate_time"] =
		flattenMonitoringAlertPolicyCreationRecordMutateTime(original["mutateTime"], d, config)
	transformed["mutated_by"] =
		flattenMonitoringAlertPolicyCreationRecordMutatedBy(original["mutatedBy"], d, config)
	return []interface{}{transformed}
}
func flattenMonitoringAlertPolicyCreationRecordMutateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenMonitoringAlertPolicyCreationRecordMutatedBy(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenMonitoringAlertPolicyEnabled(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenMonitoringAlertPolicyConditions(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed = append(transformed, map[string]interface{}{
			"condition_absent":                    flattenMonitoringAlertPolicyConditionsConditionAbsent(original["conditionAbsent"], d, config),
			"name":                                flattenMonitoringAlertPolicyConditionsName(original["name"], d, config),
			"condition_monitoring_query_language": flattenMonitoringAlertPolicyConditionsConditionMonitoringQueryLanguage(original["conditionMonitoringQueryLanguage"], d, config),
			"condition_threshold":                 flattenMonitoringAlertPolicyConditionsConditionThreshold(original["conditionThreshold"], d, config),
			"display_name":                        flattenMonitoringAlertPolicyConditionsDisplayName(original["displayName"], d, config),
			"condition_matched_log":               flattenMonitoringAlertPolicyConditionsConditionMatchedLog(original["conditionMatchedLog"], d, config),
			"condition_prometheus_query_language": flattenMonitoringAlertPolicyConditionsConditionPrometheusQueryLanguage(original["conditionPrometheusQueryLanguage"], d, config),
		})
	}
	return transformed
}
func flattenMonitoringAlertPolicyConditionsConditionAbsent(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["aggregations"] =
		flattenMonitoringAlertPolicyConditionsConditionAbsentAggregations(original["aggregations"], d, config)
	transformed["trigger"] =
		flattenMonitoringAlertPolicyConditionsConditionAbsentTrigger(original["trigger"], d, config)
	transformed["duration"] =
		flattenMonitoringAlertPolicyConditionsConditionAbsentDuration(original["duration"], d, config)
	transformed["filter"] =
		flattenMonitoringAlertPolicyConditionsConditionAbsentFilter(original["filter"], d, config)
	return []interface{}{transformed}
}
func flattenMonitoringAlertPolicyConditionsConditionAbsentAggregations(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed = append(transformed, map[string]interface{}{
			"per_series_aligner":   flattenMonitoringAlertPolicyConditionsConditionAbsentAggregationsPerSeriesAligner(original["perSeriesAligner"], d, config),
			"group_by_fields":      flattenMonitoringAlertPolicyConditionsConditionAbsentAggregationsGroupByFields(original["groupByFields"], d, config),
			"alignment_period":     flattenMonitoringAlertPolicyConditionsConditionAbsentAggregationsAlignmentPeriod(original["alignmentPeriod"], d, config),
			"cross_series_reducer": flattenMonitoringAlertPolicyConditionsConditionAbsentAggregationsCrossSeriesReducer(original["crossSeriesReducer"], d, config),
		})
	}
	return transformed
}
func flattenMonitoringAlertPolicyConditionsConditionAbsentAggregationsPerSeriesAligner(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenMonitoringAlertPolicyConditionsConditionAbsentAggregationsGroupByFields(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenMonitoringAlertPolicyConditionsConditionAbsentAggregationsAlignmentPeriod(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenMonitoringAlertPolicyConditionsConditionAbsentAggregationsCrossSeriesReducer(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenMonitoringAlertPolicyConditionsConditionAbsentTrigger(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["percent"] =
		flattenMonitoringAlertPolicyConditionsConditionAbsentTriggerPercent(original["percent"], d, config)
	transformed["count"] =
		flattenMonitoringAlertPolicyConditionsConditionAbsentTriggerCount(original["count"], d, config)
	return []interface{}{transformed}
}
func flattenMonitoringAlertPolicyConditionsConditionAbsentTriggerPercent(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenMonitoringAlertPolicyConditionsConditionAbsentTriggerCount(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenMonitoringAlertPolicyConditionsConditionAbsentDuration(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenMonitoringAlertPolicyConditionsConditionAbsentFilter(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenMonitoringAlertPolicyConditionsName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenMonitoringAlertPolicyConditionsConditionMonitoringQueryLanguage(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["query"] =
		flattenMonitoringAlertPolicyConditionsConditionMonitoringQueryLanguageQuery(original["query"], d, config)
	transformed["duration"] =
		flattenMonitoringAlertPolicyConditionsConditionMonitoringQueryLanguageDuration(original["duration"], d, config)
	transformed["trigger"] =
		flattenMonitoringAlertPolicyConditionsConditionMonitoringQueryLanguageTrigger(original["trigger"], d, config)
	transformed["evaluation_missing_data"] =
		flattenMonitoringAlertPolicyConditionsConditionMonitoringQueryLanguageEvaluationMissingData(original["evaluationMissingData"], d, config)
	return []interface{}{transformed}
}
func flattenMonitoringAlertPolicyConditionsConditionMonitoringQueryLanguageQuery(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenMonitoringAlertPolicyConditionsConditionMonitoringQueryLanguageDuration(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenMonitoringAlertPolicyConditionsConditionMonitoringQueryLanguageTrigger(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["percent"] =
		flattenMonitoringAlertPolicyConditionsConditionMonitoringQueryLanguageTriggerPercent(original["percent"], d, config)
	transformed["count"] =
		flattenMonitoringAlertPolicyConditionsConditionMonitoringQueryLanguageTriggerCount(original["count"], d, config)
	return []interface{}{transformed}
}
func flattenMonitoringAlertPolicyConditionsConditionMonitoringQueryLanguageTriggerPercent(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenMonitoringAlertPolicyConditionsConditionMonitoringQueryLanguageTriggerCount(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenMonitoringAlertPolicyConditionsConditionMonitoringQueryLanguageEvaluationMissingData(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenMonitoringAlertPolicyConditionsConditionThreshold(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["threshold_value"] =
		flattenMonitoringAlertPolicyConditionsConditionThresholdThresholdValue(original["thresholdValue"], d, config)
	transformed["denominator_filter"] =
		flattenMonitoringAlertPolicyConditionsConditionThresholdDenominatorFilter(original["denominatorFilter"], d, config)
	transformed["denominator_aggregations"] =
		flattenMonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregations(original["denominatorAggregations"], d, config)
	transformed["duration"] =
		flattenMonitoringAlertPolicyConditionsConditionThresholdDuration(original["duration"], d, config)
	transformed["forecast_options"] =
		flattenMonitoringAlertPolicyConditionsConditionThresholdForecastOptions(original["forecastOptions"], d, config)
	transformed["comparison"] =
		flattenMonitoringAlertPolicyConditionsConditionThresholdComparison(original["comparison"], d, config)
	transformed["trigger"] =
		flattenMonitoringAlertPolicyConditionsConditionThresholdTrigger(original["trigger"], d, config)
	transformed["aggregations"] =
		flattenMonitoringAlertPolicyConditionsConditionThresholdAggregations(original["aggregations"], d, config)
	transformed["filter"] =
		flattenMonitoringAlertPolicyConditionsConditionThresholdFilter(original["filter"], d, config)
	transformed["evaluation_missing_data"] =
		flattenMonitoringAlertPolicyConditionsConditionThresholdEvaluationMissingData(original["evaluationMissingData"], d, config)
	return []interface{}{transformed}
}
func flattenMonitoringAlertPolicyConditionsConditionThresholdThresholdValue(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenMonitoringAlertPolicyConditionsConditionThresholdDenominatorFilter(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenMonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregations(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed = append(transformed, map[string]interface{}{
			"per_series_aligner":   flattenMonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsPerSeriesAligner(original["perSeriesAligner"], d, config),
			"group_by_fields":      flattenMonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsGroupByFields(original["groupByFields"], d, config),
			"alignment_period":     flattenMonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsAlignmentPeriod(original["alignmentPeriod"], d, config),
			"cross_series_reducer": flattenMonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsCrossSeriesReducer(original["crossSeriesReducer"], d, config),
		})
	}
	return transformed
}
func flattenMonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsPerSeriesAligner(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenMonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsGroupByFields(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenMonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsAlignmentPeriod(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenMonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsCrossSeriesReducer(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenMonitoringAlertPolicyConditionsConditionThresholdDuration(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenMonitoringAlertPolicyConditionsConditionThresholdForecastOptions(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["forecast_horizon"] =
		flattenMonitoringAlertPolicyConditionsConditionThresholdForecastOptionsForecastHorizon(original["forecastHorizon"], d, config)
	return []interface{}{transformed}
}
func flattenMonitoringAlertPolicyConditionsConditionThresholdForecastOptionsForecastHorizon(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenMonitoringAlertPolicyConditionsConditionThresholdComparison(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenMonitoringAlertPolicyConditionsConditionThresholdTrigger(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["percent"] =
		flattenMonitoringAlertPolicyConditionsConditionThresholdTriggerPercent(original["percent"], d, config)
	transformed["count"] =
		flattenMonitoringAlertPolicyConditionsConditionThresholdTriggerCount(original["count"], d, config)
	return []interface{}{transformed}
}
func flattenMonitoringAlertPolicyConditionsConditionThresholdTriggerPercent(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenMonitoringAlertPolicyConditionsConditionThresholdTriggerCount(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenMonitoringAlertPolicyConditionsConditionThresholdAggregations(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed = append(transformed, map[string]interface{}{
			"per_series_aligner":   flattenMonitoringAlertPolicyConditionsConditionThresholdAggregationsPerSeriesAligner(original["perSeriesAligner"], d, config),
			"group_by_fields":      flattenMonitoringAlertPolicyConditionsConditionThresholdAggregationsGroupByFields(original["groupByFields"], d, config),
			"alignment_period":     flattenMonitoringAlertPolicyConditionsConditionThresholdAggregationsAlignmentPeriod(original["alignmentPeriod"], d, config),
			"cross_series_reducer": flattenMonitoringAlertPolicyConditionsConditionThresholdAggregationsCrossSeriesReducer(original["crossSeriesReducer"], d, config),
		})
	}
	return transformed
}
func flattenMonitoringAlertPolicyConditionsConditionThresholdAggregationsPerSeriesAligner(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenMonitoringAlertPolicyConditionsConditionThresholdAggregationsGroupByFields(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenMonitoringAlertPolicyConditionsConditionThresholdAggregationsAlignmentPeriod(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenMonitoringAlertPolicyConditionsConditionThresholdAggregationsCrossSeriesReducer(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenMonitoringAlertPolicyConditionsConditionThresholdFilter(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenMonitoringAlertPolicyConditionsConditionThresholdEvaluationMissingData(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenMonitoringAlertPolicyConditionsDisplayName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenMonitoringAlertPolicyConditionsConditionMatchedLog(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["filter"] =
		flattenMonitoringAlertPolicyConditionsConditionMatchedLogFilter(original["filter"], d, config)
	transformed["label_extractors"] =
		flattenMonitoringAlertPolicyConditionsConditionMatchedLogLabelExtractors(original["labelExtractors"], d, config)
	return []interface{}{transformed}
}
func flattenMonitoringAlertPolicyConditionsConditionMatchedLogFilter(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenMonitoringAlertPolicyConditionsConditionMatchedLogLabelExtractors(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenMonitoringAlertPolicyConditionsConditionPrometheusQueryLanguage(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["query"] =
		flattenMonitoringAlertPolicyConditionsConditionPrometheusQueryLanguageQuery(original["query"], d, config)
	transformed["duration"] =
		flattenMonitoringAlertPolicyConditionsConditionPrometheusQueryLanguageDuration(original["duration"], d, config)
	transformed["evaluation_interval"] =
		flattenMonitoringAlertPolicyConditionsConditionPrometheusQueryLanguageEvaluationInterval(original["evaluationInterval"], d, config)
	transformed["labels"] =
		flattenMonitoringAlertPolicyConditionsConditionPrometheusQueryLanguageLabels(original["labels"], d, config)
	transformed["rule_group"] =
		flattenMonitoringAlertPolicyConditionsConditionPrometheusQueryLanguageRuleGroup(original["ruleGroup"], d, config)
	transformed["alert_rule"] =
		flattenMonitoringAlertPolicyConditionsConditionPrometheusQueryLanguageAlertRule(original["alertRule"], d, config)
	return []interface{}{transformed}
}
func flattenMonitoringAlertPolicyConditionsConditionPrometheusQueryLanguageQuery(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenMonitoringAlertPolicyConditionsConditionPrometheusQueryLanguageDuration(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenMonitoringAlertPolicyConditionsConditionPrometheusQueryLanguageEvaluationInterval(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenMonitoringAlertPolicyConditionsConditionPrometheusQueryLanguageLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenMonitoringAlertPolicyConditionsConditionPrometheusQueryLanguageRuleGroup(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenMonitoringAlertPolicyConditionsConditionPrometheusQueryLanguageAlertRule(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenMonitoringAlertPolicyNotificationChannels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenMonitoringAlertPolicyAlertStrategy(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["notification_rate_limit"] =
		flattenMonitoringAlertPolicyAlertStrategyNotificationRateLimit(original["notificationRateLimit"], d, config)
	transformed["auto_close"] =
		flattenMonitoringAlertPolicyAlertStrategyAutoClose(original["autoClose"], d, config)
	transformed["notification_channel_strategy"] =
		flattenMonitoringAlertPolicyAlertStrategyNotificationChannelStrategy(original["notificationChannelStrategy"], d, config)
	return []interface{}{transformed}
}
func flattenMonitoringAlertPolicyAlertStrategyNotificationRateLimit(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["period"] =
		flattenMonitoringAlertPolicyAlertStrategyNotificationRateLimitPeriod(original["period"], d, config)
	return []interface{}{transformed}
}
func flattenMonitoringAlertPolicyAlertStrategyNotificationRateLimitPeriod(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenMonitoringAlertPolicyAlertStrategyAutoClose(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenMonitoringAlertPolicyAlertStrategyNotificationChannelStrategy(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed = append(transformed, map[string]interface{}{
			"notification_channel_names": flattenMonitoringAlertPolicyAlertStrategyNotificationChannelStrategyNotificationChannelNames(original["notificationChannelNames"], d, config),
			"renotify_interval":          flattenMonitoringAlertPolicyAlertStrategyNotificationChannelStrategyRenotifyInterval(original["renotifyInterval"], d, config),
		})
	}
	return transformed
}
func flattenMonitoringAlertPolicyAlertStrategyNotificationChannelStrategyNotificationChannelNames(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenMonitoringAlertPolicyAlertStrategyNotificationChannelStrategyRenotifyInterval(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenMonitoringAlertPolicyUserLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenMonitoringAlertPolicyDocumentation(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["content"] =
		flattenMonitoringAlertPolicyDocumentationContent(original["content"], d, config)
	transformed["mime_type"] =
		flattenMonitoringAlertPolicyDocumentationMimeType(original["mimeType"], d, config)
	return []interface{}{transformed}
}
func flattenMonitoringAlertPolicyDocumentationContent(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenMonitoringAlertPolicyDocumentationMimeType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}
