package main

import (
	"flag"
	"fmt"
	"github.com/vickynygaard/funtemps/conv"
	"math"
	"strconv"
)

// Definerer flag-variablene i hoved-"scope"
var fahr float64
var out string
var funfacts string
var cels float64
var kel float64
var temp string

// Bruker init (som anbefalt i dokumentasjonen) for å sikre at flagvariablene
// er initialisert.
func init() {

	/*
	   Her er eksempler på hvordan man implementerer parsing av flagg.
	   For eksempel, kommando
	       funtemps -F 0 -out C skal returnere output: 0°F er -17.78°C
	*/

	// Definerer og initialiserer flagg-variablene
	flag.Float64Var(&fahr, "F", 0.0, "temperatur i grader fahrenheit")
	// Du må selv definere flag-variablene for "C" og "K"
	flag.StringVar(&out, "out", "C", "beregne temperatur i C - celsius, F - fahrenheit, K- Kelvin")
	flag.StringVar(&funfacts, "funfacts", "sun", "\"fun-facts\" om sun - Solen, luna - Månen og terra - Jorden")
	// Du må selv definere flag-variabelen for -t flagget, som bestemmer
	// hvilken temperaturskala skal brukes når funfacts skal vises
	flag.Float64Var(&cels, "C", 0.0, "temperatur i grader Celsius")
	flag.Float64Var(&kel, "K", 0.0, "temperatur i grader kelvin")
	flag.StringVar(&temp, "t", "C", "temperatur i sola")

}

