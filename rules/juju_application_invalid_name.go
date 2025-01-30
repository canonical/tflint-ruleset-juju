package rules

import (
	"fmt"

	"github.com/juju/names/v5"
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// JujuApplicationInvalidNameRule checks whether application name is valid.
type JujuApplicationInvalidNameRule struct {
	tflint.DefaultRule
}

// NewJujuApplicationInvalidNameRule returns a new rule.
func NewJujuApplicationInvalidNameRule() *JujuApplicationInvalidNameRule {
	return &JujuApplicationInvalidNameRule{}
}

// Name returns the rule name.
func (r *JujuApplicationInvalidNameRule) Name() string {
	return "juju_application_invalid_name"
}

// Enabled returns whether the rule is enabled by default.
func (r *JujuApplicationInvalidNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity.
func (r *JujuApplicationInvalidNameRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link.
func (r *JujuApplicationInvalidNameRule) Link() string {
	return ""
}

// Check checks whether application name is valid.
func (r *JujuApplicationInvalidNameRule) Check(runner tflint.Runner) error {
	// This rule is an example to get a top-level resource attribute.
	resources, err := runner.GetResourceContent("juju_application", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "name"},
		},
	}, nil)
	if err != nil {
		return err
	}

	// Put a log that can be output with `TFLINT_LOG=debug`
	logger.Debug(fmt.Sprintf("Get %d applications", len(resources.Blocks)))

	for _, resource := range resources.Blocks {
		attribute, exists := resource.Body.Attributes["name"]
		if !exists {
			continue
		}

		err := runner.EvaluateExpr(attribute.Expr, func(applicationName string) error {
			err := names.ValidateApplicationName(applicationName)
			if err != nil {
				return runner.EmitIssue(
					r,
					err.Error(),
					attribute.Expr.Range(),
				)
			}
			return nil
		}, nil)
		if err != nil {
			return err
		}
	}

	return nil
}
