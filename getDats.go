package main

import (
	"log"
)

func searDat(data []string, option string)(arID []int, arDat []string){
    arID = []int{}
    arDat = []string{}
    for num, result := range data {
        if result == option {
		    arID = append(arID, num)
            arDat = append(arDat, result)
        }
    }
    return arID, arDat
}

func searchDat(data []string, option string)[]int{
    arID := []int{}
    for num, result := range data {
        if result == option {
		    arID = append(arID, num)
        }
    }
    return arID
}

func addDat(data []int, ns int)[]int{
    arIDdd := []int{}
    for _, result := range data {
		arIDdd = append(arIDdd, result+ns)
    }
    //log.Println(arIDdd)
	return arIDdd
}

func searName(id []int, data []string )[]string{
    arDat := []string{}
    for num, result := range data {
        for _, ress := range id {
			if num == ress {
				arDat = append(arDat, result)
			}
		}
    }
    //log.Println(arDat)
	return arDat
}

// func remove(slice []string, ids []int) []string {
// 	dats := []string{}
// 	for id, rs := range slice{
// 		for _, rsd := range ids{
// 			if id == rsd {
// 				dats = append(slice[:rs], slice[rs+1]...)
// 			}
// 		}
// 	}
//     return dats
// }

func searchSpace(data []string)(arID []int, arDa []string){
    arID = []int{}
	arDa = []string{}
    for num, result := range data {
        if result == "" {
			arDa = append(data[:num], data[num+1:]...) 
			arID = append(arID, num)
        }
    }
	//log.Println(arID)
	//log.Println(arDa)
    return arID, arDa
}

func removeSpace(){
	data := getRowss()
	log.Println("-------")
	searchSpace(data)
}