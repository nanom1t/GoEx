// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewUserConfirmWithdrawalParams creates a new UserConfirmWithdrawalParams object
// with the default values initialized.
func NewUserConfirmWithdrawalParams() *UserConfirmWithdrawalParams {
	var ()
	return &UserConfirmWithdrawalParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUserConfirmWithdrawalParamsWithTimeout creates a new UserConfirmWithdrawalParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUserConfirmWithdrawalParamsWithTimeout(timeout time.Duration) *UserConfirmWithdrawalParams {
	var ()
	return &UserConfirmWithdrawalParams{

		timeout: timeout,
	}
}

// NewUserConfirmWithdrawalParamsWithContext creates a new UserConfirmWithdrawalParams object
// with the default values initialized, and the ability to set a context for a request
func NewUserConfirmWithdrawalParamsWithContext(ctx context.Context) *UserConfirmWithdrawalParams {
	var ()
	return &UserConfirmWithdrawalParams{

		Context: ctx,
	}
}

// NewUserConfirmWithdrawalParamsWithHTTPClient creates a new UserConfirmWithdrawalParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUserConfirmWithdrawalParamsWithHTTPClient(client *http.Client) *UserConfirmWithdrawalParams {
	var ()
	return &UserConfirmWithdrawalParams{
		HTTPClient: client,
	}
}

/*UserConfirmWithdrawalParams contains all the parameters to send to the API endpoint
for the user confirm withdrawal operation typically these are written to a http.Request
*/
type UserConfirmWithdrawalParams struct {

	/*Token*/
	Token string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the user confirm withdrawal params
func (o *UserConfirmWithdrawalParams) WithTimeout(timeout time.Duration) *UserConfirmWithdrawalParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user confirm withdrawal params
func (o *UserConfirmWithdrawalParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user confirm withdrawal params
func (o *UserConfirmWithdrawalParams) WithContext(ctx context.Context) *UserConfirmWithdrawalParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user confirm withdrawal params
func (o *UserConfirmWithdrawalParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user confirm withdrawal params
func (o *UserConfirmWithdrawalParams) WithHTTPClient(client *http.Client) *UserConfirmWithdrawalParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user confirm withdrawal params
func (o *UserConfirmWithdrawalParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithToken adds the token to the user confirm withdrawal params
func (o *UserConfirmWithdrawalParams) WithToken(token string) *UserConfirmWithdrawalParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the user confirm withdrawal params
func (o *UserConfirmWithdrawalParams) SetToken(token string) {
	o.Token = token
}

// WriteToRequest writes these params to a swagger request
func (o *UserConfirmWithdrawalParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// form param token
	frToken := o.Token
	fToken := frToken
	if fToken != "" {
		if err := r.SetFormParam("token", fToken); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}