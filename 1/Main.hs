module Main where
import Data.List

solve x = sum ( union [y | y <- x, mod y 3== 0] [z | z <- x, mod z 5 == 0] )

main = print (solve [1..999])
