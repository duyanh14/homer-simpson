from locust import HttpUser, task
import argparse

class HelloWorldUser(HttpUser):
    print("executing my_task")

    @task
    def userAccess(self):
        self.client.get("user/access")

