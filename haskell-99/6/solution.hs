isPalindromeHelper :: Eq a => [a] -> [a] -> Bool
isPalindromeHelper [] [] = True
isPalindromeHelper (x:xs) (y:ys)
  | x == y    = isPalindromeHelper xs ys
  | otherwise = False

isPalindrome :: Eq a => [a] -> Bool
isPalindrome x = isPalindromeHelper x (reverse x)

main :: IO()
main = return()
