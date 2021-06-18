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

func GetComputeInstanceIamPolicyCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	return newComputeInstanceIamAsset(d, config, expandIamPolicyBindings)
}

func GetComputeInstanceIamBindingCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	return newComputeInstanceIamAsset(d, config, expandIamRoleBindings)
}

func GetComputeInstanceIamMemberCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	return newComputeInstanceIamAsset(d, config, expandIamMemberBindings)
}

func MergeComputeInstanceIamPolicy(existing, incoming Asset) Asset {
	existing.IAMPolicy = incoming.IAMPolicy
	return existing
}

func MergeComputeInstanceIamBinding(existing, incoming Asset) Asset {
	return mergeIamAssets(existing, incoming, mergeAuthoritativeBindings)
}

func MergeComputeInstanceIamBindingDelete(existing, incoming Asset) Asset {
	return mergeDeleteIamAssets(existing, incoming, mergeDeleteAuthoritativeBindings)
}

func MergeComputeInstanceIamMember(existing, incoming Asset) Asset {
	return mergeIamAssets(existing, incoming, mergeAdditiveBindings)
}

func MergeComputeInstanceIamMemberDelete(existing, incoming Asset) Asset {
	return mergeDeleteIamAssets(existing, incoming, mergeDeleteAdditiveBindings)
}

func newComputeInstanceIamAsset(
	d TerraformResourceData,
	config *Config,
	expandBindings func(d TerraformResourceData) ([]IAMBinding, error),
) ([]Asset, error) {
	bindings, err := expandBindings(d)
	if err != nil {
		return []Asset{}, fmt.Errorf("expanding bindings: %v", err)
	}

	name, err := assetName(d, config, "//compute.googleapis.com/{{instance}}")
	if err != nil {
		return []Asset{}, err
	}

	return []Asset{{
		Name: name,
		Type: "compute.googleapis.com/Instance",
		IAMPolicy: &IAMPolicy{
			Bindings: bindings,
		},
	}}, nil
}

func FetchComputeInstanceIamPolicy(d TerraformResourceData, config *Config) (Asset, error) {
	// Check if the identity field returns a value
	if _, ok := d.GetOk("{{instance}}"); !ok {
		return Asset{}, ErrEmptyIdentityField
	}

	return fetchIamPolicy(
		ComputeInstanceIamUpdaterProducer,
		d,
		config,
		"//compute.googleapis.com/{{instance}}",
		"compute.googleapis.com/Instance",
	)
}
