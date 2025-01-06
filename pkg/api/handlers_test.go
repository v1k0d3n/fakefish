package api

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"strings"
	"testing"

	"github.com/v1k0d3n/fakefish/pkg/config"
)

func TestRootHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/redfish/v1/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(RootHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{
  "@odata.type": "#ServiceRoot.v1_5_0.ServiceRoot",
  "Id": "FakeFishService",
  "Name": "FakeFish Service",
  "RedfishVersion": "1.5.0",
  "UUID": "not-that-production-ready",
  "Systems": {
    "@odata.id": "/redfish/v1/Systems"
  },
  "Managers": {
    "@odata.id": "/redfish/v1/Managers"
  },
  "@odata.id": "/redfish/v1/",
  "@Redfish.Copyright": "Copyright 2014-2016 Distributed Management Task Force, Inc. (DMTF). For the full DMTF copyright policy, see http://www.dmtf.org/about/policies/copyright."
}`
	if strings.TrimSpace(rr.Body.String()) != strings.TrimSpace(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestSystemsHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/redfish/v1/Systems/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(SystemsHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{
  "@odata.type": "#ComputerSystem.v1_1_0.ComputerSystem",
  "Id": "1",
  "Name": "fake-system",
  "UUID": "1",
  "Manufacturer": "FakeFish",
  "Status": {
    "State": "Enabled",
    "Health": "OK",
    "HealthRollUp": "OK"
  },
  "PowerState": "On",
  "Boot": {
    "BootSourceOverrideEnabled": "Continuous",
    "BootSourceOverrideTarget": "Hdd",
    "BootSourceOverrideTarget@Redfish.AllowableValues": [
      "Pxe",
      "Cd",
      "Hdd"
    ],
    "BootSourceOverrideMode": "UEFI",
    "UefiTargetBootSourceOverride": "/0x31/0x33/0x01/0x01"
  },
  "ProcessorSummary": {
    "Count": 8,
    "Status": {
      "State": "Enabled",
      "Health": "OK",
      "HealthRollUp": "OK"
    }
  },
  "MemorySummary": {
    "TotalSystemMemoryGiB": 15,
    "Status": {
      "State": "Enabled",
      "Health": "OK",
      "HealthRollUp": "OK"
    }
  },
  "Bios": {
    "@odata.id": "/redfish/v1/Systems/1/BIOS"
  },
  "Processors": {
    "@odata.id": "/redfish/v1/Systems/1/Processors"
  },
  "Memory": {
    "@odata.id": "/redfish/v1/Systems/1/Memory"
  },
  "EthernetInterfaces": {
    "@odata.id": "/redfish/v1/Systems/1/EthernetInterfaces"
  },
  "SimpleStorage": {
    "@odata.id": "/redfish/v1/Systems/1/SimpleStorage"
  },
  "Storage": {
    "@odata.id": "/redfish/v1/Systems/1/Storage"
  },
  "IndicatorLED": "Lit",
  "Links": {
    "Chassis": [
      {
        "@odata.id": "/redfish/v1/Chassis/fake-chassis"
      }
    ],
    "ManagedBy": [
      {
        "@odata.id": "/redfish/v1/Managers/1"
      }
    ]
  },
  "Actions": {
    "#ComputerSystem.Reset": {
      "target": "/redfish/v1/Systems/1/Actions/ComputerSystem.Reset",
      "ResetType@Redfish.AllowableValues": [
        "On",
        "ForceOff",
        "GracefulShutdown",
        "GracefulRestart",
        "ForceRestart",
        "Nmi",
        "ForceOn"
      ]
    }
  },
  "@odata.context": "/redfish/v1/$metadata#ComputerSystem.ComputerSystem",
  "@odata.id": "/redfish/v1/Systems/1",
  "@Redfish.Copyright": "Copyright 2014-2016 Distributed Management Task Force, Inc. (DMTF). For the full DMTF copyright policy, see http://www.dmtf.org/about/policies/copyright."
}`
	if strings.TrimSpace(rr.Body.String()) != strings.TrimSpace(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestManagersHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/redfish/v1/Managers/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ManagersHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{
  "@odata.type": "#Manager.v1_1_0.Manager",
  "Id": "1",
  "Name": "fake-manager",
  "ManagerType": "BMC",
  "UUID": "1",
  "Model": "FakeManagerModel",
  "FirmwareVersion": "1.0.0",
  "Status": {
    "State": "Enabled",
    "Health": "OK",
    "HealthRollUp": "OK"
  },
  "EthernetInterfaces": {
    "@odata.id": "/redfish/v1/Managers/1/EthernetInterfaces"
  },
  "Links": {
    "Chassis": [
      {
        "@odata.id": "/redfish/v1/Chassis/fake-chassis"
      }
    ],
    "ManagedBy": [
      {
        "@odata.id": "/redfish/v1/Managers/1"
      }
    ]
  },
  "@odata.context": "/redfish/v1/$metadata#Manager.Manager",
  "@odata.id": "/redfish/v1/Managers/1",
  "@Redfish.Copyright": "Copyright 2014-2016 Distributed Management Task Force, Inc. (DMTF). For the full DMTF copyright policy, see http://www.dmtf.org/about/policies/copyright."
}`
	if strings.TrimSpace(rr.Body.String()) != strings.TrimSpace(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestBootFromCdOnceHandler(t *testing.T) {
	scriptPath, _ := filepath.Abs("../../scripts/testing/bootfromcdonce.sh")
	scriptConfig := &config.Config{
		Commands: map[string]string{
			"bootFromCdOnce": scriptPath,
		},
		Server: config.ServerConfig{
			Port: 8080,
		},
		TLS: config.TLSConfig{
			Enabled:    false,
			CaCert:     "tls/ca.crt",
			ClientCert: "tls/client.crt",
			ClientKey:  "tls/client.key",
		},
	}
	InitHandlers(scriptConfig)

	req, err := http.NewRequest("GET", "/bootfromcdonce", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(BootFromCdOnceHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "Boot from CD once executed successfully"
	if strings.TrimSpace(rr.Body.String()) != strings.TrimSpace(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestMountCdHandler(t *testing.T) {
	scriptPath, _ := filepath.Abs("../../scripts/testing/mountcd.sh")
	scriptConfig := &config.Config{
		Commands: map[string]string{
			"mountCd": scriptPath,
		},
		Server: config.ServerConfig{
			Port: 8080,
		},
		TLS: config.TLSConfig{
			Enabled:    false,
			CaCert:     "tls/ca.crt",
			ClientCert: "tls/client.crt",
			ClientKey:  "tls/client.key",
		},
	}
	InitHandlers(scriptConfig)

	reqBody := bytes.NewBufferString(`{"Image": "http://example.com/path/to/iso"}`)
	req, err := http.NewRequest("POST", "/redfish/v1/Managers/1/VirtualMedia/CD/Actions/VirtualMedia.InsertMedia", reqBody)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(MountCdHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "Mount CD executed successfully"
	if strings.TrimSpace(rr.Body.String()) != strings.TrimSpace(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestPowerOffHandler(t *testing.T) {
	scriptPath, _ := filepath.Abs("../../scripts/testing/poweroff.sh")
	scriptConfig := &config.Config{
		Commands: map[string]string{
			"powerOff": scriptPath,
		},
		Server: config.ServerConfig{
			Port: 8080,
		},
		TLS: config.TLSConfig{
			Enabled:    false,
			CaCert:     "tls/ca.crt",
			ClientCert: "tls/client.crt",
			ClientKey:  "tls/client.key",
		},
	}
	InitHandlers(scriptConfig)

	req, err := http.NewRequest("POST", "/redfish/v1/Managers/1/Actions/Manager.Reset", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(PowerOffHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "Power off executed successfully"
	if strings.TrimSpace(rr.Body.String()) != strings.TrimSpace(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestPowerOnHandler(t *testing.T) {
	scriptPath, _ := filepath.Abs("../../scripts/testing/poweron.sh")
	scriptConfig := &config.Config{
		Commands: map[string]string{
			"powerOn": scriptPath,
		},
		Server: config.ServerConfig{
			Port: 8080,
		},
		TLS: config.TLSConfig{
			Enabled:    false,
			CaCert:     "tls/ca.crt",
			ClientCert: "tls/client.crt",
			ClientKey:  "tls/client.key",
		},
	}
	InitHandlers(scriptConfig)

	req, err := http.NewRequest("POST", "/redfish/v1/Managers/1/Actions/Manager.Start", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(PowerOnHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "Power on executed successfully"
	if strings.TrimSpace(rr.Body.String()) != strings.TrimSpace(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestUnmountCdHandler(t *testing.T) {
	scriptPath, _ := filepath.Abs("../../scripts/testing/unmountcd.sh")
	scriptConfig := &config.Config{
		Commands: map[string]string{
			"unmountCd": scriptPath,
		},
		Server: config.ServerConfig{
			Port: 8080,
		},
		TLS: config.TLSConfig{
			Enabled:    false,
			CaCert:     "tls/ca.crt",
			ClientCert: "tls/client.crt",
			ClientKey:  "tls/client.key",
		},
	}
	InitHandlers(scriptConfig)

	req, err := http.NewRequest("POST", "/redfish/v1/Managers/1/VirtualMedia/CD/Actions/VirtualMedia.EjectMedia", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(UnmountCdHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "Unmount CD executed successfully"
	if strings.TrimSpace(rr.Body.String()) != strings.TrimSpace(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
