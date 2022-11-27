from diagrams import Cluster, Diagram
from diagrams.onprem.network import Nginx
from diagrams.k8s.network import Service
from diagrams.k8s.compute import Pod
from diagrams.onprem.database import MongoDB
from diagrams.onprem.inmemory import Redis

with Diagram("architecture diagram", show=False):

    ingress = Nginx("ingress")

    serviceBE = Service("service")

    podBE = [
        Pod("backend1"),
        Pod("backend2"),
        Pod("backend3")
    ]

    serviceDB = Service("service")
    
    mongoDB = MongoDB("mongoDB")
    
    serviceMemory = Service("service")

    redis = Redis("redis")

    serviceAPIContract = Service("service")

    podAPIContract = Pod("swagger")
        
    ingress >> serviceBE >> podBE
    ingress >> serviceAPIContract >> podAPIContract
    podBE >> serviceDB >> mongoDB
    podBE >> serviceMemory >> redis