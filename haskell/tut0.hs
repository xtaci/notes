import Data.List
import System.IO

fib 0 = 0
fib 1 = 1
fib n = fib (n-1) + fib (n-2)

nats = 1 : map (+1) nats

removeNonUppercase :: [Char] -> [Char]  
removeNonUppercase st = [ c | c <- st, c `elem` ['A'..'Z']]   

factorial :: Integer -> Integer  
factorial n = product [1..n]  

circumference :: Float -> Float  
circumference r = 2 * pi * r  


circumference' :: Double -> Double  
circumference' r = 2 * pi * r  

