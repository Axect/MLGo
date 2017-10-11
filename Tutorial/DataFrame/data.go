package main

import (
	"fmt"
	"log"

	"github.com/Axect/csv"
	df "github.com/kniren/gota/dataframe"
)

func main() {
	irisCSV := csv.Read("../iris/iris.csv")
	irisData := make([][]string, len(irisCSV)+1, len(irisCSV)+1)

	irisData[0] = []string{"sepal_length", "sepal_width", "petal_length", "petal_width", "species"}
	for i := range irisCSV {
		irisData[i+1] = irisCSV[i]
	}

	irisDF := df.LoadRecords(
		irisData,
	)

	fmt.Println(irisDF)

	filter := df.F{
		Colname:    "species",
		Comparator: "==",
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
