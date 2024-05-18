package main

import "fmt"


type NumberStorer interface {
	GetAll() ([]int, error)
	Put(int) error
}

type ApiServer struct {
   numberStore NumberStorer 
}


type MongoDBNumberStore struct {
	// / some values
}

func (m MongoDBNumberStore) GetAll() ([]int,error){
	return []int{1,2,3}, nil
}

func (m MongoDBNumberStore) Put(number int) (error){
	fmt.Println("store the number into the mongodb storage")
	return nil;
}



func main() {

	apiServer := ApiServer{
		numberStore: MongoDBNumberStore{},
	}

	numbers, err := apiServer.numberStore.GetAll();
	if err != nil {
		panic(err)
	}

	// no error!
	fmt.Println(numbers)


	// numbers := []string{"a", "b"}
	// for loops!
	// for i := 0; i < 10; i++ {
	// 	fmt.Println("iter: %v", i)
	// }

	// range!
	// for i := range numbers {
	// 	fmt.Println(i)
	// }

}
