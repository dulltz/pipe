package prometheus

import (
	"context"
	"testing"

	"github.com/prometheus/common/model"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	"github.com/pipe-cd/pipe/pkg/config"
)

func TestType(t *testing.T) {
	fake := fakeAPI{
		value: newScalar(10),
	}
	p := Provider{
		api:     fake,
		timeout: defaultTimeout,
		logger:  zap.NewNop(),
	}
	assert.Equal(t, ProviderType, p.Type())
}

func TestRunQuery(t *testing.T) {
	cases := []struct {
		name       string
		value      model.Value
		expected   config.AnalysisExpected
		wantResult bool
		wantErr    bool
	}{
		{
			name:  "successfully with scalar value",
			value: newScalar(1),
			expected: config.AnalysisExpected{
				Min: float64Pointer(0),
				Max: float64Pointer(2),
			},
			wantResult: true,
			wantErr:    false,
		},
		{
			name:  "successfully with vector value",
			value: newVector(1),
			expected: config.AnalysisExpected{
				Min: float64Pointer(0),
				Max: float64Pointer(2),
			},
			wantResult: true,
			wantErr:    false,
		},
		{
			name:  "failure with scalar value",
			value: newScalar(1),
			expected: config.AnalysisExpected{
				Min: float64Pointer(2),
				Max: float64Pointer(3),
			},
			wantResult: false,
			wantErr:    false,
		},
		{
			name:  "failure with vector value",
			value: newVector(1),
			expected: config.AnalysisExpected{
				Min: float64Pointer(2),
				Max: float64Pointer(3),
			},
			wantResult: false,
			wantErr:    false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			p := Provider{
				api: fakeAPI{
					value: tc.value,
				},
				timeout: defaultTimeout,
				logger:  zap.NewNop(),
			}
			res, err := p.RunQuery(context.Background(), "dummy", tc.expected)
			assert.Equal(t, tc.wantErr, err != nil)
			assert.Equal(t, res, tc.wantResult)
		})
	}

}

func float64Pointer(i float64) *float64 { return &i }

func newScalar(f float64) model.Value {
	return &model.Scalar{
		Value:     model.SampleValue(f),
		Timestamp: model.Time(0),
	}
}

func newVector(f float64) model.Value {
	return model.Vector{
		{
			Value: model.SampleValue(f),
		},
	}
}
