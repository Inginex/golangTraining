package main

import (
	"os"
	"fmt"
	"strings"
	"bufio"
)

func main(){
	const input = "Oh! ye whose dead lie buried beneath the green grass; who standing among flowers can say--here, HERE lies my beloved; ye know not the desolation that broods in bosoms like these.  What bitter blanks in those black-bordered marbles which cover no ashes!  What despair in thos immovable inscriptions!  What deadly voids and unbidden infidelities in the lines that seem to gnaw upon all Faith, and refuse resurrections to the beings who have placelessly perished without a grave.  As well might those tablets stand in the cave of Elephanta as here. "

	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
}