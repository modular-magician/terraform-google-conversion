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

import "fmt"

func GetIapAppEngineServiceIamPolicyCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	return newBucketIamAsset(d, config, expandIamPolicyBindings)
}

func GetIapAppEngineServiceIamBindingCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	return newBucketIamAsset(d, config, expandIamRoleBindings)
}

func GetIapAppEngineServiceIamMemberCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	return newIapAppEngineServiceIamAsset(d, config, expandIamMemberBindings)
}

func MergeIapAppEngineServiceIamPolicy(existing, incoming Asset) Asset {
	existing.IAMPolicy = incoming.IAMPolicy
	return existing
}

func MergeIapAppEngineServiceIamBinding(existing, incoming Asset) Asset {
	return mergeIamAssets(existing, incoming, mergeAuthoritativeBindings)
}

func MergeIapAppEngineServiceIamBindingDelete(existing, incoming Asset) Asset {
	return mergeDeleteIamAssets(existing, incoming, mergeDeleteAuthoritativeBindings)
}

func MergeIapAppEngineServiceIamMember(existing, incoming Asset) Asset {
	return mergeIamAssets(existing, incoming, mergeAdditiveBindings)
}

func MergeIapAppEngineServiceIamMemberDelete(existing, incoming Asset) Asset {
	return mergeDeleteIamAssets(existing, incoming, mergeDeleteAdditiveBindings)
}

func newIapAppEngineServiceIamAsset(
	d TerraformResourceData,
	config *Config,
	expandBindings func(d TerraformResourceData) ([]IAMBinding, error),
) ([]Asset, error) {
	bindings, err := expandBindings(d)
	if err != nil {
		return []Asset{}, fmt.Errorf("expanding bindings: %v", err)
	}

	name, err := assetName(d, config, "//storage.googleapis.com/{{iap}}")
	if err != nil {
		return []Asset{}, err
	}

	return []Asset{{
		Name: name,
		Type: "storage.googleapis.com/IapAppEngineService",
		IAMPolicy: &IAMPolicy{
			Bindings: bindings,
		},
	}}, nil
}

func FetchIapAppEngineServiceIamPolicy(d TerraformResourceData, config *Config) (Asset, error) {
	// Check if the identity field returns a value
	if _, ok := d.GetOk("{{iap}}"); !ok {
		return Asset{}, ErrEmptyIdentityField
	}

	return fetchIamPolicy(
		IapAppEngineServiceIamUpdaterProducer,
		d,
		config,
		"//storage.googleapis.com/{{iap}}",
		"storage.googleapis.com/IapAppEngineService",
	)
}
