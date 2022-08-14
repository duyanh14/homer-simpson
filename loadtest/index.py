from locust import HttpUser, task,between
import argparse

class HelloWorldUser(HttpUser):
    wait_time = between(1, 5)

    def on_start(self):
         self.client.headers = {'Content-Type': 'application/xml', 'User-Agent': 'locust','Authorization': 'Bearer eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjAwNTAyNjUsImlzcyI6InNpbXBzb24iLCJ1c2VybmFtZSI6IjFmZHNzMTFkMmR1YyIsInJlZmVzaF90b2tlbiI6IiIsInVzZXJfaWQiOjEsImVtYWlsIjoieHNmeHhAZ21haWwuY29tIiwicGhvbmUiOiIwMzMyMjgxNzc3In0.SNUVPBLy_ftqUnV5KaRA4p4KuHsC6HPNlUCD0PqLMsW40h80wOg0sEOjvfrtXmUZc78hsRE_dr4TAQVUnjQqS2zRo0wVMQDwr_7zm6hwnU1l0iyRulQptWszk2R3eAbeOeZXz3nEFL6ZYgiqRu7M0vReZrcbEGwhTMeT4sMvaJfv7DtUjfPJp5zZ7zzfyqNdVNGl5z3nUR9O_FYhSyejxsckEShWURfo3V0-PUKFGkUGPVrQfkfQfHK4r7vEBRaePnGoOo5LLYVau7d0hSe-u6X-L9cuj13mMUqSVkfIwCDJB_R-CAV8ysoMKNLkj07GWNeqlIF96tLSxeFB12ybuFTOsbfibREvoxCLyd-llrvo8-hWZ8n3PfDbEek8hFX8arvhuzJ3VMnxYAKfOHGVbOF4tsFE6VSwbRoAmnrMyEU0uFQDNfzpSs4emRjKDuExzj_K4tk_-fyvHhsoAcW74_fwQI7QoWfTnsdIudW3TbqvY6VIffY4B_fgsxWh8nLQrQjJ_M-Zp2RTKN6ky6AgsNqnn24c9MXunjp7PxsVLzDecIfTuICt-cnYG-F1aFQjiziMi_oe2NaJpoIIpNx9OdjaQq4l5EX3U6sdEg8GlG_wCnM0BuK9vcAmf8uivWTciluuvpm5IqTWtWsb_PvsjL_vvCgAsh8tSXuOaK2DoeU' }

    @task(1)
    def userAccess(self):
        response = self.client.get("user/access?code=CREATE1D")
        print(response)


