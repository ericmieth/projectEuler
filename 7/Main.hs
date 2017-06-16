module Main where

{-

By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13,
we can see that the 6th prime is 13.

What is the 10 001st prime number?

-}

isPrime :: Int -> Bool
isPrime x = 
  if null [ y | y  <- [2..x-1], mod x y == 0]
  then True
  else False
  
  

main = print ( last (take 10001 [z | z <- [2..], isPrime z]) )
