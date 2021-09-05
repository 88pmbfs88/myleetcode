def TwoSum(arr, target):
    arr_len = len(arr)
    result = []
    for i in range(arr_len):
        for j in range(i+1, arr_len):
            if arr[i] + arr[j] == target:
                result.append([i, j])
    return result


if __name__ == "__main__":
    arr = [2, 3, 10, 12, 8, 4]
    target = 12
    result = TwoSum(arr, target)
    print(result)
