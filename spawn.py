import gevent
import gevent.monkey
gevent.monkey.patch_socket()

import requests
import sys 


URL = 'http://127.0.0.1:8000/'
STEP = 100
MAX = 20000

def open_conn():
    try:
        resp = requests.get(URL)
    except Exception, exc:
	print exc
    print resp.text


if __name__ == '__main__':
    if len(sys.argv) > 1:
       URL = sys.argv[1]

    sum = 0
    while sum < MAX:

        threads = [gevent.spawn(open_conn) for i in xrange(STEP)]
        gevent.sleep(0.1)
        sum+=STEP
        print sum

    gevent.sleep(100)
