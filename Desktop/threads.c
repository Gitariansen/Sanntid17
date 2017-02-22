#include <pthread.h>
#include <stdio.h>


int i = 0;

void* thread_1Function() {
	int x = 0;
	while(x < 1000000){
		i++;
		x++;
	}
	return NULL;
}

void* thread_2Function() {
	int x = 0;
	while(x < 1000001){
		i--;
		x++;
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

	return 0;
}
