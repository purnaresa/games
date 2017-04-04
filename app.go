package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

func main() {
	var tes float64 = 12321

	nums := []float64{}
	for index := 1; index < 100000; index = index * 10 {
		indexFloat := float64(index)
		num := math.Mod(math.Floor(tes/indexFloat), 10)

		if num != 0 {
			nums = append(nums, num)
		}
	}
	log.Println(nums)
	lengthSlice := len(nums)
	result := false
	for index, val := range nums {
		reverseIndex := lengthSlice - 1 - index
		log.Printf("%v <--> %v ", val, nums[reverseIndex])

		if val == nums[reverseIndex] {
			result = true
		} else {
			result = false
			log.Println(result)
			return
		}
	}
	log.Println(result)
	// palindrome()
	// magicScure()
}

func palindrome() {
	tes := 99022099
	tesString := strconv.Itoa(tes)
	tesSlice := strings.Split(tesString, "")
	lengthSlice := len(tesSlice)

	result := false
	for index, val := range tesSlice {
		reverseIndex := lengthSlice - 1 - index
		log.Printf("%s <--> %s ", val, tesSlice[reverseIndex])

		if val == tesSlice[reverseIndex] {
			result = true
		} else {
			result = false
			log.Println(result)
			return
		}
	}
	log.Println(result)
}

func magicScure() {
	result := [9][9]int{}
	radius := 9

	num := []int{}
	for index := 1; index <= radius*radius; index++ {
		num = append(num, index)
	}
	row := 0
	col := 0
	for _, val := range num {
		result[row][col] = val
		if row == radius-1 {
			row = 0
		} else {
			row++
		}
		if col == radius-1 {
			col = 0
		} else {
			col++
		}
		if result[row][col] != 0 {
			row++
		}
	}

	for _, val := range result {
		sum := 0
		for _, v := range val {
			fmt.Printf("%02d ", v)
			sum = sum + v
		}
		fmt.Printf("==> %d \n", sum)
	}

}
