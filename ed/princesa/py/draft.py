def printv(vector, size, sword):
    print("[ ", end="")
    for i in range(size):
        print(vector[i], end="")
        if i == sword:
            print(">", end="")
        print(" ", end="")
    print("]")

def main():
    size, sword = map(int, input().split()); sword -= 1
    vector = [i for i in range(1, size+1)]

    while size >= 1:
        printv(vector, size, sword)

        dead = (sword+1) % size
        for i in range(dead, size-1):
            vector[i] = vector[i+1]

        vector.pop()
        size -= 1

        if size != 0: sword = dead % size
    

if __name__ == "__main__": main()