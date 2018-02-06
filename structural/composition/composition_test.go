package composition

import "testing"

func TestSwimerCompsotionA(t *testing.T) {
	localSwim := Swim
	swimer := CompositeSwimerA{
		MySwim: &localSwim,
	}
	swimer.MyAthlete.Train()
	(*swimer.MySwim)()
}
