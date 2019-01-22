# Choosing a language

In Exercise 3, 4 (Network exercises) and the project, you will be using a language of your own choice. You are of course free to change your mind at any time, but to help avoid this situation (and all its associated costs) it is worth doing some research already now.

You should start by looking at the [Programming Language part of the Project](https://github.com/TTK4145/Project/tree/master#programming-language). Send in your suggestions if you find more and/or better resources.

Here are a few things you should consider:
 - Think about how want to move data around (reading buttons, network, setting motor & lights, state machines, etc). Do you think in a shared-variable way or a message-passing way? Will you be using concurrency at all?
Message passing is the way to go.
 - How will you split into modules? Functions, objects, threads? Think about what modules you need, and how they need to interact. This is an iterative design process that will take you many tries to get "right" (if such a thing even exists!).
    - A Client sending things to the network.
    - A Network node receiving state from the network.
    - A elevator controller whose single responsibility is receiving "GOTO floor X" and doing that in the most efficient way in relation to other orders. 
        - Not allowed to switch direction if it has a GOTO in "front" of itself, but it can stop at floors
        - Should broadcast 
        - Keep a slice of GOTOs sorted after priority. (Priority depends on the __current elevator direction__)
(Maybe combine some of these)
 - The networking part is often difficult. Can you find anything useful in the standard libraries, or other libraries?
The network module for Go can be imported as a package. Raft? – favors consistency Mqtt? – needs a broker (kubernetes??) We need to favor availability over consistency. 
 - While working on new sections on the project you'll want to avoid introducing bugs to the parts that already work properly. Does the language have a framework for making and running tests, or can you create one? Testing multithreaded code is especially difficult.
Golang has a great testing package
 - Code analysis/debugging/IDE support?
Goland <3
