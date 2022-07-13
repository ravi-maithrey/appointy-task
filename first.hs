module Main where

import System.Environment

main :: IO ()
main = do
  args <- getArgs
  let num1 = read args !! 0 :: Int
  let num2 = read args !! 1 :: Int
  let sum = num1 + num2
  let value = show sum
  putStrLn value
