package conv

/*
  I denne pakken skal alle konverteringfunksjonene
  implementeres. Bruk engelsk.
    FarhenheitToCelsius
    CelsiusToFahrenheit
    KelvinToFarhenheit
    ...
*/

// Konverterer Farhenheit til Celsius
func FahrenheitToCelsius(value float64) float64 {
	return (value - 32)*(5/9)
      // Her skal du implementere funksjonen
	// Du skal ikke formattere float64 i denne funksjonen
	// Gjør formattering i main.go med fmt.Printf eller
	// lignende
	   
}

func FahrenheitToKelvin(value float64) float64 {
        return (value - 32) * (5/9) + 273.15
          
}

func CelsiusToFahrenheit(value float64) float64 {
       return (value*(9/5) + 32)
       
}

func CelsiusToKelvin(value float64) float64 {
       return (value + 273.15)
          
}

func KelvinToFahrenheit(value float64) float64 {
       return (value - 273.15)*(9/5) + 32
           
}

func KelvinToCelsius(value float64) float64 {
         return (value - 273.15)
           
}

// De andre konverteringsfunksjonene implementere her
// ...
