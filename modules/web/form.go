package web

import (
	"strings"

	"github.com/astaxie/beego/validation"
	"github.com/fatih/structs"

	"gogs.ballantine.tech/gballan1/gowis/modules/base"
)

// Form - base form class
type Form struct {
	Valid  validation.Validation
	Fields []*structs.Field
}

// GetErrors - returns the errors in a nice map
func (f *Form) GetErrors() map[string][]string {
	errors := make(map[string][]string)

	// loop through field names
	for _, field := range f.Fields {
		if field.Name() != "Form" {
			// get lowercase version of field name
			name := strings.ToLower(field.Name())

			// initalize the array
			errors[name] = make([]string, 0)

			// loop through any errors
			for _, err := range f.Valid.Errors {
				if strings.HasPrefix(err.Key, name) {
					errors[name] = base.Append(errors[name], err.Message)
				}
			}
		}
	}

	return errors
}

// HasErrors - returns true if the form has validation errors
func (f *Form) HasErrors() bool {
	return f.Valid.HasErrors()
}
