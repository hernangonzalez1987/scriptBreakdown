package entity

import (
	"reflect"
	"testing"
)

func TestScriptBreakdown_GetSceneBreakdownByNumber(t *testing.T) {
	t.Parallel()

	type fields struct {
		SceneBreakdowns []SceneBreakdown
	}

	type args struct {
		number int
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   *SceneBreakdown
	}{
		{
			name:   "should find scene",
			fields: fields{SceneBreakdowns: []SceneBreakdown{{Number: 1}, {Number: 2}, {Number: 3}}},
			args:   args{number: 1},
			want:   &SceneBreakdown{Number: 1},
		},
		{
			name:   "should not find scene",
			fields: fields{SceneBreakdowns: []SceneBreakdown{{Number: 1}, {Number: 2}, {Number: 3}}},
			args:   args{number: 4},
			want:   nil,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			ref := &ScriptBreakdown{
				SceneBreakdowns: testCase.fields.SceneBreakdowns,
			}
			if got := ref.GetSceneBreakdownByNumber(testCase.args.number); !reflect.DeepEqual(got, testCase.want) {
				t.Errorf("ScriptBreakdown.GetSceneBreakdownByNumber() = %v, want %v", got, testCase.want)
			}
		})
	}
}
