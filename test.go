package main

import (
    "fmt"
    "math"
    "errors"
    "time"
)

func mypow(base float64, n int) float64 {
    newbase := 1.0
    for i := 0; i < n; i++ {
        newbase *= base
    }

    return newbase
}

func prime_pair() (int, int) {
    return 2, 3
}

func sequence() func() int {
    i := 1
    return func() int {
        i *= 2
        return i
    }
}

type animal struct {
    name string
    age int
}

func new_dog() *animal {
    dog := animal{ name: "dog", age: 1 }
    return &dog
}

type rectangle struct {
    width int
    height int
}

// implement functions for struct 'rectangle'
func (r rectangle) area() int {
    return r.width * r.height
}

// implement function for pointer to struct 'rectangle'
func (r *rectangle) get_address() *rectangle {
    return r
}

// variadic function
func sum(nums ...int) int {
    sum := 0
    for _, num := range nums {
        sum += num
    }

    return sum
}

// interfaces
type vehicle interface {
    model() string
    maxspeed() float64
}

// derive from interface
type car struct {  }

func (c car) model() string {
    return "Bugatti"
}

func (c car) maxspeed() float64 {
    return 500.0
}

func show_vehicle(v vehicle) {
    fmt.Println("Model: ", v.model())
    fmt.Println("Max Speed: ", v.maxspeed())
}

// inherit fields from other structs (struct embedding)
type electronic struct {
    name string
    price float64
}

type computer struct {
    electronic // inherit all fields from struct 'electronic'
    rating int
}

func (e electronic) show() {
    fmt.Println("Name: ", e.name)
    fmt.Println("Price: ", e.price)
}

// generics (linked list)
type item[T any] struct {
    val T
    next *item[T]
}

type list[T any] struct {
    head, tail *item[T]
}

func (v *list[T]) push(val T) { // more pointers and null pointers...
    if v.tail == nil {
        v.head = &item[T] { val, nil }
        v.tail = v.head
    } else {
        v.tail.next = &item[T] { val, nil }
        v.tail = v.tail.next
    }
}

func (v *list[T]) data() []T {
    var items []T
    for it := v.head; it != nil; it = it.next {
        items = append(items, it.val)
    }
    return items
}

// errors
type myerror struct {  }

func (e myerror) Error() string {
    return "my custom error"
}

func bad_function() (int, error) {
    return -1, errors.New("bruh........ what is even this")
}

func another_bad_function() (int, error) {
    return -1, myerror{  }
}

// goroutines
func do_something() {
    fmt.Println("Gorountines is a weird name, but it's fine")
}

// channel direction
func ping(c chan<- string) { // c is a channel that can only send data
    c <- "ping"
}

func pong(c <-chan string, out chan<- string) { // c is a channel that can only receive data
    out <- <-c // send the ping to the out channel (this syntax is horrendous)
    out <- "pong" // send the pong to the out channel
}

