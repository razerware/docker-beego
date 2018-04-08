package elbController

type ElasticInfo struct {
	UpperLimit int `json:"UpperLimit"`
	LowerLimit int `json:"LowerLimit"`
	Step int `json:"Step"`
	CpuLower int `json:"CpuLower"`
	CpuUpper int `json:"CpuUpper"`
	MemLower int `json:"MemLower"`
	MemUpper int `json:"MemUpper"`
}