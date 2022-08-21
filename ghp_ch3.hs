-- if even increment
ifEvenInc n = if even n
    then n + 1
    else n
-- if even double
ifEvenDouble n = if even n
    then n*2
    else n
-- if even square
ifEvenSquare n = if even n
    then n^2
    else n
-- we can combine these into a single function to minimize lines of code
ifEven myFunction n = if even n
    then myFunction n
    else n

inc n = n + 1
double n = n * 2
square n = n ^ 2


