package main

// Params are the parameters passed to this extension via custom stage properties
type Params struct {
	Formula              string `json:"formula,omitempty" yaml:"formula,omitempty"`
	Description          string `json:"description,omitempty" yaml:"description,omitempty"`
	Homepage             string `json:"homepage,omitempty" yaml:"homepage,omitempty"`
	BinaryURL            string `json:"binaryURL,omitempty" yaml:"binaryURL,omitempty"`
	Version              string `json:"version,omitempty" yaml:"version,omitempty"`
	TapReposityDirectory string `json:"tapRepoDir,omitempty" yaml:"tapRepoDir,omitempty"`
	FormulaDirectory     string `json:"formulaDir,omitempty" yaml:"formulaDir,omitempty"`
}

// SetDefaults sets some sane defaults
func (p *Params) SetDefaults(buildVersion string) {
	if p.FormulaDirectory == "" {
		p.FormulaDirectory = "Formula"
	}
	if p.Version == "" {
		p.Version = buildVersion
	}
}

// Validate checks if the parameters are valid
func (p *Params) Validate() (valid bool, warnings []string) {

	if p.Formula == "" {
		warnings = append(warnings, "Parameter formula is not set")
	}
	if p.BinaryURL == "" {
		warnings = append(warnings, "Parameter binaryURL is not set")
	}
	if p.Version == "" {
		warnings = append(warnings, "Parameter version is not set")
	}
	if p.TapReposityDirectory == "" {
		warnings = append(warnings, "Parameter tapRepoDir is not set")
	}

	return len(warnings) == 0, warnings
}
