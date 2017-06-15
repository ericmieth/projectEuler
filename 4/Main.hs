module Main where

isPalindrome :: [Int] -> Bool
isPalindrome x 
  | length x == 0 = True
  | length x == 1 = isPalindrome (init (tail x++x))
  | head x == last x = isPalindrome (init (tail x))
  | otherwise = False

intToList :: Int -> [Int]
intToList x 
  | x <= 9 = [x]
  | otherwise = (intToList (div x 10)) ++ [mod x 10]


main =
  let range = [100..999]
  in print (maximum [ x*y | x <- range, y <- range, isPalindrome (intToList(x*y)) ])
