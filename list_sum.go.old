package main

import "fmt"
import "strconv"
import "strings"
// import "slices"
import "os"
import "bufio"

func main() {
  fmt.Println(os.Args)
  
  reader := bufio.NewReader(os.Stdin)
  var strsacc string = ""
  var continueReading bool = true
  for (continueReading) {
      linebytes, isPrefix, lineError := reader.ReadLine()
      
      fmt.Println(len(linebytes))
  
      if (lineError != nil) {
        panic(lineError)
      }
 
      var s string = string(linebytes)
      strsacc = strsacc + s
      continueReading = isPrefix
  }
  
  strs := strings.Split(strsacc, " ")

  var suma float32 = 0
  for _, n := range strs {
    f, e := strconv.ParseFloat(n, 32)
    if (e != nil) {
      fmt.Println(e)
    } else {
      suma += float32(f)
    }
  }
  
  fmt.Println("Resultado final:", suma)
}
