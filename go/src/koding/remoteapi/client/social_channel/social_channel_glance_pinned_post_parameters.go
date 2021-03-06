package social_channel

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

	"koding/remoteapi/models"
)

// NewSocialChannelGlancePinnedPostParams creates a new SocialChannelGlancePinnedPostParams object
// with the default values initialized.
func NewSocialChannelGlancePinnedPostParams() *SocialChannelGlancePinnedPostParams {
	var ()
	return &SocialChannelGlancePinnedPostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSocialChannelGlancePinnedPostParamsWithTimeout creates a new SocialChannelGlancePinnedPostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSocialChannelGlancePinnedPostParamsWithTimeout(timeout time.Duration) *SocialChannelGlancePinnedPostParams {
	var ()
	return &SocialChannelGlancePinnedPostParams{

		timeout: timeout,
	}
}

// NewSocialChannelGlancePinnedPostParamsWithContext creates a new SocialChannelGlancePinnedPostParams object
// with the default values initialized, and the ability to set a context for a request
func NewSocialChannelGlancePinnedPostParamsWithContext(ctx context.Context) *SocialChannelGlancePinnedPostParams {
	var ()
	return &SocialChannelGlancePinnedPostParams{

		Context: ctx,
	}
}

/*SocialChannelGlancePinnedPostParams contains all the parameters to send to the API endpoint
for the social channel glance pinned post operation typically these are written to a http.Request
*/
type SocialChannelGlancePinnedPostParams struct {

	/*Body
	  body of the request

	*/
	Body models.DefaultSelector

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the social channel glance pinned post params
func (o *SocialChannelGlancePinnedPostParams) WithTimeout(timeout time.Duration) *SocialChannelGlancePinnedPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the social channel glance pinned post params
func (o *SocialChannelGlancePinnedPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the social channel glance pinned post params
func (o *SocialChannelGlancePinnedPostParams) WithContext(ctx context.Context) *SocialChannelGlancePinnedPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the social channel glance pinned post params
func (o *SocialChannelGlancePinnedPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the social channel glance pinned post params
func (o *SocialChannelGlancePinnedPostParams) WithBody(body models.DefaultSelector) *SocialChannelGlancePinnedPostParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the social channel glance pinned post params
func (o *SocialChannelGlancePinnedPostParams) SetBody(body models.DefaultSelector) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *SocialChannelGlancePinnedPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
