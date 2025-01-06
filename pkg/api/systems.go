package api

import (
	"net/http"
	"strings"
)

type Status struct {
	State        string `json:"State"`
	Health       string `json:"Health"`
	HealthRollUp string `json:"HealthRollUp"`
}

type Boot struct {
	BootSourceOverrideEnabled               string   `json:"BootSourceOverrideEnabled"`
	BootSourceOverrideTarget                string   `json:"BootSourceOverrideTarget"`
	BootSourceOverrideTargetAllowableValues []string `json:"BootSourceOverrideTarget@Redfish.AllowableValues"`
	BootSourceOverrideMode                  string   `json:"BootSourceOverrideMode"`
	UefiTargetBootSourceOverride            string   `json:"UefiTargetBootSourceOverride"`
}

type ProcessorSummary struct {
	Count  int    `json:"Count"`
	Status Status `json:"Status"`
}

type MemorySummary struct {
	TotalSystemMemoryGiB int    `json:"TotalSystemMemoryGiB"`
	Status               Status `json:"Status"`
}

type Links struct {
	Chassis   []Link `json:"Chassis"`
	ManagedBy []Link `json:"ManagedBy"`
}

type Actions struct {
	Reset struct {
		Target                   string   `json:"target"`
		ResetTypeAllowableValues []string `json:"ResetType@Redfish.AllowableValues"`
	} `json:"#ComputerSystem.Reset"`
}

type System struct {
	OdataType          string           `json:"@odata.type"`
	Id                 string           `json:"Id"`
	Name               string           `json:"Name"`
	UUID               string           `json:"UUID"`
	Manufacturer       string           `json:"Manufacturer"`
	Status             Status           `json:"Status"`
	PowerState         string           `json:"PowerState"`
	Boot               Boot             `json:"Boot"`
	ProcessorSummary   ProcessorSummary `json:"ProcessorSummary"`
	MemorySummary      MemorySummary    `json:"MemorySummary"`
	Bios               Link             `json:"Bios"`
	Processors         Link             `json:"Processors"`
	Memory             Link             `json:"Memory"`
	EthernetInterfaces Link             `json:"EthernetInterfaces"`
	SimpleStorage      Link             `json:"SimpleStorage"`
	Storage            Link             `json:"Storage"`
	IndicatorLED       string           `json:"IndicatorLED"`
	Links              Links            `json:"Links"`
	Actions            Actions          `json:"Actions"`
	OdataContext       string           `json:"@odata.context"`
	OdataID            string           `json:"@odata.id"`
	RedfishCopyright   string           `json:"@Redfish.Copyright"`
}

type SystemsResponse struct {
	OdataID string `json:"@odata.id"`
	Name    string `json:"Name"`
}

func SystemsHandler(w http.ResponseWriter, r *http.Request) {
	var response interface{}
	if strings.HasSuffix(r.URL.Path, "/1") {
		response = System{
			OdataType:    "#ComputerSystem.v1_1_0.ComputerSystem",
			Id:           "1",
			Name:         "fake-system",
			UUID:         "1",
			Manufacturer: "FakeFish",
			Status: Status{
				State:        "Enabled",
				Health:       "OK",
				HealthRollUp: "OK",
			},
			PowerState: "On",
			Boot: Boot{
				BootSourceOverrideEnabled:               "Continuous",
				BootSourceOverrideTarget:                "Hdd",
				BootSourceOverrideTargetAllowableValues: []string{"Pxe", "Cd", "Hdd"},
				BootSourceOverrideMode:                  "UEFI",
				UefiTargetBootSourceOverride:            "/0x31/0x33/0x01/0x01",
			},
			ProcessorSummary: ProcessorSummary{
				Count: 8,
				Status: Status{
					State:        "Enabled",
					Health:       "OK",
					HealthRollUp: "OK",
				},
			},
			MemorySummary: MemorySummary{
				TotalSystemMemoryGiB: 15,
				Status: Status{
					State:        "Enabled",
					Health:       "OK",
					HealthRollUp: "OK",
				},
			},
			Bios:               Link{OdataID: "/redfish/v1/Systems/1/BIOS"},
			Processors:         Link{OdataID: "/redfish/v1/Systems/1/Processors"},
			Memory:             Link{OdataID: "/redfish/v1/Systems/1/Memory"},
			EthernetInterfaces: Link{OdataID: "/redfish/v1/Systems/1/EthernetInterfaces"},
			SimpleStorage:      Link{OdataID: "/redfish/v1/Systems/1/SimpleStorage"},
			Storage:            Link{OdataID: "/redfish/v1/Systems/1/Storage"},
			IndicatorLED:       "Lit",
			Links: Links{
				Chassis:   []Link{{OdataID: "/redfish/v1/Chassis/fake-chassis"}},
				ManagedBy: []Link{{OdataID: "/redfish/v1/Managers/1"}},
			},
			Actions: Actions{
				Reset: struct {
					Target                   string   `json:"target"`
					ResetTypeAllowableValues []string `json:"ResetType@Redfish.AllowableValues"`
				}{
					Target:                   "/redfish/v1/Systems/1/Actions/ComputerSystem.Reset",
					ResetTypeAllowableValues: []string{"On", "ForceOff", "GracefulShutdown", "GracefulRestart", "ForceRestart", "Nmi", "ForceOn"},
				},
			},
			OdataContext:     "/redfish/v1/$metadata#ComputerSystem.ComputerSystem",
			OdataID:          "/redfish/v1/Systems/1",
			RedfishCopyright: "Copyright 2014-2016 Distributed Management Task Force, Inc. (DMTF). For the full DMTF copyright policy, see http://www.dmtf.org/about/policies/copyright.",
		}
	} else {
		response = SystemsResponse{
			OdataID: "/redfish/v1/Systems",
			Name:    "System Collection",
		}
	}

	WritePrettyJSON(w, response)
}
