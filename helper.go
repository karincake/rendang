package risoles

import (
	"net/http"

	d "github.com/karincake/dodol"
	lg "github.com/karincake/lepet"
)

// Process error required for string
func requiredString(w http.ResponseWriter, fieldName, input string) bool {
	if input == "" {
		WriteJSON(w, http.StatusBadRequest, d.II{"errors": d.FieldErrors{
			fieldName: d.FieldError{
				Code:    "val-required",
				Message: lg.I.Msg("val-required"),
			},
		}}, nil)
		return false
	}
	return true
}
