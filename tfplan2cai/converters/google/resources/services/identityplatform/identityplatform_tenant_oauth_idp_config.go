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

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const IdentityPlatformTenantOauthIdpConfigAssetType string = "identitytoolkit.googleapis.com/TenantOauthIdpConfig"

func ResourceConverterIdentityPlatformTenantOauthIdpConfig() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType: IdentityPlatformTenantOauthIdpConfigAssetType,
		Convert:   GetIdentityPlatformTenantOauthIdpConfigCaiObject,
	}
}

func GetIdentityPlatformTenantOauthIdpConfigCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	name, err := tpgresource.AssetName(d, config, "//identitytoolkit.googleapis.com/projects/{{project}}/tenants/{{tenant}}/oauthIdpConfigs/{{name}}")
	if err != nil {
		return []tpgresource.Asset{}, err
	}
	if obj, err := GetIdentityPlatformTenantOauthIdpConfigApiObject(d, config); err == nil {
		return []tpgresource.Asset{{
			Name: name,
			Type: IdentityPlatformTenantOauthIdpConfigAssetType,
			Resource: &tpgresource.AssetResource{
				Version:              "v2",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/identitytoolkit/v2/rest",
				DiscoveryName:        "TenantOauthIdpConfig",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []tpgresource.Asset{}, err
	}
}

func GetIdentityPlatformTenantOauthIdpConfigApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandIdentityPlatformTenantOauthIdpConfigName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	displayNameProp, err := expandIdentityPlatformTenantOauthIdpConfigDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	enabledProp, err := expandIdentityPlatformTenantOauthIdpConfigEnabled(d.Get("enabled"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("enabled"); !tpgresource.IsEmptyValue(reflect.ValueOf(enabledProp)) && (ok || !reflect.DeepEqual(v, enabledProp)) {
		obj["enabled"] = enabledProp
	}
	issuerProp, err := expandIdentityPlatformTenantOauthIdpConfigIssuer(d.Get("issuer"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("issuer"); !tpgresource.IsEmptyValue(reflect.ValueOf(issuerProp)) && (ok || !reflect.DeepEqual(v, issuerProp)) {
		obj["issuer"] = issuerProp
	}
	clientIdProp, err := expandIdentityPlatformTenantOauthIdpConfigClientId(d.Get("client_id"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("client_id"); !tpgresource.IsEmptyValue(reflect.ValueOf(clientIdProp)) && (ok || !reflect.DeepEqual(v, clientIdProp)) {
		obj["clientId"] = clientIdProp
	}
	clientSecretProp, err := expandIdentityPlatformTenantOauthIdpConfigClientSecret(d.Get("client_secret"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("client_secret"); !tpgresource.IsEmptyValue(reflect.ValueOf(clientSecretProp)) && (ok || !reflect.DeepEqual(v, clientSecretProp)) {
		obj["clientSecret"] = clientSecretProp
	}

	return obj, nil
}

func expandIdentityPlatformTenantOauthIdpConfigName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformTenantOauthIdpConfigDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformTenantOauthIdpConfigEnabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformTenantOauthIdpConfigIssuer(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformTenantOauthIdpConfigClientId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformTenantOauthIdpConfigClientSecret(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
