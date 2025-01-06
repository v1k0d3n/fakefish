package api

import (
	"net/http"
	"strings"
)

type Manager struct {
	OdataType          string `json:"@odata.type"`
	Id                 string `json:"Id"`
	Name               string `json:"Name"`
	ManagerType        string `json:"ManagerType"`
	UUID               string `json:"UUID"`
	Model              string `json:"Model"`
	FirmwareVersion    string `json:"FirmwareVersion"`
	Status             Status `json:"Status"`
	EthernetInterfaces Link   `json:"EthernetInterfaces"`
	Links              Links  `json:"Links"`
	OdataContext       string `json:"@odata.context"`
	OdataID            string `json:"@odata.id"`
	RedfishCopyright   string `json:"@Redfish.Copyright"`
}

type ManagersResponse struct {
	OdataID string `json:"@odata.id"`
	Name    string `json:"Name"`
}

func ManagersHandler(w http.ResponseWriter, r *http.Request) {
	var response interface{}
	if strings.HasSuffix(r.URL.Path, "/1") {
		response = Manager{
			OdataType:       "#Manager.v1_1_0.Manager",
			Id:              "1",
			Name:            "fake-manager",
			ManagerType:     "BMC",
			UUID:            "1",
			Model:           "FakeManagerModel",
			FirmwareVersion: "1.0.0",
			Status: Status{
				State:        "Enabled",
				Health:       "OK",
				HealthRollUp: "OK",
			},
			EthernetInterfaces: Link{OdataID: "/redfish/v1/Managers/1/EthernetInterfaces"},
			Links: Links{
				Chassis:   []Link{{OdataID: "/redfish/v1/Chassis/fake-chassis"}},
				ManagedBy: []Link{{OdataID: "/redfish/v1/Managers/1"}},
			},
			OdataContext:     "/redfish/v1/$metadata#Manager.Manager",
			OdataID:          "/redfish/v1/Managers/1",
			RedfishCopyright: "Copyright 2014-2016 Distributed Management Task Force, Inc. (DMTF). For the full DMTF copyright policy, see http://www.dmtf.org/about/policies/copyright.",
		}
	} else {
		response = ManagersResponse{
			OdataID: "/redfish/v1/Managers",
			Name:    "Manager Collection",
		}
	}

	WritePrettyJSON(w, response)
}
