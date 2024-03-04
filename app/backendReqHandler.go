package app

import (
	"backendReq/dto"
	"backendReq/logger"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type BackendReqHandler struct {
}

// retrive list of handler code starts here//
func (rloph BackendReqHandler) GetRetriveListOfProduct(w http.ResponseWriter, r *http.Request) {

	const baseurl = "https://frameplay.com/api/v2/catalog/products"
	client := http.Client{}
	req, err := http.NewRequest(http.MethodGet, baseurl, nil)

	req.SetBasicAuth("abc", "abc")

	req.Header.Set("Content-Type", "application/json") // => your content-type
	if err != nil {
		logger.Error("error while sending request" + err.Error())
		return
	}
	resp, err := client.Do(req)
	if err != nil {
		logger.Error("error while getting response" + err.Error())
		return
	}
	defer resp.Body.Close()

	receivedResponse, err := io.ReadAll(resp.Body)
	if err != nil {
		writeResponse(w, resp.StatusCode, err.Error())
	}
	//now transform the reseviced response into the RESPONSE DTO//
	var rppnres dto.ProductsResponse

	if err := json.Unmarshal(receivedResponse, &rppnres); err != nil {

		logger.Error("Error unmarshalling received response:" + err.Error())
		return

	}
	//apply filter on the response data//
	var filtered []dto.PayLoadDetails
	for _, v := range rppnres.PayLoad {
		if v.Category == "GiftCard" {
			filtered = append(filtered, v)
		}
	}
	// Flatten the JSON data
	dataMap := make(map[string]string)
	for _, obj := range filtered {
		flatten(obj, []string{}, dataMap)
	}
	// // Create the uploads folder if it doesn't already exist
	err = os.MkdirAll("./ppnuploads", os.ModePerm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Create a CSV file
	csvFile, err := os.Create("ppnuploads/" + "productlist-" + time.Now().Format(time.DateTime))
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()
	// Write the CSV header
	csvWriter := csv.NewWriter(csvFile)
	// Write the CSV data
	dataRows := make([][]string, 0)
	for _, row := range filtered {
		dataRow := []string{strconv.Itoa(row.SKUID), strconv.Itoa(row.ProductID), row.ProductName, row.Category, strconv.FormatBool(row.IsSalesTaxCharged), row.CountryCode, row.BenefitType, row.Validity, row.ProductDesc, strconv.Itoa(row.LocalPhoneNumberLength), strconv.FormatBool(row.AllowDecimal), strconv.FormatFloat(row.Fee, 'f', 1, 64), strconv.Itoa(row.OperatorId), row.OperatorName, row.ImageUrl, row.AdditionalInformation, strconv.FormatFloat(row.MinFaceValue.FaceValue, 'f', 1, 64), row.MinFaceValue.FaceValueCurrency, strconv.FormatFloat(row.MinFaceValue.Cost, 'f', 1, 64), row.MinFaceValue.CostCurrency, strconv.FormatFloat(row.MaxFaceValue.FaceValue, 'f', 1, 64), row.MaxFaceValue.FaceValueCurrency, strconv.FormatFloat(row.MaxFaceValue.Cost, 'f', 1, 64), row.MaxFaceValue.CostCurrency}
		dataRows = append(dataRows, dataRow)
	}
	err = csvWriter.WriteAll(dataRows)
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully converted slice to CSV file")

	// send the response as per the dto//
	writeResponse(w, http.StatusOK, "product saved")

}

// This function recursively extracts all key-value pairs from a nested JSON object
func flatten(data interface{}, path []string, result map[string]string) {
	switch d := data.(type) {
	case map[string]interface{}:
		for key, value := range d {
			newPath := append(path, key)
			flatten(value, newPath, result)
		}
	case []interface{}:
		for i, value := range d {
			newPath := append(path, fmt.Sprintf("%d", i))
			flatten(value, newPath, result)
		}
	case string, int, float64, bool:
		result[strings.Join(path, ".")] = fmt.Sprintf("%v", d)
	}
}
