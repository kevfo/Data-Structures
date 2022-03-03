
# Question 1

for i in [i**2 for i in range(0, 10) if i % 2 == 0]:
    print(i - 10 if i > 10 else i - 5 if i > 5 else i)


# Question 5

def fn1(rowNum):
    store = [[1], [1, 1]]

    if rowNum <= 1: return(store[rowNum])

    for i in range(3, rowNum + 1):
        item = [1] * i
        for j in range(1, item.length - 1):
            item[j] = store[-1][j] + store[-1][j - 1]
        store.append(item)

    return(store[-1])
