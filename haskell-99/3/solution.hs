elementAt :: [a] -> Int -> a
elementAt (x:_) 1 = x
elementAt [] _ = error "array length out of bounds"
elementAt (x:xs) k 
  | k > 0  = elementAt xs (k-1) 
  | otherwise = error "index must be positive"

main :: IO()
main = return()
