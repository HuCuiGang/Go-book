package client

import (
	"ch1/package/series"
	"testing"
)

func TestPackage(t *testing.T){
	t.Log(series.GetFibonacci(6))
}
