module Main where

import System.IO
import System.Random

main = do
    hSetBuffering stdin LineBuffering
    num <- randomRIO(1 :: Int, 100)
    putStrLn "I'm thinking of a number between 1 and 100"
    doGuessing num

doGuessing num = do
    putStrLn "Enter your guess:"
    guess <- getLine
    let guessNum = read guess
    if guessNum < num
        then putStrLn "Too Low"
            doGuessing num
        else if guessNum > num
            then putStrLn "Too High"
                doGuessing num
            else putStrLn "You Got It!"