prime :: Int -> Bool
prime n = length [x | x <- [1..n], mod n x == 0] == 2

prime_list :: Int -> [Int]
prime_list n = take n [x | x <- [2..], prime x]

formatList :: [Int] -> String
formatList xs = "[" ++ go xs ++ "]"
  where
    go [] = ""
    go [x] = show x
    go (x:xs) = show x ++ ", " ++ go xs

main :: IO ()
main = do
  input <- getLine
  putStrLn (formatList (prime_list (read input)))
