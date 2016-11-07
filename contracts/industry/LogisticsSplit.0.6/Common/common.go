package Common;
import (
    "encoding/json"
)
const BLSTATEKEY string = "BLSTATEKEY" 
const CONTSTATEKEY string = "CONTSTATEKEY"

const BLSTATE string = "_STATE"
// const CONTHIST string = "_HIST" 

const MYVERSION string = "2.0.0"


type BLContractState struct {
    Version      string                        `json:"version"`
    ContainerCC  string                        `json:"containercc"`
    ComplianceCC string                        `json:"compliancecc"`
}   

type ContContractState struct {
    Version         string                        `json:"version"`
    ComplianceCC    string                        `json:"compliancecc"`
} 

type Geolocation struct {
    Latitude    float64 `json:"latitude,omitempty"`
    Longitude   float64 `json:"longitude,omitempty"`
}

type Airquality struct {
    Oxygen          float64 `json:"oxygen,omitempty"`
    Carbondioxide   float64 `json:"carbondioxide,omitempty"`
    Ethylene        float64 `json:"ethylene,omitempty"`
}


// This is  optional. It stands for the 'acceptable range', say 1 degree of lat and long
// at which the container should be, before it is considered 'arrived' at 'Notified Party' location'
// If not sent in, some default could be assumed (say 1 degree or )
type NotifyRange struct {
    LatRange        float64 `json:"latrange,omitempty"`
    LongRange       float64 `json:"longrange,omitempty"`
}

// This is a logistics contract, written in the context of shipping. It tracks the progress of a Bill of Lading 
// and associated containers, and raises alerts in case of violations in expected conditions

// Assumption 1. Bill of Lading is sacrosanct - Freight forwarders may issue intermediary freight bills, but
// the original B/L is the document we trackend to end. Similarly a 'Corrected B/L' scenario is not considered

// Assumption 2. A Bill of Lading can have multiple containers attached to it. We are, for simplicity, assuming that
// the same transit rules in terms of allowed ranges in temperature, humidity etc. apply across the B/L - i.e. 
// applies to all containers attached to a Bill of Lading

// Assumption 3. A shipment may switch from one container to another in transit, for various reasons. We are assuming,
// for simplicity, that the same containers are used for end to end transit.

// Initial registration of the Bill of Lading. Sets out the constrains for B/L data and the Notification details
type BillOfLadingRegistration struct {
    BLNo                 string                  `json:"blno"` 
    ContainerNos         string                  `json:"containernos"`    // Comma separated container numbers - keep json simple    
    Hazmat               bool                    `json:"hazmat,omitempty"`     // shipment hazardous ?
    MinTemperature       float64                 `json:"mintemperature,omitempty"` //split range to min and max: Jeff's input
    MaxTemperature       float64                 `json:"maxtemperature,omitempty"` 
    MinHumidity          float64                 `json:"minhumidity,omitempty"` //split range to min and max: Jeff's input
    MaxHumidity          float64                 `json:"maxhumidity,omitempty"`
    MinLight             float64                 `json:"minlight,omitempty"` //split range to min and max: Jeff's input
    MaxLight             float64                 `json:"maxlight,omitempty"` 
    MinAcceleration      float64                 `json:"minacceleration,omitempty"` //split range to min and max: Jeff's input
    MaxAcceleration      float64                 `json:"maxacceleration,omitempty"`
 //NotifyLocations      *[]Geolocation            `json:"notifylocations,omitempty"` // No implementation right now
 //NotifyRange          *NotifyRange              `json:"notifyrange,omitempty"`     // To be integrated when shipping part gets sorted out
    TransitComplete      bool                    `json:"transitcomplete,omitempty"`
    Timestamp            string                  `json:"timestamp,omitempty"`
}
//Structure for logistics data at the container level
type ContainerLogistics struct {
    ContainerNo         string                         `json:"containerno"`    
    BLNo                string                         `json:"blno,omitempty"`    
    Location            Geolocation                    `json:"location,omitempty"`       // current asset location
    Carrier             string                         `json:"carrier,omitempty"`        // the name of the carrier
    Timestamp           string                         `json:"timestamp"`          
    Temperature         float64                        `json:"temperature,omitempty"`    // celcius
    Humidity            float64                        `json:"humidity,omitempty"` // percent
    Light               float64                        `json:"light,omitempty"` // lumen
    Acceleration        float64                        `json:"acceleration,omitempty"`
    DoorClosed          bool                           `json:"doorclosed,omitempty"`
    AirQuality          Airquality                     `json:"airquality,omitempty"`
    Extra               json.RawMessage                `json:"extra,omitempty"`  
    AlertRecord         string                         `json:"alerts,omitempty"`  
    TransitComplete     bool                           `json:"transitcomplete,omitempty"`
}

// Compliance record structure
type ComplianceState struct {
    BLNo                 string                     `json:"blno"` 
    Type                 string                     `json:"type"` // Default: DEFTYPE
    Compliance           bool                       `json:"compliance"`
    AssetAlerts          map[string]string          `json:"assetalerts"`
    Active               bool                       `json:"active,omitempty"`
    Timestamp            string                      `json:"timestamp"`
}


type Variation string

const (
    Normal Variation ="normal"
    Above ="above"
    Below = "below" 
) 
  

// These are common alerts reported by sensor. Example Tetis. 
// http://www.starcomsystems.com/download/Tetis_ENG.pdf 

type Alerts struct {
     TempAlert      Variation `json:"tempalert,omitempty"`
     HumAlert       Variation `json:"humalert,omitempty"`
     LightAlert     Variation `json:"lightalert,omitempty"` 
     AccAlert       Variation `json:"accalert,omitempty"`
     DoorAlert      bool      `json:"dooralert,omitempty"`
}


// This is a logistics contract, written in the context of shipping. It tracks the progress of a Bill of Lading 
// and associated containers, and raises alerts in case of violations in expected conditions

// Assumption 1. Bill of Lading is sacrosanct - Freight forwarders may issue intermediary freight bills, but
// the original B/L is the document we trackend to end. Similarly a 'Corrected B/L' scenario is not considered

// Assumption 2. A Bill of Lading can have multiple containers attached to it. We are, for simplicity, assuming that
// the same transit rules in terms of allowed ranges in temperature, humidity etc. apply across the B/L - i.e. 
// applies to all containers attached to a Bill of Lading

// Assumption 3. A shipment may switch from one container to another in transit, for various reasons. We are assuming,
// for simplicity, that the same containers are used for end to end transit.


