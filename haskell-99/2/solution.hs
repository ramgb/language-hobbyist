myButLastHelper:: [a] -> a ->  a
myButLastHelper (_:[]) x = x
myButLastHelper (y:ys) x = myButLastHelper ys y

myButLast ::  [a] -> a
myButLast [] = error "empty list"
myButLast (_:[]) = error "only one element list"
myButLast (x:xs) = myButLastHelper xs x

main :: IO()
main = return()
