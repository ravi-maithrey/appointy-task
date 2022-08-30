import           Test.Tasty
import           Test.Tasty.HUnit      (Assertion, testCase, (@?=))
import qualified Test.Tasty.QuickCheck as QC

runningSum :: [Int] -> [Int]
runningSum (x1:x2:xs) = x1:(x1+x2):runningSum (x2:xs)
runningSum [x]        = [0+x]
runningSum []         = [0]

group3 [] = []
group3 xs = g : group3 ys where (g,ys) = splitAt 3 xs

-- the quick check property can be
-- length (group3 xs) == ((length xs `div` 3) + (length xs `mod` 3))

-- another property
-- concat all lists after group3 == list before group 3
-- xs == concat (group3 xs)
