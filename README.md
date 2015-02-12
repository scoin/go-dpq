Generic Double Priority Queue
==============================

This was written in response to the 2015/02/11 reddit [/r/dailyprogrammer](http://www.reddit.com/r/dailyprogrammer/comments/2vkwgb/20150211_challenge_201_practical_exercise_get/) challenge.

There's probably not much need for this data type in the wild, but it was a fun exercise.

Usage
------

#####New

Returns a pointer to new dpQueue instance with zero values.

		myQueue = dpQueue.New()

#####Enqueue 

This method accepts three parameters - an generic item, priority value A, and priority value B, where the priority values are real numbers.

		myQueue.Enqueue("Cake", 4.99, 1280)
		myQueue.Enqueue("Pie", 7.99, 674)
		myQueue.Enqueue(42, 500, 7.8989)

######DequeueA 

This method removes and returns the string from the priority queue with the highest priority A value. If two entries have the same A priority, returns whichever was enqueued first.

Returns an error type if list is empty.

		myQueue.DequeueA() // returns 42

######DequeueB 

This method removes and returns the string from the priority queue with the highest priority B value. If two entries have the same B priority, returns whichever was enqueued first.

Returns an error type if list is empty.

		myQueue.DequeueB() // returns "Cake"

######Count 

This method returns the number of strings in the queue.

		myQueue.Count() // returns 1

######Clear 

Removes all entries from the queue.

		myQueue.Clear()
		fmt.Println(myQueue) >> {}
