package controllers

import (
	"testing"
	"flag"
	"os"
)

func TestMain(m *testing.M) {
	flag.Set("alsologtostderr", "true")
	//flag.Set("log_dir", "/tmp")
	flag.Set("v", "1")
	flag.Parse()

	ret := m.Run()
	os.Exit(ret)
}
func TestClusterExpand1(t *testing.T) {
	if ok,err:=ClusterExpand("10.109.252.172","10.109.252.180",
		"SWMTKN-1-1ad5cg3o98z1v1tbu2clg5h5u0753m3dh6ml7xg8b1llnrwiu3-3otyvh8qnoj8fxjni1dutmoqu");!ok{
		t.Error(err)
	}else{
		t.Log(ok)
	}
}

