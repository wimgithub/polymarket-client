package shared

import (
	"encoding/json"
	"testing"
	"time"
)

func TestFlexibleNumericTypes(t *testing.T) {
	var got struct {
		ID    String  `json:"id"`
		Size  Float64 `json:"size"`
		Count Int     `json:"count"`
		U     Uint64  `json:"u"`
	}
	if err := json.Unmarshal([]byte(`{"id":12345,"size":"12.34","count":"7","u":"99"}`), &got); err != nil {
		t.Fatal(err)
	}
	if got.ID.String() != "12345" {
		t.Fatalf("id = %q", got.ID)
	}
	if float64(got.Size) != 12.34 || int(got.Count) != 7 || uint64(got.U) != 99 {
		t.Fatalf("unexpected numerics: %+v", got)
	}
}

func TestFlexibleSliceTypes(t *testing.T) {
	var got struct {
		Outcomes []string     `json:"outcomes"`
		Prices   Float64Slice `json:"prices"`
		IDs      StringSlice  `json:"ids"`
	}
	if err := json.Unmarshal([]byte(`{
		"outcomes":["Yes","No"],
		"prices":"[\"0.42\",0.58]",
		"ids":"1,2,3"
	}`), &got); err != nil {
		t.Fatal(err)
	}
	if len(got.Prices) != 2 || float64(got.Prices[0]) != 0.42 || len(got.IDs) != 3 || got.IDs[2] != "3" {
		t.Fatalf("unexpected slices: %+v", got)
	}
}

func TestFlexibleTimeTypes(t *testing.T) {
	var got struct {
		Seconds           Time `json:"seconds"`
		Millis            Time `json:"millis"`
		RFC3339           Time `json:"rfc3339"`
		PosgresOffsetHour Time `json:"posgres_offset_hour"`
		Date              Date `json:"date"`
	}
	if err := json.Unmarshal([]byte(`{
		"seconds":"1713398400",
		"millis":1713398400000,
		"rfc3339":"2026-04-25T10:11:12Z",
		"posgres_offset_hour": "2024-04-08 22:29:46.138+00",
		"date":"2026-04-25"
	}`), &got); err != nil {
		t.Fatal(err)
	}
	want := time.Date(2024, 4, 18, 0, 0, 0, 0, time.UTC)
	if !got.Seconds.Time().Equal(want) || !got.Millis.Time().Equal(want) {
		t.Fatalf("seconds=%v millis=%v want=%s", got.Seconds.Time(), got.Millis.Time(), want)
	}
	if got.Date.Time().Format("2006-01-02") != "2026-04-25" {
		t.Fatalf("date = %s", got.Date.Time().Format("2006-01-02"))
	}
}
