package keeper_test

import (
	"testing"

	testkeeper "duke/testutil/keeper"
	"duke/x/autoburn/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.AutoburnKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
