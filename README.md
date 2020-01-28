Example to practice GO.

# random_number_hashing

This program execute a for loop (init/condition/post) which generate 10 Random Float64 numbers. 
It can generate int, uint, float64 & more, just change the type of number you want to iterate over the loop.

In func main() numbers are generated, then converted to strings with fmt.Sprintln just for conversion fun.
At each iteration of the for loop the number generated gets hashed crypto/sha256 package.

