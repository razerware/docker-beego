package models

type ElasticInfo struct {
	UpperLimit int `json:"upper_limit"`
	LowerLimit int `json:"lower_limit"`
	Step       int `json:"step"`
	CpuLower   int `json:"cpu_lower"`
	CpuUpper   int `json:"cpu_upper"`
	MemLower   int `json:"mem_lower"`
	MemUpper   int `json:"mem_upper"`
}
