#include <pthread.h>
#include <stdio.h>
#include <stdlib.h>


int i = 0;
pthread_mutex_t lock;

void* thread_1Function() {

	int x = 0;
	while(x < 1000000){
		pthread_mutex_lock(&lock);
		i++;
		x++;
		pthread_mutex_unlock(&lock);
	}

	return NULL;
}

void* thread_2Function() {

	int x = 0;
	while(x < 1000000){
		pthread_mutex_lock(&lock);
		i--;
		x++;
		pthread_mutex_unlock(&lock);
	}

	return NULL;
}




int main() {
	pthread_t thread_1;
	pthread_create(&thread_1, NULL, thread_1Function, NULL);

	pthread_t thread_2;
	pthread_create(&thread_2, NULL, thread_2Function, NULL);

	pthread_join(thread_1, NULL);
	pthread_join(thread_2, NULL);

	printf("%d\n", i);

	pthread_mutex_destroy(&lock);
	return 0;
}
