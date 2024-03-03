myLengthHelper :: [a] -> Int -> Int
myLengthHelper [] ctr = ctr
myLengthHelper (_:xs) ctr = myLengthHelper xs (ctr+1)

myLength :: [a] -> Int
myLength [] = 0
myLength (x:xs) = myLengthHelper (x:xs) 0

main :: IO()
main = return()
