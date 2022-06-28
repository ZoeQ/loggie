/*
Copyright 2022 Loggie Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package condition

import (
	"github.com/loggie-io/loggie/pkg/core/api"
	"github.com/loggie-io/loggie/pkg/util/eventops"
	"github.com/pkg/errors"
)

const (
	IsNullName     = "isNull"
	IsNullUsageMsg = "usage: isNull(key)"
)

// IsNull check if the fields value is null or not exist in the event
type IsNull struct {
	field string
}

func init() {
	RegisterCondition(IsNullName, func(args []string) (Condition, error) {
		return NewIsNull(args)
	})
}

func NewIsNull(args []string) (*IsNull, error) {
	if len(args) != 1 {
		return nil, errors.Errorf("invalid args, %s", IsNullUsageMsg)
	}

	return &IsNull{
		field: args[0],
	}, nil
}

func (et *IsNull) Check(e api.Event) bool {
	fieldVal := eventops.Get(e, et.field)
	if fieldVal == nil {
		return true
	}

	return false
}
