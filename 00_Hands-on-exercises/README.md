# List of exercises to practice in GO
## Exercises - Level 1
###### Hands-on exercise #1
1. Using the short declaration operator, ASSIGN these VALUES to VARIABLES with the IDENTIFIERS "x" and "y" and "z" 
a. 42 
b. "James Bond" 
c. true 
2. Now print the values stored in those variables using  
a. a single print statement 
b. multiple print statements 

###### Hands-on exercise #2 
1. Use var to DECLARE three VARIABLES. The variables should have package level scope. Do not assign VALUES to the variables. Use the following IDENTIFIERS for the variables and make sure the variables are of the following TYPE (meaning they can store VALUES of that TYPE). 
a. identifier "x" type int 
b. identifier "y" type string 
c. identifier "z" type bool 
2. in func main 
a. print out the values for each identifier 
b. The compiler assigned values to the variables. What are these values called? 

###### Hands-on exercise #3 
Using the code from the previous exercise, 
1. At the package level scope, assign the following values to the three variables a. for x assign 42 
b. for y assign "James Bond" 
c. for z assign true 
2. in func main 
a. use fmt.Sprintf to print all of the VALUES to one single string. ASSIGN the returned value of TYPE string using the short declaration operator to a VARIABLE with the IDENTIFIER "s" 
b. print out the value stored by variable "s"

###### Hands-on exercise #4 
 FYI - nice documentation and new terminology "underlying type" ? https://golang.org/ref/spec#Types  For this exercise 
1. Create your own type. Have the underlying type be an int. 
2. create a VARIABLE of your new TYPE with the IDENTIFIER "x" using the "VAR" keyword 
3. in func main 
a. print out the value of the variable "x" 
b. print out the type of the variable "x" 
c. assign 42 to the VARIABLE "x" using the "=" OPERATOR 
d. print out the value of the variable "x" 

###### Hands-on exercise #5 
Building on the code from the previous example 
1. at the package level scope, using the "var" keyword, create a VARIABLE with the IDENTIFIER "y". The variable should be of the UNDERLYING TYPE of your custom TYPE "x" 
a. eg:
```
type hotdog int
var x hotdog
var y hotdog
```
2. in func main 
a. this should already be done 
i. print out the value of the variable "x" 
ii. print out the type of the variable "x" 
iii. assign your own VALUE to the VARIABLE "x" using the "=" OPERATOR 
iv. print out the value of the variable "x" 
b. now do this 
i. now use CONVERSION to convert the TYPE of the VALUE stored in "x" to the UNDERLYING TYPE 
1. then use the "=" operator to ASSIGN that value to "y" 
2. print out the value stored in "y" 3. print out the type of "y" 

