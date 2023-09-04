package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	server()
}

func server() {
	http.HandleFunc("/calculatePackages", calculateHandler)
	fmt.Println("Server is starting on port 10000")
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func calculateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	dto := &CalculatePackagesDto{}
	err := json.NewDecoder(r.Body).Decode(dto)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i := 0; i < len(dto.PackageSizeList); i++ {
		if dto.PackageSizeList[i] <= 0 {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	calculatedPackages := calculatePackages(*dto)

	bytes, err1 := json.Marshal(calculatedPackages)
	if err1 != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err2 := w.Write(bytes)
	if err2 != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func calculatePackages(dto CalculatePackagesDto) CalculatedPackages {
	var calculatedPackages CalculatedPackages
	calculatedPackages.ItemsOrdered = dto.OrderedItems
	sortedPackageList := sortIntegerList(dto.PackageSizeList)
	var lastPartsPackage int32 = 0
	var remaining int32 = 0

	for i := len(sortedPackageList) - 1; i >= 0; i-- {
		packageSize := sortedPackageList[i]

		if packageSize == 0 {
			continue
		}

		remainingOrderedItems := dto.OrderedItems

		if remaining != 0 {
			remainingOrderedItems = remaining
		}

		numOfPackage := remainingOrderedItems / packageSize

		if numOfPackage > 0 && (lastPartsPackage == 0 || remainingOrderedItems == packageSize) {
			calculatedPackages.CorrectNumberOfPacks = append(calculatedPackages.CorrectNumberOfPacks,
				NumberOfPack{
					PackageSize:     packageSize,
					NumberOfPackage: numOfPackage,
				})

			remaining = remainingOrderedItems - (numOfPackage * packageSize)
			lastPartsPackage = packageSize
		} else if numOfPackage == 0 {
			lastPartsPackage = packageSize
		}
	}

	if len(calculatedPackages.CorrectNumberOfPacks) == 0 || remaining != 0 {
		existsInList := false

		for i := 0; i < len(calculatedPackages.CorrectNumberOfPacks); i++ {
			packNo := calculatedPackages.CorrectNumberOfPacks[i].PackageSize
			if packNo == lastPartsPackage {
				existsInList = true
				calculatedPackages.CorrectNumberOfPacks[i].NumberOfPackage = calculatedPackages.CorrectNumberOfPacks[i].NumberOfPackage + 1
				break
			}
		}
		if !existsInList {
			calculatedPackages.CorrectNumberOfPacks = append(calculatedPackages.CorrectNumberOfPacks,
				NumberOfPack{
					PackageSize:     lastPartsPackage,
					NumberOfPackage: 1,
				})
		}
	}

	return calculatedPackages
}

func sortIntegerList(list []int32) []int32 {
	if len(list) == 0 {
		return list
	}

	for i := 0; i < len(list)-1; i++ {
		for j := 0; j < len(list)-i-1; j++ {
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
			}
		}
	}
	return list
}

/**************
* Models
***************/

type CalculatePackagesDto struct {
	PackageSizeList []int32 `json:"packageSizeList"`
	OrderedItems    int32   `json:"orderedItems"`
}

type NumberOfPack struct {
	PackageSize     int32 `json:"packageSize"`
	NumberOfPackage int32 `json:"numberOfPackage"`
}

type CalculatedPackages struct {
	ItemsOrdered         int32          `json:"itemsOrdered"`
	CorrectNumberOfPacks []NumberOfPack `json:"correctNumberOfPacks"`
}
