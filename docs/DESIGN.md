## Design Docs 

### Concepts 

#### Controller

Controllers are the CloudScale control plane. They handle configuration changes, reload routing and load balancing, and 
handle proxying traffic to the destination servers 

#### Agent 

An optional component that can be deployed to backend servers, which transmits current resource usage data to the controllers. 
Used for the Resource Usage strategy

#### Listeners 

An incoming endpoint that receives connections

#### Target Groups 

Configured destinations that incoming [Listener](#listeners-) traffic is routed to
