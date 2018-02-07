package builder

import "testing"

func TestManufacturing(t *testing.T) {

	manufacturing := ManufacturingDirector{}

	carBuilder := &CarBuilder{}

	manufacturing.SetBuilder(carBuilder)
	manufacturing.Construct()

	car := carBuilder.GetVehicle()

	if car.Wheels != 4 {
		t.Errorf("number of wheels is diferent to 4: %d", car.Wheels)
	}

	bikeBuilder := &BikeBuilder{}

	manufacturing.SetBuilder(bikeBuilder)
	manufacturing.Construct()

	motobike := bikeBuilder.GetVehicle()
	motobike.Seats = 1

	if motobike.Wheels != 2 {
		t.Errorf("number of wheels is diferent to 2: %d", motobike.Wheels)
	}

}

func TestStream(t *testing.T) {

	streamerDirector := StreamerCreation{}

	hlsStreamBuilder := &HLS{}

	streamerDirector.SetBuilder(hlsStreamBuilder)
	streamerDirector.Construct()

	hls := hlsStreamBuilder.GetStreamInfo()

	if hls.Region != 4 {
		t.Errorf("number of region is diferent to 4: %d", hls.Region)
	}

	dashStreamBuilder := &Dash{}

	streamerDirector.SetBuilder(dashStreamBuilder)
	streamerDirector.Construct()

	dash := dashStreamBuilder.GetStreamInfo()
	dash.Region = 1

	if dash.CDN != 1 {
		t.Errorf("number of cdn is diferent to 1: %d", dash.CDN)
	}

}
