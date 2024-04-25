package toUpperCase

import (
	"math"
	"testing"
	"time"

	"github.com/go-graphite/carbonapi/expr/interfaces"
	"github.com/go-graphite/carbonapi/expr/metadata"
	"github.com/go-graphite/carbonapi/expr/types"
	"github.com/go-graphite/carbonapi/pkg/parser"
	th "github.com/go-graphite/carbonapi/tests"
)

var (
	md []interfaces.FunctionMetadata = New("")
)

func init() {
	for _, m := range md {
		metadata.RegisterFunction(m.Name, m.F)
	}
}

func TestToUpperCaseFunction(t *testing.T) {
	now32 := time.Now().Unix()

	tests := []th.EvalTestItem{
		{
			"upper(metric.test.foo)",
			map[parser.MetricRequest][]*types.MetricData{
				{Metric: "metric.test.foo", From: 0, Until: 1}: {types.MakeMetricData("metric.test.foo", []float64{1, 2, 0, 7, 8, 20, 30, math.NaN()}, 1, now32)},
			},
			[]*types.MetricData{types.MakeMetricData("METRIC.TEST.FOO",
				[]float64{1, 2, 0, 7, 8, 20, 30, math.NaN()}, 1, now32)},
		},
		{
			"upper(metric.test.foo,7)",
			map[parser.MetricRequest][]*types.MetricData{
				{Metric: "metric.test.foo", From: 0, Until: 1}: {types.MakeMetricData("metric.test.foo", []float64{1, 2, 0, 7, 8, 20, 30, math.NaN()}, 1, now32)},
			},
			[]*types.MetricData{types.MakeMetricData("metric.Test.foo",
				[]float64{1, 2, 0, 7, 8, 20, 30, math.NaN()}, 1, now32)},
		},
		{
			"upper(metric.test.foo,-3)",
			map[parser.MetricRequest][]*types.MetricData{
				{Metric: "metric.test.foo", From: 0, Until: 1}: {types.MakeMetricData("metric.test.foo", []float64{1, 2, 0, 7, 8, 20, 30, math.NaN()}, 1, now32)},
			},
			[]*types.MetricData{types.MakeMetricData("metric.test.Foo",
				[]float64{1, 2, 0, 7, 8, 20, 30, math.NaN()}, 1, now32)},
		},
		{
			"upper(metric.test.foo,0,7,12)",
			map[parser.MetricRequest][]*types.MetricData{
				{Metric: "metric.test.foo", From: 0, Until: 1}: {types.MakeMetricData("metric.test.foo", []float64{1, 2, 0, 7, 8, 20, 30, math.NaN()}, 1, now32)},
			},
			[]*types.MetricData{types.MakeMetricData("Metric.Test.Foo",
				[]float64{1, 2, 0, 7, 8, 20, 30, math.NaN()}, 1, now32)},
		},
		{
			"toUpperCase(metric.test.foo)",
			map[parser.MetricRequest][]*types.MetricData{
				{Metric: "metric.test.foo", From: 0, Until: 1}: {types.MakeMetricData("metric.test.foo", []float64{1, 2, 0, 7, 8, 20, 30, math.NaN()}, 1, now32)},
			},
			[]*types.MetricData{types.MakeMetricData("METRIC.TEST.FOO",
				[]float64{1, 2, 0, 7, 8, 20, 30, math.NaN()}, 1, now32)},
		},
	}

	for _, tt := range tests {
		testName := tt.Target
		t.Run(testName, func(t *testing.T) {
			eval := th.EvaluatorFromFunc(md[0].F)
			th.TestEvalExpr(t, eval, &tt)
		})
	}
}
