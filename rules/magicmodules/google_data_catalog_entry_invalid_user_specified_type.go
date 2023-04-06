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

// GoogleDataCatalogEntryInvalidUserSpecifiedTypeRule checks the pattern is valid
type GoogleDataCatalogEntryInvalidUserSpecifiedTypeRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
}

// NewGoogleDataCatalogEntryInvalidUserSpecifiedTypeRule returns new rule with default attributes
func NewGoogleDataCatalogEntryInvalidUserSpecifiedTypeRule() *GoogleDataCatalogEntryInvalidUserSpecifiedTypeRule {
	return &GoogleDataCatalogEntryInvalidUserSpecifiedTypeRule{
		resourceType:  "google_data_catalog_entry",
		attributeName: "user_specified_type",
	}
}

// Name returns the rule name
func (r *GoogleDataCatalogEntryInvalidUserSpecifiedTypeRule) Name() string {
	return "google_data_catalog_entry_invalid_user_specified_type"
}

// Enabled returns whether the rule is enabled by default
func (r *GoogleDataCatalogEntryInvalidUserSpecifiedTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *GoogleDataCatalogEntryInvalidUserSpecifiedTypeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *GoogleDataCatalogEntryInvalidUserSpecifiedTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *GoogleDataCatalogEntryInvalidUserSpecifiedTypeRule) Check(runner tflint.Runner) error {
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

		validateFunc := validateRegexp(`^[A-z_][A-z0-9_]{0,63}$`)

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
