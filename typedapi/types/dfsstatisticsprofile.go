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
	"strconv"
)

// DfsStatisticsProfile type.
//
// https://github.com/elastic/elasticsearch-specification/blob/48e2d9de9de2911b8cb1cf715e4bc0a2b1f4b827/specification/_global/search/_types/profile.ts#L159-L167
type DfsStatisticsProfile struct {
	Breakdown   DfsStatisticsBreakdown     `json:"breakdown"`
	Children    []DfsStatisticsProfile     `json:"children,omitempty"`
	Debug       map[string]json.RawMessage `json:"debug,omitempty"`
	Description string                     `json:"description"`
	Time        Duration                   `json:"time,omitempty"`
	TimeInNanos int64                      `json:"time_in_nanos"`
	Type        string                     `json:"type"`
}

func (s *DfsStatisticsProfile) UnmarshalJSON(data []byte) error {

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

		case "breakdown":
			if err := dec.Decode(&s.Breakdown); err != nil {
				return fmt.Errorf("%s | %w", "Breakdown", err)
			}

		case "children":
			if err := dec.Decode(&s.Children); err != nil {
				return fmt.Errorf("%s | %w", "Children", err)
			}

		case "debug":
			if s.Debug == nil {
				s.Debug = make(map[string]json.RawMessage, 0)
			}
			if err := dec.Decode(&s.Debug); err != nil {
				return fmt.Errorf("%s | %w", "Debug", err)
			}

		case "description":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Description", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Description = o

		case "time":
			if err := dec.Decode(&s.Time); err != nil {
				return fmt.Errorf("%s | %w", "Time", err)
			}

		case "time_in_nanos":
			if err := dec.Decode(&s.TimeInNanos); err != nil {
				return fmt.Errorf("%s | %w", "TimeInNanos", err)
			}

		case "type":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Type", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Type = o

		}
	}
	return nil
}

// NewDfsStatisticsProfile returns a DfsStatisticsProfile.
func NewDfsStatisticsProfile() *DfsStatisticsProfile {
	r := &DfsStatisticsProfile{
		Debug: make(map[string]json.RawMessage, 0),
	}

	return r
}
