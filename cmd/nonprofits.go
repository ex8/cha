package cmd

import (
	"bytes"
	"fmt"

	"github.com/spf13/cobra"
)

type nonprofit struct {
	IconURL        string `json:"icon_url"`
	ID             string `json:"id"`
	Name           string `json:"name"`
	Ein            string `json:"ein"`
	Memo           string `json:"memo"`
	AddressLine    string `json:"address_line"`
	City           string `json:"city"`
	Classification string `json:"classification"`
	Mission        string `json:"mission"`
	State          string `json:"state"`
	Website        string `json:"website"`
	ZipCode        string `json:"zip_code"`
}

type searchNonprofitsResponse struct {
	Nonprofits []nonprofit `json:"nonprofits"`
	Page       int         `json:"page"`
}

var nonprofitsCmd = &cobra.Command{
	Use:   "nonprofits [command]",
	Short: "List or search Nonprofits",
	Long:  "Nonprofits CLI enables you to list or search nonprofits with the Change API.",
}

var nonprofitsName string
var nonprofitsPage int

var searchNonprofits = &cobra.Command{
	Use:   "search",
	Short: "Search all non-profits by name",
	Long:  "A list of searchable non-profits within the Change platform.",
	Run: func(cmd *cobra.Command, args []string) {
		payload := []byte(fmt.Sprintf(`{"name":"%v","page":"%v"}`, nonprofitsName, nonprofitsPage))
		var p searchNonprofitsResponse
		if err := client.Run("GET", "nonprofits", bytes.NewBuffer(payload), &p); err != nil {
			panic(err)
		}
		fmt.Printf(prettyPrint(p))
	},
}

var fetchNonprofitById = &cobra.Command{
	Use:   "find [id]",
	Short: "Find a Nonprofit by ID.",
	Long:  "Find your Nonprofit by using the unique identifier.",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		endpoint := fmt.Sprintf("nonprofits/%s", args[0])
		var p nonprofit
		if err := client.Run("GET", endpoint, &bytes.Buffer{}, &p); err != nil {
			panic(err)
		}
		fmt.Printf(prettyPrint(p))
	},
}

func init() {
	searchNonprofits.Flags().StringVarP(&nonprofitsName, "name", "n", "", "The non-profit name to search.")
	searchNonprofits.Flags().IntVarP(&nonprofitsPage, "page", "p", 1, "The cursor pagination page number.")

	nonprofitsCmd.AddCommand(searchNonprofits)
	nonprofitsCmd.AddCommand(fetchNonprofitById)
}
