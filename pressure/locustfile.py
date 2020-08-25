# -*- coding:utf-8 -*-
# from __future__ import absolute_import
# from __future__ import unicode_literals
# from __future__ import print_function

import json
import time
import gevent

from websocket import create_connection

from locust import HttpUser, TaskSet, task, events

# class DanmuTestSet(TaskSet):
#     def on_start(self):
#         self.rid = 1
#         self.uid = 1

#         ws = create_connection('ws://127.0.0.1/ws/' +
#                                str(self.uid)+'/'+str(self.rid))
#         self.ws = ws
#         def _receive():
#             while True:
#                 res = ws.recv()
#                 data = json.loads(res)
#                 end_at = time.time()
#                 response_time = int((end_at - data['start_at']) * 1000000)

#                 events.request_success.fire(
#                     self=self,
#                     request_type='ws',
#                     name='test/ws/message',
#                     response_time=response_time,
#                     response_length=len(res),
#                 )
#         gevent.spawn(_receive)

#     def on_quit(self):
#         self.ws.close()

#     @task
#     def sent(self):
#         data = {
#             'uid': self.uid,
#             'rid': self.rid,
#             'content': 'test',
#             'color': 'red'
#         }
#         self.client.post("/login", json=data)
#     #     request_success.fire(
#     #         request_type='WebSocket Sent',
#     #         name='test/ws/echo',
#     #         response_time=int((time.time() - start_at) * 1000000),
#     #         response_length=len(body),
#     #     )


class DanmuLocust(HttpUser):
    # task_set = DanmuTestSet
    min_wait = 0
    max_wait = 500
    def on_start(self):
        self.rid = 1
        self.uid = 1

        ws = create_connection('ws://localhost:8888/ws/' +
                               str(self.uid)+'/'+str(self.rid))
        self.ws = ws
        def _receive():
            while True:
                res = ws.recv()
                response_time = int(1)

                events.request_success.fire(
                    request_type='ws',
                    name='test/ws/message',
                    response_time=response_time,
                    response_length=len(res),
                )
        gevent.spawn(_receive)

    def on_quit(self):
        self.ws.close()

    @task
    def sent(self):
        data = {
            'uid': self.uid,
            'rid': self.rid,
            'content': 'test',
            'color': 'red'
        }
        self.client.post("/danmu/send", json=data)
