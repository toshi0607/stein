package lint

import (
	"fmt"

	"github.com/hashicorp/hcl2/hcl"
)

// Policy represents the rule set against config files of arguments
type Policy struct {
	Config  *Config  `hcl:"config,block"`
	Rules   Rules    `hcl:"rule,block"`
	Outputs []Output `hcl:"output,block"`

	Debugs []Debug `hcl:"debug,block"`

	Remain hcl.Body `hcl:",remain"`
}

// Config represents the configuration of the linter itself
type Config struct {
	Report ReportConfig `hcl:"report,block"`
}

// ReportConfig represents the configuration of the report itself
type ReportConfig struct {
	Format string `hcl:"format,optional"`
	Style  string `hcl:"style,optional"`
	Color  bool   `hcl:"color,optional"`
}

// Rule represents the linting rule
type Rule struct {
	Name string `hcl:"name,label"`

	Description  string        `hcl:"description"`
	Dependencies []string      `hcl:"depends_on,optional"`
	Precondition *Precondition `hcl:"precondition,block"`
	Conditions   []bool        `hcl:"conditions"`
	Report       Report        `hcl:"report,block"`

	Debugs []string `hcl:"debug,optional"`
}

// Rules is the collenction of Rule object
type Rules []Rule

// Report represents the rule of reporting style
type Report struct {
	// Level takes ERROR or WARN
	// In case of ERROR, the report message of the failed rule is shown and the linter returns false
	// In case of WARN, the report message of the failed rule is shown and the linter returns true
	Level string `hcl:"level"`

	// Message is shown when the rule is failed
	Message string `hcl:"message"`
}

// Precondition represents a condition that determines whether the rule should be executed
type Precondition struct {
	Cases []bool `hcl:"cases"`
}

// Output is WIP
type Output struct {
	Name  string         `hcl:"name,label"`
	Value *hcl.Attribute `hcl:"value"`
}

// Debug is WIP
type Debug struct {
	Name  string         `hcl:"name,label"`
	Value *hcl.Attribute `hcl:"value"`
}

// Validate validates policy syntax
func (p *Policy) Validate() error {
	validations := []struct {
		rule    bool
		message string
	}{
		{
			// inline is the secret format for now
			p.Config.Report.Style == "console" || p.Config.Report.Style == "inline",
			fmt.Sprintf("%s: console is only acceptable for report style", p.Config.Report.Style),
		},
	}

	for _, validation := range validations {
		if !validation.rule {
			return fmt.Errorf(validation.message)
		}
	}

	return nil
}
