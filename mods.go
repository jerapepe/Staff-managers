package main

import(
	"fmt"
	"github.com/xuri/excelize/v2"
)

func getRowss() []string{
    f, err := excelize.OpenFile("checador.xlsx")
    if err != nil {
        fmt.Println(err)
    }
    //Get all the rows in the Sheet1.
    rows, err := f.GetRows("Hoja2")
    if err != nil {
        fmt.Println(err)
    }
    data := []string{}
    for _, row := range rows {
        for _, colCell := range row {
            //fmt.Print(colCell, "\t")
            data = append(data, colCell)
            //data = append(data, "\n")
        }
   }
   //fmt.Println(data)
   return data
}

func ilas(){
    f, err := excelize.OpenFile("checador.xlsx")
    if err != nil {
        fmt.Println(err)
    }
    cols, err := f.GetCols("Hoja2")
    if err != nil {
        fmt.Println(err)
        return
    }
    for _, col := range cols {
        for _, rowCell := range col {
            fmt.Print(rowCell, "\t")
        }
        fmt.Println()
    }
}

func readExcel(){
    f, err := excelize.OpenFile("checador.xlsx")
    if err != nil {
        fmt.Println(err)
        return
    }
    // Get value from cell by given worksheet name and axis.
	val := "5"
	vallor := ("A" + val)
    cell, err := f.GetCellValue("Hoja1", "C2")
	cell2, err := f.GetCellValue("Hoja1", vallor)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(cell)
	fmt.Println(cell2)
}

func getVal(sheet, cells string ) string{ 
    f, err := excelize.OpenFile("checador.xlsx")
    if err != nil {
        fmt.Println(err)
    }
	//vallor := ("A" + val)
    cell, err := f.GetCellValue(sheet, cells)
	//cell2, err := f.GetCellValue("Hoja1", vallor)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(cell)
    return cell
	//fmt.Println(cell2)
}