package main 

import "fmt"

type vehicle struct {
	doors int
	color string
}
type truck struct {
	vehicle
	fourWheel bool
}
type sedan struct {
	vehicle
	luxury bool
}

func main(){
	car1 := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "Blue",
		},
		luxury: false,
	}
	car2 := truck{
		vehicle : vehicle{
			doors: 2,
			color: "Red",
		},
		fourWheel: true,
	}

	fmt.Printf("Info Car1: Doors = %v, Color = %v, Luxury = %v.\n", car1.doors, car1.color, car1.luxury)
	fmt.Printf("Info Car2: Doors = %v, Color = %v, FourWheels = %v.\n", car2.doors, car2.color, car2.fourWheel)
}