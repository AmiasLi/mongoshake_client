package models

type Sentinel struct {
	OplogDump      int  `json:"OplogDump"`
	DuplicatedDump bool `json:"DuplicatedDump"`
	Pause          bool `json:"Pause"`
	TPS            int  `json:"TPS"`
	TargetDelay    int  `json:"TargetDelay"`
	ExitPoint      int  `json:"ExitPoint"`
	Shutdown       bool `json:"Shutdown"`
}
