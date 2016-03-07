import Data.List
import System.IO
maxInt = maxBound::Int
minInt = minBound::Int
bigFloat = 3.99999999999999999

doubleMe x = x + x  
doubleUs x y = doubleMe x + doubleMe y

doubleSmallNumber x = if x > 100  
                        then x  
                        else x*2   

boomBangs xs = [ if x < 10 then "BOOM!" else "BANG!" | x <- xs, odd x]   
removeNonUppercase st = [ c | c <- st, c `elem` ['A'..'Z']]   

evensub xxs= [ [ x | x <- xs, even x ] | xs <- xxs]  
