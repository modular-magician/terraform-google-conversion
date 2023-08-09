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

package certificatemanager

import (
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func certManagerDefaultScopeDiffSuppress(_, old, new string, diff *schema.ResourceData) bool {
	if old == "" && new == "DEFAULT" || old == "DEFAULT" && new == "" {
		return true
	}
	return false
}

const CertificateManagerCertificateAssetType string = "certificatemanager.googleapis.com/Certificate"

func ResourceConverterCertificateManagerCertificate() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType: CertificateManagerCertificateAssetType,
		Convert:   GetCertificateManagerCertificateCaiObject,
	}
}

func GetCertificateManagerCertificateCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	name, err := tpgresource.AssetName(d, config, "//certificatemanager.googleapis.com/projects/{{project}}/locations/{{location}}/certificates/{{name}}")
	if err != nil {
		return []tpgresource.Asset{}, err
	}
	if obj, err := GetCertificateManagerCertificateApiObject(d, config); err == nil {
		return []tpgresource.Asset{{
			Name: name,
			Type: CertificateManagerCertificateAssetType,
			Resource: &tpgresource.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/certificatemanager/v1/rest",
				DiscoveryName:        "Certificate",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []tpgresource.Asset{}, err
	}
}

func GetCertificateManagerCertificateApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	descriptionProp, err := expandCertificateManagerCertificateDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	labelsProp, err := expandCertificateManagerCertificateLabels(d.Get("labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	scopeProp, err := expandCertificateManagerCertificateScope(d.Get("scope"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("scope"); !tpgresource.IsEmptyValue(reflect.ValueOf(scopeProp)) && (ok || !reflect.DeepEqual(v, scopeProp)) {
		obj["scope"] = scopeProp
	}
	selfManagedProp, err := expandCertificateManagerCertificateSelfManaged(d.Get("self_managed"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("self_managed"); !tpgresource.IsEmptyValue(reflect.ValueOf(selfManagedProp)) && (ok || !reflect.DeepEqual(v, selfManagedProp)) {
		obj["selfManaged"] = selfManagedProp
	}
	managedProp, err := expandCertificateManagerCertificateManaged(d.Get("managed"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("managed"); !tpgresource.IsEmptyValue(reflect.ValueOf(managedProp)) && (ok || !reflect.DeepEqual(v, managedProp)) {
		obj["managed"] = managedProp
	}

	return obj, nil
}

func expandCertificateManagerCertificateDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCertificateManagerCertificateLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandCertificateManagerCertificateScope(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCertificateManagerCertificateSelfManaged(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedCertificatePem, err := expandCertificateManagerCertificateSelfManagedCertificatePem(original["certificate_pem"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCertificatePem); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["certificatePem"] = transformedCertificatePem
	}

	transformedPrivateKeyPem, err := expandCertificateManagerCertificateSelfManagedPrivateKeyPem(original["private_key_pem"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPrivateKeyPem); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["privateKeyPem"] = transformedPrivateKeyPem
	}

	transformedPemCertificate, err := expandCertificateManagerCertificateSelfManagedPemCertificate(original["pem_certificate"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPemCertificate); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["pemCertificate"] = transformedPemCertificate
	}

	transformedPemPrivateKey, err := expandCertificateManagerCertificateSelfManagedPemPrivateKey(original["pem_private_key"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPemPrivateKey); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["pemPrivateKey"] = transformedPemPrivateKey
	}

	return transformed, nil
}

func expandCertificateManagerCertificateSelfManagedCertificatePem(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCertificateManagerCertificateSelfManagedPrivateKeyPem(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCertificateManagerCertificateSelfManagedPemCertificate(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCertificateManagerCertificateSelfManagedPemPrivateKey(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCertificateManagerCertificateManaged(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedDomains, err := expandCertificateManagerCertificateManagedDomains(original["domains"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDomains); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["domains"] = transformedDomains
	}

	transformedDnsAuthorizations, err := expandCertificateManagerCertificateManagedDnsAuthorizations(original["dns_authorizations"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDnsAuthorizations); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["dnsAuthorizations"] = transformedDnsAuthorizations
	}

	transformedIssuanceConfig, err := expandCertificateManagerCertificateManagedIssuanceConfig(original["issuance_config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedIssuanceConfig); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["issuanceConfig"] = transformedIssuanceConfig
	}

	transformedState, err := expandCertificateManagerCertificateManagedState(original["state"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedState); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["state"] = transformedState
	}

	transformedProvisioningIssue, err := expandCertificateManagerCertificateManagedProvisioningIssue(original["provisioning_issue"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedProvisioningIssue); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["provisioningIssue"] = transformedProvisioningIssue
	}

	transformedAuthorizationAttemptInfo, err := expandCertificateManagerCertificateManagedAuthorizationAttemptInfo(original["authorization_attempt_info"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAuthorizationAttemptInfo); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["authorizationAttemptInfo"] = transformedAuthorizationAttemptInfo
	}

	return transformed, nil
}

func expandCertificateManagerCertificateManagedDomains(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCertificateManagerCertificateManagedDnsAuthorizations(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCertificateManagerCertificateManagedIssuanceConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCertificateManagerCertificateManagedState(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCertificateManagerCertificateManagedProvisioningIssue(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedReason, err := expandCertificateManagerCertificateManagedProvisioningIssueReason(original["reason"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedReason); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["reason"] = transformedReason
	}

	transformedDetails, err := expandCertificateManagerCertificateManagedProvisioningIssueDetails(original["details"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDetails); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["details"] = transformedDetails
	}

	return transformed, nil
}

func expandCertificateManagerCertificateManagedProvisioningIssueReason(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCertificateManagerCertificateManagedProvisioningIssueDetails(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCertificateManagerCertificateManagedAuthorizationAttemptInfo(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedDomain, err := expandCertificateManagerCertificateManagedAuthorizationAttemptInfoDomain(original["domain"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDomain); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["domain"] = transformedDomain
		}

		transformedState, err := expandCertificateManagerCertificateManagedAuthorizationAttemptInfoState(original["state"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedState); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["state"] = transformedState
		}

		transformedFailureReason, err := expandCertificateManagerCertificateManagedAuthorizationAttemptInfoFailureReason(original["failure_reason"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedFailureReason); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["failureReason"] = transformedFailureReason
		}

		transformedDetails, err := expandCertificateManagerCertificateManagedAuthorizationAttemptInfoDetails(original["details"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDetails); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["details"] = transformedDetails
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandCertificateManagerCertificateManagedAuthorizationAttemptInfoDomain(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCertificateManagerCertificateManagedAuthorizationAttemptInfoState(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCertificateManagerCertificateManagedAuthorizationAttemptInfoFailureReason(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCertificateManagerCertificateManagedAuthorizationAttemptInfoDetails(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
