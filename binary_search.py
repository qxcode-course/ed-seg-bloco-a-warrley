def binary_search(arr, target):
    """
    Perform binary search on a sorted array to find the target value.
    
    Args:
        arr: A sorted list of elements
        target: The value to search for
        
    Returns:
        int: The index of the target if found, otherwise -1
    """
    low = 0
    high = len(arr) - 1
    
    while low <= high:
        mid = (low + high) // 2
        if arr[mid] == target:
            return mid
        elif arr[mid] < target:
            low = mid + 1
        else:
            high = mid - 1
    
    return -1

# Example usage
if __name__ == "__main__":
    # Test array (must be sorted)
    test_arr = [1, 3, 5, 7, 9, 11, 13, 15, 17, 19]
    
    # Test cases
    print(f"Searching for 7 in {test_arr}: Index = {binary_search(test_arr, 7)}")
    print(f"Searching for 1 in {test_arr}: Index = {binary_search(test_arr, 1)}")
    print(f"Searching for 19 in {test_arr}: Index = {binary_search(test_arr, 19)}")
    print(f"Searching for 4 in {test_arr}: Index = {binary_search(test_arr, 4)}")
    print(f"Searching for 20 in {test_arr}: Index = {binary_search(test_arr, 20)}")