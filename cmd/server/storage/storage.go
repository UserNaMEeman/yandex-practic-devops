package storage

// import(
// 	"fmt"
// )


type DataStore struct{
	Name string
	Type string
	ValueF float64
	ValueC int64
}

// func (data *DataStore)SaveData(){
// 	var storage map[string]float64			//Change string - add storage method
// 	storage = make(map[string]float64)
// 	storage[data.Name] = data.Value
// }

// func (data *Counter)SaveDataD(){
// 	var randV int64 = 7				//Change string - add storage method
// 	data.Value = data.Value + randV
// 	// fmt.Println(data.PollCount)
// }