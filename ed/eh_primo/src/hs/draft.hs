main :: IO ()
prime :: Int -> Bool
prime n = length [x | x <- [1..n], mod n x == 0] == 2

main = do 
  input <- getLine
  if prime (read input) 
    then putStrLn "true"
  else putStrLn "false"
