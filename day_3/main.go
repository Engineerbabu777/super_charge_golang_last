


package main

type Storage interface {
	Get(id int) (any,error)
	Put(id int, val any) error
}

type Server struct{
     store Storage
}

type FooStorage struct{

}

func (s *FooStorage) Get(id int)(any, error){
	return nil, nil
}

func (s *FooStorage) Put(id int, val any) error{
    return nil
}

func updateValue(id int, val any, store Storage) error {
	return store.Put(id, val)
}

func main(){
	s := &Server{
		store: &FooStorage{},
	};
	s.store.Get(1)
	updateValue(1,"Foo", s.store)
}


