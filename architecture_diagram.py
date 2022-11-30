from diagrams import Cluster, Diagram
from diagrams.onprem.network import Nginx
from diagrams.k8s.network import Service
from diagrams.k8s.compute import Pod
from diagrams.onprem.database import MySQL
from diagrams.onprem.inmemory import Redis

with Diagram("architecture diagram", show=False):

    ingress = Nginx("ingress")

    serviceBE = Service("service")

    podBE = [
        Pod("backend1"),
        Pod("backend2"),
        Pod("backend3")
    ]

    
    serviceDBSecurity = Service("service")
    mysqlDBSecurity = MySQL("db security")

    serviceDBUser = Service("service")
    mysqlDBUser = MySQL("db user")
    
    serviceMemory = Service("service")

    redis = Redis("redis")

    serviceAPIContract = Service("service")

    podAPIContract = Pod("swagger")
        
    ingress >> serviceBE >> podBE
    ingress >> serviceAPIContract >> podAPIContract
    podBE >> serviceDBSecurity >> mysqlDBSecurity
    podBE >> serviceDBUser >> mysqlDBUser
    podBE >> serviceMemory >> redis