func main() {

	flag.Parse()

	/**
	    Her må logikken for flaggene og kall til funksjoner fra conv og funfacts pakkene implementeres.

	    Det er anbefalt å sette opp en tabell med alle mulige kombinasjoner
	    av flagg. flag-pakken har funksjoner som man kan bruke for å teste
	    hvor mange flagg og argumenter er spesifisert på kommandolinje.

	        fmt.Println("len(flag.Args())", len(flag.Args()))
			    fmt.Println("flag.NFlag()", flag.NFlag())

	    Enkelte kombinasjoner skal ikke være gyldige og da må kontrollstrukturer
	    brukes for å utelukke ugyldige kombinasjoner:
	    -F, -C, -K kan ikke brukes samtidig disse tre kan brukes med -out, men ikke med -funfacts
	    -funfacts kan brukes kun med -t
	    ...
	    Jobb deg gjennom alle tilfellene. Vær obs på at det er en del sjekk
	    implementert i flag-pakken og at den vil skrive ut "Usage" med
	    beskrivelsene av flagg-variablene, som angitt i parameter fire til
	    funksjonene Float64Var og StringVar
	*/

	// Her er noen eksempler du kan bruke i den manuelle testingen
	// fmt.Println(fahr, out, funfacts)

	// fmt.Println("len(flag.Args())", len(flag.Args()))
	// fmt.Println("flag.NFlag()", flag.NFlag())

	// fmt.Println(isFlagPassed("out"))

	// Eksempel på enkel logikk
	if out == "C" && isFlagPassed("F") {
		celsius := conv.FahrenheitToCelsius(fahr)
		if celsius == math.Trunc(celsius) {
			celsiusInt := int(celsius)
			celsiusStr := strconv.FormatInt(int64(celsiusInt), 10)

			// Insert thousands separator
			if len(celsiusStr) > 3 {
				thousands := celsiusStr[:len(celsiusStr)-3]
				remainder := celsiusStr[len(celsiusStr)-3:]
				celsiusStr = thousands + " " + remainder
			}

			fmt.Printf("%.2f °F er %s °C.", fahr, celsiusStr)
		} else {
			fmt.Printf("%.2f °F er %.2f °C.", fahr, celsius)
		}

	} else if out == "C" && isFlagPassed("K") {
		celsius := conv.KelvinToCelsius(kel)
		if celsius == math.Trunc(celsius) {
			celsiusInt := int(celsius)
			celsiusStr := strconv.FormatInt(int64(celsiusInt), 10)

			// Insert thousands separator
			if len(celsiusStr) > 3 {
				thousands := celsiusStr[:len(celsiusStr)-3]
				remainder := celsiusStr[len(celsiusStr)-3:]
				celsiusStr = thousands + " " + remainder
			}

			fmt.Printf("%.2f °K er %s °C.", kel, celsiusStr)
		} else {
			fmt.Printf("%.2f °K er %.2f °C.", kel, celsius)
		}

	} else if out == "F" && isFlagPassed("C") {
		farhenheit := conv.CelsiusToFarhenheit(cels)
		if farhenheit == math.Trunc(farhenheit) {
			farhenheitInt := int(farhenheit)
			farhenheitStr := strconv.FormatInt(int64(farhenheitInt), 10)

			// Insert thousands separator
			if len(farhenheitStr) > 3 {
				thousands := farhenheitStr[:len(farhenheitStr)-3]
				remainder := farhenheitStr[len(farhenheitStr)-3:]
				farhenheitStr = thousands + " " + remainder
			}

			fmt.Printf("%.2f °C er %s °F.", cels, farhenheitStr)
		} else {
			fmt.Printf("%.2f °C er %.2f °F.", cels, farhenheit)
		}

	} else if out == "F" && isFlagPassed("K") {
		farhenheit := conv.KelvinToFarhenheit(kel)
		if farhenheit == math.Trunc(farhenheit) {
			farhenheitInt := int(farhenheit)
			farhenheitStr := strconv.FormatInt(int64(farhenheitInt), 10)

			// Insert thousands separator
			if len(farhenheitStr) > 3 {
				thousands := farhenheitStr[:len(farhenheitStr)-3]
				remainder := farhenheitStr[len(farhenheitStr)-3:]
				farhenheitStr = thousands + " " + remainder
			}

			fmt.Printf("%.2f °K er %s °F.", kel, farhenheitStr)
		} else {
			fmt.Printf("%.2f °K er %.2f °F.", kel, farhenheit)
		}

	} else if out == "K" && isFlagPassed("C") {
		kelvin := conv.CelsiusToKelvin(cels)
		if kelvin == math.Trunc(kelvin) {
			kelvinInt := int(kelvin)
			kelvinStr := strconv.FormatInt(int64(kelvinInt), 10)

			// Insert thousands separator
			if len(kelvinStr) > 3 {
				thousands := kelvinStr[:len(kelvinStr)-3]
				remainder := kelvinStr[len(kelvinStr)-3:]
				kelvinStr = thousands + " " + remainder
			}

			fmt.Printf("%.2f °C er %s °K.", cels, kelvinStr)
		} else {
			fmt.Printf("%.2f °C er %.2f °K.", cels, kelvin)
		}

	} else if out == "K" && isFlagPassed("F") {
		kelvin := conv.FarhenheitToKevlin(fahr)
		if kelvin == math.Trunc(kelvin) {
			kelvinInt := int(kelvin)
			kelvinStr := strconv.FormatInt(int64(kelvinInt), 10)

			// Insert thousands separator
			if len(kelvinStr) > 3 {
				thousands := kelvinStr[:len(kelvinStr)-4]
				remainder := kelvinStr[len(kelvinStr)-4:]
				kelvinStr = thousands + " " + remainder
			}

			fmt.Printf("%.2f °F er %s °K.", fahr, kelvinStr)

		} else {
			fmt.Printf("%2.f °F er %2.f °K.", fahr, kelvin)
		}
	}

	if temp == "C" && isFlagPassed("funfacts") && funfacts != "luna" && funfacts != "terra" {
		fmt.Print("Temperatur på ytre lag av Solen 5506.85°C.\n")
		fmt.Print("Temperatur i Solens kjerne er 15 000 000°C.")
	}

	if funfacts == "luna" && isFlagPassed("funfacts") && funfacts != "terra" {
		fmt.Print("Temperatur på Månens overflate om natten -183°C\n")
		fmt.Print("Temperatur på Månens overflate om dagen 106°C")

	} else if funfacts == "terra" && isFlagPassed("funfacts") {
		fmt.Println("Høyeste temperatur målt på Jordens overflate 134F°")
		fmt.Println("Laveste temperatur målt på Jordens overflate -89.4°C")
		fmt.Println("Temperatur i Jordens indre kjerne 9392°K")

	}

}

// Funksjonen sjekker om flagget er spesifisert på kommandolinje
// Du trenger ikke å bruke den, men den kan hjelpe med logikken
func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}
