-- written by copilot
compress :: Eq a => [a] -> [a]
compress [] = []
compress (x:xs) = x : (compress $ dropWhile (==x) xs)

-- written by ramgb
compressHelper :: Eq a => [a] -> a -> [a]
compressHelper [] _ = []
compressHelper (x:xs) y
  | x == y = compressHelper xs y
  | otherwise = x : (compressHelper xs x)

compress2 :: Eq a => [a] -> [a]
compress2 [] = []
compress2 (x:xs) = x : (compressHelper xs x)

main :: IO ()
main = return()


