package main

import(
	"math/rand"
	"math"
)
func DualPivotYank(A []int,left,right int)([]int,int,int){
	Cn:=0
	Sn:=0
	if right-left>=1{
		p:=A[left]
		q:=A[right]
		if p>q{
			p=A[right]
			q=A[left]
			A[left],A[right] = A[right],A[left]
			Sn++
		}
		Cn++
		l:=left+1
		g:=right-1
		k:=l
		for k<=g{
			Cn++
			if A[k]<p{
				A[k],A[l] = A[l],A[k]
				Sn++
				l++
		}else{
			Cn++
			if A[k]>q {
				Cn++
				for A[g] > q && k < g {
					Cn++
					g--
				}
				A[k], A[g] = A[g], A[k]
				Sn++
				g--
				Cn++
				A[k], A[g] = A[g], A[k]
				Sn = Sn + 1
				g = g - 1
				Cn = Cn + 1
				if A[k] < p {
					A[k], A[l] = A[l], A[k]
					Sn = Sn + 1
					l = l + 1
				}
			}
		}
		k++
		}
		l--
		g++
		A[left],A[l]=A[l],A[left]
		A[right],A[g]=A[g],A[right]
		Sn+=2
		_,C0,S0:= DualPivotYank(A,left,l-1)
		_,C1,S1:=DualPivotYank(A,l+1,g-1)
		_,C2,S2:= DualPivotYank(A,g+1,right)
		Cn+=C0+C1+C2
		Sn+=S0+S1+S2

	}


	return A,Cn,Sn
}

func TestDualPivotYank(testnumber int, maxarraysize int) ([]float64,[]float64,[]float64,[]float64){
	listCn:=make([]float64,0)
	listSn:=make([]float64,0)
	listVarCNn:=make([]float64,0)
	listVarSn:=make([]float64,0)
	for i:=10;i<=maxarraysize;i+=10{
		templistCn:=make([]int,testnumber)
		templistSn:=make([]int,testnumber)
		for j:=0;j<testnumber;j++ {
			list := rand.Perm(i)
			for k, _ := range list {
				list[k]++
			}
			_,templistCn[j],templistSn[j] = DualPivotYank(list,0,len(list)-1)

		}
		listCn = append(listCn,meanfromSlice(templistCn))
		listSn = append(listSn,meanfromSlice(templistSn))
		listVarCNn = append(listVarCNn,VAriancefromSLice(templistCn))
		listVarSn = append(listVarSn,VAriancefromSLice(templistSn))

	}
	return listCn,listSn,listVarCNn,listVarSn
}

func TheorySwapPivot(slice []float64) []float64{
	x:=make([]float64,len(slice))
	for i:=range slice{
		x[i]=0.6*slice[i]*math.Log(slice[i])+0.08*slice[i]+math.Log(slice[i])
	}
	return x
}

func TheoryComparePivot(slice []float64) []float64{
	x:=make([]float64,len(slice))
	for i:=range slice{
		x[i]=1.9*slice[i]*math.Log(slice[i])-2.46*slice[i]
	}
	return x
}

