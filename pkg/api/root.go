package api

import (
	"net/http"
)

type RootResponse struct {
	OdataType        string `json:"@odata.type"`
	ID               string `json:"Id"`
	Name             string `json:"Name"`
	RedfishVersion   string `json:"RedfishVersion"`
	UUID             string `json:"UUID"`
	Systems          Link   `json:"Systems"`
	Managers         Link   `json:"Managers"`
	OdataID          string `json:"@odata.id"`
	RedfishCopyright string `json:"@Redfish.Copyright"`
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	response := RootResponse{
		OdataType:        "#ServiceRoot.v1_5_0.ServiceRoot",
		ID:               "FakeFishService",
		Name:             "FakeFish Service",
		RedfishVersion:   "1.5.0",
		UUID:             "not-that-production-ready",
		Systems:          Link{OdataID: "/redfish/v1/Systems"},
		Managers:         Link{OdataID: "/redfish/v1/Managers"},
		OdataID:          "/redfish/v1/",
		RedfishCopyright: "Copyright 2014-2016 Distributed Management Task Force, Inc. (DMTF). For the full DMTF copyright policy, see http://www.dmtf.org/about/policies/copyright.",
	}

	WritePrettyJSON(w, response)
}
