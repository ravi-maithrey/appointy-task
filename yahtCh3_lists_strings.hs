import qualified Data.Char
{- Use map to convert a string into a list of booleans, each element in the 
new list representing whether or not the original element was a lower-case 
character. That is, it should take the string "aBCde"
and return [True,False,False,True,True]. -}

map Data.Char.isLower "aBCde"

{- Use the functions mentioned in this section (you will need two of them) to 
compute the number of lower-case letters in a string. For
instance, on "aBCde" it should return 3. -}

length (map Data.Char.isLower "abCDe")

{- We've seen how to calculate sums and products using folding functions. 
Given that the function max returns the maximum of two numbers, write a 
function using a fold that will return the maximum value in a list (and 
zero if the list is empty). So, when applied to [5,10,2,8,1] it will return 
10. Assume that the values in the list are always ≥ 0 Explain to yourself why 
it works. -}

foldr max 0 [5,10,2,8,1]

{- Write a function that takes a list of pairs of length at least 2 and 
returns the first component of the second element in the list. So, 
when provided with [(5,'b'),(1,'c'),(6,'a')], it will return
1. -}

fst (head (tail [(5,'b'),(1,'c'),(6,'a')]))