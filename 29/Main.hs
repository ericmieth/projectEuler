module Main where
import Data.List

solve :: [Integer] -> Int
solve x =
  length (nub (sort ([ z^y | z <- x, y <- x])))

main = print (solve [2..100])
