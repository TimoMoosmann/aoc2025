def clickDial(oldDial, clicks, isDirectionRight):
    zeroVisits = 0
    if isDirectionRight:
        dial = oldDial + clicks
        while (dial >= 100):
            zeroVisits += 1
            dial = dial - 100
    else:
        dial = oldDial - clicks
        while (dial < 0):
            dial = 100 + dial
            if (oldDial != 0):
                zeroVisits += 1

            oldDial = dial

        if (dial == 0):
            zeroVisits += 1
    return (dial, zeroVisits)

if __name__ == "__main__":
    dial = 50
    
    with open('input') as f:
        input = f.read()

    zeroHolds = 0
    zeroVisits = 0
    
    for line in input.splitlines():
        isDirectionRight = line[0] == "R"
        clicks = int(line[1:])
        (dial, newVisits) = clickDial(dial, clicks, isDirectionRight)
        zeroVisits += newVisits

        if (dial == 0):
            zeroHolds += 1

    print(zeroVisits)
