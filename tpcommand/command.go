// Package tpcommand and its subpackages contains low-level types for TP-Link
// device command requests and responses.
//
// Note that certain fields may only be for requests or responses.
package tpcommand

// Checked should is embedded into a Module or Command to allow the error fields
// to be parsed. The struct it is embedded into must not have overlapping
// fields. If it is embedded into a struct with a custom UnmarshalJSON, it
// should also call UnmarshalJSON on Checked.
//
// Based on:
// - com.tplink.kasa_android@2.35.0.1021/com.tplinkra.tpcommon.model.smartlife.iot.Method
type Checked struct {
	ErrCode *int    `json:"err_code,omitempty"`
	ErrMsg  *string `json:"err_msg,omitempty"`
}

// CheckError returns a TP-Link API error, if any.
func (c Checked) CheckError() error {
	var code int
	var message string
	if c.ErrCode != nil {
		code = *c.ErrCode
	}
	if c.ErrMsg != nil {
		message = *c.ErrMsg
	}
	if code >= 0 && !(code == 0 && message != "") {
		return nil
	}
	return Error{
		Code:    ErrorCode(code),
		Message: message,
	}
}

// com.tplink.kasa_android@2.35.0.1021/com.tplinkra.tpcommon.model.smartlife.iot.Method
type Method struct {
	Checked
}

// com.tplink.kasa_android@2.35.0.1021/com.tplinkra.tpcommon.model.smartlife.iot.Module
type Module struct {
	Checked
}
