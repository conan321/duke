package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "duke/testutil/keeper"
	"duke/x/duke/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.DukeKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
