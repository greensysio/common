package distance

import (
	"fmt"
	"testing"
)

// Distance stores Origin Location and Destination Location
type Distance2 struct {
	Origin      Location2
	Destination Location2
}

type Location2 struct {
	Name string
}

func TestName(t *testing.T) {
	requestLocations := []Distance2{
		{Location2{"1"}, Location2{"2"}},
		{Location2{"2"}, Location2{"3"}},
		{Location2{"3"}, Location2{"4"}},
		{Location2{"4"}, Location2{"5"}},
		{Location2{"5"}, Location2{"6"}},
		{Location2{"6"}, Location2{"7"}},
		{Location2{"7"}, Location2{"8"}},
		{Location2{"8"}, Location2{"9"}},
		{Location2{"9"}, Location2{"10"}},
		{Location2{"10"}, Location2{"11"}},
		{Location2{"11"}, Location2{"12"}},
	}

	max := 10
	loopNumber := int(len(requestLocations) / (max + 1))
	fmt.Printf("loopNumber: %d \n", loopNumber)
	for i := 0; i <= loopNumber; i++ {
		fmt.Printf("- %d \n", i)
		if i != loopNumber {
			temp := requestLocations[i*max : (i+1)*max]
			fmt.Printf("--- %+v \n", temp)
		} else {
			// Last loop
			temp := requestLocations[i*max:]
			fmt.Printf("--- %+v \n", temp)
		}
	}
}
