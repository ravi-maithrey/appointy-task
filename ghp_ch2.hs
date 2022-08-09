calcChange owed given =
  if change > 0
    then change
    else 0
  where
    change = owed - given