import http from 'k6/http';
import { check, sleep } from 'k6';

export let options = {
  vus: 1000,     // Number of virtual users
  duration: '30s',  // Duration of the test
};

export default function () {
  let res = http.get('http://localhost:8090/increment');

  // Optional: Add checks to validate response
  check(res, {
    'status is 200': (r) => r.status === 200,
    // 'body contains Counter': (r) => r.body.includes('Counter'),
  });

  // Optional: Add think time between requests
  // sleep(0.1); // 100ms delay between requests
}