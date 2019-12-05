package goverisure

import (
	"encoding/json"
	"time"
)

/*
{"accountPermissions":{"accountPermissionsHash":"d1Xh4Wcnl5ew4kIYMd3qil3Q5Vs="},"armState":{"statusType":"DISARMED","date":"2019-12-05T06:38:24.000Z","changedVia":"HOMEPAD"},"armstateCompatible":true,"controlPlugs":[{"deviceId":"34747937","deviceLabel":"2AB5 EAA9","area":"Game Room","profile":"OTHER","currentState":"ON","pendingState":"NONE"},{"deviceId":"34912196","deviceLabel":"2ACD HK57","area":"Kerst rechts","profile":"OTHER","currentState":"ON","pendingState":"NONE"},{"deviceId":"34912197","deviceLabel":"2ACD HKD6","area":"Sonos achter","profile":"OTHER","currentState":"ON","pendingState":"NONE"},{"deviceId":"34912200","deviceLabel":"2ACD HL5B","area":"Tv meubel","profile":"OTHER","currentState":"ON","pendingState":"NONE"},{"deviceId":"34912408","deviceLabel":"2ACD K83Q","area":"Kerst links","profile":"OTHER","currentState":"ON","pendingState":"NONE"},{"deviceId":"35977182","deviceLabel":"2ALH HRJH","area":"Sonos sub","profile":"OTHER","currentState":"ON","pendingState":"NONE"},{"deviceId":"35977186","deviceLabel":"2ALH HSKM","area":"Computer eveline","profile":"OTHER","currentState":"ON","pendingState":"NONE"}],"smartPlugs":[{"icon":"OTHER","isHazardous":false,"deviceLabel":"2AB5 EAA9","area":"Game Room","currentState":"ON","pendingState":"NONE"},{"icon":"OTHER","isHazardous":false,"deviceLabel":"2ACD HK57","area":"Kerst rechts","currentState":"ON","pendingState":"NONE"},{"icon":"OTHER","isHazardous":false,"deviceLabel":"2ACD HKD6","area":"Sonos achter","currentState":"ON","pendingState":"NONE"},{"icon":"OTHER","isHazardous":false,"deviceLabel":"2ACD HL5B","area":"Tv meubel","currentState":"ON","pendingState":"NONE"},{"icon":"OTHER","isHazardous":false,"deviceLabel":"2ACD K83Q","area":"Kerst links","currentState":"ON","pendingState":"NONE"},{"icon":"OTHER","isHazardous":false,"deviceLabel":"2ALH HRJH","area":"Sonos sub","currentState":"ON","pendingState":"NONE"},{"icon":"OTHER","isHazardous":false,"deviceLabel":"2ALH HSKM","area":"Computer eveline","currentState":"ON","pendingState":"NONE"}],"doorLockStatusList":[],"totalSmsCount":0,"climateValues":[{"deviceLabel":"2AB7 H5YA","deviceArea":"slaapkamer","deviceType":"HOMEPAD1","temperature":19.1,"time":"2019-12-05T13:18:24.000Z"},{"deviceLabel":"2ADC 3EA7","deviceArea":"overloop","deviceType":"SMOKE2","temperature":19.8,"humidity":56.0,"time":"2019-12-05T11:19:57.000Z"},{"deviceLabel":"2ADC 3GAF","deviceArea":"Kantoor ruimte","deviceType":"SMOKE2","temperature":19.4,"humidity":56.0,"time":"2019-12-05T11:22:26.000Z"},{"deviceLabel":"2AFP TJ67","deviceArea":"Tuin pad","deviceType":"SMARTCAMERA1","temperature":26.3,"time":"2017-08-12T18:15:42.000Z"},{"deviceLabel":"2AHG DK3X","deviceArea":"Wasmachine","deviceType":"WATER1","temperature":18.6,"time":"2019-12-05T13:09:43.000Z"},{"deviceLabel":"2S5R NGU2","deviceArea":"keuken","deviceType":"SIREN1","temperature":20.5,"time":"2019-12-05T13:16:50.000Z"},{"deviceLabel":"2W2P 5B9M","deviceArea":"Kantoor ruimte","deviceType":"PIR2","temperature":16.0,"time":"2019-12-05T13:14:01.000Z"}],"installationErrorList":[{"cid":"00791980","zone":"018","statusType":"SUPV","state":true,"date":"2019-07-12T19:11:57.000Z","name":"US018","deviceLabel":"2ALH HRJH","deviceType":"CONTROLPLUG1","notificationDeviceType":"CONTROLPLUG","area":"Sonos sub"},{"cid":"00791980","zone":"017","statusType":"SUPV","state":true,"date":"2017-08-12T20:38:58.000Z","name":"US017","deviceLabel":"2AFP TJ67","deviceType":"SMARTCAMERA1","notificationDeviceType":"SMARTCAMERA","area":"Tuin pad"},{"cid":"00791980","zone":"014","statusType":"SUPV","state":true,"date":"2019-07-12T19:11:57.000Z","name":"US014","deviceLabel":"2ACD HKD6","deviceType":"CONTROLPLUG1","notificationDeviceType":"CONTROLPLUG","area":"Sonos achter"},{"cid":"00791980","zone":"013","statusType":"SUPV","state":true,"date":"2019-01-26T11:41:51.000Z","name":"US013","deviceLabel":"2ACD HK57","deviceType":"CONTROLPLUG1","notificationDeviceType":"CONTROLPLUG","area":"Kerst rechts"},{"cid":"00791980","zone":"016","statusType":"SUPV","state":true,"date":"2019-01-26T11:41:51.000Z","name":"US016","deviceLabel":"2ACD K83Q","deviceType":"CONTROLPLUG1","notificationDeviceType":"CONTROLPLUG","area":"Kerst links"}],"pendingChanges":0,"ethernetModeActive":true,"ethernetConnectedNow":true,"heatPumps":[],"smartCameras":[{"motionDetectorActive":true,"area":"Tuin pad","deviceLabel":"2AFP TJ67","numberOfNotViewedImageSeries":0}],"latestEthernetStatus":{"latestEthernetTestResult":true,"testDate":"2019-12-04T19:21:20.000Z","protectedArea":"meterkast","deviceLabel":"2676 DZFM"},"customerImageCameras":[{"motionDetectorMode":"ACTIVE","area":"Tuin pad","deviceLabel":"2AFP TJ67","numberOfNotViewedImageSeries":0,"imageResolution":"HIGHER","accelerometerMode":"INACTIVE","capability":"USER_MONITORED_CUSTOMER_IMAGE_CAMERA"}],"batteryProcess":{"active":false},"userTracking":{"installationStatus":"NOT_CONFIGURED"},"eventCounts":[],"doorWindow":{"reportState":true,"doorWindowDevice":[{"deviceLabel":"2JCM MG6J","area":"voordeur","state":"CLOSE","wired":false,"reportTime":"2019-12-05T12:29:30.000Z"},{"deviceLabel":"2JCM MSQC","area":"Tuindeun rechts","state":"CLOSE","wired":false,"reportTime":"2017-10-25T15:27:18.000Z"},{"deviceLabel":"2JCM MFEX","area":"Tuindeur links","state":"CLOSE","wired":false,"reportTime":"2019-12-01T07:49:14.000Z"},{"deviceLabel":"2JHY LAGM","area":"Babykamer raam","state":"CLOSE","wired":false,"reportTime":"2019-09-30T08:09:33.000Z"},{"deviceLabel":"2JJ6 VLF2","area":"Woonkamer raam","state":"CLOSE","wired":false,"reportTime":"2017-10-26T11:40:08.000Z"},{"deviceLabel":"2JJ7 SDUR","area":"Badkamer raam","state":"CLOSE","wired":false,"reportTime":"2019-04-23T09:10:41.000Z"},{"deviceLabel":"2JJ6 WLB5","area":"Kantoor deur","state":"CLOSE","wired":false,"reportTime":"2019-08-30T07:55:46.000Z"},{"deviceLabel":"2JJ3 GJMK","area":"Raam eethoek","state":"CLOSE","wired":false,"reportTime":"2018-06-01T15:20:08.000Z"},{"deviceLabel":"2JJ4 XRGR","area":"Keuken raam","state":"CLOSE","wired":false,"reportTime":"2019-08-22T07:32:18.000Z"}]}}
*/

type ArmState struct {
	StatusType string    `json:"statusType"`
	Date       time.Time `json:"date"`
	ChangedVia string    `json:"changedVia"`
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

func (h *Handler) Overview(installationID string) ([]Installation, error) {
	var i []Installation
	data, err := h.Get("GET", urlGetInstallations, h.apiURL, installationID)

	err = json.Unmarshal(data, i)
	if err != nil {
		return nil, err
	}

	return i, nil
}
