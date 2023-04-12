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
func FarhenheitToCelsius(value float64) float64 {
	return (value - 32)*(5/9)
      // Her skal du implementere funksjonen
	// Du skal ikke formattere float64 i denne funksjonen
	// Gjør formattering i main.go med fmt.Printf eller
	// lignende
	   
}

func FarhenheitToKelvin(value float64) float64 {
        return (value - 32) * (5/9) + 273.15
          
}

func CelsiusToFarhenheit(value float64) float64 {
       return (value*(9/5) + 32)
       
}

func CelsiusToKelvin(value float64) float64 {
       return (value + 273.15)
          
}

func KelvinToFarhenheit(value float64) float64 {
       return (value - 273.15)*(9/5) + 32
           
}

func KelvinToCelsius(value float64) float64 {
         return (value - 273.15)
           
}

// De andre konverteringsfunksjonene implementere her
// ...
