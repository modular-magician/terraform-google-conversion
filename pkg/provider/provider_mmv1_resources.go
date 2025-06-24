package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/pkg/tfplan2cai/converters/services/compute"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/pkg/tfplan2cai/converters/services/resourcemanager"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/pkg/tpgiamresource"
)

var handwrittenTfplan2caiResources = map[string]*schema.Resource{
	// ####### START handwritten resources ###########
	"google_compute_instance":             compute.ResourceComputeInstance(),
	"google_project":                      resourcemanager.ResourceGoogleProject(),
	"google_compute_instance_iam_binding": tpgiamresource.ResourceIamBinding(compute.ComputeInstanceIamSchema, compute.ComputeInstanceIamUpdaterProducer, compute.ComputeInstanceIdParseFunc),
	"google_compute_instance_iam_member":  tpgiamresource.ResourceIamMember(compute.ComputeInstanceIamSchema, compute.ComputeInstanceIamUpdaterProducer, compute.ComputeInstanceIdParseFunc),
	"google_compute_instance_iam_policy":  tpgiamresource.ResourceIamPolicy(compute.ComputeInstanceIamSchema, compute.ComputeInstanceIamUpdaterProducer, compute.ComputeInstanceIdParseFunc),
	// ####### END handwritten resources ###########
}

// Generated resources: 2
var generatedResources = map[string]*schema.Resource{
	"google_compute_address":    compute.ResourceComputeAddress(),
	"google_compute_autoscaler": compute.ResourceComputeAutoscaler(),
}
