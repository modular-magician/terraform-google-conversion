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

package google

import (
	"fmt"

	"github.com/hashicorp/errwrap"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"google.golang.org/api/cloudresourcemanager/v1"

	transport_tpg "github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/transport"
)

var VertexAIFeaturestoreIamSchema = map[string]*schema.Schema{
	"project": {
		Type:     schema.TypeString,
		Computed: true,
		Optional: true,
		ForceNew: true,
	},
	"region": {
		Type:     schema.TypeString,
		Computed: true,
		Optional: true,
		ForceNew: true,
	},
	"featurestore": {
		Type:             schema.TypeString,
		Required:         true,
		ForceNew:         true,
		DiffSuppressFunc: compareSelfLinkOrResourceName,
	},
}

type VertexAIFeaturestoreIamUpdater struct {
	project      string
	region       string
	featurestore string
	d            TerraformResourceData
	Config       *transport_tpg.Config
}

func VertexAIFeaturestoreIamUpdaterProducer(d TerraformResourceData, config *transport_tpg.Config) (ResourceIamUpdater, error) {
	values := make(map[string]string)

	project, _ := getProject(d, config)
	if project != "" {
		if err := d.Set("project", project); err != nil {
			return nil, fmt.Errorf("Error setting project: %s", err)
		}
	}
	values["project"] = project
	region, _ := getRegion(d, config)
	if region != "" {
		if err := d.Set("region", region); err != nil {
			return nil, fmt.Errorf("Error setting region: %s", err)
		}
	}
	values["region"] = region
	if v, ok := d.GetOk("featurestore"); ok {
		values["featurestore"] = v.(string)
	}

	// We may have gotten either a long or short name, so attempt to parse long name if possible
	m, err := getImportIdQualifiers([]string{"projects/(?P<project>[^/]+)/locations/(?P<region>[^/]+)/featurestores/(?P<featurestore>[^/]+)", "(?P<project>[^/]+)/(?P<region>[^/]+)/(?P<featurestore>[^/]+)", "(?P<region>[^/]+)/(?P<featurestore>[^/]+)", "(?P<featurestore>[^/]+)"}, d, config, d.Get("featurestore").(string))
	if err != nil {
		return nil, err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &VertexAIFeaturestoreIamUpdater{
		project:      values["project"],
		region:       values["region"],
		featurestore: values["featurestore"],
		d:            d,
		Config:       config,
	}

	if err := d.Set("project", u.project); err != nil {
		return nil, fmt.Errorf("Error setting project: %s", err)
	}
	if err := d.Set("region", u.region); err != nil {
		return nil, fmt.Errorf("Error setting region: %s", err)
	}
	if err := d.Set("featurestore", u.GetResourceId()); err != nil {
		return nil, fmt.Errorf("Error setting featurestore: %s", err)
	}

	return u, nil
}

func VertexAIFeaturestoreIdParseFunc(d *schema.ResourceData, config *transport_tpg.Config) error {
	values := make(map[string]string)

	project, _ := getProject(d, config)
	if project != "" {
		values["project"] = project
	}

	region, _ := getRegion(d, config)
	if region != "" {
		values["region"] = region
	}

	m, err := getImportIdQualifiers([]string{"projects/(?P<project>[^/]+)/locations/(?P<region>[^/]+)/featurestores/(?P<featurestore>[^/]+)", "(?P<project>[^/]+)/(?P<region>[^/]+)/(?P<featurestore>[^/]+)", "(?P<region>[^/]+)/(?P<featurestore>[^/]+)", "(?P<featurestore>[^/]+)"}, d, config, d.Id())
	if err != nil {
		return err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &VertexAIFeaturestoreIamUpdater{
		project:      values["project"],
		region:       values["region"],
		featurestore: values["featurestore"],
		d:            d,
		Config:       config,
	}
	if err := d.Set("featurestore", u.GetResourceId()); err != nil {
		return fmt.Errorf("Error setting featurestore: %s", err)
	}
	d.SetId(u.GetResourceId())
	return nil
}

func (u *VertexAIFeaturestoreIamUpdater) GetResourceIamPolicy() (*cloudresourcemanager.Policy, error) {
	url, err := u.qualifyFeaturestoreUrl("getIamPolicy")
	if err != nil {
		return nil, err
	}

	project, err := getProject(u.d, u.Config)
	if err != nil {
		return nil, err
	}
	var obj map[string]interface{}

	userAgent, err := generateUserAgentString(u.d, u.Config.UserAgent)
	if err != nil {
		return nil, err
	}

	policy, err := transport_tpg.SendRequest(u.Config, "POST", project, url, userAgent, obj)
	if err != nil {
		return nil, errwrap.Wrapf(fmt.Sprintf("Error retrieving IAM policy for %s: {{err}}", u.DescribeResource()), err)
	}

	out := &cloudresourcemanager.Policy{}
	err = Convert(policy, out)
	if err != nil {
		return nil, errwrap.Wrapf("Cannot convert a policy to a resource manager policy: {{err}}", err)
	}

	return out, nil
}

func (u *VertexAIFeaturestoreIamUpdater) SetResourceIamPolicy(policy *cloudresourcemanager.Policy) error {
	json, err := ConvertToMap(policy)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	obj["policy"] = json

	url, err := u.qualifyFeaturestoreUrl("setIamPolicy")
	if err != nil {
		return err
	}
	project, err := getProject(u.d, u.Config)
	if err != nil {
		return err
	}

	userAgent, err := generateUserAgentString(u.d, u.Config.UserAgent)
	if err != nil {
		return err
	}

	_, err = transport_tpg.SendRequestWithTimeout(u.Config, "POST", project, url, userAgent, obj, u.d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return errwrap.Wrapf(fmt.Sprintf("Error setting IAM policy for %s: {{err}}", u.DescribeResource()), err)
	}

	return nil
}

func (u *VertexAIFeaturestoreIamUpdater) qualifyFeaturestoreUrl(methodIdentifier string) (string, error) {
	urlTemplate := fmt.Sprintf("{{VertexAIBasePath}}%s:%s", fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", u.project, u.region, u.featurestore), methodIdentifier)
	url, err := ReplaceVars(u.d, u.Config, urlTemplate)
	if err != nil {
		return "", err
	}
	return url, nil
}

func (u *VertexAIFeaturestoreIamUpdater) GetResourceId() string {
	return fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", u.project, u.region, u.featurestore)
}

func (u *VertexAIFeaturestoreIamUpdater) GetMutexKey() string {
	return fmt.Sprintf("iam-vertexai-featurestore-%s", u.GetResourceId())
}

func (u *VertexAIFeaturestoreIamUpdater) DescribeResource() string {
	return fmt.Sprintf("vertexai featurestore %q", u.GetResourceId())
}
