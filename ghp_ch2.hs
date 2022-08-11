calcChange owed given =
  if change > 0
    then change
    else 0
  where
    change = owed - given

doublePlusTwo x = doubleX + 2
  where
    doubleX = x + x

{- Write a function that takes a value n. If n is even, the function returns
n - 2, and if the number is odd, the function returns 3 × n + 1. To check whether
the number is even, you can use either Haskell’s even function or mod
(Haskell’s modulo function). -}

numCheck n =
  if rem == 0
    then n -2
    else 3 * (n -1)
  where
    rem = mod n 2

{- implementing comparison using where -}
sumSquareOrSquareSum x y =
  if sumSquare > squareSum
    then sumSquare
    else squareSum
  where
    sumSquare = x ^ 2 + y ^ 2
    squareSum = (x + y) ^ 2

{- implementing the above without where and using lambda -}

secondSumSquareOrSquareSum x y =
  ( \sumSquare squareSum ->
      if sumSquare > squareSum
        then sumSquare
        else squareSum
  )
    (x ^ 2 + y ^ 2)
    ((x + y) ^ 2)