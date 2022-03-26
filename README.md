# Gilded Rose Kata

Kata challenge attempt. Go files copied from go folder in [here](https://github.com/emilybache/GildedRose-Refactoring-Kata).

## Approach

After some deliberation, I'm going with Table Driven Tests for the sake of simplicity, but it's probably better to split them by product and/or Item field in this case.

The last commit is a refactor of the function Update quality. I left commented the original one with the "Conjured" item update for faster checking. Both should work.

### Refactoring UpdateQuality function

Since there are very few items with specific conditions and quality variations are very small, I decided to use and reuse small incremental functions, but ideally there should be a parameter in the function signature with the value that is going to be added or substracted. It would also leave just 2 extra functions instead of the current 3 and simplify a bit switch-case part. This way would allow easier refactoring and/or updating functionality.

## Environment

* WSL2 with Ubuntu
* Go version -> 1.18
* Visual Studio Code
