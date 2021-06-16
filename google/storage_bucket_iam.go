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

func GetStorageBucketIamPolicyCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	return newBucketIamAsset(d, config, expandIamPolicyBindings)
}

func GetStorageBucketIamBindingCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	return newBucketIamAsset(d, config, expandIamRoleBindings)
}

func GetStorageBucketIamMemberCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	return newStorageBucketIamAsset(d, config, expandIamMemberBindings)
}

func MergeStorageBucketIamPolicy(existing, incoming Asset) Asset {
	existing.IAMPolicy = incoming.IAMPolicy
	return existing
}

func MergeStorageBucketIamBinding(existing, incoming Asset) Asset {
	return mergeIamAssets(existing, incoming, mergeAuthoritativeBindings)
}

func MergeStorageBucketIamBindingDelete(existing, incoming Asset) Asset {
	return mergeDeleteIamAssets(existing, incoming, mergeDeleteAuthoritativeBindings)
}

func MergeStorageBucketIamMember(existing, incoming Asset) Asset {
	return mergeIamAssets(existing, incoming, mergeAdditiveBindings)
}

func MergeStorageBucketIamMemberDelete(existing, incoming Asset) Asset {
	return mergeDeleteIamAssets(existing, incoming, mergeDeleteAdditiveBindings)
}

func newStorageBucketIamAsset(
	d TerraformResourceData,
	config *Config,
	expandBindings func(d TerraformResourceData) ([]IAMBinding, error),
) ([]Asset, error) {
	bindings, err := expandBindings(d)
	if err != nil {
		return []Asset{}, fmt.Errorf("expanding bindings: %v", err)
	}

	name, err := assetName(d, config, "//storage.googleapis.com/{{storage}}")
	if err != nil {
		return []Asset{}, err
	}

	return []Asset{{
		Name: name,
		Type: "storage.googleapis.com/StorageBucket",
		IAMPolicy: &IAMPolicy{
			Bindings: bindings,
		},
	}}, nil
}

func FetchStorageBucketIamPolicy(d TerraformResourceData, config *Config) (Asset, error) {
	// Check if the identity field returns a value
	if _, ok := d.GetOk("{{storage}}"); !ok {
		return Asset{}, ErrEmptyIdentityField
	}

	return fetchIamPolicy(
		StorageBucketIamUpdaterProducer,
		d,
		config,
		"//storage.googleapis.com/{{storage}}",
		"storage.googleapis.com/StorageBucket",
	)
}
