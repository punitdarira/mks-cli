# mks-cli

This CLI reads mks.yaml file and creates k8s namespace, deployment, loadbalancer and a grafana dashboard with 4 graphs (cpu, memory, disk and loki logs)


# mks.yaml file
This file has 4 properties.
app, image, replica and port <br/>
app- A new namespace would be created by the name of this app property. This property would also be used to create a grafana dashboard  <br/>
A container with specified image would be created <br/>
replica- no of replicas <br/>
port- the loadbalancer would point to this port of the container. <br/>



