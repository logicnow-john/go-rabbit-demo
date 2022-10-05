import http from "k6/http";
import encoding from "k6/encoding";
import { sleep } from "k6";
import { Counter } from "k6/metrics";
import { textSummary } from "https://jslib.k6.io/k6-summary/0.0.1/index.js";
import { htmlReport } from "https://raw.githubusercontent.com/benc-uk/k6-reporter/main/dist/bundle.js";
import { generateXrayJUnitXML } from "https://raw.githubusercontent.com/Xray-App/tutorial-js-k6/main/junitXray.js";
import { expect } from "https://jslib.k6.io/k6chaijs/4.3.4.1/index.js";

// Reporting seems to be better with multiple counters, rather than 1 counter + tags
// Format of string is chosen to stay consistent with the k6 values e.g. `http_req_failed`
export const clientErrorCounter = new Counter("client_errors");
export const serverErrorCounter = new Counter("server_errors");

export const options = {
  stages: [
    { duration: "5s", target: 3 },
    { duration: "5s", target: 2 },
    { duration: "5s", target: 1 },
  ],
  thresholds: {
    client_errors: ["count==0"],
    server_errors: ["count==0"],
    http_req_duration: ["p(95) < 2000"],
  },
};

export function handleSummary(data) {
  const reportsDirectory = "k6/reports/";

  return {
    stdout: textSummary(data, { indent: " ", enableColors: true }), // Command Line report
    [`${reportsDirectory}results.json`]: JSON.stringify(data), // JSON report
    [`${reportsDirectory}results.html`]: htmlReport(data), // HTML report
    [`${reportsDirectory}results.xml`]: generateXrayJUnitXML(
      data,
      `${reportsDirectory}results.json`,
      encoding.b64encode(JSON.stringify(data))
    ), // Xray Compatible JUnit report
  };
}

export default function (setupData) {
  const response = sendRequest(setupData);

  if (response.status >= 500) {
    serverErrorCounter.add(1);
  } else if (response.status >= 400) {
    clientErrorCounter.add(1);
  }

  sleep(1);
}

function sendRequest(setupData) {
  const pathValues = ["andy4", "john1", "liam3"];
  const pathValue = pathValues[Math.floor(Math.random() * pathValues.length)];
  const url = `http://localhost:8090/result?type=${pathValue}`;
  // const payload = createAuditEntryPayload();
  // const params = {
  //   headers: {
  //     "Content-Type": "application/json",
  //     Authorization: `Bearer ${setupData.bearerToken}`,
  //   },
  // };

  return http.get(url);
}