**code solution:** [Here](https://github.com/Inginex/golangTraining/tree/master/00_Hands-on-exercises/%231_Level1).


## Exercises - Level 2
###### Hands-on exercise #1 
1. Write a program that prints a number in decimal, binary, and hex.

###### Hands-on exercise #2
1. Using the following operators, write expressions and assign their values to variables:
``` 
g. == 
h. <= 
i. >= 
j. != 
k. < 
l. > 
``` 
Now print each of the variables.

###### Hands-on exercise #3
Create TYPED and UNTYPED constants. Print the values of the constants. 

###### Hands-on exercise #4 
Write a program that assigns an int to a variable prints that int in decimal, binary, and hex shifts the bits of that int over 1 position to the left, and assigns that to a variable prints that variable in decimal, binary, and hex. 

###### Hands-on exercise #5 
Create a variable of type string using a raw string literal. Print it.
 
###### Hands-on exercise #6 
Using iota, create 4 constants for the last 4 years. Print the constant values.

**code solution:** [Here](https://github.com/Inginex/golangTraining/tree/master/00_Hands-on-exercises/%232_Level2).


## Exercises - Level 3
###### Hands-on exercise #1 
Print every number from 1 to 10,000.

###### Hands-on exercise #2 
Print every rune code point of the uppercase alphabet three times. Your output should look like this: 
```
65 
U+0041 'A' U+0041 'A' U+0041 'A' 
66 
U+0042 'B' U+0042 'B' U+0042 'B'  
``` 
" through the rest of the alphabet characters.

###### Hands-on exercise #3 
Create a for loop using this syntax for condition { } Have it print out the years you have been alive. 

###### Hands-on exercise #4 
Create a for loop using this syntax  for { } Have it print out the years you have been alive.

###### Hands-on exercise #5 
Print out the remainder (modulus) which is found for each number between 10 and 100 when it is divided by 4.

###### Hands-on exercise #6 
Create a program that shows the "if statement" in action. 

###### Hands-on exercise #7 
Building on the previous hands-on exercise, create a program that uses "else if" and "else". 

###### Hands-on exercise #8 
Create a program that uses a switch statement with no switch expression specified.

###### Hands-on exercise #9 
Create a program that uses a switch statement with the switch expression specified as a variable of TYPE string with the IDENTIFIER "favSport".

###### Hands-on exercise #10 
Write down what these print: 
- fmt.Println(?true && true?) 
- fmt.Println(?true && false?)
- fmt.Println(?true || true?)
- fmt.Println(?true || false?)
- fmt.Println(?!true?) 

**code solution:** [Here](https://github.com/Inginex/golangTraining/tree/master/00_Hands-on-exercises/%233_Level3).


## Exercises - Level 4
###### Hands-on exercise #1 
Using a COMPOSITE LITERAL: create an ARRAY which holds 5 VALUES of TYPE int assign VALUES to each index position. Range over the array and print the values out. Using format printing print out the TYPE of the array.

###### Hands-on exercise #2
Using a COMPOSITE LITERAL: create a SLICE of TYPE int assign 10 VALUES Range over the slice and print the values out. Using format printing print out the TYPE of the slice.

###### Hands-on exercise #3 
Using the code from the previous example, use SLICING to create the following new slices which are then printed: 
```
[42 43 44 45 46]
[47 48 49 50 51]
[44 45 46 47 48] 
[43 44 45 46 47] 
```

###### Hands-on exercise #4 
Follow these steps: 
- start with this slice 
  * x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51} 
- append to that slice this value ? 
  * 52 
- print out the slice 
- in ONE STATEMENT append to that slice these values
  * 53
  * 54
  * 55
- print out the slice 
- append to the slice this slice
  * y := []int{56, 57, 58, 59, 60}
- print out the slice  

###### Hands-on exercise #5 
To DELETE from a slice, we use APPEND along with SLICING. For this hands-on exercise, follow these steps:
- start with this slice 
  * x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51} 
- use APPEND & SLICING to get these values here which you should ASSIGN to a variable "y" and then print: 
  * [42, 43, 44, 48, 49, 50, 51]

###### Hands-on exercise #6 
Create a slice to store the names of all of the states in the United States of America. What is the length of your slice? What is the capacity? Print out all of the values, along with their index position in the slice, without using the range clause. Here is a list of the states: 
```
` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`,  
```

###### Hands-on exercise #7 
Create a slice of a slice of string ([][]string). Store the following data in the multi-dimensional slice: 
``` 
"James", "Bond", "Shaken, not stirred" "Miss", "Moneypenny", "Helloooooo, James." 
```
Range over the records, then range over the data in each record.

###### Hands-on exercise #8 
Create a map with a key of TYPE string which is a person"s "last_first" name, and a value of TYPE []string which stores their favorite things. Store three records in your map. Print out all of the values, along with their index position in the slice. 
```
`bond_james`, `Shaken, not stirred`, `Martinis`, `Women` `moneypenny_miss`, `James Bond`, `Literature`, `Computer Science` `no_dr`, `Being evil`, `Ice cream`, `Sunsets` 
```

###### Hands-on exercise #9 
Using the code from the previous example, add a record to your map. Now print the map out using the "range" loop.

###### Hands-on exercise #10 
Using the code from the previous example, delete a record from your map. Now print the map out using the "range" loop.

**code solution:** [Here](https://github.com/Inginex/golangTraining/tree/master/00_Hands-on-exercises/%234_Level4).


## Exercises - Level 5
###### Hands-on exercise #1
Create your own type "person" which will have an underlying type of "struct" so that it can store the following data: 
- first name 
- last name 
- favorite ice cream flavors 
Create two VALUES of TYPE person. Print out the values, ranging over the elements in the slice which stores the favorite flavors.

###### Hands-on exercise #2 
Take the code from the previous exercise, then store the values of type person in a map with the key of last name. Access each value in the map. Print out the values, ranging over the slice.

###### Hands-on exercise #3 
Create a new type: vehicle.  
- The underlying type is a struct.  
- The fields:  
  * doors 
  * color  
- Create two new types: truck & sedan.
  * The underlying type of each of these new types is a struct.  
  * Embed the "vehicle" type in both truck & sedan.  
  * Give truck the field "fourWheel" which will be set to bool.  
  * Give sedan the field "luxury" which will be set to bool. solution  
- Using the vehicle, truck, and sedan structs:  
  * using a composite literal, create a value of type truck and assign values to the fields; 
  * using a composite literal, create a value of type sedan and assign values to the fields.  
- Print out each of these values. 
- Print out a single field from each of these values. 

###### Hands-on exercise #4 
Create and use an anonymous struct.

**code solution:** [Here](https://github.com/Inginex/golangTraining/tree/master/00_Hands-on-exercises/%235_Level5).

## Exercises - Level 6
###### Hands-on exercise #1
- create a func with the identifier foo that returns an int
- create a func with the identifier bar that returns an int and a string
- call both funcs
- print out their results

###### Hands-on exercise #2
- create a func with the identifier foo that 
* takes in a variadic parameter of type int
* pass in a value of type []int into your func (unfurl the []int)
* returns the sum of all values of type int passed in
- create a func with the identifier bar that 
* takes in a parameter of type []int
* returns the sum of all values of type int passed in


###### Hands-on exercise #3
- Use the "defer" keyword to show that a deferred func runs after the func containing it exits.

###### Hands-on exercise #4
- Create a user defined struct with 
* the identifier “person”
* the fields:
  * first
  * last
  * age
* attach a method to type person with
* the identifier “speak”
* the method should have the person say their name and age
- create a value of type person
* call the method from the value of type person

###### Hands-on exercise #5
- create a type SQUARE
- create a type CIRCLE
- attach a method to each that calculates AREA and returns it
  * circle area= π r 2
  * square area = L * W
- create a type SHAPE that defines an interface as anything that has the AREA method
- create a func INFO which takes type shape and then prints the area
- create a value of type square
- create a value of type circle
- use func info to print the area of square
- use func info to print the area of circle

###### Hands-on exercise #6
- Build and use an anonymous func 

###### Hands-on exercise #7
- Assign a func to a variable, then call that func

###### Hands-on exercise #8
- Create a func which returns a func
- assign the returned func to a variable
- call the returned func

###### Hands-on exercise #9
- A “callback” is when we pass a func into a func as an argument. For this exercise, 
* pass a func into a func as an argument 

###### Hands-on exercise #10
- Closure is when we have “enclosed” the scope of a variable in some code block. For this hands-on exercise, create a func which “encloses” the scope of a variable:

**code solution:** [Here](https://github.com/Inginex/golangTraining/tree/master/00_Hands-on-exercises/%236_Level6).

## Exercises - Level 7
###### Hands-on exercise #1
- Create a value and assign it to a variable. 
- Print the address of that value.

###### Hands-on exercise #2
- create a person struct
- create a func called “changeMe” with a *person as a parameter
  - change the value stored at the *person address
    * important
      * to dereference a struct, use (*value).field 
        * p1.first
        * (*p1).first
      * “As an exception, if the type of x is a named pointer type and (*x).f is a valid selector expression denoting a field (but not a method), x.f is shorthand for (*x).f.”
- create a value of type person
  * print out the value
- in func main
  * call “changeMe”
- in func main
  * print out the value
      
**code solution:** [Here](https://github.com/Inginex/golangTraining/tree/master/00_Hands-on-exercises/%237_Level7).

## Exercises - Level 8
###### Hands-on exercise #1
- Marshal the []user to JSON. There is a little bit of a curve ball in this hands-on exercise - remember to ask yourself what you need to do to EXPORT a value from a package.

###### Hands-on exercise #2
- Unmarshal the JSON into a Go data structure.

###### Hands-on exercise #3
- Encode to JSON the []user sending the results to Stdout. Hint: you will need to use json.NewEncoder(os.Stdout).encode(v interface{})

###### Hands-on exercise #4
- Sort the []int and []string for each person.

###### Hands-on exercise #5
- Sort the []user by 
  * age
  * last
- Also sort each []string “Sayings” for each user
  * print everything in a way that is pleasant

**code solution:** [Here](https://github.com/Inginex/golangTraining/tree/master/00_Hands-on-exercises/%238_Level8).