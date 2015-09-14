package bpost

import (
	"log"
	"testing"
)

func TestTrackingFetch(t *testing.T) {
	client := NewClient(nil, "https://api.bpost.be/services/trackedmail/item/", "116907", "001CrazyChicken")
	resp, err := client.FetchTracking("323211690759906489247126")
	if err != nil {
		t.Error(err)
	}
	log.Println(resp)
}
