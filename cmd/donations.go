package cmd

import (
	"bytes"
	"fmt"

	"github.com/spf13/cobra"
)

type donation struct {
	Amount         int    `json:"amount"`
	NonprofitId    string `json:"nonprofit_id"`
	FundsCollected bool   `json:"funds_collected"`
	ZipCode        string `json:"zip_code"`
	OrderValue     int    `json:"order_value"`
	ExternalId     string `json:"external_id"`
	LiveMode       bool   `json:"live_mode"`
}

type listDonationsResponse struct {
	Donations []donation `json:"donations"`
	Page      int        `json:"page"`
}

var donationsCmd = &cobra.Command{
	Use:   "donations [command]",
	Short: "List, find, or create Donations",
	Long:  "Donations CLI enables you to list, find, or create donations using the Change API.",
}

var donationPageCursor int

var listDonations = &cobra.Command{
	Use:   "list",
	Short: "List your Donations.",
	Long:  "A list of donations associated with your account.",
	Run: func(cmd *cobra.Command, args []string) {
		payload := []byte(fmt.Sprintf(`{"page":"%v"}`, donationPageCursor))
		var d listDonationsResponse
		if err := client.Run("GET", "donations", bytes.NewBuffer(payload), &d); err != nil {
			panic(err)
		}
		fmt.Printf(prettyPrint(d))
	},
}

var findDonationById = &cobra.Command{
	Use:   "find [id]",
	Short: "Find a Donation by ID.",
	Long:  "Find your Donation by using the unique identifier.",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		endpoint := fmt.Sprintf("donations/%s", args[0])
		var d donation
		if err := client.Run("GET", endpoint, &bytes.Buffer{}, &d); err != nil {
			panic(err)
		}
		fmt.Printf(prettyPrint(d))
	},
}

var donationAmount int
var donationNonProfitId string
var donationFundsCollected bool
var donationZipCode string
var donationOrderValue int
var donationExternalId string

var createDonation = &cobra.Command{
	Use:   "create",
	Short: "Create a new Donation",
	Long:  "Create a new Donation record.",
	Run: func(cmd *cobra.Command, args []string) {
		vars := []interface{}{
			donationAmount,
			donationNonProfitId,
			donationFundsCollected,
			donationZipCode,
			donationOrderValue,
			donationExternalId,
		}
		payload := []byte(fmt.Sprintf(`{"amount":"%v","nonprofit_id":"%v","funds_collected":"%v","zip_code":"%v","order_value":"%v","external_id":"%v"}`, vars...))
		var d donation
		if err := client.Run("POST", "donations", bytes.NewBuffer(payload), &d); err != nil {
			panic(err)
		}
		fmt.Printf(prettyPrint(d))
	},
}

func init() {
	listDonations.Flags().IntVarP(&donationPageCursor, "page", "p", 1, "The cursor pagination page number.")

	createDonation.Flags().IntVarP(&donationAmount, "amount", "a", 0, "The donation amount (in cents).")
	createDonation.Flags().StringVarP(&donationNonProfitId, "nonprofitId", "n", "", "The nonprofit id you want to donate to.")
	createDonation.Flags().BoolVarP(&donationFundsCollected, "fundsCollected", "f", false, "Whether you are collecting payment for the carbon offset.")
	createDonation.Flags().StringVarP(&donationZipCode, "zipCode", "z", "", "The zip code to associate with the donation.")
	createDonation.Flags().IntVarP(&donationOrderValue, "orderValue", "o", 0, "The cart or order volume (in cents) associated with the donation.")
	createDonation.Flags().StringVarP(&donationExternalId, "externalId", "e", "", "An external ID associated with the donation.")
	createDonation.MarkFlagRequired("amount")
	createDonation.MarkFlagRequired("nonprofitId")
	createDonation.MarkFlagRequired("fundsCollected")

	donationsCmd.AddCommand(listDonations)
	donationsCmd.AddCommand(findDonationById)
	donationsCmd.AddCommand(createDonation)
}
