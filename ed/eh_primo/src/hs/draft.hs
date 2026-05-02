main :: IO ()
main = do 
  input <- getLine 
  let n = read input :: Int 
  putStrLn (if prime n then "true" else "false")

prime :: Int -> Bool
prime n 
  | n <= 1 = False
  | otherwise = length [x | x <- [2..(n-1)], n `mod` x == 0] == 0