func main() {
    // variables and printing
    fmt.Println("bruh")
    var result int = 2 * 10
    fmt.Println("2 * 10 = ", result)

    some_text := "This variable was declared and initialized"
    fmt.Println(some_text)

    // math
    const num int64 = 2e10
    fmt.Println(num)

    fmt.Printf("Sine of PI: %.2f", math.Sin(math.Pi))
    fmt.Println()

    for i := 0; i < 10; i++ {
        fmt.Print(i, " ")
    }
    fmt.Println()

    // switch and inline func
    print_type := func(i interface{}) {
        switch t := i.(type) {
        case int:
            fmt.Println("The value is an integer")
        case bool:
            fmt.Println("The value is a boolean")
        default:
            fmt.Printf("The value is unknown: %T", t)
            fmt.Println()
        }
    }

    print_type(10)
    print_type(true)
    print_type('A')
    print_type("C is better")

    // arrays
    var arr [4]int
    arr[0] = 10
    fmt.Println("Array: ", arr)

    matrix := [2][2]int {{ 1, 0 }, { 0, 1 }}
    fmt.Println("Matrix: ", matrix)

    // slices
    s := make([]string, 2)
    s[0] = "Hello"
    s[1] = "World"
    fmt.Println("Slice: ", s)
    fmt.Println("Slice Length: ", len(s))

    c := make([]string, len(s))
    copy(c, s)
    c = append(c, "Testing")
    c = append(c, "GoLang")

    fmt.Println("New Slice: ", c)
    fmt.Println("New Slice Length: ", len(c))

    numbers := []int {0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
    slice := numbers[1:4]
    fmt.Println("Number Slice: ", slice)

    // maps
    menu := map[string]float64 {
        "pizza" : 20.00,
        "burger" : 10.00, "soda" : 2.00, // this comma is MANDATORY, for some reason
    }

    fmt.Println("Menu (Map): ", menu)
    delete(menu, "soda")
    fmt.Println("New Menu: ", menu)

    fruits := []string{ "apple", "banana", "orange", "lemon" }
   
    fmt.Print("Fruits: [")
    for i := range fruits {
        fmt.Print(fruits[i], " ")
    }
    fmt.Println("]")

    stocks := map[string]string{ "GOOGL" : "$100.00", "TSLA" : "$500.00", "MSFT" : "$0.00" }
    for key, value := range stocks {
        fmt.Println(key, " -> ", value)
    }

    // functions
    base := 2.0
    exponent := 10
    newbase := mypow(base, exponent)
    fmt.Println(base, "^", exponent, "=", newbase) // it automatically puts spaces between the arguments for some reason

    prime0, prime1 := prime_pair()
    fmt.Println("Prime Pair: ", prime0, prime1)

    // closures
    next := sequence()
    fmt.Println("Sequence: ", next(), next(), next(), next(), next(), next())

    // pointers
    value := 10
    var pvalue *int = &value
    fmt.Println("Value: ", value)
    fmt.Println("Address of Value: ", pvalue)
    *pvalue = 1337
    fmt.Println("New Value: ", value)

    // strings and "runes" (fancy word for int32 character or unicode character)
    var a rune = 'a'
    fmt.Println("Rune: ", a)
    fmt.Printf("Rune Unicode: %U#", a)
    fmt.Println()
    fmt.Println("Formatted Rune: ", string(a))

    // structs
    dog := new_dog()
    dog.age = 10
    fmt.Println("Animal: ", dog)

    cat := animal { "cat", 2 }
    fmt.Println("Animal: ", cat)

    // methods for structs
    rect := rectangle{5, 2}
    fmt.Println("Rectangle Area: ", rect.area())

    prect := &rect
    fmt.Println("Rectangle Address: ", prect.get_address())

    // variadic function
    result = sum(1, 2, 3, 4, 5)
    fmt.Println("Variadic sum: ", result)

    nums := []int {10, 20, 30, 40}
    result = sum(nums...)
    fmt.Println("Another variadic sum: ", result)

    // interfaces
    mycar := car{  }
    show_vehicle(mycar)

    // struct embedding
    pc := computer { electronic: electronic { name: "Gaming PC", price: 1000.0 }, rating: 5 } 
    pc.show()

    // generics
    mylist := list[string]{  }
    mylist.push("Testing")
    mylist.push("Generics")
    fmt.Println(mylist.data())

    // errors
    fmt.Println(bad_function())
    fmt.Println(another_bad_function())

    // "goroutines" (a.k.a coroutines)
    go do_something()

    go func() {
        fmt.Println("Goroutine finished")
    }()

    time.Sleep(time.Second) // wait for coroutines to run

    // channels
    messages := make(chan string)
    go func () {
        fmt.Println("ping")
        messages <- "ping" // send message to channel
    }()

    _ = <-messages // wait and receive message from channel
    fmt.Println("pong")

    // buffered channels
    messages = make(chan string, 2)
    messages <- "Buffered"
    messages <- "Channels"
    fmt.Println(<-messages)
    fmt.Println(<-messages)

    // channel direction
    cpong := make(chan string, 1)
    cout := make(chan string, 1)

    go ping(cpong)
    go pong(cpong, cout)
    fmt.Println(<-cout)
    fmt.Println(<-cout)

    // select (receive messages from different channels once they arrive)
    c1 := make(chan string)
    c2 := make(chan string)

    go func() {
        time.Sleep(1 * time.Second)
        c1 <- "one"
    }()

    go func() {
        time.Sleep(2 * time.Second)
        c2 <- "two"
    }()

    for i := 0; i < 2; i++ {
        select {
        case msg1 := <-c1:
            fmt.Println("received", msg1)
        case msg2 := <-c2:
            fmt.Println("received", msg2) // gets received 1s after the first one, even though it delays for 2 seconds and 'msg1 := <-c1' would normally wait for the channel to send a message, therefore making it 3 seconds. that's because of the select
        }
    }
}
