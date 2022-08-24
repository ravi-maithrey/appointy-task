import           Data.Char (toLower)
-- writing my version of take
myTake _ []     = []
myTake 0 _      = []
myTake n (x:xs) = x:(myTake (n-1) xs)
-- myTake 1 xs:x = xs:[]
greatestCommonDivisor a b = if a `rem` b == 0
    then b
    else greatestCommonDivisor b (a `rem` b)

-- writing my version of reverse
myreverse []     = []
myreverse (x:xs) = myreverse xs ++ [x]
-- writing the elem function
-- myElem a (x:xs) = filter (a == x) (x:xs)

{-- Your isPalindrome function from lesson 6 doesn’t handle sentences with spaces or
capitals. Use map and filter to make sure the phrase “A man a plan a canal Panama”
is recognized as a palindrome. --}
-- basic is a palindrome definition
isPalindrome x = x == reverse x
-- adding to the above definition to handle spaces
modifiedIsPalindrome xs = isPalindrome (filter (/=' ') (map toLower xs))
