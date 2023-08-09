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

package cloudidentity

import (
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const CloudIdentityGroupMembershipAssetType string = "cloudidentity.googleapis.com/GroupMembership"

func ResourceConverterCloudIdentityGroupMembership() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType: CloudIdentityGroupMembershipAssetType,
		Convert:   GetCloudIdentityGroupMembershipCaiObject,
	}
}

func GetCloudIdentityGroupMembershipCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	name, err := tpgresource.AssetName(d, config, "//cloudidentity.googleapis.com/{{name}}")
	if err != nil {
		return []tpgresource.Asset{}, err
	}
	if obj, err := GetCloudIdentityGroupMembershipApiObject(d, config); err == nil {
		return []tpgresource.Asset{{
			Name: name,
			Type: CloudIdentityGroupMembershipAssetType,
			Resource: &tpgresource.AssetResource{
				Version:              "v1beta1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/cloudidentity/v1beta1/rest",
				DiscoveryName:        "GroupMembership",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []tpgresource.Asset{}, err
	}
}

func GetCloudIdentityGroupMembershipApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	memberKeyProp, err := expandCloudIdentityGroupMembershipMemberKey(d.Get("member_key"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("member_key"); !tpgresource.IsEmptyValue(reflect.ValueOf(memberKeyProp)) && (ok || !reflect.DeepEqual(v, memberKeyProp)) {
		obj["memberKey"] = memberKeyProp
	}
	preferredMemberKeyProp, err := expandCloudIdentityGroupMembershipPreferredMemberKey(d.Get("preferred_member_key"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("preferred_member_key"); !tpgresource.IsEmptyValue(reflect.ValueOf(preferredMemberKeyProp)) && (ok || !reflect.DeepEqual(v, preferredMemberKeyProp)) {
		obj["preferredMemberKey"] = preferredMemberKeyProp
	}
	rolesProp, err := expandCloudIdentityGroupMembershipRoles(d.Get("roles"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("roles"); !tpgresource.IsEmptyValue(reflect.ValueOf(rolesProp)) && (ok || !reflect.DeepEqual(v, rolesProp)) {
		obj["roles"] = rolesProp
	}

	return obj, nil
}

func expandCloudIdentityGroupMembershipMemberKey(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedId, err := expandCloudIdentityGroupMembershipMemberKeyId(original["id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["id"] = transformedId
	}

	transformedNamespace, err := expandCloudIdentityGroupMembershipMemberKeyNamespace(original["namespace"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNamespace); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["namespace"] = transformedNamespace
	}

	return transformed, nil
}

func expandCloudIdentityGroupMembershipMemberKeyId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudIdentityGroupMembershipMemberKeyNamespace(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudIdentityGroupMembershipPreferredMemberKey(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedId, err := expandCloudIdentityGroupMembershipPreferredMemberKeyId(original["id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["id"] = transformedId
	}

	transformedNamespace, err := expandCloudIdentityGroupMembershipPreferredMemberKeyNamespace(original["namespace"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNamespace); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["namespace"] = transformedNamespace
	}

	return transformed, nil
}

func expandCloudIdentityGroupMembershipPreferredMemberKeyId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudIdentityGroupMembershipPreferredMemberKeyNamespace(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudIdentityGroupMembershipRoles(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedName, err := expandCloudIdentityGroupMembershipRolesName(original["name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["name"] = transformedName
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandCloudIdentityGroupMembershipRolesName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
