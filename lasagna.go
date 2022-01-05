package lasagna

// TODO: define the 'OvenTime' constant
const OvenTime = 40;

// RemainingOvenTime returns the remaining minutes based on the `actual` minutes already in the oven.
func RemainingOvenTime(actualMinutesInOven int) int {
	var RemainingTime = OvenTime - actualMinutesInOven;
    return RemainingTime; 
}


// PreparationTime calculates the time needed to prepare the lasagna based on the amount of layers.
func PreparationTime(numberOfLayers int) int {
	var layerTime = 2;
    var preparationTime = numberOfLayers * layerTime;
    return preparationTime;
}

// ElapsedTime calculates the total time needed to create and bake a lasagna.
func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	var totalTime = numberOfLayers * 2 + actualMinutesInOven;
    return totalTime;
}
