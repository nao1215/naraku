package model

import "time"

// OS is operating system information.
type OS struct {
	// Name is os Name.
	Name string `json:"name"`
	// Version is os Version.
	Version string `json:"version"`
	// CodeName is os CodeName.
	CodeName string `json:"code_name,omitempty"`
	// Release is os Release.
	Release time.Time `json:"release"`
	// APILevel is os APILevel.
	APILevel []uint `json:"api_level,omitempty"`
}
