package conv

func FahrenheitToCelsius(fahrenheit float64) float64 {

	celsius := (fahrenheit - 32) * 5 / 9

	// Her skal du implementere funksjonen
	// Du skal ikke formattere float64 i denne funksjonen
	// Gjør formattering i main.go med fmt.Printf eller
	// lignende
	return celsius
}

func FahrenheitToKelvin(Fahrenheit float64) float64 {

	Kelvin := (Fahrenheit-32)*5/9 + 273.15

	return Kelvin
}

// //c
func CelsiusToFahrenheit(Celsius float64) float64 {

	fahrenheit := Celsius*9/5 + 32

	return fahrenheit
}

func CelsiusToKelvin(Celsius float64) float64 {
	Kelvin := Celsius + 273.15

	return Kelvin
}

// k
func KelvinToCelsius(Kelvin float64) float64 {
	Celsius := Kelvin - 273.15

	return Celsius
}

func KelvinToFahrenheit(Kelvin float64) float64 {
	Fahrenheit := (Kelvin-273.15)*9/5 + 32

	return Fahrenheit
}

func GetSunTemp() string {
	return "Temperatur på ytre lag av Solen 5506.85°C.\nTemperatur i Solens kjerne er 15 000 000°C."
}
