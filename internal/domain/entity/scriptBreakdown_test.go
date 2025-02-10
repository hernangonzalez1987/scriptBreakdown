package entity

import (
	"reflect"
	"testing"
)

func TestScriptBreakdown_GetSceneBreakdownByNumber(t *testing.T) {
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
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ref := &ScriptBreakdown{
				SceneBreakdowns: tt.fields.SceneBreakdowns,
			}
			if got := ref.GetSceneBreakdownByNumber(tt.args.number); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ScriptBreakdown.GetSceneBreakdownByNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
