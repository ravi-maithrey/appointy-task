{--
Haskell has a function called repeat that takes a value and repeats it infinitely.
Using the functions you’ve learned so far, implement your own version of repeat.
--}
myRepeat n = cycle [n]

{--Write a function subseq that takes three arguments: a start position, an end position,
and a list. The function should return the subsequence between the start and end.--}
subseq start end list = take difference (drop start list)
    where difference = end - start

{--
Write a function inFirstHalf that returns True if an element is in the first half of a
list, and otherwise returns False.
--}
-- since elem returns a bool anyway, we just return it and we drop the revrse of the list as we want the first half
inFirstHalf element list = element `elem` (drop half (reverse list))
    where half = length list `div` 2 -- div instead of / as / gives a frac int
