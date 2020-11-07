#include <winsock2.h>
#include <windows.h>
#include <stdlib.h>
#include <string.h>
#include <stdio.h>
#include <conio.h>

#define LOG_IN 0x01
#define HEART_BEAT 0x02
#define MESSAGE 0x03
#define DISCONNECT 0x04

#define NEW_CONN 0x0A
#define NEW_DISC 0x0B

#define LOCAL_ADDR "127.0.0.1"
#define PORT 3200

typedef DWORD WINAPI thread_f;

const char *STR_CH = "->";

void disconnectClient(){
    system("cls");
    printf("Server Crashed...\n");
    Sleep(5000);
    exit(0);
}

thread_f HeartBeat(void *lparam){
    SOCKET client = *(SOCKET *)lparam;

    for(;;){
        char hb_packet[2] = {HEART_BEAT, '\n'};
        send(client, hb_packet, 2, 0);
        Sleep(100);
    }
}

thread_f Receive(void *lparam){
    SOCKET client = *(SOCKET *)lparam;
    HANDLE console = GetStdHandle(STD_OUTPUT_HANDLE);
    for(;;){
        char buffer[255];
        recv(client, buffer, 255, 0);
        if(strlen(buffer) != 0){
            switch (buffer[0])
            {
                case MESSAGE:
                    SetConsoleTextAttribute(console, FOREGROUND_GREEN);
                    printf("\r%s: ", buffer + 1);
                    int len = strlen(buffer);
                    SetConsoleTextAttribute(console, 0x0007);
                    printf("%s\n", buffer + 1 + len);
                    printf(STR_CH);
                    break;

                case NEW_CONN:
                    SetConsoleTextAttribute(console, FOREGROUND_GREEN);
                    printf("\r%s joined the chat.\n", buffer + 1);
                    SetConsoleTextAttribute(console, 0x0007);
                    printf(STR_CH);
                    break;

                case NEW_DISC:
                    SetConsoleTextAttribute(console, FOREGROUND_RED);
                    printf("\r%s left the chat.\n", buffer + 1);
                    SetConsoleTextAttribute(console, 0x0007);
                    printf(STR_CH);
                    break;
                    break;
                    
                default:
                    break;
            }
            
        }
        else{
            disconnectClient();
        }
        Sleep(100);
    }
    
}

int main(){
    WSADATA data;
    WSAStartup(MAKEWORD(2,2), &data);

    SOCKET sock = socket(AF_INET, SOCK_STREAM, IPPROTO_TCP);
    
    SOCKADDR_IN server_addr;
    server_addr.sin_addr.s_addr = inet_addr(LOCAL_ADDR);
    server_addr.sin_family = AF_INET;
    server_addr.sin_port = htons(PORT);

    int cs = connect(sock, &server_addr, sizeof(server_addr));
    if(cs == SOCKET_ERROR){
        printf("Cannot connect.");
        getchar();
        return -1;
    }

    printf("Connected\n");

    char packet[32];
    char name[30];
    printf("Type your name: ");
    gets(name);
    system("cls");
    sprintf(packet, "%c%s\n", LOG_IN, name);
    send(sock, packet, strlen(packet), 0);
    
    HANDLE HB = CreateThread(NULL, 0, HeartBeat, &sock, 0, NULL);  
    HANDLE RC = CreateThread(NULL, 0, Receive, &sock, 0, NULL);  

    for(;;){
        char msg[255];
        char snd[255];
        printf(STR_CH);
        gets(msg);

        if(!strcmp(msg, "/exit")){
            sprintf(snd, "%c\n", DISCONNECT);
            send(sock, snd, strlen(snd), 0);
            break;
        }
        else{
            sprintf(snd, "%c%s\n", MESSAGE, msg);
            send(sock, snd, strlen(snd), 0);
        }
    }  

    return 0;
}