package handler

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHandler_jsonError(t *testing.T) {
	msg := "JSON Error"
	result := jsonError(msg)
	require.Equal(t, string([]byte(`{"message":"JSON Error"}`)), string(result))
}
