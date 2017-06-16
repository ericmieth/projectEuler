module Main where

smallest :: [Int] -> [Int] -> Int
smallest x d
  | length d == 1 = head x
  | mod (head x) (last d) == 0 = smallest x (init d)
  | otherwise = smallest (tail x) [1..20]

main = print ( smallest [20,40..] [1..20] )
