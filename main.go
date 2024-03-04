package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
	"unicode/utf8"
)

var wg = sync.WaitGroup{}
var results = []string{}
var mut = sync.RWMutex{}

func main() {

	fmt.Println("Something")

	var integer0 int16 = 389

	fmt.Println("Signed integer: ", integer0)

	var integer1 uint16 = 38945

	fmt.Println("Un-signed integer: ", integer1)

	var float0 float32 = 389.89065

	fmt.Println("Float: ", float0)

	// cast

	fmt.Println("Cast float to int. Result: ", int(float0))

	// division
	fmt.Println("int/int e.g 3/2 results in rounded-down int. Result: ", 3/2)

	var my_string1 = "Hello World!"
	var my_string2 string = "Hello \n World!"

	fmt.Println("String 1: ", my_string1)
	fmt.Println("String 2: ", my_string2)

	fmt.Println(`
	This is
	a 
	string block or whareva`)

	fmt.Println("String len() and RuneCount are different as len is len in bytes: ")
	fmt.Println(`len("θ") is :`, len("θ"))

	fmt.Println(`utf8.RuneCountInString("θ") is :`, utf8.RuneCountInString("θ"))

	var myRune rune = 'u'

	fmt.Println("Rune: ", myRune)

	runeShort := 'e'

	fmt.Println(runeShort)

	var myBoolean bool = true

	fmt.Println("Boolean: ", myBoolean)

	var someType = "jrklg"

	fmt.Println(someType)

	var1, var2, var3 := "jhjj", "jkd", 3

	fmt.Println(var1, var2, var3)

	const myConst string = "Something!!"

	print(myConst)

	print(fmt.Sprint(intDivision(2, 3)))

	result, remainder, error := intDivision(3, 8)

	if error != nil {
		print("Error encountered!")
	} else if remainder == 0 {
		print("Perfect division")
	}

	fmt.Printf("Dividing %v by %v, result is %v and remainder is %v \n", 8, 3, result, remainder)

	// arrays
	var intArr [3]int

	intArr[0] = 3
	intArr[2] = 5356
	intArr[1] = 7777

	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3])

	fmt.Println(&intArr[0])

	var strArr [3]string = [3]string{"akd", "hjik", "jdjmkf"}

	fmt.Println(strArr)

	strArr2 := [...]string{"akd", "hjik", "jdjmkf"}

	fmt.Println(strArr2)

	// slices
	mySlice := []string{"f", "rt", "f"}

	fmt.Printf("The length of mySlice before update is %v and capacity is %v \n", len(mySlice), cap(mySlice))

	mySlice = append(mySlice, "wdv")

	fmt.Printf("The length of mySlice before update is %v and capacity is %v \n", len(mySlice), cap(mySlice))

	mySlice = append(mySlice, mySlice...)

	fmt.Println(mySlice)

	// maps
	var myMap map[string]string = make(map[string]string)
	fmt.Println(myMap)

	var map2 = map[string]string{"A": "B", "C": "D"}
	fmt.Println(map2)
	fmt.Println(map2["A"])
	fmt.Println(map2["YS"])

	var value, containsKey = map2["weck"]

	fmt.Printf("Value: %v, containsKey? %v \n", value, containsKey)

	delete(map2, "C")

	fmt.Println(map2)

	/// loops
	// for
	for item := range intArr {
		fmt.Println(item)
	}

	// while
	// impl 1
	var i int = 0
	for i < 10 {
		fmt.Println("Loop in i= ", i)
		i++
	}

	// impl 2
	var g int = 0
	for {
		if g >= 12 {
			break
		}
		fmt.Println("Loop in g= ", g)
		g++
	}

	// impl 3
	for i := 0; i <= 12; i++ {
		fmt.Println("Loop in i= ", i)
	}

	// structs
	var engine gasEngine = gasEngine{"hbjd", 200, owner{"hjk"}}
	fmt.Println(engine)

	engine.capacity = 300
	engine.owner.name = "Ian N."
	fmt.Println(engine.capacity)
	fmt.Println(engine)
	fmt.Println(engine.name)

	// anonymous structs
	var engine2 = struct {
		mpg  int16
		name string
	}{2, "RTX 456"}

	fmt.Println(engine2)

	print(engine.nameAndCC())

	// goroutines
	t0 := time.Now()
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go someFunction(i)
	}
	wg.Wait()
	fmt.Println("Excecution time: ", time.Since(t0))
	fmt.Println("Results: ", results)

	// generic function
	var sliceA = []int64{8, 4, 6}
	var sumA = sumSlice[int64](sliceA)
	fmt.Println(sumA)

	fmt.Println(isEmpty(sliceA))

	// generic struct
	var petrolCar = car[petrolEngine]{
		ownerName: "Ian N",
		age:       12,
		engine: petrolEngine{
			capacity: 1235,
			fuel:     "petrol",
		},
	}

	var elecCar = car[electricEngine]{
		ownerName: "Ian N",
		age:       12,
		engine: electricEngine{
			rangeInKms: 1235,
			kWh:        156,
		},
	}

	fmt.Println("Petrol Car: ", petrolCar)
	fmt.Println("Electric Car: ", elecCar)

}

func print(print string) {
	fmt.Println(print)
}

func intDivision(denom int, numer int) (int, int, error) {

	if denom == 0 {
		var error error = errors.New("division by zero")
		return 0, 0, error
	}
	return numer / denom, numer % denom, nil
}

type gasEngine struct {
	variant  string
	capacity int16
	owner
}

func (g gasEngine) nameAndCC() string {
	return g.variant + " - " + fmt.Sprint(g.capacity) + " cc"
}

type owner struct {
	name string
}

func someFunction(i int) {
	var t int = i * 1000
	time.Sleep(time.Duration(t))
	fmt.Println("Finished Routine A...when i= ", i)
	mut.RLock()
	results = append(results, fmt.Sprint(i))
	mut.RUnlock()
	wg.Done()
}

// generics
// generic func
func sumSlice[T int | float32 | int64 | int32 | int16 | float64](slice []T) T {
	var sum T
	for _, i := range slice {
		sum += i
	}
	return sum
}

func isEmpty[T any](slice []T) bool {
	return len(slice) == 0
}

// generic struct
type petrolEngine struct {
	capacity int
	fuel     string
}

type electricEngine struct {
	rangeInKms int
	kWh        int
}

type car[T petrolEngine | electricEngine] struct {
	age       int
	ownerName string
	engine    T
}
