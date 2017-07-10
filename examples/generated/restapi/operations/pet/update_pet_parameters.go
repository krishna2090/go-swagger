package pet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/krishna2090/go-swagger/examples/generated/models"
)

// NewUpdatePetParams creates a new UpdatePetParams object
// with the default values initialized.
func NewUpdatePetParams() UpdatePetParams {
	var ()
	return UpdatePetParams{}
}

// UpdatePetParams contains all the bound params for the update pet operation
// typically these are obtained from a http.Request
//
// swagger:parameters updatePet
type UpdatePetParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*Pet object that needs to be added to the store
	  In: body
	*/
	Body *models.Pet
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *UpdatePetParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.Pet
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			res = append(res, errors.NewParseError("body", "body", "", err))
		} else {
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Body = &body
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
