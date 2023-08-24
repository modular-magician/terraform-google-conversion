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

package networksecurity

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const NetworkSecurityGatewaySecurityPolicyRuleAssetType string = "networksecurity.googleapis.com/GatewaySecurityPolicyRule"

func ResourceConverterNetworkSecurityGatewaySecurityPolicyRule() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: NetworkSecurityGatewaySecurityPolicyRuleAssetType,
		Convert:   GetNetworkSecurityGatewaySecurityPolicyRuleCaiObject,
	}
}

func GetNetworkSecurityGatewaySecurityPolicyRuleCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//networksecurity.googleapis.com/projects/{{project}}/locations/{{location}}/gatewaySecurityPolicies/{{gateway_security_policy}}/rules/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetNetworkSecurityGatewaySecurityPolicyRuleApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: NetworkSecurityGatewaySecurityPolicyRuleAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1beta1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/networksecurity/v1beta1/rest",
				DiscoveryName:        "GatewaySecurityPolicyRule",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetNetworkSecurityGatewaySecurityPolicyRuleApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	enabledProp, err := expandNetworkSecurityGatewaySecurityPolicyRuleEnabled(d.Get("enabled"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("enabled"); !tpgresource.IsEmptyValue(reflect.ValueOf(enabledProp)) && (ok || !reflect.DeepEqual(v, enabledProp)) {
		obj["enabled"] = enabledProp
	}
	priorityProp, err := expandNetworkSecurityGatewaySecurityPolicyRulePriority(d.Get("priority"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("priority"); !tpgresource.IsEmptyValue(reflect.ValueOf(priorityProp)) && (ok || !reflect.DeepEqual(v, priorityProp)) {
		obj["priority"] = priorityProp
	}
	descriptionProp, err := expandNetworkSecurityGatewaySecurityPolicyRuleDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	sessionMatcherProp, err := expandNetworkSecurityGatewaySecurityPolicyRuleSessionMatcher(d.Get("session_matcher"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("session_matcher"); !tpgresource.IsEmptyValue(reflect.ValueOf(sessionMatcherProp)) && (ok || !reflect.DeepEqual(v, sessionMatcherProp)) {
		obj["sessionMatcher"] = sessionMatcherProp
	}
	applicationMatcherProp, err := expandNetworkSecurityGatewaySecurityPolicyRuleApplicationMatcher(d.Get("application_matcher"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("application_matcher"); !tpgresource.IsEmptyValue(reflect.ValueOf(applicationMatcherProp)) && (ok || !reflect.DeepEqual(v, applicationMatcherProp)) {
		obj["applicationMatcher"] = applicationMatcherProp
	}
	tlsInspectionEnabledProp, err := expandNetworkSecurityGatewaySecurityPolicyRuleTlsInspectionEnabled(d.Get("tls_inspection_enabled"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("tls_inspection_enabled"); !tpgresource.IsEmptyValue(reflect.ValueOf(tlsInspectionEnabledProp)) && (ok || !reflect.DeepEqual(v, tlsInspectionEnabledProp)) {
		obj["tlsInspectionEnabled"] = tlsInspectionEnabledProp
	}
	basicProfileProp, err := expandNetworkSecurityGatewaySecurityPolicyRuleBasicProfile(d.Get("basic_profile"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("basic_profile"); !tpgresource.IsEmptyValue(reflect.ValueOf(basicProfileProp)) && (ok || !reflect.DeepEqual(v, basicProfileProp)) {
		obj["basicProfile"] = basicProfileProp
	}

	return obj, nil
}

func expandNetworkSecurityGatewaySecurityPolicyRuleEnabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkSecurityGatewaySecurityPolicyRulePriority(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkSecurityGatewaySecurityPolicyRuleDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkSecurityGatewaySecurityPolicyRuleSessionMatcher(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkSecurityGatewaySecurityPolicyRuleApplicationMatcher(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkSecurityGatewaySecurityPolicyRuleTlsInspectionEnabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkSecurityGatewaySecurityPolicyRuleBasicProfile(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
