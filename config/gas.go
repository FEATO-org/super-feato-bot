package config

import "os"

var SCRIPT_ID string

type GAS struct {
	ScriptID string
}

func init() {
	SCRIPT_ID = os.Getenv("SCRIPT_ID")
}

func NewGAS() GAS {
	return GAS{ScriptID: SCRIPT_ID}
}
