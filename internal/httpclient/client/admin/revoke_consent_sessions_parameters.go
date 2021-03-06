// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewRevokeConsentSessionsParams creates a new RevokeConsentSessionsParams object
// with the default values initialized.
func NewRevokeConsentSessionsParams() *RevokeConsentSessionsParams {
	var ()
	return &RevokeConsentSessionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRevokeConsentSessionsParamsWithTimeout creates a new RevokeConsentSessionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRevokeConsentSessionsParamsWithTimeout(timeout time.Duration) *RevokeConsentSessionsParams {
	var ()
	return &RevokeConsentSessionsParams{

		timeout: timeout,
	}
}

// NewRevokeConsentSessionsParamsWithContext creates a new RevokeConsentSessionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewRevokeConsentSessionsParamsWithContext(ctx context.Context) *RevokeConsentSessionsParams {
	var ()
	return &RevokeConsentSessionsParams{

		Context: ctx,
	}
}

// NewRevokeConsentSessionsParamsWithHTTPClient creates a new RevokeConsentSessionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRevokeConsentSessionsParamsWithHTTPClient(client *http.Client) *RevokeConsentSessionsParams {
	var ()
	return &RevokeConsentSessionsParams{
		HTTPClient: client,
	}
}

/*RevokeConsentSessionsParams contains all the parameters to send to the API endpoint
for the revoke consent sessions operation typically these are written to a http.Request
*/
type RevokeConsentSessionsParams struct {

	/*Client
	  If set, deletes only those consent sessions by the Subject that have been granted to the specified OAuth 2.0 Client ID

	*/
	Client *string
	/*Subject
	  The subject (Subject) who's consent sessions should be deleted.

	*/
	Subject string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the revoke consent sessions params
func (o *RevokeConsentSessionsParams) WithTimeout(timeout time.Duration) *RevokeConsentSessionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the revoke consent sessions params
func (o *RevokeConsentSessionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the revoke consent sessions params
func (o *RevokeConsentSessionsParams) WithContext(ctx context.Context) *RevokeConsentSessionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the revoke consent sessions params
func (o *RevokeConsentSessionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the revoke consent sessions params
func (o *RevokeConsentSessionsParams) WithHTTPClient(client *http.Client) *RevokeConsentSessionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the revoke consent sessions params
func (o *RevokeConsentSessionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClient adds the client to the revoke consent sessions params
func (o *RevokeConsentSessionsParams) WithClient(client *string) *RevokeConsentSessionsParams {
	o.SetClient(client)
	return o
}

// SetClient adds the client to the revoke consent sessions params
func (o *RevokeConsentSessionsParams) SetClient(client *string) {
	o.Client = client
}

// WithSubject adds the subject to the revoke consent sessions params
func (o *RevokeConsentSessionsParams) WithSubject(subject string) *RevokeConsentSessionsParams {
	o.SetSubject(subject)
	return o
}

// SetSubject adds the subject to the revoke consent sessions params
func (o *RevokeConsentSessionsParams) SetSubject(subject string) {
	o.Subject = subject
}

// WriteToRequest writes these params to a swagger request
func (o *RevokeConsentSessionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Client != nil {

		// query param client
		var qrClient string
		if o.Client != nil {
			qrClient = *o.Client
		}
		qClient := qrClient
		if qClient != "" {
			if err := r.SetQueryParam("client", qClient); err != nil {
				return err
			}
		}

	}

	// query param subject
	qrSubject := o.Subject
	qSubject := qrSubject
	if qSubject != "" {
		if err := r.SetQueryParam("subject", qSubject); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
