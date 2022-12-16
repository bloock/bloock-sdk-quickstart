package main

import (
	"fmt"

	"github.com/bloock/bloock-sdk-go/v2/builder"
	"github.com/bloock/bloock-sdk-go/v2/client"
	"github.com/bloock/bloock-sdk-go/v2/client/entity"
	"github.com/bloock/bloock-sdk-quickstart/utils"
	"github.com/bloock/bloock-sdk-quickstart/utils/logger"
)

func main() {
	utils.Sample("set proof", func(c utils.Config) error {
		record, err := builder.NewRecordBuilderFromString("Hello world").Build()
		if err != nil {
			return err
		}
		// we can set the proof of a record
        record.SetProof(entity.Proof{
        	Leaves: []string{"b20422fcfbe5b818ee305b44d5aaf427d103f8c5791b79c772ce82c747b2cd0d"},
        	Nodes:  []string{"b20422fcfbe5b818ee305b44d5aaf427d103f8c5791b79c772ce82c747b2cd0d"},
        	Depth:  "10101010",
        	Bitmap: "01010101",
        	Anchor: entity.ProofAnchor{
        		AnchorID: 10,
        		Networks: []entity.AnchorNetwork{{
        			Name:   "network",
        			State:  "state",
        			TxHash: "b20422fcfbe5b818ee305b44d5aaf427d103f8c5791b79c772ce82c747b2cd0d",
        		}},
        		Root:     "b20422fcfbe5b818ee305b44d5aaf427d103f8c5791b79c772ce82c747b2cd0d",
        		Status:   "success",
        	},
        })

        sdk := client.NewClient()
        proof, err := sdk.GetProof([]entity.Record{record})
        if err != nil {
            return err
        }

        logger.Success(fmt.Sprintf("The following proof was set: %+v", proof))
		return nil
	})
}
