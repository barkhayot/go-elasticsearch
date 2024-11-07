// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification/tree/48e2d9de9de2911b8cb1cf715e4bc0a2b1f4b827

// Update an API key.
//
// Updates attributes of an existing API key.
// Users can only update API keys that they created or that were granted to
// them.
// Use this API to update API keys created by the create API Key or grant API
// Key APIs.
// If you need to apply the same update to many API keys, you can use bulk
// update API Keys to reduce overhead.
// It’s not possible to update expired API keys, or API keys that have been
// invalidated by invalidate API Key.
// This API supports updates to an API key’s access scope and metadata.
// The access scope of an API key is derived from the `role_descriptors` you
// specify in the request, and a snapshot of the owner user’s permissions at the
// time of the request.
// The snapshot of the owner’s permissions is updated automatically on every
// call.
// If you don’t specify `role_descriptors` in the request, a call to this API
// might still change the API key’s access scope.
// This change can occur if the owner user’s permissions have changed since the
// API key was created or last modified.
// To update another user’s API key, use the `run_as` feature to submit a
// request on behalf of another user.
// IMPORTANT: It’s not possible to use an API key as the authentication
// credential for this API.
// To update an API key, the owner user’s credentials are required.
package updateapikey

import (
	gobytes "bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

const (
	idMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type UpdateApiKey struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	id string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewUpdateApiKey type alias for index.
type NewUpdateApiKey func(id string) *UpdateApiKey

// NewUpdateApiKeyFunc returns a new instance of UpdateApiKey with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewUpdateApiKeyFunc(tp elastictransport.Interface) NewUpdateApiKey {
	return func(id string) *UpdateApiKey {
		n := New(tp)

		n._id(id)

		return n
	}
}

// Update an API key.
//
// Updates attributes of an existing API key.
// Users can only update API keys that they created or that were granted to
// them.
// Use this API to update API keys created by the create API Key or grant API
// Key APIs.
// If you need to apply the same update to many API keys, you can use bulk
// update API Keys to reduce overhead.
// It’s not possible to update expired API keys, or API keys that have been
// invalidated by invalidate API Key.
// This API supports updates to an API key’s access scope and metadata.
// The access scope of an API key is derived from the `role_descriptors` you
// specify in the request, and a snapshot of the owner user’s permissions at the
// time of the request.
// The snapshot of the owner’s permissions is updated automatically on every
// call.
// If you don’t specify `role_descriptors` in the request, a call to this API
// might still change the API key’s access scope.
// This change can occur if the owner user’s permissions have changed since the
// API key was created or last modified.
// To update another user’s API key, use the `run_as` feature to submit a
// request on behalf of another user.
// IMPORTANT: It’s not possible to use an API key as the authentication
// credential for this API.
// To update an API key, the owner user’s credentials are required.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-update-api-key.html
func New(tp elastictransport.Interface) *UpdateApiKey {
	r := &UpdateApiKey{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),

		buf: gobytes.NewBuffer(nil),

		req: NewRequest(),
	}

	if instrumented, ok := r.transport.(elastictransport.Instrumented); ok {
		if instrument := instrumented.InstrumentationEnabled(); instrument != nil {
			r.instrument = instrument
		}
	}

	return r
}

// Raw takes a json payload as input which is then passed to the http.Request
// If specified Raw takes precedence on Request method.
func (r *UpdateApiKey) Raw(raw io.Reader) *UpdateApiKey {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *UpdateApiKey) Request(req *Request) *UpdateApiKey {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *UpdateApiKey) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	if len(r.deferred) > 0 {
		for _, f := range r.deferred {
			deferredErr := f(r.req)
			if deferredErr != nil {
				return nil, deferredErr
			}
		}
	}

	if r.raw == nil && r.req != nil {

		data, err := json.Marshal(r.req)

		if err != nil {
			return nil, fmt.Errorf("could not serialise request for UpdateApiKey: %w", err)
		}

		r.buf.Write(data)

	}

	if r.buf.Len() > 0 {
		r.raw = r.buf
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == idMask:
		path.WriteString("/")
		path.WriteString("_security")
		path.WriteString("/")
		path.WriteString("api_key")
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "id", r.id)
		}
		path.WriteString(r.id)

		method = http.MethodPut
	}

	r.path.Path = path.String()
	r.path.RawQuery = r.values.Encode()

	if r.path.Path == "" {
		return nil, ErrBuildPath
	}

	if ctx != nil {
		req, err = http.NewRequestWithContext(ctx, method, r.path.String(), r.raw)
	} else {
		req, err = http.NewRequest(method, r.path.String(), r.raw)
	}

	req.Header = r.headers.Clone()

	if req.Header.Get("Content-Type") == "" {
		if r.raw != nil {
			req.Header.Set("Content-Type", "application/vnd.elasticsearch+json;compatible-with=8")
		}
	}

	if req.Header.Get("Accept") == "" {
		req.Header.Set("Accept", "application/vnd.elasticsearch+json;compatible-with=8")
	}

	if err != nil {
		return req, fmt.Errorf("could not build http.Request: %w", err)
	}

	return req, nil
}

