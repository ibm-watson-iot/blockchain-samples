package main

type Device struct {
    DeviceID                    *string                    `json:"deviceid,omitempty"`    
    MinimumUsageCost            *float64                   `json:"minimumusagecost,omitempty"`    
    MinimumUsageTime            *int32                    `json:"minimumusagetime,omitempty"` 
    OvertimeUsageCost           *float64                   `json:"overtimeusagecost,omitempty"`    
    OvertimeUsageTime           *int32                    `json:"overtimeusagetime,omitempty"`       
    Available                   *bool                      `json:"available,omitempty"`        
}

type Usage struct {
    DeviceID                    string                         `json:"deviceid,omitempty"`    
    StartTime                   string                         `json:"starttime,omitempty"`    
    EndTime                     string                         `json:"endtime,omitempty"`       // current asset location
    Duration                    int64                         `json:"duration,omitempty"`        // the name of the carrier
    UsageCost                   float64                         `json:"usagecost,omitempty"`          
    ActualEndtime               string                        `json:"actualendtime,omitempty"`    // celcius
    OvertimeCost                float64                        `json:"overtimecost,omitempty"` // percent
    TotalCost                   float64                        `json:"totalcost,omitempty"` // percent
}



const CONTSTATEKEY string = "STATE"
const DEVICESKEY string = "DEVICES"
const USAGEKEY string = "USAGE"
const USAGEHIST string = "USAGEHIST"
const ALERTKEY string = "ALERT"
const LISTKEY  string = "DEVLIST"

const MAXHIST int = 10

const BufferTime int = 2

type AlertLevels string

const (
    Available AlertLevels = "available"
    Confirm  ="confirm"
    HalfTime = "half-time"
    Warning ="warning"
    Overtime = "overtime" 
) 
  

const MYVERSION string = "1.0.0"


type ContractState struct {
    Version         string                        `json:"version"`
} 

//Usage History
type UsageHistory struct {
	History []string `json:"history"`
}
//Device List
type DevList struct {
	Devices []string `json:"devices"`
}
var contractState = ContractState{MYVERSION}

