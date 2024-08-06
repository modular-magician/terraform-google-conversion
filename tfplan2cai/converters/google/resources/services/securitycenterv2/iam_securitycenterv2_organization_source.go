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

package securitycenterv2

import (
	"fmt"

	"github.com/hashicorp/errwrap"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"google.golang.org/api/cloudresourcemanager/v1"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgiamresource"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

var SecurityCenterV2OrganizationSourceIamSchema = map[string]*schema.Schema{
	"organization": {
		Type:     schema.TypeString,
		Required: true,
		ForceNew: true,
	},
	"source": {
		Type:             schema.TypeString,
		Required:         true,
		ForceNew:         true,
		DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
	},
}

type SecurityCenterV2OrganizationSourceIamUpdater struct {
	organization string
	source       string
	d            tpgresource.TerraformResourceData
	Config       *transport_tpg.Config
}

func SecurityCenterV2OrganizationSourceIamUpdaterProducer(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (tpgiamresource.ResourceIamUpdater, error) {
	values := make(map[string]string)

	if v, ok := d.GetOk("organization"); ok {
		values["organization"] = v.(string)
	}

	if v, ok := d.GetOk("source"); ok {
		values["source"] = v.(string)
	}

	// We may have gotten either a long or short name, so attempt to parse long name if possible
	m, err := tpgresource.GetImportIdQualifiers([]string{"organizations/(?P<organization>[^/]+)/sources/(?P<source>[^/]+)", "(?P<organization>[^/]+)/(?P<source>[^/]+)", "(?P<source>[^/]+)"}, d, config, d.Get("source").(string))
	if err != nil {
		return nil, err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &SecurityCenterV2OrganizationSourceIamUpdater{
		organization: values["organization"],
		source:       values["source"],
		d:            d,
		Config:       config,
	}

	if err := d.Set("organization", u.organization); err != nil {
		return nil, fmt.Errorf("Error setting organization: %s", err)
	}
	if err := d.Set("source", u.GetResourceId()); err != nil {
		return nil, fmt.Errorf("Error setting source: %s", err)
	}

	return u, nil
}

func SecurityCenterV2OrganizationSourceIdParseFunc(d *schema.ResourceData, config *transport_tpg.Config) error {
	values := make(map[string]string)

	m, err := tpgresource.GetImportIdQualifiers([]string{"organizations/(?P<organization>[^/]+)/sources/(?P<source>[^/]+)", "(?P<organization>[^/]+)/(?P<source>[^/]+)", "(?P<source>[^/]+)"}, d, config, d.Id())
	if err != nil {
		return err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &SecurityCenterV2OrganizationSourceIamUpdater{
		organization: values["organization"],
		source:       values["source"],
		d:            d,
		Config:       config,
	}
	if err := d.Set("source", u.GetResourceId()); err != nil {
		return fmt.Errorf("Error setting source: %s", err)
	}
	d.SetId(u.GetResourceId())
	return nil
}

func (u *SecurityCenterV2OrganizationSourceIamUpdater) GetResourceIamPolicy() (*cloudresourcemanager.Policy, error) {
	url, err := u.qualifyOrganizationSourceUrl("getIamPolicy")
	if err != nil {
		return nil, err
	}

	var obj map[string]interface{}

	userAgent, err := tpgresource.GenerateUserAgentString(u.d, u.Config.UserAgent)
	if err != nil {
		return nil, err
	}

	policy, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    u.Config,
		Method:    "POST",
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
	})
	if err != nil {
		return nil, errwrap.Wrapf(fmt.Sprintf("Error retrieving IAM policy for %s: {{err}}", u.DescribeResource()), err)
	}

	out := &cloudresourcemanager.Policy{}
	err = tpgresource.Convert(policy, out)
	if err != nil {
		return nil, errwrap.Wrapf("Cannot convert a policy to a resource manager policy: {{err}}", err)
	}

	return out, nil
}

func (u *SecurityCenterV2OrganizationSourceIamUpdater) SetResourceIamPolicy(policy *cloudresourcemanager.Policy) error {
	json, err := tpgresource.ConvertToMap(policy)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	obj["policy"] = json

	url, err := u.qualifyOrganizationSourceUrl("setIamPolicy")
	if err != nil {
		return err
	}

	userAgent, err := tpgresource.GenerateUserAgentString(u.d, u.Config.UserAgent)
	if err != nil {
		return err
	}

	_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    u.Config,
		Method:    "POST",
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   u.d.Timeout(schema.TimeoutCreate),
	})
	if err != nil {
		return errwrap.Wrapf(fmt.Sprintf("Error setting IAM policy for %s: {{err}}", u.DescribeResource()), err)
	}

	return nil
}

func (u *SecurityCenterV2OrganizationSourceIamUpdater) qualifyOrganizationSourceUrl(methodIdentifier string) (string, error) {
	urlTemplate := fmt.Sprintf("{{SecurityCenterV2BasePath}}%s:%s", fmt.Sprintf("organizations/%s/sources/%s", u.organization, u.source), methodIdentifier)
	url, err := tpgresource.ReplaceVars(u.d, u.Config, urlTemplate)
	if err != nil {
		return "", err
	}
	return url, nil
}

func (u *SecurityCenterV2OrganizationSourceIamUpdater) GetResourceId() string {
	return fmt.Sprintf("organizations/%s/sources/%s", u.organization, u.source)
}

func (u *SecurityCenterV2OrganizationSourceIamUpdater) GetMutexKey() string {
	return fmt.Sprintf("iam-securitycenterv2-organizationsource-%s", u.GetResourceId())
}

func (u *SecurityCenterV2OrganizationSourceIamUpdater) DescribeResource() string {
	return fmt.Sprintf("securitycenterv2 organizationsource %q", u.GetResourceId())
}
