package solarwatch

import "testing"

func TestSolarWatchService(t *testing.T) {
	service := Service{}

	resp, err := service.GetRiport(51.50732, -0.1276474)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(resp)
}
