package main

func main(){
	Comp,Swaps,_,_:=TestDualPivotYank(100,10000)
	x:= makeX(Swaps,10)
	teorySwap:=TheorySwapPivot(x)
	theoryComp:=TheoryComparePivot(x)
	plot("Size","Swaps","Swap yank ",2,
		"Empiric","Theory",x,Swaps,teorySwap,"YankSwap.png")
	plot("Size","Comprasions","Comp yank",2,
		"Empiric","Theory",x,Comp,theoryComp,"YankComp.png")

}
