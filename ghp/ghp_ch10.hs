-- we can store values by calling closures
-- example: coffeecup

-- read the hlint suggesstions to look at other ways
-- here we are passing message as a function to our lambda
cup :: t1 -> (t1 -> t2) -> t2
cup size = \message -> message size

-- it is not enough to store the valuue, we want to make use of that as well
smallCup = cup 6
-- storing a cup of size 6 in smallCup

-- creating getters
getSize cup = cup (\size -> size)
-- try getSize smallCup
