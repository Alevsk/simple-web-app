package routes

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Item struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	Price    int    `json:"price"`
	Id       string `json:"id"`
}

type ItemListPrice struct {
	Items []Item `json:"item"`
	Total int    `json:"total"`
}

var itemsBunch ItemListPrice

func add(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if r.Method == "POST" {
		w.WriteHeader(http.StatusOK)
		bodyBytes, _ := ioutil.ReadAll(r.Body)
		var element Item
		err := json.Unmarshal(bodyBytes, &element)
		if err != nil {
			log.Println(err)
		}
		itemsBunch.Items = append(itemsBunch.Items, element)
		totalProduct := productoTotalPrice(element.Price, element.Quantity)
		itemsBunch.Total = newTotal(itemsBunch.Total, totalProduct)
		if err := json.NewEncoder(w).Encode(itemsBunch); err != nil {
			panic(err)
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}

}

func edit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if r.Method == "POST" {
		w.WriteHeader(http.StatusOK)
		bodyBytes, _ := ioutil.ReadAll(r.Body)
		var element Item
		err := json.Unmarshal(bodyBytes, &element)
		if err != nil {
			log.Println(err)
		}
		var total int
		var tmpArray []Item
		for _, item := range itemsBunch.Items {
			if item.Id == element.Id {
				total += productoTotalPrice(element.Price, element.Quantity)
				tmpArray = append(tmpArray, element)
			} else {
				total += productoTotalPrice(item.Price, item.Quantity)
				tmpArray = append(tmpArray, item)
			}
		}
		itemsBunch.Items = tmpArray
		itemsBunch.Total = total
		if err := json.NewEncoder(w).Encode(itemsBunch); err != nil {
			panic(err)
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if r.Method == "POST" {
		w.WriteHeader(http.StatusOK)
		bodyBytes, _ := ioutil.ReadAll(r.Body)
		var element Item
		err := json.Unmarshal(bodyBytes, &element)
		if err != nil {
			log.Println(err)
		}
		var total int
		var tmpArray []Item
		for _, item := range itemsBunch.Items {
			if item.Id != element.Id {
				total += productoTotalPrice(item.Price, item.Quantity)
				tmpArray = append(tmpArray, item)
			}
		}
		itemsBunch.Items = tmpArray
		itemsBunch.Total = total
		if err := json.NewEncoder(w).Encode(itemsBunch); err != nil {
			panic(err)
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func productoTotalPrice(productPrice, quantity int) int {
	return productPrice * quantity
}

func newTotal(total, newElement int) int {
	return total + newElement
}
