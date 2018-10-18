package main

import (
	"math/rand"
	"math"
)
func QuickSort(slice []int) []int {
	length := len(slice)

	if length <= 1 {
		sliceCopy := make([]int, length)
		copy(sliceCopy, slice)
		return sliceCopy
	}

	m := slice[rand.Intn(length)]

	less := make([]int, 0, length)
	middle := make([]int, 0, length)
	more := make([]int, 0, length)

	for _, item := range slice {
		switch {
		case item < m:
			less = append(less, item)
		case item == m:
			middle = append(middle, item)

		case item > m:
			more = append(more, item)

		}
	}

	less, more = QuickSort(less), QuickSort(more)

	less = append(less, middle...)
	less = append(less, more...)

	return less
}

func StandardQuickSort(slice []int,left,right int) ([]int,int,int){
	Cn:=0
	Sn:=0
	if right>left{
		pivot:=slice[right]
		i,j:=left,right
		for i<=j{
			for slice[i]<pivot{
				i++
				Cn++
			}
			for slice[j]>pivot{
				j--
				Cn++
			}
			Cn+=2
			if i<=j{
				slice[i],slice[j] = slice[j],slice[i]
				Sn++
				i++
				j--
			}

		}
		_,Co,So := StandardQuickSort(slice,left,j)
		_,Ci,Si := StandardQuickSort(slice,i,right)
		Cn+=Co+Ci
		Sn+=So+Si
	}
	return slice,Cn,Sn
}
func makeRange(min, max int) []float64 {
	a := make([]float64, max-min+1)
	for i := range a {
		a[i] = float64(min + i)
	}
	return a
}
func makeX(slice []float64,step int) []float64{
	x:=make([]float64,len(slice))
	steps:=step
		for i:= range slice{
			x[i]=float64(steps)
			steps+=step
		}
		return x
}
func TheorySwapList(slice []float64) []float64{
	x:=make([]float64,len(slice))
	for i:=range slice{
		x[i]=0.33*slice[i]*math.Log(slice[i])-0.58*slice[i]+math.Log(slice[i])
	}
	return x
}
func meanfromSlice(slice []int)float64{
	var total float64 = 0
	for _,value := range slice{
		total+=float64(value)

	}
	return total/float64(len(slice))
}

func VAriancefromSLice(slice []int) float64{
	mean := meanfromSlice(slice)
	var total float64 = 0
	for _,value := range slice{
		total+=math.Pow((float64(value)-float64(mean)),2)
	}
	return total/float64(len(slice))
}
func TestQuiclkSort(testnumber int, maxarraysize int) ([]float64,[]float64){
	listCn:=make([]float64,0)
	listSn:=make([]float64,0)
	for i:=10;i<=maxarraysize;i+=10{
		templistCn:=make([]int,testnumber)
		templistSn:=make([]int,testnumber)
		for j:=0;j<testnumber;j++ {
			list := rand.Perm(i)
			for k, _ := range list {
				list[k]++
			}
			_,templistCn[j],templistSn[j] = StandardQuickSort(list,0,len(list)-1)

		}
		listCn = append(listCn,meanfromSlice(templistCn))
		listSn = append(listSn,meanfromSlice(templistSn))

	}
	return listCn,listSn
}