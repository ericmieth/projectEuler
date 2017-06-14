module Main where 

next :: [Int] -> [Int]
next x = do
  let next' = sum (take 2 (reverse x))
  if next' < 4000000
  then next (x++[next'])
  else x

main = print (sum [ x | x <- next [1,2], mod x 2 == 0 ] )
