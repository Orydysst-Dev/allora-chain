package integration_test

import (
	"context"
	"os"
	"testing"

	chain_test "github.com/allora-network/allora-chain/integration/chain"
	"github.com/stretchr/testify/require"
)

type TestMetadata struct {
	t   *testing.T
	ctx context.Context
	n   chain_test.Node
}

func Setup(t *testing.T) TestMetadata {
	ret := TestMetadata{}
	ret.t = t
	var err error
	ret.ctx = context.Background()
	// userHomeDir, _ := os.UserHomeDir()
	// home := filepath.Join(userHomeDir, ".allorad")
	node, err := chain_test.NewNode(
		t,
		chain_test.NodeConfig{
			NodeRPCAddress: "http://localhost:26657",
			AlloraHomeDir:  "./devnet/genesis",
		},
	)
	require.NoError(t, err)
	ret.n = node
	return ret
}

func TestExternalTestSuite(t *testing.T) {
	if _, isIntegration := os.LookupEnv("INTEGRATION"); isIntegration == false {
		t.Skip("Skipping Integration Test Outside CI")
	}
	t.Log(">>> Setting up connection to local node <<<")
	m := Setup(t)
	t.Log(">>> Test Getting Chain Params <<<")
	GetParams(m)
	t.Log(">>> Test Topic Creation <<<")
	CreateTopic(m)
	// TODO Uncomment as we stand up devnet
	// t.Log(">>> Test Distribution Checks <<<")
	// DistributionChecks(m)
	t.Log(">>> Test Actor Registration <<<")
	RegistrationChecks(m)
	t.Log(">>> Test Update Params <<<")
	UpdateParamsChecks(m)
	t.Log(">>> Test Reputer Staking <<<")
	StakingChecks(m)
	t.Log(">>> Test Topic Funding and Activation <<<")
	TopicFundingChecks(m)
	t.Log(">>> Test Making Inference <<<")
	WorkerInferenceAndForecastChecks(m)
}
