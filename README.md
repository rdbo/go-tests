# go-tests
Learning by Go Lang by example

## Thoughts
I started learning GoLang in case libmem gets ported to it. There are some things I'm finding odd about it:
- Resolving dependencies on the source code (e.g `import github.com/someuser/somerepo`). I think there should be a central repository like Rust's crates.io or Python's PyPi, that would avoid package name conflicts and fetching the same package from different URLs
- The compiler throws errors on unused imports and unused variables instead of warnings. Let's say I imported all of the modules I want to use first, and then I started coding. The compiler won't let me run the code because of unused modules. Official FAQ tells you to "pretend" you're using the module, which I didn't like: `var _ = unused.Item  // TODO: Delete before committing!` (what if you forget to delete it later?)
- The function `fmt.Println` inserts spaces automatically for you between each argument. I don't get why, I might be string to just append some stuff that doesn't have spaces.
- It has pointers, and it seems rather common to do operations with them. Almost like in C. I was reading an example of generic types in `gobyexample.com` and there is a sample code for a linked list. It is practically the same thing as the C version. I think pointers should be avoided on languages like this, they only bring trouble. I'm pretty sure Python passes everything (except primitive types) as reference, so if you want to pass an argument to a function that needs to be modified, pass it as a class, because the class will always be a reference. This doesn't seem to be true in Go.
- Enforcing a coding style. In Go, you **have** to do this:
```go
func main() {
  // ...
}
```
Instead of this (this won't even compile):
```go
func main()
{

}
```
I don't think it's that big of a deal, but it's still not nice.
- Syntax is a bit weird. I was working with Goroutines (which I found pretty cool), but then I ended up with this: `out <- <-c` (read the message from channel `c` into the channel `out`). This looks unusual. Also, the functions from the standard library always start uppercase: `fmt.Println`, which I dislike. Some more examples:
```go
func pong(c <-chan string, out chan<- string) {
```
```go
func (v *list[T]) data() []T {
```
```go
menu := map[string]float64 {
        "pizza" : 20.00,
        "burger" : 10.00, "soda" : 2.00, // this comma is MANDATORY, for some reason
}
```
- Returning pointers to local variables is valid. My inner C programmer doesn't like that:
```go
func new_dog() *animal {
    dog := animal{ name: "dog", age: 1 }
    return &dog
}
```
- No function signature overloading + no specific interface implementation. On Rust for example, you have to specify which trait you're implementing to your type. Like this:
```rust
impl fmt::Display for MyType {
  // ...
}
```
In Go, this is not the case. You just implement it as if you were defining normal functions to your struct. This is problematic, what if two interfaces that you need to implement have the same function name and different signatures? Even better, what if they have the same function name and same signature, but different code? In Rust, because you're implementing the traits one by one, it can differenciate. This is not the case in Go.

- Strange keywords. First example I would say is the keyword `range`. It can do many different things, like iterating through a map, arrays and slices, and variadic arguments. What it can't do is iterate through a range of numbers. So why is it called range? Second example is the keyword `rune`. They purposefully avoid using the word `char` (for character) because a rune is supposed to represent Unicode code points instead of characters - I think this is overthinking about semantics.

These things aside, there are some good things I like about it:
- It has a very powerful standard library, which I think is the major selling point for the language. Unlike Rust, it has a very robust standard library with so many different modules for everyday use. I think that is a positive thing for GoLang.
- The goroutines are a pretty interesting implementation of asynchronous communication. It is built in the language, and everything is very simple. Another plus for it.
- Fast runtimes. The fact that GoLang is a compiled language directly to machine code makes the execution much faster than, for example, Python or Java.

### TLDR;
I think the language has a great potential, especially because of the standard library and the fast execution times. But there also are a lot of problems that cannot be ignored.
