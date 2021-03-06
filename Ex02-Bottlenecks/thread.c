// gcc 4.7.2 +
// gcc -std=gnu99 -Wall -g -o helloworld_c helloworld_c.c -lpthread

#include <pthread.h>
#include <stdio.h>
#include <stdlib.h>

pthread_mutex_t mutex;
int i = 2;

// Note the return type: void*
void* someThreadFunction1(){
    for (int j = 0; j < 1000023; j++) {
	pthread_mutex_lock(&mutex);
	i++;
	pthread_mutex_unlock(&mutex);
    }
    return NULL;
}

void* someThreadFunction2(){
    for (int j = 0; j < 1000000; j++) {
	pthread_mutex_lock(&mutex);
	i--;
	pthread_mutex_unlock(&mutex);
    }
    return NULL;
}



int main(){
    pthread_mutex_init(&mutex, NULL);

    pthread_t someThread1;
    pthread_t someThread2;
    pthread_create(&someThread1, NULL, someThreadFunction1, NULL);
    pthread_create(&someThread2, NULL, someThreadFunction2, NULL);
    
    pthread_join(someThread1, NULL);
    pthread_join(someThread2, NULL);
    printf("%d\n", i);

    pthread_mutex_destroy(&mutex);
    return 0;
    
}
