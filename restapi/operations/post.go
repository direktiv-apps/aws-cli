// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PostHandlerFunc turns a function with the right signature into a post handler
type PostHandlerFunc func(PostParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostHandlerFunc) Handle(params PostParams) middleware.Responder {
	return fn(params)
}

// PostHandler interface for that can handle valid post params
type PostHandler interface {
	Handle(PostParams) middleware.Responder
}

// NewPost creates a new http.Handler for the post operation
func NewPost(ctx *middleware.Context, handler PostHandler) *Post {
	return &Post{Context: ctx, Handler: handler}
}

/* Post swagger:route POST / post

Post post API

*/
type Post struct {
	Context *middleware.Context
	Handler PostHandler
}

func (o *Post) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// PostBody post body
//
// swagger:model PostBody
type PostBody struct {

	// AWS access key.
	// Example: ABCABCABCDABCABCABCD
	// Required: true
	AccessKey *string `json:"access-key"`

	// Array of AWS cli commands. Does NOT include 'aws'.
	// Example: ["ecr get-login-password","ec2 describe-instances"]
	Commands []string `json:"commands"`

	// If set to true all commands are getting executed and errors ignored.
	// Example: true
	Continue *bool `json:"continue,omitempty"`

	// Region the commands should be executed in.
	// Example: eu-central-1
	Region *string `json:"region,omitempty"`

	// AWS secret key.
	// Example: Abcd45sa01234+ThIsIsSuPeRsEcReT
	// Required: true
	SecretKey *string `json:"secret-key"`
}

// Validate validates this post body
func (o *PostBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAccessKey(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSecretKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostBody) validateAccessKey(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"access-key", "body", o.AccessKey); err != nil {
		return err
	}

	return nil
}

func (o *PostBody) validateSecretKey(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"secret-key", "body", o.SecretKey); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this post body based on context it is used
func (o *PostBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostBody) UnmarshalBinary(b []byte) error {
	var res PostBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostOKBody post o k body
//
// swagger:model PostOKBody
type PostOKBody struct {

	// output
	Output []*PostOKBodyOutputItems0 `json:"output"`
}

// Validate validates this post o k body
func (o *PostOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateOutput(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostOKBody) validateOutput(formats strfmt.Registry) error {
	if swag.IsZero(o.Output) { // not required
		return nil
	}

	for i := 0; i < len(o.Output); i++ {
		if swag.IsZero(o.Output[i]) { // not required
			continue
		}

		if o.Output[i] != nil {
			if err := o.Output[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("postOK" + "." + "output" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("postOK" + "." + "output" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this post o k body based on the context it is used
func (o *PostOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateOutput(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostOKBody) contextValidateOutput(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Output); i++ {

		if o.Output[i] != nil {
			if err := o.Output[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("postOK" + "." + "output" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("postOK" + "." + "output" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostOKBody) UnmarshalBinary(b []byte) error {
	var res PostOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostOKBodyOutputItems0 post o k body output items0
//
// swagger:model PostOKBodyOutputItems0
type PostOKBodyOutputItems0 struct {

	// result
	// Required: true
	Result interface{} `json:"result"`

	// success
	// Required: true
	Success *bool `json:"success"`
}

// Validate validates this post o k body output items0
func (o *PostOKBodyOutputItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSuccess(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostOKBodyOutputItems0) validateResult(formats strfmt.Registry) error {

	if o.Result == nil {
		return errors.Required("result", "body", nil)
	}

	return nil
}

func (o *PostOKBodyOutputItems0) validateSuccess(formats strfmt.Registry) error {

	if err := validate.Required("success", "body", o.Success); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this post o k body output items0 based on context it is used
func (o *PostOKBodyOutputItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostOKBodyOutputItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostOKBodyOutputItems0) UnmarshalBinary(b []byte) error {
	var res PostOKBodyOutputItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
