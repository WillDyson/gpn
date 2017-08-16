# GPN
A golang RPN calculator

## Usage
```
> gpn
56
78
*
12
34
*
-
.
(3960.000000)
```


## Supported Operations
```
> gpn
.com
     +  (y+x) adds x and y on the stack
     -  (y-x) takes x from y on the stack
     *  (y*x) multiplies x with y on the stack
     /  (y/x) divides y by x on the stack
     ^  (y**x) raises y to the power of x on the stack
     $  (e**x) raises e to the power of x on the stack
   log  (log x) logs x on the stack
   sin  (sin x) sins x on the stack
   cos  (cos x) coses x on the stack
   tan  (tan x) tans x on the stack
  asin  (asin x) asins x on the stack
  acos  (acos x) acoses x on the stack
  atan  (atan x) atans x on the stack
   sum  sums all values on the stack
  mean  takes the mean of all the values on the stack
 clear  clears the stack
 count  counts the number of values on the stack
   pop  pops the element at the top of the stack
     .  prints the top value of stack
    .q  quits the program
  .com  lists all availiable commands
```
