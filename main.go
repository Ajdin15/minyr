package main

import (
    "fmt"
    "github.com/Ajdin15/funtemps/conv"
    "io"
    "log"
    "os"
    "strings"
    "strconv"
    "bufio"
    "encoding/csv"
    
   
    
        
    
)

func main() {
    src, err := os.Open("table.csv")
    //src, err := os.Open("/home/ajdins/minyr/kjevik-temp-celsius-20220318-20230318.csv")
    if err != nil {
        log.Fatal(err)
    }
    defer src.Close()
    log.Println(src)

    var buffer []byte
    var linebuf []byte // nil
    buffer = make([]byte, 1)
    bytesCount := 0
    for {
        _, err := src.Read(buffer)
        if err != nil && err != io.EOF {
            log.Fatal(err)
        }

        bytesCount++
        //log.Printf("%c ", buffer[:n])
        if buffer[0] == 0x0A {
            log.Println(string(linebuf))
            // Her
            elementArray := strings.Split(string(linebuf), ";")
            if len(elementArray) > 3 {
                celsius := elementArray[3]
                var fahr = CelsiusToFarhenheitString(celsius)
                log.Println(elementArray[3])
                fmt.Println(fahr)
            }
            linebuf = nil
        } else {
            linebuf = append(linebuf, buffer[0])
        }
        //log.Println(string(linebuf))
        if err == io.EOF {
            break
        }
var input string
scanner := bufio.NewScanner(os.Stdin)

for scanner.Scan() {
    input = scanner.Text()
    if input == "q" || input == "exit" {
        fmt.Println("exit")
        os.Exit(0)
    } else if input == "convert" {
        fmt.Println("Konverterer alle målingene gitt i grader Celsius til grader Fahrenheit.")
        // funksjon som åpner fil, leser linjer, gjør endringer og lagrer nye linjer i en ny fil

    // flere else-if setninger     
    } else {
        fmt.Println("Venligst velg convert, average eller exit:")

    }
  fd, error := os.Open("kjevik-temp-celsius-20220318-20230318.csv")
  if error != nil {
    fmt.Println(error)
  }
  fmt.Println("Successfully opened the CSV file")
  defer fd.Close()
    
  // read CSV file
  fileReader := csv.NewReader(fd)
  records, error := fileReader.ReadAll()
  if error != nil {
      fmt.Println(error)
  }
  fmt.Println(records)

}

    }
}

func CelsiusToFarhenheitString(celsius string)  string { 
	var fahrFloat float64
      	if celsiusFloat, err := strconv.ParseFloat(celsius, 64); err == nil {
		fahrFloat = conv.CelsiusToFarhenheit(celsiusFloat)
	}
	fahrString := fmt.Sprintf("%.1f", fahrFloat)
	return fahrString 
}
    




