# Othello AI
https://othello-asami.appspot.com/

## Strategy
Divide the game into 2 parts.

**1. part1 (number of pieces < 32)**  
``func selectNearCenter()``  
Place it in the center of the board as much as possible.

**2. part2 (number of pieces >= 32)**  
``func minmax()``  
Minimax method with a depth of 3.

