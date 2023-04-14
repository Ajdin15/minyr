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
//f _, err := os.Stat("kjevik-temp-fahr-20220318-20230318.csv"); err == nil {
    fmt.Print("Filen kjevik-temp-fahr-20220318-20230318.csv finnes allerede. Vil du generere den på nytt? (j/n):")
    scanner.Scan()
    answer := strings.ToLower(scanner.Text())
    if answer != "j" && answer != "n" {
        log.Fatal("Ugyldig svar")
    } else if answer == "n" {
        return
    }
}
 
convertedTemperatures, err := ConvertTemperatures()
if err != nil {
    log.Fatal(err)
}

if err := WriteLines(convertedTemperatures, "kjevik-temp-fahr-20220318-20230318.csv"); err != nil {
    log.Fatal(err)
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

func WriteLines(lines []string, filename string) error {
file, err := os.Create(filename)
if err != nil {
return err
}

writer := bufio.NewWriter(file)
defer writer.Flush()

fmt.Fprintln(writer, "Navn;Stasjon;Tid(norsk normaltid);Lufttemperatur (F)")

for _, line := range lines {
    fmt.Fprintln(writer, line)
}

fmt.Fprint(writer, "Data er basert på gyldig data (per 18.03.2023) (CC BY 4.0) fra Meteorologisk institutt (MET);endringen er gjort av Ajdin Smajic")

return nil
}

func ConvertTemperatures() ([]string, error) {
file, err := os.Open("kjevik-temp-celsius-20220318-20230318.csv")
if err != nil {
return nil, err
}

scanner := bufio.NewScanner(file)

ConvertedTemperatures := make([]string, 0)
for i := 0; scanner.Scan(); i++ {
    line := scanner.Text()

    if i == 0 {
        continue
    }

    fields := strings.Split(line, ";")
    if len(fields) != 4 {
        return nil, fmt.Errorf("uventet antall felt i linje %d: %d", i, len(fields))
    }

    if fields[3] == "" {
        continue
    }

    TemperatureCelsius, err := strconv.ParseFloat(fields[3], 64)

    if err != nil {
        return nil, fmt.Errorf("kunne ikke parse temperatur i linje %d: %s", i, err)
    }
    TemperatureFarhenheit := conv.CelsiusToFarhenheit(TemperatureCelsius)

    ConvertedTemperature := fmt.Sprintf("%s;%s;%.2fF", fields[0], strings.Join(fields[1:3], ";"), TemperatureFarhenheit)
    ConvertedTemperatures = append(ConvertedTemperatures, ConvertedTemperature)
}

if err := scanner.Err(); err != nil {
    return nil, err
}

return ConvertedTemperatures, nil
}
    




