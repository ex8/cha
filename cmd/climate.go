package cmd

import (
	"github.com/spf13/cobra"
)

type draftCarbonOffset struct {
	Amount   int    `json:"amount"`
	LiveMode bool   `json:"live_mode"`
	OffsetId string `json:"offset_id"`
}

type createCarbonOffset struct {
	Amount     int    `json:"amount"`
	ID         string `json:"id"`
	LiveMode   bool   `json:"live_mode"`
	MerchantId string `json:"merchant_id"`
	OrderValue int    `json:"order_value"`
	ZipCode    string `json:"zip_code"`
	ExternalId string `json:"external_id"`
}

var climateCmd = &cobra.Command{
	Use:   "climate [command]",
	Short: "Draft or create Climate offsets",
	Long:  "The Climate CLI enables you to draft or create carbon offsets with the Change API.",
}
