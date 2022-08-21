import Data.List -- for the sorting done later
-- if even increment
{--
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
--}
-- we can combine these into a single function to minimize lines of code
ifEven myFunction n = if even n
    then myFunction n
    else n
-- we use the below functions in place of the myFunction above
inc n = n + 1
double n = n * 2
square n = n ^ 2
-- we can also combine these directly in code
ifEvenInc = ifEven inc
ifEvenSquare = ifEven square
ifEvenDouble = ifEven double

-- a function for cubing x and passing to ifEven
ifEvenCube = ifEven (\x -> x^3)


-- implementing a custom sorting by last names then first names in a tuple
names = [("Ian", "Curtis"),
         ("Bernard","Sumner"),
         ("James", "Hook"),
         ("Peter", "Hook"),
         ("Stephen","Morris")]

compareLastNames name1 name2 = if lastName1 > lastName2
                               then GT
                               else if lastName1 < lastName2
                                    then LT
                                    else if firstName1 > firstName2
                                        then GT
                                        else if firstName1 < firstName2
                                            then LT
                                            else EQ
  where lastName1 = snd name1
        lastName2 = snd name2
        firstName1 = fst name1
        firstName2 = fst name2

-- we check this with 'sortBy compareLastNames names'