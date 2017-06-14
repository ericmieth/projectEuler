module Main where

factorize :: [Int] -> Int -> [Int]
factorize divisors x = do
  if x > 1
  then
    if mod x (head divisors) == 0 
    then factorize divisors (div x (head divisors))
    else factorize (succ (head divisors):divisors) x
  else
    divisors

main = print(head (factorize [2] 600851475143))
