import Data.List
import System.IO

fib 0 = 0
fib 1 = 1
fib n = fib (n-1) + fib (n-2)

nats = 1 : map (+1) nats

removeNonUppercase :: [Char] -> [Char]  
removeNonUppercase st = [ c | c <- st, c `elem` ['A'..'Z']]   

circumference :: Float -> Float  
circumference r = 2 * pi * r  


circumference' :: Double -> Double  
circumference' r = 2 * pi * r  


charName :: Char -> String  
charName 'a' = "Albert"  
charName 'b' = "Broseph"  
charName 'c' = "Cecil"  

factorial :: (Integral a) => a -> a  
factorial 0 = 1  
factorial n = n * factorial (n - 1)  

addVectors :: (Num a) => (a, a) -> (a, a) -> (a, a)  
addVectors (x1, y1) (x2, y2) = (x1 + x2, y1 + y2)  