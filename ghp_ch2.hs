calcChange owed given =
  if change > 0
    then change
    else 0
  where
    change = owed - given

doublePlusTwo x = doubleX + 2
  where
    doubleX = x + x
