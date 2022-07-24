module Main where

-- import IO

-- import IO
import Data.ByteString ()
import System.IO (BufferMode (LineBuffering), hSetBuffering, stdin)

main = do
  hSetBuffering stdin LineBuffering
  putStrLn "Please enter your name: "
  name <- getLine
  putStrLn ("Hello, " ++ name)