package cache

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	t.Parallel()

	ttl := time.Millisecond * 10
	expectedValue := "someOtherValue"

	stringCache := Cache[string]{
		mu: sync.RWMutex{},
		data: map[string]element[string]{
			"someExpiredKey": {
				value:      "someValue",
				insertTime: time.Now().Add(time.Millisecond * 10 * -1),
			},
			"someNonExpiredKey": {
				value:      expectedValue,
				insertTime: time.Now(),
			},
		},
		ttl: &ttl,
	}

	tests := []struct {
		name      string
		arg       string
		wantValue *string
		wantFound bool
	}{
		{
			name:      "should not get expired object",
			arg:       "someExpiredKey",
			wantValue: nil,
			wantFound: false,
		},
		{
			name:      "should not get non existing object",
			arg:       "SomeNonExistingObject",
			wantValue: nil,
			wantFound: false,
		},
		{
			name:      "should get non expired object",
			arg:       "someNonExpiredKey",
			wantValue: &expectedValue,
			wantFound: true,
		},
	}
	for _, scenario := range tests {
		t.Run(scenario.name, func(t *testing.T) {
			t.Parallel()

			value, found := stringCache.Get(scenario.arg)
			assert.Equal(t, scenario.wantValue, value)
			assert.Equal(t, scenario.wantFound, found)
		})
	}
}

func TestSave(t *testing.T) {
	t.Parallel()

	ttl := time.Millisecond * 10

	stringCache := Cache[string]{
		mu:   sync.RWMutex{},
		data: map[string]element[string]{},
		ttl:  &ttl,
	}

	tests := []struct {
		name     string
		argKey   string
		argValue string
	}{
		{
			name:     "should save object",
			argKey:   "someKey",
			argValue: "someValue",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			stringCache.Save(tt.argKey, tt.argValue)
			assert.Equal(t, tt.argValue, stringCache.data[tt.argKey].value)
		})
	}
}
