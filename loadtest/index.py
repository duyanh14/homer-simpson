from locust import HttpUser, task,between
import argparse

class HelloWorldUser(HttpUser):
    wait_time = between(1, 5)

    def on_start(self):
         self.client.headers = {'Content-Type': 'application/xml', 'User-Agent': 'locust','Authorization': 'Bearer eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTk3Nzg5NDEsImlzcyI6InNpbXBzb24iLCJ1c2VybmFtZSI6IjFmZHNzMTFkMmR1YyIsInJlZmVzaF90b2tlbiI6IiIsInVzZXJfaWQiOjEsImVtYWlsIjoieHNmeHhAZ21haWwuY29tIiwicGhvbmUiOiIwMzMyMjgxNzc3In0.WcP-HtbjT0SCFKeIHNrgpwkrEhc_DYtp_hJvu4LGtA8cvkDqcjWPwY858Q-6RovQSCcBGZzC8EigFsfpqTy4GtiKDxJfSHnH9GU3S7Oo5uW6YqZPtrZKLK8jjNbYtdpA8hRwIeJ5iCmEYB9pVWElIRjm1jguBsXgfIqrXhPlpfEK-xhFcTXP7W5b2CZyL9iL896ys9bkiSj5H14VDQ3tC4FUwNnTqsaLnywdcTSFgmvYUjhJLdP9JfuYJhYz5WZJZDIUzlMNj7eY3sQntwdgoeu20pHKVygGYR6d2WyD5rGO1PWzI995QfoX1MnJWeMIjlR929taHJttfRizjg-KlnklyFfhuYYwcKQU-C2T1YGHFwKmTgrA1fCHvb3Rvf_KbFFPbmTyKlA3t5xHYDXyLBko9rp-RvZeU3gXZDE3zI8DksgtO5sJYzeNuZNFgcdktw1KuDVNc_AP0qNBJjz7P8efNtsLXdY4HhRsSNxQcixHVIljARA1hOdJijzf8rVhr3VBJCLYe4rYGDOt-NL3ivZGRqQTYeA1b92veP3xDTly_Cqh1oNgBl7rePYvPBbc3WE-e1k2u41C4H0vxQraUj-zyoZJNqVqX8qwe9vMtMBQPP8yFmMvAjsx3zyQ6OwNxQumzI7x2Jd4AeO3rgqV84cFqOkaakig_7E_8f9rI0o' }

    @task(1)
    def userAccess(self):
        response = self.client.get("user/access?code=CREATE1D")
        print(response)


