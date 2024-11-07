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

// TimingStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/48e2d9de9de2911b8cb1cf715e4bc0a2b1f4b827/specification/ml/_types/DataframeAnalytics.ts#L562-L567
type TimingStats struct {
	// ElapsedTime Runtime of the analysis in milliseconds.
	ElapsedTime int64 `json:"elapsed_time"`
	// IterationTime Runtime of the latest iteration of the analysis in milliseconds.
	IterationTime *int64 `json:"iteration_time,omitempty"`
}

func (s *TimingStats) UnmarshalJSON(data []byte) error {

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

		case "elapsed_time":
			if err := dec.Decode(&s.ElapsedTime); err != nil {
				return fmt.Errorf("%s | %w", "ElapsedTime", err)
			}

		case "iteration_time":
			if err := dec.Decode(&s.IterationTime); err != nil {
				return fmt.Errorf("%s | %w", "IterationTime", err)
			}

		}
	}
	return nil
}

// NewTimingStats returns a TimingStats.
func NewTimingStats() *TimingStats {
	r := &TimingStats{}

	return r
}
