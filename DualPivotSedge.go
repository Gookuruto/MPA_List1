package main

import ("math"
"math/rand")

func DualPivotSedge(A []int,left,right int)([]int,int,int)  {
	Sn:=0
	Cn:=0
if left-right>1{
	i:=left
	i1:=left
	j:=right
	j1:=right
	p:=A[left]
	q:=A[right]
	if p>q{
		q, p = p,q
		Sn++

	}
	Cn++
	for{
		i++
		for A[i]<=q{
			Cn++
			if i>=j{
				break
			}
			if A[i]<p{
				Cn++
				A[i1],A[i]=A[i],A[i1+1]
				Sn++
				i1++
			}
			i++
		}
		for A[j]>=p{
			Cn++
			if A[j]>q{
				Cn++
				A[j1],A[j]=A[j],A[j1-1]
				Sn++
				j1--
			}
			if i>=j{
				break
			}

		}
		A[i1]=A[j]
		A[j1]=A[i]
		i1++
		j1--
		A[i] = A[i1]
		A[j]=A[j1]
		Sn+=2


	}
	A[i1]=p
	A[j1]=q
	DualPivotSedge(A,left,i1-1)
	DualPivotSedge(A,i1+1,j1-1)
	DualPivotSedge(A,j1+1,right)



}

return A,Sn,Cn
}

func TestDualPivotsedge(testnumber int, maxarraysize int) ([]float64,[]float64){
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
			_,templistCn[j],templistSn[j] = DualPivotSedge(list,0,len(list)-1)

		}
		listCn = append(listCn,meanfromSlice(templistCn))
		listSn = append(listSn,meanfromSlice(templistSn))

	}
	return listCn,listSn
}

func TheorySwapPivotSedg(slice []float64) []float64{
	x:=make([]float64,len(slice))
	for i:=range slice{
		x[i]=0.6*slice[i]*math.Log(slice[i])+0.08*slice[i]+math.Log(slice[i])
	}
	return x
}