// Perform runs the http.Request through the provided transport and returns an http.Response.
func (r UpdateApiKey) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "security.update_api_key")
			defer instrument.Close(ctx)
		}
	}
	if ctx == nil {
		ctx = providedCtx
	}

	req, err := r.HttpRequest(ctx)
	if err != nil {
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return nil, err
	}

	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.BeforeRequest(req, "security.update_api_key")
		if reader := instrument.RecordRequestBody(ctx, "security.update_api_key", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "security.update_api_key")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the UpdateApiKey query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a updateapikey.Response
func (r UpdateApiKey) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "security.update_api_key")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	response := NewResponse()

	res, err := r.Perform(ctx)
	if err != nil {
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode < 299 {
		err = json.NewDecoder(res.Body).Decode(response)
		if err != nil {
			if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
				instrument.RecordError(ctx, err)
			}
			return nil, err
		}

		return response, nil
	}

	errorResponse := types.NewElasticsearchError()
	err = json.NewDecoder(res.Body).Decode(errorResponse)
	if err != nil {
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return nil, err
	}

	if errorResponse.Status == 0 {
		errorResponse.Status = res.StatusCode
	}

	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.RecordError(ctx, errorResponse)
	}
	return nil, errorResponse
}

// Header set a key, value pair in the UpdateApiKey headers map.
func (r *UpdateApiKey) Header(key, value string) *UpdateApiKey {
	r.headers.Set(key, value)

	return r
}

// Id The ID of the API key to update.
// API Name: id
func (r *UpdateApiKey) _id(id string) *UpdateApiKey {
	r.paramSet |= idMask
	r.id = id

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *UpdateApiKey) ErrorTrace(errortrace bool) *UpdateApiKey {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *UpdateApiKey) FilterPath(filterpaths ...string) *UpdateApiKey {
	tmp := []string{}
	for _, item := range filterpaths {
		tmp = append(tmp, fmt.Sprintf("%v", item))
	}
	r.values.Set("filter_path", strings.Join(tmp, ","))

	return r
}

// Human When set to `true` will return statistics in a format suitable for humans.
// For example `"exists_time": "1h"` for humans and
// `"eixsts_time_in_millis": 3600000` for computers. When disabled the human
// readable values will be omitted. This makes sense for responses being
// consumed
// only by machines.
// API name: human
func (r *UpdateApiKey) Human(human bool) *UpdateApiKey {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *UpdateApiKey) Pretty(pretty bool) *UpdateApiKey {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// Expiration Expiration time for the API key.
// API name: expiration
func (r *UpdateApiKey) Expiration(duration types.Duration) *UpdateApiKey {
	r.req.Expiration = duration

	return r
}

// Metadata Arbitrary metadata that you want to associate with the API key. It supports
// nested data structure. Within the metadata object, keys beginning with _ are
// reserved for system usage.
// API name: metadata
func (r *UpdateApiKey) Metadata(metadata types.Metadata) *UpdateApiKey {
	r.req.Metadata = metadata

	return r
}

// RoleDescriptors An array of role descriptors for this API key. This parameter is optional.
// When it is not specified or is an empty array, then the API key will have a
// point in time snapshot of permissions of the authenticated user. If you
// supply role descriptors then the resultant permissions would be an
// intersection of API keys permissions and authenticated user’s permissions
// thereby limiting the access scope for API keys. The structure of role
// descriptor is the same as the request for create role API. For more details,
// see create or update roles API.
// API name: role_descriptors
func (r *UpdateApiKey) RoleDescriptors(roledescriptors map[string]types.RoleDescriptor) *UpdateApiKey {

	r.req.RoleDescriptors = roledescriptors

	return r
}
