package cmd

import (
	"bytes"
	"fmt"

	"github.com/spf13/cobra"
)

var shippingClimateCmd = &cobra.Command{
	Use:   "shipping",
	Short: "Draft and create Shipping Carbon Offsets",
	Long:  "Climate Shipping CLI enables you to draft or create carbon offsets using the Change API.",
}

var draftShippingTransportMethod string
var draftShippingWeight int
var draftShippingDestinationAddressZipCode int
var draftShippingDistanceMiles int
var draftShippingOriginAddressZipCode int
var draftShippingOffsetId string

var draftShippingOffset = &cobra.Command{
	Use:   "draft",
	Short: "Draft Shipping Carbon Offset",
	Long:  "Draft Shipping Carbon Offset.",
	Run: func(cmd *cobra.Command, args []string) {
		vars := []interface{}{
			draftShippingTransportMethod,
			draftShippingWeight,
			draftShippingDestinationAddressZipCode,
			draftShippingDistanceMiles,
		}
		payload := []byte(fmt.Sprintf(`{"transportation_method":"%v","weight_lb":"%v","destination_address":"%v","distance_mi":"%v","origin_address":"%v"}`, vars...))
		var o draftCarbonOffset
		if err := client.Run("GET", "climate/shipping_offset", bytes.NewBuffer(payload), &o); err != nil {
			panic(err)
		}
		fmt.Printf(prettyPrint(o))
	},
}

var createShippingFundsCollected bool
var createShippingDestinationAddress int
var createShippingDistanceMiles int
var createShippingOffsetId string
var createShippingOriginAddress int
var createShippingTransportationMethod string
var createShippingWeight int
var createShippingZipCode int

var createShippingOffset = &cobra.Command{
	Use:   "create",
	Short: "Create Shipping Carbon Offset",
	Long:  "Create Shipping Carbon Offset.",
	Run: func(cmd *cobra.Command, args []string) {
		vars := []interface{}{
			createShippingFundsCollected,
			createShippingDestinationAddress,
			createShippingDistanceMiles,
			createShippingOffsetId,
			createShippingOriginAddress,
			createShippingTransportationMethod,
			createShippingWeight,
			createShippingZipCode,
		}
		payload := []byte(fmt.Sprintf(`{"funds_collected":"%v","destination_address":"%v","distance_mi":"%v","offset_id":"%v","origin_address":"%v","transportation_method":"%v","weight_lb":"%v","zip_code":"%v"}`, vars...))
		var o createCarbonOffset
		if err := client.Run("POST", "climate/shipping_offset", bytes.NewBuffer(payload), &o); err != nil {
			panic(err)
		}
		fmt.Printf(prettyPrint(o))
	},
}

func init() {
	climateCmd.AddCommand(shippingClimateCmd)

	// Draft shipping
	draftShippingOffset.Flags().StringVarP(&draftShippingTransportMethod, "transportMethod", "t", "air", "The transporation method; enum values: air|truck|rail|sea")
	draftShippingOffset.Flags().IntVarP(&draftShippingWeight, "weight", "w", 0, "The total weight in pounds of the shipment.")
	draftShippingOffset.Flags().IntVarP(&draftShippingDestinationAddressZipCode, "destinationAddress", "d", 0, "The destination zip code (US only) of the shipment. If you send this parameter, also send origin_address.")
	draftShippingOffset.Flags().IntVarP(&draftShippingDistanceMiles, "distanceMiles", "m", 0, "The total distance (in miles) of the shipment. You can use this parameter in place of origin_address and destination_address.")
	draftShippingOffset.Flags().IntVarP(&draftShippingOriginAddressZipCode, "originAddress", "s", 0, "The origin zip code (US only) of the shipment. If you send this parameter, also send destination_address.")
	draftShippingOffset.MarkFlagRequired("transportMethod")
	draftShippingOffset.MarkFlagRequired("weight")

	// Create shipping
	createShippingOffset.Flags().BoolVarP(&createShippingFundsCollected, "fundsCollected", "f", false, "Whether you are collecting payment for the carbon offset.")
	createShippingOffset.Flags().IntVarP(&createShippingDestinationAddress, "destinationAddress", "d", 0, "The destination zip code (US only) of the shipment. If you send this parameter, also send origin_address.")
	createShippingOffset.Flags().StringVarP(&createShippingOffsetId, "offsetId", "o", "", "The ID for a drafted carbon offset. You can use this parameter in place of origin_address, destination_address, distance_mi, transporation_method and weight_lb.")
	createShippingOffset.Flags().IntVarP(&createShippingOriginAddress, "originAddress", "s", 0, "The origin zip code (US only) of the shipment. If you send this parameter, also send destination_address.")
	createShippingOffset.Flags().StringVarP(&createShippingTransportationMethod, "transportMethod", "t", "air", "The transporation method; enum values: air|truck|rail|sea")
	createShippingOffset.Flags().IntVarP(&createShippingWeight, "weight", "w", 0, "The total weight in pounds of the shipment.")
	createShippingOffset.Flags().IntVarP(&createShippingZipCode, "zipCode", "z", 0, "The customer's zip code. Provide this to unlock geographic insights.")
	createShippingOffset.MarkFlagRequired("fundsCollected")

	shippingClimateCmd.AddCommand(draftShippingOffset)
	shippingClimateCmd.AddCommand(createShippingOffset)
}
