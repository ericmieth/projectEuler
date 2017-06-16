module Main where

{-

2^15 = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.

What is the sum of the digits of the number 2^1000?

-}

intToList :: Integer -> [Integer]
intToList x 
  | x <= 9 = [x] 
  | otherwise = (intToList (div x 10)) ++ [mod x 10] 

main = print ( sum (intToList (2^1000)))
