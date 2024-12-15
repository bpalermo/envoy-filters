# Envoy filters
Envoy filter example with load tests

# Load testing

## Lua Filter
```bash
$ k6 run lua/.loadtest/script.js

         /\      Grafana   /‾‾/  
    /\  /  \     |\  __   /  /   
   /  \/    \    | |/ /  /   ‾‾\ 
  /          \   |   (  |  (‾)  |
 / __________ \  |_|\_\  \_____/ 

     execution: local
        script: lua/.loadtest/script.js
        output: -

     scenarios: (100.00%) 1 scenario, 10 max VUs, 10m30s max duration (incl. graceful stop):
              * default: 5000 iterations for each of 10 VUs (maxDuration: 10m0s, gracefulStop: 30s)


     data_received..................: 6.4 MB 2.0 MB/s
     data_sent......................: 4.0 MB 1.2 MB/s
     http_req_blocked...............: avg=5.57µs   min=856ns    med=4.52µs   max=729.88µs p(95)=8.51µs   p(99)=21.68µs  p(99.9)=239.6µs 
     http_req_connecting............: avg=56ns     min=0s       med=0s       max=339.47µs p(95)=0s       p(99)=0s       p(99.9)=0s      
     http_req_duration..............: avg=492.2µs  min=118.05µs med=468.44µs max=15.94ms  p(95)=781.11µs p(99)=1.27ms   p(99.9)=1.92ms  
       { expected_response:true }...: avg=492.2µs  min=118.05µs med=468.44µs max=15.94ms  p(95)=781.11µs p(99)=1.27ms   p(99.9)=1.92ms  
     http_req_failed................: 0.00%  0 out of 50000
     http_req_receiving.............: avg=50.51µs  min=7.91µs   med=41.12µs  max=1.14ms   p(95)=81.86µs  p(99)=382.33µs p(99.9)=742.21µs
     http_req_sending...............: avg=15.86µs  min=2.56µs   med=12.67µs  max=1.03ms   p(95)=26.04µs  p(99)=66µs     p(99.9)=537.82µs
     http_req_tls_handshaking.......: avg=0s       min=0s       med=0s       max=0s       p(95)=0s       p(99)=0s       p(99.9)=0s      
     http_req_waiting...............: avg=425.82µs min=98.33µs  med=407.54µs max=15.28ms  p(95)=664.34µs p(99)=1ms      p(99.9)=1.69ms  
     http_reqs......................: 50000  15445.550763/s
     iteration_duration.............: avg=576.07µs min=147.75µs med=550.02µs max=17.08ms  p(95)=920.07µs p(99)=1.5ms    p(99.9)=2.33ms  
     iterations.....................: 50000  15445.550763/s
     vus............................: 4      min=4          max=10
     vus_max........................: 10     min=10         max=10


running (00m03.2s), 00/10 VUs, 50000 complete and 0 interrupted iterations
default ✓ [======================================] 10 VUs  00m03.2s/10m0s  50000/50000 iters, 5000 per VU
```

## External Processor

```bash
$ k6 run ext-proc/.loadtest/script.js

         /\      Grafana   /‾‾/  
    /\  /  \     |\  __   /  /   
   /  \/    \    | |/ /  /   ‾‾\ 
  /          \   |   (  |  (‾)  |
 / __________ \  |_|\_\  \_____/ 

     execution: local
        script: ext-proc/.loadtest/script.js
        output: -

     scenarios: (100.00%) 1 scenario, 10 max VUs, 10m30s max duration (incl. graceful stop):
              * default: 5000 iterations for each of 10 VUs (maxDuration: 10m0s, gracefulStop: 30s)


     data_received..................: 6.8 MB 1.7 MB/s
     data_sent......................: 4.0 MB 978 kB/s
     http_req_blocked...............: avg=2.88µs   min=946ns    med=2.17µs   max=659.1µs  p(95)=6.12µs  p(99)=9.43µs   p(99.9)=53.57µs 
     http_req_connecting............: avg=48ns     min=0s       med=0s       max=304.12µs p(95)=0s      p(99)=0s       p(99.9)=0s      
     http_req_duration..............: avg=571.92µs min=256.29µs med=478.6µs  max=19.65ms  p(95)=1.3ms   p(99)=1.6ms    p(99.9)=2.08ms  
       { expected_response:true }...: avg=571.92µs min=256.29µs med=478.6µs  max=19.65ms  p(95)=1.3ms   p(99)=1.6ms    p(99.9)=2.08ms  
     http_req_failed................: 0.00%  0 out of 50000
     http_req_receiving.............: avg=26.72µs  min=7.54µs   med=21µs     max=757.44µs p(95)=64.5µs  p(99)=134.28µs p(99.9)=335.95µs
     http_req_sending...............: avg=7.84µs   min=2.68µs   med=5.93µs   max=1.48ms   p(95)=17.67µs p(99)=30.53µs  p(99.9)=183.92µs
     http_req_tls_handshaking.......: avg=0s       min=0s       med=0s       max=0s       p(95)=0s      p(99)=0s       p(99.9)=0s      
     http_req_waiting...............: avg=537.35µs min=235.4µs  med=449.51µs max=19.43ms  p(95)=1.22ms  p(99)=1.51ms   p(99.9)=1.98ms  
     http_reqs......................: 50000  12378.153232/s
     iteration_duration.............: avg=616.51µs min=277.22µs med=516.39µs max=20.56ms  p(95)=1.4ms   p(99)=1.71ms   p(99.9)=2.25ms  
     iterations.....................: 50000  12378.153232/s
     vus............................: 4      min=4          max=10
     vus_max........................: 10     min=10         max=10


running (00m04.0s), 00/10 VUs, 50000 complete and 0 interrupted iterations
default ✓ [======================================] 10 VUs  00m04.0s/10m0s  50000/50000 iters, 5000 per VU
```