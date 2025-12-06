def getInput():
    with open('input', 'r', encoding="utf-8") as inputFile:
        return inputFile.read()

def isEntirlyBelowRange(range1, range2):
    return range2[1] < range1[0]

def touchesRange(range1, range2):
    # Check if the ranges touch each other
    if range1[1] + 1 == range2[0] or range2[1] + 1 == range1[0]:
        return True

    # Check if the ranges intersect each other
    if (range2[0] >= range1[0] and range2[0] <= range1[1]
            or (range2[1] >= range1[0] and range2[1] <= range1[1])):
        return True

    # Check if on range containes another
    if (range1[0] < range2[0] and range1[1] > range2[1]) or \
        (range2[0] < range1[0] and range2[1] > range1[1]):
        return True

    return False

def getUpdatedSortedRanges(sortedRanges, newRange):
    i = 0
    while i < len(sortedRanges):
        currentRange = sortedRanges[i]
        if (touchesRange(currentRange, newRange)):
            followingIntersections = 0
            while i + followingIntersections + 1 < len(sortedRanges):
                if touchesRange(sortedRanges[i + followingIntersections + 1], newRange):
                    followingIntersections += 1
                else:
                    break

            combinedRange = [
                min(currentRange[0], newRange[0]),
                max(sortedRanges[i + followingIntersections][1], newRange[1])
            ]

            newSortedRanges = sortedRanges[: i] \
                + [combinedRange] \
                + sortedRanges[i + followingIntersections + 1:]
            return newSortedRanges
        if (isEntirlyBelowRange(currentRange, newRange)):
            return sortedRanges[: i] + [newRange] + sortedRanges[i :]

        i += 1

    updatedRange = sortedRanges + [newRange]
    return updatedRange

def isInSortedRanges(sortedRanges, id):
    for rangeItem in sortedRanges:
        if(id >= rangeItem[0] and id <= rangeItem[1]):
            return True
    return False

if __name__ == "__main__":
    input = getInput()
    lines = input.splitlines()

    availableFreshIngredientsCount = 0
    sortedRanges = []
    hasProcessedRanges = False
    for line in lines:
        if (len(line) == 0):
            hasProcessedRanges = True
            continue

        if (hasProcessedRanges):
            availableIngredientId = int(line)
            if isInSortedRanges(sortedRanges, availableIngredientId):
                availableFreshIngredientsCount += 1
        else:
            [rangeStart, rangeEnd] = [int(rangeStr) for rangeStr in line.split('-')]
            idRange = [rangeStart, rangeEnd]
            sortedRanges = getUpdatedSortedRanges(sortedRanges, idRange)

    allFreshIdsCount = 0

    for idRange in sortedRanges:
        rangeIdsCount = idRange[1] - idRange[0] + 1
        allFreshIdsCount += rangeIdsCount

    print(f'{availableFreshIngredientsCount} fresh ingredients are stored.')
    print(f'In total there are {allFreshIdsCount} fresh ingredients available.')

