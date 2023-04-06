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

// GoogleCloudiotRegistryInvalidLogLevelRule checks the pattern is valid
type GoogleCloudiotRegistryInvalidLogLevelRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
}

// NewGoogleCloudiotRegistryInvalidLogLevelRule returns new rule with default attributes
func NewGoogleCloudiotRegistryInvalidLogLevelRule() *GoogleCloudiotRegistryInvalidLogLevelRule {
	return &GoogleCloudiotRegistryInvalidLogLevelRule{
		resourceType:  "google_cloudiot_registry",
		attributeName: "log_level",
	}
}

// Name returns the rule name
func (r *GoogleCloudiotRegistryInvalidLogLevelRule) Name() string {
	return "google_cloudiot_registry_invalid_log_level"
}

// Enabled returns whether the rule is enabled by default
func (r *GoogleCloudiotRegistryInvalidLogLevelRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *GoogleCloudiotRegistryInvalidLogLevelRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *GoogleCloudiotRegistryInvalidLogLevelRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *GoogleCloudiotRegistryInvalidLogLevelRule) Check(runner tflint.Runner) error {
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

		validateFunc := validation.StringInSlice([]string{"NONE", "ERROR", "INFO", "DEBUG", ""}, false)

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
