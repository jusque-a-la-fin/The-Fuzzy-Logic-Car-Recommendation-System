package entities

import "car-recommendation-service/api/proto/generated/cars"

// GetCarsForStorage переводит данные из типа []*Car в тип []*cars.Car
//
//gocyclo:ignore
func GetCarsForStorage(crs []*Car) []*cars.Car {
	storageCars := make([]*cars.Car, len(crs))
	for idx, car := range crs {
		storageCar := cars.Car{
			Specs: &cars.Specifications{
				Engine:        &cars.Engine{},
				SteeringWheel: &cars.SteeringWheel{},
				Suspension:    &cars.Suspension{},
				Brakes:        &cars.Brakes{},
				Tires:         &cars.Tires{},
			},
			Features: &cars.Features{
				SafetyAndMotionControlSystem: &cars.SafetyAndMotionControlSystems{},
				Lights:                       &cars.Lights{},
				Interior:                     &cars.Interior{},
				CabinMicroclimate:            &cars.CabinMicroclimate{},
				ElectricOptions:              &cars.SetOfElectricOptions{},
				Airbags:                      &cars.SetOfAirbags{},
				MultimediaSystems:            &cars.MultimediaSystems{},
			},
			Offering: &cars.Offering{},
		}

		storageCar.ID = int32(car.ID)
		storageCar.FullName = car.FullName
		storageCar.Description = car.Description
		storageCar.Generation = car.Generation
		storageCar.TrimLevel = car.TrimLevel
		storageCar.Specs.Body = car.Specs.Body
		storageCar.Specs.Length = car.Specs.Length
		storageCar.Specs.Width = car.Specs.Width
		storageCar.Specs.Height = car.Specs.Height
		storageCar.Specs.GroundClearance = car.Specs.GroundClearance
		storageCar.Specs.DragCoefficient = car.Specs.DragCoefficient
		storageCar.Specs.FrontTrackWidth = car.Specs.FrontTrackWidth
		storageCar.Specs.BackTrackWidth = car.Specs.BackTrackWidth
		storageCar.Specs.Wheelbase = car.Specs.Wheelbase
		storageCar.Specs.Acceleration0To100 = car.Specs.Acceleration0To100
		storageCar.Specs.MaxSpeed = car.Specs.MaxSpeed
		storageCar.Specs.CityFuelConsumption = car.Specs.CityFuelConsumption
		storageCar.Specs.HighwayFuelConsumption = car.Specs.HighwayFuelConsumption
		storageCar.Specs.MixedFuelConsumption = car.Specs.MixedFuelConsumption
		storageCar.Specs.NumberOfSeats = int32(car.Specs.NumberOfSeats)
		storageCar.Specs.TrunkVolume = car.Specs.TrunkVolume
		storageCar.Specs.Mass = car.Specs.Mass
		storageCar.Specs.Gearbox = car.Specs.Gearbox
		storageCar.Specs.Drive = car.Specs.Drive
		storageCar.Specs.CrashTestEstimate = car.Specs.CrashTestEstimate
		storageCar.Specs.Engine.Capacity = car.Specs.Engine.Capacity
		storageCar.Specs.Engine.EngineType = car.Specs.Engine.EngineType
		storageCar.Specs.Engine.FuelUsed = car.Specs.Engine.FuelUsed
		storageCar.Specs.Engine.MaxPower = car.Specs.Engine.MaxPower
		storageCar.Specs.Engine.MaxTorque = car.Specs.Engine.MaxTorque
		switch car.Specs.SteeringWheel.PowerSteering {
		case ElectricPS:
			storageCar.Specs.SteeringWheel.PowerSteering = cars.PowerSteering_ELECTRIC_PS
		case ElectrohydraulicPS:
			storageCar.Specs.SteeringWheel.PowerSteering = cars.PowerSteering_ELECTROHYDRAULIC_PS
		case HydraulicPS:
			storageCar.Specs.SteeringWheel.PowerSteering = cars.PowerSteering_HYDRAULIC_PS
		case NoPS:
			storageCar.Specs.SteeringWheel.PowerSteering = cars.PowerSteering_NO_PS
		case UndefinedPS:
			storageCar.Specs.SteeringWheel.PowerSteering = cars.PowerSteering_UNDEFINED_PS
		}

		switch car.Specs.SteeringWheel.SteeringWheelPosition {
		case LeftPos:
			storageCar.Specs.SteeringWheel.SteeringWheelPosition = cars.SteeringWheelPosition_LEFT_POS
		case RightPos:
			storageCar.Specs.SteeringWheel.SteeringWheelPosition = cars.SteeringWheelPosition_RIGHT_POS
		case UndefinedPos:
			storageCar.Specs.SteeringWheel.SteeringWheelPosition = cars.SteeringWheelPosition_UNDEFINED_POS
		}

		switch car.Specs.Suspension.FrontStabilizer {
		case YesValue:
			storageCar.Specs.Suspension.FrontStabilizer = cars.Availability_YES_VALUE
		case NoValue:
			storageCar.Specs.Suspension.FrontStabilizer = cars.Availability_NO_VALUE
		case OptionValue:
			storageCar.Specs.Suspension.FrontStabilizer = cars.Availability_OPTION_VALUE
		case UndefinedValue:
			storageCar.Specs.Suspension.FrontStabilizer = cars.Availability_UNDEFINED_VALUE
		}

		switch car.Specs.Suspension.BackStabilizer {
		case YesValue:
			storageCar.Specs.Suspension.BackStabilizer = cars.Availability_YES_VALUE
		case NoValue:
			storageCar.Specs.Suspension.BackStabilizer = cars.Availability_NO_VALUE
		case OptionValue:
			storageCar.Specs.Suspension.BackStabilizer = cars.Availability_OPTION_VALUE
		case UndefinedValue:
			storageCar.Specs.Suspension.BackStabilizer = cars.Availability_UNDEFINED_VALUE
		}

		storageCar.Specs.Suspension.FrontSuspension = car.Specs.Suspension.FrontSuspension
		storageCar.Specs.Suspension.BackSuspension = car.Specs.Suspension.BackSuspension
		storageCar.Specs.Brakes.FrontBrakes = car.Specs.Brakes.FrontBrakes
		storageCar.Specs.Brakes.BackBrakes = car.Specs.Brakes.BackBrakes
		storageCar.Specs.Brakes.ParkingBrake = car.Specs.Brakes.ParkingBrake
		storageCar.Specs.Tires.FrontTiresWidth = int32(car.Specs.Tires.FrontTiresWidth)
		storageCar.Specs.Tires.BackTiresWidth = int32(car.Specs.Tires.BackTiresWidth)
		storageCar.Specs.Tires.FrontTiresAspectRatio = int32(car.Specs.Tires.FrontTiresAspectRatio)
		storageCar.Specs.Tires.BackTiresAspectRatio = int32(car.Specs.Tires.BackTiresAspectRatio)
		storageCar.Specs.Tires.FrontTiresRimDiameter = int32(car.Specs.Tires.FrontTiresRimDiameter)
		storageCar.Specs.Tires.BackTiresRimDiameter = int32(car.Specs.Tires.BackTiresRimDiameter)

		switch car.Features.SafetyAndMotionControlSystem.ABS {
		case YesValue:
			storageCar.Features.SafetyAndMotionControlSystem.ABS = cars.Availability_YES_VALUE
		case NoValue:
			storageCar.Features.SafetyAndMotionControlSystem.ABS = cars.Availability_NO_VALUE
		case OptionValue:
			storageCar.Features.SafetyAndMotionControlSystem.ABS = cars.Availability_OPTION_VALUE
		case UndefinedValue:
			storageCar.Features.SafetyAndMotionControlSystem.ABS = cars.Availability_UNDEFINED_VALUE
		}

		switch car.Features.SafetyAndMotionControlSystem.ESP {
		case YesValue:
			storageCar.Features.SafetyAndMotionControlSystem.ESP = cars.Availability_YES_VALUE
		case NoValue:
			storageCar.Features.SafetyAndMotionControlSystem.ESP = cars.Availability_NO_VALUE
		case OptionValue:
			storageCar.Features.SafetyAndMotionControlSystem.ESP = cars.Availability_OPTION_VALUE
		case UndefinedValue:
			storageCar.Features.SafetyAndMotionControlSystem.ESP = cars.Availability_UNDEFINED_VALUE
		}

		switch car.Features.SafetyAndMotionControlSystem.EBD {
		case YesValue:
			storageCar.Features.SafetyAndMotionControlSystem.EBD = cars.Availability_YES_VALUE
		case NoValue:
			storageCar.Features.SafetyAndMotionControlSystem.EBD = cars.Availability_NO_VALUE
		case OptionValue:
			storageCar.Features.SafetyAndMotionControlSystem.EBD = cars.Availability_OPTION_VALUE
		case UndefinedValue:
			storageCar.Features.SafetyAndMotionControlSystem.EBD = cars.Availability_UNDEFINED_VALUE
		}

		switch car.Features.SafetyAndMotionControlSystem.BAS {
		case YesValue:
			storageCar.Features.SafetyAndMotionControlSystem.BAS = cars.Availability_YES_VALUE
		case NoValue:
			storageCar.Features.SafetyAndMotionControlSystem.BAS = cars.Availability_NO_VALUE
		case OptionValue:
			storageCar.Features.SafetyAndMotionControlSystem.BAS = cars.Availability_OPTION_VALUE
		case UndefinedValue:
			storageCar.Features.SafetyAndMotionControlSystem.BAS = cars.Availability_UNDEFINED_VALUE
		}

		switch car.Features.SafetyAndMotionControlSystem.TCS {
		case YesValue:
			storageCar.Features.SafetyAndMotionControlSystem.TCS = cars.Availability_YES_VALUE
		case NoValue:
			storageCar.Features.SafetyAndMotionControlSystem.TCS = cars.Availability_NO_VALUE
		case OptionValue:
			storageCar.Features.SafetyAndMotionControlSystem.TCS = cars.Availability_OPTION_VALUE
		case UndefinedValue:
			storageCar.Features.SafetyAndMotionControlSystem.TCS = cars.Availability_UNDEFINED_VALUE
		}

		switch car.Features.SafetyAndMotionControlSystem.FrontParkingSensor {
		case YesValue:
			storageCar.Features.SafetyAndMotionControlSystem.FrontParkingSensor = cars.Availability_YES_VALUE
		case NoValue:
			storageCar.Features.SafetyAndMotionControlSystem.FrontParkingSensor = cars.Availability_NO_VALUE
		case OptionValue:
			storageCar.Features.SafetyAndMotionControlSystem.FrontParkingSensor = cars.Availability_OPTION_VALUE
		case UndefinedValue:
			storageCar.Features.SafetyAndMotionControlSystem.FrontParkingSensor = cars.Availability_UNDEFINED_VALUE
		}

		switch car.Features.SafetyAndMotionControlSystem.BackParkingSensor {
		case YesValue:
			storageCar.Features.SafetyAndMotionControlSystem.BackParkingSensor = cars.Availability_YES_VALUE
		case NoValue:
			storageCar.Features.SafetyAndMotionControlSystem.BackParkingSensor = cars.Availability_NO_VALUE
		case OptionValue:
			storageCar.Features.SafetyAndMotionControlSystem.BackParkingSensor = cars.Availability_OPTION_VALUE
		case UndefinedValue:
			storageCar.Features.SafetyAndMotionControlSystem.BackParkingSensor = cars.Availability_UNDEFINED_VALUE
		}

		switch car.Features.SafetyAndMotionControlSystem.RearViewCamera {
		case YesValue:
			storageCar.Features.SafetyAndMotionControlSystem.RearViewCamera = cars.Availability_YES_VALUE
		case NoValue:
			storageCar.Features.SafetyAndMotionControlSystem.RearViewCamera = cars.Availability_NO_VALUE
		case OptionValue:
			storageCar.Features.SafetyAndMotionControlSystem.RearViewCamera = cars.Availability_OPTION_VALUE
		case UndefinedValue:
			storageCar.Features.SafetyAndMotionControlSystem.RearViewCamera = cars.Availability_UNDEFINED_VALUE
		}

		switch car.Features.SafetyAndMotionControlSystem.CruiseControl {
		case YesValue:
			storageCar.Features.SafetyAndMotionControlSystem.CruiseControl = cars.Availability_YES_VALUE
		case NoValue:
			storageCar.Features.SafetyAndMotionControlSystem.CruiseControl = cars.Availability_NO_VALUE
		case OptionValue:
			storageCar.Features.SafetyAndMotionControlSystem.CruiseControl = cars.Availability_OPTION_VALUE
		case UndefinedValue:
			storageCar.Features.SafetyAndMotionControlSystem.CruiseControl = cars.Availability_UNDEFINED_VALUE
		}

		car.Features.Lights.Headlights = storageCar.Features.Lights.Headlights
		switch car.Features.Lights.LEDRunningLights {
		case YesValue:
			storageCar.Features.Lights.LEDRunningLights = cars.Availability_YES_VALUE
		case NoValue:
			storageCar.Features.Lights.LEDRunningLights = cars.Availability_NO_VALUE
		case OptionValue:
			storageCar.Features.Lights.LEDRunningLights = cars.Availability_OPTION_VALUE
		case UndefinedValue:
			storageCar.Features.Lights.LEDRunningLights = cars.Availability_UNDEFINED_VALUE
		}

		switch car.Features.Lights.LEDTailLights {
		case YesValue:
			storageCar.Features.Lights.LEDTailLights = cars.Availability_YES_VALUE
		case NoValue:
			storageCar.Features.Lights.LEDTailLights = cars.Availability_NO_VALUE
		case OptionValue:
			storageCar.Features.Lights.LEDTailLights = cars.Availability_OPTION_VALUE
		case UndefinedValue:
			storageCar.Features.Lights.LEDTailLights = cars.Availability_UNDEFINED_VALUE
		}

		switch car.Features.Lights.LightSensor {
		case YesValue:
			storageCar.Features.Lights.LightSensor = cars.Availability_YES_VALUE
		case NoValue:
			storageCar.Features.Lights.LightSensor = cars.Availability_NO_VALUE
		case OptionValue:
			storageCar.Features.Lights.LightSensor = cars.Availability_OPTION_VALUE
		case UndefinedValue:
			storageCar.Features.Lights.LightSensor = cars.Availability_UNDEFINED_VALUE
		}

		switch car.Features.Lights.FrontFogLights {
		case YesValue:
			storageCar.Features.Lights.FrontFogLights = cars.Availability_YES_VALUE
		case NoValue:
			storageCar.Features.Lights.FrontFogLights = cars.Availability_NO_VALUE
		case OptionValue:
			storageCar.Features.Lights.FrontFogLights = cars.Availability_OPTION_VALUE
		case UndefinedValue:
			storageCar.Features.Lights.FrontFogLights = cars.Availability_UNDEFINED_VALUE
		}

		switch car.Features.Lights.BackFogLights {
		case YesValue:
			storageCar.Features.Lights.BackFogLights = cars.Availability_YES_VALUE
		case NoValue:
			storageCar.Features.Lights.BackFogLights = cars.Availability_NO_VALUE
		case OptionValue:
			storageCar.Features.Lights.BackFogLights = cars.Availability_OPTION_VALUE
		case UndefinedValue:
			storageCar.Features.Lights.BackFogLights = cars.Availability_UNDEFINED_VALUE
		}

		storageCar.Features.Interior.Upholstery = car.Features.Interior.Upholstery

		switch car.Features.CabinMicroclimate.AirConditioner {
		case YesValue:
			storageCar.Features.CabinMicroclimate.AirConditioner = cars.Availability_YES_VALUE
		case NoValue:
			storageCar.Features.CabinMicroclimate.AirConditioner = cars.Availability_NO_VALUE
		case OptionValue:
			storageCar.Features.CabinMicroclimate.AirConditioner = cars.Availability_OPTION_VALUE
		case UndefinedValue:
			storageCar.Features.CabinMicroclimate.AirConditioner = cars.Availability_UNDEFINED_VALUE
		}

		switch car.Features.CabinMicroclimate.ClimateControl {
		case YesValue:
			storageCar.Features.CabinMicroclimate.ClimateControl = cars.Availability_YES_VALUE
		case NoValue:
			storageCar.Features.CabinMicroclimate.ClimateControl = cars.Availability_NO_VALUE
		case OptionValue:
			storageCar.Features.CabinMicroclimate.ClimateControl = cars.Availability_OPTION_VALUE
		case UndefinedValue:
			storageCar.Features.CabinMicroclimate.ClimateControl = cars.Availability_UNDEFINED_VALUE
		}

		switch car.Features.ElectricOptions.ElectricFrontSideWindowsLifts {
		case YesValue:
			storageCar.Features.ElectricOptions.ElectricFrontSideWindowsLifts = cars.Availability_YES_VALUE
		case NoValue:
			storageCar.Features.ElectricOptions.ElectricFrontSideWindowsLifts = cars.Availability_NO_VALUE
		case OptionValue:
			storageCar.Features.ElectricOptions.ElectricFrontSideWindowsLifts = cars.Availability_OPTION_VALUE
		case UndefinedValue:
			storageCar.Features.ElectricOptions.ElectricFrontSideWindowsLifts = cars.Availability_UNDEFINED_VALUE
		}

		switch car.Features.ElectricOptions.ElectricBackSideWindowsLifts {
		case YesValue:
			storageCar.Features.ElectricOptions.ElectricBackSideWindowsLifts = cars.Availability_YES_VALUE
		case NoValue:
			storageCar.Features.ElectricOptions.ElectricBackSideWindowsLifts = cars.Availability_NO_VALUE
		case OptionValue:
			storageCar.Features.ElectricOptions.ElectricBackSideWindowsLifts = cars.Availability_OPTION_VALUE
		case UndefinedValue:
			storageCar.Features.ElectricOptions.ElectricBackSideWindowsLifts = cars.Availability_UNDEFINED_VALUE
		}

		switch car.Features.ElectricOptions.ElectricHeatingOfFrontSeats {
		case YesValue:
			storageCar.Features.ElectricOptions.ElectricHeatingOfFrontSeats = cars.Availability_YES_VALUE
		case NoValue:
			storageCar.Features.ElectricOptions.ElectricHeatingOfFrontSeats = cars.Availability_NO_VALUE
		case OptionValue:
			storageCar.Features.ElectricOptions.ElectricHeatingOfFrontSeats = cars.Availability_OPTION_VALUE
		case UndefinedValue:
			storageCar.Features.ElectricOptions.ElectricHeatingOfFrontSeats = cars.Availability_UNDEFINED_VALUE
		}

		switch car.Features.ElectricOptions.ElectricHeatingOfBackSeats {
		case YesValue:
			storageCar.Features.ElectricOptions.ElectricHeatingOfBackSeats = cars.Availability_YES_VALUE
		case NoValue:
			storageCar.Features.ElectricOptions.ElectricHeatingOfBackSeats = cars.Availability_NO_VALUE
		case OptionValue:
			storageCar.Features.ElectricOptions.ElectricHeatingOfBackSeats = cars.Availability_OPTION_VALUE
		case UndefinedValue:
			storageCar.Features.ElectricOptions.ElectricHeatingOfBackSeats = cars.Availability_UNDEFINED_VALUE
		}

		switch car.Features.ElectricOptions.ElectricHeatingOfSteeringWheel {
		case YesValue:
			storageCar.Features.ElectricOptions.ElectricHeatingOfSteeringWheel = cars.Availability_YES_VALUE
		case NoValue:
			storageCar.Features.ElectricOptions.ElectricHeatingOfSteeringWheel = cars.Availability_NO_VALUE
		case OptionValue:
			storageCar.Features.ElectricOptions.ElectricHeatingOfSteeringWheel = cars.Availability_OPTION_VALUE
		case UndefinedValue:
			storageCar.Features.ElectricOptions.ElectricHeatingOfSteeringWheel = cars.Availability_UNDEFINED_VALUE
		}

		switch car.Features.ElectricOptions.ElectricHeatingOfWindshield {
		case YesValue:
			storageCar.Features.ElectricOptions.ElectricHeatingOfWindshield = cars.Availability_YES_VALUE
		case NoValue:
			storageCar.Features.ElectricOptions.ElectricHeatingOfWindshield = cars.Availability_NO_VALUE
		case OptionValue:
			storageCar.Features.ElectricOptions.ElectricHeatingOfWindshield = cars.Availability_OPTION_VALUE
		case UndefinedValue:
			storageCar.Features.ElectricOptions.ElectricHeatingOfWindshield = cars.Availability_UNDEFINED_VALUE
		}

		switch car.Features.ElectricOptions.ElectricHeatingOfRearWindow {
		case YesValue:
			storageCar.Features.ElectricOptions.ElectricHeatingOfRearWindow = cars.Availability_YES_VALUE
		case NoValue:
			storageCar.Features.ElectricOptions.ElectricHeatingOfRearWindow = cars.Availability_NO_VALUE
		case OptionValue:
			storageCar.Features.ElectricOptions.ElectricHeatingOfRearWindow = cars.Availability_OPTION_VALUE
		case UndefinedValue:
			storageCar.Features.ElectricOptions.ElectricHeatingOfRearWindow = cars.Availability_UNDEFINED_VALUE
		}

		switch car.Features.ElectricOptions.ElectricHeatingOfSideMirrors {
		case YesValue:
			storageCar.Features.ElectricOptions.ElectricHeatingOfSideMirrors = cars.Availability_YES_VALUE
		case NoValue:
			storageCar.Features.ElectricOptions.ElectricHeatingOfSideMirrors = cars.Availability_NO_VALUE
		case OptionValue:
			storageCar.Features.ElectricOptions.ElectricHeatingOfSideMirrors = cars.Availability_OPTION_VALUE
		case UndefinedValue:
			storageCar.Features.ElectricOptions.ElectricHeatingOfSideMirrors = cars.Availability_UNDEFINED_VALUE
		}

		switch car.Features.ElectricOptions.ElectricDriveOfDriverSeat {
		case YesValue:
			storageCar.Features.ElectricOptions.ElectricDriveOfDriverSeat = cars.Availability_YES_VALUE
		case NoValue:
			storageCar.Features.ElectricOptions.ElectricDriveOfDriverSeat = cars.Availability_NO_VALUE
		case OptionValue:
			storageCar.Features.ElectricOptions.ElectricDriveOfDriverSeat = cars.Availability_OPTION_VALUE
		case UndefinedValue:
			storageCar.Features.ElectricOptions.ElectricDriveOfDriverSeat = cars.Availability_UNDEFINED_VALUE
		}

		switch car.Features.ElectricOptions.ElectricDriveOfFrontSeats {
		case YesValue:
			storageCar.Features.ElectricOptions.ElectricDriveOfFrontSeats = cars.Availability_YES_VALUE
		case NoValue:
			storageCar.Features.ElectricOptions.ElectricDriveOfFrontSeats = cars.Availability_NO_VALUE
		case OptionValue:
			storageCar.Features.ElectricOptions.ElectricDriveOfFrontSeats = cars.Availability_OPTION_VALUE
		case UndefinedValue:
			storageCar.Features.ElectricOptions.ElectricDriveOfFrontSeats = cars.Availability_UNDEFINED_VALUE
		}

		switch car.Features.ElectricOptions.ElectricDriveOfSideMirrors {
		case YesValue:
			storageCar.Features.ElectricOptions.ElectricDriveOfSideMirrors = cars.Availability_YES_VALUE
		case NoValue:
			storageCar.Features.ElectricOptions.ElectricDriveOfSideMirrors = cars.Availability_NO_VALUE
		case OptionValue:
			storageCar.Features.ElectricOptions.ElectricDriveOfSideMirrors = cars.Availability_OPTION_VALUE
		case UndefinedValue:
			storageCar.Features.ElectricOptions.ElectricDriveOfSideMirrors = cars.Availability_UNDEFINED_VALUE
		}

		switch car.Features.ElectricOptions.ElectricTrunkOpener {
		case YesValue:
			storageCar.Features.ElectricOptions.ElectricTrunkOpener = cars.Availability_YES_VALUE
		case NoValue:
			storageCar.Features.ElectricOptions.ElectricTrunkOpener = cars.Availability_NO_VALUE
		case OptionValue:
			storageCar.Features.ElectricOptions.ElectricTrunkOpener = cars.Availability_OPTION_VALUE
		case UndefinedValue:
			storageCar.Features.ElectricOptions.ElectricTrunkOpener = cars.Availability_UNDEFINED_VALUE
		}

		switch car.Features.ElectricOptions.RainSensor {
		case YesValue:
			storageCar.Features.ElectricOptions.RainSensor = cars.Availability_YES_VALUE
		case NoValue:
			storageCar.Features.ElectricOptions.RainSensor = cars.Availability_NO_VALUE
		case OptionValue:
			storageCar.Features.ElectricOptions.RainSensor = cars.Availability_OPTION_VALUE
		case UndefinedValue:
			storageCar.Features.ElectricOptions.RainSensor = cars.Availability_UNDEFINED_VALUE
		}

		switch car.Features.Airbags.DriverAirbag {
		case YesValue:
			storageCar.Features.Airbags.DriverAirbag = cars.Availability_YES_VALUE
		case NoValue:
			storageCar.Features.Airbags.DriverAirbag = cars.Availability_NO_VALUE
		case OptionValue:
			storageCar.Features.Airbags.DriverAirbag = cars.Availability_OPTION_VALUE
		case UndefinedValue:
			storageCar.Features.Airbags.DriverAirbag = cars.Availability_UNDEFINED_VALUE
		}

		switch car.Features.Airbags.FrontPassengerAirbag {
		case YesValue:
			storageCar.Features.Airbags.FrontPassengerAirbag = cars.Availability_YES_VALUE
		case NoValue:
			storageCar.Features.Airbags.FrontPassengerAirbag = cars.Availability_NO_VALUE
		case OptionValue:
			storageCar.Features.Airbags.FrontPassengerAirbag = cars.Availability_OPTION_VALUE
		case UndefinedValue:
			storageCar.Features.Airbags.FrontPassengerAirbag = cars.Availability_UNDEFINED_VALUE
		}

		switch car.Features.Airbags.SideAirbags {
		case YesValue:
			storageCar.Features.Airbags.SideAirbags = cars.Availability_YES_VALUE
		case NoValue:
			storageCar.Features.Airbags.SideAirbags = cars.Availability_NO_VALUE
		case OptionValue:
			storageCar.Features.Airbags.SideAirbags = cars.Availability_OPTION_VALUE
		case UndefinedValue:
			storageCar.Features.Airbags.SideAirbags = cars.Availability_UNDEFINED_VALUE
		}

		switch car.Features.Airbags.CurtainAirbags {
		case YesValue:
			storageCar.Features.Airbags.CurtainAirbags = cars.Availability_YES_VALUE
		case NoValue:
			storageCar.Features.Airbags.CurtainAirbags = cars.Availability_NO_VALUE
		case OptionValue:
			storageCar.Features.Airbags.CurtainAirbags = cars.Availability_OPTION_VALUE
		case UndefinedValue:
			storageCar.Features.Airbags.CurtainAirbags = cars.Availability_UNDEFINED_VALUE
		}

		switch car.Features.MultimediaSystems.HandsFreeSupport {
		case YesValue:
			storageCar.Features.MultimediaSystems.HandsFreeSupport = cars.Availability_YES_VALUE
		case NoValue:
			storageCar.Features.MultimediaSystems.HandsFreeSupport = cars.Availability_NO_VALUE
		case OptionValue:
			storageCar.Features.MultimediaSystems.HandsFreeSupport = cars.Availability_OPTION_VALUE
		case UndefinedValue:
			storageCar.Features.MultimediaSystems.HandsFreeSupport = cars.Availability_UNDEFINED_VALUE
		}

		switch car.Features.MultimediaSystems.MP3Support {
		case YesValue:
			storageCar.Features.MultimediaSystems.MP3Support = cars.Availability_YES_VALUE
		case NoValue:
			storageCar.Features.MultimediaSystems.MP3Support = cars.Availability_NO_VALUE
		case OptionValue:
			storageCar.Features.MultimediaSystems.MP3Support = cars.Availability_OPTION_VALUE
		case UndefinedValue:
			storageCar.Features.MultimediaSystems.MP3Support = cars.Availability_UNDEFINED_VALUE
		}

		switch car.Features.MultimediaSystems.OnBoardComputer {
		case YesValue:
			storageCar.Features.MultimediaSystems.OnBoardComputer = cars.Availability_YES_VALUE
		case NoValue:
			storageCar.Features.MultimediaSystems.OnBoardComputer = cars.Availability_NO_VALUE
		case OptionValue:
			storageCar.Features.MultimediaSystems.OnBoardComputer = cars.Availability_OPTION_VALUE
		case UndefinedValue:
			storageCar.Features.MultimediaSystems.OnBoardComputer = cars.Availability_UNDEFINED_VALUE
		}

		switch car.Features.CarAlarm {
		case YesValue:
			storageCar.Features.CarAlarm = cars.Availability_YES_VALUE
		case NoValue:
			storageCar.Features.CarAlarm = cars.Availability_NO_VALUE
		case OptionValue:
			storageCar.Features.CarAlarm = cars.Availability_OPTION_VALUE
		case UndefinedValue:
			storageCar.Features.CarAlarm = cars.Availability_UNDEFINED_VALUE
		}

		storageCar.Features.Color = car.Features.Color

		storageCar.Offering.Price = car.Offering.Price
		storageCar.Offering.Year = int32(car.Offering.Year)
		storageCar.Offering.Kilometrage = car.Offering.Kilometrage
		storageCar.Offering.PhotoURLs = car.Offering.PhotoURLs
		storageCars[idx] = &storageCar
	}
	return storageCars
}

