package main

import "fmt"

func sumaACero(nums []int) []int {
	if nums == nil {
		return nil
	}

	var result []int
	i := 0
	for i < len(nums) {
		sum := 0
		resultaCero := false
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum == 0 {
				// Si encontramos una suma de cero, omitimos esta secuencia
				i = j + 1
				resultaCero = true
				break
			}
		}
		// Si no encontramos una suma de cero, añadimos nums[i] al resultado
		if !resultaCero {
			result = append(result, nums[i])
			i++
		}
	}
	return result
}

func main() {
	nums := []int{3, 4, -7, 5, -6, 2, 5, -1, 8}
	result := sumaACero(nums)
	fmt.Println(result) // Esperado: [5, 8]
}
