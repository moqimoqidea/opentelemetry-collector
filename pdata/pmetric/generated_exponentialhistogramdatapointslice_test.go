// Copyright The OpenTelemetry Authors
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

// Code generated by "pdata/internal/cmd/pdatagen/main.go". DO NOT EDIT.
// To regenerate this file run "make genpdata".

package pmetric

import (
	"testing"

	"github.com/stretchr/testify/assert"

	otlpmetrics "go.opentelemetry.io/collector/pdata/internal/data/protogen/metrics/v1"
)

func TestExponentialHistogramDataPointSlice(t *testing.T) {
	es := NewExponentialHistogramDataPointSlice()
	assert.Equal(t, 0, es.Len())
	es = newExponentialHistogramDataPointSlice(&[]*otlpmetrics.ExponentialHistogramDataPoint{})
	assert.Equal(t, 0, es.Len())

	es.EnsureCapacity(7)
	emptyVal := newExponentialHistogramDataPoint(&otlpmetrics.ExponentialHistogramDataPoint{})
	testVal := generateTestExponentialHistogramDataPoint()
	assert.Equal(t, 7, cap(*es.orig))
	for i := 0; i < es.Len(); i++ {
		el := es.AppendEmpty()
		assert.Equal(t, emptyVal, el)
		fillTestExponentialHistogramDataPoint(el)
		assert.Equal(t, testVal, el)
	}
}

func TestExponentialHistogramDataPointSlice_CopyTo(t *testing.T) {
	dest := NewExponentialHistogramDataPointSlice()
	// Test CopyTo to empty
	NewExponentialHistogramDataPointSlice().CopyTo(dest)
	assert.Equal(t, NewExponentialHistogramDataPointSlice(), dest)

	// Test CopyTo larger slice
	generateTestExponentialHistogramDataPointSlice().CopyTo(dest)
	assert.Equal(t, generateTestExponentialHistogramDataPointSlice(), dest)

	// Test CopyTo same size slice
	generateTestExponentialHistogramDataPointSlice().CopyTo(dest)
	assert.Equal(t, generateTestExponentialHistogramDataPointSlice(), dest)
}

func TestExponentialHistogramDataPointSlice_EnsureCapacity(t *testing.T) {
	es := generateTestExponentialHistogramDataPointSlice()

	// Test ensure smaller capacity.
	const ensureSmallLen = 4
	es.EnsureCapacity(ensureSmallLen)
	assert.Less(t, ensureSmallLen, es.Len())
	assert.Equal(t, es.Len(), cap(*es.orig))
	assert.Equal(t, generateTestExponentialHistogramDataPointSlice(), es)

	// Test ensure larger capacity
	const ensureLargeLen = 9
	es.EnsureCapacity(ensureLargeLen)
	assert.Less(t, generateTestExponentialHistogramDataPointSlice().Len(), ensureLargeLen)
	assert.Equal(t, ensureLargeLen, cap(*es.orig))
	assert.Equal(t, generateTestExponentialHistogramDataPointSlice(), es)
}

func TestExponentialHistogramDataPointSlice_MoveAndAppendTo(t *testing.T) {
	// Test MoveAndAppendTo to empty
	expectedSlice := generateTestExponentialHistogramDataPointSlice()
	dest := NewExponentialHistogramDataPointSlice()
	src := generateTestExponentialHistogramDataPointSlice()
	src.MoveAndAppendTo(dest)
	assert.Equal(t, generateTestExponentialHistogramDataPointSlice(), dest)
	assert.Equal(t, 0, src.Len())
	assert.Equal(t, expectedSlice.Len(), dest.Len())

	// Test MoveAndAppendTo empty slice
	src.MoveAndAppendTo(dest)
	assert.Equal(t, generateTestExponentialHistogramDataPointSlice(), dest)
	assert.Equal(t, 0, src.Len())
	assert.Equal(t, expectedSlice.Len(), dest.Len())

	// Test MoveAndAppendTo not empty slice
	generateTestExponentialHistogramDataPointSlice().MoveAndAppendTo(dest)
	assert.Equal(t, 2*expectedSlice.Len(), dest.Len())
	for i := 0; i < expectedSlice.Len(); i++ {
		assert.Equal(t, expectedSlice.At(i), dest.At(i))
		assert.Equal(t, expectedSlice.At(i), dest.At(i+expectedSlice.Len()))
	}
}

func TestExponentialHistogramDataPointSlice_RemoveIf(t *testing.T) {
	// Test RemoveIf on empty slice
	emptySlice := NewExponentialHistogramDataPointSlice()
	emptySlice.RemoveIf(func(el ExponentialHistogramDataPoint) bool {
		t.Fail()
		return false
	})

	// Test RemoveIf
	filtered := generateTestExponentialHistogramDataPointSlice()
	pos := 0
	filtered.RemoveIf(func(el ExponentialHistogramDataPoint) bool {
		pos++
		return pos%3 == 0
	})
	assert.Equal(t, 5, filtered.Len())
}

func generateTestExponentialHistogramDataPointSlice() ExponentialHistogramDataPointSlice {
	tv := NewExponentialHistogramDataPointSlice()
	fillTestExponentialHistogramDataPointSlice(tv)
	return tv
}

func fillTestExponentialHistogramDataPointSlice(tv ExponentialHistogramDataPointSlice) {
	*tv.orig = make([]*otlpmetrics.ExponentialHistogramDataPoint, 7)
	for i := 0; i < 7; i++ {
		(*tv.orig)[i] = &otlpmetrics.ExponentialHistogramDataPoint{}
		fillTestExponentialHistogramDataPoint(newExponentialHistogramDataPoint((*tv.orig)[i]))
	}
}