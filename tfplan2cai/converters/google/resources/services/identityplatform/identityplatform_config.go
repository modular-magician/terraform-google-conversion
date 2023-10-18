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

package identityplatform

import (
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const IdentityPlatformConfigAssetType string = "identitytoolkit.googleapis.com/Config"

func ResourceConverterIdentityPlatformConfig() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: IdentityPlatformConfigAssetType,
		Convert:   GetIdentityPlatformConfigCaiObject,
	}
}

func GetIdentityPlatformConfigCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//identitytoolkit.googleapis.com/projects/{{project}}/config")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetIdentityPlatformConfigApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: IdentityPlatformConfigAssetType,
			Resource: &cai.AssetResource{
				Version:              "v2",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/identitytoolkit/v2/rest",
				DiscoveryName:        "Config",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetIdentityPlatformConfigApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	autodeleteAnonymousUsersProp, err := expandIdentityPlatformConfigAutodeleteAnonymousUsers(d.Get("autodelete_anonymous_users"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("autodelete_anonymous_users"); !tpgresource.IsEmptyValue(reflect.ValueOf(autodeleteAnonymousUsersProp)) && (ok || !reflect.DeepEqual(v, autodeleteAnonymousUsersProp)) {
		obj["autodeleteAnonymousUsers"] = autodeleteAnonymousUsersProp
	}
	signInProp, err := expandIdentityPlatformConfigSignIn(d.Get("sign_in"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("sign_in"); !tpgresource.IsEmptyValue(reflect.ValueOf(signInProp)) && (ok || !reflect.DeepEqual(v, signInProp)) {
		obj["signIn"] = signInProp
	}
	blockingFunctionsProp, err := expandIdentityPlatformConfigBlockingFunctions(d.Get("blocking_functions"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("blocking_functions"); !tpgresource.IsEmptyValue(reflect.ValueOf(blockingFunctionsProp)) && (ok || !reflect.DeepEqual(v, blockingFunctionsProp)) {
		obj["blockingFunctions"] = blockingFunctionsProp
	}
	quotaProp, err := expandIdentityPlatformConfigQuota(d.Get("quota"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("quota"); !tpgresource.IsEmptyValue(reflect.ValueOf(quotaProp)) && (ok || !reflect.DeepEqual(v, quotaProp)) {
		obj["quota"] = quotaProp
	}
	authorizedDomainsProp, err := expandIdentityPlatformConfigAuthorizedDomains(d.Get("authorized_domains"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("authorized_domains"); !tpgresource.IsEmptyValue(reflect.ValueOf(authorizedDomainsProp)) && (ok || !reflect.DeepEqual(v, authorizedDomainsProp)) {
		obj["authorizedDomains"] = authorizedDomainsProp
	}
	smsRegionConfigProp, err := expandIdentityPlatformConfigSmsRegionConfig(d.Get("sms_region_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("sms_region_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(smsRegionConfigProp)) && (ok || !reflect.DeepEqual(v, smsRegionConfigProp)) {
		obj["smsRegionConfig"] = smsRegionConfigProp
	}

	return obj, nil
}

func expandIdentityPlatformConfigAutodeleteAnonymousUsers(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformConfigSignIn(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedEmail, err := expandIdentityPlatformConfigSignInEmail(original["email"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEmail); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["email"] = transformedEmail
	}

	transformedPhoneNumber, err := expandIdentityPlatformConfigSignInPhoneNumber(original["phone_number"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPhoneNumber); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["phoneNumber"] = transformedPhoneNumber
	}

	transformedAnonymous, err := expandIdentityPlatformConfigSignInAnonymous(original["anonymous"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAnonymous); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["anonymous"] = transformedAnonymous
	}

	transformedAllowDuplicateEmails, err := expandIdentityPlatformConfigSignInAllowDuplicateEmails(original["allow_duplicate_emails"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAllowDuplicateEmails); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["allowDuplicateEmails"] = transformedAllowDuplicateEmails
	}

	return transformed, nil
}

func expandIdentityPlatformConfigSignInEmail(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedEnabled, err := expandIdentityPlatformConfigSignInEmailEnabled(original["enabled"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEnabled); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["enabled"] = transformedEnabled
	}

	transformedPasswordRequired, err := expandIdentityPlatformConfigSignInEmailPasswordRequired(original["password_required"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPasswordRequired); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["passwordRequired"] = transformedPasswordRequired
	}

	return transformed, nil
}

func expandIdentityPlatformConfigSignInEmailEnabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformConfigSignInEmailPasswordRequired(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformConfigSignInPhoneNumber(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedEnabled, err := expandIdentityPlatformConfigSignInPhoneNumberEnabled(original["enabled"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEnabled); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["enabled"] = transformedEnabled
	}

	transformedTestPhoneNumbers, err := expandIdentityPlatformConfigSignInPhoneNumberTestPhoneNumbers(original["test_phone_numbers"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTestPhoneNumbers); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["testPhoneNumbers"] = transformedTestPhoneNumbers
	}

	return transformed, nil
}

func expandIdentityPlatformConfigSignInPhoneNumberEnabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformConfigSignInPhoneNumberTestPhoneNumbers(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandIdentityPlatformConfigSignInAnonymous(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedEnabled, err := expandIdentityPlatformConfigSignInAnonymousEnabled(original["enabled"], d, config)
	if err != nil {
		return nil, err
	} else {
		transformed["enabled"] = transformedEnabled
	}

	return transformed, nil
}

func expandIdentityPlatformConfigSignInAnonymousEnabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformConfigSignInAllowDuplicateEmails(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformConfigBlockingFunctions(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedTriggers, err := expandIdentityPlatformConfigBlockingFunctionsTriggers(original["triggers"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTriggers); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["triggers"] = transformedTriggers
	}

	transformedForwardInboundCredentials, err := expandIdentityPlatformConfigBlockingFunctionsForwardInboundCredentials(original["forward_inbound_credentials"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedForwardInboundCredentials); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["forwardInboundCredentials"] = transformedForwardInboundCredentials
	}

	return transformed, nil
}

func expandIdentityPlatformConfigBlockingFunctionsTriggers(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	if v == nil {
		return map[string]interface{}{}, nil
	}
	m := make(map[string]interface{})
	for _, raw := range v.(*schema.Set).List() {
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedFunctionUri, err := expandIdentityPlatformConfigBlockingFunctionsTriggersFunctionUri(original["function_uri"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedFunctionUri); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["functionUri"] = transformedFunctionUri
		}

		transformedEventType, err := tpgresource.ExpandString(original["event_type"], d, config)
		if err != nil {
			return nil, err
		}
		m[transformedEventType] = transformed
	}
	return m, nil
}

func expandIdentityPlatformConfigBlockingFunctionsTriggersFunctionUri(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformConfigBlockingFunctionsForwardInboundCredentials(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedIdToken, err := expandIdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsIdToken(original["id_token"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedIdToken); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["idToken"] = transformedIdToken
	}

	transformedAccessToken, err := expandIdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsAccessToken(original["access_token"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAccessToken); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["accessToken"] = transformedAccessToken
	}

	transformedRefreshToken, err := expandIdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsRefreshToken(original["refresh_token"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRefreshToken); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["refreshToken"] = transformedRefreshToken
	}

	return transformed, nil
}

func expandIdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsIdToken(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsAccessToken(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsRefreshToken(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformConfigQuota(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedSignUpQuotaConfig, err := expandIdentityPlatformConfigQuotaSignUpQuotaConfig(original["sign_up_quota_config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSignUpQuotaConfig); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["signUpQuotaConfig"] = transformedSignUpQuotaConfig
	}

	return transformed, nil
}

func expandIdentityPlatformConfigQuotaSignUpQuotaConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedQuota, err := expandIdentityPlatformConfigQuotaSignUpQuotaConfigQuota(original["quota"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedQuota); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["quota"] = transformedQuota
	}

	transformedStartTime, err := expandIdentityPlatformConfigQuotaSignUpQuotaConfigStartTime(original["start_time"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedStartTime); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["startTime"] = transformedStartTime
	}

	transformedQuotaDuration, err := expandIdentityPlatformConfigQuotaSignUpQuotaConfigQuotaDuration(original["quota_duration"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedQuotaDuration); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["quotaDuration"] = transformedQuotaDuration
	}

	return transformed, nil
}

func expandIdentityPlatformConfigQuotaSignUpQuotaConfigQuota(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformConfigQuotaSignUpQuotaConfigStartTime(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformConfigQuotaSignUpQuotaConfigQuotaDuration(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformConfigAuthorizedDomains(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformConfigSmsRegionConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedAllowByDefault, err := expandIdentityPlatformConfigSmsRegionConfigAllowByDefault(original["allow_by_default"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAllowByDefault); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["allowByDefault"] = transformedAllowByDefault
	}

	transformedAllowlistOnly, err := expandIdentityPlatformConfigSmsRegionConfigAllowlistOnly(original["allowlist_only"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAllowlistOnly); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["allowlistOnly"] = transformedAllowlistOnly
	}

	return transformed, nil
}

func expandIdentityPlatformConfigSmsRegionConfigAllowByDefault(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedDisallowedRegions, err := expandIdentityPlatformConfigSmsRegionConfigAllowByDefaultDisallowedRegions(original["disallowed_regions"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDisallowedRegions); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["disallowedRegions"] = transformedDisallowedRegions
	}

	return transformed, nil
}

func expandIdentityPlatformConfigSmsRegionConfigAllowByDefaultDisallowedRegions(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformConfigSmsRegionConfigAllowlistOnly(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedAllowedRegions, err := expandIdentityPlatformConfigSmsRegionConfigAllowlistOnlyAllowedRegions(original["allowed_regions"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAllowedRegions); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["allowedRegions"] = transformedAllowedRegions
	}

	return transformed, nil
}

func expandIdentityPlatformConfigSmsRegionConfigAllowlistOnlyAllowedRegions(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
