fibRabbits :: Int -> Int -> Int
fibRabbits 1 _ = 1
fibRabbits 2 _ = 1
fibRabbits n k = fibRabbits(n-1) k + k * fibRabbits(n-2) k 

main :: IO ()
main = do 
  line <- getLine
  let [n, k] = map read (words line)
  print (fibRabbits n k)
