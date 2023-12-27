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

package clouddeploy

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const ClouddeployAutomationAssetType string = "clouddeploy.googleapis.com/Automation"

func ResourceConverterClouddeployAutomation() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: ClouddeployAutomationAssetType,
		Convert:   GetClouddeployAutomationCaiObject,
	}
}

func GetClouddeployAutomationCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//clouddeploy.googleapis.com/projects/{{project}}/locations/{{location}}/deliveryPipelines/{{delivery_pipeline}}/automations/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetClouddeployAutomationApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: ClouddeployAutomationAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/clouddeploy/v1/rest",
				DiscoveryName:        "Automation",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetClouddeployAutomationApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	descriptionProp, err := expandClouddeployAutomationDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	suspendedProp, err := expandClouddeployAutomationSuspended(d.Get("suspended"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("suspended"); ok || !reflect.DeepEqual(v, suspendedProp) {
		obj["suspended"] = suspendedProp
	}
	serviceAccountProp, err := expandClouddeployAutomationServiceAccount(d.Get("service_account"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("service_account"); !tpgresource.IsEmptyValue(reflect.ValueOf(serviceAccountProp)) && (ok || !reflect.DeepEqual(v, serviceAccountProp)) {
		obj["serviceAccount"] = serviceAccountProp
	}
	selectorProp, err := expandClouddeployAutomationSelector(d.Get("selector"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("selector"); !tpgresource.IsEmptyValue(reflect.ValueOf(selectorProp)) && (ok || !reflect.DeepEqual(v, selectorProp)) {
		obj["selector"] = selectorProp
	}
	rulesProp, err := expandClouddeployAutomationRules(d.Get("rules"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("rules"); !tpgresource.IsEmptyValue(reflect.ValueOf(rulesProp)) && (ok || !reflect.DeepEqual(v, rulesProp)) {
		obj["rules"] = rulesProp
	}
	annotationsProp, err := expandClouddeployAutomationEffectiveAnnotations(d.Get("effective_annotations"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("effective_annotations"); !tpgresource.IsEmptyValue(reflect.ValueOf(annotationsProp)) && (ok || !reflect.DeepEqual(v, annotationsProp)) {
		obj["annotations"] = annotationsProp
	}
	labelsProp, err := expandClouddeployAutomationEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	return obj, nil
}

func expandClouddeployAutomationDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandClouddeployAutomationSuspended(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandClouddeployAutomationServiceAccount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandClouddeployAutomationSelector(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedTargets, err := expandClouddeployAutomationSelectorTargets(original["targets"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTargets); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["targets"] = transformedTargets
	}

	return transformed, nil
}

func expandClouddeployAutomationSelectorTargets(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedId, err := expandClouddeployAutomationSelectorTargetsId(original["id"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["id"] = transformedId
		}

		transformedLabels, err := expandClouddeployAutomationSelectorTargetsLabels(original["labels"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedLabels); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["labels"] = transformedLabels
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandClouddeployAutomationSelectorTargetsId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandClouddeployAutomationSelectorTargetsLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandClouddeployAutomationRules(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedPromoteReleaseRule, err := expandClouddeployAutomationRulesPromoteReleaseRule(original["promote_release_rule"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedPromoteReleaseRule); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["promoteReleaseRule"] = transformedPromoteReleaseRule
		}

		transformedAdvanceRolloutRule, err := expandClouddeployAutomationRulesAdvanceRolloutRule(original["advance_rollout_rule"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedAdvanceRolloutRule); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["advanceRolloutRule"] = transformedAdvanceRolloutRule
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandClouddeployAutomationRulesPromoteReleaseRule(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedId, err := expandClouddeployAutomationRulesPromoteReleaseRuleId(original["id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["id"] = transformedId
	}

	transformedWait, err := expandClouddeployAutomationRulesPromoteReleaseRuleWait(original["wait"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedWait); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["wait"] = transformedWait
	}

	transformedDestinationTargetId, err := expandClouddeployAutomationRulesPromoteReleaseRuleDestinationTargetId(original["destination_target_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDestinationTargetId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["destinationTargetId"] = transformedDestinationTargetId
	}

	transformedDestinationPhase, err := expandClouddeployAutomationRulesPromoteReleaseRuleDestinationPhase(original["destination_phase"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDestinationPhase); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["destinationPhase"] = transformedDestinationPhase
	}

	return transformed, nil
}

func expandClouddeployAutomationRulesPromoteReleaseRuleId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandClouddeployAutomationRulesPromoteReleaseRuleWait(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandClouddeployAutomationRulesPromoteReleaseRuleDestinationTargetId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandClouddeployAutomationRulesPromoteReleaseRuleDestinationPhase(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandClouddeployAutomationRulesAdvanceRolloutRule(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedId, err := expandClouddeployAutomationRulesAdvanceRolloutRuleId(original["id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["id"] = transformedId
	}

	transformedWait, err := expandClouddeployAutomationRulesAdvanceRolloutRuleWait(original["wait"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedWait); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["wait"] = transformedWait
	}

	transformedSourcePhases, err := expandClouddeployAutomationRulesAdvanceRolloutRuleSourcePhases(original["source_phases"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSourcePhases); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["sourcePhases"] = transformedSourcePhases
	}

	return transformed, nil
}

func expandClouddeployAutomationRulesAdvanceRolloutRuleId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandClouddeployAutomationRulesAdvanceRolloutRuleWait(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandClouddeployAutomationRulesAdvanceRolloutRuleSourcePhases(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandClouddeployAutomationEffectiveAnnotations(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandClouddeployAutomationEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
