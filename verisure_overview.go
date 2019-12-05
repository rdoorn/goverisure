package goverisure

import (
	"encoding/json"
	"time"
)

type ArmState struct {
	Cid        string    `json:"cid"`
	StatusType string    `json:"statusType"`
	Date       time.Time `json:"date"`
	ChangedVia string    `json:"changedVia"`
	Name       string    `json:"name"`
}

type ControlPlug struct {
	DeviceID     string `json:"deviceId"`
	DeviceLabel  string `json:"deviceLabel"`
	Area         string `json:"area"`
	Profile      string `json:"profile"`
	CurrentState string `json:"currentState"`
	PendingState string `json:"pendingState"`
}

type SmartPlug struct {
	Icon         string `json:"icon"`
	IsHazardous  bool   `json:"isHazardous"`
	DeviceLabel  string `json:"deviceLabel"`
	Area         string `json:"area"`
	CurrentState string `json:"currentState"`
	PendingState string `json:"pendingState"`
}

type ClimateValue struct {
	DeviceLabel string    `json:"deviceLabel"`
	DeviceArea  string    `json:"deviceArea"`
	DeviceType  string    `json:"deviceType"`
	Temperature float64   `json:"temperature"`
	Time        time.Time `json:"time"`
	Humidity    float64   `json:"humidity,omitempty"`
}

type InstallationError struct {
	Cid                    string    `json:"cid"`
	Zone                   string    `json:"zone"`
	StatusType             string    `json:"statusType"`
	State                  bool      `json:"state"`
	Date                   time.Time `json:"date"`
	Name                   string    `json:"name"`
	DeviceLabel            string    `json:"deviceLabel"`
	DeviceType             string    `json:"deviceType"`
	NotificationDeviceType string    `json:"notificationDeviceType"`
	Area                   string    `json:"area"`
}

type DoorWindow struct {
	ReportState      bool               `json:"reportState"`
	DoorWindowDevice []DoorWindowDevice `json:"doorWindowDevice"`
}

type DoorWindowDevice struct {
	DeviceLabel string    `json:"deviceLabel"`
	Area        string    `json:"area"`
	State       string    `json:"state"`
	Wired       bool      `json:"wired"`
	ReportTime  time.Time `json:"reportTime"`
}

type Overview struct {
	/*AccountPermissions struct {
		 AccountPermissionsHash string `json:"accountPermissionsHash"`
	 } `json:"accountPermissions"`*/
	ArmState           ArmState      `json:"armState"`
	ArmstateCompatible bool          `json:"armstateCompatible"`
	ControlPlugs       []ControlPlug `json:"controlPlugs"`
	SmartPlugs         []SmartPlug   `json:"smartPlugs"`
	//DoorLockStatusList []interface{} `json:"doorLockStatusList"`
	TotalSmsCount         int                 `json:"totalSmsCount"`
	ClimateValues         []ClimateValue      `json:"climateValues"`
	InstallationErrorList []InstallationError `json:"installationErrorList"`
	PendingChanges        int                 `json:"pendingChanges"`
	EthernetModeActive    bool                `json:"ethernetModeActive"`
	EthernetConnectedNow  bool                `json:"ethernetConnectedNow"`
	//	HeatPumps            []interface{} `json:"heatPumps"`
	/*SmartCameras         []struct {
		MotionDetectorActive         bool   `json:"motionDetectorActive"`
		Area                         string `json:"area"`
		DeviceLabel                  string `json:"deviceLabel"`
		NumberOfNotViewedImageSeries int    `json:"numberOfNotViewedImageSeries"`
	} `json:"smartCameras"`
	LatestEthernetStatus struct {
		LatestEthernetTestResult bool      `json:"latestEthernetTestResult"`
		TestDate                 time.Time `json:"testDate"`
		ProtectedArea            string    `json:"protectedArea"`
		DeviceLabel              string    `json:"deviceLabel"`
	} `json:"latestEthernetStatus"`
	CustomerImageCameras []struct {
		MotionDetectorMode           string `json:"motionDetectorMode"`
		Area                         string `json:"area"`
		DeviceLabel                  string `json:"deviceLabel"`
		NumberOfNotViewedImageSeries int    `json:"numberOfNotViewedImageSeries"`
		ImageResolution              string `json:"imageResolution"`
		AccelerometerMode            string `json:"accelerometerMode"`
		Capability                   string `json:"capability"`
	} `json:"customerImageCameras"`*/
	/*BatteryProcess struct {
		Active bool `json:"active"`
	} `json:"batteryProcess"`*/
	/*UserTracking struct {
		InstallationStatus string `json:"installationStatus"`
	} `json:"userTracking"`
	EventCounts []interface{} `json:"eventCounts"`*/
	DoorWindow DoorWindow `json:"doorWindow"`
}

func (n ClimateValue) String() string {
	b, err := json.Marshal(n)
	if err != nil {
		return err.Error()
	}

	return string(b)
}

func (n DoorWindowDevice) String() string {
	b, err := json.Marshal(n)
	if err != nil {
		return err.Error()
	}

	return string(b)
}

func (n ArmState) String() string {
	b, err := json.Marshal(n)
	if err != nil {
		return err.Error()
	}

	return string(b)
}

func (h *Handler) Overview(installationID string) (Overview, error) {
	o := Overview{}
	data, err := h.Get("GET", urlGetOverview, h.apiURL, installationID)

	err = json.Unmarshal(data, &o)
	if err != nil {
		return Overview{}, err
	}

	return o, nil
}
