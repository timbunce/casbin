// Copyright 2017 The casbin Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package casbin

// addPolicy adds a rule to the current policy.
func (e *Enforcer) addPolicy(sec string, ptype string, rule []string) bool {
	ruleAdded := e.model.AddPolicy(sec, ptype, rule)

	if ruleAdded {
		if e.adapter != nil && e.autoSave {
			err := e.adapter.AddPolicy(sec, ptype, rule)
			if err != nil && err.Error() != "not implemented" {
				panic(err)
			}
		}
	}

	return ruleAdded
}

// removePolicy removes a rule from the current policy.
func (e *Enforcer) removePolicy(sec string, ptype string, rule []string) bool {
	ruleRemoved := e.model.RemovePolicy(sec, ptype, rule)

	if ruleRemoved {
		if e.adapter != nil && e.autoSave {
			err := e.adapter.RemovePolicy(sec, ptype, rule)
			if err != nil && err.Error() != "not implemented" {
				panic(err)
			}
		}
	}

	return ruleRemoved
}

// removeFilteredPolicy removes rules based on field filters from the current policy.
func (e *Enforcer) removeFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) bool {
	ruleRemoved := e.model.RemoveFilteredPolicy(sec, ptype, fieldIndex, fieldValues...)

	if ruleRemoved {
		if e.adapter != nil && e.autoSave {
			err := e.adapter.RemoveFilteredPolicy(sec, ptype, fieldIndex, fieldValues...)
			if err != nil && err.Error() != "not implemented" {
				panic(err)
			}
		}
	}

	return ruleRemoved
}
