# Mutex and Channel basics

### What is an atomic operation?
> Informally, an operation that happens in a single step relative to other threads. This definition is often extended, especially in relation to database operations, to include operations which can be caught in an intermediate state by another thread. For example, consider some data being written to a file. This operation would of course take many steps, however, we can make the operation "atomic" in this extended sense by writing to a temporary file first, and then moving this file to its intended location. The move operation is atomic in most systems.

### What is a semaphore?
> A semaphore is a variable used for signaling that a thread is using a certain resource. A simple semaphore can be for instance a flag or a counter which is toggled or decremented. When this thread is finished using the resource, the semaphore is reset, signaling to other threads that the resource is safe to use. 

### What is a mutex?
> A Mutex (as in mutual exclusion) is a locking mechanism ensuring that two threads do not enter their _critical section_ at the same time. 

### What is the difference between a mutex and a binary semaphore?
> Mutexes are similar to binary semaphores, but have somewhat different use cases. Mutexes are used to lock down a resource. The _owner_ of the mutex then has access to the resource and all other threads have to wait. A binary semaphore on the other hand is used to signal to other threads that an event has occured. 

### What is a critical section?
> A critical section for some subroutine is one where it _really_ should not be interfered with by another thread. In most cases, this is relating to shared memory. The critical section will then be the part where a subroutine is performing some read and write operations over possibly many cycles, at which some other program might read or write to the same data while it's in an intermediate state.

### What is the difference between race conditions and data races?
> Data races are when two or more processes try to access the same data simultaneously and in an unpredictable way and at least one of them is a write operation. Race conditions on the other hand are any situation where the ordering of a (multiprocess) routine's operations affect the outcome of that routine.

### List some advantages of using message passing over lock-based synchronization primitives.
> For many tasks, it results in much simpler, more understandable code, less prone to errors. Data races are also effectively circumvented. 

### List some advantages of using lock-based synchronization primitives over message passing.
> In some situations, a single global variable may be much less complex than a message passing alternative. Also, if the state to be represented is big, message passing can become very tedious. For example, a database is effectively a huge store of shared variables. Handled by systematic synchronization mechanisms to avoid data races.
