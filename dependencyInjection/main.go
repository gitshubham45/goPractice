package main

import "fmt"

type SafetyPlacer interface {
	PlaceSafeties()
}

// type SafetyPlacer struct {
// 	kind int
// }

type RockClimber struct {
	kind         int
	rocksClimbed int
	sp           SafetyPlacer
}

func (rc *RockClimber) ClimbRock() {
	rc.rocksClimbed++
	if rc.rocksClimbed == 10 {
		rc.sp.PlaceSafeties()
	}
}

type IceSafetyPlacer struct {
	// db
	// data
	// api
}

func newRockClimber(sp SafetyPlacer) *RockClimber{
	 return &RockClimber{
		sp : sp,
	 }
}

func (sp IceSafetyPlacer) PlaceSafeties(){
	fmt.Println("placing my Ice Safeties")
}

type NOPSafetyPlacer struct{}

func (sp NOPSafetyPlacer) PlaceSafeties(){
	fmt.Println("placing No Safeties")
}


// func (sp *SafetyPlacer) PlaceSafeties(){
// 	switch sp.kind{
// 	case 1:
// 		// Ice
// 	case 2:
// 		// sand
// 	case 3 :
// 		// Concrete
// 	}
// 	fmt.Println("Safety Placed")
// }

func main() {
	rc := newRockClimber(NOPSafetyPlacer{})
	for i := 0; i < 11; i++ {
		rc.ClimbRock()
	}
}
