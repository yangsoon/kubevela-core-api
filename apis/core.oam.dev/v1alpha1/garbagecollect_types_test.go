/*
Copyright 2021 The KubeVela Authors.

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

package v1alpha1

import (
	"testing"

	"github.com/stretchr/testify/require"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

	"github.com/oam-dev/kubevela-core-api/pkg/oam"
)

func TestGarbageCollectPolicySpec_FindStrategy(t *testing.T) {
	testCases := map[string]struct {
		rules          []GarbageCollectPolicyRule
		input          *unstructured.Unstructured
		notFound       bool
		expectStrategy GarbageCollectStrategy
	}{
		"trait rule match": {
			rules: []GarbageCollectPolicyRule{{
				Selector: GarbageCollectPolicyRuleSelector{TraitTypes: []string{"a"}},
				Strategy: GarbageCollectStrategyNever,
			}},
			input: &unstructured.Unstructured{Object: map[string]interface{}{
				"metadata": map[string]interface{}{
					"labels": map[string]interface{}{oam.TraitTypeLabel: "a"},
				},
			}},
			expectStrategy: GarbageCollectStrategyNever,
		},
		"trait rule mismatch": {
			rules: []GarbageCollectPolicyRule{{
				Selector: GarbageCollectPolicyRuleSelector{TraitTypes: []string{"a"}},
				Strategy: GarbageCollectStrategyNever,
			}},
			input:    &unstructured.Unstructured{Object: map[string]interface{}{}},
			notFound: true,
		},
		"trait rule multiple match": {
			rules: []GarbageCollectPolicyRule{{
				Selector: GarbageCollectPolicyRuleSelector{TraitTypes: []string{"a"}},
				Strategy: GarbageCollectStrategyOnAppDelete,
			}, {
				Selector: GarbageCollectPolicyRuleSelector{TraitTypes: []string{"a"}},
				Strategy: GarbageCollectStrategyNever,
			}},
			input: &unstructured.Unstructured{Object: map[string]interface{}{
				"metadata": map[string]interface{}{
					"labels": map[string]interface{}{oam.TraitTypeLabel: "a"},
				},
			}},
			expectStrategy: GarbageCollectStrategyOnAppDelete,
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			r := require.New(t)
			spec := GarbageCollectPolicySpec{Rules: tc.rules}
			strategy := spec.FindStrategy(tc.input)
			if tc.notFound {
				r.Nil(strategy)
			} else {
				r.Equal(tc.expectStrategy, *strategy)
			}
		})
	}
}
