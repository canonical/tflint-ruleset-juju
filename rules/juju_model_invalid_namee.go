package rules

import (
	"fmt"

	"github.com/juju/names/v5"
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// JujuModelInvalidNameRule checks whether application name is valid.
type JujuModelInvalidNameRule struct {
	tflint.DefaultRule
}

// NewJujuModelInvalidNameRule returns a new rule.
func NewJujuModelInvalidNameRule() *JujuModelInvalidNameRule {
	return &JujuModelInvalidNameRule{}
}

// Name returns the rule name.
func (r *JujuModelInvalidNameRule) Name() string {
	return "juju_model_invalid_name"
}

// Enabled returns whether the rule is enabled by default.
func (r *JujuModelInvalidNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity.
func (r *JujuModelInvalidNameRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link.
func (r *JujuModelInvalidNameRule) Link() string {
	return ""
}

// Check checks whether model name is valid.
func (r *JujuModelInvalidNameRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("juju_model", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "name"},
		},
	}, nil)
	if err != nil {
		return err
	}

	logger.Debug(fmt.Sprintf("Get %d models", len(resources.Blocks)))

	for _, resource := range resources.Blocks {
		attribute, exists := resource.Body.Attributes["name"]
		if !exists {
			continue
		}

		err := runner.EvaluateExpr(attribute.Expr, func(modelName string) error {
			if !names.IsValidModelName(modelName) {
				return runner.EmitIssue(
					r,
					fmt.Sprintf("%s is not a valid model name", modelName),
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
