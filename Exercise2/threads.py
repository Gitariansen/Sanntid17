from threading import Thread
import threading

i = 0
lock = threading.Lock()
	


def thread_1_Function():
	global i
	for x in range(0,1000001):
		lock.acquire()
		i+=1
		lock.release()
	

def thread_2_Function():
	global i	
	for x in range(0,1000000):
		lock.acquire()
		i-=1
		lock.release()


def main():
	thread_1 = Thread(target = thread_1_Function, args = (),)
	thread_2 = Thread(target = thread_2_Function, args = (),)
	
	thread_1.start()
	thread_2.start()

	thread_1.join()
	thread_2.join()
	
	print(i)

main()
