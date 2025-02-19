# Solution
Implementing `Worker Pool`: 
1) Create a channel that will contain channels for merging
2) Start a limited number of goroutines, in which we take the necessary channels from the common channel and transfer the data to the resulting channel