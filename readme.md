# Chat

This chat was made with C for the client and Golang for the server. It was made to learn the foundation of online applications, queues, and packet structuring. Some things are hard coded because this is suppossed to be a practice project.

To use it, you have to compile the client using the Makefile (preferably with the MinGW's GCC) in the _Client_ folder, and the server using the _go build_ command in the server folder. 

In order to use it you have two options:

1. First you can open the port 3200 (you can change it, but it must be the same in the client.c and the main.go) in your router, write your ip ([find it using this link](https://www.myip.com/)) in the **#define LOCAL_ADDR "0.0.0.0"** in the [client](Client/src/client.c), build it, send the client executable to your friends and finally run the server executable on your computer.
   
2. It also will work on localhost if the **LOCAL_ADDR** value is **127.0.0.1**, but you can only connect to the server from your own computer.

This code it's free. You can use it however you want, but I'll appreciate if you give me credits. Thank you for visiting this repo.

