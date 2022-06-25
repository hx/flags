package states_test

import (
	. "github.com/hx/flags/states"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestState_Get(t *testing.T) {
	assert.True(t, State(5).Get(0))
	assert.False(t, State(5).Get(1))
	assert.True(t, State(5).Get(2))
}

func TestState_Set(t *testing.T) {
	assert.EqualValues(t, 5, State(4).Set(0, true))
	assert.EqualValues(t, 4, State(4).Set(0, false))
	assert.EqualValues(t, 6, State(4).Set(1, true))
	assert.EqualValues(t, 4, State(4).Set(1, false))
	assert.EqualValues(t, 4, State(4).Set(2, true))
	assert.EqualValues(t, 0, State(4).Set(2, false))
	assert.EqualValues(t, 5, State(1).Set(2, true))
}

func TestState_Apply(t *testing.T) {
	assert.EqualValues(t, 2, State(1).Apply(Diff{0: false, 1: true}))
	assert.EqualValues(t, 2, State(2).Apply(Diff{1: true}))
	assert.EqualValues(t, 3, State(1).Apply(Diff{1: true}))
}

func TestState_Diff(t *testing.T) {
	assert.Equal(t, Diff{}, State(0).Diff(0))
	assert.Equal(t, Diff{0: false}, State(0).Diff(1))
	assert.Equal(t, Diff{0: true}, State(3).Diff(2))
}

func TestState_Count(t *testing.T) {
	assert.Equal(t, 2, State(5).Count())
}
