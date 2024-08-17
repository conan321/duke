package duke_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "duke/testutil/keeper"
	"duke/testutil/nullify"
	"duke/x/duke"
	"duke/x/duke/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.DukeKeeper(t)
	duke.InitGenesis(ctx, *k, genesisState)
	got := duke.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
