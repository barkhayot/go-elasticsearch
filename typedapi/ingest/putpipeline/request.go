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
// https://github.com/elastic/elasticsearch-specification/tree/ec3159eb31c62611202a4fb157ea88fa6ff78e1a


package putpipeline

import (
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Request holds the request body struct for the package putpipeline
//
// https://github.com/elastic/elasticsearch-specification/blob/ec3159eb31c62611202a4fb157ea88fa6ff78e1a/specification/ingest/put_pipeline/PutPipelineRequest.ts#L25-L74
type Request struct {

	// Description Description of the ingest pipeline.
	Description *string `json:"description,omitempty"`
	// Meta_ Optional metadata about the ingest pipeline. May have any contents. This map
	// is not automatically generated by Elasticsearch.
	Meta_ map[string]interface{} `json:"_meta,omitempty"`
	// OnFailure Processors to run immediately after a processor failure. Each processor
	// supports a processor-level `on_failure` value. If a processor without an
	// `on_failure` value fails, Elasticsearch uses this pipeline-level parameter as
	// a fallback. The processors in this parameter run sequentially in the order
	// specified. Elasticsearch will not attempt to run the pipeline's remaining
	// processors.
	OnFailure []types.ProcessorContainer `json:"on_failure,omitempty"`
	// Processors Processors used to perform transformations on documents before indexing.
	// Processors run sequentially in the order specified.
	Processors []types.ProcessorContainer `json:"processors,omitempty"`
	// Version Version number used by external systems to track ingest pipelines. This
	// parameter is intended for external systems only. Elasticsearch does not use
	// or validate pipeline version numbers.
	Version *int64 `json:"version,omitempty"`
}

// NewRequest returns a Request
func NewRequest() *Request {
	r := &Request{}
	return r
}

// FromJSON allows to load an arbitrary json into the request structure
func (rb *Request) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Putpipeline request: %w", err)
	}

	return &req, nil
}
