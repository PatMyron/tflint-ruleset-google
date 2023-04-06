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

package magicmodules

import (
	"fmt"
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// GoogleComputeRegionBackendServiceInvalidSessionAffinityRule checks the pattern is valid
type GoogleComputeRegionBackendServiceInvalidSessionAffinityRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
}

// NewGoogleComputeRegionBackendServiceInvalidSessionAffinityRule returns new rule with default attributes
func NewGoogleComputeRegionBackendServiceInvalidSessionAffinityRule() *GoogleComputeRegionBackendServiceInvalidSessionAffinityRule {
	return &GoogleComputeRegionBackendServiceInvalidSessionAffinityRule{
		resourceType:  "google_compute_region_backend_service",
		attributeName: "session_affinity",
	}
}

// Name returns the rule name
func (r *GoogleComputeRegionBackendServiceInvalidSessionAffinityRule) Name() string {
	return "google_compute_region_backend_service_invalid_session_affinity"
}

// Enabled returns whether the rule is enabled by default
func (r *GoogleComputeRegionBackendServiceInvalidSessionAffinityRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *GoogleComputeRegionBackendServiceInvalidSessionAffinityRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *GoogleComputeRegionBackendServiceInvalidSessionAffinityRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *GoogleComputeRegionBackendServiceInvalidSessionAffinityRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{{Name: r.attributeName}},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		attribute, exists := resource.Body.Attributes[r.attributeName]
		if !exists {
			continue
		}

		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		validateFunc := validation.StringInSlice([]string{"NONE", "CLIENT_IP", "CLIENT_IP_PORT_PROTO", "CLIENT_IP_PROTO", "GENERATED_COOKIE", "HEADER_FIELD", "HTTP_COOKIE", "CLIENT_IP_NO_DESTINATION", ""}, false)

		err = runner.EnsureNoError(err, func() error {
			_, errors := validateFunc(val, r.attributeName)
			for _, err := range errors {
				runner.EmitIssue(r, err.Error(), attribute.Expr.Range())
			}
			return nil
		})
		if err != nil {
			return err
		}
	}

	return nil
}
