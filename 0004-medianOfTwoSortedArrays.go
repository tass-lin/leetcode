package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums := combine(nums1,nums2)
	return medianOf(nums)
}

func combine(mis, njs []int) []int {
	lenMis, i := len(mis), 0
	lenNjs, j := len(njs), 0
	res := make([]int, lenMis+lenNjs)

	for k:=0;k < lenMis+lenNjs ;k++  {
		if i == lenMis || (i < lenMis && j < lenNjs && mis[i] > njs[j] ) {
			res[k] = njs[j]
			j++

			continue
		}

		if j == lenNjs || (i < lenMis && j < lenNjs && mis[i] <= njs[j] ) {
			res[k] = mis[i]
			i++
		}
	}
	return res
}

func medianOf(num []int) float64 {
	l := len(num)

	if l == 0 {
		panic("長度為0,無法執行")
	}

	if l%2 == 0 {
		return float64(num[l/2]+ num[l/2-1]) / 2.0
	}
	return float64(num[l/2])
}