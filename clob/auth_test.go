package clob

import "testing"

func TestBuildHMACSignature(t *testing.T) {
	got, err := BuildHMACSignature("c2VjcmV0", 1713398400, "GET", "/data/orders", nil)
	if err != nil {
		t.Fatal(err)
	}
	const want = "Fhwg4QPTdvtjCN-5ibPvlCGaHItdnOkdhZdMqxx8gHU="
	if got != want {
		t.Fatalf("signature = %q, want %q", got, want)
	}
}
