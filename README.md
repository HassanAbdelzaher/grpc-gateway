# gRPC Gateway
## Project Goal

Build a transparent reverse proxy for gRPC targets that will make it easy to expose gRPC services
over the Internet.
This includes:
* route grpc-web requests from browser to gRPC microservices
* no needed knowledge of the semantics of requests exchanged in the call (independent rollouts)
* easy declarative definition of backends and their mappings to frontends
* simple round-robin load balancing of inbound requests from a single connection to multiple backends

# Configurations
* host : define the gateway binding address
* http_port: define which http port will the gateway binding to
* tls_port: define the tls port which the gateway will binding to
*  run_tls_server : control if the gateway will listen on http or not
*  run_http_server: control if the gateway will listen on tls or not
* allow_all_origin : if it is true all origins will be allowed
* allowed_origins: define allowed origin in case of allow_all_origins is false
* allowed_headers: define the headers the will transfeer from the http1.1 request to http2 request header
* use_websockets used to control if the gateway will accept request of websocket transport
* max_call_recv_msg_size : maximum size of protobuff message
* websocket_ping_interval : timeout for ping signal to websockets clients
* http_max_write_timeout : max write timeout 
* http_max_read_timeout: max read timeout
* services define array of microservice of the system
* * service_name:define the service name as package.servicename
## BackEnd Configurations
* * is_using_tls: define is microservice using tls
* * tls_no_verify:Whether to ignore TLS verification checks (cert validity, hostname)
* * tls_client_cert:Path to the PEM certificate used when the backend requires client certificates for TLS.
* * tls_client_key:Path to the PEM key used when the backend requires client certificates for TLS.
* * max_call_recv_msg_size:Maximum receive message size limit. If not specified, the default of 4MB will be used
* * tls_ca:Paths (comma separated) to PEM certificate chains used for verification of backend certificates. If empty, host CA chain will be used.
* * default_authority:Default value to use for the HTTP/2 :authority header commonly used for routing gRPC calls through a backend gateway.
* * backoff_max_delayMaximum delay when backing off after failed connection attempts to the backend.

# Instllation
Build from source files or using Docker image

# Using
* first create javascript client using protoc 
```
var gatewayAddress="http://localhost:8089"
const __client=new StoreManagerClient(gatewayAddress);
```
* second create javascript function for each grpc method
```
export const sotresCount=()=>{
  return new Promise<number>((resolve,reject)=>{
    let req=new StoresCountRequest();
    req.setCustomerId(customerId)
    req.setIncludeDeleted(true)
    const client=getClient();
    client.storesCount(req,(err,resp)=>{
      if(err!=null){
        reject(err);
      }
      else{
        let count:number=resp?.getCount()||0;
        resolve(count);
      }
    })
  });
}
```
## config.json sample
```batch
{
  "host": "0.0.0.0",
  "http_port": 8089,
  "tls_port": 8090,
  "run_tls_server": false,
  "run_http_server": true,
  "allow_all_origin": true,
  "allowed_origins": [],
  "allowed_headers": [],
  "use_websockets": false,
  "max_call_recv_msg_size": 5242880,
  "websocket_ping_interval": "10s",
  "http_max_write_timeout": "10s",
  "http_max_read_timeout": "10s",
  "services": [{
    "service_name":"storemanager.StoreManager",
    "backends": [{
      "address": "host.docker.internal:8020",
      "is_using_tls": false,
      "tls_no_verify": true,
      "tls_client_cert": null,
      "tls_client_key": null,
      "max_call_recv_msg_size": 5242880,
      "tls_ca": [],
      "default_authority": null,
      "backoff_max_delay": null
    }]
  }]
}
```
Note:for testing you can use
BloomRPC application which support both grpc and grpc-web request