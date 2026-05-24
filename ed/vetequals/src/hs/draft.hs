equals :: [Int] -> [Int] -> Bool
equals [] [] = True
equals [] y = False
equals x [] = False
equals (x:xs) (y:ys)
  | x == y = equals xs ys
  | otherwise = False
-- equals (x:xs) (y:ys) = (x == y)
 
parse :: String -> [Int]
parse s = map read (words (filter isValid s))
  where 
    isValid c = c /= '[' && c /= ']'

main :: IO ()
main = do
  a <- getLine 
  b <- getLine

  let v1 = parse a
  let v2 = parse b

  if equals v1 v2 then putStrLn "iguais" else putStrLn "diferentes"
