package main

//we can import packages inside a module with/without alias
//like here we have aliased package "golang_tut" as aliasForSomePackage
import aliasForSomePackage "golang_tut/externalFile"

func main() {
	for i := 0; i < 10; i++ {
		aliasForSomePackage.Hey()
	}
}
