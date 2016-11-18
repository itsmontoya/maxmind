package maxmind

import "testing"

func TestBasic(t *testing.T) {
	if _, err := FilterFile("./worldcitiespop.txt", "", "", "US"); err != nil {
		t.Fatal(err)
		return
	}
}
