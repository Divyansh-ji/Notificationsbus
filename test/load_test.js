import http from "k6/http";
import { sleep, check } from "k6";

export let options = {
  vus: 1000,          // virtual users
  duration: "10s",   // total test duration
  thresholds: {
    http_req_duration: ["p(95)<200"], // 95% requests must complete <500ms
  },
};

export default function () {
  const payload = JSON.stringify({
    name: "TestUser",
    email: "test@mail.com"
  });

  const params = {
    headers: { "Content-Type": "application/json" },
  };

  const res = http.post("http://localhost:8080/User", payload, params);

  // Assertions for correctness
  check(res, {
    "status is 200": (r) => r.status === 200,
  });

  sleep(1); // keeps load realistic
}
