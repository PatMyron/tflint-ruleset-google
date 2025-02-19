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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// GoogleNotebooksInstanceInvalidBootDiskTypeRule checks the pattern is valid
type GoogleNotebooksInstanceInvalidBootDiskTypeRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
}

// NewGoogleNotebooksInstanceInvalidBootDiskTypeRule returns new rule with default attributes
func NewGoogleNotebooksInstanceInvalidBootDiskTypeRule() *GoogleNotebooksInstanceInvalidBootDiskTypeRule {
	return &GoogleNotebooksInstanceInvalidBootDiskTypeRule{
		resourceType:  "google_notebooks_instance",
		attributeName: "boot_disk_type",
	}
}

// Name returns the rule name
func (r *GoogleNotebooksInstanceInvalidBootDiskTypeRule) Name() string {
	return "google_notebooks_instance_invalid_boot_disk_type"
}

// Enabled returns whether the rule is enabled by default
func (r *GoogleNotebooksInstanceInvalidBootDiskTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *GoogleNotebooksInstanceInvalidBootDiskTypeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *GoogleNotebooksInstanceInvalidBootDiskTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *GoogleNotebooksInstanceInvalidBootDiskTypeRule) Check(runner tflint.Runner) error {
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

		validateFunc := validation.StringInSlice([]string{"DISK_TYPE_UNSPECIFIED", "PD_STANDARD", "PD_SSD", "PD_BALANCED", ""}, false)

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
