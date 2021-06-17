package cmd

import (
	"bytes"
	"fmt"

	"github.com/spf13/cobra"
)

var cryptoClimateCmd = &cobra.Command{
	Use:   "crypto",
	Short: "Draft and create Crypto Carbon Offsets",
	Long:  "Climate Crypto CLI enables you to draft or create carbon offsets using the Change API.",
}

var draftCryptoCurrency string
var draftCryptoCount int

var draftCryptoOffset = &cobra.Command{
	Use:   "draft",
	Short: "Draft Crypto Carbon Offset",
	Long:  "Draft Crypto Carbon Offset.",
	Run: func(cmd *cobra.Command, args []string) {
		vars := []interface{}{draftCryptoCurrency, draftCryptoCount}
		payload := []byte(fmt.Sprintf(`{"currency":"%v","count":"%v"}`, vars...))
		var o draftCarbonOffset
		if err := client.Run("GET", "climate/crypto_offset", bytes.NewBuffer(payload), &o); err != nil {
			panic(err)
		}
		fmt.Printf(prettyPrint(o))
	},
}

var createCryptoFundsCollected bool
var createCryptoCount int
var createCryptoCurrency string
var createCryptoOffsetId string
var createCryptoZipCode string

var createCryptoOffset = &cobra.Command{
	Use:   "create",
	Short: "Create Crypto Carbon Offset",
	Long:  "Create Crypto Carbon Offset.",
	Run: func(cmd *cobra.Command, args []string) {
		vars := []interface{}{
			createCryptoFundsCollected,
			createCryptoCount,
			createCryptoCurrency,
			createCryptoOffsetId,
			createCryptoZipCode,
		}
		payload := []byte(fmt.Sprintf(`{"funds_collected":"%v","count":"%v","offset_id":"%v","zip_code":"%v"}`, vars...))
		var o createCarbonOffset
		if err := client.Run("POST", "climate/crypto_offset", bytes.NewBuffer(payload), &o); err != nil {
			panic(err)
		}
		fmt.Printf(prettyPrint(o))
	},
}

func init() {
	climateCmd.AddCommand(cryptoClimateCmd)

	// Draft
	draftCryptoOffset.Flags().StringVarP(&draftCryptoCurrency, "currency", "c", "eth", "The currency of the transaction.")
	draftCryptoOffset.Flags().IntVarP(&draftCryptoCount, "count", "k", 1, "The number of transactions to offset.")
	draftCryptoOffset.MarkFlagRequired("currency")

	// Create
	createCryptoOffset.Flags().BoolVarP(&createCryptoFundsCollected, "fundsCollected", "f", false, "Whether you are collecting payment for the carbon offset.")
	createCryptoOffset.Flags().IntVarP(&createCryptoCount, "count", "k", 0, "The number of transactions to offset.")
	createCryptoOffset.Flags().StringVarP(&createCryptoCurrency, "currency", "c", "eth", "The currency of the transaction.")
	createCryptoOffset.Flags().StringVarP(&createCryptoOffsetId, "offsetId", "o", "", "The ID for a drafted carbon offset. You can use this parameter in place of currency and count.")
	createCryptoOffset.Flags().StringVarP(&createCryptoZipCode, "zipCode", "z", "", "The customer's zip code. Provide this to unlock geographic insights.")
	createCryptoOffset.MarkFlagRequired("fundsCollected")

	cryptoClimateCmd.AddCommand(draftCryptoOffset)
	cryptoClimateCmd.AddCommand(createCryptoOffset)
}
