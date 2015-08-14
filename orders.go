package bpost

import "encoding/xml"

type OrderInfo struct {
	AccountID      int         `xml:"accountId"`
	OrderReference string      `xml:"reference"`
	Box            []Box       `xml:"box"`
	Status         string      `xml:"status"`
	CostCenter     string      `xml:"costCenter"`
	Orders         []OrderLine `xml:"orderLines"`
}

type Box struct {
	Sender   Customer `xml:"sender"`
	Receiver Customer `xml:"nationalBox>atHome>receiver"`
	Status   string   `xml:"status"`
	Barcode  string   `xml:"barcode"`
}

type Customer struct {
	Name    string  `xml:"name"`
	Company string  `xml:"company"`
	Email   string  `xml:"emailAddress"`
	Address Address `xml:"address"`
}

type Address struct {
	Street      string `xml:"streetName"`
	Number      string `xml:"number"`
	Zip         string `xml:"postalCode"`
	City        string `xml:"locality"`
	CountryCode string `xml:"countryCode"`
}

type OrderLine struct {
	Text  string `xml:"text"`
	Count int    `xml:"nbOfItems"`
}

// FetchOrder retrieves a single order by BPOST reference.
func (c *Client) FetchOrder(ref string) (*OrderInfo, error) {
	req, err := c.NewRequest("GET", "orders/"+ref, nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	response := OrderInfo{}
	if err := xml.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}
	return &response, nil
}
