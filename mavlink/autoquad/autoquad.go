package autoquad

import (
	"github.com/SpaceLeap/go-mavlink/mavlink"
	"github.com/SpaceLeap/go-mavlink/mavlink/common"
)

const (
	PROTOCOL_NAME    = "autoquad"
	PROTOCOL_VERSION = "3"
	PROTOCOL_INCLUDE = common.PROTOCOL_NAME
)

func init() {
	mavlink.ProtocolName = PROTOCOL_NAME
	mavlink.ProtocolVersion = PROTOCOL_VERSION

	mavlink.MessageFactory[150] = func() mavlink.Message { return new(AqTelemetryF) }
}

////////////////////////////////////////////////////////////////////////////////
// Enums
////////////////////////////////////////////////////////////////////////////////

// AUTOQUAD_NAV_STATUS: Available operating modes/statuses for AutoQuad flight controller.
// 				Bitmask up to 32 bits. Low side bits for base modes, high side for
// 				additional active features/modifiers/constraints.
const (
	AQ_NAV_STATUS_INIT            = 0 // System is initializing
	AQ_NAV_STATUS_STANDBY         = 0 // System is standing by, not active
	AQ_NAV_STATUS_MANUAL          = 0 // Stabilized, under full manual control
	AQ_NAV_STATUS_ALTHOLD         = 0 // Altitude hold engaged
	AQ_NAV_STATUS_POSHOLD         = 0 // Position hold engaged
	AQ_NAV_STATUS_DVH             = 0 // Dynamic Velocity Hold is active
	AQ_NAV_STATUS_MISSION         = 0 // Autonomous mission execution mode
	AQ_NAV_STATUS_FAILSAFE        = 0 // System is in failsafe recovery mode
	AQ_NAV_STATUS_RTH             = 0 // Automatic Return to Home is active
	AQ_NAV_STATUS_HF_LOCKED       = 0 // Heading-Free locked mode active
	AQ_NAV_STATUS_HF_DYNAMIC      = 0 // Heading-Free dynamic mode active
	AQ_NAV_STATUS_CEILING         = 0 // Ceiling altitude is set
	AQ_NAV_STATUS_CEILING_REACHED = 0 // Craft is at ceiling altitude
)

// MAV_CMD:
const (
	MAV_CMD_AQ_TELEMETRY       = 2 // Start/stop AutoQuad telemetry values stream.
	MAV_CMD_AQ_FOLLOW          = 3 // Command AutoQuad to go to a particular place at a set speed.
	MAV_CMD_AQ_REQUEST_VERSION = 4 // Request AutoQuad firmware version number.
)

////////////////////////////////////////////////////////////////////////////////
// Messages
////////////////////////////////////////////////////////////////////////////////

// Sends up to 20 raw float values.
type AqTelemetryF struct {
	Value20 float32 // value20
	Value19 float32 // value19
	Value18 float32 // value18
	Value17 float32 // value17
	Value16 float32 // value16
	Value15 float32 // value15
	Value14 float32 // value14
	Value13 float32 // value13
	Value12 float32 // value12
	Value11 float32 // value11
	Value10 float32 // value10
	Value9  float32 // value9
	Value8  float32 // value8
	Value7  float32 // value7
	Value6  float32 // value6
	Value5  float32 // value5
	Value4  float32 // value4
	Value3  float32 // value3
	Value2  float32 // value2
	Value1  float32 // value1
	Index   uint16  // Index of message
}

func (self *AqTelemetryF) TypeID() uint8 {
	return 150
}

func (self *AqTelemetryF) TypeName() string {
	return "AQ_TELEMETRY_F"
}

func (self *AqTelemetryF) TypeSize() uint8 {
	return 82
}

func (self *AqTelemetryF) TypeCRCExtra() uint8 {
	return 21
}
