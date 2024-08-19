package autoburn_test

import (
	"testing"

	keepertest "duke/testutil/keeper"
	"duke/testutil/nullify"
	"duke/x/autoburn"
	"duke/x/autoburn/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.AutoburnKeeper(t)
	autoburn.InitGenesis(ctx, *k, genesisState)
	got := autoburn.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
