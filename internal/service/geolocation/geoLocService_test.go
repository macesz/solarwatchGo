package geolocation

import (
	"testing"
)

func TestGeoLocService_GetReport(t *testing.T) {
	service := Service{
		apiId: "3c4308a7a3e66d61d8e7f4b1cc5ec4bc",
	}

	resp, err := service.GetReport("London", "GB", "")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(resp)
}
