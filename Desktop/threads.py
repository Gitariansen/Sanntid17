from threading import Thread

i = 0

	


def thread_1_Function():
	global i
	for x in range(0,1000000):
		i+=1
	print(i)
	

def thread_2_Function():
	global i	
	for x in range(0,1000000):
		i-=1
	print(i)


def main():
	print(i)
	thread_1 = Thread(target = thread_1_Function, args = (),)
	thread_2 = Thread(target = thread_2_Function, args = (),)
	
	thread_1.start()
	thread_2.start()

	thread_1.join()
	thread_2.join()
	
	print(i)

main()
