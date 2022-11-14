package main

import (
	"github.com/bloock/bloock-sdk-go/v2"
	"github.com/bloock/bloock-sdk-go/v2/builder"
	"github.com/bloock/bloock-sdk-go/v2/client"
	"github.com/bloock/bloock-sdk-quickstart/utils"
)

func main() {
	utils.Sample("client", func(c utils.Config) error {
        bloock.ApiKey = c.ApiKey
        bloock.ApiHost = c.ApiHost
        sdk := client.NewClient()

		record, err := builder.NewRecordBuilderFromString("Hello world 1").Build()
		if err != nil {
			return err
		}
		hash, err := record.GetHash()
		if err != nil {
			return err
		}
		records := []string{hash}

		receipt, err := sdk.SendRecords(records)
		require.NoError(t, err)
		assert.Greater(t, len(receipt), 0)
		assert.NotEqual(t, entity.RecordReceipt{}, receipt[0])

		anchor, err := sdk.WaitAnchor(receipt[0].Anchor, entity.NewAnchorParams())
		require.NoError(t, err)
		assert.Equal(t, receipt[0].Anchor, anchor.Id)

		proof, err := sdk.GetProof(records)
		require.NoError(t, err)
		assert.NotEqual(t, entity.Proof{}, proof)

		root, err := sdk.VerifyProof(proof)
		require.NoError(t, err)
		assert.NotEqual(t, "", proof)

		timestampValidateRoot, err := sdk.ValidateRoot(root, entity.ListOfNetworks().BloockChain)
		require.NoError(t, err)
		assert.Greater(t, timestampValidateRoot, uint64(0))

		network := entity.NewNetworkParams()
		network.Network = entity.ListOfNetworks().BloockChain
		timestampVerifyRecords, err := sdk.VerifyRecords(records, network)
		require.NoError(t, err)
		assert.Greater(t, timestampVerifyRecords, uint64(0))

		assert.Equal(t, timestampValidateRoot, timestampVerifyRecords)

		return nil
	})
}
