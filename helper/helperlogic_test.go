package helper

import (
	"reflect"
	"testing"
)

func TestCalculateMean(t *testing.T) {
	type args struct {
		data []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Test Calculate Mean",
			args: args{data: []float64{1, 3}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateMean(tt.args.data); got != tt.want {
				t.Errorf("CalculateMean() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculateTStar(t *testing.T) {
	type args struct {
		alpha            float64
		degreesOfFreedom int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Test Calculate T Star",
			args: args{alpha: 0.05, degreesOfFreedom: 10},
			want: 2.228,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateTStar(tt.args.alpha, tt.args.degreesOfFreedom); got != tt.want {
				t.Errorf("CalculateTStar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculateStandardDeviation(t *testing.T) {
	type args struct {
		data []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Test Calculate Standard Deviation",
			args: args{data: []float64{1, 3}},
			want: 1.414,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateStandardDeviation(tt.args.data); got != tt.want {
				t.Errorf("CalculateStandardDeviation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBootstrapResample(t *testing.T) {
	type args struct {
		data []float64
	}
	tests := []struct {
		name string
		args args
		want [][]float64
	}{
		{
			name: "Test Bootstrap Resample",
			args: args{data: []float64{1, 3}},
			//Bootstrapping can have multiple outputs:
			want: [][]float64{
				{1, 3},
				{3, 1},
				{1, 1},
				{3, 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BootstrapResample(tt.args.data)

			// Check want list for any one of the possible outcomes
			var found bool
			for _, acceptableOutcome := range tt.want {
				if reflect.DeepEqual(got, acceptableOutcome) {
					found = true
					break
				}
			}

			if !found {
				t.Errorf("BootstrapResample() = %v, want one of %v", got, tt.want)
			}
		})
	}
}
func TestCalculateConfidenceInterval(t *testing.T) {
	type args struct {
		data  []float64
		level float64
	}
	tests := []struct {
		name string
		args args
		want [2]float64
	}{
		{
			name: "Test Calculate Confidence Interval",
			args: args{data: []float64{1, 3}, level: 0.95},
			want: [2]float64{1, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateConfidenceInterval(tt.args.data, tt.args.level); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CalculateConfidenceInterval() = %v, want %v", got, tt.want)
			}
		})
	}
}
