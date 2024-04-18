# Concurrent Sudoku solver in GO
This program attempts to solve easy level sudoku puzzles using goroutines in go.

### Approach
- A gouroutine is created for each un-numbered cell with a state which is a map that holds all the 1 to 9 digits. 
- The goroutines will check their horizontal, vertical and 3X3 grid for available numbers and remove them from the state.
- As the state gets left with only one number, the gorutine will write the number in the sudoku (shared 2D 9X9 array) and Broadcast to all the waiting goroutines.
- Otherwise current goroutine will also wait for other goroutines to broadcast.
- If all the alive goroutines goes to waiting state the sudoku cannot be solved by this program.