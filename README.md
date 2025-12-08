We found that for our advanced program, the concurrent approach is actually slower.  

This is due to the backend management that has to occur when managing concurrent programs, ex mutexes and synchronization.

Our concurrent approach would have yielded better results if there were more than 4 branching paths per node.  This is where bfs really excels.


