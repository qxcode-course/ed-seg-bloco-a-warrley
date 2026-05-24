prime :: Int -> Bool
prime n = length [x | x <- [1..n], mod n x == 0] == 2

n_prime :: Int -> Int
n_prime n = last(take n [x | x <- [2..], prime x])

main :: IO ()
main = do 
  input <- getLine
  let n = read(input)
  print(n_prime(n))
