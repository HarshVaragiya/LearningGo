# Exercise: Slices
Implement Pic. It should return a slice of length `dy`, each element of which is a slice of `dx` 8-bit unsigned integers. When you run the program, it will display your picture, interpreting the integers as grayscale (well, bluescale) values.

The choice of image is up to you. Interesting functions include `(x+y)/2`, `x*y`, and` x^y`.

## Hints
1. You need to use a loop to allocate each []uint8 inside the [][]uint8.
2. Use uint8(intValue) to convert between types.

## Output

- Averaging the x and y positions gives us the image 
![Output_Avg](Output/Avg.jpg)


- Gradients in x and y directions generate the following image
![Output_Gradient1](Output/Gradient1.jpg)
![Output_Gradient2](Output/Gradient2.jpg)


- Raising x to power of y gives us the following image 
![Output_Exp](Output/Exp.jpg)


- Multiplying x and y to generate the image gives us a nice pattern 
![Output_Mul](Output/Mul.jpg)


- Changing the x and y values using left and right shifts and multiplying gives us
![Output_Mul2](Output/Mul2.jpg)