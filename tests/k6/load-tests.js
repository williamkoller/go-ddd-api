import http from 'k6/http';
import { check, sleep } from 'k6';
import { uuidv4 } from 'https://jslib.k6.io/k6-utils/1.4.0/index.js';

export const options = {
    stages: [
        { duration: '30s', target: 100 },  // Sobe rapidamente para 100 VUs
        { duration: '1m', target: 200 },   // Aumenta para 200 VUs
        { duration: '2m', target: 200 },   // Mantém 200 VUs
        { duration: '30s', target: 0 },    // Reduz gradualmente para 0 VUs
    ],
    thresholds: {
        http_req_duration: ['p(95)<5000'],   // 95% das requisições devem ser < 5000ms
        http_req_failed: ['rate<0.01'],     // Falhas devem ser < 1%
    },
};

const BASE_URL = 'http://localhost:3003'; // Altere se necessário

export default function () {
    const payload = JSON.stringify({
        name: `Heavy User ${__VU}-${__ITER}`,
        email: `${uuidv4()}@loadtest.com`,
        password: 'LoadTest#2025',
    });

    const headers = {
        'Content-Type': 'application/json',
    };

    const res = http.post(`${BASE_URL}/users/register`, payload, { headers });

    check(res, {
        'status is 201': (r) => r.status === 201,
        'response is JSON': (r) => r.headers['Content-Type'].includes('application/json'),
    });

    sleep(0.5); // Menor tempo entre iterações para aumentar intensidade
}
