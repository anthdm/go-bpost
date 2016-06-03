package bpost

import "encoding/xml"

type TrackingInfo struct {
	ItemCode   string      `xml:"itemCode"`
	CostCenter string      `xml:"costCenter"`
	StateInfo  []StateInfo `xml:"stateInfo"`
}

type StateInfo struct {
	Time        string `xml:"time"`
	Code        string `xml:"stateCode"`
	Description string `xml:"stateDescription"`
}

// FetchTracking retrieves a single order by BPOST barcode.
func (c *Client) FetchTracking(ref string) (*TrackingInfo, error) {
	req, err := c.NewRequest("GET", ref+"/trackingInfo", true, nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	response := TrackingInfo{}
	if resp.StatusCode == http.StatusOK {
		if err := xml.NewDecoder(resp.Body).Decode(&response); err != nil {
			return nil, err
		}
	}

	return &response, nil
}
