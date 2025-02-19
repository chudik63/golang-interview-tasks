# Solution
1) A type is created that represents a channel of the same type (O_o).

2) A buffered channel of length 1 is created.

3) The same channel is added to the channel (o_О).

4) It is important to understand that when multiple cases in a select-case statement are ready at the same time, one is chosen randomly.

5) On the first iteration, either the first or the second case will definitely trigger.

Case 1: Suppose the first case `case <-c`: triggers.\
We read a value from the channel => the channel is now empty, and on the next iteration the default case will trigger => we exit the loop with the output 1.

Case 2: Suppose the second case `case <-c: c <- c` triggers.\
We read from the channel and then put the same channel back into it => nothing changes. This process repeats.