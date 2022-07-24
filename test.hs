module Main
    where

x = 5
y = (6, "Hello")
z = x * fst y

main = putStrLn "Hello World"

square x = x*x

signum x = 
    if x < 0 
        then -1
    else if x > 0
        then 1
        else 0

f x =
    case x of
        1 -> 2
        2 -> 5
        -1 -> 3
        _ -> 0

-- using let and in
roots a b c =
    let det = sqrt (b*b - 4*a*c)
        twice_a = 2*a
    in ((-b + det)/ twice_a,
        (-b - det)/ twice_a)

factorial 1 = 1
factorial n = n * factorial (n-1)

-- demonstrating recursion 
my_filter predicate [] = []
my_filter predicate (x:xs) = -- x first element and xs rest of list (since in haskell lists are a:b:c etc)
    if predicate x -- if predicate is true on x
        then x : my_filter predicate xs -- add x to list beind returned and call my_filter on rest of list
        else my_filter predicate xs -- if false dont add x to list and call my_filter on rest of list

-- fibonacci sum
fibonacci 1 = 1
fibonacci 2 = 1
fibonacci n = fibonacci (n-2) + fibonacci (n-1)

-- multiplication using addition
mult a 1 = a
mult a b = a + mult a (b-1)

-- writing a my_map using recursion
my_map func [] = [] -- base case
my_map func (x:xs) = 
    (func x) : my_map func xs