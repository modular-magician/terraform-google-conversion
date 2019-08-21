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
	"regexp"
)

func defaultBinaryAuthorizationPolicy(project string) map[string]interface{} {
	return map[string]interface{}{
		"name": fmt.Sprintf("projects/%s/policy", project),
		"admissionWhitelistPatterns": []interface{}{
			map[string]interface{}{
				"namePattern": "gcr.io/google_containers/*",
			},
		},
		"defaultAdmissionRule": map[string]interface{}{
			"evaluationMode":  "ALWAYS_ALLOW",
			"enforcementMode": "ENFORCED_BLOCK_AND_AUDIT_LOG",
		},
	}
}

func GetBinaryAuthorizationPolicyCaiObject(d TerraformResourceData, config *Config) (Asset, error) {
	name, err := assetName(d, config, "//binaryauthorization.googleapis.com/projects/{{project}}/policy")
	if err != nil {
		return Asset{}, err
	}
	if obj, err := GetBinaryAuthorizationPolicyApiObject(d, config); err == nil {
		return Asset{
			Name: name,
			Type: "binaryauthorization.googleapis.com/Policy",
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/binaryauthorization/v1/rest",
				DiscoveryName:        "Policy",
				Data:                 obj,
			},
		}, nil
	} else {
		return Asset{}, err
	}
}

func GetBinaryAuthorizationPolicyApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	descriptionProp, err := expandBinaryAuthorizationPolicyDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	globalPolicyEvaluationModeProp, err := expandBinaryAuthorizationPolicyGlobalPolicyEvaluationMode(d.Get("global_policy_evaluation_mode"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("global_policy_evaluation_mode"); !isEmptyValue(reflect.ValueOf(globalPolicyEvaluationModeProp)) && (ok || !reflect.DeepEqual(v, globalPolicyEvaluationModeProp)) {
		obj["globalPolicyEvaluationMode"] = globalPolicyEvaluationModeProp
	}
	admissionWhitelistPatternsProp, err := expandBinaryAuthorizationPolicyAdmissionWhitelistPatterns(d.Get("admission_whitelist_patterns"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("admission_whitelist_patterns"); !isEmptyValue(reflect.ValueOf(admissionWhitelistPatternsProp)) && (ok || !reflect.DeepEqual(v, admissionWhitelistPatternsProp)) {
		obj["admissionWhitelistPatterns"] = admissionWhitelistPatternsProp
	}
	clusterAdmissionRulesProp, err := expandBinaryAuthorizationPolicyClusterAdmissionRules(d.Get("cluster_admission_rules"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("cluster_admission_rules"); !isEmptyValue(reflect.ValueOf(clusterAdmissionRulesProp)) && (ok || !reflect.DeepEqual(v, clusterAdmissionRulesProp)) {
		obj["clusterAdmissionRules"] = clusterAdmissionRulesProp
	}
	defaultAdmissionRuleProp, err := expandBinaryAuthorizationPolicyDefaultAdmissionRule(d.Get("default_admission_rule"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("default_admission_rule"); !isEmptyValue(reflect.ValueOf(defaultAdmissionRuleProp)) && (ok || !reflect.DeepEqual(v, defaultAdmissionRuleProp)) {
		obj["defaultAdmissionRule"] = defaultAdmissionRuleProp
	}

	return obj, nil
}

func expandBinaryAuthorizationPolicyDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBinaryAuthorizationPolicyGlobalPolicyEvaluationMode(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBinaryAuthorizationPolicyAdmissionWhitelistPatterns(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedNamePattern, err := expandBinaryAuthorizationPolicyAdmissionWhitelistPatternsNamePattern(original["name_pattern"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedNamePattern); val.IsValid() && !isEmptyValue(val) {
			transformed["namePattern"] = transformedNamePattern
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandBinaryAuthorizationPolicyAdmissionWhitelistPatternsNamePattern(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBinaryAuthorizationPolicyClusterAdmissionRules(v interface{}, d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	if v == nil {
		return map[string]interface{}{}, nil
	}
	m := make(map[string]interface{})
	for _, raw := range v.(*schema.Set).List() {
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedEvaluationMode, err := expandBinaryAuthorizationPolicyClusterAdmissionRulesEvaluationMode(original["evaluation_mode"], d, config)
		if err != nil {
			return nil, err
		}
		transformed["evaluationMode"] = transformedEvaluationMode
		transformedRequireAttestationsBy, err := expandBinaryAuthorizationPolicyClusterAdmissionRulesRequireAttestationsBy(original["require_attestations_by"], d, config)
		if err != nil {
			return nil, err
		}
		transformed["requireAttestationsBy"] = transformedRequireAttestationsBy
		transformedEnforcementMode, err := expandBinaryAuthorizationPolicyClusterAdmissionRulesEnforcementMode(original["enforcement_mode"], d, config)
		if err != nil {
			return nil, err
		}
		transformed["enforcementMode"] = transformedEnforcementMode

		m[original["cluster"].(string)] = transformed
	}
	return m, nil
}

func expandBinaryAuthorizationPolicyClusterAdmissionRulesEvaluationMode(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBinaryAuthorizationPolicyClusterAdmissionRulesRequireAttestationsBy(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	r := regexp.MustCompile("projects/(.+)/attestors/(.+)")

	// It's possible that all entries in the list will specify a project, in
	// which case the user wouldn't necessarily have to specify a provider
	// project.
	var project string
	var err error
	for _, s := range v.(*schema.Set).List() {
		if !r.MatchString(s.(string)) {
			project, err = getProject(d, config)
			if err != nil {
				return []interface{}{}, err
			}
			break
		}
	}

	return convertAndMapStringArr(v.(*schema.Set).List(), func(s string) string {
		if r.MatchString(s) {
			return s
		}

		return fmt.Sprintf("projects/%s/attestors/%s", project, s)
	}), nil
}

func expandBinaryAuthorizationPolicyClusterAdmissionRulesEnforcementMode(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBinaryAuthorizationPolicyDefaultAdmissionRule(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedEvaluationMode, err := expandBinaryAuthorizationPolicyDefaultAdmissionRuleEvaluationMode(original["evaluation_mode"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEvaluationMode); val.IsValid() && !isEmptyValue(val) {
		transformed["evaluationMode"] = transformedEvaluationMode
	}

	transformedRequireAttestationsBy, err := expandBinaryAuthorizationPolicyDefaultAdmissionRuleRequireAttestationsBy(original["require_attestations_by"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRequireAttestationsBy); val.IsValid() && !isEmptyValue(val) {
		transformed["requireAttestationsBy"] = transformedRequireAttestationsBy
	}

	transformedEnforcementMode, err := expandBinaryAuthorizationPolicyDefaultAdmissionRuleEnforcementMode(original["enforcement_mode"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEnforcementMode); val.IsValid() && !isEmptyValue(val) {
		transformed["enforcementMode"] = transformedEnforcementMode
	}

	return transformed, nil
}

func expandBinaryAuthorizationPolicyDefaultAdmissionRuleEvaluationMode(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBinaryAuthorizationPolicyDefaultAdmissionRuleRequireAttestationsBy(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	r := regexp.MustCompile("projects/(.+)/attestors/(.+)")

	// It's possible that all entries in the list will specify a project, in
	// which case the user wouldn't necessarily have to specify a provider
	// project.
	var project string
	var err error
	for _, s := range v.(*schema.Set).List() {
		if !r.MatchString(s.(string)) {
			project, err = getProject(d, config)
			if err != nil {
				return []interface{}{}, err
			}
			break
		}
	}

	return convertAndMapStringArr(v.(*schema.Set).List(), func(s string) string {
		if r.MatchString(s) {
			return s
		}

		return fmt.Sprintf("projects/%s/attestors/%s", project, s)
	}), nil
}

func expandBinaryAuthorizationPolicyDefaultAdmissionRuleEnforcementMode(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
