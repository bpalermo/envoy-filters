import http from 'k6/http';

export const options = {
  scenarios: {
    default: {
      executor: 'per-vu-iterations',
      vus: 10,
      iterations: 5000,
    },
  },
  summaryTrendStats: ['avg', 'min', 'med', 'max', 'p(95)', 'p(99)', 'p(99.9)'],
};

export default function () {
  http.get('http://localhost/lua/');
}
