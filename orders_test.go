package bpost

import (
	"log"
	"testing"
)

func TestOrdersFetch(t *testing.T) {
	client := NewClient(nil, "https://api.bpost.be/services/shm/", "116907", "001CrazyChicken")
	resp, err := client.FetchOrder("116907-1439460160262")
	if err != nil {
		t.Error(err)
	}
	log.Println(resp)
}
