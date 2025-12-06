def isIllegal(id):
    idStr = str(id)
    sliceLength = 1

    while(sliceLength <= len(idStr) // 2):
        if (len(idStr) % sliceLength != 0):
            sliceLength += 1
            continue

        slice = idStr[-sliceLength:]
        sliceVal = int(slice)

        if (sliceVal == 0):
            sliceLength += 1
            continue

        idRem = id // (10 ** sliceLength)

        while (idRem % (10 ** sliceLength) == sliceVal):
            idRem = idRem // (10 ** sliceLength)

            if (idRem == 0):
                return True

        sliceLength += 1


    return False


if __name__ == "__main__":
    with open('input') as inputFile:
        input = inputFile.read()

    res = 0

    for rangeStr in input.split(','):
        [lowerBound, upperBound] = rangeStr.split('-')

        for id in range(int(lowerBound), int(upperBound) + 1):
            if isIllegal(id):
                res += id

    print (res)

