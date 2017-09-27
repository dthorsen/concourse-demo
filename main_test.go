package main

import "testing"

func TestUnitMultiply(t *testing.T) {
	if Multiply(8, 3) != 24 {
		t.Fail()
	}
}

func TestIntegrationGetGithubZen(t *testing.T) {
	if GetGithubZen() == "" {
		t.Fail()
	}
}