// GetCars переводит данные из типа []*cars.Car в тип []Car
//
//gocyclo:ignore
func GetCars(storageCars []*cars.Car) []*Car {
	crs := make([]*Car, len(storageCars))
	for idx, storageCar := range storageCars {
		car := Car{}
		car.ID = int(storageCar.ID)
		car.FullName = storageCar.FullName
		car.Description = storageCar.Description
		car.Generation = storageCar.Generation
		car.TrimLevel = storageCar.TrimLevel
		car.Specs.Body = storageCar.Specs.Body
		car.Specs.Length = storageCar.Specs.Length
		car.Specs.Width = storageCar.Specs.Width
		car.Specs.Height = storageCar.Specs.Height
		car.Specs.GroundClearance = storageCar.Specs.GroundClearance
		car.Specs.DragCoefficient = storageCar.Specs.DragCoefficient
		car.Specs.FrontTrackWidth = storageCar.Specs.FrontTrackWidth
		car.Specs.BackTrackWidth = storageCar.Specs.BackTrackWidth
		car.Specs.Wheelbase = storageCar.Specs.Wheelbase
		car.Specs.Acceleration0To100 = storageCar.Specs.Acceleration0To100
		car.Specs.MaxSpeed = storageCar.Specs.MaxSpeed
		car.Specs.CityFuelConsumption = storageCar.Specs.CityFuelConsumption
		car.Specs.HighwayFuelConsumption = storageCar.Specs.HighwayFuelConsumption
		car.Specs.MixedFuelConsumption = storageCar.Specs.MixedFuelConsumption
		car.Specs.NumberOfSeats = int(storageCar.Specs.NumberOfSeats)
		car.Specs.TrunkVolume = storageCar.Specs.TrunkVolume
		car.Specs.Mass = storageCar.Specs.Mass
		car.Specs.Gearbox = storageCar.Specs.Gearbox
		car.Specs.Drive = storageCar.Specs.Drive
		car.Specs.CrashTestEstimate = storageCar.Specs.CrashTestEstimate
		car.Specs.Engine.Capacity = storageCar.Specs.Engine.Capacity
		car.Specs.Engine.EngineType = storageCar.Specs.Engine.EngineType
		car.Specs.Engine.FuelUsed = storageCar.Specs.Engine.FuelUsed
		car.Specs.Engine.MaxPower = storageCar.Specs.Engine.MaxPower
		car.Specs.Engine.MaxTorque = storageCar.Specs.Engine.MaxTorque
		switch storageCar.Specs.SteeringWheel.PowerSteering {
		case cars.PowerSteering_ELECTRIC_PS:
			car.Specs.SteeringWheel.PowerSteering = ElectricPS
		case cars.PowerSteering_ELECTROHYDRAULIC_PS:
			car.Specs.SteeringWheel.PowerSteering = ElectrohydraulicPS
		case cars.PowerSteering_HYDRAULIC_PS:
			car.Specs.SteeringWheel.PowerSteering = HydraulicPS
		case cars.PowerSteering_NO_PS:
			car.Specs.SteeringWheel.PowerSteering = NoPS
		case cars.PowerSteering_UNDEFINED_PS:
			car.Specs.SteeringWheel.PowerSteering = UndefinedPS
		}

		switch storageCar.Specs.SteeringWheel.SteeringWheelPosition {
		case cars.SteeringWheelPosition_LEFT_POS:
			car.Specs.SteeringWheel.SteeringWheelPosition = LeftPos
		case cars.SteeringWheelPosition_RIGHT_POS:
			car.Specs.SteeringWheel.SteeringWheelPosition = RightPos
		case cars.SteeringWheelPosition_UNDEFINED_POS:
			car.Specs.SteeringWheel.SteeringWheelPosition = UndefinedPos
		}

		switch storageCar.Specs.Suspension.FrontStabilizer {
		case cars.Availability_YES_VALUE:
			car.Specs.Suspension.FrontStabilizer = YesValue
		case cars.Availability_NO_VALUE:
			car.Specs.Suspension.FrontStabilizer = NoValue
		case cars.Availability_OPTION_VALUE:
			car.Specs.Suspension.FrontStabilizer = OptionValue
		case cars.Availability_UNDEFINED_VALUE:
			car.Specs.Suspension.FrontStabilizer = UndefinedValue
		}

		switch storageCar.Specs.Suspension.BackStabilizer {
		case cars.Availability_YES_VALUE:
			car.Specs.Suspension.BackStabilizer = YesValue
		case cars.Availability_NO_VALUE:
			car.Specs.Suspension.BackStabilizer = NoValue
		case cars.Availability_OPTION_VALUE:
			car.Specs.Suspension.BackStabilizer = OptionValue
		case cars.Availability_UNDEFINED_VALUE:
			car.Specs.Suspension.BackStabilizer = UndefinedValue
		}

		car.Specs.Suspension.FrontSuspension = storageCar.Specs.Suspension.FrontSuspension
		car.Specs.Suspension.BackSuspension = storageCar.Specs.Suspension.BackSuspension
		car.Specs.Brakes.FrontBrakes = storageCar.Specs.Brakes.FrontBrakes
		car.Specs.Brakes.BackBrakes = storageCar.Specs.Brakes.BackBrakes
		car.Specs.Brakes.ParkingBrake = storageCar.Specs.Brakes.ParkingBrake
		car.Specs.Tires.FrontTiresWidth = int(storageCar.Specs.Tires.FrontTiresWidth)
		car.Specs.Tires.BackTiresWidth = int(storageCar.Specs.Tires.BackTiresWidth)
		car.Specs.Tires.FrontTiresAspectRatio = int(storageCar.Specs.Tires.FrontTiresAspectRatio)
		car.Specs.Tires.BackTiresAspectRatio = int(storageCar.Specs.Tires.BackTiresAspectRatio)
		car.Specs.Tires.FrontTiresRimDiameter = int(storageCar.Specs.Tires.FrontTiresRimDiameter)
		car.Specs.Tires.BackTiresRimDiameter = int(storageCar.Specs.Tires.BackTiresRimDiameter)

		switch storageCar.Features.SafetyAndMotionControlSystem.ABS {
		case cars.Availability_YES_VALUE:
			car.Features.SafetyAndMotionControlSystem.ABS = YesValue
		case cars.Availability_NO_VALUE:
			car.Features.SafetyAndMotionControlSystem.ABS = NoValue
		case cars.Availability_OPTION_VALUE:
			car.Features.SafetyAndMotionControlSystem.ABS = OptionValue
		case cars.Availability_UNDEFINED_VALUE:
			car.Features.SafetyAndMotionControlSystem.ABS = UndefinedValue
		}

		switch storageCar.Features.SafetyAndMotionControlSystem.ESP {
		case cars.Availability_YES_VALUE:
			car.Features.SafetyAndMotionControlSystem.ESP = YesValue
		case cars.Availability_NO_VALUE:
			car.Features.SafetyAndMotionControlSystem.ESP = NoValue
		case cars.Availability_OPTION_VALUE:
			car.Features.SafetyAndMotionControlSystem.ESP = OptionValue
		case cars.Availability_UNDEFINED_VALUE:
			car.Features.SafetyAndMotionControlSystem.ESP = UndefinedValue
		}

		switch storageCar.Features.SafetyAndMotionControlSystem.EBD {
		case cars.Availability_YES_VALUE:
			car.Features.SafetyAndMotionControlSystem.EBD = YesValue
		case cars.Availability_NO_VALUE:
			car.Features.SafetyAndMotionControlSystem.EBD = NoValue
		case cars.Availability_OPTION_VALUE:
			car.Features.SafetyAndMotionControlSystem.EBD = OptionValue
		case cars.Availability_UNDEFINED_VALUE:
			car.Features.SafetyAndMotionControlSystem.EBD = UndefinedValue
		}

		switch storageCar.Features.SafetyAndMotionControlSystem.BAS {
		case cars.Availability_YES_VALUE:
			car.Features.SafetyAndMotionControlSystem.BAS = YesValue
		case cars.Availability_NO_VALUE:
			car.Features.SafetyAndMotionControlSystem.BAS = NoValue
		case cars.Availability_OPTION_VALUE:
			car.Features.SafetyAndMotionControlSystem.BAS = OptionValue
		case cars.Availability_UNDEFINED_VALUE:
			car.Features.SafetyAndMotionControlSystem.BAS = UndefinedValue
		}

		switch storageCar.Features.SafetyAndMotionControlSystem.TCS {
		case cars.Availability_YES_VALUE:
			car.Features.SafetyAndMotionControlSystem.TCS = YesValue
		case cars.Availability_NO_VALUE:
			car.Features.SafetyAndMotionControlSystem.TCS = NoValue
		case cars.Availability_OPTION_VALUE:
			car.Features.SafetyAndMotionControlSystem.TCS = OptionValue
		case cars.Availability_UNDEFINED_VALUE:
			car.Features.SafetyAndMotionControlSystem.TCS = UndefinedValue
		}

		switch storageCar.Features.SafetyAndMotionControlSystem.FrontParkingSensor {
		case cars.Availability_YES_VALUE:
			car.Features.SafetyAndMotionControlSystem.FrontParkingSensor = YesValue
		case cars.Availability_NO_VALUE:
			car.Features.SafetyAndMotionControlSystem.FrontParkingSensor = NoValue
		case cars.Availability_OPTION_VALUE:
			car.Features.SafetyAndMotionControlSystem.FrontParkingSensor = OptionValue
		case cars.Availability_UNDEFINED_VALUE:
			car.Features.SafetyAndMotionControlSystem.FrontParkingSensor = UndefinedValue
		}

		switch storageCar.Features.SafetyAndMotionControlSystem.BackParkingSensor {
		case cars.Availability_YES_VALUE:
			car.Features.SafetyAndMotionControlSystem.BackParkingSensor = YesValue
		case cars.Availability_NO_VALUE:
			car.Features.SafetyAndMotionControlSystem.BackParkingSensor = NoValue
		case cars.Availability_OPTION_VALUE:
			car.Features.SafetyAndMotionControlSystem.BackParkingSensor = OptionValue
		case cars.Availability_UNDEFINED_VALUE:
			car.Features.SafetyAndMotionControlSystem.BackParkingSensor = UndefinedValue
		}

		switch storageCar.Features.SafetyAndMotionControlSystem.RearViewCamera {
		case cars.Availability_YES_VALUE:
			car.Features.SafetyAndMotionControlSystem.RearViewCamera = YesValue
		case cars.Availability_NO_VALUE:
			car.Features.SafetyAndMotionControlSystem.RearViewCamera = NoValue
		case cars.Availability_OPTION_VALUE:
			car.Features.SafetyAndMotionControlSystem.RearViewCamera = OptionValue
		case cars.Availability_UNDEFINED_VALUE:
			car.Features.SafetyAndMotionControlSystem.RearViewCamera = UndefinedValue
		}

		switch storageCar.Features.SafetyAndMotionControlSystem.CruiseControl {
		case cars.Availability_YES_VALUE:
			car.Features.SafetyAndMotionControlSystem.CruiseControl = YesValue
		case cars.Availability_NO_VALUE:
			car.Features.SafetyAndMotionControlSystem.CruiseControl = NoValue
		case cars.Availability_OPTION_VALUE:
			car.Features.SafetyAndMotionControlSystem.CruiseControl = OptionValue
		case cars.Availability_UNDEFINED_VALUE:
			car.Features.SafetyAndMotionControlSystem.CruiseControl = UndefinedValue
		}

		storageCar.Features.Lights.Headlights = car.Features.Lights.Headlights
		switch storageCar.Features.Lights.LEDRunningLights {
		case cars.Availability_YES_VALUE:
			car.Features.Lights.LEDRunningLights = YesValue
		case cars.Availability_NO_VALUE:
			car.Features.Lights.LEDRunningLights = NoValue
		case cars.Availability_OPTION_VALUE:
			car.Features.Lights.LEDRunningLights = OptionValue
		case cars.Availability_UNDEFINED_VALUE:
			car.Features.Lights.LEDRunningLights = UndefinedValue
		}

		switch storageCar.Features.Lights.LEDTailLights {
		case cars.Availability_YES_VALUE:
			car.Features.Lights.LEDTailLights = YesValue
		case cars.Availability_NO_VALUE:
			car.Features.Lights.LEDTailLights = NoValue
		case cars.Availability_OPTION_VALUE:
			car.Features.Lights.LEDTailLights = OptionValue
		case cars.Availability_UNDEFINED_VALUE:
			car.Features.Lights.LEDTailLights = UndefinedValue
		}

		switch storageCar.Features.Lights.LightSensor {
		case cars.Availability_YES_VALUE:
			car.Features.Lights.LightSensor = YesValue
		case cars.Availability_NO_VALUE:
			car.Features.Lights.LightSensor = NoValue
		case cars.Availability_OPTION_VALUE:
			car.Features.Lights.LightSensor = OptionValue
		case cars.Availability_UNDEFINED_VALUE:
			car.Features.Lights.LightSensor = UndefinedValue
		}

		switch storageCar.Features.Lights.FrontFogLights {
		case cars.Availability_YES_VALUE:
			car.Features.Lights.FrontFogLights = YesValue
		case cars.Availability_NO_VALUE:
			car.Features.Lights.FrontFogLights = NoValue
		case cars.Availability_OPTION_VALUE:
			car.Features.Lights.FrontFogLights = OptionValue
		case cars.Availability_UNDEFINED_VALUE:
			car.Features.Lights.FrontFogLights = UndefinedValue
		}

		switch storageCar.Features.Lights.BackFogLights {
		case cars.Availability_YES_VALUE:
			car.Features.Lights.BackFogLights = YesValue
		case cars.Availability_NO_VALUE:
			car.Features.Lights.BackFogLights = NoValue
		case cars.Availability_OPTION_VALUE:
			car.Features.Lights.BackFogLights = OptionValue
		case cars.Availability_UNDEFINED_VALUE:
			car.Features.Lights.BackFogLights = UndefinedValue
		}

		car.Features.Interior.Upholstery = storageCar.Features.Interior.Upholstery

		switch storageCar.Features.CabinMicroclimate.AirConditioner {
		case cars.Availability_YES_VALUE:
			car.Features.CabinMicroclimate.AirConditioner = YesValue
		case cars.Availability_NO_VALUE:
			car.Features.CabinMicroclimate.AirConditioner = NoValue
		case cars.Availability_OPTION_VALUE:
			car.Features.CabinMicroclimate.AirConditioner = OptionValue
		case cars.Availability_UNDEFINED_VALUE:
			car.Features.CabinMicroclimate.AirConditioner = UndefinedValue
		}

		switch storageCar.Features.CabinMicroclimate.ClimateControl {
		case cars.Availability_YES_VALUE:
			car.Features.CabinMicroclimate.ClimateControl = YesValue
		case cars.Availability_NO_VALUE:
			car.Features.CabinMicroclimate.ClimateControl = NoValue
		case cars.Availability_OPTION_VALUE:
			car.Features.CabinMicroclimate.ClimateControl = OptionValue
		case cars.Availability_UNDEFINED_VALUE:
			car.Features.CabinMicroclimate.ClimateControl = UndefinedValue
		}

		switch storageCar.Features.ElectricOptions.ElectricFrontSideWindowsLifts {
		case cars.Availability_YES_VALUE:
			car.Features.ElectricOptions.ElectricFrontSideWindowsLifts = YesValue
		case cars.Availability_NO_VALUE:
			car.Features.ElectricOptions.ElectricFrontSideWindowsLifts = NoValue
		case cars.Availability_OPTION_VALUE:
			car.Features.ElectricOptions.ElectricFrontSideWindowsLifts = OptionValue
		case cars.Availability_UNDEFINED_VALUE:
			car.Features.ElectricOptions.ElectricFrontSideWindowsLifts = UndefinedValue
		}

		switch storageCar.Features.ElectricOptions.ElectricBackSideWindowsLifts {
		case cars.Availability_YES_VALUE:
			car.Features.ElectricOptions.ElectricBackSideWindowsLifts = YesValue
		case cars.Availability_NO_VALUE:
			car.Features.ElectricOptions.ElectricBackSideWindowsLifts = NoValue
		case cars.Availability_OPTION_VALUE:
			car.Features.ElectricOptions.ElectricBackSideWindowsLifts = OptionValue
		case cars.Availability_UNDEFINED_VALUE:
			car.Features.ElectricOptions.ElectricBackSideWindowsLifts = UndefinedValue
		}

		switch storageCar.Features.ElectricOptions.ElectricHeatingOfFrontSeats {
		case cars.Availability_YES_VALUE:
			car.Features.ElectricOptions.ElectricHeatingOfFrontSeats = YesValue
		case cars.Availability_NO_VALUE:
			car.Features.ElectricOptions.ElectricHeatingOfFrontSeats = NoValue
		case cars.Availability_OPTION_VALUE:
			car.Features.ElectricOptions.ElectricHeatingOfFrontSeats = OptionValue
		case cars.Availability_UNDEFINED_VALUE:
			car.Features.ElectricOptions.ElectricHeatingOfFrontSeats = UndefinedValue
		}

		switch storageCar.Features.ElectricOptions.ElectricHeatingOfBackSeats {
		case cars.Availability_YES_VALUE:
			car.Features.ElectricOptions.ElectricHeatingOfBackSeats = YesValue
		case cars.Availability_NO_VALUE:
			car.Features.ElectricOptions.ElectricHeatingOfBackSeats = NoValue
		case cars.Availability_OPTION_VALUE:
			car.Features.ElectricOptions.ElectricHeatingOfBackSeats = OptionValue
		case cars.Availability_UNDEFINED_VALUE:
			car.Features.ElectricOptions.ElectricHeatingOfBackSeats = UndefinedValue
		}

		switch storageCar.Features.ElectricOptions.ElectricHeatingOfSteeringWheel {
		case cars.Availability_YES_VALUE:
			car.Features.ElectricOptions.ElectricHeatingOfSteeringWheel = YesValue
		case cars.Availability_NO_VALUE:
			car.Features.ElectricOptions.ElectricHeatingOfSteeringWheel = NoValue
		case cars.Availability_OPTION_VALUE:
			car.Features.ElectricOptions.ElectricHeatingOfSteeringWheel = OptionValue
		case cars.Availability_UNDEFINED_VALUE:
			car.Features.ElectricOptions.ElectricHeatingOfSteeringWheel = UndefinedValue
		}

		switch storageCar.Features.ElectricOptions.ElectricHeatingOfWindshield {
		case cars.Availability_YES_VALUE:
			car.Features.ElectricOptions.ElectricHeatingOfWindshield = YesValue
		case cars.Availability_NO_VALUE:
			car.Features.ElectricOptions.ElectricHeatingOfWindshield = NoValue
		case cars.Availability_OPTION_VALUE:
			car.Features.ElectricOptions.ElectricHeatingOfWindshield = OptionValue
		case cars.Availability_UNDEFINED_VALUE:
			car.Features.ElectricOptions.ElectricHeatingOfWindshield = UndefinedValue
		}

		switch storageCar.Features.ElectricOptions.ElectricHeatingOfRearWindow {
		case cars.Availability_YES_VALUE:
			car.Features.ElectricOptions.ElectricHeatingOfRearWindow = YesValue
		case cars.Availability_NO_VALUE:
			car.Features.ElectricOptions.ElectricHeatingOfRearWindow = NoValue
		case cars.Availability_OPTION_VALUE:
			car.Features.ElectricOptions.ElectricHeatingOfRearWindow = OptionValue
		case cars.Availability_UNDEFINED_VALUE:
			car.Features.ElectricOptions.ElectricHeatingOfRearWindow = UndefinedValue
		}

		switch storageCar.Features.ElectricOptions.ElectricHeatingOfSideMirrors {
		case cars.Availability_YES_VALUE:
			car.Features.ElectricOptions.ElectricHeatingOfSideMirrors = YesValue
		case cars.Availability_NO_VALUE:
			car.Features.ElectricOptions.ElectricHeatingOfSideMirrors = NoValue
		case cars.Availability_OPTION_VALUE:
			car.Features.ElectricOptions.ElectricHeatingOfSideMirrors = OptionValue
		case cars.Availability_UNDEFINED_VALUE:
			car.Features.ElectricOptions.ElectricHeatingOfSideMirrors = UndefinedValue
		}

		switch storageCar.Features.ElectricOptions.ElectricDriveOfDriverSeat {
		case cars.Availability_YES_VALUE:
			car.Features.ElectricOptions.ElectricDriveOfDriverSeat = YesValue
		case cars.Availability_NO_VALUE:
			car.Features.ElectricOptions.ElectricDriveOfDriverSeat = NoValue
		case cars.Availability_OPTION_VALUE:
			car.Features.ElectricOptions.ElectricDriveOfDriverSeat = OptionValue
		case cars.Availability_UNDEFINED_VALUE:
			car.Features.ElectricOptions.ElectricDriveOfDriverSeat = UndefinedValue
		}

		switch storageCar.Features.ElectricOptions.ElectricDriveOfFrontSeats {
		case cars.Availability_YES_VALUE:
			car.Features.ElectricOptions.ElectricDriveOfFrontSeats = YesValue
		case cars.Availability_NO_VALUE:
			car.Features.ElectricOptions.ElectricDriveOfFrontSeats = NoValue
		case cars.Availability_OPTION_VALUE:
			car.Features.ElectricOptions.ElectricDriveOfFrontSeats = OptionValue
		case cars.Availability_UNDEFINED_VALUE:
			car.Features.ElectricOptions.ElectricDriveOfFrontSeats = UndefinedValue
		}

		switch storageCar.Features.ElectricOptions.ElectricDriveOfSideMirrors {
		case cars.Availability_YES_VALUE:
			car.Features.ElectricOptions.ElectricDriveOfSideMirrors = YesValue
		case cars.Availability_NO_VALUE:
			car.Features.ElectricOptions.ElectricDriveOfSideMirrors = NoValue
		case cars.Availability_OPTION_VALUE:
			car.Features.ElectricOptions.ElectricDriveOfSideMirrors = OptionValue
		case cars.Availability_UNDEFINED_VALUE:
			car.Features.ElectricOptions.ElectricDriveOfSideMirrors = UndefinedValue
		}

		switch storageCar.Features.ElectricOptions.ElectricTrunkOpener {
		case cars.Availability_YES_VALUE:
			car.Features.ElectricOptions.ElectricTrunkOpener = YesValue
		case cars.Availability_NO_VALUE:
			car.Features.ElectricOptions.ElectricTrunkOpener = NoValue
		case cars.Availability_OPTION_VALUE:
			car.Features.ElectricOptions.ElectricTrunkOpener = OptionValue
		case cars.Availability_UNDEFINED_VALUE:
			car.Features.ElectricOptions.ElectricTrunkOpener = UndefinedValue
		}

		switch storageCar.Features.ElectricOptions.RainSensor {
		case cars.Availability_YES_VALUE:
			car.Features.ElectricOptions.RainSensor = YesValue
		case cars.Availability_NO_VALUE:
			car.Features.ElectricOptions.RainSensor = NoValue
		case cars.Availability_OPTION_VALUE:
			car.Features.ElectricOptions.RainSensor = OptionValue
		case cars.Availability_UNDEFINED_VALUE:
			car.Features.ElectricOptions.RainSensor = UndefinedValue
		}

		switch storageCar.Features.Airbags.DriverAirbag {
		case cars.Availability_YES_VALUE:
			car.Features.Airbags.DriverAirbag = YesValue
		case cars.Availability_NO_VALUE:
			car.Features.Airbags.DriverAirbag = NoValue
		case cars.Availability_OPTION_VALUE:
			car.Features.Airbags.DriverAirbag = OptionValue
		case cars.Availability_UNDEFINED_VALUE:
			car.Features.Airbags.DriverAirbag = UndefinedValue
		}

		switch storageCar.Features.Airbags.FrontPassengerAirbag {
		case cars.Availability_YES_VALUE:
			car.Features.Airbags.FrontPassengerAirbag = YesValue
		case cars.Availability_NO_VALUE:
			car.Features.Airbags.FrontPassengerAirbag = NoValue
		case cars.Availability_OPTION_VALUE:
			car.Features.Airbags.FrontPassengerAirbag = OptionValue
		case cars.Availability_UNDEFINED_VALUE:
			car.Features.Airbags.FrontPassengerAirbag = UndefinedValue
		}

		switch storageCar.Features.Airbags.SideAirbags {
		case cars.Availability_YES_VALUE:
			car.Features.Airbags.SideAirbags = YesValue
		case cars.Availability_NO_VALUE:
			car.Features.Airbags.SideAirbags = NoValue
		case cars.Availability_OPTION_VALUE:
			car.Features.Airbags.SideAirbags = OptionValue
		case cars.Availability_UNDEFINED_VALUE:
			car.Features.Airbags.SideAirbags = UndefinedValue
		}

		switch storageCar.Features.Airbags.CurtainAirbags {
		case cars.Availability_YES_VALUE:
			car.Features.Airbags.CurtainAirbags = YesValue
		case cars.Availability_NO_VALUE:
			car.Features.Airbags.CurtainAirbags = NoValue
		case cars.Availability_OPTION_VALUE:
			car.Features.Airbags.CurtainAirbags = OptionValue
		case cars.Availability_UNDEFINED_VALUE:
			car.Features.Airbags.CurtainAirbags = UndefinedValue
		}

		switch storageCar.Features.MultimediaSystems.HandsFreeSupport {
		case cars.Availability_YES_VALUE:
			car.Features.MultimediaSystems.HandsFreeSupport = YesValue
		case cars.Availability_NO_VALUE:
			car.Features.MultimediaSystems.HandsFreeSupport = NoValue
		case cars.Availability_OPTION_VALUE:
			car.Features.MultimediaSystems.HandsFreeSupport = OptionValue
		case cars.Availability_UNDEFINED_VALUE:
			car.Features.MultimediaSystems.HandsFreeSupport = UndefinedValue
		}

		switch storageCar.Features.MultimediaSystems.MP3Support {
		case cars.Availability_YES_VALUE:
			car.Features.MultimediaSystems.MP3Support = YesValue
		case cars.Availability_NO_VALUE:
			car.Features.MultimediaSystems.MP3Support = NoValue
		case cars.Availability_OPTION_VALUE:
			car.Features.MultimediaSystems.MP3Support = OptionValue
		case cars.Availability_UNDEFINED_VALUE:
			car.Features.MultimediaSystems.MP3Support = UndefinedValue
		}

		switch storageCar.Features.MultimediaSystems.OnBoardComputer {
		case cars.Availability_YES_VALUE:
			car.Features.MultimediaSystems.OnBoardComputer = YesValue
		case cars.Availability_NO_VALUE:
			car.Features.MultimediaSystems.OnBoardComputer = NoValue
		case cars.Availability_OPTION_VALUE:
			car.Features.MultimediaSystems.OnBoardComputer = OptionValue
		case cars.Availability_UNDEFINED_VALUE:
			car.Features.MultimediaSystems.OnBoardComputer = UndefinedValue
		}

		switch storageCar.Features.CarAlarm {
		case cars.Availability_YES_VALUE:
			car.Features.CarAlarm = YesValue
		case cars.Availability_NO_VALUE:
			car.Features.CarAlarm = NoValue
		case cars.Availability_OPTION_VALUE:
			car.Features.CarAlarm = OptionValue
		case cars.Availability_UNDEFINED_VALUE:
			car.Features.CarAlarm = UndefinedValue
		}

		car.Features.Color = storageCar.Features.Color

		car.Offering.Price = storageCar.Offering.Price
		car.Offering.Year = int(storageCar.Offering.Year)
		car.Offering.Kilometrage = storageCar.Offering.Kilometrage
		car.Offering.PhotoURLs = storageCar.Offering.PhotoURLs
		crs[idx] = &car
	}
	return crs
}
