runningSum :: [Int] -> [Int]
runningSum (x1:x2:xs) = x1:(x1+x2):runningSum (x2:xs)
runningSum [x] = [0+x]
runningSum [] = [0]