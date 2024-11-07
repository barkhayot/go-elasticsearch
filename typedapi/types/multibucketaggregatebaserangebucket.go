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

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

// MultiBucketAggregateBaseRangeBucket type.
//
// https://github.com/elastic/elasticsearch-specification/blob/48e2d9de9de2911b8cb1cf715e4bc0a2b1f4b827/specification/_types/aggregations/Aggregate.ts#L357-L359
type MultiBucketAggregateBaseRangeBucket struct {
	Buckets BucketsRangeBucket `json:"buckets"`
	Meta    Metadata           `json:"meta,omitempty"`
}

func (s *MultiBucketAggregateBaseRangeBucket) UnmarshalJSON(data []byte) error {

	dec := json.NewDecoder(bytes.NewReader(data))

	for {
		t, err := dec.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}

		switch t {

		case "buckets":

			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			source := bytes.NewReader(rawMsg)
			localDec := json.NewDecoder(source)
			switch rawMsg[0] {
			case '{':
				o := make(map[string]RangeBucket, 0)
				if err := localDec.Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "Buckets", err)
				}
				s.Buckets = o
			case '[':
				o := []RangeBucket{}
				if err := localDec.Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "Buckets", err)
				}
				s.Buckets = o
			}

		case "meta":
			if err := dec.Decode(&s.Meta); err != nil {
				return fmt.Errorf("%s | %w", "Meta", err)
			}

		}
	}
	return nil
}

// NewMultiBucketAggregateBaseRangeBucket returns a MultiBucketAggregateBaseRangeBucket.
func NewMultiBucketAggregateBaseRangeBucket() *MultiBucketAggregateBaseRangeBucket {
	r := &MultiBucketAggregateBaseRangeBucket{}

	return r
}
