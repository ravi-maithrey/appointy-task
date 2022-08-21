-- a function that can build any of the ifEven type functions
-- we definined in the prev example
-- the below function generates more functions
genIfEven f = (\x -> ifEven f x)
-- takes any f and applies it to ifEven
-- we have captured the value of x inside the lambda
-- our new func will be waiting for an x


-- Quick check 5.1 Write a function genIfXEven 
-- that creates a closure with x and returns a new
-- function that allows the user to pass in a function 
-- to apply to x if x is even.
genIfXEven x = (\f -> ifEven f x)

