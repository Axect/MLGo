package main

import (
	"fmt"
	"log"
	"os"

	df "github.com/kniren/gota/dataframe"
	sr "github.com/kniren/gota/series"
)

func main() {
	irisFile, err := os.Open("../iris/iris.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer irisFile.Close()

	irisDF := df.ReadCSV(irisFile)
	fmt.Println(irisDF)

	filter := df.F{
		Colname:    "species",
		Comparator: sr.Eq,
		Comparando: "Iris-versicolor",
	}

	versicolorDF := irisDF.Filter(filter)
	if versicolorDF.Err != nil {
		log.Fatal(versicolorDF.Err)
	}

	fmt.Println(versicolorDF)

	versicolorDF = irisDF.Filter(filter).Select([]string{"sepal_width", "species"})
	fmt.Println(versicolorDF)

	versicolorDF = irisDF.Filter(filter).Select([]string{"sepal_width", "species"}).Subset([]int{0, 1, 2})
	fmt.Println(versicolorDF)

	return
}
