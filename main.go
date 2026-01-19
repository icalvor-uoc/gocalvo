package main

import "github.com/icalvor-uoc/gocalvo/package1"
import "github.com/icalvor-uoc/gocalvo/submodule"
import "github.com/icalvor-uoc/gocalvo/submodule/saycat"

func main() {
  package1.SayPackage()
  submodule.Submodule()
  saycat.Saycat()
}
