package gateway

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	discoverInterface, err := DiscoverInterface()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(discoverInterface)
}
