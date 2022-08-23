{--
Haskell has a function called repeat that takes a value and repeats it infinitely.
Using the functions you’ve learned so far, implement your own version of repeat.
--}
myRepeat n = cycle [n]

{--Write a function subseq that takes three arguments: a start position, an end position,
and a list. The function should return the subsequence between the start and end.--}
subseq start end list = take difference (drop start list)
    where difference = end - start
